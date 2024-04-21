package models

type Settings struct {
	AppParams  Params
	DBSettings DBSettings
}

type Params struct {
	PortRun string
}

type DBSettings struct {
	User     string
	Password string
	Host     string
	Port     string
	Protocol string
	Database string
	SSLMode  string
}
