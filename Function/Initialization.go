package main

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "fmt"
	_ "net/mail"
)

//user is a struct type that declares user information.
type user struct {
	ID  int
	Name string
}

//updateStats provides update stats
type updateStats struct {
	Modified int
	Duration float64
	Success bool
	Message string
}

func main() {
	//Retire the user profile.
	// Retrieve the user profile.
	u, err := retrieveUser("Hoanh")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the user profile
	// Since the returned u is an address, use * to get the value.
	fmt.Printf("%+v\n", *u)

	// Update user name. Don't care about the update stats.
	// This _ is called blank identifier.
	// Since we don't need anything outside the scope of if, we can use the compact syntax.
	if _, err := updateUser(u); err != nil {
		fmt.Println(err)
		return
	}

	// Display the update was successful.
	fmt.Println("Updated user record for ID", u.ID)

}

//functions


//retrieveUser retrieve the user doc for specified user
func retrieveUser (name string) (*user, error) {
	//Make a call to get user in a json response.
	//set getter
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	// Goal: Unmarshal the json document into a value of the user struct type.
	// Create a value type user.
	var u user

	// Share the value down the call stack, which is completely safe so the Unmarshal function can
	// read the document and initialize it.
	err = json.Unmarshal([]byte(r), &u)

	// Share it back up the call stack.
	// Because of this line, we know that this create an allocation.
	// The value is the previous step is not on the stack but on the heap.
	return &u, err
}

//GETTER
//getUser simulates a web call that returns a json
//document for the specified user
func getUser(name  string)(string, error) {

	response := `{"ID":101, "Name":"Wyatt"}`
	return response, nil
}

//updateUser updates the specified user document

func updateUser(u*user)(*updateStats,error){
	//response simulates a JSON response
	response := `{"Modified":1, "Duration":0.005, "Success" : true, "Message": "updated"}`

	//Unmarshal the json document into  a var of userStats struct type.

	var us updateStats

	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}
	// Check the update status to verify the update is successful.
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}