package models

import "time"

type Items struct {
	ID          int      `json:"item_id"`
	ContentType string   `json:"content_type"`
	Title       string   `json:"title"`
	ReleaseYear float32  `json:"release_year"`
	Genres      []string `json:"genres"`
	Countries   []string `json:"countries"`
	ForKids     float64  `json:"for_kids"`
	AgeRating   float64  `json:"age_rating"`
	Studios     []string `json:"studios"`
	Directors   []string `json:"directors"`
	Actors      []string `json:"actors"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}
type Interactions struct {
	UserID      int       `json:"user_id"`
	ItemID      int       `json:"item_id"`
	LastWatchDT time.Time `json:"last_watch_dt"`
	TotalDur    int       `json:"total_dur"`
	WatchedPCT  float64   `json:"watched_pct"`
}