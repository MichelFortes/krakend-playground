package repository

import "vet/doctors/pkg"

func populate() {
	Create(pkg.Doctor{Id: "e215033f-86d0-4957-92c4-0cc9b78c6884", Name: "Dra Márcia"})
	Create(pkg.Doctor{Id: "b4baf67f-4464-4260-b561-dc2d69687f4a", Name: "Dr Aroldo"})
	Create(pkg.Doctor{Id: "d441cb2d-8546-42f0-b79b-30c59ba9261a", Name: "Dra Gabriella"})
}
