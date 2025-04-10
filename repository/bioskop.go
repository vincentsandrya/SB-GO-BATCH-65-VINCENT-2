package repository

import (
	"SB-GO-BATCH-65-VINCENT-2/model"
	"database/sql"
)

func GetAllBioskop(db *sql.DB) (result []model.Bioskop, err error) {
	sql := "SELECT * FROM Bioskop"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var Bioskop model.Bioskop

		err = rows.Scan(&Bioskop.ID, &Bioskop.Nama, &Bioskop.Lokasi, &Bioskop.Rating)
		if err != nil {
			return
		}

		result = append(result, Bioskop)
	}

	return
}

func GetBioskopById(db *sql.DB, id int) (result model.Bioskop, err error) {
	sql := "SELECT * FROM Bioskop where id = $1"

	err = db.QueryRow(sql, id).Scan(
		&result.ID,
		&result.Nama,
		&result.Lokasi,
		&result.Rating,
	)
	if err != nil {
		return result, err
	}

	return result, nil
}

func InsertBioskop(db *sql.DB, Bioskop model.Bioskop) (err error) {
	sql := "INSERT INTO Bioskop(nama, lokasi, rating) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, Bioskop.Nama, Bioskop.Lokasi, Bioskop.Rating)

	return errs.Err()
}

func UpdateBioskop(db *sql.DB, Bioskop model.Bioskop) (err error) {
	sql := "UPDATE Bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4"

	errs := db.QueryRow(sql, Bioskop.Nama, Bioskop.Lokasi, Bioskop.Rating, Bioskop.ID)

	return errs.Err()
}

func DeleteBioskop(db *sql.DB, Bioskop model.Bioskop) (err error) {
	sql := "DELETE FROM Bioskop WHERE id = $1"

	errs := db.QueryRow(sql, Bioskop.ID)
	return errs.Err()
}
