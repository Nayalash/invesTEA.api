package models

import "time"

type Favourite struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Ticker    string `json:"ticker"`
	CreatedAt time.Time
	UserRefer int  `json:"user_id"`
	User      User `gorm:"foreignKey:UserRefer"`
}
