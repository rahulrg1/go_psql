package middleware

import (
	"database/sql"
	// "errors"
	"fmt"
	"go_valid/models"
	"log"
	"net/http"
	"os"
	// "regexp"
	"strconv"

	"encoding/json"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type responseValidation struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

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

func GetAllData(w http.ResponseWriter, r *http.Request) {

	data, err := getAllData()
	if err != nil {
		log.Fatalf("data retrival failed! %v", err)
	}
	json.NewEncoder(w).Encode(data)
}

func CreateForm(w http.ResponseWriter, r *http.Request) {
	var form models.Form
	json.NewDecoder(r.Body).Decode(&form)
	if form.Firstname == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Firstname required"})
		return
	} else if form.Lastname == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Lastname required"})
		return
	}

	err := createForm(&form)
	if err != nil {
		log.Fatalf("Error in createForm : %v", err)
	}
	json.NewEncoder(w).Encode(form)
}

func createForm(form *models.Form) error {
	db := createConnection()
	defer db.Close()
	// if form.Phonenumber != "" {
	// 	match, _ := regexp.MatchString("[0-9]{10}", form.Phonenumber)
	// 	if !match {
	// 		return errors.New("Not a valid Phone Number")
	// 	}
	// }
	// if form.Email != "" {
	// 	match, _ := regexp.MatchString("*@gmail([.])com", form.Email)
	// 	if !match {
	// 		return errors.New("Not a valid email")
	// 	}
	// }

	sqlStatement := `INSERT INTO form(firstname,lastname,email,phonenumber) VALUES($1,$2,$3,$4)`
	// fmt.Println(form.Email)
	// fmt.Println(form.Firstname)
	_,err := db.Exec(sqlStatement, form.Firstname, form.Lastname, form.Email, form.Phonenumber)
	if err != nil {
		log.Fatalf("Insertion failed ! : %v", err)
	}
	fmt.Printf("data inserted")
	return nil
}

func getAllData() ([]models.Form, error) {
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
