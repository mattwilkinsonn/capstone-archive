package db

import (
	"github.com/bxcodec/faker/v3"
)

func GenerateFakeCapstones(int num) []Capstone {

	list := make([]Capstone, 0)
	for i := 0; i < num; i++ {
		capstone := Capstone{}
		err := faker.FakeData(&capstone)
		if err != nil {
			panic(err)
		}
		list = append(list, capstone)
	}
	return list

}
