module github.com/yangliang4488/log_agent_system

go 1.16

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/Shopify/sarama v1.29.1
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)
