package employee

import (
	"challenge/employee/calculate"
	"fmt"
)

type Data struct {
	Name   string
	Office string
	Wage   float64
}

func (f *Data) CalculateWage(e Data) {
	switch f.Office {
	case "DEV":
		e.verifyWageBonusDeveloper()
	case "DBAs", "TESTER":
		e.verifyWageBonusDbaAndTest()
	}

	fmt.Println(e)
}

func (f *Data) verifyWageBonusDeveloper() {
	if f.Wage <= 3000 {
		f.Wage += calculate.ReturnWageWithTwentyPercent(f.Wage)
	} else {
		f.Wage += calculate.ReturnWageWithTenPercent(f.Wage)
	}
}

func (f *Data) verifyWageBonusDbaAndTest() {
	if f.Wage <= 3000 {
		f.Wage += calculate.ReturnWageWithTwentyPercent(f.Wage)
	} else {
		f.Wage += calculate.ReturnWageWithFifteenPercent(f.Wage)
	}
}
