package db

import (
	"2FAServer/models"
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DbContext defines an abstract Datbase connection.
type DbContext struct {
	connection *gorm.DB
}

// NewDbContext DbContext constructor.
func NewDbContext() ContextInterface {
	dbc := new(DbContext)
	dbc.connection = createConnection()

	if !dbc.connection.HasTable("Keys") {
		fmt.Println("Generating table schemas...")
		dbc.connection.CreateTable(new(models.Key))
	}

	return dbc
}

// Private functions.
func createConnection() *gorm.DB {
	configMap := make(map[string]string)
	dialect := "postgres"

	// PG configuration schema
	// https://godoc.org/github.com/lib/pq
	configMap["user"] = os.Getenv("PG_USERNAME")
	configMap["password"] = os.Getenv("PG_PASSWORD")
	configMap["dbname"] = os.Getenv("PG_DATABASE")
	configMap["sslmode"] = "disable"

	hostname := os.Getenv("PG_HOSTNAME")
	port := os.Getenv("PG_PORT")
	if hostname == "" {
		hostname = "127.0.0.1"
	}

	if port == "" {
		port = "5432"
	}

	configMap["host"] = hostname
	configMap["port"] = port

	configString := ""
	for k, v := range configMap {
		configString += k + "=" + v + " "
	}

	fmt.Println("Connecting to " + hostname + ":" + port)

	db, err := gorm.Open(dialect, configString)
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
		// TODO: Log the error. Panic for now.
		panic(err)
		return false
	}

	// Fetch success.
	return true
}

func (dbc *DbContext) GetWithWhere(model interface{}, refArray interface{},
	whereClause string, params ...interface{}) {

	dbc.connection.Where(&model).Find(refArray)
	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		// TODO: Log the error. Panic for now.
		panic(err)
	}
}

func (dbc *DbContext) InsertModel(model interface{}) bool {
	isNew := dbc.connection.NewRecord(model)
	if !isNew {
		return false
	}

	dbc.connection.Create(model)

	fmt.Println(model)

	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		// TODO: Log the error. Panic for now.
		panic(err)
		return false
	}

	// Fetch success.
	return true
}

func (dbc *DbContext) UpdateModel(model interface{}) bool {
	dbc.connection.Save(model)
	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		// TODO: Log the error. Panic for now.
		return false
	}

	return true
}

func (dbc *DbContext) DeleteModel(model interface{}) bool {
	dbc.connection.Delete(model)
	err := dbc.connection.GetErrors()
	if len(err) > 0 {
		// TODO: Log the error. Panic for now.
		return false
	}

	return true
}
