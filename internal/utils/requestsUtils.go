package utils

import (
	"errors"
	"reflect"
)

type IDExtractorFunc func(interface{}) uint

func ValidateItemIDs(existingItems []interface{}, itemIDs []uint, missingErrorMessage string, idExtractor IDExtractorFunc) error {
	if len(itemIDs) > 0 && len(existingItems) != len(itemIDs) {
		missingIDs := make([]uint, 0)
		existingIDs := make(map[uint]bool)
		for _, item := range existingItems {
			existingIDs[idExtractor(item)] = true
		}
		for _, id := range itemIDs {
			if !existingIDs[id] {
				missingIDs = append(missingIDs, id)
			}
		}
		return errors.New(missingErrorMessage + UintSliceToString(missingIDs))
	}
	return nil
}

func SliceToInterface(slice interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("sliceToInterface called with a non-slice type")
	}
	length := s.Len()
	result := make([]interface{}, length)
	for i := 0; i < length; i++ {
		result[i] = s.Index(i).Interface()
	}
	return result, nil
}
