package movie

type Movie struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title string `json:"title" gorm:"not null"`
	URL   string `json:"url" gorm:"not null"`
}
