package tests

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/utils"
	"gotest.tools/v3/assert"
	"testing"
)

func TestUintSliceToStringShouldPass(t *testing.T) {
	var result string
	result = utils.UintSliceToString([]uint{})
	assert.Equal(t, result, "", "Result should be empty string.")
}
