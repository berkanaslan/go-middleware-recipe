package config

type Profile string

const (
	Local Profile = "local"
	Test  Profile = "test"
)

var Profiles = []Profile{Local, Test}

var ActiveProfile Profile
