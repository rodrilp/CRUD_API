package main

import "testing"

func TestCreatePerson(t *testing.T) {
	id := createPerson(Person{Name: "Rodri", Age: 26})

	p, ok := datalayer[1]

	if id == 1 && ok && p.Name == "Rodri" && p.Age == 26 {
		t.Log("Person created!")
		return
	}

	t.Error("Create Person failed", id, p, ok, datalayer)
}

func TestReadPerson(t *testing.T) {
	p, err := readPerson(1)

	if err != nil {
		t.Error("Read person error", err)
	}

	if p.Name != "Rodri" || p.Age != 26 {
		t.Error("Invalid person. Expected Rodri, got: ", p.Name)
	}

	t.Log("Read person successful", p)
}

func TestUpdatePerson(t *testing.T) {
	err := updatePerson(1, Person{Name: "Rodri", Age: 30})

	if err != nil {
		t.Error("Update person error", err)
		return
	}

	p := datalayer[1]
	if p.Name != "Rodri" && p.Age != 30 {
		t.Error("Update failed. Expected 30, got: ", p.Age)
		return
	}

	t.Log("Update person successful", p)
}

func TestDeletePerson(t *testing.T) {
	err := deletePerson(1)

	if err != nil {
		t.Error("Delete person error", err)
		return
	}

	p, ok := datalayer[1]
	if ok {
		t.Error("Delete failed. Expected nill, got: ", ok, p)
		return
	}

	t.Log("Delete person successful", p)
}
