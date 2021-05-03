package configs

import (
	"errors"
	"os"

	godotenv "github.com/joho/godotenv"
)

type Config struct {
	AppId     string
	AuthToken string
}

func Load() (*Config, error) {
	r := Config{"", ""}

	err := godotenv.Load()
	if err != nil {
		return &r, errors.New("Load `$PROJRCT_ROOT/.env` file failed")
	}

	r.AppId, r.AuthToken = os.Getenv("DISCORD_APP_ID"), os.Getenv("DISCORD_AUTH_TOKEN")
	return &r, nil
}
