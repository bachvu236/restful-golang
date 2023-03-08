package models

import (
	"database/sql"
	"errors"
	"html"
	"log"
	"strings"
	"time"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32
	Nickname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *sql.DB) (*User, error) {

	var err error
	result, err := db.Exec("INSERT INTO employee(nickname, email, password,created_at,updated_at) VALUES(?, ?, ?, ?, ?)", u.Nickname, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return &User{}, err
	}
	uid, _ := result.LastInsertId()
	u.ID = uint32(uid)
	return u, nil
}

func (u *User) FindAllUsers(db *sql.DB) (*[]User, error) {
	var err error
	user := User{}
	users := []User{}
	rows, err := db.Query("SELECT id, nickname, email, password, created_at, updated_at FROM employee")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			users = append(users, user)
		}
	}
	return &users, err
}

func (u *User) FindUserByID(db *sql.DB, uid uint32) (*User, error) {
	var err error
	user := User{}
	result, err := db.Query("SELECT id, nickname, email, password, created_at, updated_at FROM employee WHERE id = ?",uid)
	if err != nil {
		panic(err.Error())
	}
		err = result.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Fatal(err.Error())
		}
	return &user, err
}

func (u *User) UpdateAUser(db *sql.DB, uid uint32) (*User, error) {
	// To hash the password
	// err := u.BeforeSave()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
	// 	map[string]interface{}{
	// 		"password":  u.Password,
	// 		"nickname":  u.Nickname,
	// 		"email":     u.Email,
	// 		"updated_at": time.Now(),
	// 	},
	// )
	// if db.Error != nil {
	// 	return &User{}, db.Error
	// }
	// This is the display the updated user
	// err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	// if err != nil {
	// 	return &User{}, err
	// }
	return u, nil
}

func (u *User) DeleteAUser(db *sql.DB, uid uint32) (int64, error) {

	// db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	// if db.Error != nil {
	// 	return 0, db.Error
	// }
	// return db.RowsAffected, nil
	return 0, nil
}
