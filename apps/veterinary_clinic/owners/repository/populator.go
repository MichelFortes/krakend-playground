package repository

import "vet/owners/pkg"

func populate() {

	Create(pkg.Owner{
		Id:           "65464199-544d-41b4-94a5-91d3099a3551",
		FirstName:    "Carlos",
		LastName:     "Augusto Barbosa",
		Phone:        "(22) 2222 2222",
		Address:      "Rua Xxxxx, xxxxx, xxxxxx",
		ProfileImage: "https://cdn1.iconfinder.com/data/icons/user-pictures/101/malecostume-512.png",
	})

	Create(pkg.Owner{
		Id:           "1dbcec23-0efd-442f-a6d9-54523a69f52b",
		FirstName:    "Amanda",
		LastName:     "Santos da Silva",
		Phone:        "(11) 1111 1111",
		Address:      "Rua Zzzzz, zzzzzz, zzzzzzz",
		ProfileImage: "https://cdn1.iconfinder.com/data/icons/user-pictures/100/female1-512.png",
	})

	Create(pkg.Owner{
		Id:        "f5b1e895-a853-4dd4-869f-d2d38f4a72c1",
		FirstName: "Pedro",
		LastName:  "Cortes Fonseca",
		Phone:     "(55) 5555 5555",
		Address:   "Rua Yyyyy, yyyyy, yyyyyy",
	})

}
