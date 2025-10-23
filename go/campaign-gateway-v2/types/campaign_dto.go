package types

import "time"

type Campaign struct {
	Id		  string `json:"id"`
	Name	  string `json:"name"`
	Cron 	  string `json:"cron"`
	Filter    string `json:"filter"`
	Enabled   bool   `json:"enabled"`
	LastRun   *time.Time `json:"last_run"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}