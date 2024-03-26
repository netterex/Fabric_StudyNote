package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type HousemngCc struct{}

func (p *HousemngCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	// id_house, owner, addr_house, type_house
	if len(args) != 4 {
		return shim.Error("parameter length must be 4, including id_house, owner, addr_house, type_house")
	}
	// 房屋编号
	id_house := args[0]
	// 房屋所有人
	owner := args[1]
	// 房屋地址
	addr_house := args[2]
	// 房屋类型
	type_house := args[3]

	hvalue := fmt.Sprintf("{\"owner\":\"%s\",\"addr_house\":\"%s\",\"type_house\":\"%s\"}", owner, addr_house, type_house)
	err := stub.PutState(id_house, []byte(hvalue))

	if err != nil {
		return shim.Error("init error")
	}

	return shim.Success([]byte("init success"))
}

// 新增房屋
func (p *HousemngCc) setValue(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("parameter length must be 4, including id_house, owner, addr_house, type_house")
	}
	// 房屋类型
	id_house := args[0]
	// 房屋所有人
	owner := args[1]
	// 房屋地址
	addr_house := args[2]
	// 房屋类型
	type_house := args[3]

	// if ledger doesn't have this house, do the follows, whether return error
	resp, err := stub.GetState(id_house)

	if string(resp) != "" {
		return shim.Error("this house has existed")
	}
	hvalue := fmt.Sprintf("{\"owner\":\"%s\",\"addr_house\":\"%s\",\"type_house\":\"%s\"}", owner, addr_house, type_house)
	err = stub.PutState(id_house, []byte(hvalue))

	if err != nil {
		return shim.Error("create error")
	}

	return shim.Success([]byte("create success"))
}

// 查询房屋
func (p *HousemngCc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("parameter length must be 1")
	}

	id_house := args[0]
	idBytes, err := stub.GetState(id_house)

	if err != nil || idBytes == nil {
		return shim.Error("donot have this house")
	}

	return shim.Success(idBytes)
}

func (p *HousemngCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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
	err := shim.Start(new(HousemngCc))

	if err != nil {
		fmt.Println("cannot start")
	}
}
