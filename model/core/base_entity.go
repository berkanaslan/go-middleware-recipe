package core

type BaseEntity interface {
	GetID() uint
}

type TableNamer interface {
	TableName() string
}
