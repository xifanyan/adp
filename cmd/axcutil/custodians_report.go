package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/urfave/cli/v2"
	"github.com/xifanyan/adp/pkg/adp"
)

func ensureEnginesAreRunning(engines adp.ListEntitiesResult) bool {
	for _, engine := range engines {
		if engine.ProcessStatus != "Running" {
			return false
		}
	}
	return true
}

func custodiansReport(ctx *cli.Context) error {

	c := adp.NewClientBuilder().
		WithDomain("localhost").
		WithPort(443).
		WithUser("adpuser").
		WithPassword("adpus3r").
		Build()

	engines, err := getEnginesByApplicationID(ctx, c)
	if err != nil {
		return err
	}

	if !ensureEnginesAreRunning(engines) {
		return fmt.Errorf("not all engines are running")
	}

	appID := ctx.String("applicationIdentifier")

	matters, err := getMatters(c, appID)
	if err != nil {
		return err
	}

	for _, matter := range matters {
		for _, category := range matter.Category {
			matterName := category.DisplayName
			custodians, err := getCustodiansByMatterID(c, appID, category.ID)
			if err != nil {
				return err
			}
			for _, custodian := range custodians {
				for _, c := range custodian.Category {
					custodianName := c.DisplayName
					if c.Count > 0 {
						fmt.Printf("%s\t%s\t%d\n", matterName, custodianName, c.Count)
					}
				}
			}
		}
	}

	return nil
}

func getMatters(c *adp.Client, appID string) (adp.TaxonomyStatisticResult, error) {
	req := adp.NewTaxonomyStatisticRequest(
		adp.WithTaxonomyStatisticApplicationIdentifier(appID),
		adp.WithTaxonomyStatisticOutputTaxonomies("rm_document_hold"),
	)

	log.Debug().Msgf("req: %+v", req)

	task := &adp.TaxonomyStatisticTask{
		Task: adp.NewTask().WithClient(c).WithRequest(req),
	}

	result, err := task.GetResult()
	// log.Debug().Msgf("result: %+v", result)

	return result, err
}

func getCustodiansByMatterID(c *adp.Client, appID string, matterID string) (adp.TaxonomyStatisticResult, error) {
	s := fmt.Sprintf("rm_document_hold=%s", matterID)
	req := adp.NewTaxonomyStatisticRequest(
		adp.WithTaxonomyStatisticApplicationIdentifier(appID),
		adp.WithTaxonomyStatisticEngineTaxonomies(s),
		adp.WithTaxonomyStatisticOutputTaxonomies("rm_custodian"),
	)

	log.Debug().Msgf("req: %+v", req)

	task := &adp.TaxonomyStatisticTask{
		Task: adp.NewTask().WithClient(c).WithRequest(req),
	}

	result, err := task.GetResult()
	// log.Debug().Msgf("result: %+v", result)

	return result, err
}
