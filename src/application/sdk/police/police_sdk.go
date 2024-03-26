package police

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"time"
)

var (
	// interface
	police_sdk *fabsdk.FabricSDK

	config_path string = "./conf/sdk_conf/police-conf.yaml"
	channel_id  string = "policechannel"
	peer_domain string = "peer0.police.example.com" // cli.Execute(channel.Request{}, channel.WithTargetEndpoints(peer_domain))
	org         string = "PoliceOrg"
	user        string = "Admin"
	err         error
)

func init() {
	police_sdk, err = fabsdk.New(config.FromFile(config_path))

	if err != nil {
		ret_err := fmt.Sprintf("Cannot read config file! Error message: ", err)
		panic(ret_err)
	}

}

func ChannelExecute(chaincode_name string, fcn string, args [][]byte) (channel.Response, error) {
	context := police_sdk.ChannelContext(channel_id, fabsdk.WithOrg(org), fabsdk.WithUser(user))

	cli, err := channel.New(context)

	resp, err := cli.Execute(channel.Request{
		ChaincodeID: chaincode_name,
		Fcn:         fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(peer_domain), channel.WithTimeout(fab.TimeoutType(3), time.Second*2))

	if err != nil && error.Error(err) != "Client Status Code: (5) TIMEOUT. Description: request timed out or been cancelled" {
		ret_err := fmt.Sprintf("Execute wrong! Error message: ", err)
		fmt.Println(ret_err)
	}

	return resp, err
}

func ChannelQuery(chaincode_name string, fcn string, args [][]byte) (channel.Response, error) {
	context := police_sdk.ChannelContext(channel_id, fabsdk.WithOrg(org), fabsdk.WithUser(user))

	cli, err := channel.New(context)

	resp, err := cli.Query(channel.Request{
		ChaincodeID: chaincode_name,
		Fcn:         fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(peer_domain))

	if err != nil {
		ret_err := fmt.Sprintf("Query wrong! Error message: ", err)
		fmt.Println(ret_err)
	}

	return resp, err
}
