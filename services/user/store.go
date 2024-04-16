package user

import (
	"database/sql"
	"fmt"

	"github.com/rodolfole/go-users-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CheckIfUserExist(correo string, telefono string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE correo = ? OR telefono = ?", correo, telefono)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) GetUserByCorreo(correo string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE correo = ?", correo)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.Usuario,
		&user.Correo,
		&user.Telefono,
		&user.Contrasena,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {

	fmt.Println(user.Correo)

	// return nil

	_, err := s.db.Exec("INSERT INTO users (usuario, correo, telefono, contrasena) VALUES (?,?,?,?)", user.Usuario, user.Correo, user.Telefono, user.Contrasena)
	if err != nil {
		return err
	}

	return nil
}
