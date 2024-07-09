package repository

import "vet/pets/pkg"

func populate() {
	Create(pkg.Pet{
		Id:       "a15c72f4-1145-45ad-a544-6ea6f5fc3bd0",
		Name:     "Pedrita",
		Breed:    "Beagle",
		OwnerId:  "65464199-544d-41b4-94a5-91d3099a3551",
		DoctorId: "e215033f-86d0-4957-92c4-0cc9b78c6884",
	})
	Create(pkg.Pet{
		Id:       "82b41dbf-cdd6-46a7-ae02-4db76702fa80",
		Name:     "Bebeto",
		Breed:    "Schnauzer",
		OwnerId:  "1dbcec23-0efd-442f-a6d9-54523a69f52b",
		DoctorId: "b4baf67f-4464-4260-b561-dc2d69687f4a",
	})
	Create(pkg.Pet{
		Id:       "68f4961c-cb5a-4e64-91f0-25a69513b9cc",
		Name:     "Totó",
		Breed:    "Caramelo",
		OwnerId:  "f5b1e895-a853-4dd4-869f-d2d38f4a72c1",
		DoctorId: "d441cb2d-8546-42f0-b79b-30c59ba9261a",
	})
}
