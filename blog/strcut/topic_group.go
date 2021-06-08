package blog_strcut

type TopicGroup struct {
	ID    uint   `json:"id" gorm:"autoincrement"`
	TOPIC string `json:"topic" gorm:"type:varchar(100);not null"`
}
