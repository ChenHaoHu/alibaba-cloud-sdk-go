package das

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

// GetSqlOptimizeAdvice invokes the das.GetSqlOptimizeAdvice API synchronously
func (client *Client) GetSqlOptimizeAdvice(request *GetSqlOptimizeAdviceRequest) (response *GetSqlOptimizeAdviceResponse, err error) {
	response = CreateGetSqlOptimizeAdviceResponse()
	err = client.DoAction(request, response)
	return
}

// GetSqlOptimizeAdviceWithChan invokes the das.GetSqlOptimizeAdvice API asynchronously
func (client *Client) GetSqlOptimizeAdviceWithChan(request *GetSqlOptimizeAdviceRequest) (<-chan *GetSqlOptimizeAdviceResponse, <-chan error) {
	responseChan := make(chan *GetSqlOptimizeAdviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSqlOptimizeAdvice(request)
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

// GetSqlOptimizeAdviceWithCallback invokes the das.GetSqlOptimizeAdvice API asynchronously
func (client *Client) GetSqlOptimizeAdviceWithCallback(request *GetSqlOptimizeAdviceRequest, callback func(response *GetSqlOptimizeAdviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSqlOptimizeAdviceResponse
		var err error
		defer close(result)
		response, err = client.GetSqlOptimizeAdvice(request)
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

// GetSqlOptimizeAdviceRequest is the request struct for api GetSqlOptimizeAdvice
type GetSqlOptimizeAdviceRequest struct {
	*requests.RpcRequest
	EndDt          string `position:"Query" name:"EndDt"`
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	Engine         string `position:"Query" name:"Engine"`
	InstanceIds    string `position:"Query" name:"InstanceIds"`
	StartDt        string `position:"Query" name:"StartDt"`
}

// GetSqlOptimizeAdviceResponse is the response struct for api GetSqlOptimizeAdvice
type GetSqlOptimizeAdviceResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateGetSqlOptimizeAdviceRequest creates a request to invoke GetSqlOptimizeAdvice API
func CreateGetSqlOptimizeAdviceRequest() (request *GetSqlOptimizeAdviceRequest) {
	request = &GetSqlOptimizeAdviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetSqlOptimizeAdvice", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSqlOptimizeAdviceResponse creates a response to parse from GetSqlOptimizeAdvice response
func CreateGetSqlOptimizeAdviceResponse() (response *GetSqlOptimizeAdviceResponse) {
	response = &GetSqlOptimizeAdviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
