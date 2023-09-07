package tests

import (
	"database/sql"
	"fmt"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/database"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
	"time"
)

var Db *sql.DB
var Router *gin.Engine

func TestMain(m *testing.M) {
	err := godotenv.Load("test.env")
	if err != nil {
		log.Fatalf("Error loading test.env file: %s", err)
	}
	var postgresUser = os.Getenv("POSTGRES_USER")
	var postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	var postgresDatabaseName = os.Getenv("POSTGRES_DB")

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=" + postgresPassword,
			"POSTGRES_USER=" + postgresUser,
			"POSTGRES_DB=" + postgresDatabaseName,
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://" + postgresUser + ":" + postgresPassword + "@" + hostAndPort + "/" + postgresDatabaseName + "?sslmode=disable")

	log.Println("Connecting to database on url: ", databaseUrl)

	err = resource.Expire(120)
	if err != nil {
		log.Println("Hard killing container. ")
		return
	} // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		Db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return Db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	//Run tests
	db, err := database.GenerateDatabase(databaseUrl, false)
	Router = app.SetupRouter(db)
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
