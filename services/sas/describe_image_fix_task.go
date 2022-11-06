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

// DescribeImageFixTask invokes the sas.DescribeImageFixTask API synchronously
func (client *Client) DescribeImageFixTask(request *DescribeImageFixTaskRequest) (response *DescribeImageFixTaskResponse, err error) {
	response = CreateDescribeImageFixTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageFixTaskWithChan invokes the sas.DescribeImageFixTask API asynchronously
func (client *Client) DescribeImageFixTaskWithChan(request *DescribeImageFixTaskRequest) (<-chan *DescribeImageFixTaskResponse, <-chan error) {
	responseChan := make(chan *DescribeImageFixTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageFixTask(request)
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

// DescribeImageFixTaskWithCallback invokes the sas.DescribeImageFixTask API asynchronously
func (client *Client) DescribeImageFixTaskWithCallback(request *DescribeImageFixTaskRequest, callback func(response *DescribeImageFixTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageFixTaskResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageFixTask(request)
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

// DescribeImageFixTaskRequest is the request struct for api DescribeImageFixTask
type DescribeImageFixTaskRequest struct {
	*requests.RpcRequest
	StartTime   requests.Integer `position:"Query" name:"StartTime"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Status      string           `position:"Query" name:"Status"`
}

// DescribeImageFixTaskResponse is the response struct for api DescribeImageFixTask
type DescribeImageFixTaskResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	PageInfo   PageInfo    `json:"PageInfo" xml:"PageInfo"`
	BuildTasks []BuildTask `json:"BuildTasks" xml:"BuildTasks"`
}

// CreateDescribeImageFixTaskRequest creates a request to invoke DescribeImageFixTask API
func CreateDescribeImageFixTaskRequest() (request *DescribeImageFixTaskRequest) {
	request = &DescribeImageFixTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeImageFixTask", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeImageFixTaskResponse creates a response to parse from DescribeImageFixTask response
func CreateDescribeImageFixTaskResponse() (response *DescribeImageFixTaskResponse) {
	response = &DescribeImageFixTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
