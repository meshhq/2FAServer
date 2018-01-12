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
type ContextInterface interface {
	CreateSchema() error
	GetModel(model interface{}) interface{}
	GetWithWhere(model interface{}, whereClause string, params ...interface{}) []interface{}
	InsertModel(model interface{}) interface{}
	UpdateModel(model interface{}) bool
	DeleteModel(model interface{}) bool
}

type DbContext struct {
	connection *pg.DB
}

// NewDbContext DbContext constructor.
func NewDbContext() ContextInterface {
	dbc := new(DbContext)
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
func (dbc *DbContext) CreateSchema() error {
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

func (dbc *DbContext) GetModel(model interface{}) interface{} {
	err := dbc.connection.Select(&model)
	if err != nil && err != pg.ErrNoRows {
		// TODO: Log the error. Panic for now.
		panic(err)
	}

	if err == pg.ErrNoRows {
		return nil
	}

	return model
}

func (dbc *DbContext) GetWithWhere(model interface{}, whereClause string, params ...interface{}) []interface{} {
	var res []interface{}
	err := dbc.connection.Model(&model).Where(whereClause, params).Select(&res)
	if err != nil && err != pg.ErrNoRows {
		// TODO: Log the error. Panic for now.
		panic(err)
	}

	return res
}

func (dbc *DbContext) InsertModel(model interface{}) interface{} {
	err := dbc.connection.Insert(&model)
	if err != nil {
		// TODO: Log the error.
		return 0
	}

	return model
}

func (dbc *DbContext) UpdateModel(model interface{}) bool {
	err := dbc.connection.Update(&model)
	if err != nil {
		// TODO: Log the error.
		return false
	}

	return true
}

func (dbc *DbContext) DeleteModel(model interface{}) bool {
	err := dbc.connection.Delete(&model)
	if err != nil {
		// TODO: Log the error.
		return false
	}

	return true
}
