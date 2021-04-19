package fake

import (
	"context"
	"fmt"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/bxcodec/faker/v3"
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
