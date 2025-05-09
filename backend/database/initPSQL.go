package database

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey;autoIncrement;not null;unique"`

	Name     string
	Email    string
	Password string

	Tokens *[]BrowserToken
	Perm   *Permission
}

type BrowserToken struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey;autoIncrement;not null;unique"`

	DeviceID string `gorm:"unique"`
	Token    string `gorm:"unique"`

	UserID uint64 `gorm:"index"`
	User   User
}

type Permission struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey;autoIncrement;not null;unique"`

	Login bool `json:"login"`
	Admin bool `json:"admin"`

	UserID uint64 `gorm:"index"`
	User   User
}

// InitPSQL Creates all the necessary tables
// Also creates an initial admin user
func InitPSQL(db *gorm.DB) {
	var err error

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Error migrating Users table: %v\n", err)
	}

	err = db.AutoMigrate(&BrowserToken{})
	if err != nil {
		log.Fatalf("Error migrating BrowserTokens table: %v\n", err)
	}

	err = db.AutoMigrate(&Permission{})
	if err != nil {
		log.Fatalf("Error migrating Permissions table: %v\n", err)
	}

	// Create initial admin user
	err = db.Where("name = admin").First(&User{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var user User
		user.Name = "admin"
		user.Password = "admin"
		user.Email = "admin@example.com"

		err = db.Create(&user).Error
		if err != nil {
			log.Fatalf("Error creating initial admin user: %v\n", err)
		}

		log.Println("Initial admin user created")
	} else {
		log.Println("Admin already exists ...")
	}

	log.Println("Database initialized")
}
