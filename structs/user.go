package structs

import (
	"goLearn/utils"
)

type User struct {
	Name string
	Age  int
}

type Class struct {
	Number int
	Char   string
}

type School struct {
	Students []User
	Classes  []Class
}

func (user User) GetName() string {
	return user.Name
}

func CreateStudentsArray(studentsNames []string) []User {
	users := []User{}

	for i := 0; i < len(studentsNames); i++ {
		user := User{
			Name: studentsNames[i],
			Age:  utils.GenerateRandomInt(100),
		}
		users = append(users, user)
	}
	return users
}
