package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	//"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Cargo struct {
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

	if function == "queryCargo" {
		return s.queryCargo(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "refreshCondition" {
		return s.refreshCondition(APIstub, args)
	} else if function == "resetConditionToNormal" {
		return s.resetConditionToNormal(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) refreshCondition(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cargoAsBytes, _ := APIstub.GetState(args[0])
	cargo := Cargo{}

	json.Unmarshal(cargoAsBytes, &cargo)
	cargo.State = "damaged"

	cargoAsBytes, _ = json.Marshal(cargo)
	APIstub.PutState(args[0], cargoAsBytes)

	return shim.Success(cargoAsBytes)
}

func (s *SmartContract) resetConditionToNormal(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cargoAsBytes, _ := APIstub.GetState(args[0])
	cargo := Cargo{}

	json.Unmarshal(cargoAsBytes, &cargo)
	cargo.State = "ok"

	cargoAsBytes, _ = json.Marshal(cargo)
	APIstub.PutState(args[0], cargoAsBytes)

	return shim.Success(cargoAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cargos := []Cargo{
		Cargo{State: "normal"},
	}

	cargoAsBytes, _ := json.Marshal(cargos[0])
	APIstub.PutState("CARGO1", cargoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryCargo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cargoAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(cargoAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
