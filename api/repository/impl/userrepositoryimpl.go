package impl

import (
	"database/sql"
	"errors"
	"github.com/vanilla/go-jwt-crud/api/entities"
	"time"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) FindAll() ([]entities.User, error) {
	row, err := r.db.Query(`
		SELECT * FROM users ORDER By id ASC
	`)
	defer row.Close()

	if err != nil {
		return nil, err
	}

	var users []entities.User
	for row.Next() {
		var user entities.User

		err := row.Scan(
				&user.ID,
				&user.Username,
				&user.Password,
				&user.Email,
				&user.Photo,
				&user.IsActive,
				&user.CreatedAt,
				&user.UpdatedAt,
			)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindById(uid uint64) (entities.User, error) {
	row, err := r.db.Prepare(`
		SELECT * FROM users WHERE id=$1
	`)
	defer row.Close()

	if err != nil {
		return entities.User{}, err
	}

	var user entities.User
	err = row.QueryRow(uid).Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Photo,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		return entities.User{}, err
	}

	if user.ID.Valid == false {
		return entities.User{}, errors.New("User not found")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Save(postUser entities.User) (bool, error) {
	query, err := r.db.Prepare(`
		SELECT id FROM users WHERE username=$1
	`)
	defer query.Close()

	if err != nil {
		return false, err
	}

	var user entities.User
	err = query.QueryRow(postUser.Username.String).Scan(&user.ID)

	if user.ID.Valid == true {
		return false, errors.New("User already exists")
	}

	row, err := r.db.Prepare(`
		INSERT INTO users (username, password, email, photo, is_active, created_at) VALUES ($1, $2, $3, $4, $5, $6)
	`)
	defer row.Close()

	if err != nil {
		return false, err
	}

	_, err = row.Exec(
			postUser.Username.String,
			postUser.Password.String,
			postUser.Email.String,
			postUser.Photo.String,
			postUser.IsActive.Bool,
			time.Now(),
		)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepositoryImpl) Update(uid uint64, postUser entities.User) (bool, error) {
	query, err := r.db.Prepare(`
		SELECT id FROM users WHERE id=$1
	`)
	defer query.Close()

	if err != nil {
		return false, err
	}

	var user entities.User
	err = query.QueryRow(uid).Scan(&user.ID)

	if user.ID.Valid == false {
		return false, errors.New("User not found")
	}

	row, err := r.db.Prepare(`
		UPDATE users SET username=$1, password=$2, email=$3, photo=$4, is_active=$5, updated_at=$6 WHERE id=$7
	`)
	defer row.Close()

	if err != nil {
		return false, nil
	}

	_, err = row.Exec(
			postUser.Username.String,
			postUser.Password.String,
			postUser.Email.String,
			postUser.Photo.String,
			postUser.IsActive.Bool,
			time.Now(),
			uid,
		)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepositoryImpl) Delete(uid uint64) (bool, error) {
	query, err := r.db.Prepare(`
		SELECT id FROM users WHERE id=$1
	`)
	defer query.Close()

	if err != nil {
		return false, err
	}

	var user entities.User
	err = query.QueryRow(uid).Scan(&user.ID)

	if user.ID.Valid == false {
		return false, errors.New("User not found")
	}

	row, err := r.db.Prepare(`
		DELETE FROM users WHERE id=$1
	`)
	defer row.Close()

	if err != nil {
		return false, nil
	}

	_, err = row.Exec(uid)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepositoryImpl) Login(username string) (entities.User, error) {
	row, err := r.db.Prepare(`
		SELECT * FROM users WHERE username=$1
	`)
	defer row.Close()

	if err != nil {
		return entities.User{}, err
	}

	var user entities.User
	err = row.QueryRow(username).Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Photo,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}