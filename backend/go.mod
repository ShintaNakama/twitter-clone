module github.com/ShintaNakama/twitter-clone/backend

go 1.15

require (
	github.com/DataDog/datadog-go v4.0.0+incompatible // indirect
	github.com/go-gorp/gorp v2.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/lib/pq v1.8.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/poy/onpar v1.1.2 // indirect
	github.com/rakyll/statik v0.1.7
	github.com/schemalex/schemalex v0.1.1
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200828194041-157a740278f4 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	golang.org/x/tools v0.0.0-20200828161849-5deb26317202 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20200831141814-d751682dd103 // indirect
	google.golang.org/grpc v1.37.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.27.0
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
	pb v1.0.0
)

replace pb v1.0.0 => ../pb
