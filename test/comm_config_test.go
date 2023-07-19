package test

import (
	"testing"

	"github.com/hao508506/micro_cart/common"
)

func TestGetConsulConfig(t *testing.T) {

	_, err := common.GetConsulConfig("38.54.94.198", 8500, "/example-consul/config")
	if err != nil {
		panic(err)
	}

}
