package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type IDField struct {
	ID int `db:"id,omitempty"`
}

type TableNamer interface {
	TableName() string
}

type User struct {
	IDField
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (u User) TableName() string {
	return "users"
}

func (u User) InsertQuery() string {
	return fmt.Sprintf("INSERT INTO %s (name, age) VALUES ($1, $2)", u.TableName())
}

type Vehicle struct {
	IDField
	Brand string `db:"brand"`
	Model string `db:"model"`
	Year  int    `db:"year"`
}

func (v Vehicle) TableName() string {
	return "vehicles"
}

func (v Vehicle) InsertQuery() string {
	return fmt.Sprintf("INSERT INTO %s (brand, model, year) VALUES ($1, $2, $3)", v.TableName())
}

func CreateTables(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            name TEXT,
            age INTEGER
        );
        CREATE TABLE IF NOT EXISTS vehicles (
            id INTEGER PRIMARY KEY,
            brand TEXT,
            model TEXT,
            year INTEGER
        );
    `)
	return err
}

func ReadFromDB[T TableNamer](db *sqlx.DB, id int) (*T, error) {
	var item T
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", item.TableName())
	err := db.Get(&item, query, id)
	return &item, err
}

func ReadAllFromDB[T TableNamer](db *sqlx.DB) ([]T, error) {
	var item T
	var items []T
	query := fmt.Sprintf("SELECT * FROM %s", item.TableName())
	err := db.Select(&items, query)
	return items, err
}

func main() {
	db, err := sqlx.Connect("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = CreateTables(db)
	if err != nil {
		panic(err)
	}

	// User'ı ekleme
	user := User{
		Name: "Umut",
		Age:  19,
	}
	_, err = db.Exec(user.InsertQuery(), user.Name, user.Age)
	if err != nil {
		panic(err)
	}

	// Vehicle'ı ekleme
	vehicle := Vehicle{
		Brand: "BMW",
		Model: "M4",
		Year:  2020,
	}

	_, err = db.Exec(vehicle.InsertQuery(), vehicle.Brand, vehicle.Model, vehicle.Year)
	if err != nil {
		panic(err)
	}

	// User'ı okuma
	u, err := ReadFromDB[User](db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("User name:", u.Name)

	users, err := ReadAllFromDB[User](db)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Println("User:", u)
	}

	// Vehicle'ı okuma
	v, err := ReadFromDB[Vehicle](db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Vehicle brand:", v.Brand)

	vehicles, err := ReadAllFromDB[Vehicle](db)
	if err != nil {
		panic(err)
	}
	for _, v := range vehicles {
		fmt.Println("Vehicle:", v)
	}
}
