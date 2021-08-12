package models

// Users is a model for the userd table
type Users struct {
	UserId int `json:"user_id" orm:"auto"`
	UserName string `json:"user_name" orm:"size(32)"`
	Email string `json:"email" orm""size(128)"`
	Password string `json:"password" orm:size(64)"`
}