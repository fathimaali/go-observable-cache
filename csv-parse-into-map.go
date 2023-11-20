package main

import (
	"encoding/csv"
	"go-observable-cache/observablecache"
	"log"
	"os"
	"strconv"
)

func csvParseIntoMap(fileName string) (map[EmployeeID]EmployeeDetails, observablecache.LocalCache) {

	file, err := os.Open(fileName)
	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	// Closes the file
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	myMap := map[EmployeeID]EmployeeDetails{}
	localC := observablecache.New()
	localC.Init()
	for _, eachrecord := range records {
		if len(eachrecord) != 4 {
			continue
		}
		input_age, _ := strconv.Atoi(eachrecord[2])
		myMap[eachrecord[0]] = EmployeeDetails{
			Name:      eachrecord[1],
			Age:       input_age,
			RoleLevel: Role{Level: eachrecord[3]},
		}
		localC.Set(eachrecord[0], eachrecord[1])
	}
	return myMap, localC
}
