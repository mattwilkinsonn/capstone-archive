package fake

import (
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/bxcodec/faker/v3"
)

func GenerateFakeCapstones(num int) []db.Capstone {

	list := make([]db.Capstone, 0)
	for i := 0; i < num; i++ {
		capstone := db.Capstone{}
		err := faker.FakeData(&capstone)
		if err != nil {
			panic(err)
		}
		list = append(list, capstone)
	}
	return list

}
