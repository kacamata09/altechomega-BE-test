package main

import (
	httpRoutes "altech-omega-api/transport/http"
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	//// driver for pgsql
	// _ "github.com/lib/pq"

	//// driver for mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	viper.SetConfigFile("env.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	dbHost := viper.GetString("database.host")
	dbUser := viper.GetString("database.username")
	dbPort := viper.GetInt("database.port")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	// connection := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	e := echo.New()
	httpRoutes.StartHttp(e, db)

	port := viper.GetInt("server.port")
	e.Start(fmt.Sprintf(":%d", port))
}
