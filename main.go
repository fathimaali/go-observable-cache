package main

import (
	"fmt"
	"os"
	"strings"
)

type Role struct {
	Level string
}

type EmployeeDetails struct {
	Name      string
	Age       int
	RoleLevel Role
}

type EmployeeID = string

func main() {

	fileName := os.Args[1]

	for {
		_, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist, do you want to try again? Enter filename: ")
			fmt.Scanln(&fileName)
		}
		if err == nil {
			break
		}
	}

	myMap, newlocalCache := csvParseIntoMap(fileName)
	fmt.Println("Contents of the resultant map are: ")
	for eID, eDetails := range myMap {
		fmt.Println(eID, eDetails)
	}
	var searchChoice string
	var inputEmployeeID EmployeeID
	for {
		// check if more values need to be entered, if not exit
		fmt.Println("Do you want to search the cache? yes / no ")
		fmt.Scanln(&searchChoice)

		if strings.ToLower(searchChoice) != "yes" {
			break
		}
		fmt.Println("Enter the Employee ID that you want to search the details for, we'll see if it's cached: ")
		fmt.Scanln(&inputEmployeeID)
		found, err := newlocalCache.Get(inputEmployeeID)
		if err == false {
			fmt.Println("Hmm, looks like we don't have results for it.")
		} else {
			fmt.Println("Employee details found! Employee ID : " + inputEmployeeID + " " + found + "!")
		}
	}
}
