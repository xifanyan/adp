package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"github.com/xifanyan/adp/pkg/adp"
)

func exportTaxonomy(ctx *cli.Context) error {

	c := newClient(ctx)

	req := adp.NewTaxonomyStatisticRequest(
		adp.WithTaxonomyStatisticEngineName(ctx.String("engineName")),
		adp.WithTaxonomyStatisticOutputTaxonomies(ctx.String("outputTaxonomies")),
		adp.WithTaxonomyStatisticComputeCounts(ctx.String("computeCounts")),
		adp.WithTaxonomyStatisticListCategoryProperties(ctx.String("listCategoryProperties")),
		adp.WithTaxonomyStatisticOutputTaxonomies(ctx.String("outputTaxonomies")),
	)

	task := &adp.TaxonomyStatisticTask{
		Task: adp.NewTask().WithClient(c).WithRequest(req),
	}

	res, err := task.GetResult()
	if err != nil {
		return err
	}

	log.Debug().Msgf("Result: %+v", res)

	for _, taxonomy := range res {
		i := 0
		for _, category := range taxonomy.Category {
			i++
			fmt.Printf("%d\t:%s\n", i, category.DisplayName)
		}
	}

	return nil
}
