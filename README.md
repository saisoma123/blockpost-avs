# BlockPost-AVS
**BlockPost-AVS** is an actively validated service that allows for the storage and retrieval of messages.

## Get started

1. [Download and install Go](https://golang.org/dl/)
2. Add Go binary to your PATH and set up Go workspace:

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```
3. [Download and install Foundry](https://book.getfoundry.sh/getting-started/installation)
4. Run `go mod tidy` to install the Go packages this app uses in the `operator` directory
5. Install `konsole` with `sudo apt install konsole`
6. Install `dotenv` with `npm install -g dotenv-cli`
7. `make npm_build` command will install all of the necessary JS packages to run the
React app (make sure to have npm installed).
8. `make` command will start the app, it runs `go run main.go --config operator.anvil.yaml` to start the operator,
then it starts a React app that you can use to interact with the ServiceManager contract.
9. `make contract_tests` command will run all of the tests for the ServiceManager contract.
10. `make operator_tests` command will run all of the tests for the operator.

This AVS design works as follows: the user (AVS consumer) sends a submitMessage
request to the ServiceManager contract, and the request emits a MessageSubmitted
event with the message and message id on chain. The operator is 
listening for this particular event and extracts the message and validates the 
message, ensuring data integrity. A signature is also created with the message 
and operator private key. The validated message and signature is sent back to the
Service Manager, which recovers the address of the sender from the signature, and
makes sure that the sender is the operator. Once this is verified, the message
is stored in the messages mapping along with its id. For message retrieval,
only the user that stored the message in the first place can access it. 

One limitation of this design is that with event emission, the message is 
technically viewable publically on chain, so to ensure complete security, I would have
added Chainlink VRF integration to generate a random key and encrypt the message,
and emit the encrypted message on chain. Once the encrypted message is validated 
and sent back to the service manager, I would decrypt it and store it to the messages
mapping. One drawback of this though is that it would require 4 additonal mappings
to correctly pass the id and message data around to the requestRandomness function,
which would require more gas, and the Chainlink usage itself would require quite
a bit more gas. But to ensure complete security, as in only messages that belong
to a particular user can only be seen by that user, this would be one of the most
secure choices. Also, I would have added the onlyOperator modifier to the 
storeValidatedMessage function. These design choices are definitely needed if this
was to be deployed to mainnet.