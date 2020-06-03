package linux

import (
	log "github.com/mgutz/logxi/v1"
)

var logger = log.New("linux")

func init() {
	logger.SetLevel(log.LevelAll)
}
