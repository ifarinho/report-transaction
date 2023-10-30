package file

import (
	"encoding/csv"
	"io"
)

func CsvReader[T any](path string, id uint, getter func(string) (*csv.Reader, error), parser func([]string, uint) (*T, error)) ([]T, error) {
	reader, err := getter(path)
	if err != nil {
		return nil, err
	}

	if _, err = reader.Read(); err != nil { // skip the header
		return nil, err
	}

	var result []T

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		ts, err := parser(row, id)
		if err != nil {
			return nil, err
		}

		result = append(result, *ts)
	}

	return result, nil
}
