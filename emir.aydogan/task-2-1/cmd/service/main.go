package main

import "fmt"

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
)

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			return
		}

		for range employeesCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				return
			}
		}
	}
}
