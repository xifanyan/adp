package main

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/urfave/cli/v2"
	adp "github.com/xifanyan/adp"
)

func engineDistribution(ctx *cli.Context) error {
	var err error

	c := newClient(ctx).Assemble()
	engines, err := getEnginesByApplicationID(ctx, c)

	filtered := []adp.Entity{}
	for _, engine := range engines {
		if engine.ProcessStatus == "Running" {
			log.Debug().Msgf("%s %s\n", engine.ID, engine.HostID)
			filtered = append(filtered, engine)
		}
	}

	log.Debug().Msgf("filtered: %+v", filtered)

	for _, engine := range filtered {
		result, err := queryEngine(ctx, c, engine)
		if err != nil {
			log.Error().Msgf("%s %s\n", engine.ID, engine.HostID)
		}

		count, err := strconv.Atoi(result.Count)
		if err != nil {
			log.Error().Msgf("%s %s\n", engine.ID, engine.HostID)
		}

		if count > 0 {
			fmt.Printf("%s %s\n", engine.ID, engine.HostID)
		}
	}

	return err
}

func getEnginesByApplicationID(ctx *cli.Context, c *adp.Client) (adp.ListEntitiesResult, error) {
	req := adp.NewListEntitiesRequest(
		adp.WithListEntitiesType("singleMindServer"),
		adp.WithListEntitiesRelatedEntity(ctx.String("applicationIdentifier")),
	)

	log.Debug().Msgf("req: %+v", req)

	task := &adp.ListEntitiesTask{
		Task: adp.NewTask().WithClient(c).WithRequest(req),
	}

	result, err := task.GetResult()
	log.Debug().Msgf("result: %+v", result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func queryEngine(ctx *cli.Context, c *adp.Client, engine adp.Entity) (adp.QueryEngineResult, error) {
	var result adp.QueryEngineResult

	req := adp.NewQueryEngineRequest(
		adp.WithQueryEngineEngineName(engine.ID),
		adp.WithQueryEngineEngineTaxonomies(ctx.String("engineTaxonomies")),
		adp.WithQueryEngineEngineQuery(ctx.String("engineQuery")),
	)

	task := &adp.QueryEngineTask{
		Task: adp.NewTask().WithClient(c).WithRequest(req),
	}

	result, err := task.GetResult()
	log.Debug().Msgf("result: %+v", result)

	if err != nil {
		return result, err
	}

	return result, err
}
