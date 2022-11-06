package dms_enterprise

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

// ResumeTaskFlowInstance invokes the dms_enterprise.ResumeTaskFlowInstance API synchronously
func (client *Client) ResumeTaskFlowInstance(request *ResumeTaskFlowInstanceRequest) (response *ResumeTaskFlowInstanceResponse, err error) {
	response = CreateResumeTaskFlowInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeTaskFlowInstanceWithChan invokes the dms_enterprise.ResumeTaskFlowInstance API asynchronously
func (client *Client) ResumeTaskFlowInstanceWithChan(request *ResumeTaskFlowInstanceRequest) (<-chan *ResumeTaskFlowInstanceResponse, <-chan error) {
	responseChan := make(chan *ResumeTaskFlowInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeTaskFlowInstance(request)
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

// ResumeTaskFlowInstanceWithCallback invokes the dms_enterprise.ResumeTaskFlowInstance API asynchronously
func (client *Client) ResumeTaskFlowInstanceWithCallback(request *ResumeTaskFlowInstanceRequest, callback func(response *ResumeTaskFlowInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeTaskFlowInstanceResponse
		var err error
		defer close(result)
		response, err = client.ResumeTaskFlowInstance(request)
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

// ResumeTaskFlowInstanceRequest is the request struct for api ResumeTaskFlowInstance
type ResumeTaskFlowInstanceRequest struct {
	*requests.RpcRequest
	DagVersion    string           `position:"Query" name:"DagVersion"`
	DagId         requests.Integer `position:"Query" name:"DagId"`
	Tid           requests.Integer `position:"Query" name:"Tid"`
	DagInstanceId requests.Integer `position:"Query" name:"DagInstanceId"`
}

// ResumeTaskFlowInstanceResponse is the response struct for api ResumeTaskFlowInstance
type ResumeTaskFlowInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateResumeTaskFlowInstanceRequest creates a request to invoke ResumeTaskFlowInstance API
func CreateResumeTaskFlowInstanceRequest() (request *ResumeTaskFlowInstanceRequest) {
	request = &ResumeTaskFlowInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ResumeTaskFlowInstance", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResumeTaskFlowInstanceResponse creates a response to parse from ResumeTaskFlowInstance response
func CreateResumeTaskFlowInstanceResponse() (response *ResumeTaskFlowInstanceResponse) {
	response = &ResumeTaskFlowInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
