// internal/models/user.go
package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model represents a user in the system
type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	Email        string
	PasswordHash string
	Role         string
}

// Register creates a new user record in the database
func (user *User) Register(db *gorm.DB) error {
	err := db.Create(user).Error
	return err
}

// ChangePassword updates the user's password hash in the database
func (u *User) ChangePassword(db *gorm.DB, newPassword string) error {
	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update the user's password hash in the database
	if err := db.Model(u).Update("PasswordHash", hashedPassword).Error; err != nil {
		return err
	}

	return nil
}
