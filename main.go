package main

import (
	"fmt"

	"github.com/mbetim/bit/pkg/cmd"
	_ "github.com/mbetim/bit/pkg/cmd/auth"
	_ "github.com/mbetim/bit/pkg/cmd/config"
	_ "github.com/mbetim/bit/pkg/cmd/pr"
	_ "github.com/mbetim/bit/pkg/cmd/repo"
	"github.com/mbetim/bit/pkg/config"
)

func main() {
	err := config.StartConfig()
	if err != nil {
		fmt.Println("Unable to start config:", err)
		return
	}

	cmd.Execute()
}
