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

// QuerySnapshotCallbackAuth invokes the live.QuerySnapshotCallbackAuth API synchronously
func (client *Client) QuerySnapshotCallbackAuth(request *QuerySnapshotCallbackAuthRequest) (response *QuerySnapshotCallbackAuthResponse, err error) {
	response = CreateQuerySnapshotCallbackAuthResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySnapshotCallbackAuthWithChan invokes the live.QuerySnapshotCallbackAuth API asynchronously
func (client *Client) QuerySnapshotCallbackAuthWithChan(request *QuerySnapshotCallbackAuthRequest) (<-chan *QuerySnapshotCallbackAuthResponse, <-chan error) {
	responseChan := make(chan *QuerySnapshotCallbackAuthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySnapshotCallbackAuth(request)
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

// QuerySnapshotCallbackAuthWithCallback invokes the live.QuerySnapshotCallbackAuth API asynchronously
func (client *Client) QuerySnapshotCallbackAuthWithCallback(request *QuerySnapshotCallbackAuthRequest, callback func(response *QuerySnapshotCallbackAuthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySnapshotCallbackAuthResponse
		var err error
		defer close(result)
		response, err = client.QuerySnapshotCallbackAuth(request)
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

// QuerySnapshotCallbackAuthRequest is the request struct for api QuerySnapshotCallbackAuth
type QuerySnapshotCallbackAuthRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// QuerySnapshotCallbackAuthResponse is the response struct for api QuerySnapshotCallbackAuth
type QuerySnapshotCallbackAuthResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	DomainName      string `json:"DomainName" xml:"DomainName"`
	CallbackReqAuth string `json:"CallbackReqAuth" xml:"CallbackReqAuth"`
	CallbackAuthKey string `json:"CallbackAuthKey" xml:"CallbackAuthKey"`
}

// CreateQuerySnapshotCallbackAuthRequest creates a request to invoke QuerySnapshotCallbackAuth API
func CreateQuerySnapshotCallbackAuthRequest() (request *QuerySnapshotCallbackAuthRequest) {
	request = &QuerySnapshotCallbackAuthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "QuerySnapshotCallbackAuth", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySnapshotCallbackAuthResponse creates a response to parse from QuerySnapshotCallbackAuth response
func CreateQuerySnapshotCallbackAuthResponse() (response *QuerySnapshotCallbackAuthResponse) {
	response = &QuerySnapshotCallbackAuthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
