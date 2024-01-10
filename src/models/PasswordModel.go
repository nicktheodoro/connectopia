package models

type PasswordModel struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
