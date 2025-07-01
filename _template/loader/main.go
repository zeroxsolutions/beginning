package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/zeroxsolutions/sazabi"
)

func main() {
	stmts, err := gormschema.New("mysql").Load()
	if err != nil {
		sazabi.Fatalf("failed to create gormschema: %v", err)
	}
	sazabi.Infof(stmts)
}
