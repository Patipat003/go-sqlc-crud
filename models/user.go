package models

import "github.com/patipat003/go-dbsql-2/db/generate"

type UserData struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email,omitempty"`
	Gender    string `json:"gender,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
	Date      string `json:"date,omitempty"`
}

func NewUserDataFrom(u generate.Mockdatum) UserData {
	return UserData{
		ID:	u.ID,
		FirstName: u.FirstName,
		LastName: u.LastName,
		Email: u.Email.String,
		Gender: u.Gender.String,
		IpAddress: u.IpAddress.String,
		Date: u.Date.Time.Format("2006-01-02"),
	}
}