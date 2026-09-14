package archive

type Archive struct {
	ID    int    `json:"id" gorm:"primaryKey,autoIncrement"`
	URL   string `json:"url" gorm:"not null;unique"`
	Depth int    `json:"depth" gorm:"not null"`
}
