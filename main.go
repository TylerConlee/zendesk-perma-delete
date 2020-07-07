package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Establish variables to store data from flags in.
var (
	subdomain string
	user      string
	key       string
)

// User represents the output of a single user from the Deleted Users
// List returned by Zendesk. We're really only interested in the ID
// here, but because the whole object is returned, we represent it here.
type User struct {
	ID                int       `json:"id"`
	url               string    `json:"url"`
	name              string    `json:"name"`
	email             string    `json:"email"`
	createdAt         time.Time `json:"created_at"`
	updatedAt         time.Time `json:"updated_at"`
	timeZone          string    `json:"time_zone"`
	phone             string    `json:"phone"`
	sharedPhoneNumber string    `json:"shared_phone_number"`
	photo             string    `json:"photo"`
	localeID          int       `json:"locale_id"`
	locale            string    `json:"locale"`
	organizationID    int       `json:"organization_id"`
	role              string    `json:"role"`
	active            bool      `json:"active"`
}

// zendesk-perma-delete grabs a list of all of the Deleted Users from
// a specified Zendesk instance and cycles through that list to permanently
// delete them
func main() {
	parseFlags()
	users := getDeletedUsers()
	for _, ID := range users {
		deleteUser(ID)
	}
}

// getDeletedUsers retrieves the Deleted Users list from the specified Zendesk
// instance. It returns a list of IDs to be used to delete individual users
func getDeletedUsers() (users []int) {
	// data represents the returned structure of the response from Zendesk
	// after it's been parsed from JSON
	var data struct {
		Users []User `json:"deleted_users"`
		Next  string `json:"next_page"`
		Prev  string `json:"previous_page"`
		Count int    `json:"count"`
	}

	// Construct the URL used in the requests
	url := fmt.Sprintf("https://%s.zendesk.com/api/v2/deleted_users.json", subdomain)

	// To handle pagination, loop through the request process until the count
	// of users on our end matches the count of users on Zendesk's end.
	for {
		// Establish what the request is
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
		}

		// Use the provided Zendesk credentials
		req.SetBasicAuth(user, key)
		var client = &http.Client{}

		// Make the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &data)

		// Add the output to the Users slice, where all of the users are stored
		for _, u := range data.Users {
			users = append(users, u.ID)
		}

		// If the count of users matches the count that Zendesk has, break the
		// loop
		if data.Count != len(users) {
			url = data.Next
		} else {
			break
		}
	}

	return users
}

// deleteUser takes the deleted users' ID returned from getDeletedUsers
// and makes a DELETE request to the permadelete endpoint to permanently
// delete the user's information from Zendesk.
func deleteUser(ID int) {
	url := fmt.Sprintf("https://%s.zendesk.com/api/v2/deleted_users/%d.json", subdomain, ID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
	}

	// Use the provided Zendesk credentials
	req.SetBasicAuth(user, key)
	var client = &http.Client{}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("User ID", ID, "Deleted Successfully")
	} else {
		fmt.Println("Error in deleting ID", ID, ". Status Code:", resp.StatusCode, resp.Status)
	}

}

// parseFlags grabs the inputs from the input flags and assigns them to the
// related variables
func parseFlags() {
	// Define what each of the flags do

	// -user: The username that will access the Zendesk API. Can be used in
	// the form user@company.com/token
	flag.StringVar(&user, "user", "test@circleci.com/token", "Zendesk username for API")

	// -url: The subdomain for your Zendesk instance:
	// https://<subdomain>.zendesk.com
	flag.StringVar(&subdomain, "url", "circleci", "Zendesk subdomain (i.e. <companyname>.zendesk.com)")

	// -key: The API key that will be used to interact with Zendesk
	// API keys can be generated at
	// https://<subdomain>.zendesk.com/agent/admin/api/settings
	flag.StringVar(&key, "key", "XXXXX", "Zendesk API key")

	// Parse the inputted flags into usable variables
	flag.Parse()

}
