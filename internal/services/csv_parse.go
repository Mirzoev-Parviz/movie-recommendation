package services

import (
	"encoding/csv"
	"os"
)

type CSV_Parser interface {
	ReadCSV(filename string) ([][]string, error)
}

type CSV_ParserService struct {
}

func NewCSV_Parser() *CSV_ParserService {
	return &CSV_ParserService{}
}

func (c *CSV_ParserService) ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
