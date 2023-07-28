package dts

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

// StartMigrationJob invokes the dts.StartMigrationJob API synchronously
func (client *Client) StartMigrationJob(request *StartMigrationJobRequest) (response *StartMigrationJobResponse, err error) {
	response = CreateStartMigrationJobResponse()
	err = client.DoAction(request, response)
	return
}

// StartMigrationJobWithChan invokes the dts.StartMigrationJob API asynchronously
func (client *Client) StartMigrationJobWithChan(request *StartMigrationJobRequest) (<-chan *StartMigrationJobResponse, <-chan error) {
	responseChan := make(chan *StartMigrationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartMigrationJob(request)
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

// StartMigrationJobWithCallback invokes the dts.StartMigrationJob API asynchronously
func (client *Client) StartMigrationJobWithCallback(request *StartMigrationJobRequest, callback func(response *StartMigrationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartMigrationJobResponse
		var err error
		defer close(result)
		response, err = client.StartMigrationJob(request)
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

// StartMigrationJobRequest is the request struct for api StartMigrationJob
type StartMigrationJobRequest struct {
	*requests.RpcRequest
	MigrationJobId string `position:"Query" name:"MigrationJobId"`
}

// StartMigrationJobResponse is the response struct for api StartMigrationJob
type StartMigrationJobResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	Success    string `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateStartMigrationJobRequest creates a request to invoke StartMigrationJob API
func CreateStartMigrationJobRequest() (request *StartMigrationJobRequest) {
	request = &StartMigrationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2017-06-01", "StartMigrationJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartMigrationJobResponse creates a response to parse from StartMigrationJob response
func CreateStartMigrationJobResponse() (response *StartMigrationJobResponse) {
	response = &StartMigrationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
