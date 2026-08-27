package main

import "fmt"

func main() {
	fmt.Println("Methods!")
	user := User{
		Name:   "Rahul",
		Email:  "rahul@gmail.com",
		Status: true,
		Age:    22,
	}

	fmt.Println("User details are:", user)
	user.GetStatus()
	user.NewEmail()
	fmt.Println("User details are:", user)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User Status:", u.Status)

}

// it dont cage the origirna it sednthe copy to the output here come the pointert things

func (u User) NewEmail() {
	u.Email = "newemail@gmail.com"
	fmt.Println("User Email updated to:", u.Email)
}
