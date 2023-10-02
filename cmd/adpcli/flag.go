package main

import (
	"github.com/urfave/cli/v2"
)

var (
	EngineUserName = &cli.StringFlag{
		Name:  "engineUserName",
		Usage: "Engine User Name, e.g., docs-pull-scripts",
	}

	EngineUserPassword = &cli.StringFlag{
		Name:  "engineUserPassword",
		Usage: "Engine User Password",
	}

	ID = &cli.StringFlag{
		Name:  "id",
		Usage: "Entity ID e.g., singleMindServer.APP001",
	}

	RelatedEntity = &cli.StringFlag{
		Name:  "relatedEntity",
		Usage: "Related Entity ID e.g., documentHold.APP001",
	}

	WhiteList = &cli.StringFlag{
		Name:  "whiteList",
		Usage: "Whitelisted Field Names",
		Value: "id,displayName,processStatus,hostName",
	}

	Type = &cli.StringFlag{
		Name:  "type",
		Usage: "Entity Type e.g., singleMindServer, mergingMeta",
	}

	Status = &cli.StringFlag{
		Name:  "status",
		Usage: "Status of the process e.g., Running",
	}

	Workspace = &cli.StringFlag{
		Name:  "workspace",
		Usage: "Name of workspace e.g., Production",
	}

	ApplicationIdentifier = &cli.StringFlag{
		Name:  "applicationIdentifier",
		Usage: "Application Identifier e.g., documentHold.APP001",
	}

	EngineName = &cli.StringFlag{
		Name:  "engineName",
		Usage: "Engine Identifier e.g., singleMindServer.APP001",
	}

	EngineTaxonomies = &cli.StringFlag{
		Name:  "engineTaxonomies",
		Usage: "Engine Taxonomies e.g., rm_loadbatch=Google,csv_guts_datatype!=docs",
	}

	EngineQuery = &cli.StringFlag{
		Name:  "engineQuery",
		Usage: "Engine Query (default: *)",
		Value: "*",
	}

	AdvancedRestrictions = &cli.StringFlag{
		Name:  "advancedRestrictions",
		Usage: "Advanced Restrictions",
	}

	OutputTaxonomies = &cli.StringFlag{
		Name:  "outputTaxonomies",
		Usage: "Output Taxonomies e.g., rm_loadbatch",
	}

	ListCategoryProperties = &cli.StringFlag{
		Name:  "listCategoryProperties",
		Usage: "List Category Properties (default: \"false\")",
		Value: "false",
	}

	ComputeCounts = &cli.StringFlag{
		Name:  "computeCounts",
		Usage: "Compute Counts",
		Value: "true",
	}

	QueryParts = &cli.StringFlag{
		Name:  "queryParts",
		Usage: "Query Parts",
	}

	ExecutionID = &cli.StringFlag{
		Name:  "executionID",
		Usage: "Execution ID",
	}

	ClientCertPath = &cli.StringFlag{
		Name:  "clientCertPath",
		Usage: "Client Cert Path",
		Value: "S:/Projects/ProcessControl/security/postgres/client.billing.crt",
	}

	ClientKeyPath = &cli.StringFlag{
		Name:  "clientKeyPath",
		Usage: "Client Key Path",
		Value: "S:/Projects/ProcessControl/security/postgres/client.billing.key",
	}

	RootCertPath = &cli.StringFlag{
		Name:  "rootCertPath",
		Usage: "Root Cert Path",
		Value: "S:/Projects/ProcessControl/security/postgres/root.billing.crt",
	}

	DbUser = &cli.StringFlag{
		Name:  "dbUser",
		Usage: "Database User (default: postgres)",
		Value: "postgres",
	}

	DbPassword = &cli.StringFlag{
		Name:  "dbPassword",
		Usage: "Database User Password",
	}

	DbConnectionURL = &cli.StringFlag{
		Name:  "dbConnectionURL",
		Usage: "DSN e.g., postgresql://dev0-db0.axcelerate.local:5432/db",
	}

	SQLQuery = &cli.StringFlag{
		Name:  "sqlQuery",
		Usage: "Query e.g., select * from ...",
	}

	SizeLimitMB = &cli.StringFlag{
		Name:  "sizeLimitMB",
		Usage: "Memory Limit in Mb e.g., 5",
	}

	Identifiers = &cli.StringFlag{
		Name:  "identifiers",
		Usage: "List of Application or Engine Identifiers e.g., documentHold.APP001;singleMindServer.APP001",
	}

	ApplicationIdentifiers = &cli.StringFlag{
		Name:  "applicationIdentifiers",
		Usage: "List of Application Identifier e.g., documentHold.APP001;documentHold.APP001",
	}

	EngineIdentifiers = &cli.StringFlag{
		Name:  "engineIdentifiers",
		Usage: "List of Engine Identifier e.g., singleMindServer.APP001;mergingMeta.APP001",
	}

	ApplicationURL = &cli.StringFlag{
		Name:  "applicationURL",
		Usage: "Variable name for application URL output",
		Value: "appURL",
	}

	ProcessIdentifiers = &cli.StringFlag{
		Name:  "processIdentifiers",
		Usage: "List of Identifiers e.g., documentHold.APP001;singleMindServer.APP002",
	}

	RemoveAssociatedStorage = &cli.StringFlag{
		Name:  "removeAssociatedStorage",
		Usage: "RemoveAssociatedStorage (default: \"true\")",
		Value: "true",
	}
)
