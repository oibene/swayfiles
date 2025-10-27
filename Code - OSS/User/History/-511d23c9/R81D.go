package main

type Customer struct {
	id_customer string `json:id_customer`
	first_name string `json:first_name`
	last_name string `json:last_name`
	email string `json:email`
	password string `json:password`
	img_URL string `json:img_URL`
}

var customers = []customer(
	{id: '1', first_name: 'Icaro', last_name: 'O', email: 'icaro@email.com', senha:'123teste123', null},
	{id: '2', first_name: 'T', last_name: 'Nagata', email: 'nagata@email.com', senha:'123teste453', null},
)

func getCustomers(c *gin.Context) {
	
}