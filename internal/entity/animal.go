package entity

import "time"

type Animal struct {
	ID          string    `json:"id,omitempty"`
	PhotoURLs   []string  `json:"photoURLs"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
