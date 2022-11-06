package edas

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

// UpdateApplicationScalingRule invokes the edas.UpdateApplicationScalingRule API synchronously
func (client *Client) UpdateApplicationScalingRule(request *UpdateApplicationScalingRuleRequest) (response *UpdateApplicationScalingRuleResponse, err error) {
	response = CreateUpdateApplicationScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateApplicationScalingRuleWithChan invokes the edas.UpdateApplicationScalingRule API asynchronously
func (client *Client) UpdateApplicationScalingRuleWithChan(request *UpdateApplicationScalingRuleRequest) (<-chan *UpdateApplicationScalingRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateApplicationScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateApplicationScalingRule(request)
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

// UpdateApplicationScalingRuleWithCallback invokes the edas.UpdateApplicationScalingRule API asynchronously
func (client *Client) UpdateApplicationScalingRuleWithCallback(request *UpdateApplicationScalingRuleRequest, callback func(response *UpdateApplicationScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateApplicationScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateApplicationScalingRule(request)
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

// UpdateApplicationScalingRuleRequest is the request struct for api UpdateApplicationScalingRule
type UpdateApplicationScalingRuleRequest struct {
	*requests.RoaRequest
	ScalingRuleName    string `position:"Query" name:"ScalingRuleName"`
	ScalingRuleEnable  string `position:"Query" name:"ScalingRuleEnable"`
	ScalingRuleTimer   string `position:"Query" name:"ScalingRuleTimer"`
	ScalingRuleMetric  string `position:"Query" name:"ScalingRuleMetric"`
	AppId              string `position:"Query" name:"AppId"`
	ScalingRuleTrigger string `position:"Query" name:"ScalingRuleTrigger"`
	ScalingRuleType    string `position:"Query" name:"ScalingRuleType"`
}

// UpdateApplicationScalingRuleResponse is the response struct for api UpdateApplicationScalingRule
type UpdateApplicationScalingRuleResponse struct {
	*responses.BaseResponse
	Code           int            `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	AppScalingRule AppScalingRule `json:"AppScalingRule" xml:"AppScalingRule"`
}

// CreateUpdateApplicationScalingRuleRequest creates a request to invoke UpdateApplicationScalingRule API
func CreateUpdateApplicationScalingRuleRequest() (request *UpdateApplicationScalingRuleRequest) {
	request = &UpdateApplicationScalingRuleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateApplicationScalingRule", "/pop/v1/eam/scale/application_scaling_rule", "edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateApplicationScalingRuleResponse creates a response to parse from UpdateApplicationScalingRule response
func CreateUpdateApplicationScalingRuleResponse() (response *UpdateApplicationScalingRuleResponse) {
	response = &UpdateApplicationScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
