package sas

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

// PublicCreateImageScanTask invokes the sas.PublicCreateImageScanTask API synchronously
func (client *Client) PublicCreateImageScanTask(request *PublicCreateImageScanTaskRequest) (response *PublicCreateImageScanTaskResponse, err error) {
	response = CreatePublicCreateImageScanTaskResponse()
	err = client.DoAction(request, response)
	return
}

// PublicCreateImageScanTaskWithChan invokes the sas.PublicCreateImageScanTask API asynchronously
func (client *Client) PublicCreateImageScanTaskWithChan(request *PublicCreateImageScanTaskRequest) (<-chan *PublicCreateImageScanTaskResponse, <-chan error) {
	responseChan := make(chan *PublicCreateImageScanTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublicCreateImageScanTask(request)
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

// PublicCreateImageScanTaskWithCallback invokes the sas.PublicCreateImageScanTask API asynchronously
func (client *Client) PublicCreateImageScanTaskWithCallback(request *PublicCreateImageScanTaskRequest, callback func(response *PublicCreateImageScanTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublicCreateImageScanTaskResponse
		var err error
		defer close(result)
		response, err = client.PublicCreateImageScanTask(request)
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

// PublicCreateImageScanTaskRequest is the request struct for api PublicCreateImageScanTask
type PublicCreateImageScanTaskRequest struct {
	*requests.RpcRequest
	SourceIp       string `position:"Query" name:"SourceIp"`
	Digests        string `position:"Query" name:"Digests"`
	RegistryTypes  string `position:"Query" name:"RegistryTypes"`
	RegionIds      string `position:"Query" name:"RegionIds"`
	Tags           string `position:"Query" name:"Tags"`
	RepoNamespaces string `position:"Query" name:"RepoNamespaces"`
	InstanceIds    string `position:"Query" name:"InstanceIds"`
	RepoIds        string `position:"Query" name:"RepoIds"`
	RepoNames      string `position:"Query" name:"RepoNames"`
}

// PublicCreateImageScanTaskResponse is the response struct for api PublicCreateImageScanTask
type PublicCreateImageScanTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreatePublicCreateImageScanTaskRequest creates a request to invoke PublicCreateImageScanTask API
func CreatePublicCreateImageScanTaskRequest() (request *PublicCreateImageScanTaskRequest) {
	request = &PublicCreateImageScanTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "PublicCreateImageScanTask", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePublicCreateImageScanTaskResponse creates a response to parse from PublicCreateImageScanTask response
func CreatePublicCreateImageScanTaskResponse() (response *PublicCreateImageScanTaskResponse) {
	response = &PublicCreateImageScanTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
