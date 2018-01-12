package mts

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

func (client *Client) AddMediaTag(request *AddMediaTagRequest) (response *AddMediaTagResponse, err error) {
	response = CreateAddMediaTagResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddMediaTagWithChan(request *AddMediaTagRequest) (<-chan *AddMediaTagResponse, <-chan error) {
	responseChan := make(chan *AddMediaTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddMediaTag(request)
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

func (client *Client) AddMediaTagWithCallback(request *AddMediaTagRequest, callback func(response *AddMediaTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddMediaTagResponse
		var err error
		defer close(result)
		response, err = client.AddMediaTag(request)
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

type AddMediaTagRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Tag                  string           `position:"Query" name:"Tag"`
	MediaId              string           `position:"Query" name:"MediaId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type AddMediaTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateAddMediaTagRequest() (request *AddMediaTagRequest) {
	request = &AddMediaTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddMediaTag", "mts", "openAPI")
	return
}

func CreateAddMediaTagResponse() (response *AddMediaTagResponse) {
	response = &AddMediaTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}