package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	//"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type SupplyContract struct {
	Name	 string `json:"name"`
	Id     int64  `json:"Id"`
	State  string `json:"state"`
}

/*
 * The Init method is called when the Smart Contract "Cargo Condition" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "Cargo Condition"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately

	if function == "checkState" {
		return s.checkState(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) checkState(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// get condition of cargo
	var response = APIstub.InvokeChaincode("cargo_condition", nil, "mychannel")

	fmt.Printf("CargoCondition invoking response: %s", response.Message)

	var payload = string(response.Payload)

	if payload == "GOT_IT" {
		return shim.Error("GOT_IT!!!!!")
	}

	return shim.Success([]byte("payload_test"))
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	contracts := []SupplyContract{
		SupplyContract{Name: "Codfish contract #1", Id: 42, State: "not-yet-started"},
	}

	i := 0
	for i < len(contracts) {
		fmt.Println("i is ", i)
		contractAsBytes, _ := json.Marshal(contracts[i])
		APIstub.PutState("SupplyContract"+strconv.Itoa(i), contractAsBytes)
		fmt.Println("Added", contracts[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryContract(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	contractAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(contractAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
