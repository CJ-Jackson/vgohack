package listjson

import "time"

type Module struct {
	Path     string
	Version  string
	Replace  *Module
	Time     *time.Time
	Update   *Module
	Main     bool
	Dir      string
	Error    *ModuleError
}

type ModuleError struct {
	Err string
}