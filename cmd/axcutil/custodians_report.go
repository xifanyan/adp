package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/urfave/cli/v2"
	adp "github.com/xifanyan/adp"
)

func ensureEnginesAreRunning(engines adp.ListEntitiesResult) bool {
	for _, engine := range engines {
		if engine.ProcessStatus != "Running" {
			return false
		}
	}
	return true
}

func custodians(ctx *cli.Context) error {

	c := newClient(ctx).Assemble()
	engines, err := getEnginesByApplicationID(ctx, c)
	if err != nil {
		return err
	}

	if !ensureEnginesAreRunning(engines) {
		return fmt.Errorf("not all engines are running")
	}

	engine := engines[0]
	log.Debug().Msgf("engine %s running on %s\n", engine.ID, engine.HostID)

	return nil
}
