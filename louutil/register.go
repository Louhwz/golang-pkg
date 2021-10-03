package louutil

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/louhwz/pkg/louconsul"
	"github.com/louhwz/pkg/lounetwork"
)

func Register(consulAddr *string, service, healthCheckURI string, port int) error {
	var (
		err          error
		ip           string
		consulClient *api.Client
	)

	consulClient, err = louconsul.NewClient(consulAddr)
	if err != nil {
		return err
	}

	ip = lounetwork.GetLocalIP()
	if ip == "" {
		return fmt.Errorf("can't get pod's local ip")
	}

	err = louconsul.Register(consulClient, service, ip, port, healthCheckURI)
	return err
}
