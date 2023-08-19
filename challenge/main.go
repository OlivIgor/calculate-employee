package main

import (
	"challenge/employee"
)

func main() {
	employee := employee.Data{
		Name:   "igor",
		Office: "TESTER",
		Wage:   5500,
	}

	employee.CalculateWage(employee)
}
