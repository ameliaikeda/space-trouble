package models

import "time"

type Booking struct {
	ID          string    `json:"id" doc:"Booking ID" example:"01904d33-d262-7531-b71c-05555c63df91" readOnly:"true"`
	FirstName   string    `json:"first_name" doc:"First Name" example:"Jane"`
	LastName    string    `json:"last_name" doc:"Last Name" example:"Doe"`
	Gender      string    `json:"gender" doc:"Gender" example:"female"`
	DateOfBirth time.Time `json:"date_of_birth" doc:"Birthday" example:"1990-01-01" format:"date"`
	LaunchpadID string    `json:"launchpad_id" doc:"Launchpad ID" example:"5e9e4502f509094188566f88"`
	LaunchDate  time.Time `json:"launch_date" doc:"Launch Date" example:"2024-01-01" format:"date"`
	Destination string    `json:"destination" doc:"Destination of flight" example:"mars" enum:"mars,moon,pluto,asteroid_belt,europa,titan,ganymede"`
}
