package controllers

import (
	"SB-GO-BATCH-65-VINCENT-2/database"
	"SB-GO-BATCH-65-VINCENT-2/model"
	"SB-GO-BATCH-65-VINCENT-2/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func validateDataBioskop(Bioskop model.Bioskop) (string, error) {
	if Bioskop.Nama == "" {
		return "Nama tidak boleh kosong", errors.New("error")
	}

	return "data valid", nil
}

func GetAllBioskop(c *gin.Context) {
	var (
		result gin.H
	)

	Bioskop, err := repository.GetAllBioskop(database.DbConnection)

	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		}

		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Success Getting Data!",
			"result":  Bioskop,
		}

		c.JSON(http.StatusOK, result)
	}

}

func GetBioskopById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var (
		result gin.H
	)

	Bioskop, err := repository.GetBioskopById(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		}

		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Success Getting Data!",
			"result":  Bioskop,
		}

		c.JSON(http.StatusOK, result)
	}
}

func InsertBioskop(c *gin.Context) {
	var Bioskop model.Bioskop

	err := c.BindJSON(&Bioskop)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
		panic(err)
	}

	message, err := validateDataBioskop(Bioskop)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": message,
			"result":  nil,
		})
		panic(err)
	}

	err = repository.InsertBioskop(database.DbConnection, Bioskop)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Success Insert Data!",
			"result":  Bioskop,
		})
	}
}

func UpdateBioskop(c *gin.Context) {
	var Bioskop model.Bioskop
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&Bioskop)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
		panic(err)
	}

	message, err := validateDataBioskop(Bioskop)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": message,
			"result":  nil,
		})

		panic(err)
	}

	Bioskop.ID = id

	err = repository.UpdateBioskop(database.DbConnection, Bioskop)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Success Update Data!",
			"result":  Bioskop,
		})
	}
}

func DeleteBioskop(c *gin.Context) {
	var Bioskop model.Bioskop
	id, _ := strconv.Atoi(c.Param("id"))

	Bioskop.ID = id
	err := repository.DeleteBioskop(database.DbConnection, Bioskop)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Success Delete Data!",
			"result":  Bioskop,
		})
	}
}
