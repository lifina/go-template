package cmd

import (
	"fmt"

	"github.com/lifina/go-template/pkg/api"
	"github.com/lifina/go-template/pkg/config"

	"github.com/spf13/cobra"
	"gopkg.in/go-playground/validator.v9"
)

const (
	cPort   = "port"
	cDBHost = "dbhost"
	cDBUser = "dbuser"
	cDBPass = "dbpass"
	cDBPort = "dbport"
	cDBName = "dbname"
)

type serveApp struct {
	*cobra.Command
}

func newServeApp() serveApp {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "start http server with configured api",
		Long:  `Starts a http server and serves the configured api`,
	}
	app := serveApp{Command: cmd}
	var port int64
	cnf := config.DBConfig{}
	cmd.Flags().Int64VarP(&port, cPort, "p", 3001, "http server port")
	cmd.Flags().StringVar(&cnf.Host, cDBHost, "127.0.0.1", "database host")
	cmd.Flags().StringVar(&cnf.User, cDBUser, "", "database user(required)")
	cmd.Flags().StringVar(&cnf.Password, cDBPass, "", "database password(required)")
	cmd.Flags().Int64Var(&cnf.Port, cDBPort, 3310, "database port")
	cmd.Flags().StringVar(&cnf.Database, cDBName, "", "database name(required)")
	cmd.MarkFlagRequired(cDBUser)
	cmd.MarkFlagRequired(cDBPass)

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		v := validator.New()
		return v.Struct(cnf)
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ds := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			cnf.User,
			cnf.Password,
			cnf.Host,
			cnf.Port,
			cnf.Database,
		)

		server, err := api.NewServer("localhost", port, ds)
		if err != nil {
			return err
		}

		server.Start()
		return nil
	}
	return app
}

func init() {
	RootCmd.AddCommand(newServeApp().Command)
}
