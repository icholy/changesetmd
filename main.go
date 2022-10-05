package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

func main() {
	// read the changeset json
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("no changeset file provided")
	}
	data, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	changes, err := ParseChangeSet(data)
	if err != nil {
		log.Fatal(err)
	}
	// convert it to markdown
	var table MarkdownTable
	table.EmptyValue = "-"
	table.WriteHeaders(
		"Action",
		"Logical ID",
		"Physical ID",
		"Resource Type",
		"Replacement",
	)
	for _, c := range changes {
		if c.Type != types.ChangeTypeResource {
			msg := fmt.Sprintf("ERROR: unknown type: %q", c.Type)
			table.WriteRow(msg)
			continue
		}
		table.WriteRow(
			string(c.ResourceChange.Action),
			aws.ToString(c.ResourceChange.LogicalResourceId),
			aws.ToString(c.ResourceChange.PhysicalResourceId),
			aws.ToString(c.ResourceChange.ResourceType),
			string(c.ResourceChange.Replacement),
		)
	}
	// output it
	fmt.Println(table.String())
}
