package test

import (
	"micro_cart/common"
	"testing"
)

func TestGetConsulConfig(t *testing.T) {

	_, err := common.GetConsulConfig("38.54.94.198", 8500, "/example-consul/config")
	if err != nil {
		panic(err)
	}

}
