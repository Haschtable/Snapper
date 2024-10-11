package main

import (
	config "github.com/Haschtable/snapper/Config"
	logger "github.com/Haschtable/snapper/Logger"
)

func main() {
	minVersion := config.NewVersion(0, 1, 0)
	Logger := logger.NewLogger(nil)
	Logger.Info("Snapper " + config.CurrentVersion.String())
	Logger.Assert(config.CurrentVersion.Min(minVersion), "Snapper is outdated. Please update to the latest version.")
}
