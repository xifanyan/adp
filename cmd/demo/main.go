package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/xifanyan/adp/pkg/adp"
	"github.com/xifanyan/adp/pkg/axc"
)

const (
	NUMBER_OF_PROCESSORS = 2
)

type queryRequest struct {
	ApplicationID string
	MatterID      string
	CustodianID   string
}

func produceQueryRequests(app *axc.Application, ch chan<- queryRequest) {
	defer close(ch)

	matters, _ := app.Matters()
	for _, matter := range matters {
		custodians, _ := app.GetCustodiansByMatter(matter.ID)

		for _, custodian := range custodians {
			if custodian.Count > 0 {
				ch <- queryRequest{
					ApplicationID: app.ID,
					MatterID:      matter.ID,
					CustodianID:   custodian.ID,
				}
			}
		}

	}
}

func query(id int, app *axc.Application, ch <-chan queryRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if req, ok := <-ch; ok {
			engineTaxonomies := fmt.Sprintf("rm_document_hold=%s;rm_custodian=%s", req.MatterID, req.CustodianID)
			res, _ := app.Query(engineTaxonomies, "*")
			fmt.Printf("%s|%s|%s|%s|%s\n", req.ApplicationID, req.MatterID, req.CustodianID, res.Count, res.Size)
		} else {
			log.Debug().Msgf("thread #%d - query channel closed", id)
			return
		}
	}
}

func main() {

	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var wg sync.WaitGroup
	queryQueue := make(chan queryRequest, NUMBER_OF_PROCESSORS)

	client := adp.NewClientBuilder().
		WithDomain("localhost").
		WithPort(443).
		WithUser("adpuser").
		WithPassword("adpus3r").
		Build()

	app := axc.NewApplicationBuilder().
		WithID("documentHold.G01610").
		WithClient(client).
		Build()

	go produceQueryRequests(app, queryQueue)

	for i := 0; i < NUMBER_OF_PROCESSORS; i++ {
		wg.Add(1)
		go query(i, app, queryQueue, &wg)
	}

	wg.Wait()
}
