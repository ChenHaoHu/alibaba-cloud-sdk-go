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

// ListLiveDelayConfig invokes the live.ListLiveDelayConfig API synchronously
func (client *Client) ListLiveDelayConfig(request *ListLiveDelayConfigRequest) (response *ListLiveDelayConfigResponse, err error) {
	response = CreateListLiveDelayConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListLiveDelayConfigWithChan invokes the live.ListLiveDelayConfig API asynchronously
func (client *Client) ListLiveDelayConfigWithChan(request *ListLiveDelayConfigRequest) (<-chan *ListLiveDelayConfigResponse, <-chan error) {
	responseChan := make(chan *ListLiveDelayConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLiveDelayConfig(request)
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

// ListLiveDelayConfigWithCallback invokes the live.ListLiveDelayConfig API asynchronously
func (client *Client) ListLiveDelayConfigWithCallback(request *ListLiveDelayConfigRequest, callback func(response *ListLiveDelayConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLiveDelayConfigResponse
		var err error
		defer close(result)
		response, err = client.ListLiveDelayConfig(request)
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

// ListLiveDelayConfigRequest is the request struct for api ListLiveDelayConfig
type ListLiveDelayConfigRequest struct {
	*requests.RpcRequest
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	Domain   string           `position:"Query" name:"Domain"`
}

// ListLiveDelayConfigResponse is the response struct for api ListLiveDelayConfig
type ListLiveDelayConfigResponse struct {
	*responses.BaseResponse
	Total           int             `json:"Total" xml:"Total"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	DelayConfigList DelayConfigList `json:"DelayConfigList" xml:"DelayConfigList"`
}

// CreateListLiveDelayConfigRequest creates a request to invoke ListLiveDelayConfig API
func CreateListLiveDelayConfigRequest() (request *ListLiveDelayConfigRequest) {
	request = &ListLiveDelayConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ListLiveDelayConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLiveDelayConfigResponse creates a response to parse from ListLiveDelayConfig response
func CreateListLiveDelayConfigResponse() (response *ListLiveDelayConfigResponse) {
	response = &ListLiveDelayConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
