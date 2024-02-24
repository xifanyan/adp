package main

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/urfave/cli/v2"
	adp "github.com/xifanyan/adp"
)

// documentDistribution displays document distrubution among engines based on search criteria.
//
// ctx: the cli.Context object containing the command-line arguments and flags.
// Returns: an error if there was an issue retrieving engines or querying the engine.
func documentDistribution(ctx *cli.Context) error {

	c := newClient(ctx)
	engines, err := getEnginesByApplicationID(ctx, c)
	if err != nil {
		return err
	}

	for _, engine := range engines {

		if engine.ProcessStatus != "Running" {
			log.Error().Msgf("%s is not in running state\n", engine.ID)
			continue
		}

		log.Debug().Msgf("engine %s running on %s\n", engine.ID, engine.HostID)

		result, err := queryEngine(ctx, c, engine)
		if err != nil {
			log.Error().Msgf("%s %s\n", engine.ID, engine.HostID)
			continue
		}

		count, err := strconv.Atoi(result.Count)
		if err != nil {
			log.Error().Msgf("%s: failed to convert count %v to integer", engine.ID, result.Count)
			continue
		}

		log.Debug().Msgf("%s: doc count %d\n", engine.ID, count)

		if count > 0 {
			fmt.Printf("%s|%s|%d\n", engine.ID, engine.HostID, count)
		}
	}

	return nil
}

// getEnginesByApplicationID retrieves a list of singleMindServer engines
// related to the given application identifier.
//
// Parameters:
// - ctx: The cli.Context object that contains the command line arguments.
// - c: The adp.Client object used for making API requests.
//
// Returns:
// - adp.ListEntitiesResult: The result of the list entities operation.
// - error: An error if the operation fails.
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

	return result, err
}

// queryEngine is a function that performs a query on the provided engine using the given context and client.
//
// It takes in the following parameters:
// - ctx: a pointer to a cli.Context object that contains the context for the function.
// - c: a pointer to an adp.Client object that represents the client for the function.
// - engine: an adp.Entity object that represents the engine on which the query is performed.
//
// It returns an adp.QueryEngineResult object and an error.
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

	return result, err
}
