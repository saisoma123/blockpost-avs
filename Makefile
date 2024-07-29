ROOT_DIR := .
FRONTEND_DIR := blockpost-frontend
OPERATOR_DIR := operator/cmd
SCRIPT_DIR := contracts/script

all:
	$(MAKE) operator &
	$(MAKE) frontend

frontend:
	konsole --hold -e "bash -c 'cd $(FRONTEND_DIR) && npm start; exec bash'" 

operator:
		@dotenv -f .env -- sh -c 'cd $(OPERATOR_DIR) && go run main.go --config operator.anvil.yaml'	

npm_build:
	cd $(FRONTEND_DIR) && npm install

deploy:
	cd $(SCRIPT_DIR) && forge script BlockPostDeployer.s.sol --rpc-url https://ethereum-holesky-rpc.publicnode.com --private-key $$PRIVATE_KEY --broadcast
	
.PHONY: all operator frontend deploy npm_build

SOL_TEST := contracts/test
OPERATOR_ROOT := operator

contract_tests: 
	cd $(SOL_TEST) && forge test

operator_tests:
	cd $(OPERATOR_ROOT) && go test

runAllTests: contract_tests operator_tests