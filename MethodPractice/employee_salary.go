package MethodPractice

import (
	"fmt"
	"math"
)

type EmployeeSalary struct {
	Salary int
}

func CheckSalary() {
	salary := &EmployeeSalary{
		Salary: 13000,
	}
	increaseRate := 0.05
	fmt.Printf("基本工资:%d\n", salary.Salary)
	fmt.Printf("工资涨幅比例:%.2f\n", increaseRate)
	newSalary := salary.AddSalary(increaseRate)
	fmt.Printf("增加后的薪水:%d\n", newSalary)
}

func (s *EmployeeSalary) AddSalary(percent float64) int {
	salaryFloat := float64(s.Salary)
	newSalaryFloat := salaryFloat * (1.0 + percent)
	updatedSalary := int(math.Round(newSalaryFloat))
	s.Salary = updatedSalary
	return s.Salary
}
