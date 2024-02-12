package main

import (
	"fmt"
	"goLearn/structs"
	"goLearn/utils"
)

func main() {

	var school = structs.School{
		Students: structs.CreateStudentsArray(utils.GetNames()),
	}
	fmt.Print(school.Students)
}
