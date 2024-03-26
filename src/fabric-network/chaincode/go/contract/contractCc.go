package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ContractCc struct{}

/*
	合同管理：
		key：合同编号
		value：租户身份证号，房屋 id ，合同图片 hash
*/

func (p *ContractCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	// id_contract, id_tenant, id_house, hash_contractImg
	if len(args) != 4 {
		return shim.Error("parameter length must be 4, including id_contract, id_tenant, id_house, hash_contractImg")
	}
	// 合同编号
	id_contract := args[0]
	// 租户身份证号
	id_tenant := args[1]
	// 房屋编号
	id_house := args[2]
	// 合同图片hash
	hash_contractImg := args[3]

	cvalue := fmt.Sprintf("{\"id_tenant\":\"%s\",\"id_house\":\"%s\",\"hash_contractImg\":\"%s\"}", id_tenant, id_house, hash_contractImg)
	err := stub.PutState(id_contract, []byte(cvalue))

	if err != nil {
		return shim.Error("init error")
	}

	return shim.Success([]byte("init success"))
}

// 新增合同
func (p *ContractCc) setValue(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("parameter length must be 4, including id_contract, id_tenant, id_house, hash_contractImg")
	}
	// 合同编号
	id_contract := args[0]
	// 租户身份证号
	id_tenant := args[1]
	// 房屋编号
	id_house := args[2]
	// 合同图片hash
	hash_contractImg := args[3]

	// if ledger doesn't have this contract, do the follows, whether return error
	resp, err := stub.GetState(id_contract)

	if string(resp) != "" {
		return shim.Error("this contract has existed")
	}
	cvalue := fmt.Sprintf("{\"id_tenant\":\"%s\",\"id_house\":\"%s\",\"hash_contractImg\":\"%s\"}", id_tenant, id_house, hash_contractImg)
	err = stub.PutState(id_contract, []byte(cvalue))

	if err != nil {
		return shim.Error("create error")
	}

	return shim.Success([]byte("create success"))
}

// 查询合同
func (p *ContractCc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("parameter length must be 1")
	}

	id_contract := args[0]
	idBytes, err := stub.GetState(id_contract)

	if err != nil || idBytes == nil {
		return shim.Error("donot have this contract")
	}

	return shim.Success(idBytes)
}

func (p *ContractCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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
	err := shim.Start(new(ContractCc))

	if err != nil {
		fmt.Println("cannot start")
	}
}
