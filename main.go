package main

import (
	"net/http"
	"log"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
	"database/sql"
	"fmt"
	"strconv"
)

type CustomerInfo struct {
	ID        int
	Firstname string
	Last_name string
	Email     string
	Age       int
	Gender    string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getAllAccountInfo)
	r.HandleFunc("/id={id}", searchByID)
	r.HandleFunc("/name={name}", searchByName)
	r.HandleFunc("/lname={lname}", searchByLastname)
	r.HandleFunc("/email={email}", searchByEmail)
	r.HandleFunc("/age1={age1}&age2={age2}", searchByAge)
	r.HandleFunc("/gender={gender}", searchByGender)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getAllAccountInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	var CustomerInfoList []CustomerInfo
	tempCustomer := CustomerInfo{}
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
	json.NewEncoder(w).Encode(CustomerInfoList)
}

func searchByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var ID = params["id"]
	var CustomerInfoList []CustomerInfo
	tempCustomer := CustomerInfo{}
	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where id = " + ID)
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
	json.NewEncoder(w).Encode(CustomerInfoList)
}

func searchByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var name = params["name"]
	var CustomerInfoList []CustomerInfo
	tempCustomer := CustomerInfo{}

	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where first_name = \"" + name + "\"")
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
	json.NewEncoder(w).Encode(CustomerInfoList)
}
func searchByLastname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var Lname = params["lname"]
	var CustomerInfoList []CustomerInfo
	tempCustomer := CustomerInfo{}

	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where last_name = \"" + Lname + "\"")
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
	json.NewEncoder(w).Encode(CustomerInfoList)
}

func searchByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var Email = params["email"]
	var CustomerInfoList []CustomerInfo
	tempCustomer := CustomerInfo{}

	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where email like \"%" + Email + "%\"")
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
	json.NewEncoder(w).Encode(CustomerInfoList)
}

func searchByGender(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var Gender = params["gender"]
	var CustomerInfoList []CustomerInfo
	tempCustomer := CustomerInfo{}

	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where gender = \"" + Gender + "\"")
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
	json.NewEncoder(w).Encode(CustomerInfoList)

}

func	searchByAge(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var Age1 = params["age1"]
	var Age2 = params["age2"]
	var CustomerInfoList []CustomerInfo
	tempCustomer := CustomerInfo{}
	database, _ := sql.Open("sqlite3", "user.db")
	rows, err := database.Query("SELECT id, first_name, last_name, email,age,gender FROM users where age >= " + Age1 + " and age <= "+ Age2)
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
	json.NewEncoder(w).Encode(CustomerInfoList)
}
/*find by age*/

