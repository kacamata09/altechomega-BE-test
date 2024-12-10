package main

import (
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../env.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	dbHost := viper.GetString("database.host")
	dbUser := viper.GetString("database.username")
	dbPort := viper.GetInt("database.port")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")

	connection := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	m, err := migrate.New(

		"file://mysql",
		connection,
	)

	if err != nil {

		log.Fatalf("Failed to create migrate: %v", err)

	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {

		log.Fatalf("Failed to apply migrations: %v", err)

	}

	fmt.Println("Migrations applied successfully!")

}
