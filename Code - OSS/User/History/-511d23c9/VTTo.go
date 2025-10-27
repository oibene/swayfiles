package main

type Customer struct {
	id_customer string `json:id_customer`
	first_name string `json:first_name`
	last_name string `json:last_name`
	email string `json:email`
	password string `json:password`
	img_URL string `json:img_URL`
}