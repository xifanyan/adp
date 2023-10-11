package main

import (
	"github.com/urfave/cli/v2"
	adp "github.com/xifanyan/adp"
)

// commands
var (
	EngineDistributionCmd = &cli.Command{
		Name:    "engineDistribution",
		Usage:   `axcutil -p * engineDistribution --applicationIdentifier APPID --engineTaxonomies rm_loadbatch=Google;csv_guts_datatype!=docs`,
		Aliases: []string{"ed"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			EngineName,
			EngineTaxonomies,
			EngineQuery,
			EngineUserName,
			EngineUserPassword,
		},
		Action: engineDistribution,
	}

	ExportTaxonomyCmd = &cli.Command{
		Name:    "exportTaxonomy",
		Usage:   `axcutil -p * exportTaxonomy --engineName singleMindServer.APP001 --outputTaxonomies rm_saved_search --ListCategoryProperties true --ComputeCounts true`,
		Aliases: []string{"et"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			EngineName,
			OutputTaxonomies,
			EngineQuery,
			ListCategoryProperties,
			ComputeCounts,
			EngineUserName,
			EngineUserPassword,
		},
		Action: exportTaxonomy,
	}

	Commands = []*cli.Command{
		EngineDistributionCmd,
		ExportTaxonomyCmd,
	}
)

// newClient creates a new ADP client with the provided context.
func newClient(ctx *cli.Context) *adp.Client {
	// Default values for client are defined when setting up the context with cli package.
	client := adp.NewClient().
		WithHost(ctx.String("host")).
		WithPort(ctx.Int("port")).
		WithUser(ctx.String("user")).
		WithPassword(ctx.String("password"))

	if ctx.String("taskAccessKey") != "" {
		client = client.WithTaskAccessKey(ctx.String("taskAccessKey"))
	}

	return client
}
