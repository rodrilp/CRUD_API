package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Data    Person `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
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
	http.HandleFunc("/person", personHandler)

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

func personHandler(w http.ResponseWriter, r *http.Request) {
	var res Response
	// send the response at the end
	defer func() {
		json.NewEncoder(w).Encode(res)
	}()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	switch r.Method {
	case http.MethodPost:
		var p Person

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			res.Error = err.Error()
			return
		}

		id := createPerson(p)
		res.Message = fmt.Sprintf("New person created with id: %d", id)

	case http.MethodGet:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("Invalid ID")
			res.Error = "Invalid ID"
			w.WriteHeader(400)
		}

		// conver id into int
		id, err := strconv.Atoi(strID)
		if err != nil {
			log.Println("Invalid id", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		p, err := readPerson(id)
		if err != nil {
			log.Println("Read erro", err)
			res.Error = err.Error()
			w.WriteHeader(400)
		}

		res.Data = p

	case http.MethodPut:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("Invalid ID")
			res.Error = "Invalid ID"
			w.WriteHeader(400)
		}

		// conver id into int
		id, err := strconv.Atoi(strID)
		if err != nil {
			log.Println("Invalid id", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		// read json body and conver to struct
		var p Person
		json.NewDecoder(r.Body).Decode(&p)

		err = updatePerson(id, p)
		if err != nil {
			log.Println("Update error", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		res.Message = "Person successfully updated."

	case http.MethodDelete:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("Invalid ID")
			res.Error = "Invalid ID"
			w.WriteHeader(400)
		}

		// conver id into int
		id, err := strconv.Atoi(strID)
		if err != nil {
			log.Println("Invalid id", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		err = deletePerson(id)
		if err != nil {
			log.Println("Delete error", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		res.Message = "Person successfully deleted."
	}
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
