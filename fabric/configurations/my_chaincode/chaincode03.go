package main

import (
	"fmt"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type InvokeCc struct {
}

func (m *InvokeCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var num string // Sum of asset holdings across accounts. Initially 0
	var numVal int // Sum of holdings
	var err error
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Initialize the chaincode
	num = args[0]
	numVal, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for num")
	}
	fmt.Printf("numVal = %d\n", numVal)

	// Write the state to the ledger
	err = stub.PutState(num, []byte(strconv.Itoa(numVal)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (m *InvokeCc) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string
	if len(args) < 3 {
		return shim.Error("parameter length is less than 3")
	}

	chaincodeName := args[0]
	num := args[1]
	ob := args[2]
	if len(args) > 3 {
		channelName = args[3]
	} else {
		channelName = ""
	}
	// 调用链码
	f := "query"
	queryArgs := util.ToChaincodeArgs(f, ob)
	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		return shim.Error("invoke wrong")
	}
	numVal, err := strconv.Atoi(string(response.Payload))
	if err != nil {
		return shim.Error("wrong numval")
	}
	fmt.Printf("sumVal = %d\n", numVal)
	err = stub.PutState(num, []byte(strconv.Itoa(numVal)))
	if err != nil {
		return shim.Error(err.Error())
	}
	ret := fmt.Sprintf("invoke success: %s value is %d\n", ob, numVal)
	return shim.Success([]byte(ret))
}

func (m *InvokeCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	f, args := stub.GetFunctionAndParameters()

	if f == "invoke" {
		return m.invoke(stub, args)
	}

	return shim.Error("function name wrong")
}

func main() {
	err := shim.Start(new(InvokeCc))
	if err != nil {
		fmt.Println("start error")
	}
}
