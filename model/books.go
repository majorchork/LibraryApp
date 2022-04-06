package models

type Books struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn" gorm:"unique"`
}
