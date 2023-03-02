package command

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"opentext.com/axcelerate/adp/client"
	"opentext.com/axcelerate/adp/task"
)

var (
	/*
		ListEntitiesCmd
		- get information by identifier
		  listEntities --id documentHold.APP001
		- get running eca applications under workspace Production
		  listEntities --status Running --workspace Production --type documentHold
		- get singleMindServers under application
		  listEntities --relatedEntity documentHold.APP001 --type singleMindServer
		- list only id and hostname of eca applications
		  listEntities --whiteList id,hostname --type documentHold
	*/
	ListEntitiesCmd = &cli.Command{
		Name:    "listEntities",
		Usage:   `adp-cli -p * listEntities --type dataSource`,
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

	/*
		QueryEngineCmd
	*/
	QueryEngineCmd = &cli.Command{
		Name:    "queryEngine",
		Usage:   `adp-cli -p * queryEngine --engineTaxonomies csv_guts_datatype=docs --applicationIdentifier documentHold.G0000 --engineQuery test`,
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

	/*
		TaxonomyStatisticCmd
		- get category properties (only singleMindServer allowed when deal with eca)
		  taxonomyStatistic --targetTaxonomy rm_loadbatch --engineName singleMindServer.APP001 --listCategoryProperties true
	*/
	TaxonomyStatisticCmd = &cli.Command{
		Name:    "taxonomyStatistic",
		Usage:   `adp-cli -p * taxonomyStatistic --engineTaxonomies csv_guts_datatype=docs --targetTaxonomy rm_loadbatch --applicationIdentifier documentHold.APP001`,
		Aliases: []string{"ts"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			EngineTaxonomies,
			EngineName,
			TargetTaxonomy,
			ListCategoryProperties,
			ComputeCounts,
			EngineUserName,
			EngineUserPassword,
		},
		Action: executeTask,
	}

	ComputeCountsCmd = &cli.Command{
		Name:    "computeCounts",
		Usage:   `adp-cli -p * computeCounts --applicationIdentifier documentHold.APP001 --queryParts csv_guts_datatype=email;cvs_guts_datatype=docs`,
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

	// QueryPostgresqlDBCmd ...
	QueryPostgresqlDBCmd = &cli.Command{
		Name:    "queryPostgresqlDB",
		Usage:   `adp-cli -p * queryPostgresqlDB -dbUser postgres -dbPassword * -dbConnectionURL jdbc:postgresql://dev1-db0.axcelerate.local:5432/db -sqlQuery "select ..."`,
		Aliases: []string{"qp"},
		Flags: []cli.Flag{
			DbUser,
			DbPassword,
			DbConnectionURL,
			SQLQuery,
			SizeLimitMB,
		},
		Action: executeTask,
	}

	// StopProcessesCmd ...
	StopProcessesCmd = &cli.Command{
		Name:    "stopProcesses",
		Usage:   `adp-cli -p * stopProcesses -processIdentifiers documentHold.APP001;axcelerate.APP001_Review"`,
		Aliases: []string{"sp"},
		Flags: []cli.Flag{
			ProcessIdentifiers,
		},
		Action: executeTask,
	}

	// StartApplicationCmd ...
	StartApplicationCmd = &cli.Command{
		Name:    "startApplication",
		Usage:   `adp-cli -p * startApplication -applicationIdentifier documentHold.APP001"`,
		Aliases: []string{"sa"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			ApplicationURL,
		},
		Action: executeTask,
	}

	// RemoveProcessesCmd ...
	RemoveProcessesCmd = &cli.Command{
		Name:    "removeProcesses",
		Usage:   `adp-cli -p * removeProcesses -processIdentifiers documentHold.APP001"`,
		Aliases: []string{"rp"},
		Flags: []cli.Flag{
			ProcessIdentifiers,
		},
		Action: executeTask,
	}

	StatusAndProgressCmd = &cli.Command{
		Name:    "statusAndProgress",
		Usage:   `go-adp -p * statusAndProgress --executionID ***`,
		Aliases: []string{"sp"},
		Flags: []cli.Flag{
			ExecutionID,
		},
		Action: executeTask,
	}

	Commands = []*cli.Command{
		ComputeCountsCmd,
		ListEntitiesCmd,
		QueryEngineCmd,
		QueryPostgresqlDBCmd,
		RemoveProcessesCmd,
		StopProcessesCmd,
		StartApplicationCmd,
		StatusAndProgressCmd,
		TaxonomyStatisticCmd,
	}
)

func NewTask(c *cli.Context) task.Tasker {
	var adp task.Tasker

	client := client.NewClient(
		client.WithHost(c.String("host")),
		client.WithUser(c.String("user")),
		client.WithPassword(c.String("password")),
		client.WithTaskAccessKey(c.String("taskAccessKey")),
	)

	switch c.Command.Name {
	case "listEntities":
		var opts []func(*task.ListEntitiesConfiguration)
		opts = append(opts,
			task.WithListEntitiesLoggingEnabled(false),
			task.WithListEntitiesExecutionPersistent(false),
			task.WithListEntitiesID(c.String("id")),
			task.WithListEntitiesRelatedEntity(c.String("relatedEntity")),
			task.WithListEntitiesType(c.String("type")),
			task.WithListEntitiesWhiteList(c.String("whiteList")),
			task.WithListEntitiesStatus(c.String("status")),
			task.WithListEntitiesWorkspace(c.String("workspace")),
		)

		adp = task.NewListEntitiesTask(
			client,
			c.Bool("async"),
			opts...,
		)
	case "queryEngine":
		var opts []func(*task.QueryEngineConfiguration)
		opts = append(opts,
			task.WithQueryEngineLoggingEnabled(false),
			task.WithQueryEngineExecutionPersistent(false),
			task.WithQueryEngineEngineTaxonomies(c.String("engineTaxonomies")),
			task.WithQueryEngineAdvancedRestrictions(c.String("advancedRestrictions")),
			task.WithQueryEngineEngineName(c.String("engineName")),
			task.WithQueryEngineEngineQuery(c.String("engineQuery")),
			task.WithQueryEngineEngineUserName(c.String("engineUserName")),
			task.WithQueryEngineEngineUserPassword(c.String("engineUserPassword")),
			task.WithQueryEngineApplicationIdentifier(c.String("applicationIdentifier")),
		)

		adp = task.NewQueryEngineTask(
			client,
			c.Bool("async"),
			opts...,
		)
	case "taxonomyStatistic":
		var opts []func(*task.TaxonomyStatisticConfiguration)
		opts = append(opts,
			task.WithTaxonomyStatisticLoggingEnabled(false),
			task.WithTaxonomyStatisticExecutionPersistent(false),
			task.WithTaxonomyStatisticEngineTaxonomies(c.String("engineTaxonomies")),
			task.WithTaxonomyStatisticEngineName(c.String("engineName")),
			task.WithTaxonomyStatisticOutputTaxonomies(c.String("targetTaxonomy")),
			task.WithTaxonomyStatisticComputeCounts(c.String("computeCounts")),
			task.WithTaxonomyStatisticListCategoryProperties(c.String("listCategoryProperties")),
			task.WithTaxonomyStatisticApplicationIdentifier(c.String("applicationIdentifier")),
			task.WithTaxonomyStatisticEngineUserName(c.String("engineUserName")),
			task.WithTaxonomyStatisticEngineUserPassword(c.String("engineUserPassword")),
		)

		adp = task.NewTaxonomyStatisticTask(
			client,
			c.Bool("async"),
			opts...,
		)
	case "computeCounts":
		var opts []func(*task.ComputeCountsConfiguration)
		opts = append(opts,
			task.WithComputeCountsLoggingEnabled(false),
			task.WithCommputeCountsExecutionPersistent(false),
			task.WithComputeCountsEngineTaxonomies(c.String("engineTaxonomies")),
			task.WithComputeCountsEngineName(c.String("engineName")),
			task.WithComputeCountsApplicationIdentifier(c.String("applicationIdentifier")),
			task.WithComputeCountsEngineUserName(c.String("engineUserName")),
			task.WithComputeCountsEngineUserPassword(c.String("engineUserPassword")),
			task.WithComputeCountsQueryParts(c.String("queryParts")),
		)

		adp = task.NewComputeCountsTask(
			client,
			c.Bool("async"),
			opts...,
		)
	case "stopProcesses":
		var opts []func(*task.StopProcessesConfiguration)
		if c.Bool("async") {
			opts = append(opts,
				task.WithStopProcessesLoggingEnabled(true),
				task.WithStopProcessesExecutionPersistent(true),
			)
		}
		opts = append(opts, task.WithStopProcessesProcessIdentifiers(c.String("processIdentifiers")))

		adp = task.NewStopProcessesTask(
			client,
			c.Bool("async"),
			opts...,
		)
	case "removeProcesses":
		var opts []func(*task.RemoveProcessesConfiguration)

		if c.Bool("async") {
			opts = append(opts,
				task.WithRemoveProcessesLoggingEnabled(true),
				task.WithRemoveProcessesExecutionPersistent(true))
		}

		opts = append(opts,
			task.WithRemoveProcessesProcessIdentifiers(c.String("processIdentifiers")),
			task.WithRemoveProcessesRemoveAssociatedStorages(c.String("removeAssociatedStorage")),
			task.WithRemoveProcessesTaskActive(true),
		)

		adp = task.NewRemoveProcessesTask(
			client,
			c.Bool("async"),
			opts...,
		)
	case "startApplication":
		var opts []func(*task.StartApplicationConfiguration)

		if c.Bool("async") {
			opts = append(opts,
				task.WithStartApplicationLoggingEnabled(true),
				task.WithStartApplicationExecutionPersistent(true),
			)
		}
		opts = append(opts,
			task.WithStartApplicationApplicationIdentifier(c.String("applicationIdentifier")),
		)

		adp = task.NewStartApplicationTask(
			client,
			c.Bool("async"),
			opts...,
		)
	case "statusAndProgress":
		adp = task.NewStatusAndProgressTask(client, c.String("executionID"))
	default:
		log.Fatal().Msgf("invalid ADP task name: ", c.Command.Name)
	}

	return adp
}

func executeTask(c *cli.Context) error {
	adp := NewTask(c)

	if c.Bool("async") {
		adp.SetAsync()
	}

	s, err := adp.StringOutput()
	if err != nil {
		return err
	}

	if c.Bool("pretty") {
		s = task.Prettify(s)
	}

	fmt.Println(s)

	return nil
}
