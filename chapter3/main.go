package main

import (
	"fmt"
	"time"
)

// User is a User interface
type User interface {
	PrintName()
	PrintDetails()
}

// Person is a Person struct
type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

// PrintName prints a Person's name
func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

// PrintDetails is A person method
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s]\n", p.Dob.String(), p.Email,
		p.Location)
}

// ChangeLocation is a person method with pointer receiver that change's the person's location
func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}

// Admin is an Admin struct that embeds the Person struct
type Admin struct {
	Person
	Roles []string
}

// PrintDetails overrides Person's PrintDetails for Admins
func (a Admin) PrintDetails() {
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

// Member is a Member struct that embeds the Person struct
type Member struct {
	Person
	Skills []string
}

// PrintDetails overrides PrintDetails
func (m Member) PrintDetails() {
	//Call person PrintDetails
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}

// Team describes a Team of Users
type Team struct {
	Name, Description string
	Users             []User
}

// GetTeamDetails outputs the team's details
func (t Team) GetTeamDetails() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of the team members:")
	for _, v := range t.Users {
		v.PrintName()
		v.PrintDetails()
	}
}

func main() {
	fmt.Println("Hello, Web Development with Go\nChapter 3")

	alex := Admin{
		Person{
			"Alex",
			"John",
			time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
			"alex@email.com",
			"New York",
		},
		[]string{"Manage Team", "Manage Tasks"},
	}
	shiju := Member{
		Person{
			"Shiju",
			"Varghese",
			time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			"shiju@email.com",
			"Kochi",
		},
		[]string{"Go", "Docker", "Kubernets"},
	}
	chris := Member{
		Person{
			"Chris",
			"Martin",
			time.Date(1978, time.March, 15, 0, 0, 0, 0, time.UTC),
			"chris@email.com",
			"Santa Clara",
		},
		[]string{"Go", "Docker"},
	}
	users := []User{alex, shiju, chris}

	team := Team{
		"Go",
		"Golang CoE esto es una descriotion",
		users,
	}
	team.GetTeamDetails()
}
