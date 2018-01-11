package db

import (
	"2FAServer/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-pg/pg/orm"

	"github.com/go-pg/pg"
)

// DbContextInterface for DB access.
type DbContextInterface interface {
	CreateSchema() error
	GetKeyByID(keyID int) models.Key
	GetKeysByUserID(userID string) []models.Key
	InsertKey(m models.Key) models.Key
	UpdateKey(keyID int, key string) bool
	DeleteKey(m models.Key) bool
}

type dbContext struct {
	connection *pg.DB
}

// NewDbContext DbContext constructor.
func NewDbContext() DbContextInterface {
	dbc := new(dbContext)
	dbc.connection = initializeDb()

	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if isDebug && err == nil {
		dbc.connection.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}

			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}

	// Create table if not exist.
	dbc.CreateSchema()

	return dbc
}

// Private functions.
func initializeDb() *pg.DB {
	user := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	database := os.Getenv("PG_DATABASE")
	hostname := os.Getenv("PG_HOSTNAME")
	port := os.Getenv("PG_PORT")

	if hostname == "" {
		hostname = "127.0.0.1"
	}

	if port == "" {
		port = "5432"
	}

	fmt.Println("Connecting to " + hostname + ":" + port)

	return pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     hostname + ":" + port,
	})
}

// CreateSchema defines a new of tables based on models.Key struct.
func (dbc *dbContext) CreateSchema() error {
	for _, model := range []interface{}{&models.Key{}} {
		err := dbc.connection.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// GetKeyByID retrieves a model by KeyID
func (dbc *dbContext) GetKeyByID(keyID int) models.Key {
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

// GetKeysByUserID retrieves a list of Keys by UserID
func (dbc *dbContext) GetKeysByUserID(userID string) []models.Key {
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

// InsertKey creates a new Key record in the database.
func (dbc *dbContext) InsertKey(m models.Key) models.Key {
	err := dbc.connection.Insert(&m)
	if err != nil {
		// TODO: Log the error.
		return models.Key{}
	}

	return m
}

// UpdateKey updates a Key records's key value.
func (dbc *dbContext) UpdateKey(keyID int, key string) bool {
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

// DeleteKey removes a Key record from the database.
func (dbc *dbContext) DeleteKey(m models.Key) bool {
	err := dbc.connection.Delete(&m)
	if err != nil {
		// TODO: Log the error.
		return false
	}

	return true
}
