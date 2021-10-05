package louconsul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

const (
	// ServiceIDFormat service-ip-port
	ServiceIDFormat = "%v-%v-%v"

	healthCheckURIFormat = "http://%v:%v%v"
)

// NewClient addr format: 127.0.0.1:8500
func NewClient(addr *string) (*api.Client, error) {
	conf := api.DefaultConfig()
	if addr != nil {
		conf.Address = *addr
	}
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, fmt.Errorf("can't new consul client: %v", err)
	}
	return client, err
}

func Register(client *api.Client, service string, ip string, port int, healthCheckURI string) error {
	return client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      fmt.Sprintf(ServiceIDFormat, service, ip, port),
		Name:    service,
		Address: ip,
		Port:    port,
		Weights: &api.AgentWeights{
			Passing: 100,
			Warning: 1,
		},
		Check: &api.AgentServiceCheck{
			CheckID:                        fmt.Sprintf("Check: %s", service),
			Name:                           fmt.Sprintf("Service %s's check", service),
			Interval:                       "10s",
			Timeout:                        "3s",
			HTTP:                           fmt.Sprintf(healthCheckURIFormat, ip, port, healthCheckURI),
			Method:                         "Get",
			TLSSkipVerify:                  true,
			DeregisterCriticalServiceAfter: "30s",
		},
	})
}
