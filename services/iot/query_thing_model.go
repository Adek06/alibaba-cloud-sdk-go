package iot

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

// QueryThingModel invokes the iot.QueryThingModel API synchronously
// api document: https://help.aliyun.com/api/iot/querythingmodel.html
func (client *Client) QueryThingModel(request *QueryThingModelRequest) (response *QueryThingModelResponse, err error) {
	response = CreateQueryThingModelResponse()
	err = client.DoAction(request, response)
	return
}

// QueryThingModelWithChan invokes the iot.QueryThingModel API asynchronously
// api document: https://help.aliyun.com/api/iot/querythingmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryThingModelWithChan(request *QueryThingModelRequest) (<-chan *QueryThingModelResponse, <-chan error) {
	responseChan := make(chan *QueryThingModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryThingModel(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryThingModelWithCallback invokes the iot.QueryThingModel API asynchronously
// api document: https://help.aliyun.com/api/iot/querythingmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryThingModelWithCallback(request *QueryThingModelRequest, callback func(response *QueryThingModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryThingModelResponse
		var err error
		defer close(result)
		response, err = client.QueryThingModel(request)
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

// QueryThingModelRequest is the request struct for api QueryThingModel
type QueryThingModelRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	IotInstanceId   string `position:"Query" name:"IotInstanceId"`
	ProductKey      string `position:"Query" name:"ProductKey"`
	ApiProduct      string `position:"Body" name:"ApiProduct"`
	ApiRevision     string `position:"Body" name:"ApiRevision"`
	ModelVersion    string `position:"Query" name:"ModelVersion"`
}

// QueryThingModelResponse is the response struct for api QueryThingModel
type QueryThingModelResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ProductKey   string `json:"ProductKey" xml:"ProductKey"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryThingModelRequest creates a request to invoke QueryThingModel API
func CreateQueryThingModelRequest() (request *QueryThingModelRequest) {
	request = &QueryThingModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryThingModel", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryThingModelResponse creates a response to parse from QueryThingModel response
func CreateQueryThingModelResponse() (response *QueryThingModelResponse) {
	response = &QueryThingModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
