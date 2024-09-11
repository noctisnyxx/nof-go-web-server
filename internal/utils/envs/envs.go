package envs

import (
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var (
	MONGO_HOST  = os.Getenv("MONGO_HOST")
	MONGO_ATLAS = os.Getenv("MONGO_ATLAS")
	MONGO_USER  = os.Getenv("MONGO_USER")
	MONGO_PASS  = os.Getenv("MONGO_PASS")
	MONGO_LOCAL = os.Getenv("MONGO_LOCAL")
)
