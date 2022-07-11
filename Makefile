include .env
run:
	swag init --parseDependency --parseInternal --parseDepth 1 --md docs/md_endpoints
	go mod tidy
	go build -o ${DAPPPATH}
	api.dapp