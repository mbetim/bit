package main

import (
	"github.com/mbetim/bit/pkg/cmd"
	_ "github.com/mbetim/bit/pkg/cmd/auth"
)

func main() {
	cmd.Execute()
}
