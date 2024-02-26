package main

import (
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

// personID: Person
var datalayer map[int]Person
var personID int

func init() {
	datalayer = make(map[int]Person)
	personID = 0

	log.Println("Data Layer initiated")

}

func main() {
	http.HandleFunc("/", healthHandler)

	log.Println("Server starting at port -> :4444")

	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal("server error ", err)
	}

	log.Println("server is stopped")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("server is running..."))
}

// Datalayer
func createPerson(p Person) int {
	pID := personID + 1
	datalayer[pID] = p

	personID = pID

	return pID
}

func readPerson(id int) (Person, error) {
	if person, ok := datalayer[id]; ok {
		return person, nil
	}

	return Person{}, fmt.Errorf("Person with id = %d not found", id)
}

func updatePerson(id int, p Person) error {
	if _, ok := datalayer[id]; ok {
		datalayer[id] = p
		return nil
	}

	return fmt.Errorf("Person with id = %d not found", id)
}

func deletePerson(id int) error {
	if _, ok := datalayer[id]; ok {
		delete(datalayer, id)
		return nil
	}

	return fmt.Errorf("Person with id = %d not found", id)
}
