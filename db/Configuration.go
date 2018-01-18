package db

import (
	"strconv"
	"strings"
)

type Configuration struct {
	Hostname string
	Port     int
	User     string
	Password string
	DbName   string
	SslMode  string
}

func (c *Configuration) GetConfig() string {
	configMap := make(map[string]string)

	hostname := c.Hostname
	port := c.Port
	if hostname == "" {
		hostname = "127.0.0.1"
	}

	if port == 0 {
		port = 5432
	}

	configMap["host"] = hostname
	configMap["port"] = strconv.Itoa(port)
	configMap["user"] = c.User
	configMap["password"] = c.Password
	configMap["dbname"] = c.DbName
	configMap["sslmode"] = c.SslMode

	configString := ""
	for k, v := range configMap {
		configString += k + "=" + v + " "
	}

	return strings.TrimSpace(configString)
}
