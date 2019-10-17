package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord is error model for no matching record
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials is error model for invalid credentials
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail is error model for duplicate email exists
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet is the snippet model struct
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User is the user model struct
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
