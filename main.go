package dependency_container

import (
	"context"
	"github.com/autorootcause/dependency-container.git/app/config"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	ctx := context.Background()
	logger := zerolog.Ctx(ctx)
	filename, valid := os.LookupEnv("DEPENDENCY_CONTAINER_CONFIG_FILE")
	if !valid {
		cfgPath, ok := os.LookupEnv("DEPENDENCY_CONTAINER_CONFIG_FILE")
		if !ok {
			logger.Fatal().Msgf("environment variable DEPENDENCY_CONTAINER_CONFIG_FILE is not set")
		}
		logger.Fatal().Msgf("cannot read config file at path: %s", cfgPath)
	}
	cfg, err := config.LoadConfig(filename)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error loading config")
	}

}
