package models

type User struct {
	ID      int    `json:"id"`
	Age     string `json:"age"`
	Income  string `json:"income"`
	Sex     string `json:"sex"`
	HasKids bool   `json:"kids_flg"`
}
