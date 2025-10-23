package types

import "time"

type UserData struct {
	Id		     string     `json:"id"`
	Name         string	    `json:"name"`
	Age          int		`json:"age"`
	Email        string	    `json:"email"`
	Mobile       string	    `json:"mobile"`
	Organization string	    `json:"organization"`
	JoiningDate  time.Time  `json:"joining_date"`
	CreatedAt    time.Time	`json:"created_at"`
	UpdatedAt    time.Time	`json:"updated_at"`
}