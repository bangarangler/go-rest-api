package main_test

import (
	// "bytes"
	// "encoding/json"
	"github.com/bangarangler/go-rest-api"
	"github.com/joho/godotenv"
	"log"
	// "net/http"
	// "net/http/httptest"
	"os"
	// "strconv"
	"testing"
)

var a main.App

func goDotEnvVar(key string) string {
	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatalf("error loading env")
	}
	return os.Getenv(key)
}

func TestMain(m *testing.M) {
	a.Initialize(
		goDotEnvVar("POSTGRES_USER"),
		goDotEnvVar("POSTGRES_PASSWORD"),
		goDotEnvVar("POSTGRES_DB"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
  id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id)
)
`
