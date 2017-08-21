package config

// Context context holder
type Context struct {
	SQL    string
	Source string
	Output string
	Pretty bool
}

// Global global vars
var Global Context
