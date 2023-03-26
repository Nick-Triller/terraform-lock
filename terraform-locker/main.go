package main

import (
	"fmt"

	"github.com/hashicorp/terraform/internal/backend"
	"github.com/hashicorp/terraform/internal/backend/remote-state/s3"
	"github.com/hashicorp/terraform/internal/states/statemgr"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	stateBackend := s3.New()
	// TODO pass config via ENV vars or something
	stateBackend.Configure(cty.ObjectVal(map[string]cty.Value{
		"bucket":         cty.StringVal("ntriller-terraform"),
		"key":            cty.StringVal("app-a"),
		"region":         cty.StringVal("eu-central-1"),
		"dynamodb_table": cty.StringVal("terraform-locks"),
	}),
	)
	stateMgr, _ := stateBackend.StateMgr(backend.DefaultStateName)
	lockInfo := statemgr.NewLockInfo()
	if _, err := stateMgr.Lock(lockInfo); err != nil {
		fmt.Printf("state is already locked: %s\n", err.Error())
	} else {
		fmt.Printf("locked state successfully: \n%s", lockInfo)
	}
}
