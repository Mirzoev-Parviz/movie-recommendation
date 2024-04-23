package dto

import "github.com/Mirzoev-Parviz/movie-recommendation/models"

// I know it's stupid, but life made me do it
var InteractionsData [][]string
var ItemsData [][]string
var UserData [][]string

var Interactions []models.Interactions
var Items []models.Item
var Users []models.User

type RCMInput struct {
	UserID int `json:"user_id"`
}
