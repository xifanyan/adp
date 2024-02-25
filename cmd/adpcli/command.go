package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/xifanyan/adp/pkg/adp"
)

var (
	ListEntitiesCmd = &cli.Command{
		Name:    "listEntities",
		Usage:   `adpcli -p * listEntities --type dataSource`,
		Aliases: []string{"le"},
		Flags: []cli.Flag{
			ID,
			RelatedEntity,
			WhiteList,
			Type,
			Status,
			Workspace,
		},
		Action: executeTask,
	}

	QueryEngineCmd = &cli.Command{
		Name:    "queryEngine",
		Usage:   `adpcli -p * queryEngine --engineTaxonomies csv_guts_datatype=docs --applicationIdentifier documentHold.G0000 --engineQuery test`,
		Aliases: []string{"qe"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			EngineName,
			EngineTaxonomies,
			EngineQuery,
			EngineUserName,
			EngineUserPassword,
		},
		Action: executeTask,
	}

	TaxonomyStatisticCmd = &cli.Command{
		Name:    "taxonomyStatistic",
		Usage:   `adpcli -p * taxonomyStatistic --engineTaxonomies csv_guts_datatype=docs --targetTaxonomy rm_loadbatch --applicationIdentifier documentHold.APP001`,
		Aliases: []string{"ts"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			EngineTaxonomies,
			EngineName,
			OutputTaxonomies,
			ListCategoryProperties,
			ComputeCounts,
			EngineUserName,
			EngineUserPassword,
		},
		Action: executeTask,
	}

	ComputeCountsCmd = &cli.Command{
		Name:    "computeCounts",
		Usage:   `adpcli -p * computeCounts --applicationIdentifier documentHold.APP001 --queryParts csv_guts_datatype=email;cvs_guts_datatype=docs`,
		Aliases: []string{"cc"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			EngineTaxonomies,
			EngineName,
			EngineUserName,
			EngineUserPassword,
			EngineQuery,
			QueryParts,
		},
		Action: executeTask,
	}

	StopProcessesCmd = &cli.Command{
		Name:    "stopProcesses",
		Usage:   `adpcli -p * stopProcesses -processIdentifiers documentHold.APP001;axcelerate.APP001_Review`,
		Aliases: []string{"sp"},
		Flags: []cli.Flag{
			ProcessIdentifiers,
		},
		Action: executeTask,
	}

	PingProjectCmd = &cli.Command{
		Name:    "pingProject",
		Usage:   `adpcli -p * pingProject --identifiers documentHold.APP001,documentHold.APP002`,
		Aliases: []string{"pp"},
		Flags: []cli.Flag{
			Identifiers,
		},
		Action: executeTask,
	}

	StartApplicationCmd = &cli.Command{
		Name:    "startApplication",
		Usage:   `apdcli -p * startApplication --applicationIdentifier documentHold.APP001`,
		Aliases: []string{"sa"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			ApplicationURL,
		},
		Action: executeTask,
	}

	Commands = []*cli.Command{
		ListEntitiesCmd,
		QueryEngineCmd,
		TaxonomyStatisticCmd,
		ComputeCountsCmd,
		StopProcessesCmd,
		PingProjectCmd,
		StartApplicationCmd,
	}
)

// newClient creates a new ADP client with the provided context.
func newClient(ctx *cli.Context) *adp.Client {
	// Default values for client are defined when setting up the context with cli package.
	builder := adp.NewClientBuilder().
		WithDomain(ctx.String("domain")).
		WithPort(ctx.Int("port")).
		WithUser(ctx.String("user")).
		WithPassword(ctx.String("password"))

	if ctx.String("taskAccessKey") != "" {
		builder = builder.WithTaskAccessKey(ctx.String("taskAccessKey"))
	}

	return builder.Build()
}

func newListEntitiesRequest(ctx *cli.Context) *adp.Request {
	return adp.NewListEntitiesRequest(
		adp.WithListEntitiesLoggingEnabled(false),
		adp.WithListEntitiesExecutionPersistent(false),
		adp.WithListEntitiesID(ctx.String("id")),
		adp.WithListEntitiesRelatedEntity(ctx.String("relatedEntity")),
		adp.WithListEntitiesType(ctx.String("type")),
		adp.WithListEntitiesWhiteList(ctx.String("whiteList")),
		adp.WithListEntitiesStatus(ctx.String("status")),
		adp.WithListEntitiesWorkspace(ctx.String("workspace")),
	)
}

func newQueryEngineRequest(ctx *cli.Context) *adp.Request {
	return adp.NewQueryEngineRequest(
		adp.WithQueryEngineLoggingEnabled(false),
		adp.WithQueryEngineExecutionPersistent(false),
		adp.WithQueryEngineApplicationIdentifier(ctx.String("applicationIdentifier")),
		adp.WithQueryEngineEngineName(ctx.String("engineName")),
		adp.WithQueryEngineEngineTaxonomies(ctx.String("engineTaxonomies")),
		adp.WithQueryEngineEngineQuery(ctx.String("engineQuery")),
		adp.WithQueryEngineEngineUserName(ctx.String("engineUserName")),
		adp.WithQueryEngineEngineUserPassword(ctx.String("engineUserPassword")),
	)
}

func newTaxonomyStatisticRequest(ctx *cli.Context) *adp.Request {
	return adp.NewTaxonomyStatisticRequest(
		adp.WithTaxonomyStatisticLoggingEnabled(false),
		adp.WithTaxonomyStatisticExecutionPersistent(false),
		adp.WithTaxonomyStatisticEngineTaxonomies(ctx.String("engineTaxonomies")),
		adp.WithTaxonomyStatisticApplicationIdentifier(ctx.String("applicationIdentifier")),
		adp.WithTaxonomyStatisticEngineName(ctx.String("engineName")),
		adp.WithTaxonomyStatisticEngineUserPassword(ctx.String("engineUserPassword")),
		adp.WithTaxonomyStatisticOutputTaxonomies(ctx.String("outputTaxonomies")),
		adp.WithTaxonomyStatisticComputeCounts(ctx.String("computeCounts")),
		adp.WithTaxonomyStatisticListCategoryProperties(ctx.String("listCategoryProperties")),
	)
}

func newComputeCountsRequest(ctx *cli.Context) *adp.Request {
	return adp.NewComputeCountsRequest(
		adp.WithComputeCountsLoggingEnabled(false),
		adp.WithComputeCountsExecutionPersistent(false),
		adp.WithComputeCountsApplicationIdentifier(ctx.String("applicationIdentifier")),
		adp.WithComputeCountsEngineName(ctx.String("engineName")),
		adp.WithComputeCountsEngineUserName(ctx.String("engineUserName")),
		adp.WithComputeCountsEngineUserPassword(ctx.String("engineUserPassword")),
		adp.WithComputeCountsEngineTaxonomies(ctx.String("engineTaxonomies")),
		adp.WithComputeCountsQueryParts(ctx.String("queryParts")),
	)
}

func newStopProcessesRequest(ctx *cli.Context) *adp.Request {
	return adp.NewStopProcessesRequest(
		adp.WithStopProcessesLoggingEnabled(false),
		adp.WithStopProcessesExecutionPersistent(false),
		adp.WithStopProcessesProcessIdentifiers(ctx.String("processIdentifiers")),
	)
}

func newPingProjectRequest(ctx *cli.Context) *adp.Request {
	return adp.NewPingProjectRequest(
		adp.WithPingProjectLoggingEnabled(false),
		adp.WithPingProjectExecutionPersistent(false),
		adp.WithPingProjectIdentifiers(ctx.String("identifier")),
	)
}

func newStartApplicationRequest(ctx *cli.Context) *adp.Request {
	return adp.NewStartApplicationRequest(
		adp.WithStartApplicationLoggingEnabled(false),
		adp.WithStartApplicationExecutionPersistent(false),
		adp.WithStartApplicationApplicationIdentifier(ctx.String("applicationIdentifier")),
		adp.WithStartApplicationApplicationURL(ctx.String("applicationURL")),
	)
}

// newTask initializes and returns a tasker based on the command name in the CLI context.
func newTask(ctx *cli.Context) adp.Tasker {

	var task *adp.Task
	var tasker adp.Tasker

	client := newClient(ctx)

	switch ctx.Command.Name {
	case "listEntities":
		req := newListEntitiesRequest(ctx)
		task = adp.NewTask().WithClient(client).WithRequest(req)
		tasker = &adp.ListEntitiesTask{Task: task}
	case "queryEngine":
		req := newQueryEngineRequest(ctx)
		task = adp.NewTask().WithClient(client).WithRequest(req)
		tasker = &adp.ListEntitiesTask{Task: task}
	case "taxonomyStatistic":
		req := newTaxonomyStatisticRequest(ctx)
		task = adp.NewTask().WithClient(client).WithRequest(req)
		tasker = &adp.TaxonomyStatisticTask{Task: task}
	case "computeCounts":
		req := newComputeCountsRequest(ctx)
		task = adp.NewTask().WithClient(client).WithRequest(req)
		tasker = &adp.ComputeCountsTask{Task: task}
	case "stopProcesses":
		req := newStopProcessesRequest(ctx)
		task = adp.NewTask().WithClient(client).WithRequest(req)
		tasker = &adp.StopProcessesTask{Task: task}
	case "pingProject":
		req := newPingProjectRequest(ctx)
		task = adp.NewTask().WithClient(client).WithRequest(req)
		tasker = &adp.PingProjectTask{Task: task}
	case "startApplication":
		req := newStartApplicationRequest(ctx)
		task = adp.NewTask().WithClient(client).WithRequest(req)
		tasker = &adp.PingProjectTask{Task: task}
	}

	return tasker
}

// executeTask runs a task based on the provided command-line arguments.
func executeTask(ctx *cli.Context) error {
	task := newTask(ctx)

	s, err := task.GetResultAsString()
	if err != nil {
		return err
	}

	// Optionally prettify the output if the "pretty" flag is set.
	if ctx.Bool("pretty") {
		s = adp.Prettify(s)
	}

	fmt.Println(s)

	return nil
}
