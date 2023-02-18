package tests

import (
	"fastgameserver/framework/logger"
	"fastgameserver/framework/store/redis"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: redis.Flags,
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(redis.Opts.RedisUrl)
}