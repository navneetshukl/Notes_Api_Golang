package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/navneetshukl/Golang_notes_API/models"
)

//var conn *sql.DB
//var err error

func DB_Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connString := os.Getenv("DB_CONNECTION_STRING")
	if connString == "" {
		log.Fatal("DB_CONNECTION_STRING not found in .env file")
	}
	conn, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Unable to connect: %v\n", err))
		return nil, err
	}

	//defer conn.Close()

	log.Println("Connected to database")

	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot Ping the database")
		return nil, err
	}
	log.Println("pinged database")

	return conn, nil

}

func InsertIntoUser(name, email, password string) {
	conn, err := DB_Connect()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	query := `insert into users(name,email,password) values($1,$2,$3)`

	_, err = conn.Exec(query, name, email, password)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a row")
}

func InsertintoNotes(title string) {

	conn, err := DB_Connect()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	query := `insert into notes(title) values($1)`

	_, err = conn.Exec(query, title)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a row")
}

func CheckUser(email, password string) (bool, error) {

	conn, err := DB_Connect()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	query := "SELECT COUNT(*) FROM users WHERE email = $1 AND password = $2"
	var count int
	err = conn.QueryRow(query, email, password).Scan(&count)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil

}

func GetDataFromNotes() ([]models.Notes, error) {
	conn, err := DB_Connect()
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	query := "SELECT * FROM notes"

	rows, err := conn.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.Notes

	for rows.Next() {
		var user models.Notes

		err := rows.Scan(&user.Id, &user.Title, &user.Email)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		return nil, err
	}

	return users, nil
}

func SaveData(title, email string) error {
	// Replace with your PostgreSQL connection details
	conn, err := DB_Connect()
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	// Prepare the SQL statement
	stmt, err := conn.Prepare("INSERT INTO notes (title,email) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user data
	_, err = stmt.Exec(title, email)
	if err != nil {
		return err
	}

	fmt.Println("Data saved successfully")

	return nil
}

func DeleteData(email string) error {
	// Replace with your PostgreSQL connection details
	conn, err := DB_Connect()
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	// Prepare the SQL statement
	stmt, err := conn.Prepare("DELETE FROM notes WHERE email = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the ID parameter
	_, err = stmt.Exec(email)
	if err != nil {
		return err
	}

	fmt.Println("Data deleted successfully")

	return nil
}
