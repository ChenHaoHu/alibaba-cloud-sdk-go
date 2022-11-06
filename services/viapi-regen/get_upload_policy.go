package viapi_regen

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

// GetUploadPolicy invokes the viapi_regen.GetUploadPolicy API synchronously
func (client *Client) GetUploadPolicy(request *GetUploadPolicyRequest) (response *GetUploadPolicyResponse, err error) {
	response = CreateGetUploadPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// GetUploadPolicyWithChan invokes the viapi_regen.GetUploadPolicy API asynchronously
func (client *Client) GetUploadPolicyWithChan(request *GetUploadPolicyRequest) (<-chan *GetUploadPolicyResponse, <-chan error) {
	responseChan := make(chan *GetUploadPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUploadPolicy(request)
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

// GetUploadPolicyWithCallback invokes the viapi_regen.GetUploadPolicy API asynchronously
func (client *Client) GetUploadPolicyWithCallback(request *GetUploadPolicyRequest, callback func(response *GetUploadPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUploadPolicyResponse
		var err error
		defer close(result)
		response, err = client.GetUploadPolicy(request)
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

// GetUploadPolicyRequest is the request struct for api GetUploadPolicy
type GetUploadPolicyRequest struct {
	*requests.RpcRequest
	Type     string           `position:"Body" name:"Type"`
	Id       requests.Integer `position:"Body" name:"Id"`
	FileName string           `position:"Body" name:"FileName"`
}

// GetUploadPolicyResponse is the response struct for api GetUploadPolicy
type GetUploadPolicyResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetUploadPolicyRequest creates a request to invoke GetUploadPolicy API
func CreateGetUploadPolicyRequest() (request *GetUploadPolicyRequest) {
	request = &GetUploadPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "GetUploadPolicy", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUploadPolicyResponse creates a response to parse from GetUploadPolicy response
func CreateGetUploadPolicyResponse() (response *GetUploadPolicyResponse) {
	response = &GetUploadPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
