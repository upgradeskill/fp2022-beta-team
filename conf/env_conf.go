package conf

import "strings"

// in mem package variable
var env = "dev"

// call this func for simple feature flag.
// when your feature need to disable in prod
func IsDev() bool {
	return strings.ToLower(env) == "dev"
}
