package louconsul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/louhwz/pkg/loustring"
	"k8s.io/klog/v2"
)

const (
	// ServiceIDFormat service-ip-port
	ServiceIDFormat = "%v-%v-%v"

	healthCheckURIFormat = "http://%v:%v%v"

	ProtocolTagName = "protocol"
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

func Register(client *api.Client, service string, ip string, port int, protocol string) error {
	id := fmt.Sprintf(ServiceIDFormat, service, ip, port)
	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    service,
		Address: ip,
		Port:    port,
		Weights: &api.AgentWeights{
			Passing: 100,
			Warning: 1,
		},
		Meta: map[string]string{
			ProtocolTagName: protocol,
		},
		Check: &api.AgentServiceCheck{
			CheckID:  fmt.Sprintf("Check for %v", id),
			Name:     fmt.Sprintf("Service %s's check", id),
			Interval: "10s",
			Timeout:  "3s",
			//HTTP:                           fmt.Sprintf(healthCheckURIFormat, ip, port, healthCheckURI),
			TCP:                            fmt.Sprintf("%v:%v", ip, port),
			Method:                         "Get",
			DeregisterCriticalServiceAfter: "15s",
		},
	}
	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		klog.Errorf("Register Error. Err=%v. Registration=%v", err, loustring.ToJsonString(registration))
	} else {
		klog.Infof("Register Success. Registration=%v", loustring.ToJsonString(registration))
	}
	return err
}
