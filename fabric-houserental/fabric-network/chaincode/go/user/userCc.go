package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type UserCc struct{}

/*
	账号管理：
		key：用户名-密码-角色-年龄-性别
		value：
			角色：1. police 2. authority（housemng） 3.credit 4. 个人
			用户名
			密码
			年龄
			性别
*/

func (p *UserCc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 5 {
		return shim.Error("parameter length must be 5, including role, username, pwd, age, sex")
	}
	// 角色
	role := args[0]
	// 用户名
	username := args[1]
	// 密码
	pwd := args[2]
	// 年龄
	age := args[3]
	// 性别
	sex := args[4]

	uvalue := fmt.Sprintf("{'role':%s,'username':%s,'pwd':%s,'age':%s,'sex':%s}", role, username, pwd, age, sex)
	k := fmt.Sprintf("%s-%s-%s-%s-%s", role, username, pwd, age, sex)
	err := stub.PutState(k, []byte(uvalue))

	if err != nil {
		return shim.Error("init error")
	}

	return shim.Success([]byte("init success"))
}

// 新增用户
func (p *UserCc) setValue(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("parameter length must be 5, including role, username, pwd, age, sex")
	}
	// 角色
	role := args[0]
	// 用户名
	username := args[1]
	// 密码
	pwd := args[2]
	// 年龄
	age := args[3]
	// 性别
	sex := args[4]
	k := fmt.Sprintf("%s-%s-%s-%s-%s", role, username, pwd, age, sex)
	// if ledger doesn't have this info, do the follows, whether return error
	_, err := stub.GetState(k)

	// if err == nil, the user has existed
	if err == nil {
		return shim.Error("this user has existed")
	}
	uvalue := fmt.Sprintf("{'role':%s,'username':%s,'pwd':%s,'age':%s,'sex':%s}", role, username, pwd, age, sex)
	err = stub.PutState(k, []byte(uvalue))

	if err != nil {
		return shim.Error("create error")
	}

	return shim.Success([]byte("create success"))
}

// 查询用户
func (p *UserCc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("parameter length must be 1")
	}

	user := args[0]
	userBytes, err := stub.GetState(user)

	if err != nil {
		return shim.Error("donot have this user")
	}

	return shim.Success(userBytes)
}

func (p *UserCc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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
	err := shim.Start(new(UserCc))

	if err != nil {
		fmt.Println("cannot start")
	}
}
