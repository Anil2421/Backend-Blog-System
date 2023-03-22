package models

type Blog struct {
	//ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ID      int    `gorm:"primary_key" json:"id,omitempty"`
	Title   string `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Content string `gorm:"not null" json:"content,omitempty"`
	Author  string `gorm:"not null" json:"author,omitempty"`
}
