package main

import (
	"fmt"
)

// Interface para definir o comportamento de cálculo de salário para diferentes cargos
type Employee interface {
	CalculateSalary(salary float64) float64
}

// Struct para representar um funcionário
type Staff struct {
	Name   string
	Salary float64
	Role   string
}

// Struct para representar um Desenvolvedor
type Developer struct {
	Staff
}

func (d Developer) CalculateSalary(salary float64) float64 {
	if salary > 3000 {
		return salary * 0.1
	}
	return salary * 0.2
}

// Struct para representar um DBA
type DBA struct {
	Staff
}

func (d DBA) CalculateSalary(salary float64) float64 {
	if salary > 3000 {
		return salary * 0.15
	}
	return salary * 0.2
}

// Struct para representar um Testador
type Tester struct {
	Staff
}

func (t Tester) CalculateSalary(salary float64) float64 {
	if salary > 3000 {
		return salary * 0.15
	}
	return salary * 0.2
}

func main() {
	dev := Developer{Staff{"Alice", 4000, "DEV"}}
	dba := DBA{Staff{"Bob", 2500, "DBA"}}
	tester := Tester{Staff{"Charlie", 3200, "TESTER"}}

	employees := []Employee{dev, dba, tester}

	for _, employee := range employees {
		calculatedSalary := employee.CalculateSalary(employee.Salary)
		fmt.Printf("%s - %s: Salário Calculado: R$%.2f\n", employee.Name, employee.Role, calculatedSalary)
	}
}
