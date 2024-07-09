package repository

import "vet/owners/pkg"

func populate() {
	Create(pkg.Owner{Id: "65464199-544d-41b4-94a5-91d3099a3551", Name: "Carlos Augusto"})
	Create(pkg.Owner{Id: "1dbcec23-0efd-442f-a6d9-54523a69f52b", Name: "Amanda"})
	Create(pkg.Owner{Id: "f5b1e895-a853-4dd4-869f-d2d38f4a72c1", Name: "Gustavo"})
}
