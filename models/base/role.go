package base

type UserRole string

const (
	Admin     UserRole = "admin"
	Moderator UserRole = "moderator"
	User      UserRole = "user"
)
