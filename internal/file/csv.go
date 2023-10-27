package file

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

func CsvWriter[T any](content []T, formatter func(T) ([]string, error)) ([]byte, error) {
	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)

	defer writer.Flush()

	for _, entry := range content {
		row, err := formatter(entry)
		if err != nil {
			return nil, err
		}
		if err = writer.Write(row); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}

func CsvReader[T any](path string, header bool, parser func([]string) (*T, error)) ([]T, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	if header {
		_, err = reader.Read()
		if err != nil {
			return nil, err
		}
	}

	result := make([]T, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		ts, err := parser(row)
		if err != nil {
			return nil, err
		}

		result = append(result, *ts)
	}

	return result, nil
}
