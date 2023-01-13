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

// ImportScript invokes the outboundbot.ImportScript API synchronously
func (client *Client) ImportScript(request *ImportScriptRequest) (response *ImportScriptResponse, err error) {
	response = CreateImportScriptResponse()
	err = client.DoAction(request, response)
	return
}

// ImportScriptWithChan invokes the outboundbot.ImportScript API asynchronously
func (client *Client) ImportScriptWithChan(request *ImportScriptRequest) (<-chan *ImportScriptResponse, <-chan error) {
	responseChan := make(chan *ImportScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportScript(request)
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

// ImportScriptWithCallback invokes the outboundbot.ImportScript API asynchronously
func (client *Client) ImportScriptWithCallback(request *ImportScriptRequest, callback func(response *ImportScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportScriptResponse
		var err error
		defer close(result)
		response, err = client.ImportScript(request)
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

// ImportScriptRequest is the request struct for api ImportScript
type ImportScriptRequest struct {
	*requests.RpcRequest
	SignatureUrl string `position:"Query" name:"SignatureUrl"`
	InstanceId   string `position:"Query" name:"InstanceId"`
}

// ImportScriptResponse is the response struct for api ImportScript
type ImportScriptResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	ScriptId       string `json:"ScriptId" xml:"ScriptId"`
}

// CreateImportScriptRequest creates a request to invoke ImportScript API
func CreateImportScriptRequest() (request *ImportScriptRequest) {
	request = &ImportScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ImportScript", "", "")
	request.Method = requests.POST
	return
}

// CreateImportScriptResponse creates a response to parse from ImportScript response
func CreateImportScriptResponse() (response *ImportScriptResponse) {
	response = &ImportScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
