package users

import "database/sql"

type UserRepository struct {
	Db *sql.DB
}

func (r UserRepository) GetAll() (*sql.Rows, error) {
	rows, err := r.Db.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	usersSlice := make([]User, 0)

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		usersSlice := append(usersSlice, user)
	}
}
