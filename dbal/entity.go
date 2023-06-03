package dbal

type Entity interface {
	TableName() string
}
