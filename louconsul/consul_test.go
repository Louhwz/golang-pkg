package louconsul

import (
	"testing"

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
