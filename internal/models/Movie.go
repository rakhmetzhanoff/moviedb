package models

type Movie struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	DirectorID int       `json:"director_id"`
	Director   *Director `gorm:"foreignKey:DirectorID" json:"director"`
}
