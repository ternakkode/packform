package csvreader

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func ReadCSVFile(absFileDir string) ([][]string, error) {
	file, err := filepath.Abs(absFileDir)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	return data[1:], nil
}
