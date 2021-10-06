package louconsul

import (
	"testing"

	"k8s.io/utils/pointer"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	client, err := NewClient(nil)
	assert.Nil(t, err, "Can't new consul client")
	err = Register(client, "cpu", "192.192.192.192", 8080, "/health")
	assert.Nil(t, err, "Can't register service")
	err = Register(client, "cpu", "192.192.192.193", 8080, "/health")
	assert.Nil(t, err, "Can't register service")
}

func TestListCatalogService(t *testing.T) {
	client, err := NewClient(nil)
	assert.Nil(t, err, "Can't new consul client")
	data, _, err := client.Catalog().Services(nil)
	assert.Nil(t, err, "Can't list catalog service")
	for serviceName := range data {
		endpoints, _, err := client.Catalog().Service(serviceName, "", nil)
		assert.Nil(t, err, "")
		t.Log(endpoints)
	}

}

func TestRegister2(t *testing.T) {
	var consulAddr = pointer.StringPtr("192.168.1.236:8500")
	client, err := NewClient(consulAddr)
	if err != nil {
		panic(err)
	}
	err = Register(client, "louhwz-test", "192.192.192.192", 8500, "/health")
	assert.Nil(t, err, "Register Err")
}
