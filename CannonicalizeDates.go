package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)


func monthHandler(DateString []string) map[string]string{

	DateMap := make(map[string]string)

	if len(DateString[0]) < 4{
		DateMap["month"] = DateString[0]
		DateMap["day"] = DateString[1]
		DateMap["year"] = DateString[2]
	}else if len(DateString[0]) == 4{
		DateMap["year"] = DateString[0]
		DateMap["month"] = DateString[1]
		DateMap["day"] = DateString[2]
	}

	return DateMap
}

func getDelimeter(dateInputs string) map[string]string  {
	delimeterMap := make(map[string]string)

	delimeterList := []string{"/"," ", "-"}

	for i:=0; i < len(delimeterList) ; i++  {

		if strings.Contains(dateInputs, delimeterList[i]){
			delimeterMap["delimeter"] = delimeterList[i]

		}

	}
	return delimeterMap

}


func main() {

//	someString := "2015/5/14"

	datesFile, err := os.Open("bad_dates.txt")
	if err != nil {
		panic("Cannot open file!")
	}
	defer datesFile.Close()

	scanner := bufio.NewScanner(datesFile)

	for scanner.Scan(){
		delimeter := getDelimeter(strings.TrimSpace(scanner.Text()))

		splitString := strings.Split(strings.TrimSpace(scanner.Text()), delimeter["delimeter"])
		monthHandlerValue := monthHandler(splitString)

		if len(monthHandlerValue["year"]) < 4{
			monthHandlerValue["year"] = "20"+ monthHandlerValue["year"]
		}
		if len(monthHandlerValue["day"]) < 2 {
			monthHandlerValue["day"] = "0"+monthHandlerValue["day"]
		}
		if len(monthHandlerValue["month"]) < 2{
			monthHandlerValue["month"] = "0"+monthHandlerValue["month"]
		}

		canonicalDate := monthHandlerValue["year"] + "-" + monthHandlerValue["month"] + "-" + monthHandlerValue["day"]


		fmt.Println(canonicalDate)
	}



}