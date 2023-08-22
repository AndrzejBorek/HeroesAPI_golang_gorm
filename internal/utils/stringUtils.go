package utils

import (
	"strconv"
	"strings"
)

func UintSliceToString(slice []uint) string {
	if len(slice) == 0 {
		return ""
	}
	stringSlice := make([]string, 0)
	for _, v := range slice {
		stringSlice = append(stringSlice, strconv.FormatUint(uint64(v), 10))
	}
	return "[" + strings.Join(stringSlice, ",") + "]"
}

func StringToUint(s string) (uint, error) {
	intValue, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	uintValue := uint(intValue)
	return uintValue, nil
}
