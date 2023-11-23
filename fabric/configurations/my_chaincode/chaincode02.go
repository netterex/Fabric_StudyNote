package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type TxCc struct {
}

/*
Init 初始化两个账户及其余额，共四个参数
Invoke 通过参数判断调用 invoke 还是 query
invoke 实现从第一个参数账户向第二个参数账户转账，共三个参数
query 实现查询某账户剩余余额的功能，共一个参数
*/
func (m *TxCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()
	// 判断参数长度
	if len(args) != 4 {
		return shim.Error("wrong lens")
	}

	A := args[0]
	Aval, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("the second arg must be integar")
	}
	B := args[2]
	Bval, err := strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("the forth arg must be integar")
	}
	// 将初始化数据存入账户对应的账本中
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("success init"))
}

func (m *TxCc) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var X int // 转账金额
	var err error
	// 判断参数长度
	if len(args) != 3 {
		return shim.Error("wrong args lens")
	}
	acFrom := args[0]
	acTo := args[1]
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("the third arg must be integar")
	}
	fromBytes, err := stub.GetState(acFrom)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	/*	// 判断转出账户是否存在
		if fromBytes == nil {
			return shim.Error("Entity not found")
		}*/
	fromval, _ := strconv.Atoi(string(fromBytes))
	// 判断转出账户的余额是否足够
	if fromval < X {
		return shim.Error("the from account's coin is not enough")
	}
	toBytes, err := stub.GetState(acTo)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	/*	// 判断转入账户是否存在
		if toBytes == nil {
			return shim.Error("Entity not found")
		}*/
	toval, _ := strconv.Atoi(string(toBytes))
	fromval = fromval - X
	toval = toval + X

	// 将余额重新写入账本中
	err = stub.PutState(acFrom, []byte(strconv.Itoa(fromval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(acTo, []byte(strconv.Itoa(toval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("invoke success"))
}

func (m *TxCc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("expect the function name, length of args must be one")
	}
	A := args[0]
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("cannot get state of " + A)
	}
	if Avalbytes == nil {
		return shim.Error("entity not found")
	}
	fmt.Printf("%s's value is is %s\n", A, string(Avalbytes))
	ret := fmt.Sprintf("success query, %s's value is %s", A, string(Avalbytes))
	return shim.Success([]byte(ret))
}

func (m *TxCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	f, args := stub.GetFunctionAndParameters()
	// 判断需要执行的函数名是invoke
	if f == "invoke" {
		return m.invoke(stub, args)
	}
	if f == "query" {
		return m.query(stub, args)
	}
	return shim.Error("function parameter wrong")
}

func main() {
	err := shim.Start(new(TxCc))
	if err != nil {
		fmt.Printf("error starting chaincode02: %s", err)
	}
}
