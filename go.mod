module github.com/suslovs/yandex-cloud

go 1.15

require (
	github.com/aws/aws-sdk-go v1.19.39
	github.com/c2h5oh/datasize v0.0.0-20200112174442-28bbd4740fee
	github.com/golang/protobuf v1.4.1
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/hashicorp/terraform-plugin-sdk v1.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/yandex-cloud/go-genproto v0.0.0-20211028095821-ae507b7629fe
	github.com/yandex-cloud/go-sdk v0.0.0-20211012083249-d09272b4a71c
	google.golang.org/genproto v0.0.0-20200323114720-3f67cca34472
	google.golang.org/grpc v1.28.0
)

replace git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
