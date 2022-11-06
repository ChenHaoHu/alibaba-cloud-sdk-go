package alikafka

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

// UpdateInstanceConfig invokes the alikafka.UpdateInstanceConfig API synchronously
func (client *Client) UpdateInstanceConfig(request *UpdateInstanceConfigRequest) (response *UpdateInstanceConfigResponse, err error) {
	response = CreateUpdateInstanceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateInstanceConfigWithChan invokes the alikafka.UpdateInstanceConfig API asynchronously
func (client *Client) UpdateInstanceConfigWithChan(request *UpdateInstanceConfigRequest) (<-chan *UpdateInstanceConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateInstanceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateInstanceConfig(request)
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

// UpdateInstanceConfigWithCallback invokes the alikafka.UpdateInstanceConfig API asynchronously
func (client *Client) UpdateInstanceConfigWithCallback(request *UpdateInstanceConfigRequest, callback func(response *UpdateInstanceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateInstanceConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateInstanceConfig(request)
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

// UpdateInstanceConfigRequest is the request struct for api UpdateInstanceConfig
type UpdateInstanceConfigRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Config     string `position:"Query" name:"Config"`
}

// UpdateInstanceConfigResponse is the response struct for api UpdateInstanceConfig
type UpdateInstanceConfigResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateInstanceConfigRequest creates a request to invoke UpdateInstanceConfig API
func CreateUpdateInstanceConfigRequest() (request *UpdateInstanceConfigRequest) {
	request = &UpdateInstanceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "UpdateInstanceConfig", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateInstanceConfigResponse creates a response to parse from UpdateInstanceConfig response
func CreateUpdateInstanceConfigResponse() (response *UpdateInstanceConfigResponse) {
	response = &UpdateInstanceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
