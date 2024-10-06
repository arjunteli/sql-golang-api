package main

type PersonInfo struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
	Age         int    `json:"age"`
}

type CreatePersonRequest struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	City        string `json:"city" binding:"required"`
	State       string `json:"state" binding:"required"`
	Street1     string `json:"street1" binding:"required"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code" binding:"required"`
	Age         int    `json:"age" binding:"required"`
}
