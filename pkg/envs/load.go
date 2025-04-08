package envs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	err := godotenv.Load(".env")
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return err
}
