package database

import (
	"SB-GO-BATCH-65-VINCENT-2/model"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "myB@nk88"
	dbname   = "gosanber"
)

var (
	DB  *sql.DB
	err error
)

func ConnectDB() {

	pSqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", pSqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected database!")
}

func GetBioskop() {

	var res = []model.Bioskop{}
	// generate ID disini

	sqlStatement := `SELECT * FROM Bioskop`

	rows, err := DB.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var bioskop model.Bioskop
		err = rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
		if err != nil {
			panic(err)
		}
		res = append(res, bioskop)
	}
	fmt.Println(res)
}

func AddBioskop(bioskop *model.Bioskop) error {

	fmt.Println(bioskop)

	sqlStatement := `
	INSERT INTO Bioskop (Nama, Lokasi, Rating) 
	VALUES ($1, $2, $3) 
	RETURNING ID`

	// Insert and retrieve the ID of the new Bioskop
	err := DB.QueryRow(sqlStatement, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskop.ID)
	if err != nil {
		return fmt.Errorf("could not insert bioskop: %v", err)
	}

	// Successfully inserted, return nil error
	fmt.Printf("Bioskop inserted with ID: %d\n", bioskop.ID)
	return nil
}
