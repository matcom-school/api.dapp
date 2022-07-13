run:
	swag init --parseDependency --parseInternal --parseDepth 1 --md docs/md_endpoints
	go mod tidy
	go mod vendor
	go build -o ./api.dapp
	./api.dapp

echo-export:
	echo export HLF_DAPP_CONFIG=${PWD}/conf.linux_and_wsl.yaml