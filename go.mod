module github.com/GeeVong/SimpleGo

go 1.20

require (
	github.com/GeeVong/SimpleGo/log v0.0.0-20230828083616-27f1fb050dc9
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.0
	github.com/liangdas/mqant v2.0.0+incompatible
	github.com/streadway/amqp v1.1.0
)

require (
	github.com/anqiansong/ketty v0.0.0-20211202021934-dbaf2e277891 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.30.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

replace github.com/rs/zerolog v1.30.0 => ./log/zerolog
replace github.com/GeeVong/SimpleGo/log v0.0.0-20230828083616-27f1fb050dc9 => ./log