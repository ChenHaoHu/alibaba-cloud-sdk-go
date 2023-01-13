package outboundbot

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

// ModifyGlobalQuestion invokes the outboundbot.ModifyGlobalQuestion API synchronously
func (client *Client) ModifyGlobalQuestion(request *ModifyGlobalQuestionRequest) (response *ModifyGlobalQuestionResponse, err error) {
	response = CreateModifyGlobalQuestionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGlobalQuestionWithChan invokes the outboundbot.ModifyGlobalQuestion API asynchronously
func (client *Client) ModifyGlobalQuestionWithChan(request *ModifyGlobalQuestionRequest) (<-chan *ModifyGlobalQuestionResponse, <-chan error) {
	responseChan := make(chan *ModifyGlobalQuestionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGlobalQuestion(request)
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

// ModifyGlobalQuestionWithCallback invokes the outboundbot.ModifyGlobalQuestion API asynchronously
func (client *Client) ModifyGlobalQuestionWithCallback(request *ModifyGlobalQuestionRequest, callback func(response *ModifyGlobalQuestionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGlobalQuestionResponse
		var err error
		defer close(result)
		response, err = client.ModifyGlobalQuestion(request)
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

// ModifyGlobalQuestionRequest is the request struct for api ModifyGlobalQuestion
type ModifyGlobalQuestionRequest struct {
	*requests.RpcRequest
	GlobalQuestionId   string `position:"Query" name:"GlobalQuestionId"`
	GlobalQuestionName string `position:"Query" name:"GlobalQuestionName"`
	Questions          string `position:"Query" name:"Questions"`
	Answers            string `position:"Query" name:"Answers"`
	ScriptId           string `position:"Query" name:"ScriptId"`
	InstanceId         string `position:"Query" name:"InstanceId"`
	GlobalQuestionType string `position:"Query" name:"GlobalQuestionType"`
}

// ModifyGlobalQuestionResponse is the response struct for api ModifyGlobalQuestion
type ModifyGlobalQuestionResponse struct {
	*responses.BaseResponse
	HttpStatusCode     int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
	Success            bool   `json:"Success" xml:"Success"`
	Code               string `json:"Code" xml:"Code"`
	Message            string `json:"Message" xml:"Message"`
	DialogueQuestionId string `json:"DialogueQuestionId" xml:"DialogueQuestionId"`
}

// CreateModifyGlobalQuestionRequest creates a request to invoke ModifyGlobalQuestion API
func CreateModifyGlobalQuestionRequest() (request *ModifyGlobalQuestionRequest) {
	request = &ModifyGlobalQuestionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ModifyGlobalQuestion", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyGlobalQuestionResponse creates a response to parse from ModifyGlobalQuestion response
func CreateModifyGlobalQuestionResponse() (response *ModifyGlobalQuestionResponse) {
	response = &ModifyGlobalQuestionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
