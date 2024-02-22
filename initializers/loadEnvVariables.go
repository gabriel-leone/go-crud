package initializers

import "os"

func LoadEnvVariables() {
	os.Setenv("PORT", "8080")
	os.Setenv("GIN_MODE", "release")
}
