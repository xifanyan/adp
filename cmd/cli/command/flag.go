package command

import (
	"github.com/urfave/cli/v2"
)

var (
	// EngineUserName ...
	EngineUserName = &cli.StringFlag{
		Name:  "engineUserName",
		Usage: "Engine User Name, e.g., docs-pull-scripts",
	}

	// EngineUserPassword ...
	EngineUserPassword = &cli.StringFlag{
		Name:  "engineUserPassword",
		Usage: "Engine User Password",
	}

	// ID ...
	ID = &cli.StringFlag{
		Name:  "id",
		Usage: "Entity ID e.g., singleMindServer.APP001",
	}

	// RelatedEntity ...
	RelatedEntity = &cli.StringFlag{
		Name:  "relatedEntity",
		Usage: "Related Entity ID e.g., documentHold.APP001",
	}

	// WhiteList ...
	WhiteList = &cli.StringFlag{
		Name:  "whiteList",
		Usage: "Whitelisted Field Names",
		Value: "id,displayName,processStatus,hostName",
	}

	// Type ...
	Type = &cli.StringFlag{
		Name:  "type",
		Usage: "Entity Type e.g., singleMindServer, mergingMeta",
	}

	// Status ...
	Status = &cli.StringFlag{
		Name:  "status",
		Usage: "Status of the process e.g., Running",
	}

	// Workspace ...
	Workspace = &cli.StringFlag{
		Name:  "workspace",
		Usage: "Name of workspace e.g., Production",
	}

	// ApplicationIdentifier ...
	ApplicationIdentifier = &cli.StringFlag{
		Name:  "applicationIdentifier",
		Usage: "Application Identifier e.g., documentHold.APP001",
	}

	// EngineName ...
	EngineName = &cli.StringFlag{
		Name:  "engineName",
		Usage: "Engine Identifier e.g., singleMindServer.APP001",
	}

	// EngineTaxonomies ...
	EngineTaxonomies = &cli.StringFlag{
		Name:  "engineTaxonomies",
		Usage: "Engine Taxonomies e.g., rm_loadbatch=Google,csv_guts_datatype!=docs",
	}

	// EngineQuery ...
	EngineQuery = &cli.StringFlag{
		Name:  "engineQuery",
		Usage: "Engine Query",
		Value: "*",
	}

	// AdvancedRestrictions ...
	AdvancedRestrictions = &cli.StringFlag{
		Name:  "advancedRestrictions",
		Usage: "Advanced Restrictions",
	}

	// TargetTaxonomy ...
	TargetTaxonomy = &cli.StringFlag{
		Name:  "targetTaxonomy",
		Usage: "Target Taxonomy e.g., rm_loadbatch",
	}

	// ListCategoryProperties ...
	ListCategoryProperties = &cli.StringFlag{
		Name:  "listCategoryProperties",
		Usage: "List Category Properties",
		Value: "false",
	}

	// ComputeCounts ...
	ComputeCounts = &cli.StringFlag{
		Name:  "computeCounts",
		Usage: "Compute Counts",
		Value: "true",
	}

	// QueryParts ...
	QueryParts = &cli.StringFlag{
		Name:  "queryParts",
		Usage: "Query Parts",
	}

	// Execution ID...
	ExecutionID = &cli.StringFlag{
		Name:  "executionID",
		Usage: "Execution ID",
	}

	// DbUser ...
	DbUser = &cli.StringFlag{
		Name:  "dbUser",
		Usage: "Database User",
		Value: "postgres",
	}

	// DbPassword ...
	DbPassword = &cli.StringFlag{
		Name:  "dbPassword",
		Usage: "Database User Password",
	}

	// DbConnectionURL ...
	DbConnectionURL = &cli.StringFlag{
		Name:  "dbConnectionURL",
		Usage: "DSN e.g., postgresql://dev0-db0.axcelerate.local:5432/db",
	}

	// SQLQuery ...
	SQLQuery = &cli.StringFlag{
		Name:  "sqlQuery",
		Usage: "Query e.g., select * from ...",
	}

	// SizeLimitMB ...
	SizeLimitMB = &cli.StringFlag{
		Name:  "sizeLimitMB",
		Usage: "Memory Limit in Mb e.g., 5",
	}

	// Identifiers ...
	Identifiers = &cli.StringFlag{
		Name:  "identifiers",
		Usage: "List of Application or Engine Identifiers e.g., documentHold.APP001;singleMindServer.APP001",
	}

	// ApplicationIdentifiers ...
	ApplicationIdentifiers = &cli.StringFlag{
		Name:  "applicationIdentifiers",
		Usage: "List of Application Identifier e.g., documentHold.APP001;documentHold.APP001",
	}

	// EngineIdentifiers ...
	EngineIdentifiers = &cli.StringFlag{
		Name:  "engineIdentifiers",
		Usage: "List of Engine Identifier e.g., singleMindServer.APP001;mergingMeta.APP001",
	}

	// ApplicationURL ...
	ApplicationURL = &cli.StringFlag{
		Name:  "applicationURL",
		Usage: "Variable name for application URL output",
		Value: "appURL",
	}

	// ProcessIdentifiers ...
	ProcessIdentifiers = &cli.StringFlag{
		Name:  "processIdentifiers",
		Usage: "List of Identifiers e.g., documentHold.APP001;singleMindServer.APP002",
	}
)
