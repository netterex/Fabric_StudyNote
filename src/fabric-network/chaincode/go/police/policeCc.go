package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type PoliceCc struct{}

func (p *PoliceCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	// id_card, name, age, address
	if len(args) != 4 {
		return shim.Error("parameter length must be 4, including id_card, name, age, address")
	}
	// 身份证号
	id_card := args[0]
	// 姓名
	name := args[1]
	// 年龄
	age := args[2]
	// 住址
	addr := args[3]

	ivalue := fmt.Sprintf("{\"name\":\"%s\",\"age\":\"%s\",\"addr\":\"%s\"}", name, age, addr)
	err := stub.PutState(id_card, []byte(ivalue))

	if err != nil {
		return shim.Error("init error")
	}

	return shim.Success([]byte("init success"))
}

// 新增账户
func (p *PoliceCc) setValue(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("parameter length must be 4, including id_card, name, age, address")
	}
	// 身份证号
	id_card := args[0]
	// 姓名
	name := args[1]
	// 年龄
	age := args[2]
	// 住址
	addr := args[3]

	// if ledger doesn't have this id, do the follows, whether return error
	resp, err := stub.GetState(id_card)

	if string(resp) != "" {
		return shim.Error("this account has existed")
	}
	ivalue := fmt.Sprintf("{\"name\":\"%s\",\"age\":\"%s\",\"addr\":\"%s\"}", name, age, addr)
	err = stub.PutState(id_card, []byte(ivalue))

	if err != nil {
		return shim.Error("create error")
	}

	return shim.Success([]byte("create success"))
}

// 查询账户
func (p *PoliceCc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (p *PoliceCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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
	err := shim.Start(new(PoliceCc))

	if err != nil {
		fmt.Println("cannot start")
	}
}
