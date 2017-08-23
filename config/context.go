package config

// Context context holder
type Context struct {
	// SQL query string
	SQL string
	// Source db dsn connection string
	Source string
	// Output to where
	Output string
	// Pretty json output
	Pretty bool
}

// Global global vars
var Global Context
