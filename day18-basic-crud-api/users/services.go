package users

import "errors"

var store = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

var nextID = 3

func getAllUsers() ([]User, error) {
	return store, nil
}

func getUserByID(id int) (User, error) {
	for _, u := range store {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, errors.New("user not found")
}

func createUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	store = append(store, u)
	return u, nil
}

func updateUser(id int, updated User) (User, error) {
	for i, u := range store {
		if u.ID == id {
			updated.ID = id
			store[i] = updated
			return updated, nil
		}
	}
	return User{}, errors.New("user not found")
}

func deleteUser(id int) error {
	for i, u := range store {
		if u.ID == id {
			store = append(store[:i], store[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
