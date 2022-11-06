package rds

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

// DescribeSecrets invokes the rds.DescribeSecrets API synchronously
func (client *Client) DescribeSecrets(request *DescribeSecretsRequest) (response *DescribeSecretsResponse, err error) {
	response = CreateDescribeSecretsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecretsWithChan invokes the rds.DescribeSecrets API asynchronously
func (client *Client) DescribeSecretsWithChan(request *DescribeSecretsRequest) (<-chan *DescribeSecretsResponse, <-chan error) {
	responseChan := make(chan *DescribeSecretsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecrets(request)
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

// DescribeSecretsWithCallback invokes the rds.DescribeSecrets API asynchronously
func (client *Client) DescribeSecretsWithCallback(request *DescribeSecretsRequest, callback func(response *DescribeSecretsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecretsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecrets(request)
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

// DescribeSecretsRequest is the request struct for api DescribeSecrets
type DescribeSecretsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Engine               string           `position:"Query" name:"Engine"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AcceptLanguage       string           `position:"Query" name:"AcceptLanguage"`
}

// DescribeSecretsResponse is the response struct for api DescribeSecrets
type DescribeSecretsResponse struct {
	*responses.BaseResponse
	RequestId  string        `json:"RequestId" xml:"RequestId"`
	PageSize   int64         `json:"PageSize" xml:"PageSize"`
	PageNumber int64         `json:"PageNumber" xml:"PageNumber"`
	Secrets    []SecretsItem `json:"Secrets" xml:"Secrets"`
}

// CreateDescribeSecretsRequest creates a request to invoke DescribeSecrets API
func CreateDescribeSecretsRequest() (request *DescribeSecretsRequest) {
	request = &DescribeSecretsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSecrets", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSecretsResponse creates a response to parse from DescribeSecrets response
func CreateDescribeSecretsResponse() (response *DescribeSecretsResponse) {
	response = &DescribeSecretsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
