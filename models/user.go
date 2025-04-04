package models

import "database/sql"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Address   string
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) All() ([]User, error) {
	rows, err := m.DB.Query("SELECT id, firstname, lastname, email, address FROM \"user\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	row := m.DB.QueryRow("SELECT id, firstname, lastname, email, address FROM \"user\" WHERE id = $1", id)

	u := &User{}
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Address)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m *UserModel) Create(firstName, lastName, email, address string) error {
	_, err := m.DB.Exec(
		"INSERT INTO \"user\" (firstname, lastname, email, address) VALUES ($1, $2, $3, $4)",
		firstName, lastName, email, address)
	return err
}

func (m *UserModel) Update(id int, firstName, lastName, email, address string) error {
	_, err := m.DB.Exec(
		"UPDATE \"user\" SET firstname = $1, lastname = $2, email = $3, address = $4 WHERE id = $5",
		firstName, lastName, email, address, id)
	return err
}

func (m *UserModel) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM \"user\" WHERE id = $1", id)
	return err
}
