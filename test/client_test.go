package test

import (
	"context"
	"log"
	"testing"

	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"

	"github.com/hao508506/micro_cart/common"
	proto "github.com/hao508506/micro_cart/proto/cart"

	opentracing_micro "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
)

func TestClient(t *testing.T) {
	registry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"38.54.94.198:8500",
		}
	})

	tc, io, err := common.NewTracer("go.micro.service.cart.client", "38.54.94.198:6831")
	if err != nil {
		panic(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tc)

	srv := micro.NewService(
		micro.Name("go.micro.client.cart.clent"),
		// micro.Address("0.0.0.0:8501"),
		micro.Registry(registry),
		micro.WrapHandler(opentracing_micro.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	service := proto.NewCartService("go.micro.service.micro_cart", srv.Client())

	rsp, err := service.AddCart(context.TODO(), &proto.CartInfo{}, client.WithAddress("38.54.94.198:8501"))
	if err != nil {
		log.Fatalln(err)
	}

	t.Logf("%v", rsp)
}
