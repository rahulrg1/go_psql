package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	// "encoding/json"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	er := godotenv.Load(".env")

	const (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	if er != nil {
		log.Fatal("Error loading .env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("sucessfully connected to postgres")

	return db
}

// func CreateStock(w http.ResponseWriter,r *http.Request){
// 	var stock models.Stock
// 	json.Decoder(r.Body).Decode(&stock)
// 	insertID := insertStock(stock)

// 	res := response{
// 		ID: insertID,
// 		Message: "stock created",
// 	}
// 	json.NewEncoder(w).Encode(res)
// }

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi{params["id"]}
	stock, err := getStock(int64(id))

}

// func GetAllStock(w http.ResponseWriter,r *http.Request){

// }

// func insertStock() int64{

// }

// func getStock(id int64) int64{

// }

// func updateStock(id int64,stock models.Stock) int64{

// }
