NAME = imgproc

# BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
# COMMIT = $(shell git rev-parse --short HEAD)
# BUILDTIME = $(shell date +%Y-%m-%dT%T%z)
# LASTTAG = $(shell git describe --tags --abbrev=0 --dirty)
# GOPATH = $(shell echo "$$GOPATH")
LOCAL = $(shell pwd)
# LD_OPTS = -ldflags="-X main.branch=${BRANCH} -X main.commit=${COMMIT} -X main.lasttag=${LASTTAG} -X main.buildtime=${BUILDTIME} -linkmode=external -w -s"

# run:
# 	cd $(LOCAL)/cmd && make build && ./$(NAME) && cd ..

build:
	cd ./cmd && go build  -o $(NAME) . && cd ..

run:
	cd $(LOCAL)/cmd && cd .. && make build  && cd cmd && ./$(NAME) && ../

dist:
	cd ./cmd && GOOS=linux GOARCH=amd64 go build  -o $(NAME) . && cd ..

