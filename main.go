package main

import (
	"github.com/mbetim/bit/pkg/cmd"
	_ "github.com/mbetim/bit/pkg/cmd/auth"
	_ "github.com/mbetim/bit/pkg/cmd/pr"
)

func main() {
	cmd.Execute()
}
