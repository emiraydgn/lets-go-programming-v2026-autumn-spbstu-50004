package main

import "fmt"

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
)

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Error: invalid departments count:", err)

		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Error: invalid employees count:", err)

			return
		}

		minTemp := minAllowedTemp
		maxTemp := maxAllowedTemp

		for range employeesCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("Error: invalid temperature request:", err)

				return
			}

			switch operator {
			case ">=":
				if temp > minTemp {
					minTemp = temp
				}
			case "<=":
				if temp < maxTemp {
					maxTemp = temp
				}
			default:
				fmt.Println("Error: invalid operator:", operator)

				return
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
