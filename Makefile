#
ARCH=$(shell uname -s)
SERVERS:=$(shell ls app)

PB_FILES = $(shell find . -name '*.proto')
PB_DIRS = $(sort $(dir $(PB_FILES)))
PB_GO_FILES = $(shell find . -name '*.pb.go')
PB_DIR_TGTS = $(addprefix _PB, $(PB_DIRS))
WORK_DIR = $(shell pwd)
PB_VER = 26.1

.PHONY: all
all: $(SERVERS)

.PHONY: $(SERVERS)
$(SERVERS):
	@echo -n "Now build $@ ... "
	@mkdir -p bin && touch bin/.gitignore
	@go mod tidy && go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" ./app/$@/
	@mv $@ bin && echo Done
# Reference: [解决《panic: proto: file “xxx.proto“ is already registered》问题](https://blog.csdn.net/zhaolinfenggg/article/details/135776526)

.PHONY: lint
lint:
	golint ./...
	golangci-lint run ./...

.PHONY: fmt
fmt:
	find . -name "*.go" | xargs goimports -e -d -local git.woa.com -w && \
    find . -name "*.go" | xargs gofmt -e -d -s -w

.PHONY: test
test:
	go test -v ./... -gcflags "all=-N -l"

.PHONY: cover
cover:
	go test -v ./... -gcflags "all=-N -l" -coverprofile=tmp_coverage.out
	go tool cover -html=tmp_coverage.out
	if [ -f "tmp_coverage.out" ]; then rm tmp_coverage.out; fi

# ======== protobuf 文件编译支持 ========

# pb 编译规则
.PHONY: pb
pb: $(PB_DIR_TGTS)

# 寻找包含 .proto 的目录并编译
.PHONY: $(PB_DIR_TGTS)
$(PB_DIR_TGTS):
	@for dir in $(subst _PB,, $@); do \
		echo Now Build proto in directory: $$dir; \
		cd $$dir; rm -rf mock; \
		export PATH=$(PATH); \
		rm -f *.pb.go; rm -f *.trpc.go; \
		find . -name '*.proto' | xargs -I DD \
			trpc create -f --protofile=DD --protocol=trpc --rpconly --nogomod --alias --mock=false --protodir=$(WORK_DIR)/proto; \
		ls *.trpc.go | xargs -I DD mockgen -source=DD -destination=mock/DD -package=mock ; \
		find `pwd` -name '*.pb.go'; \
		find `pwd` -name '*.pb.go' | xargs -I XXXX sed -i 's/err_code,omitempty/err_code/g' XXXX; \
	done

_PROTOC_PKG_URL=https://github.com/protocolbuffers/protobuf/releases/download/v$(PB_VER)/protoc-$(PB_VER)-linux-x86_64.zip

# installpb 用于在设备上安装 protobuf 编译器, 仅适用于 Linux 环境。
# 如果环境 OK 那么不用执行
.PHONY: installpb
installpb:
	wget $(_PROTOC_PKG_URL)
	7z x $(notdir $(_PROTOC_PKG_URL)) -o/usr/local -y
	rm -f $(notdir $(_PROTOC_PKG_URL))*
	chmod +x /usr/local/bin/protoc
	@echo ---- $@ done ----
	@protoc --version | xargs echo "Protobuf version:"

.PHONY: installtrpc
installtrpc:
	go install trpc.group/trpc-go/trpc-cmdline/trpc@latest
	@echo ---- $@ done ----
	@trpc version | xargs echo "TRPC version:"

# ======== 自动 go generate 支持 ========

GO_GENERATE_DIRS=
ifeq ($(ARCH), Darwin)
	GO_GENERATE_DIRS=$(sort $(dir $(shell grep -lr --include='*.go' '//go:generate ' .)))
else ifeq ($(ARCH), Linux)
	GO_GENERATE_DIRS=$(sort $(dir $(shell grep -lr --include='*.go' '//go:generate ')))
else
	$(error 不支持的系统: $(ARCH))
endif

.PHONY: gogenerate
gogenerate: mockinstall $(GO_GENERATE_DIRS)
	@go mod tidy

.PHONY: installmock
installmock:
	go install go.uber.org/mock/mockgen@latest
	@echo ---- $@ done ----
	@mockgen -version | xargs echo "mockgen version:"

.PHONY: $(GO_GENERATE_DIRS)
$(GO_GENERATE_DIRS):
	@for dir in $@; do \
		echo ==== go generate $$dir ====; \
		cd $$dir; \
		go generate; \
	done