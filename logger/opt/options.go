package opt

import "github.com/google/uuid"

type LoggerOpts struct {
	Opts          *GeneralOpts
	StdLoggerOpts *StdLoggerOpts
}

type GeneralOpts struct {
	AppVersion string
	InstanceID uuid.UUID
	Env        string
	AppName    string
	Level      string
}

type StdLoggerOpts struct {
	LogFile  string
	Stdout   bool
	Disabled bool
}
