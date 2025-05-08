package lib

import "time"

type Password struct {
	ID	uint
	Name string
	Password string
	CreatedAt time.Time
}

type CreateNewPasswordManual struct {
	ID	uint
	Name	string
	Password	string
	RepeatPassword	string
	CreatedAt 	time.Time
}

type CreateNewPasswordAutomatic struct {
	ID	uint
	Name	string
	Password	string
	CreatedAt 	time.Time
}