package services

import(
	m "../model"
	"database/sql"
	"fmt"
	"strconv"
	)

type CustomerServices struct {

}

func (c CustomerServices) GetAllAccountInfo() []m.CustomerInfo  {
	var CustomerInfoList []m.CustomerInfo
	tempCustomer := m.CustomerInfo{}
	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users ")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	var id int
	var firstname string
	var last_name string
	var email string
	var age int
	var gender string
	for rows.Next() {
		rows.Scan(&id, &firstname, &last_name, &email, &age, &gender)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + last_name + " " + email + " " + strconv.Itoa(age) + " " + gender)
		tempCustomer.ID = id
		tempCustomer.Firstname = firstname
		tempCustomer.Last_name = last_name
		tempCustomer.Email = email
		tempCustomer.Age = age
		tempCustomer.Gender = gender
		CustomerInfoList = append(CustomerInfoList, tempCustomer)
	}
	return CustomerInfoList
}

func (c CustomerServices) GetAccountInfoByID(idParam string) []m.CustomerInfo  {
	var CustomerInfoList []m.CustomerInfo
	tempCustomer := m.CustomerInfo{}
	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where id = " + idParam)
	if err != nil {
		panic(err)
	}
	var id int
	var firstname string
	var last_name string
	var email string
	var age int
	var gender string
	for rows.Next() {
		rows.Scan(&id, &firstname, &last_name, &email, &age, &gender)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + last_name + " " + email + " " + strconv.Itoa(age) + " " + gender)
		tempCustomer.ID = id
		tempCustomer.Firstname = firstname
		tempCustomer.Last_name = last_name
		tempCustomer.Email = email
		tempCustomer.Age = age
		tempCustomer.Gender = gender
		CustomerInfoList = append(CustomerInfoList, tempCustomer)
	}
	return CustomerInfoList
}

func (c CustomerServices) GetAccountInfoByFName(nameParam string) []m.CustomerInfo {
	var CustomerInfoList []m.CustomerInfo
	tempCustomer := m.CustomerInfo{}

	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where first_name = \"" + nameParam + "\"")
	if err != nil {
		panic(err)
	}
	var id int
	var firstname string
	var last_name string
	var email string
	var age int
	var gender string
	for rows.Next() {
		rows.Scan(&id, &firstname, &last_name, &email, &age, &gender)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + last_name + " " + email + " " + strconv.Itoa(age) + " " + gender)
		tempCustomer.ID = id
		tempCustomer.Firstname = firstname
		tempCustomer.Last_name = last_name
		tempCustomer.Email = email
		tempCustomer.Age = age
		tempCustomer.Gender = gender
		CustomerInfoList = append(CustomerInfoList, tempCustomer)
	}
	return CustomerInfoList
}

func (c CustomerServices) GetAccountInfoByLName(lNameParam string) []m.CustomerInfo{
	var CustomerInfoList []m.CustomerInfo
	tempCustomer := m.CustomerInfo{}

	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where last_name = \"" + lNameParam + "\"")
	if err != nil {
		panic(err)
	}
	var id int
	var firstname string
	var last_name string
	var email string
	var age int
	var gender string
	for rows.Next() {
		rows.Scan(&id, &firstname, &last_name, &email, &age, &gender)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + last_name + " " + email + " " + strconv.Itoa(age) + " " + gender)
		tempCustomer.ID = id
		tempCustomer.Firstname = firstname
		tempCustomer.Last_name = last_name
		tempCustomer.Email = email
		tempCustomer.Age = age
		tempCustomer.Gender = gender
		CustomerInfoList = append(CustomerInfoList, tempCustomer)
	}
	return CustomerInfoList
}

func (c CustomerServices)GetAccountInfoByEmail (emailParam string) []m.CustomerInfo  {
	var CustomerInfoList []m.CustomerInfo
	tempCustomer := m.CustomerInfo{}

	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where email like \"%" + emailParam + "%\"")
	if err != nil {
		panic(err)
	}
	var id int
	var firstname string
	var last_name string
	var email string
	var age int
	var gender string
	for rows.Next() {
		rows.Scan(&id, &firstname, &last_name, &email, &age, &gender)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + last_name + " " + email + " " + strconv.Itoa(age) + " " + gender)
		tempCustomer.ID = id
		tempCustomer.Firstname = firstname
		tempCustomer.Last_name = last_name
		tempCustomer.Email = email
		tempCustomer.Age = age
		tempCustomer.Gender = gender
		CustomerInfoList = append(CustomerInfoList, tempCustomer)
	}
	return CustomerInfoList
}

func (c CustomerServices) GetAccountInfoByGender(genderParam string) []m.CustomerInfo{
	var CustomerInfoList []m.CustomerInfo
	tempCustomer := m.CustomerInfo{}

	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where gender = \"" + genderParam + "\"")
	if err != nil {
		panic(err)
	}
	var id int
	var firstname string
	var last_name string
	var email string
	var age int
	var gender string
	for rows.Next() {
		rows.Scan(&id, &firstname, &last_name, &email, &age, &gender)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + last_name + " " + email + " " + strconv.Itoa(age) + " " + gender)
		tempCustomer.ID = id
		tempCustomer.Firstname = firstname
		tempCustomer.Last_name = last_name
		tempCustomer.Email = email
		tempCustomer.Age = age
		tempCustomer.Gender = gender
		CustomerInfoList = append(CustomerInfoList, tempCustomer)
	}
	return CustomerInfoList
}

func (c CustomerServices) GetAccountInfoByAge(ageParam1 string, ageParam2 string) []m.CustomerInfo  {
	var CustomerInfoList []m.CustomerInfo
	tempCustomer := m.CustomerInfo{}
	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where age >= " + ageParam1 + " and age <= " + ageParam2)
	if err != nil {
		panic(err)
	}
	var id int
	var firstname string
	var last_name string
	var email string
	var age int
	var gender string
	for rows.Next() {
		rows.Scan(&id, &firstname, &last_name, &email, &age, &gender)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + last_name + " " + email + " " + strconv.Itoa(age) + " " + gender)
		tempCustomer.ID = id
		tempCustomer.Firstname = firstname
		tempCustomer.Last_name = last_name
		tempCustomer.Email = email
		tempCustomer.Age = age
		tempCustomer.Gender = gender
		CustomerInfoList = append(CustomerInfoList, tempCustomer)
	}
	return CustomerInfoList
}