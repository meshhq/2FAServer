package db

import (
	"2FAServer/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
)

type dbContext struct {
	connection *pg.DB
}

func NewDbContext() *dbContext {
	dbc := new(dbContext)
	dbc.connection = initializeDb()

	dbc.connection.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	// err := dbc.createSchema()
	// if err != nil {
	// 	panic(err)
	// }

	return dbc
}

func initializeDb() *pg.DB {
	user := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	database := os.Getenv("PG_DATABASE")

	return pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})
}

func (dbc dbContext) createSchema() error {
	for _, model := range []interface{}{&models.Key{}} {
		err := dbc.connection.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dbc dbContext) GetModel(keyID string) models.Key {
	aKey := models.Key{KeyID: keyID}
	err := dbc.connection.Select(&aKey)
	if err != nil && err != pg.ErrNoRows {
		// TODO: Log the error. Panic for now.
		panic(err)
	}

	if err == pg.ErrNoRows {
		return models.Key{}
	}

	return aKey
}

func (dbc dbContext) GetModels(userID string) []models.Key {
	var res []models.Key
	err := dbc.connection.Model(&models.Key{}).Where("user_id = ?", userID).Select(&res)
	if err != nil && err != pg.ErrNoRows {
		// TODO: Log the error. Panic for now.
		panic(err)
	}

	if res == nil {
		res = []models.Key{}
	}

	return res
}

func (dbc dbContext) InsertModel(m models.Key) bool {
	err := dbc.connection.Insert(&m)
	if err != nil {
		// TODO: Log the error.
		return false
	}

	return true
}

func (dbc dbContext) UpdateModel(keyID, key string) bool {
	aKey := models.Key{
		KeyID: keyID,
		Key:   key,
	}

	res, err := dbc.connection.Model(&aKey).Column("key").Update()
	if err != nil {
		// TODO: Log the error.
		return false
	}

	fmt.Println(res)

	return true
}

func (dbc dbContext) DeleteModel(m models.Key) bool {
	err := dbc.connection.Delete(&m)
	if err != nil {
		// TODO: Log the error.
		return false
	}

	return true
}
