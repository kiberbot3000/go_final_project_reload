package settings

import (
	"os"
)

const Template = "20060102"

type TodoEnv struct {
	Port     string
	DBFile   string
	Password string
}

func checkENV(environment, baseValue string) string {
	if value, ok := os.LookupEnv(environment); ok {
		return value
	}
	return baseValue
}

func GetEnv() *TodoEnv {
	port := checkENV("TODO_PORT", "7540")
	db := checkENV("TODO_DBFILE", "")
	password := checkENV("TODO_PASSWORD", "")

	return &TodoEnv{
		Port:     port,
		DBFile:   db,
		Password: password,
	}
}
