package main

// zendesk-perma-delete grabs a list of all of the Deleted Users from
// a specified Zendesk instance and cycles through that list to permanently
// delete them
func main() {
	users := getDeletedUsers()
	for user := range users {
		deleteUser(user)
	}
}

// getDeletedUsers retrieves the Deleted Users list from the specified Zendesk
// instance. It returns a list of IDs to be used to delete individual users
func getDeletedUsers() (users []int) {
	return users
}

// deleteUser takes the deleted users' ID returned from getDeletedUsers
// and makes a DELETE request to the permadelete endpoint to permanently
// delete the user's information from Zendesk.
func deleteUser(ID int) {

}
