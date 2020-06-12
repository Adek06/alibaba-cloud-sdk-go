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

// QueryDeviceGroupTagList invokes the iot.QueryDeviceGroupTagList API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicegrouptaglist.html
func (client *Client) QueryDeviceGroupTagList(request *QueryDeviceGroupTagListRequest) (response *QueryDeviceGroupTagListResponse, err error) {
	response = CreateQueryDeviceGroupTagListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceGroupTagListWithChan invokes the iot.QueryDeviceGroupTagList API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicegrouptaglist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceGroupTagListWithChan(request *QueryDeviceGroupTagListRequest) (<-chan *QueryDeviceGroupTagListResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceGroupTagListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceGroupTagList(request)
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

// QueryDeviceGroupTagListWithCallback invokes the iot.QueryDeviceGroupTagList API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicegrouptaglist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceGroupTagListWithCallback(request *QueryDeviceGroupTagListRequest, callback func(response *QueryDeviceGroupTagListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceGroupTagListResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceGroupTagList(request)
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

// QueryDeviceGroupTagListRequest is the request struct for api QueryDeviceGroupTagList
type QueryDeviceGroupTagListRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	GroupId       string `position:"Query" name:"GroupId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// QueryDeviceGroupTagListResponse is the response struct for api QueryDeviceGroupTagList
type QueryDeviceGroupTagListResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceGroupTagList `json:"Data" xml:"Data"`
}

// CreateQueryDeviceGroupTagListRequest creates a request to invoke QueryDeviceGroupTagList API
func CreateQueryDeviceGroupTagListRequest() (request *QueryDeviceGroupTagListRequest) {
	request = &QueryDeviceGroupTagListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceGroupTagList", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceGroupTagListResponse creates a response to parse from QueryDeviceGroupTagList response
func CreateQueryDeviceGroupTagListResponse() (response *QueryDeviceGroupTagListResponse) {
	response = &QueryDeviceGroupTagListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
