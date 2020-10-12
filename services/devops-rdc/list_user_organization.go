package devops_rdc

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

// ListUserOrganization invokes the devops_rdc.ListUserOrganization API synchronously
func (client *Client) ListUserOrganization(request *ListUserOrganizationRequest) (response *ListUserOrganizationResponse, err error) {
	response = CreateListUserOrganizationResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserOrganizationWithChan invokes the devops_rdc.ListUserOrganization API asynchronously
func (client *Client) ListUserOrganizationWithChan(request *ListUserOrganizationRequest) (<-chan *ListUserOrganizationResponse, <-chan error) {
	responseChan := make(chan *ListUserOrganizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserOrganization(request)
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

// ListUserOrganizationWithCallback invokes the devops_rdc.ListUserOrganization API asynchronously
func (client *Client) ListUserOrganizationWithCallback(request *ListUserOrganizationRequest, callback func(response *ListUserOrganizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserOrganizationResponse
		var err error
		defer close(result)
		response, err = client.ListUserOrganization(request)
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

// ListUserOrganizationRequest is the request struct for api ListUserOrganization
type ListUserOrganizationRequest struct {
	*requests.RpcRequest
	RealPk string `position:"Body" name:"RealPk"`
}

// ListUserOrganizationResponse is the response struct for api ListUserOrganization
type ListUserOrganizationResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       string `json:"Object" xml:"Object"`
}

// CreateListUserOrganizationRequest creates a request to invoke ListUserOrganization API
func CreateListUserOrganizationRequest() (request *ListUserOrganizationRequest) {
	request = &ListUserOrganizationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "ListUserOrganization", "", "")
	request.Method = requests.POST
	return
}

// CreateListUserOrganizationResponse creates a response to parse from ListUserOrganization response
func CreateListUserOrganizationResponse() (response *ListUserOrganizationResponse) {
	response = &ListUserOrganizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
