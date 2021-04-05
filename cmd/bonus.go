package cmd

import (
	"os"

	"github.com/MihaiBlebea/coffee-shop/bonus/conn"
	"github.com/MihaiBlebea/coffee-shop/bonus/event"
	"github.com/MihaiBlebea/coffee-shop/bonus/server"
	"github.com/MihaiBlebea/coffee-shop/bonus/server/handler"
	"github.com/MihaiBlebea/coffee-shop/bonus/stamp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startBonusCmd)
}

var startBonusCmd = &cobra.Command{
	Use:   "start-bonus",
	Short: "Start the bonus server api.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		c, err := conn.ConnectSQL()
		if err != nil {
			return err
		}

		stampService := stamp.New(c)

		bus, err := event.New(stampService)
		if err != nil {
			return err
		}

		go bus.Listen()

		s := server.New(handler.New(stampService, l), l)

		s.Run()

		return nil
	},
}
