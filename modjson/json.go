package modjson

type Module struct {
	Path string
	Version string
}

type Replace struct{ Old, New Module }

type GoMod struct {
	Module Module
	Require []Module
	Exclude []Module
	Replace []Replace
}