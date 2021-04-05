package cmd

import (
	"os"

	"github.com/MihaiBlebea/coffee-shop/order/conn"
	"github.com/MihaiBlebea/coffee-shop/order/payment"
	"github.com/MihaiBlebea/coffee-shop/order/server"
	"github.com/MihaiBlebea/coffee-shop/order/server/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startOrderCmd)
}

var startOrderCmd = &cobra.Command{
	Use:   "start-order",
	Short: "Start the order server api.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		c, err := conn.ConnectSQL()
		if err != nil {
			return err
		}

		paymentService := payment.New(c)

		s := server.New(handler.New(paymentService, l), l)

		s.Run()

		return nil
	},
}
