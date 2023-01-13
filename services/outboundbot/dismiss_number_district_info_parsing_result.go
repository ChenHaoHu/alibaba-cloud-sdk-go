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

// DismissNumberDistrictInfoParsingResult invokes the outboundbot.DismissNumberDistrictInfoParsingResult API synchronously
func (client *Client) DismissNumberDistrictInfoParsingResult(request *DismissNumberDistrictInfoParsingResultRequest) (response *DismissNumberDistrictInfoParsingResultResponse, err error) {
	response = CreateDismissNumberDistrictInfoParsingResultResponse()
	err = client.DoAction(request, response)
	return
}

// DismissNumberDistrictInfoParsingResultWithChan invokes the outboundbot.DismissNumberDistrictInfoParsingResult API asynchronously
func (client *Client) DismissNumberDistrictInfoParsingResultWithChan(request *DismissNumberDistrictInfoParsingResultRequest) (<-chan *DismissNumberDistrictInfoParsingResultResponse, <-chan error) {
	responseChan := make(chan *DismissNumberDistrictInfoParsingResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DismissNumberDistrictInfoParsingResult(request)
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

// DismissNumberDistrictInfoParsingResultWithCallback invokes the outboundbot.DismissNumberDistrictInfoParsingResult API asynchronously
func (client *Client) DismissNumberDistrictInfoParsingResultWithCallback(request *DismissNumberDistrictInfoParsingResultRequest, callback func(response *DismissNumberDistrictInfoParsingResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DismissNumberDistrictInfoParsingResultResponse
		var err error
		defer close(result)
		response, err = client.DismissNumberDistrictInfoParsingResult(request)
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

// DismissNumberDistrictInfoParsingResultRequest is the request struct for api DismissNumberDistrictInfoParsingResult
type DismissNumberDistrictInfoParsingResultRequest struct {
	*requests.RpcRequest
	VersionId string `position:"Query" name:"VersionId"`
}

// DismissNumberDistrictInfoParsingResultResponse is the response struct for api DismissNumberDistrictInfoParsingResult
type DismissNumberDistrictInfoParsingResultResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDismissNumberDistrictInfoParsingResultRequest creates a request to invoke DismissNumberDistrictInfoParsingResult API
func CreateDismissNumberDistrictInfoParsingResultRequest() (request *DismissNumberDistrictInfoParsingResultRequest) {
	request = &DismissNumberDistrictInfoParsingResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DismissNumberDistrictInfoParsingResult", "", "")
	request.Method = requests.POST
	return
}

// CreateDismissNumberDistrictInfoParsingResultResponse creates a response to parse from DismissNumberDistrictInfoParsingResult response
func CreateDismissNumberDistrictInfoParsingResultResponse() (response *DismissNumberDistrictInfoParsingResultResponse) {
	response = &DismissNumberDistrictInfoParsingResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
