package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type CreditCc struct{}

func (p *CreditCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	// 身份证号，征信级别（A，B，C信誉依次递减，C不允许租房）
	if len(args) != 2 {
		return shim.Error("parameter length must be 2, including id_card, creditLevel")
	}
	// 身份证号
	id_card := args[0]
	// 征信级别
	creditLevel := args[1]

	err := stub.PutState(id_card, []byte(creditLevel))

	if err != nil {
		return shim.Error("init error")
	}

	return shim.Success([]byte("init success"))
}

// 新增用户征信信息
func (p *CreditCc) setValue(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("parameter length must be 2, including id_card, creditLevel")
	}
	// 身份证号
	id_card := args[0]
	// 征信级别
	creditLevel := args[1]

	// if ledger doesn't have this id, do the follows, whether return error
	resp, err := stub.GetState(id_card)

	// if id_card doesnot exist, the err is still nil. So judge if the resp is ""
	if string(resp) != "" {
		return shim.Error("this account has existed")
	}

	err = stub.PutState(id_card, []byte(creditLevel))

	if err != nil {
		return shim.Error("create error")
	}

	return shim.Success([]byte("create success"))
}

// 查询用户征信信息
func (p *CreditCc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("parameter length must be 1")
	}

	id_card := args[0]
	idBytes, err := stub.GetState(id_card)

	if err != nil || idBytes == nil {
		return shim.Error("donot have this id")
	}

	return shim.Success(idBytes)
}

func (p *CreditCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	f, args := stub.GetFunctionAndParameters()

	if f == "set" {
		return p.setValue(stub, args)
	}

	if f == "query" {
		return p.query(stub, args)
	}

	return shim.Error("wrong function name")
}

func main() {
	err := shim.Start(new(CreditCc))

	if err != nil {
		fmt.Println("cannot start")
	}
}
