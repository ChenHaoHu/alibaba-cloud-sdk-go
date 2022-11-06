package cloudesl

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

// GetAddMaterialStatus invokes the cloudesl.GetAddMaterialStatus API synchronously
func (client *Client) GetAddMaterialStatus(request *GetAddMaterialStatusRequest) (response *GetAddMaterialStatusResponse, err error) {
	response = CreateGetAddMaterialStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetAddMaterialStatusWithChan invokes the cloudesl.GetAddMaterialStatus API asynchronously
func (client *Client) GetAddMaterialStatusWithChan(request *GetAddMaterialStatusRequest) (<-chan *GetAddMaterialStatusResponse, <-chan error) {
	responseChan := make(chan *GetAddMaterialStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAddMaterialStatus(request)
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

// GetAddMaterialStatusWithCallback invokes the cloudesl.GetAddMaterialStatus API asynchronously
func (client *Client) GetAddMaterialStatusWithCallback(request *GetAddMaterialStatusRequest, callback func(response *GetAddMaterialStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAddMaterialStatusResponse
		var err error
		defer close(result)
		response, err = client.GetAddMaterialStatus(request)
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

// GetAddMaterialStatusRequest is the request struct for api GetAddMaterialStatus
type GetAddMaterialStatusRequest struct {
	*requests.RpcRequest
}

// GetAddMaterialStatusResponse is the response struct for api GetAddMaterialStatus
type GetAddMaterialStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetAddMaterialStatusRequest creates a request to invoke GetAddMaterialStatus API
func CreateGetAddMaterialStatusRequest() (request *GetAddMaterialStatusRequest) {
	request = &GetAddMaterialStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "GetAddMaterialStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAddMaterialStatusResponse creates a response to parse from GetAddMaterialStatus response
func CreateGetAddMaterialStatusResponse() (response *GetAddMaterialStatusResponse) {
	response = &GetAddMaterialStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
