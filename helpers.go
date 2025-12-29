package main

const appVersion = "v1"

// VersionedGreeting returns a greeting with version prefix.
func VersionedGreeting(name string) string {
	return "[" + appVersion + "] Hello, " + name
}
