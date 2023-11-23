package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type FirstCc struct {
}

/*
传两个字段以及对应的值
*/
func (m *FirstCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	// 判断长度
	if len(args) != 4 {
		return shim.Error("the length of args must be four")
	}

	key1 := args[0]
	key1_value := args[1]
	key2 := args[2]
	key2_value, err := strconv.Atoi(args[3])

	if err != nil {
		return shim.Error("the forth arg must be integar")
	}

	err = stub.PutState(key1, []byte(key1_value))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(key2, []byte(strconv.Itoa(key2_value)))
	if err != nil {
		return shim.Error(err.Error())
	}

	ret := fmt.Sprintf("success init, your %s is %s, your %s is %d\n", key1, key1_value, key2, key2_value)

	return shim.Success([]byte(ret))
}

func (m *FirstCc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("expect the function name, length of args must be one")
	}
	A := args[0]
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("cannot get state of " + A)
	}
	fmt.Printf("%s's value is %s", A, string(Avalbytes))
	ret := fmt.Sprintf("success query, %s's value is %s", A, Avalbytes)
	return shim.Success([]byte(ret))
}

func (m *FirstCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	f, args := stub.GetFunctionAndParameters()

	if f == "query" {
		return m.query(stub, args)
	}

	return shim.Error("function parameter must be query")
}

func main() {
	err := shim.Start(new(FirstCc))
	if err != nil {
		fmt.Printf("error starting chaincode01: %s", err)
	}
}
