package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
	response = CreateCreateNatGatewayResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateNatGatewayWithChan(request *CreateNatGatewayRequest) (<-chan *CreateNatGatewayResponse, <-chan error) {
	responseChan := make(chan *CreateNatGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNatGateway(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) CreateNatGatewayWithCallback(request *CreateNatGatewayRequest, callback func(response *CreateNatGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNatGatewayResponse
		var err error
		defer close(result)
		response, err = client.CreateNatGateway(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type CreateNatGatewayRequest struct {
	*requests.RpcRequest
	Spec                 string                              `position:"Query" name:"Spec"`
	BandwidthPackage     *[]CreateNatGatewayBandwidthPackage `position:"Query" name:"BandwidthPackage"  type:"Repeated"`
	ClientToken          string                              `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string                              `position:"Query" name:"ResourceOwnerAccount"`
	Description          string                              `position:"Query" name:"Description"`
	Name                 string                              `position:"Query" name:"Name"`
	ResourceOwnerId      requests.Integer                    `position:"Query" name:"ResourceOwnerId"`
	VpcId                string                              `position:"Query" name:"VpcId"`
	OwnerAccount         string                              `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                    `position:"Query" name:"OwnerId"`
}

type CreateNatGatewayBandwidthPackage struct {
	IpCount            string `name:"IpCount"`
	Bandwidth          string `name:"Bandwidth"`
	Zone               string `name:"Zone"`
	ISP                string `name:"ISP"`
	InternetChargeType string `name:"InternetChargeType"`
}

type CreateNatGatewayResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	NatGatewayId    string `json:"NatGatewayId" xml:"NatGatewayId"`
	ForwardTableIds struct {
		ForwardTableId []string `json:"ForwardTableId" xml:"ForwardTableId"`
	} `json:"ForwardTableIds" xml:"ForwardTableIds"`
	SnatTableIds struct {
		SnatTableId []string `json:"SnatTableId" xml:"SnatTableId"`
	} `json:"SnatTableIds" xml:"SnatTableIds"`
	BandwidthPackageIds struct {
		BandwidthPackageId []string `json:"BandwidthPackageId" xml:"BandwidthPackageId"`
	} `json:"BandwidthPackageIds" xml:"BandwidthPackageIds"`
}

func CreateCreateNatGatewayRequest() (request *CreateNatGatewayRequest) {
	request = &CreateNatGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateNatGateway", "vpc", "openAPI")
	return
}

func CreateCreateNatGatewayResponse() (response *CreateNatGatewayResponse) {
	response = &CreateNatGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}