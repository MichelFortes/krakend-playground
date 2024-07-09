package pkg

type Pet struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Breed    string `json:"breed"`
	OwnerId  string `json:"ownerId"`
	DoctorId string `json:"doctorId"`
}
