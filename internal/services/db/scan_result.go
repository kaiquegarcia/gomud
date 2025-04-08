package db

import "github.com/go-mysql-org/go-mysql/mysql"

func ScanResult[T any](result *mysql.Result) ([]*T, error) {
	defer result.Close()
	output := make([]*T, 0)

	for row := range result.RowNumber() {
		var e T
		if err := scanRow(result.Resultset, row, &e); err != nil {
			return nil, err
		}

		output = append(output, &e)
	}

	return output, nil
}
