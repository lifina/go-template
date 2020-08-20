package bgo

import (
	"github.com/lifina/go-template/internal/cmd/server"
	"github.com/spf13/cobra"
)

func NewCommand(name string) *cobra.Command {

	c := &cobra.Command{
		Use:   name,
		Short: "bgo",
		Long:  `bgo`,
	}
	c.AddCommand(
		server.NewCommand(),
	)
	return c
}
