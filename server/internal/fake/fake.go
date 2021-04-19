package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/bxcodec/faker/v3"
	"github.com/gofrs/uuid"
)

type FakeCapstoneText struct {
	Title       string `faker:"sentence,unique"`
	Description string `faker:"paragraph"`
	Author      string `faker:"name"`
	Semester    string `faker:"oneof: Fall 2019, Spring 2020, Fall 2020, Spring 2021"`
}

func GenerateFakeCapstones(num int) []db.CreateCapstoneParams {

	list := make([]db.CreateCapstoneParams, 0)
	for i := 0; i < num; i++ {
		capstone := db.CreateCapstoneParams{}
		text := FakeCapstoneText{}
		err := faker.FakeData(&capstone)
		if err != nil {
			panic(err)
		}
		err = faker.FakeData(&text)
		if err != nil {
			panic(err)
		}

		capstone.Title = text.Title
		capstone.Description = text.Description
		capstone.Author = text.Author
		capstone.Semester = text.Semester

		list = append(list, capstone)
	}
	return list

}

func AddFakeCapstonesIfEmpty(queries *db.Queries) {
	capstones, err := queries.GetCapstones(context.Background(), 5)

	if err != nil {
		panic(err)
	}

	if len(capstones) <= 0 {
		fmt.Println("No capstones found. Adding fakes")
		fakes := GenerateFakeCapstones(100)

		for _, fake := range fakes {
			_, err := queries.CreateCapstone(context.Background(), fake)

			if err != nil {
				panic(err)
			}
		}

	}

}

// TODO REMOVE THIS IF ACTUALLY DEPLOYING
func AddTestAdminUserIfEmpty(queries *db.Queries) {
	_, err := queries.GetUserByEmail(context.Background(), "admin@test.com")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}

		addUser := db.CreateUserParams{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Username:  "admin",
			Email:     "admin@test.com",
			Password:  "hunter2",
		}

		_, err = queries.CreateUser(context.Background(), addUser)
		if err != nil {
			panic(err)
		}
	}
}
