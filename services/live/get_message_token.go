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

// GetMessageToken invokes the live.GetMessageToken API synchronously
func (client *Client) GetMessageToken(request *GetMessageTokenRequest) (response *GetMessageTokenResponse, err error) {
	response = CreateGetMessageTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetMessageTokenWithChan invokes the live.GetMessageToken API asynchronously
func (client *Client) GetMessageTokenWithChan(request *GetMessageTokenRequest) (<-chan *GetMessageTokenResponse, <-chan error) {
	responseChan := make(chan *GetMessageTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMessageToken(request)
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

// GetMessageTokenWithCallback invokes the live.GetMessageToken API asynchronously
func (client *Client) GetMessageTokenWithCallback(request *GetMessageTokenRequest, callback func(response *GetMessageTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMessageTokenResponse
		var err error
		defer close(result)
		response, err = client.GetMessageToken(request)
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

// GetMessageTokenRequest is the request struct for api GetMessageToken
type GetMessageTokenRequest struct {
	*requests.RpcRequest
	DeviceId   string `position:"Body" name:"DeviceId"`
	UserId     string `position:"Body" name:"UserId"`
	DeviceType string `position:"Body" name:"DeviceType"`
	AppId      string `position:"Body" name:"AppId"`
}

// GetMessageTokenResponse is the response struct for api GetMessageToken
type GetMessageTokenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateGetMessageTokenRequest creates a request to invoke GetMessageToken API
func CreateGetMessageTokenRequest() (request *GetMessageTokenRequest) {
	request = &GetMessageTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "GetMessageToken", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMessageTokenResponse creates a response to parse from GetMessageToken response
func CreateGetMessageTokenResponse() (response *GetMessageTokenResponse) {
	response = &GetMessageTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
