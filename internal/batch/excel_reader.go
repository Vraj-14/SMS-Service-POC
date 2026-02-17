package batch

import (
	"errors"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ExcelReader struct{}

func NewExcelReader() *ExcelReader {
	return &ExcelReader{}
}

func (r *ExcelReader) ReadPhoneNumbers(filePath string) ([]string, error) {

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sheetName := file.GetSheetName(0)

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, errors.New("excel file contains no data")
	}

	var numbers []string

	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}

		if len(row) == 0 {
			continue
		}

		number := strings.TrimSpace(row[0])
		if number != "" {
			numbers = append(numbers, number)
		}
	}

	return numbers, nil
}
