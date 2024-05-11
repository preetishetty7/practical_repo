package main

type Person struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}

type getRequest struct {
	PersonId int `json:"person_id"`
}

type createRequest struct {
	Person Person `json:"person"`
}
