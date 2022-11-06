package ahas_openapi

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

// EnableSystemRule invokes the ahas_openapi.EnableSystemRule API synchronously
func (client *Client) EnableSystemRule(request *EnableSystemRuleRequest) (response *EnableSystemRuleResponse, err error) {
	response = CreateEnableSystemRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableSystemRuleWithChan invokes the ahas_openapi.EnableSystemRule API asynchronously
func (client *Client) EnableSystemRuleWithChan(request *EnableSystemRuleRequest) (<-chan *EnableSystemRuleResponse, <-chan error) {
	responseChan := make(chan *EnableSystemRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableSystemRule(request)
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

// EnableSystemRuleWithCallback invokes the ahas_openapi.EnableSystemRule API asynchronously
func (client *Client) EnableSystemRuleWithCallback(request *EnableSystemRuleRequest, callback func(response *EnableSystemRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableSystemRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableSystemRule(request)
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

// EnableSystemRuleRequest is the request struct for api EnableSystemRule
type EnableSystemRuleRequest struct {
	*requests.RpcRequest
	AhasRegionId string           `position:"Query" name:"AhasRegionId"`
	RuleId       requests.Integer `position:"Query" name:"RuleId"`
}

// EnableSystemRuleResponse is the response struct for api EnableSystemRule
type EnableSystemRuleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEnableSystemRuleRequest creates a request to invoke EnableSystemRule API
func CreateEnableSystemRuleRequest() (request *EnableSystemRuleRequest) {
	request = &EnableSystemRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "EnableSystemRule", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableSystemRuleResponse creates a response to parse from EnableSystemRule response
func CreateEnableSystemRuleResponse() (response *EnableSystemRuleResponse) {
	response = &EnableSystemRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
