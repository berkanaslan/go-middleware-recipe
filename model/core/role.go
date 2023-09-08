package core

type UserRole string

const (
	Admin     UserRole = "admin"
	Moderator UserRole = "moderator"
	StdUser   UserRole = "user"
)
