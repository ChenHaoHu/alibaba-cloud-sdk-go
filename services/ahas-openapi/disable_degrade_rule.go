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

// DisableDegradeRule invokes the ahas_openapi.DisableDegradeRule API synchronously
func (client *Client) DisableDegradeRule(request *DisableDegradeRuleRequest) (response *DisableDegradeRuleResponse, err error) {
	response = CreateDisableDegradeRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DisableDegradeRuleWithChan invokes the ahas_openapi.DisableDegradeRule API asynchronously
func (client *Client) DisableDegradeRuleWithChan(request *DisableDegradeRuleRequest) (<-chan *DisableDegradeRuleResponse, <-chan error) {
	responseChan := make(chan *DisableDegradeRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableDegradeRule(request)
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

// DisableDegradeRuleWithCallback invokes the ahas_openapi.DisableDegradeRule API asynchronously
func (client *Client) DisableDegradeRuleWithCallback(request *DisableDegradeRuleRequest, callback func(response *DisableDegradeRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableDegradeRuleResponse
		var err error
		defer close(result)
		response, err = client.DisableDegradeRule(request)
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

// DisableDegradeRuleRequest is the request struct for api DisableDegradeRule
type DisableDegradeRuleRequest struct {
	*requests.RpcRequest
	AhasRegionId string           `position:"Query" name:"AhasRegionId"`
	RuleId       requests.Integer `position:"Query" name:"RuleId"`
}

// DisableDegradeRuleResponse is the response struct for api DisableDegradeRule
type DisableDegradeRuleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDisableDegradeRuleRequest creates a request to invoke DisableDegradeRule API
func CreateDisableDegradeRuleRequest() (request *DisableDegradeRuleRequest) {
	request = &DisableDegradeRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "DisableDegradeRule", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableDegradeRuleResponse creates a response to parse from DisableDegradeRule response
func CreateDisableDegradeRuleResponse() (response *DisableDegradeRuleResponse) {
	response = &DisableDegradeRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
