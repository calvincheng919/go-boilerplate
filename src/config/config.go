package config

import "os"

var MyEnvVar = os.Getenv("MY_ENV_VAR")
