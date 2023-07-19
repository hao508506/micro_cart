package main

import (
	"fmt"
	"log"
	"micro_cart/common"
	"micro_cart/domain/repository"
	"micro_cart/domain/service"
	"micro_cart/handler"
	proto "micro_cart/proto/cart"

	"github.com/go-micro/plugins/v4/registry/consul"
	ratelimiter "github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber"
	opentracing_micro "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 参考：https://www.cnblogs.com/bossma/p/16188603.html

const QPS = 100

func main() {

	consulConfig, err := common.GetConsulConfig("38.54.94.198", 8500, "/micro/config")
	if err != nil {
		panic(err)
	}

	consulRegistry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"38.54.94.198:8500",
		}
	})

	tc, io, err := common.NewTracer("go.micro.service.cart", "jaeger:6831")
	if err != nil {
		panic(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tc)

	// Create service
	// regCheckFunc := func(ctx context.Context) error {
	// 	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " do register check")
	// 	if 1+1 == 2 {
	// 		return nil
	// 	}
	// 	return errors.New("this not earth")
	// }

	rpcServer := server.NewServer(
		server.Name("go.micro.service.micro_cart"),
		server.Address("0.0.0.0:8501"),
		server.Registry(consulRegistry),
		// server.RegisterCheck(regCheckFunc),
		// server.RegisterInterval(10*time.Second), //指定程序去刷新TLL的频率。
		// server.RegisterTTL(20*time.Second),      //指定TTL的生存周期，如果超过这个时间没有刷新TTL，则Consul会认为服务是不健康
	)
	srv := micro.NewService(
		micro.Server(rpcServer),
		//绑定链路追踪
		micro.WrapHandler(opentracing_micro.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimiter.NewHandlerWrapper(QPS)),
	)

	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlInfo.User, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Name, "30s")
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "tbl_", // 添加表名前缀（可选）
			SingularTable: true, // 禁用表名的复数形式
		},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	cartRepository := repository.NewCartRepository(_db)
	cartRepository.InitTable()

	cartService := service.NewCartService(cartRepository)

	// Register handler
	proto.RegisterCartHandler(rpcServer, &handler.Cart{CartService: cartService})

	srv.Init() //需要注意不要使用service.Init()，因为这里边会覆盖 RegisterInterval 和 RegisterTTL 的设置，除非你不关心这两个参数

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
