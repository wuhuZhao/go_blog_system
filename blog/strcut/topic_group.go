package blog_strcut

type TopicGroup struct {
	ID    uint `gorm:"autoincrement"`
	TOPIC string
}
