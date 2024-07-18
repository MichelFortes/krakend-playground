package pkg

type Owner struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName,omitempty"`
	ProfileImage string `json:"profileImage,omitempty"`
	Address      string `json:"address,omitempty"`
	Phone        string `json:"phone,omitempty"`
}
