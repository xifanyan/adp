package main

import (
	"flag"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/xifanyan/adp"
)

var (
	domain   = flag.String("domain", "localhost", "Domain to connect to.")
	port     = flag.Int("port", 8443, "Port to connect to.")
	user     = flag.String("user", "adpuser", "ADP User.")
	password = flag.String("password", "", "ADP User Password.")
	debug    = flag.Bool("debug", false, "sets log level to debug")
)

func main() {
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		// zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	client := adp.NewClientBuilder().
		WithDomain(*domain).WithPort(*port).
		WithUser(*user).WithPassword(*password).
		Build()

	svc := adp.Service{
		ADPClient: client,
	}

	documentHolds, err := svc.ListDocumentHoldsByUser("user1")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", documentHolds)

}
