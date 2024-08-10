package user

import (
	"database/sql"
	"fmt"

	"github.com/Govind516/E-Commerce-Backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(database *sql.DB) *Store{
	return &Store{db:database}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error){
	rows, err := s.db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil{
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil{
			return nil, err
		}	
	}
	if u.ID == 0{
		return nil, fmt.Errorf("user not found")
	}
	return u, err
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error){
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
	)
	if err != nil{
		return nil, err
	}	
	return user, err
}

func (s *Store) GetUserById(id int) (*types.User, error){
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error{
	return nil
}