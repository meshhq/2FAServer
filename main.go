package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	// Create new key
	server.POST("/keys", func(c echo.Context) (err error) {
		nk := new(NewKey)

		if err = c.Bind(nk); err != nil {
			response := JsonResponse{Message: "Invalid request payload\n", TimeStamp: time.Now().Unix()}
			return c.JSON(http.StatusBadRequest, response)
		}

		// Search for key in db.
		// if it doesnt exist, create.

		// Return new value.
		response := JsonResponse{Message: nk.Key + nk.Provider + nk.UserID, TimeStamp: time.Now().Unix()}
		return c.JSON(http.StatusOK, response)
	})

	// Retrieve all keys in storage
	server.GET("/keys", func(c echo.Context) error {
		var userID = c.QueryParam("user_id")
		if userID == "" {
			err := JsonResponse{Message: "'user_id' is missing.", TimeStamp: time.Now().Unix()}
			return c.JSON(http.StatusBadRequest, err)
		}

		// Fetch keys from Db
		keys := [5]string{
			"bdfajfsnkjav",
			"bdfajfsnkjav",
			"bdfajfsnkjav",
			"bdfajfsnkjav",
			"bdfajfsnkjav"}

		return c.JSON(http.StatusOK, keys)
	})

	// Update existing key by key_id
	server.PUT("/keys/:key_id", func(c echo.Context) error {
		keyID := c.Param("key_id")

		// Search for key in db.
		// if it doesnt exist, do nothing

		// Modify key property

		response := JsonResponse{Message: keyID, TimeStamp: time.Now().Unix()}
		return c.JSON(http.StatusOK, response)
	})

	// Delete existing key by key_id
	server.DELETE("/keys/:key_id", func(c echo.Context) error {
		keyID := c.Param("key_id")

		// Search for key in db.
		// if it doesnt exist, do nothing

		// Remove entry

		response := JsonResponse{Message: keyID, TimeStamp: time.Now().Unix()}
		return c.JSON(http.StatusOK, response)
	})

	server.Logger.Fatal(server.Start(":1323"))
}
