package live

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

// CancelMuteAllGroupUser invokes the live.CancelMuteAllGroupUser API synchronously
func (client *Client) CancelMuteAllGroupUser(request *CancelMuteAllGroupUserRequest) (response *CancelMuteAllGroupUserResponse, err error) {
	response = CreateCancelMuteAllGroupUserResponse()
	err = client.DoAction(request, response)
	return
}

// CancelMuteAllGroupUserWithChan invokes the live.CancelMuteAllGroupUser API asynchronously
func (client *Client) CancelMuteAllGroupUserWithChan(request *CancelMuteAllGroupUserRequest) (<-chan *CancelMuteAllGroupUserResponse, <-chan error) {
	responseChan := make(chan *CancelMuteAllGroupUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelMuteAllGroupUser(request)
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

// CancelMuteAllGroupUserWithCallback invokes the live.CancelMuteAllGroupUser API asynchronously
func (client *Client) CancelMuteAllGroupUserWithCallback(request *CancelMuteAllGroupUserRequest, callback func(response *CancelMuteAllGroupUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelMuteAllGroupUserResponse
		var err error
		defer close(result)
		response, err = client.CancelMuteAllGroupUser(request)
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

// CancelMuteAllGroupUserRequest is the request struct for api CancelMuteAllGroupUser
type CancelMuteAllGroupUserRequest struct {
	*requests.RpcRequest
	GroupId        string `position:"Body" name:"GroupId"`
	AppId          string `position:"Body" name:"AppId"`
	OperatorUserId string `position:"Body" name:"OperatorUserId"`
}

// CancelMuteAllGroupUserResponse is the response struct for api CancelMuteAllGroupUser
type CancelMuteAllGroupUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCancelMuteAllGroupUserRequest creates a request to invoke CancelMuteAllGroupUser API
func CreateCancelMuteAllGroupUserRequest() (request *CancelMuteAllGroupUserRequest) {
	request = &CancelMuteAllGroupUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CancelMuteAllGroupUser", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelMuteAllGroupUserResponse creates a response to parse from CancelMuteAllGroupUser response
func CreateCancelMuteAllGroupUserResponse() (response *CancelMuteAllGroupUserResponse) {
	response = &CancelMuteAllGroupUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
