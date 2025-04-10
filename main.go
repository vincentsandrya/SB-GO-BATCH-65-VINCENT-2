package main

import (
	"SB-GO-BATCH-65-VINCENT-2/controllers"
	"SB-GO-BATCH-65-VINCENT-2/database"
	"SB-GO-BATCH-65-VINCENT-2/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	// database.GetBioskop()

	// routing
	router := gin.Default()

	// http.Handle("/bioskop", http.HandlerFunc(HandleBioskop))
	router.GET("/bioskop", controllers.GetAllBioskop)
	router.GET("//bioskop/:id", controllers.GetBioskopById)
	router.POST("/bioskop", controllers.InsertBioskop)
	router.PUT("/bioskop/:id", controllers.UpdateBioskop)
	router.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	router.Run(":8080")

	// jalankan server
	fmt.Println("server running at http://localhost:8080")

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
