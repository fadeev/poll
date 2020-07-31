PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NewApp \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=polld \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=pollcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/polld
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/pollcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

init-pre:
	rm -rf ~/.pollcli
	rm -rf ~/.polld
	polld init mynode --chain-id poll
	pollcli config keyring-backend test

init-user1:
	pollcli keys add user1 --output json 2>&1

init-user2:
	pollcli keys add user2 --output json 2>&1

init-post:
	polld add-genesis-account $$(pollcli keys show user1 -a) 1000token,100000000stake
	polld add-genesis-account $$(pollcli keys show user2 -a) 500token
	pollcli config chain-id poll
	pollcli config output json
	pollcli config indent true
	pollcli config trust-node true
	polld gentx --name user1 --keyring-backend test
	polld collect-gentxs

init: init-pre init-user1 init-user2 init-post