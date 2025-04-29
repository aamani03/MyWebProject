package dbInterface

import (
	"database/sql"
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
	DB *sql.DB
}

func (u *DbRepo) GetUser(id string) (*User, error) {
	row := u.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	var user User
	err := row.Scan(&user.UserID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}

	return &user, err
}

func (u *DbRepo) CreateUser(usr *User) (string, error) { return "", nil }
func (u *DbRepo) UpdateUser(usr *User) error           { return nil }
func (u *DbRepo) DeleteUser(usr *User) error           { return nil }
func (u *DbRepo) LoginUser(usr *User) error            { return nil }
func (u *DbRepo) LogoutUser(usr *User) error           { return nil }

type TaskService interface {
	GetTask(string) (*Task, error)
}

type TaskRepo struct {
	DB *sql.DB
}

func (u *TaskRepo) GetTask(string) (*Task, error) { return nil, nil }
