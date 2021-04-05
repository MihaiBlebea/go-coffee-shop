package cmd

import (
	"os"

	"github.com/MihaiBlebea/coffee-shop/account/conn"
	"github.com/MihaiBlebea/coffee-shop/account/event"
	"github.com/MihaiBlebea/coffee-shop/account/server"
	"github.com/MihaiBlebea/coffee-shop/account/server/handler"
	"github.com/MihaiBlebea/coffee-shop/account/user"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startAccountCmd)
}

var startAccountCmd = &cobra.Command{
	Use:   "start-account",
	Short: "Start the account server api.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		c, err := conn.ConnectSQL()
		if err != nil {
			return err
		}

		evStore, err := event.New()
		if err != nil {
			return err
		}

		userService := user.New(c, evStore)

		s := server.New(handler.New(userService, l), l)

		s.Run()

		return nil
	},
}
