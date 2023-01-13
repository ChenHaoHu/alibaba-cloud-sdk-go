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

// RecordFailure invokes the outboundbot.RecordFailure API synchronously
func (client *Client) RecordFailure(request *RecordFailureRequest) (response *RecordFailureResponse, err error) {
	response = CreateRecordFailureResponse()
	err = client.DoAction(request, response)
	return
}

// RecordFailureWithChan invokes the outboundbot.RecordFailure API asynchronously
func (client *Client) RecordFailureWithChan(request *RecordFailureRequest) (<-chan *RecordFailureResponse, <-chan error) {
	responseChan := make(chan *RecordFailureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecordFailure(request)
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

// RecordFailureWithCallback invokes the outboundbot.RecordFailure API asynchronously
func (client *Client) RecordFailureWithCallback(request *RecordFailureRequest, callback func(response *RecordFailureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecordFailureResponse
		var err error
		defer close(result)
		response, err = client.RecordFailure(request)
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

// RecordFailureRequest is the request struct for api RecordFailure
type RecordFailureRequest struct {
	*requests.RpcRequest
	CallId          string           `position:"Query" name:"CallId"`
	ActualTime      requests.Integer `position:"Query" name:"ActualTime"`
	CallingNumber   string           `position:"Query" name:"CallingNumber"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	DispositionCode string           `position:"Query" name:"DispositionCode"`
	CalledNumber    string           `position:"Query" name:"CalledNumber"`
	InstanceOwnerId string           `position:"Query" name:"InstanceOwnerId"`
	TaskId          string           `position:"Query" name:"TaskId"`
	ExceptionCodes  string           `position:"Query" name:"ExceptionCodes"`
}

// RecordFailureResponse is the response struct for api RecordFailure
type RecordFailureResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateRecordFailureRequest creates a request to invoke RecordFailure API
func CreateRecordFailureRequest() (request *RecordFailureRequest) {
	request = &RecordFailureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "RecordFailure", "", "")
	request.Method = requests.POST
	return
}

// CreateRecordFailureResponse creates a response to parse from RecordFailure response
func CreateRecordFailureResponse() (response *RecordFailureResponse) {
	response = &RecordFailureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
