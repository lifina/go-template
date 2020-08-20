package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lifina/go-template/internal/server"
	"github.com/spf13/cobra"
)

const (
	defaultPort   = "8080"
	defaultDBHost = "127.0.0.1"
	defaultDBPort = 3306
	defaultDBNet  = "tcp"
)

func NewCommand() *cobra.Command {

	var (
		config = server.Config{
			Port:   defaultPort,
			DBHost: defaultDBHost,
			DBPort: defaultDBPort,
			DBNet:  defaultDBNet,
		}
	)

	var command = &cobra.Command{
		Use:   "server",
		Short: "Run the bgo server",
		Long:  "Run the bgo server",
		RunE: func(c *cobra.Command, args []string) error {

			srv, err := server.NewServer(&config)
			if err != nil {
				return err
			}

			fmt.Fprint(os.Stdout, "start server")
			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					fmt.Fprintf(os.Stderr, "closed server, %s", err.Error())
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
			fmt.Printf("SIGNAL %d received, so server shutting down now...\n", <-quit)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				fmt.Fprintf(os.Stderr, "failed to gracefully shutdown, %s", err.Error())
			}

			fmt.Fprintf(os.Stdout, "server shutdown completed")
			return nil
		},
	}

	command.Flags().StringVar(&config.Port, "port", config.Port, "server port")
	command.Flags().StringVar(&config.DBNet, "dbNet", config.DBNet, "connection method to database")
	command.Flags().StringVar(&config.DBHost, "dbHost", config.DBHost, "database host")
	command.Flags().Int64Var(&config.DBPort, "dbPort", config.DBPort, "database port")
	command.Flags().StringVar(&config.DBUser, "dbUser", "", "database user")
	command.Flags().StringVar(&config.DBPassword, "dbPass", "", "database password")
	command.Flags().StringVar(&config.DBName, "dbName", "", "database name")
	return command
}
