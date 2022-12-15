package loader

import (
	"os"
)

func GetPort() string {
	return os.Getenv("PORT")
}

func GetMongoUri() string {
	return os.Getenv("MONGO_URI")
}
