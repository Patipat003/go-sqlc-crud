package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/patipat003/go-dbsql-2/db/generate"
	"github.com/patipat003/go-dbsql-2/models"
)

func ReadUser(db *sql.DB) error {
	queries := 	generate.New(db)

	users, err := queries.GetAllUsers(context.Background())
    if err != nil {
        return err
    }

		var userData []models.UserData
		for _, u := range users {
			userData = append(userData, models.NewUserDataFrom(u))
		}

			jsonBytes, err := json.MarshalIndent(userData, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(jsonBytes))
		return nil
}

func ReadUserById(db *sql.DB, id int32) error {
	queries := generate.New(db)

	user, err := queries.GetUserByID(context.Background(), id)
	if err != nil {
		return err
	}

	userData := models.NewUserDataFrom(user)

	jsonBytes, err := json.MarshalIndent(userData, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))
	return nil
}

func AddUser(db *sql.DB) error {
	queries := generate.New(db)

	params := generate.CreateUserParams{
		FirstName: "adsasdasd",
		LastName:  "a23232s",
		Email:     sql.NullString{String: "xasd222x@email.com", Valid: true},
		Gender:    sql.NullString{String: "Male", Valid: true},
		IpAddress: sql.NullString{String: "192.128.1.200", Valid: true},
		Date:      sql.NullTime{Time: time.Now(), Valid: true},
	}

	newUser, err := queries.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("User created:", newUser)

	return nil
}

func DelUser(db *sql.DB, id int32) error {
	queries := generate.New(db)

	user, err := queries.DeleteUser(context.Background(), id)
	if err != nil {
		return err
	}

	fmt.Printf("User deleted ID: %v, Name: %v", user.ID, user.FirstName)

	return nil
	
}

func UpdateDataUser(db *sql.DB, firstName string, lastName string, id int32) error {
	params := generate.UpdateUserParams{
		FirstName: firstName,
		LastName: lastName,
		ID: id,
	}

	queries := generate.New(db)
	
	data, err := queries.UpdateUser(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println(data)
	return nil
}