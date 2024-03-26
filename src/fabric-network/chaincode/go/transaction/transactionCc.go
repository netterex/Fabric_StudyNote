package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type TxCc struct{}

/*
	交易管理：
		key: 订单编号-期数
		value: 订单编号,分期,房东身份证号,租户身份证号,房东姓名,
		租户姓名,房屋编号,租金,交易时间,合同编号,备注
*/

func (p *TxCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 11 {
		return shim.Error("parameter length must be 11")
	}
	// 订单编号
	id_order := args[0]
	// 分期
	period := args[1]
	// 房东身份证号
	id_landlord := args[2]
	// 租户身份证号
	id_tenant := args[3]
	// 房东姓名
	name_landlord := args[4]
	// 租户姓名
	name_tenant := args[5]
	// 房屋编号
	id_house := args[6]
	// 租金
	rent := args[7]
	// 时间
	time_tx := args[8]
	// 合同编号
	id_contract := args[9]
	// 备注
	desc := args[10]

	tvalue := fmt.Sprintf("{\"id_order\":\"%s\",\"period\":\"%s\",\"id_landlord\":\"%s\",\"id_tenant\":\"%s\",\"name_landlord\":\"%s\",\"name_tenant\":\"%s\",\"id_house\":\"%s\",\"rent\":\"%s\",\"time_tx\":\"%s\",\"id_contract\":\"%s\",\"desc\":\"%s\"}", id_order, period, id_landlord, id_tenant, name_landlord, name_tenant, id_house, rent, time_tx, id_contract, desc)
	k := fmt.Sprintf("%s-%s", id_order, period)
	err := stub.PutState(k, []byte(tvalue))

	if err != nil {
		return shim.Error("init error")
	}

	return shim.Success([]byte("init success"))
}

// 新增订单
func (p *TxCc) setValue(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 11 {
		return shim.Error("parameter length must be 11")
	}
	// 订单编号
	id_order := args[0]
	// 分期
	period := args[1]
	// 房东身份证号
	id_landlord := args[2]
	// 租户身份证号
	id_tenant := args[3]
	// 房东姓名
	name_landlord := args[4]
	// 租户姓名
	name_tenant := args[5]
	// 房屋编号
	id_house := args[6]
	// 租金
	rent := args[7]
	// 时间
	time_tx := args[8]
	// 合同编号
	id_contract := args[9]
	// 备注
	desc := args[10]

	// 订单编号-期数
	k := fmt.Sprintf("%s-%s", id_order, period)

	// if ledger doesn't have this order, do the follows, whether return error
	resp, err := stub.GetState(k)

	if string(resp) != "" {
		return shim.Error("this order has existed")
	}
	tvalue := fmt.Sprintf("{\"id_order\":\"%s\",\"period\":\"%s\",\"id_landlord\":\"%s\",\"id_tenant\":\"%s\",\"name_landlord\":\"%s\",\"name_tenant\":\"%s\",\"id_house\":\"%s\",\"rent\":\"%s\",\"time_tx\":\"%s\",\"id_contract\":\"%s\",\"desc\":\"%s\"}", id_order, period, id_landlord, id_tenant, name_landlord, name_tenant, id_house, rent, time_tx, id_contract, desc)
	err = stub.PutState(k, []byte(tvalue))

	if err != nil {
		return shim.Error("create error")
	}

	return shim.Success([]byte("create success"))
}

// 查询订单
func (p *TxCc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("parameter length must be 2")
	}

	id_order := args[0]
	period := args[1]
	k := fmt.Sprintf("%s-%s", id_order, period)

	idBytes, err := stub.GetState(k)

	if err != nil || idBytes == nil {
		return shim.Error("donot have this order")
	}

	return shim.Success(idBytes)
}

func (p *TxCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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
	err := shim.Start(new(TxCc))

	if err != nil {
		fmt.Println("cannot start")
	}
}
