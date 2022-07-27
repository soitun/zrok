package main

import (
	"github.com/openziti-test-kitchen/zrok/http"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok/tunnel"
	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
	"github.com/openziti-test-kitchen/zrok/zrokdir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var httpCmd = &cobra.Command{
	Use:   "http <endpoint>",
	Short: "Start an http terminator",
	Args:  cobra.ExactArgs(1),
	Run:   handleHttp,
}

func handleHttp(_ *cobra.Command, args []string) {
	idCfg, err := zrokdir.IdentityConfigFile()
	if err != nil {
		panic(err)
	}
	cfg := &http.Config{
		IdentityPath:    idCfg,
		EndpointAddress: args[0],
	}
	id, err := zrokdir.ReadIdentityId()
	if err != nil {
		panic(err)
	}

	zrok := newZrokClient()
	req := tunnel.NewTunnelParams()
	req.Body = &rest_model_zrok.TunnelRequest{
		Endpoint: cfg.EndpointAddress,
		Identity: id,
	}
	resp, err := zrok.Tunnel.Tunnel(req)
	if err != nil {
		panic(err)
	}
	cfg.Service = resp.Payload.Service

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanupHttp(cfg, zrok)
		os.Exit(1)
	}()

	if err := http.Run(cfg); err != nil {
		panic(err)
	}
}

func cleanupHttp(cfg *http.Config, zrok *rest_client_zrok.Zrok) {
	logrus.Infof("shutting down '%v'", cfg.Service)
}