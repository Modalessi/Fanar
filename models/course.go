package models

import "github.com/google/uuid"

type Course struct {
	ID           *uuid.UUID
	Title        string
	Code         string
	Description  string
	CreditHours  int
	ContactHours int
}

func NewCourse(title string, code string, creditHours int, contactHours int) *Course {
	return &Course{
		Title:        title,
		Code:         code,
		CreditHours:  creditHours,
		ContactHours: contactHours,
	}
}

func (c *Course) SetDescription(description string) {
	c.Description = description
}
