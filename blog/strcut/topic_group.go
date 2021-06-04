package blog_strcut

type TopicGroup struct {
	ID    uint `gorm:"autoincrement"`
	TOPIC string `gorm:"type:varchar(100);not null"`
}
