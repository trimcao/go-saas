package engine

type key int

const (
	// ContextOriginalPath holds the original request URL
	ContextOriginalPath key = iota
	// ContextRequestStart holds the request start time
	ContextRequestStart
	// ContextUserID holds the user ID (this is just for demo)
	ContextUserID
	// ContextTestKey is just a demo key
	ContextTestKey
	// ContextDatabase stores the database context
	ContextDatabase
	// ContextAuth stores the auth context?
	ContextAuth
	// ContextMinimumRole stores the minimum role of the call
	ContextMinimumRole
)
