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

// DescribeScriptVoiceConfig invokes the outboundbot.DescribeScriptVoiceConfig API synchronously
func (client *Client) DescribeScriptVoiceConfig(request *DescribeScriptVoiceConfigRequest) (response *DescribeScriptVoiceConfigResponse, err error) {
	response = CreateDescribeScriptVoiceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScriptVoiceConfigWithChan invokes the outboundbot.DescribeScriptVoiceConfig API asynchronously
func (client *Client) DescribeScriptVoiceConfigWithChan(request *DescribeScriptVoiceConfigRequest) (<-chan *DescribeScriptVoiceConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeScriptVoiceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScriptVoiceConfig(request)
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

// DescribeScriptVoiceConfigWithCallback invokes the outboundbot.DescribeScriptVoiceConfig API asynchronously
func (client *Client) DescribeScriptVoiceConfigWithCallback(request *DescribeScriptVoiceConfigRequest, callback func(response *DescribeScriptVoiceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScriptVoiceConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeScriptVoiceConfig(request)
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

// DescribeScriptVoiceConfigRequest is the request struct for api DescribeScriptVoiceConfig
type DescribeScriptVoiceConfigRequest struct {
	*requests.RpcRequest
	ScriptId            string `position:"Query" name:"ScriptId"`
	ScriptVoiceConfigId string `position:"Query" name:"ScriptVoiceConfigId"`
	InstanceId          string `position:"Query" name:"InstanceId"`
}

// DescribeScriptVoiceConfigResponse is the response struct for api DescribeScriptVoiceConfig
type DescribeScriptVoiceConfigResponse struct {
	*responses.BaseResponse
	HttpStatusCode    int               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code              string            `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	ScriptVoiceConfig ScriptVoiceConfig `json:"ScriptVoiceConfig" xml:"ScriptVoiceConfig"`
}

// CreateDescribeScriptVoiceConfigRequest creates a request to invoke DescribeScriptVoiceConfig API
func CreateDescribeScriptVoiceConfigRequest() (request *DescribeScriptVoiceConfigRequest) {
	request = &DescribeScriptVoiceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DescribeScriptVoiceConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScriptVoiceConfigResponse creates a response to parse from DescribeScriptVoiceConfig response
func CreateDescribeScriptVoiceConfigResponse() (response *DescribeScriptVoiceConfigResponse) {
	response = &DescribeScriptVoiceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
