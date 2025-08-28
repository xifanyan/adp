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

	/*
		Examples:
		#1: List all applications by user
		documentHolds, err := svc.ListDocumentHoldsByUser("user1")
		users, groups, err := svc.GetAllUsersAndGroups()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n%+v\n", users, groups)

		users, err = svc.GetUsersByID("demouser1")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", users)

		groups, err = svc.GetGroupsByID("demogroup1")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", groups)

		u, err := svc.GetUserByID("demouser1")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", u)

		g, err := svc.GetGroupByID("demogroup1")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", g)

		groups, err := svc.GetGroupsByUserID("demouser2")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", groups)

		user3 := adp.UserDefinition{
			UserName: "demouser3",
			Password: "demous3r3",
		}
		err = svc.AddUser([]adp.UserDefinition{user3})
		if err != nil {
			panic(err)
		}

		group3 := adp.GroupDefinition{
			GroupName: "demogroup3",
		}

		err = svc.AddGroup([]adp.GroupDefinition{group3})
		if err != nil {
			panic(err)
		}

		err = svc.AddUsersToGroup([]string{"demouser3"}, "demogroup3")
		if err != nil {
			panic(err)
		}

		users, groups, err := svc.GetUsersAndGroupsByApplicationID("documentHold.demo00001")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n---\n%+v\n", users, groups)

		res, err := svc.CreateApplication(
			adp.WithCreateApplicationApplicationType("axcelerateStandalone"),
			adp.WithCreateApplicationApplicationName("DEMO1Review"),
			adp.WithCreateApplicationApplicationTemplate("axcelerate._DEMO_Review_Template"),
			adp.WithCreateApplicationApplicationWorkspace("Workspace1"),
			adp.WithCreateApplicationApplicationHost("vm-rhauswirth2.otxlab.net"),
		)

		res, err := svc.CreateApplication(
			adp.WithCreateApplicationApplicationType("documentHold"),
			adp.WithCreateApplicationApplicationName("newDocumentHold"),
			adp.WithCreateApplicationApplicationTemplate("documentHold._Disney_Template_v1"),
			adp.WithCreateApplicationApplicationWorkspace("Workspace1"),
			adp.WithCreateApplicationApplicationHost("vm-rhauswirth2.otxlab.net"),
		)
	*/

	/*
		gs, _ := svc.ListGlobalSearches()
		fmt.Printf("List %+v\n", gs)

		svc.DeleteGlobalSearches([]string{"savedSearch.xyz", "savedSearch.abc"})
	*/
	// gs, _ = svc.ListGlobalSearches()
	// fmt.Printf("List %+v\n", gs)
	/*
		data := []adp.GlobalSearchDefinition{
			{
				DisplayName: "xyz",
				Queries:     []string{"1234", "hahah"},
			},
		}
	*/

	// gs, _ = svc.CreateGlobalSearches(data)
	// fmt.Printf("Create %+v\n", gs)

	/*
			data1 := []adp.GlobalSearchDefinition{
				{
					ID:      "savedSearch.xyz",
					Queries: []string{"hello", "world"},
				},
			}
			gs, _ = svc.UpdateGlobalSearches(data1)
			fmt.Printf("Create %+v\n", gs)

		roles := []adp.ApplicationRoles{
			{
				Enabled:               false,
				GroupOrUserName:       "demouser1",
				ApplicationIdentifier: "documentHold.demo00001",
				Roles:                 "Standard User",
			},
		}
		svc.AssignUsersOrGroupsToApplication(roles)

	*/

	documentHolds, err := svc.ListDocumentHolds()
	if err != nil {
		panic(err)
	}

	hasAccess, _ := svc.FindApplicationsUserHasAccess(documentHolds, "pyan")
	fmt.Printf("%+v\n", hasAccess)

}
