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

// DescribeLiveDomainSpecialConfig invokes the live.DescribeLiveDomainSpecialConfig API synchronously
func (client *Client) DescribeLiveDomainSpecialConfig(request *DescribeLiveDomainSpecialConfigRequest) (response *DescribeLiveDomainSpecialConfigResponse, err error) {
	response = CreateDescribeLiveDomainSpecialConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainSpecialConfigWithChan invokes the live.DescribeLiveDomainSpecialConfig API asynchronously
func (client *Client) DescribeLiveDomainSpecialConfigWithChan(request *DescribeLiveDomainSpecialConfigRequest) (<-chan *DescribeLiveDomainSpecialConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainSpecialConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainSpecialConfig(request)
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

// DescribeLiveDomainSpecialConfigWithCallback invokes the live.DescribeLiveDomainSpecialConfig API asynchronously
func (client *Client) DescribeLiveDomainSpecialConfigWithCallback(request *DescribeLiveDomainSpecialConfigRequest, callback func(response *DescribeLiveDomainSpecialConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainSpecialConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainSpecialConfig(request)
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

// DescribeLiveDomainSpecialConfigRequest is the request struct for api DescribeLiveDomainSpecialConfig
type DescribeLiveDomainSpecialConfigRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Name       string           `position:"Query" name:"Name"`
}

// DescribeLiveDomainSpecialConfigResponse is the response struct for api DescribeLiveDomainSpecialConfig
type DescribeLiveDomainSpecialConfigResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Configs   []Config `json:"Configs" xml:"Configs"`
}

// CreateDescribeLiveDomainSpecialConfigRequest creates a request to invoke DescribeLiveDomainSpecialConfig API
func CreateDescribeLiveDomainSpecialConfigRequest() (request *DescribeLiveDomainSpecialConfigRequest) {
	request = &DescribeLiveDomainSpecialConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainSpecialConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainSpecialConfigResponse creates a response to parse from DescribeLiveDomainSpecialConfig response
func CreateDescribeLiveDomainSpecialConfigResponse() (response *DescribeLiveDomainSpecialConfigResponse) {
	response = &DescribeLiveDomainSpecialConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
