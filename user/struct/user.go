package user

type Admin struct {
	Username string `gorm:"primaryKey"`
	Password string
	Auth     string
}
