package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"go_valid/models"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

func createConnection() *sql.DB {

	err := godotenv.Load(".env")
	var (
		host     = os.Getenv("DB_HOST")
		port, _  = strconv.Atoi(os.Getenv("DB_PORT"))
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	if err != nil {
		log.Fatal("Error loading .env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, er := sql.Open("postgres", psqlInfo)
	if er != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("connection not created %v", err)
	}

	return db
}

func CreateForm(form *models.Form) error {
	db := createConnection()
	defer db.Close()
	if form.Phonenumber != 0 {
		phonenumber := fmt.Sprint(form.Phonenumber)
		match, _ := regexp.MatchString("^[0-9]{10}$", phonenumber)
		if !match {
			return errors.New("not a valid Phone Number")
		}
	}
	if form.Email != "" {
		match, _ := regexp.MatchString(`^(.*)@(.*)([\.])com$`, form.Email)
		if !match {
			return errors.New("not a valid Gmail")
		}
	}

	sqlStatement := `Select email from form WHERE email=$1`
	var email string
	row := db.QueryRow(sqlStatement, form.Email)
	row.Scan(&email)
	
	// fmt.Printf("%v", email)
	if email!=""{
		return errors.New("this email is already entered , Please add unique only! ")
	}
	sqlStatement = `INSERT INTO form(firstname,lastname,email,phonenumber) VALUES($1,$2,$3,$4)`
	_, err := db.Exec(sqlStatement, form.Firstname, form.Lastname, form.Email, form.Phonenumber)

	if err != nil {
		log.Fatalf("Insertion failed ! : %v", err)
	}
	fmt.Println("data inserted!")
	return nil
}

func GetAllData() ([]models.Form, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM form`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute %v", err)
	}
	defer rows.Close()
	var forms []models.Form

	for rows.Next() {
		var form models.Form
		err = rows.Scan(&form.Firstname, &form.Lastname, &form.Email, &form.Phonenumber)
		if err != nil {
			log.Fatalf("retrive data failed : %v", err)
		}
		forms = append(forms, form)
	}
	return forms, err
}
