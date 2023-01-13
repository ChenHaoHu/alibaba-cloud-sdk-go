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

// SaveAfterAnswerDelayPlayback invokes the outboundbot.SaveAfterAnswerDelayPlayback API synchronously
func (client *Client) SaveAfterAnswerDelayPlayback(request *SaveAfterAnswerDelayPlaybackRequest) (response *SaveAfterAnswerDelayPlaybackResponse, err error) {
	response = CreateSaveAfterAnswerDelayPlaybackResponse()
	err = client.DoAction(request, response)
	return
}

// SaveAfterAnswerDelayPlaybackWithChan invokes the outboundbot.SaveAfterAnswerDelayPlayback API asynchronously
func (client *Client) SaveAfterAnswerDelayPlaybackWithChan(request *SaveAfterAnswerDelayPlaybackRequest) (<-chan *SaveAfterAnswerDelayPlaybackResponse, <-chan error) {
	responseChan := make(chan *SaveAfterAnswerDelayPlaybackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveAfterAnswerDelayPlayback(request)
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

// SaveAfterAnswerDelayPlaybackWithCallback invokes the outboundbot.SaveAfterAnswerDelayPlayback API asynchronously
func (client *Client) SaveAfterAnswerDelayPlaybackWithCallback(request *SaveAfterAnswerDelayPlaybackRequest, callback func(response *SaveAfterAnswerDelayPlaybackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveAfterAnswerDelayPlaybackResponse
		var err error
		defer close(result)
		response, err = client.SaveAfterAnswerDelayPlayback(request)
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

// SaveAfterAnswerDelayPlaybackRequest is the request struct for api SaveAfterAnswerDelayPlayback
type SaveAfterAnswerDelayPlaybackRequest struct {
	*requests.RpcRequest
	StrategyLevel            requests.Integer `position:"Query" name:"StrategyLevel"`
	EntryId                  string           `position:"Query" name:"EntryId"`
	AfterAnswerDelayPlayback requests.Integer `position:"Query" name:"AfterAnswerDelayPlayback"`
}

// SaveAfterAnswerDelayPlaybackResponse is the response struct for api SaveAfterAnswerDelayPlayback
type SaveAfterAnswerDelayPlaybackResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateSaveAfterAnswerDelayPlaybackRequest creates a request to invoke SaveAfterAnswerDelayPlayback API
func CreateSaveAfterAnswerDelayPlaybackRequest() (request *SaveAfterAnswerDelayPlaybackRequest) {
	request = &SaveAfterAnswerDelayPlaybackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "SaveAfterAnswerDelayPlayback", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveAfterAnswerDelayPlaybackResponse creates a response to parse from SaveAfterAnswerDelayPlayback response
func CreateSaveAfterAnswerDelayPlaybackResponse() (response *SaveAfterAnswerDelayPlaybackResponse) {
	response = &SaveAfterAnswerDelayPlaybackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
