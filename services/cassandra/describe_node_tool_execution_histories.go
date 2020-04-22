package cassandra

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

// DescribeNodeToolExecutionHistories invokes the cassandra.DescribeNodeToolExecutionHistories API synchronously
// api document: https://help.aliyun.com/api/cassandra/describenodetoolexecutionhistories.html
func (client *Client) DescribeNodeToolExecutionHistories(request *DescribeNodeToolExecutionHistoriesRequest) (response *DescribeNodeToolExecutionHistoriesResponse, err error) {
	response = CreateDescribeNodeToolExecutionHistoriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNodeToolExecutionHistoriesWithChan invokes the cassandra.DescribeNodeToolExecutionHistories API asynchronously
// api document: https://help.aliyun.com/api/cassandra/describenodetoolexecutionhistories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNodeToolExecutionHistoriesWithChan(request *DescribeNodeToolExecutionHistoriesRequest) (<-chan *DescribeNodeToolExecutionHistoriesResponse, <-chan error) {
	responseChan := make(chan *DescribeNodeToolExecutionHistoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNodeToolExecutionHistories(request)
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

// DescribeNodeToolExecutionHistoriesWithCallback invokes the cassandra.DescribeNodeToolExecutionHistories API asynchronously
// api document: https://help.aliyun.com/api/cassandra/describenodetoolexecutionhistories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNodeToolExecutionHistoriesWithCallback(request *DescribeNodeToolExecutionHistoriesRequest, callback func(response *DescribeNodeToolExecutionHistoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNodeToolExecutionHistoriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeNodeToolExecutionHistories(request)
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

// DescribeNodeToolExecutionHistoriesRequest is the request struct for api DescribeNodeToolExecutionHistories
type DescribeNodeToolExecutionHistoriesRequest struct {
	*requests.RpcRequest
	ClusterId  string           `position:"Query" name:"ClusterId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeNodeToolExecutionHistoriesResponse is the response struct for api DescribeNodeToolExecutionHistories
type DescribeNodeToolExecutionHistoriesResponse struct {
	*responses.BaseResponse
	RequestId  string                                        `json:"RequestId" xml:"RequestId"`
	PageNumber int                                           `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                                           `json:"PageSize" xml:"PageSize"`
	TotalCount int64                                         `json:"TotalCount" xml:"TotalCount"`
	Histories  HistoriesInDescribeNodeToolExecutionHistories `json:"Histories" xml:"Histories"`
}

// CreateDescribeNodeToolExecutionHistoriesRequest creates a request to invoke DescribeNodeToolExecutionHistories API
func CreateDescribeNodeToolExecutionHistoriesRequest() (request *DescribeNodeToolExecutionHistoriesRequest) {
	request = &DescribeNodeToolExecutionHistoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "DescribeNodeToolExecutionHistories", "Cassandra", "openAPI")
	return
}

// CreateDescribeNodeToolExecutionHistoriesResponse creates a response to parse from DescribeNodeToolExecutionHistories response
func CreateDescribeNodeToolExecutionHistoriesResponse() (response *DescribeNodeToolExecutionHistoriesResponse) {
	response = &DescribeNodeToolExecutionHistoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
