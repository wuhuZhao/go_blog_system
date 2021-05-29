package blog_strcut

import "time"

type Blog struct {
	ID         uint      `json:"blog_id" gorm:"autoincrement"`
	Title      string    `json:"title"`
	TopicGroup string    `json:"topic_group"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
}
