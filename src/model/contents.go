package model

import "time"

type Content struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Winner    uint      `json:"winner"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LimitDate time.Time `json:"limit_date"`
	Link      string    `json:"link"`
	Provider  string    `json:"provider"`
	Way       string    `json:"way"`
	Category  []string  `json:"category"`
	IsOneclick bool 		`json:"is_oneclick"`
	TwitterWays []string `json:"twitter_way"`
	TwitterRaw string		`json:"twitter_raw"`
}
