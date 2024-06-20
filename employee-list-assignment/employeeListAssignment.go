package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/Pallinder/go-randomdata"
)

// Forms struct for employee information
type Employees struct {
	employeeID int
	name       string
	age        int
	city       string
	salary     float64
}

func main() {
	employeeList := []Employees{}

	// Forms array of struct objects of employee information for employees
	for i := 1; i <= 1000; i++ {
		employeeList = append(employeeList, Employees{i, randomdata.FullName(randomdata.RandomGender), randomdata.Number(25, 66), randomdata.City(), float64(randomdata.Number(25000, 200001))})
	}
	// Call for task 1
	sortByID(employeeList)
	fmt.Println(employeeList)
	// Call for task 1
	sortByAge(employeeList)
	fmt.Println(employeeList)
	// Call for task 3
	fmt.Println(top5(employeeList))
	// Call for task 4
	fmt.Println(topPaidCity(employeeList))
	// Call for task 5
	fmt.Println(averageSalaryForBerkhamsted(employeeList))
	// Call for task 6
	fmt.Println(moreThanTwoVowels(employeeList))
	// Call for task 7
	fmt.Println(cityEmployeeCount(employeeList))
	// Call for task 8
	fmt.Println(between45To55(employeeList))
	// Call for task 9
	fmt.Println(citySalaryPercentages(employeeList))
}

// Task 1
func sortByID(employeeList []Employees) {
	slices.SortStableFunc(employeeList, func(a, b Employees) int {
		return cmp.Compare(a.employeeID, b.employeeID)
	})
}

// Task 2
func sortByAge(employeeList []Employees) {
	slices.SortStableFunc(employeeList, func(a, b Employees) int {
		return cmp.Compare(a.name, b.name)
	})

	slices.SortStableFunc(employeeList, func(a, b Employees) int {
		return cmp.Compare(a.age, b.age)
	})
}

// Task 3
func top5(employeeList []Employees) map[string]int {
	top5List := map[string]int{}
	slices.SortStableFunc(employeeList, func(a, b Employees) int {
		return cmp.Compare(a.salary, b.salary)
	})
	for i := 1; i <= 5; i++ {
		top5List[employeeList[len(employeeList)-i].name] = employeeList[len(employeeList)-i].employeeID
	}
	return top5List
}

// Task 4
func topPaidCity(employeeList []Employees) (string, float64) {
	var topCity string
	var topCityAverageSalary float64
	cityAverageSalaries := map[string]float64{}
	cityCount := map[string]float64{}
	for _, value := range employeeList {
		cityAverageSalaries[value.city] += value.salary
		cityCount[value.city]++
	}
	for city, count := range cityCount {
		cityAverageSalaries[city] = cityAverageSalaries[city] / count
	}
	for city, averageCitySalary := range cityAverageSalaries {
		if averageCitySalary > topCityAverageSalary {
			topCityAverageSalary = averageCitySalary
			topCity = city
		}
	}
	topCityAverageSalary = math.Round(topCityAverageSalary*100) / 100
	return topCity, topCityAverageSalary
}

// Task 5
func averageSalaryForBerkhamsted(employeeList []Employees) float64 {
	var totalSalaryForBerkhamsted float64
	var countForBerkhamsted float64
	var averageSalaryForBerkhamsted float64
	for _, value := range employeeList {
		if value.city == "Berkhamsted" {
			totalSalaryForBerkhamsted += value.salary
			countForBerkhamsted++
		}
	}
	averageSalaryForBerkhamsted = math.Round((totalSalaryForBerkhamsted/countForBerkhamsted)*100) / 100
	return averageSalaryForBerkhamsted
}

// Task 6
func moreThanTwoVowels(employeeList []Employees) int {
	var employeeCount int
	for _, value := range employeeList {
		if strings.Count(value.name, "a")+strings.Count(value.name, "e")+strings.Count(value.name, "i")+strings.Count(value.name, "o")+strings.Count(value.name, "u")+strings.Count(value.name, "A")+strings.Count(value.name, "E")+strings.Count(value.name, "I")+strings.Count(value.name, "O")+strings.Count(value.name, "U") > 2 {
			employeeCount++
		}
	}
	return employeeCount
}

// Task 7
func cityEmployeeCount(employeeList []Employees) map[string]int {
	cityEmployeeCount := map[string]int{}
	for _, value := range employeeList {
		cityEmployeeCount[value.city]++
	}
	return cityEmployeeCount
}

// Task 8
func between45To55(employeeList []Employees) int {
	countOfEmployees := 0
	for _, value := range employeeList {
		if value.age >= 45 && value.age <= 55 {
			countOfEmployees++
		}
	}
	return countOfEmployees
}

// Task 9
func citySalaryPercentages(employeeList []Employees) map[string]float64 {
	citySalaryPercentages := map[string]float64{}
	var totalSalary float64
	for _, value := range employeeList {
		totalSalary += value.salary
		citySalaryPercentages[value.city] += value.salary
	}
	for city := range citySalaryPercentages {
		citySalaryPercentages[city] = math.Round((citySalaryPercentages[city]/totalSalary)*10000) / 100
	}
	return citySalaryPercentages
}
