module shippy

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/globalsign/mgo v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/labstack/gommon v0.3.0
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-plugins v1.5.1
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb // indirect
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20191109021931-daa7c04131f5
	google.golang.org/grpc v1.25.1
)

replace github.com/globalsign/mgo => github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
