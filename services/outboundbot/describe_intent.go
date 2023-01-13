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

// DescribeIntent invokes the outboundbot.DescribeIntent API synchronously
func (client *Client) DescribeIntent(request *DescribeIntentRequest) (response *DescribeIntentResponse, err error) {
	response = CreateDescribeIntentResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIntentWithChan invokes the outboundbot.DescribeIntent API asynchronously
func (client *Client) DescribeIntentWithChan(request *DescribeIntentRequest) (<-chan *DescribeIntentResponse, <-chan error) {
	responseChan := make(chan *DescribeIntentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIntent(request)
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

// DescribeIntentWithCallback invokes the outboundbot.DescribeIntent API asynchronously
func (client *Client) DescribeIntentWithCallback(request *DescribeIntentRequest, callback func(response *DescribeIntentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIntentResponse
		var err error
		defer close(result)
		response, err = client.DescribeIntent(request)
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

// DescribeIntentRequest is the request struct for api DescribeIntent
type DescribeIntentRequest struct {
	*requests.RpcRequest
	IntentId   string `position:"Query" name:"IntentId"`
	ScriptId   string `position:"Query" name:"ScriptId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeIntentResponse is the response struct for api DescribeIntent
type DescribeIntentResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Intent         Intent `json:"Intent" xml:"Intent"`
}

// CreateDescribeIntentRequest creates a request to invoke DescribeIntent API
func CreateDescribeIntentRequest() (request *DescribeIntentRequest) {
	request = &DescribeIntentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DescribeIntent", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeIntentResponse creates a response to parse from DescribeIntent response
func CreateDescribeIntentResponse() (response *DescribeIntentResponse) {
	response = &DescribeIntentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
