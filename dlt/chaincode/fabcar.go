package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"math/rand"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

/**
* ==============================================================================
* ==============================================================================
* DIGITAL ASSETS DEFINITIONS
* ==============================================================================
* ==============================================================================
*/
type Share struct {
	ShareId   string `json:"share_hash_id"`
	ShareCorporateId  string `json:"share_corporate_id"`
	ShareOwnerId string `json:"share_owner_id"`
}
func (s *SmartContract) pupolateInitialShareData(APIstub shim.ChaincodeStubInterface) sc.Response {
	shares := []Share{
		Share{
			ShareId: "85029b86a41d8cfed975c6f69856fb9759bec4548f208078d3211b771f24ef30",
			ShareCorporateId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
			ShareOwnerId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
		},
		Share{
			ShareId: "123871e5c5b5e783b1a3255a7c3d52ff828431d717f330d1c75fb69ee14faf5a",
			ShareCorporateId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
			ShareOwnerId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
		},
		Share{
			ShareId: "f15872c286a26f7b8c15a2fc8a984fbab3de270072c604c7889debac6f6ac698",
			ShareCorporateId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
			ShareOwnerId: "3d3f1d244a7a9d7b22217a9eed2cd12e584ed286d6db659c101d5ad8230ae61b",
		},
		Share{
			ShareId: "2f975d9d7e488d8956d2b3942e1b48f82405c064a2de8910850a053a4bc6b2e0",
			ShareCorporateId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
			ShareOwnerId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
		},
		Share{
			ShareId: "3550442e6b99ffbd377256296c85209af505fa11bd8059225ab4068ee12d37f6",
			ShareCorporateId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
			ShareOwnerId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
		},
		Share{
			ShareId: "4da1bd9087b12d1213d2ed59dbeb3b779020ad49317af537ae168e5ec8b83e58",
			ShareCorporateId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
			ShareOwnerId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
		},
		Share{
			ShareId: "485d7863b0c05ac2bb3ea0143653788eb3a3a099c3dbc7dff52d6af9a86a2a4e",
			ShareCorporateId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
			ShareOwnerId: "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def",
		},
		Share{
			ShareId: "d475f6cc9cbe5ca4428e50364c4a2192d1e35cb2cda565fa6c8163d185c2adb6",
			ShareCorporateId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
			ShareOwnerId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
		},
		Share{
			ShareId: "6860a5e722e0adb051df477363a78c58c8b882ecf2776fb552c21533cb800247",
			ShareCorporateId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
			ShareOwnerId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
		},
		Share{
			ShareId: "b8b23c7335192fd96e328dd55a0d6036726d38b691cdf88768e89310557a8467",
			ShareCorporateId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
			ShareOwnerId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
		},
		Share{
			ShareId: "de119fa0f2da067aa4d9e4189486c74ddd6f4aa386199a54dc0360c30b3ac153",
			ShareCorporateId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
			ShareOwnerId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
		},
		Share{
			ShareId: "b8c828cfa82142bfc77fbf6975ffe0f23d3f42c5cf3d823889bcca7dfab1e011",
			ShareCorporateId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
			ShareOwnerId: "f3e5b88f1e2226b924c75c902ba6a617815bd09d03b581f69bacd517871a79f6",
		},
	}

	i := 0
	for i < len(shares) {
		fmt.Println("i is ", i)
		AssetsAsBytes, _ := json.Marshal(shares[i])
		APIstub.PutState("SHARE"+strconv.Itoa(i), AssetsAsBytes)
		fmt.Println("Added", shares[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) fetchSharesByCorporateId(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	target_corporate_id := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"share_corporate_id\":\"%s\"}}", target_corporate_id)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (s *SmartContract) fetchSharesByUserId(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	target_user_id := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"share_owner_id\":\"%s\"}}", target_user_id)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) fetchAllShares(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "SHARE0"
	endKey := "SHARE9999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) fetchPortfolioInfo(
	APIstub shim.ChaincodeStubInterface,
	args []string) sc.Response {

	target_user_id := "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7"
	queryString := fmt.Sprintf("{\"selector\":{\"share_owner_id\":\"%s\"}}", target_user_id)
	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil { return shim.Error("ERROR!!!!!!") }

	defer resultsIterator.Close()

	var buffer bytes.Buffer
	// bArrayMemberAlreadyWritten := false
	ItrCounter := 0

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil { return shim.Error("ERROR!!!!!!") }

		share := Share{}
		json.Unmarshal(queryResponse.Value, &share)

		if strings.TrimRight(string(share.ShareCorporateId), "\n") == "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def" { ItrCounter++ }
	}

	buffer.WriteString("[")
	buffer.WriteString("{\"asset\":")
	buffer.WriteString("\"")
	buffer.WriteString("ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def")
	buffer.WriteString("\"},{\"number_share\":")
	buffer.WriteString("\"")
	buffer.WriteString(strconv.Itoa(ItrCounter))
	buffer.WriteString("\"},{\"market_price\":")
	buffer.WriteString("\"")
	buffer.WriteString("120")
	buffer.WriteString("\"}")
	buffer.WriteString("]")

	fmt.Printf("\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) fetchCorporateInfo(
	APIstub shim.ChaincodeStubInterface,
	args []string) sc.Response {

	target_company_id := "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def"
	queryString := fmt.Sprintf("{\"selector\":{\"share_corporate_id\":\"%s\"}}", target_company_id)
	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil { return shim.Error("ERROR!!!!!!") }

	defer resultsIterator.Close()
	var buffer bytes.Buffer
	// bArrayMemberAlreadyWritten := false
	ItrCounter := 0
	totalItrCounter := 0

	for resultsIterator.HasNext() {
		totalItrCounter++
		queryResponse, err := resultsIterator.Next()
		if err != nil { return shim.Error("ERROR!!!!!!") }
		share := Share{}
		json.Unmarshal(queryResponse.Value, &share)
		fmt.Printf("TARGET SHARE = \n%s\n", share)
		fmt.Printf("TARGET SHARE OWNER ID = \n%s\n", share.ShareOwnerId)

		if strings.TrimRight(string(share.ShareOwnerId), "\n") == "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def" { ItrCounter++ }
	}

	buffer.WriteString("[")
	buffer.WriteString("{\"company_owned_shares\":")
	buffer.WriteString("\"")
	buffer.WriteString(strconv.Itoa(ItrCounter))
	buffer.WriteString("\"},{\"company_owned_shares\":")
	buffer.WriteString("\"")
	buffer.WriteString(strconv.Itoa(totalItrCounter))
	buffer.WriteString("\"}")
	buffer.WriteString("]")

	fmt.Printf("\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) buyShare(
	APIstub shim.ChaincodeStubInterface,
	args []string) sc.Response {

	// if len(args) != 3 {
	// 	// Corporate_ID
	// 	// Number of Share
	// 	// Owner id
	// 	return shim.Error("Incorrect number of arguments. Expecting 3")
	// }

	target_company_id := "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def"
	queryString := fmt.Sprintf("{\"selector\":{\"share_owner_id\":\"%s\"}}", target_company_id)

	fmt.Printf("- #################:\n%s\n", queryString)

	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil { return shim.Error("ERROR!!!!!!") }

	defer resultsIterator.Close()
	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	ItrCounter := 0

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		fmt.Printf("QUERY RESPONSE = \n%s\n", queryResponse)
		ItrCounter++
		fmt.Printf("ITR COUNTER = \n%s\n", ItrCounter)

		if ItrCounter > 2 { break }

		share := Share{}
		json.Unmarshal(queryResponse.Value, &share)
		fmt.Printf("TARGET SHARE = \n%s\n", share)
		share.ShareOwnerId = "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7"

		var shareAsBytes, _ = json.Marshal(share)
		APIstub.PutState(queryResponse.Key, shareAsBytes)

		if err != nil { return shim.Error("ERROR!!!!!!") }
		if bArrayMemberAlreadyWritten == true { buffer.WriteString(",") }

		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	fmt.Printf("\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

type Company struct {
	CorporateId   string `json:"corporate_hash_id"`
	CorporateType  string `json:"corporate_type"`
}
func (s *SmartContract) pupolateInitialCompanyData(APIstub shim.ChaincodeStubInterface) sc.Response {
	shipments := []Shipment{
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
	}

	i := 0
	for i < len(shipments) {
		fmt.Println("i is ", i)
		shipmentAsBytes, _ := json.Marshal(shipments[i])
		APIstub.PutState("SHIPMENT"+strconv.Itoa(i), shipmentAsBytes)
		fmt.Println("Added", shipments[i])
		i = i + 1
	}

	return shim.Success(nil)
}

type User struct {
	UserId   string `json:"user_hash_id"`
	UserVerimiId   string `json:"user_hash_id"`
	UserType  string `json:"corporate_type"`
}
func (s *SmartContract) pupolateInitialUserData(APIstub shim.ChaincodeStubInterface) sc.Response {
	users := []User{
		User{
			UserId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
			UserVerimiId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
			UserType: "HUMAN",
		},
	}

	i := 0
	for i < len(users) {
		fmt.Println("i is ", i)
		AssetAsBytes, _ := json.Marshal(users[i])
		APIstub.PutState("USER"+strconv.Itoa(i), AssetAsBytes)
		fmt.Println("Added", users[i])
		i = i + 1
	}

	return shim.Success(nil)
}
func (s *SmartContract) fetchAllUser(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "USER0"
	endKey := "USER9999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

type BidAsk struct {
	BASahreId   string `json:"ba_share_hash_id"`
	BAType   string `json:"ba_type"`
	BAUserId   string `json:"ba_user_id"`
	BATimestampUTC   string `json:"ba_timestamp_utc"`
	BAPrice   string `json:"ba_price"`
	BANumber   string `json:"ba_number"`
}
func (s *SmartContract) pupolateInitialBidAskData(APIstub shim.ChaincodeStubInterface) sc.Response {
	bidasks := []BidAsk{
		BidAsk{
			BASahreId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
			BAType: "ASK",
		  BAUserId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
		  BATimestampUTC: "018-08-28T05:09:35+00:00",
		  BAPrice: "200.00",
		  BANumber: "85",
		},
	}

	i := 0
	for i < len(bidasks) {
		fmt.Println("i is ", i)
		AssetsAsBytes, _ := json.Marshal(bidasks[i])
		APIstub.PutState("BIDASK"+strconv.Itoa(i), AssetsAsBytes)
		fmt.Println("Added", bidasks[i])
		i = i + 1
	}

	return shim.Success(nil)
}
func (s *SmartContract) fetchAllBidAsk(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "BIDASK0"
	endKey := "BIDASK9999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

type Transaction struct {
	TransactionId   string `json:"txn_hash_id"`
	TransactionFromUserId   string `json:"txn_from_user_id"`
	TransactionToUserId   string `json:"txn_to_user_id"`
	TransactionTimestampUTC   string `json:"txn_timestamp_utc"`
	TransactionNumber   string `json:"txn_number"`
	TransactionPrice   string `json:"txn_price"`
}
func (s *SmartContract) pupolateInitialTransactionData(APIstub shim.ChaincodeStubInterface) sc.Response {
	transactions := []Transaction{
		Transaction{
			TransactionId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
			TransactionFromUserId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
			TransactionToUserId: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7",
			TransactionTimestampUTC: "018-08-28T05:09:35+00:00",
			TransactionNumber: "200.00",
			TransactionPrice: "85",
		},
	}

	i := 0
	for i < len(transactions) {
		fmt.Println("i is ", i)
		AssetAsBytes, _ := json.Marshal(transactions[i])
		APIstub.PutState("TRANSACTION"+strconv.Itoa(i), AssetAsBytes)
		fmt.Println("Added", transactions[i])
		i = i + 1
	}

	return shim.Success(nil)
}
func (s *SmartContract) fetchAllTransaction(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "TRANSACTION0"
	endKey := "TRANSACTION9999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}



type Car struct {
  // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

type Shipment struct {
  // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Identity   string `json:"identity"`
	Name  string `json:"name"`
	Latitude string `json:"latitude"`
	Lognitude  string `json:"lognitude"`
	Value  string `json:"value"`
	Status  string `json:"status"`
	ShipperId  string `json:"shipperId"`
	CarrierId  string `json:"carrierId"`
}

type Location struct {
	ObjectType string `json:"docType"`
	LocationShipmentKey   string `json:"locationShipmentKey"`
	LocationShipmentIdentity   string `json:"locationShipmentIdentity"`
	Latitude string `json:"latitude"`
	Lognitude  string `json:"lognitude"`
	LocationDatetimeUtc  string `json:"locationDatetimeUtc"`
}

/**
* ==============================================================================
* ==============================================================================
* SMART CONTRACT LIFECYCLE
* ==============================================================================
* ==============================================================================
*/

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryCar" {
		return s.queryCar(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "initShipmentLedger" {
		return s.initShipmentLedger(APIstub)

		// BIDASK
	} else if function == "pupolateInitialBidAskData" {
		return s.pupolateInitialBidAskData(APIstub)
	} else if function == "fetchAllShares" {
		return s.fetchAllShares(APIstub)
		// SHARE
	} else if function == "pupolateInitialShareData" {
		return s.pupolateInitialShareData(APIstub)
	} else if function == "fetchAllBidAsk" {
		return s.fetchAllBidAsk(APIstub)
	} else if function == "fetchSharesByUserId" {
		return s.fetchSharesByUserId(APIstub, args)
	} else if function == "fetchSharesByCorporateId" {
		return s.fetchSharesByCorporateId(APIstub, args)
	} else if function == "buyShare" {
		return s.buyShare(APIstub, args)
	} else if function == "fetchCorporateInfo" {
		return s.fetchCorporateInfo(APIstub, args)
	} else if function == "fetchPortfolioInfo" {
		return s.fetchPortfolioInfo(APIstub, args)
		// TRANSACTION
	} else if function == "pupolateInitialTransactionData" {
		return s.pupolateInitialTransactionData(APIstub)
	} else if function == "fetchAllTransaction" {
		return s.fetchAllTransaction(APIstub)
		// USER 
	} else if function == "pupolateInitialUserData" {
		return s.pupolateInitialUserData(APIstub)
	} else if function == "fetchAllUser" {
		return s.fetchAllUser(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/**
* ==============================================================================
* ==============================================================================
* POPULATION DATA
* ==============================================================================
* ==============================================================================
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	}

	i := 0
	for i < len(cars) {
		fmt.Println("i is ", i)
		carAsBytes, _ := json.Marshal(cars[i])
		APIstub.PutState("CAR"+strconv.Itoa(i), carAsBytes)
		fmt.Println("Added", cars[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) initShipmentLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	shipments := []Shipment{
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
		Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""},
	}

	i := 0
	for i < len(shipments) {
		fmt.Println("i is ", i)
		shipmentAsBytes, _ := json.Marshal(shipments[i])
		APIstub.PutState("SHIPMENT"+strconv.Itoa(i), shipmentAsBytes)
		fmt.Println("Added", shipments[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}

	carAsBytes, _ := json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}
/**
* ==============================================================================
* ==============================================================================
* UTILITIES
* ==============================================================================
* ==============================================================================
*/
func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}

func random(min, max int) int {
  rand.Seed(time.Now().Unix())
  return rand.Intn(max - min) + min
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
	// =========================================================================================
	// getQueryResultForQueryString executes the passed in query string.
	// Result set is built and returned as a byte array containing the JSON results.
	// =========================================================================================
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

/**
* ==============================================================================
* ==============================================================================
* QUERIES (NO CONSENSUS VALIDATION)
* ==============================================================================
* ==============================================================================
*/
func (s *SmartContract) createShipment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var shipment = Shipment{Identity: "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7", Name: "very expensive material 100098", Latitude: "50.936508", Lognitude: "6.939782", Value: "200.00", Status: "waiting for carrier", ShipperId: "awesome shipper", CarrierId: ""}

	shipmentAsBytes, _ := json.Marshal(shipment)

	var keyId = random(10, 999)
	APIstub.PutState("SHIPMENT"+strconv.Itoa(keyId), shipmentAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllShipments(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "SHIPMENT0"
	endKey := "SHIPMENT999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CAR0"
	endKey := "CAR999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/**
* ==============================================================================
* ==============================================================================
* INVOKE (CONSENSUS METHOD)
* ==============================================================================
* ==============================================================================
*/
func (s *SmartContract) assignCarrier(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	shipmentAsBytes, _ := APIstub.GetState(args[0])
	shipment := Shipment{}
	json.Unmarshal(shipmentAsBytes, &shipment)
	shipment.CarrierId = "carrier_id_921"

	shipmentAsBytes, _ = json.Marshal(shipment)
	APIstub.PutState(args[0], shipmentAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateShipmentStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	shipmentAsBytes, _ := APIstub.GetState(args[0])
	shipment := Shipment{}
	json.Unmarshal(shipmentAsBytes, &shipment)
	shipment.Status = args[1]

	shipmentAsBytes, _ = json.Marshal(shipment)
	APIstub.PutState(args[0], shipmentAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.Owner = args[1]

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}


func (s *SmartContract) queryShipmentsByCarrierId(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	//   0
	// "carrier_id_921"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carrier_id := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"carrierId\":\"%s\"}}", carrier_id)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) queryShipmentsByShipperId(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	shipper_id := strings.ToLower(args[0])
	queryString := fmt.Sprintf("{\"selector\":{\"shipperId\":\"%s\"}}", shipper_id)
	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (s *SmartContract) logLocationForShipment(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var location = Location{
		LocationShipmentKey: args[0],
		LocationShipmentIdentity: "",
		Latitude: "50.000000",
		Lognitude: "6.000000",
		LocationDatetimeUtc: time.Now().UTC().String()}

	locaitionAsBytes, _ := json.Marshal(location)

	var keyid = random(10, 999)
	stub.PutState("LOCATION"+strconv.Itoa(keyid), locaitionAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllLocationsForShipment(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// shipper_id := strings.ToLower(args[0])
	queryString := fmt.Sprintf("{\"selector\":{\"locationShipmentKey\":\"%s\"}}", args[0])
	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

/**
* ==============================================================================
* ==============================================================================
* MAIN FUNCTION
* ==============================================================================
* ==============================================================================
*/
// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
