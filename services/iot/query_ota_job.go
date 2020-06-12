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

// QueryOTAJob invokes the iot.QueryOTAJob API synchronously
// api document: https://help.aliyun.com/api/iot/queryotajob.html
func (client *Client) QueryOTAJob(request *QueryOTAJobRequest) (response *QueryOTAJobResponse, err error) {
	response = CreateQueryOTAJobResponse()
	err = client.DoAction(request, response)
	return
}

// QueryOTAJobWithChan invokes the iot.QueryOTAJob API asynchronously
// api document: https://help.aliyun.com/api/iot/queryotajob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryOTAJobWithChan(request *QueryOTAJobRequest) (<-chan *QueryOTAJobResponse, <-chan error) {
	responseChan := make(chan *QueryOTAJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryOTAJob(request)
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

// QueryOTAJobWithCallback invokes the iot.QueryOTAJob API asynchronously
// api document: https://help.aliyun.com/api/iot/queryotajob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryOTAJobWithCallback(request *QueryOTAJobRequest, callback func(response *QueryOTAJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryOTAJobResponse
		var err error
		defer close(result)
		response, err = client.QueryOTAJob(request)
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

// QueryOTAJobRequest is the request struct for api QueryOTAJob
type QueryOTAJobRequest struct {
	*requests.RpcRequest
	JobId         string `position:"Query" name:"JobId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// QueryOTAJobResponse is the response struct for api QueryOTAJob
type QueryOTAJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryOTAJobRequest creates a request to invoke QueryOTAJob API
func CreateQueryOTAJobRequest() (request *QueryOTAJobRequest) {
	request = &QueryOTAJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryOTAJob", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryOTAJobResponse creates a response to parse from QueryOTAJob response
func CreateQueryOTAJobResponse() (response *QueryOTAJobResponse) {
	response = &QueryOTAJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
