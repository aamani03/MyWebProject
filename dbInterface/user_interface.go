package dbInterface

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserService interface {
	GetUser(string) (*User, error)
	CreateUser(*User) (string, error)
	UpdateUser(*User) error
	DeleteUser(*User) error
	LoginUser(*User) error
	LogoutUser(*User) error
}

type DbRepo struct {
	SqlConnection *sql.DB
}

func (u *DbRepo) GetUser(id string) (*User, error) {
	row := u.SqlConnection.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	var user User
	err := row.Scan(&user.UserID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}

	return &user, err
}

func (u *DbRepo) CreateUser(usr *User) (string, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	result, err := u.SqlConnection.Exec(query, usr.Name, usr.Email, usr.Password)

	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("failed to retrive last inserted id %v", err)
	}

	return fmt.Sprint("id", id), nil
}

func (u *DbRepo) UpdateUser(usr *User) error {
	query := `UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`

	_, err := u.SqlConnection.Exec(query, usr.Name, usr.Email, usr.Password, usr.UserID)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (u *DbRepo) DeleteUser(usr *User) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := u.SqlConnection.Exec(query, usr.UserID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check if user was deleted: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d", usr.UserID)
	}

	return nil
}

func (u *DbRepo) LoginUser(usr *User) error  { return nil }
func (u *DbRepo) LogoutUser(usr *User) error { return nil }
