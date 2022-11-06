package iot

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

// PageQuerySharedSpeechOpen invokes the iot.PageQuerySharedSpeechOpen API synchronously
func (client *Client) PageQuerySharedSpeechOpen(request *PageQuerySharedSpeechOpenRequest) (response *PageQuerySharedSpeechOpenResponse, err error) {
	response = CreatePageQuerySharedSpeechOpenResponse()
	err = client.DoAction(request, response)
	return
}

// PageQuerySharedSpeechOpenWithChan invokes the iot.PageQuerySharedSpeechOpen API asynchronously
func (client *Client) PageQuerySharedSpeechOpenWithChan(request *PageQuerySharedSpeechOpenRequest) (<-chan *PageQuerySharedSpeechOpenResponse, <-chan error) {
	responseChan := make(chan *PageQuerySharedSpeechOpenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PageQuerySharedSpeechOpen(request)
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

// PageQuerySharedSpeechOpenWithCallback invokes the iot.PageQuerySharedSpeechOpen API asynchronously
func (client *Client) PageQuerySharedSpeechOpenWithCallback(request *PageQuerySharedSpeechOpenRequest, callback func(response *PageQuerySharedSpeechOpenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PageQuerySharedSpeechOpenResponse
		var err error
		defer close(result)
		response, err = client.PageQuerySharedSpeechOpen(request)
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

// PageQuerySharedSpeechOpenRequest is the request struct for api PageQuerySharedSpeechOpen
type PageQuerySharedSpeechOpenRequest struct {
	*requests.RpcRequest
	PageId        requests.Integer `position:"Body" name:"PageId"`
	IotId         string           `position:"Body" name:"IotId"`
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	ShareTaskCode string           `position:"Body" name:"ShareTaskCode"`
	ProductKey    string           `position:"Body" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Body" name:"DeviceName"`
	Status        requests.Integer `position:"Body" name:"Status"`
}

// PageQuerySharedSpeechOpenResponse is the response struct for api PageQuerySharedSpeechOpen
type PageQuerySharedSpeechOpenResponse struct {
	*responses.BaseResponse
	RequestId    string                          `json:"RequestId" xml:"RequestId"`
	Success      bool                            `json:"Success" xml:"Success"`
	Code         string                          `json:"Code" xml:"Code"`
	ErrorMessage string                          `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInPageQuerySharedSpeechOpen `json:"Data" xml:"Data"`
}

// CreatePageQuerySharedSpeechOpenRequest creates a request to invoke PageQuerySharedSpeechOpen API
func CreatePageQuerySharedSpeechOpenRequest() (request *PageQuerySharedSpeechOpenRequest) {
	request = &PageQuerySharedSpeechOpenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "PageQuerySharedSpeechOpen", "", "")
	request.Method = requests.POST
	return
}

// CreatePageQuerySharedSpeechOpenResponse creates a response to parse from PageQuerySharedSpeechOpen response
func CreatePageQuerySharedSpeechOpenResponse() (response *PageQuerySharedSpeechOpenResponse) {
	response = &PageQuerySharedSpeechOpenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
