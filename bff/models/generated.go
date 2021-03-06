// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

// post input
type PostInput struct {
	// ID
	ID string `json:"id"`
	// ユーザー
	UserID string `json:"userID"`
	// 本文
	Body string `json:"body"`
}

// user input
type UserInput struct {
	// ID
	ID string `json:"id"`
	// email
	Email string `json:"email"`
	// name
	Name string `json:"name"`
	// image
	Image string `json:"image"`
}
