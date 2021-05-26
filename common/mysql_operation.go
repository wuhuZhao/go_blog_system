package common

type MysqlOperation interface {
	Add()
	Remove()
	Delete()
	Get()
}
