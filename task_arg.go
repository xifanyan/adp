package adp

import "strings"

// Delimiters used for parsing Args.
const (
	ARGDELIMITER              = ";"
	NE                        = "!="
	EQ                        = "="
	MAX_NUMBER_OF_CATEGORIES  = 1000
	Engine_BOX_DOC_THREADHOLD = "10000000"
)

// EngineTaxonomyArg represents data structure for category/taxonmy query statement,
// e.g., csv_guts_datatype=email.
type EngineTaxonomyArg struct {
	Taxonomy string `json:"Taxonomy"`
	Negation bool   `json:"Negation"`
	Query    string `json:"Query"`
}

func ParseEngineTaxonomiesArgs(s string) []EngineTaxonomyArg {
	var engineTaxonomyArgs []EngineTaxonomyArg

	for _, taxonomy := range strings.Split(s, ARGDELIMITER) {
		var fields []string
		var negation bool

		if strings.Contains(taxonomy, NE) {
			negation = true
			fields = strings.SplitN(taxonomy, NE, 2)
		} else {
			fields = strings.SplitN(taxonomy, EQ, 2)
		}

		if len(fields) == 2 {
			// remove leading and trailing spaces of each field
			tax := strings.TrimSpace(fields[0])
			qry := strings.TrimSpace(fields[1])

			if len(tax) > 0 {
				engineTaxonomyArgs = append(engineTaxonomyArgs,
					EngineTaxonomyArg{
						Taxonomy: tax,
						Negation: negation,
						Query:    qry,
					})
			}
		}
	}

	return engineTaxonomyArgs
}

// AdvancedRestrictionsArg ...
type AdvancedRestrictionsArg struct {
	AdvancedSearchField string `json:"Advanced Search Field"`
	Negation            bool   `json:"Negation"`
	FuzzySearch         bool   `json:"Fuzzy Search"`
	AdvancedExpression  string `json:"Advanced Expression"`
}

func ParseAdvancedRestrictionsArg(s string) []AdvancedRestrictionsArg {
	var advancedRestrictionsArgs []AdvancedRestrictionsArg

	for _, restriction := range strings.Split(s, ARGDELIMITER) {
		var fields []string
		var negation bool

		if strings.Contains(restriction, NE) {
			negation = true
			fields = strings.SplitN(restriction, NE, 2)
		} else {
			fields = strings.SplitN(restriction, EQ, 2)
		}

		if len(fields) == 2 {
			fld := strings.TrimSpace(fields[0])
			exp := strings.TrimSpace(fields[1])

			if len(fld) > 0 {
				advancedRestrictionsArgs = append(advancedRestrictionsArgs,
					AdvancedRestrictionsArg{
						AdvancedSearchField: fld,
						Negation:            negation,
						FuzzySearch:         false,
						AdvancedExpression:  exp,
					})
			}
		}
	}
	return advancedRestrictionsArgs
}

func parseQueryParts(s string) []string {
	return strings.Split(s, ARGDELIMITER)
}

// OutputTaxonomiesArg ...
type OutputTaxonomiesArg struct {
	Taxonomy                  string `json:"Taxonomy"`
	Mode                      string `json:"Mode"`
	MaximumNumberOfCategories int    `json:"Maximum number of categories"`
}

func ParseOutputTaxonomiesArg(s string) []OutputTaxonomiesArg {
	var outputTaxonomiesArgs []OutputTaxonomiesArg

	for _, taxonomy := range strings.Split(s, ARGDELIMITER) {
		outputTaxonomiesArgs = append(outputTaxonomiesArgs,
			OutputTaxonomiesArg{
				Taxonomy:                  strings.TrimSpace(taxonomy),
				Mode:                      "Category counts",
				MaximumNumberOfCategories: MAX_NUMBER_OF_CATEGORIES,
			})
	}

	return outputTaxonomiesArgs
}

// StopProcessesProcessIdentifierArg ...
type StopProcessesProcessIdentifierArg struct {
	ProcessIdenifier string `json:"Process identifier"`
	StopRecursive    bool   `json:"Stop recursive"`
}

func ParseStopProcessesProcessIdentifiersArg(s string) []StopProcessesProcessIdentifierArg {
	var processIdentifierArgs []StopProcessesProcessIdentifierArg
	for _, proc := range strings.Split(s, ARGDELIMITER) {
		proc = strings.TrimSpace(proc)
		if len(proc) > 0 {
			processIdentifierArgs = append(processIdentifierArgs,
				StopProcessesProcessIdentifierArg{
					ProcessIdenifier: proc,
					StopRecursive:    true, // always recursive
				},
			)
		}
	}
	return processIdentifierArgs
}

type ConfigTableMapsArg struct {
	Action       string `json:"Action"`
	Column       string `json:"Column"`
	Row          int    `json:"Row"`
	Substitution string `json:"Substitution"`
	TableName    string `json:"Table name"`
	Value        string `json:"Value"`
}

type ConfigTableMapsArgs []ConfigTableMapsArg

type CliParameterArg struct {
	Parameter string `json:"Parameter"`
}

type FieldMappingArg struct {
	CSVFieldName   string `json:"CSV Field Name"`
	TextType       string `json:"Text type"`
	ValueDelimiter string `json:"Value Delimiter"`
}

// example #1: s = "dataModel.G00000;export_base_directory,printing.storageDirectory"
// example #2: s = "dataModel.G00000;ax_coded_qc_change,ax_coded_pl_do_not_log;x"
//
// {
//   ConfigurationID : "dataModel.G00000"
//   DynamicComponentNames : "export_base_directory,printing.storageDirectory"
//   FieldList : "name,value,path,host"
//   NameValueList : "ax_coded_qc_change,ax_coded_pl_do_not_log"
// }
type ConfigurationArg struct {
	ConfigurationID       string `json:"Configuration ID"`
	DynamicComponentNames string `json:"Dynamic Component Names"`
	FieldList             string `json:"Field list"`
	NameValueList         string `json:"Name value list"`
	ApplicationType       string `json:"Application type"`
	EntityType            string `json:"Entity type"`
}

// Configure Entity Args
type SimpleConfigParamArg struct {
	Value        string `json:"Value"`
	Substitution string `json:"Substitution"`
	Parameter    string `json:"Parameter"`
}

type SimpleConfigParamArgs []SimpleConfigParamArg

type ConfigTableArg struct {
	Action       string `json:"Action"`
	Column       string `json:"Column"`
	Row          int    `json:"Row"`
	Substitution string `json:"Substitution"`
	TableName    string `json:"Table name"`
	Value        string `json:"Value"`
}
type ConfigTableArgs []ConfigTableArg

type ConfigStringListArg struct {
	Value        string `json:"Value"`
	Substitution string `json:"Substitution"`
	ListName     string `json:"List name"`
	Row          int    `json:"Row"`
	Action       string `json:"Action"`
}
type ConfigStringListArgs []ConfigStringListArg

type DynamicComponentsArg struct {
	Type       string `json:"Type"`
	Identifier string `json:"Identifier"`
	Position   int    `json:"Position"`
	Action     string `json:"Action"`
}
type DynamicComponentsArgs []DynamicComponentsArg

type DynamicComponentsSimpleConfigParamArg struct {
	Type       string `json:"Type"`
	Identifier string `json:"Identifier"`
	SimpleConfigParamArg
}
type DynamicComponentsSimpleConfigParamArgs []DynamicComponentsSimpleConfigParamArg

type DynamicComponentsConfigTableArg struct {
	Type       string `json:"Type"`
	Identifier string `json:"Identifier"`
	ConfigTableArg
}
type DynamicComponentsConfigTableArgs []DynamicComponentsConfigTableArg

type DynamicComponentsConfigStringListArg struct {
	Type       string `json:"Type"`
	Identifier string `json:"Identifier"`
	ConfigStringListArg
}
type DynamicComponentsConfigStringListArgs []DynamicComponentsConfigStringListArg

// ActionOnEntityArg:  Task "Configure Service Properties"
type ActionOnEntityArg struct {
	Identifier string `json:"Identifier"`
	Action     string `json:"Action"`
	Value      string `json:"Value"`
}

/*
type RemoveProcessesProcessIdentifierArg struct {
	ProcessIdentifier   string `json:"Process identifier"`
	RemoveRecursive     bool   `json:"Remove recursive"`
	PreserveFiles       bool   `json:"Preserve files"`
	SourceFilesLocation string `json:"Source files location"`
	TargetLocation      string `json:"Target location"`
}

func parseRemoveProcessesProcessIdentifierArgs(s string) []RemoveProcessesProcessIdentifierArg {
	var removeProcessesProcessIdentifierArgs []RemoveProcessesProcessIdentifierArg
	for _, proc := range strings.Split(s, ARGDELIMITER) {
		removeProcessesProcessIdentifierArgs = append(removeProcessesProcessIdentifierArgs,
			RemoveProcessesProcessIdentifierArg{
				ProcessIdentifier:   proc,
				RemoveRecursive:     true,
				PreserveFiles:       false,
				SourceFilesLocation: "",
				TargetLocation:      "",
			})
	}
	return removeProcessesProcessIdentifierArgs
}
*/
