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

// GetMultiRateConfigList invokes the live.GetMultiRateConfigList API synchronously
func (client *Client) GetMultiRateConfigList(request *GetMultiRateConfigListRequest) (response *GetMultiRateConfigListResponse, err error) {
	response = CreateGetMultiRateConfigListResponse()
	err = client.DoAction(request, response)
	return
}

// GetMultiRateConfigListWithChan invokes the live.GetMultiRateConfigList API asynchronously
func (client *Client) GetMultiRateConfigListWithChan(request *GetMultiRateConfigListRequest) (<-chan *GetMultiRateConfigListResponse, <-chan error) {
	responseChan := make(chan *GetMultiRateConfigListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMultiRateConfigList(request)
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

// GetMultiRateConfigListWithCallback invokes the live.GetMultiRateConfigList API asynchronously
func (client *Client) GetMultiRateConfigListWithCallback(request *GetMultiRateConfigListRequest, callback func(response *GetMultiRateConfigListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMultiRateConfigListResponse
		var err error
		defer close(result)
		response, err = client.GetMultiRateConfigList(request)
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

// GetMultiRateConfigListRequest is the request struct for api GetMultiRateConfigList
type GetMultiRateConfigListRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// GetMultiRateConfigListResponse is the response struct for api GetMultiRateConfigList
type GetMultiRateConfigListResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Message   string    `json:"Message" xml:"Message"`
	Code      int       `json:"Code" xml:"Code"`
	GroupInfo GroupInfo `json:"GroupInfo" xml:"GroupInfo"`
}

// CreateGetMultiRateConfigListRequest creates a request to invoke GetMultiRateConfigList API
func CreateGetMultiRateConfigListRequest() (request *GetMultiRateConfigListRequest) {
	request = &GetMultiRateConfigListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "GetMultiRateConfigList", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMultiRateConfigListResponse creates a response to parse from GetMultiRateConfigList response
func CreateGetMultiRateConfigListResponse() (response *GetMultiRateConfigListResponse) {
	response = &GetMultiRateConfigListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
