package db

import (
	"strconv"
	"strings"
)

type Configuration struct {
	hostname string
	port     int
	User     string
	Password string
	DbName   string
	SslMode  string
}

func (c *Configuration) GetConfig() string {
	configMap := make(map[string]string)

	configMap["host"] = c.hostname
	configMap["port"] = strconv.Itoa(c.port)
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

func (c *Configuration) GetPort() int {
	return c.port
}

func (c *Configuration) GetHost() string {
	return c.hostname
}

func (c *Configuration) SetPort(port int) {
	if port == 0 {
		c.port = 5432
	} else {
		c.port = port
	}
}

func (c *Configuration) SetHost(hostname string) {
	if hostname == "" {
		c.hostname = "127.0.0.1"
	} else {
		c.hostname = hostname
	}
}
