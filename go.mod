module micro_cart

go 1.18

require (
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	gorm.io/gorm v1.25.1
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-acme/lego/v4 v4.4.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.0.4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/hashicorp/consul/api v1.9.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-hclog v0.12.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/hashicorp/serf v0.9.5 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/rogpeppe/go-internal v1.6.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/sync v0.2.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)

require (
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/evanphx/json-patch/v5 v5.5.0 // indirect
	github.com/go-micro/plugins/v4/config/source/consul v1.2.0
	github.com/go-micro/plugins/v4/registry/consul v1.2.1
	github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber v1.2.0
	github.com/go-micro/plugins/v4/wrapper/trace/opentracing v1.2.0
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	go-micro.dev/v4 v4.10.2
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.9.1 // indirect
	google.golang.org/protobuf v1.31.0
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gorm.io/driver/mysql v1.5.1
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
