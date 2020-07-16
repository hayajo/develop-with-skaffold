package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	env "github.com/caarlos0/env/v6"
	_ "github.com/lib/pq"
)

// SQLConn ...
var SQLConn *sql.DB

// DBConfig ...
type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
}

func main() {
	dbConf := DBConfig{}
	if err := env.Parse(&dbConf); err != nil {
		log.Fatal(err)
	}

	if err := CreateDBConnection(dbConf); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Hello ...
func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Print(err)
	}
}

// CreateDBConnection ...
func CreateDBConnection(conf DBConfig) error {
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password)
	SQLConn, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	err = SQLConn.Ping()
	if err != nil {
		return err
	}

	return nil
}
