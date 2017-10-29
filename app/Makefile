PACKAGES = $(shell glide novendor)
default: test

all: deps fmt test build build_image push_image

deps:
	glide install

vet:
	go vet $(PACKAGES)

test: vet
	go test -v -cover $(PACKAGES)

dev:
	go run main.go -database_user root -database_password koe7POut -database_IP 104.236.47.92:3306 -database_name childre_qu -debug true -super_admin_token 1 -email_auth_name jay@kidsdynamic.com -email_auth_password VeG4cLt3GM2fUK5q -base_url http://localhost:8111
local:
	go run main.go -database_user root -database_IP localhost:3306 -database_name childre_qu -debug true -super_admin_token 1 -email_auth_name jay@kidsdynamic.com -email_auth_password VeG4cLt3GM2fUK5q -base_url http://localhost:8111

