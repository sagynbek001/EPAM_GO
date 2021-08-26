package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "address"
)

type Contact struct {
	Name       string
	Phone      string
	Group_name string
}

func addContact(contact Contact) error {
	_, err := db.Exec("INSERT INTO contacts (name, phone, group_name) VALUES ($1, $2, $3)", contact.Name, contact.Phone, contact.Group_name)
	if err != nil {
		return fmt.Errorf("addContact: %v", err)
	}
	return nil
}

func assignContact(id int64, group_name string) error {
	_, err := db.Exec("UPDATE contacts SET group_name = $1 WHERE contact_id = $2", group_name, strconv.Itoa(int(id)))
	if err != nil {
		return fmt.Errorf("assignContact: %v", err)
	}
	return nil
}

func listContacts(group_name string) error {
	var (
		contact_id int
		name       string
		phone      string
		group      string
	)
	rows, err := db.Query("SELECT * FROM contacts WHERE group_name=$1", group_name)
	if err != nil {
		return fmt.Errorf("listContacts: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&contact_id, &name, &phone, &group)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(contact_id, name, phone, group)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	exampleContact1 := Contact{
		Name:       "Vasia",
		Phone:      "+77887788978",
		Group_name: "Friends",
	}

	exampleContact2 := Contact{
		Name:       "Ania",
		Phone:      "+77887788786",
		Group_name: "Family",
	}

	exampleContact3 := Contact{
		Name:       "Maks",
		Phone:      "+77887788123",
		Group_name: "undeclared",
	}

	err = addContact(exampleContact1)
	if err != nil {
		fmt.Print(err)
	}
	err = addContact(exampleContact2)
	if err != nil {
		fmt.Print(err)
	}
	err = addContact(exampleContact3)
	if err != nil {
		fmt.Print(err)
	}

	err = assignContact(0, "Friends")
	if err != nil {
		fmt.Print(err)
	}
	err = assignContact(1, "Friends")
	if err != nil {
		fmt.Print(err)
	}
	err = assignContact(2, "Friends")
	if err != nil {
		fmt.Print(err)
	}

	err = listContacts("Friends")
	if err != nil {
		fmt.Print(err)
	}
}
