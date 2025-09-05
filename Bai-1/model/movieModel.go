package model

type Movie struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	Year  int    `json:"year" gorm:"not null"`
	Genre string `json:"genre" gorm:"not null"`
}
