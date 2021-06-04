package blog_strcut

import "time"

type Blog struct {
	ID         uint      `json:"blog_id" gorm:"autoincrement"`
	Title      string    `json:"title" gorm:"type:varchar(100);not null"`
	TopicGroup string    `json:"topic_group" gorm:"type:varchar(100);not null"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Content    string    `json:"content" gorm:"type:text;not null"`
	Author     string    `json:"author" gorm:"type:varchar(100);not null"`
}
