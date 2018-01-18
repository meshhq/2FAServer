package db

import (
	"2FAServer/models"
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	// Used to import a specific dialect.
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DbContext defines an Database connection.
type DbContext struct {
	connection *gorm.DB
}

// NewDbContext DbContext constructor.
func NewDbContext() ContextInterface {
	dbc := new(DbContext)
	dbc.connection = createConnection()

	// TODO: Move this to a Migration function.
	if !dbc.connection.HasTable("Keys") {
		fmt.Println("Generating table schemas...")
		dbc.connection.CreateTable(new(models.Key))
	}

	return dbc
}

// Private functions.
func createConnection() *gorm.DB {
	dialect := "postgres"
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	config := &Configuration{
		User:     os.Getenv("PG_USERNAME"),
		Password: os.Getenv("PG_PASSWORD"),
		DbName:   os.Getenv("PG_DATABASE"),
		SslMode:  "disable",
		Hostname: os.Getenv("PG_HOSTNAME"),
		Port:     port,
	}

	fmt.Println("Connecting to " + config.Hostname + ":" + strconv.Itoa(config.Port))

	db, err := gorm.Open(dialect, config.GetConfig())
	if err != nil {
		panic(err)
	}

	enableLog, _ := strconv.ParseBool(os.Getenv("DB_ENABLE_LOGS"))
	db.LogMode(enableLog)

	return db
}

func (dbc *DbContext) GetModel(model interface{}) bool {
	dbc.connection.First(model)
	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		return false
	}

	// Fetch success.
	return true
}

func (dbc *DbContext) GetWithWhere(refArray interface{}, whereClause string, params ...interface{}) {
	dbc.connection.Where(whereClause, params...).Find(refArray)
	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		panic(err)
	}
}

func (dbc *DbContext) InsertModel(model interface{}) bool {
	isNew := dbc.connection.NewRecord(model)
	if !isNew {
		return false
	}

	dbc.connection.Create(model)

	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		return false
	}

	// Fetch success.
	return true
}

func (dbc *DbContext) UpdateModel(model interface{}) bool {
	dbc.connection.Save(model)
	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		return false
	}

	// Update success.
	return true
}

func (dbc *DbContext) DeleteModel(model interface{}) bool {
	dbc.connection.Delete(model)
	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		return false
	}

	// Delete success.
	return true
}
