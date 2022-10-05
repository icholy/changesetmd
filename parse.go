package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

// ParseChangeSet attempts to parse the json in the two commonly
// available json formats.
func ParseChangeSet(data []byte) ([]*types.Change, error) {
	var aux struct {
		Changes []*types.Change
	}
	// try to parse the output of describe-change-set
	if err := json.Unmarshal(data, &aux); err != nil {
		// try to parse the output from the Changes JSON tab
		if err := json.Unmarshal(data, &aux.Changes); err != nil {
			return nil, fmt.Errorf("invalid changeset: %v", err)
		}
	}
	return aux.Changes, nil
}
