package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	isClose     bool
	Rating      float64
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		return User{}
	}

	if age <= 0 || age >= 100 {
		return User{}
	}

	if phoneNumber == "" {
		return User{}
	}

	if rating < 0.0 || rating > 10.0 {
		return User{}
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		isClose:     isClose,
		Rating:      rating,
	}

}

func (u *User) changeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) changeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) changePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	}
}

func (u *User) closeAccount() {
	u.isClose = true
}

func (u *User) openAccount() {
	u.isClose = false
}

func (u *User) ratingUP(newRating float64) {
	if u.Rating+newRating <= 10.0 {
		u.Rating += newRating
	}
}

func (u *User) ratingDown(newRating float64) {
	if u.Rating-newRating >= 0.0 {
		u.Rating -= newRating
	}
}

func main() {
	user := NewUser(
		"Mary",
		50,
		"689954",
		false,
		7,
	)

	fmt.Println(user)
}
