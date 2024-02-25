package adp

import "strings"

// Delimiters used for parsing Args.
const (
	ARGDELIMITER = ";"
	NE           = "!="
	EQ           = "="
)

// EngineTaxonomyArg represents data structure for category/taxonmy query statement,
// e.g., csv_guts_datatype=email.
type EngineTaxonomyArg struct {
	Taxonomy string `json:"Taxonomy"`
	Negation bool   `json:"Negation"`
	Query    string `json:"Query"`
}

func parseEngineTaxonomiesArgs(s string) []EngineTaxonomyArg {
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

func parseAdvancedRestrictionsArg(s string) []AdvancedRestrictionsArg {
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

func parseOutputTaxonomiesArg(s string) []OutputTaxonomiesArg {
	var outputTaxonomiesArgs []OutputTaxonomiesArg

	for _, taxonomy := range strings.Split(s, ARGDELIMITER) {
		outputTaxonomiesArgs = append(outputTaxonomiesArgs,
			OutputTaxonomiesArg{
				Taxonomy:                  strings.TrimSpace(taxonomy),
				Mode:                      "Category counts",
				MaximumNumberOfCategories: 1500,
			})
	}

	return outputTaxonomiesArgs
}

// StopProcessesProcessIdentifierArg ...
type StopProcessesProcessIdentifierArg struct {
	ProcessIdenifier string `json:"Process identifier"`
	StopRecursive    bool   `json:"Stop recursive"`
}

func parseStopProcessesProcessIdentifiersArg(s string) []StopProcessesProcessIdentifierArg {
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

/*
// RemoveProcessesProcessIdentifierArg ...
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
