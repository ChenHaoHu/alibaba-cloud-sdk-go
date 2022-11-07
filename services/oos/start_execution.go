package oos

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

// StartExecution invokes the oos.StartExecution API synchronously
func (client *Client) StartExecution(request *StartExecutionRequest) (response *StartExecutionResponse, err error) {
	response = CreateStartExecutionResponse()
	err = client.DoAction(request, response)
	return
}

// StartExecutionWithChan invokes the oos.StartExecution API asynchronously
func (client *Client) StartExecutionWithChan(request *StartExecutionRequest) (<-chan *StartExecutionResponse, <-chan error) {
	responseChan := make(chan *StartExecutionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartExecution(request)
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

// StartExecutionWithCallback invokes the oos.StartExecution API asynchronously
func (client *Client) StartExecutionWithCallback(request *StartExecutionRequest, callback func(response *StartExecutionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartExecutionResponse
		var err error
		defer close(result)
		response, err = client.StartExecution(request)
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

// StartExecutionRequest is the request struct for api StartExecution
type StartExecutionRequest struct {
	*requests.RpcRequest
	ClientToken       string                 `position:"Query" name:"ClientToken"`
	Description       string                 `position:"Query" name:"Description"`
	TemplateURL       string                 `position:"Query" name:"TemplateURL"`
	Mode              string                 `position:"Query" name:"Mode"`
	ResourceGroupId   string                 `position:"Query" name:"ResourceGroupId"`
	TemplateVersion   string                 `position:"Query" name:"TemplateVersion"`
	TemplateName      string                 `position:"Query" name:"TemplateName"`
	LoopMode          string                 `position:"Query" name:"LoopMode"`
	SafetyCheck       string                 `position:"Query" name:"SafetyCheck"`
	Tags              map[string]interface{} `position:"Query" name:"Tags"`
	TemplateContent   string                 `position:"Query" name:"TemplateContent"`
	ParentExecutionId string                 `position:"Query" name:"ParentExecutionId"`
	Parameters        string                 `position:"Query" name:"Parameters"`
}

// StartExecutionResponse is the response struct for api StartExecution
type StartExecutionResponse struct {
	*responses.BaseResponse
	RequestId string                    `json:"RequestId" xml:"RequestId"`
	Execution ExecutionInStartExecution `json:"Execution" xml:"Execution"`
}

// CreateStartExecutionRequest creates a request to invoke StartExecution API
func CreateStartExecutionRequest() (request *StartExecutionRequest) {
	request = &StartExecutionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "StartExecution", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartExecutionResponse creates a response to parse from StartExecution response
func CreateStartExecutionResponse() (response *StartExecutionResponse) {
	response = &StartExecutionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
