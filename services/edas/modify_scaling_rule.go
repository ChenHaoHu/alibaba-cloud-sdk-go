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

// ModifyScalingRule invokes the edas.ModifyScalingRule API synchronously
func (client *Client) ModifyScalingRule(request *ModifyScalingRuleRequest) (response *ModifyScalingRuleResponse, err error) {
	response = CreateModifyScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScalingRuleWithChan invokes the edas.ModifyScalingRule API asynchronously
func (client *Client) ModifyScalingRuleWithChan(request *ModifyScalingRuleRequest) (<-chan *ModifyScalingRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingRule(request)
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

// ModifyScalingRuleWithCallback invokes the edas.ModifyScalingRule API asynchronously
func (client *Client) ModifyScalingRuleWithCallback(request *ModifyScalingRuleRequest, callback func(response *ModifyScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyScalingRule(request)
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

// ModifyScalingRuleRequest is the request struct for api ModifyScalingRule
type ModifyScalingRuleRequest struct {
	*requests.RoaRequest
	InStep               string `position:"Query" name:"InStep"`
	OutInstanceNum       string `position:"Query" name:"OutInstanceNum"`
	OutRT                string `position:"Query" name:"OutRT"`
	InInstanceNum        string `position:"Query" name:"InInstanceNum"`
	VSwitchIds           string `position:"Query" name:"VSwitchIds"`
	TemplateInstanceId   string `position:"Query" name:"TemplateInstanceId"`
	AcceptEULA           string `position:"Query" name:"AcceptEULA"`
	OutStep              string `position:"Query" name:"OutStep"`
	OutCPU               string `position:"Query" name:"OutCPU"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	Password             string `position:"Query" name:"Password"`
	TemplateVersion      string `position:"Query" name:"TemplateVersion"`
	InCondition          string `position:"Query" name:"InCondition"`
	InRT                 string `position:"Query" name:"InRT"`
	InCpu                string `position:"Query" name:"InCpu"`
	OutDuration          string `position:"Query" name:"OutDuration"`
	MultiAzPolicy        string `position:"Query" name:"MultiAzPolicy"`
	OutLoad              string `position:"Query" name:"OutLoad"`
	InLoad               string `position:"Query" name:"InLoad"`
	GroupId              string `position:"Query" name:"GroupId"`
	ResourceFrom         string `position:"Query" name:"ResourceFrom"`
	OutEnable            string `position:"Query" name:"OutEnable"`
	TemplateId           string `position:"Query" name:"TemplateId"`
	ScalingPolicy        string `position:"Query" name:"ScalingPolicy"`
	OutCondition         string `position:"Query" name:"OutCondition"`
	InDuration           string `position:"Query" name:"InDuration"`
	InEnable             string `position:"Query" name:"InEnable"`
	AppId                string `position:"Query" name:"AppId"`
	VpcId                string `position:"Query" name:"VpcId"`
	TemplateInstanceName string `position:"Query" name:"TemplateInstanceName"`
}

// ModifyScalingRuleResponse is the response struct for api ModifyScalingRule
type ModifyScalingRuleResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyScalingRuleRequest creates a request to invoke ModifyScalingRule API
func CreateModifyScalingRuleRequest() (request *ModifyScalingRuleRequest) {
	request = &ModifyScalingRuleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ModifyScalingRule", "/pop/v5/app/scaling_rules", "edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyScalingRuleResponse creates a response to parse from ModifyScalingRule response
func CreateModifyScalingRuleResponse() (response *ModifyScalingRuleResponse) {
	response = &ModifyScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
