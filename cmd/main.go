package main

import (
	"flag"
	"fmt"

	"opentext.com/axcelerate/adp/client"
	"opentext.com/axcelerate/adp/task"
)

func main() {

	flag.Parse()

	client := client.NewClient(
		client.WithHost("localhost"),
		client.WithUser("adpuser"),
		client.WithPassword("adpus3r"),
	)

	listEntitiesTask := task.NewListEntitiesTask(
		client,
		task.WithListEntitiesType("singleMindServer"),
		task.WithListEntitiesWhiteList(task.DefaultListEntitiesWhitelist),
	)

	s, _ := listEntitiesTask.StringOutput()
	fmt.Println(task.Beautify(s))
}
