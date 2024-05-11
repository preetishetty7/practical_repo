package main

import "database/sql"

type Service interface {
	GetPersonInfo(personId int) (Person, error)
}

type MyService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &MyService{db: db}
}

func (s *MyService) GetPersonInfo(personId int) (Person, error) {
	var person Person
	query := `
		select p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code
		from person p
		left join phone ph on p.id = ph.person_id 
		left join address_join aj on p.id = aj.person_id
		left join address a on a.id = aj.address_id
		where p.id = ?
	`
	err := s.db.QueryRow(query, personId).Scan(&person.Name, &person.PhoneNumber, &person.City,
		&person.State, &person.Street1, &person.Street2, &person.ZipCode)
	if err != nil {
		return Person{}, err
	}
	return person, nil
}
