package main

import (
	"SB-GO-BATCH-65-VINCENT-2/database"
	"SB-GO-BATCH-65-VINCENT-2/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	database.ConnectDB()

	database.GetBioskop()

	// konfigurasi server
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/bioskop", http.HandlerFunc(HandleBioskop))

	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()

}

func HandleBioskop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bioskop model.Bioskop

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)

			if err := decodeJSON.Decode(&bioskop); err != nil {
				log.Fatal(err)
			}

		} else {

			var strId string = r.PostFormValue("ID")
			var strRating string = r.PostFormValue("Rating")
			Id, _ := strconv.Atoi(strId)
			Rating, _ := strconv.ParseFloat(strRating, 64)

			bioskop = model.Bioskop{
				ID:     Id,
				Nama:   r.PostFormValue("Nama"),
				Lokasi: r.PostFormValue("Lokasi"),
				Rating: Rating,
			}
		}

		if bioskop.Nama == "" {
			w.Write([]byte("<h1>Nama tidak boleh kosong</h1>"))
		} else if bioskop.Lokasi == "" {
			w.Write([]byte("<h1>Lokasi tidak boleh kosong</h1>"))
		} else {
			database.AddBioskop(&bioskop)
			w.Write([]byte("<h1>Anda Berhasil Memasukan data Bioskop</h1>"))
			return
		}

		http.Error(w, "ERROR....", http.StatusNotFound)
	}
}
