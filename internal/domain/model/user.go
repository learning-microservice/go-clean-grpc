package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id           string // ユーザID
	name         string // ユーザ名
	email        string // メールアドレス
	passwordHash []byte // パスワードHash
}

func NewUser(id, name, email string, passwordHash []byte) *User {
	return &User{
		id:           id,
		name:         name,
		email:        email,
		passwordHash: passwordHash,
	}
}

// Getter function

func (user *User) ID() string {
	return user.id
}

func (user *User) Name() string {
	return user.name
}

func (user *User) Email() string {
	return user.email
}

func (user *User) PasswordHash() []byte {
	return user.passwordHash
}

// Domain Logic

func (user *User) VerifyPassword(plainPassword string) error {
	return bcrypt.CompareHashAndPassword(user.passwordHash, []byte(plainPassword))
}
