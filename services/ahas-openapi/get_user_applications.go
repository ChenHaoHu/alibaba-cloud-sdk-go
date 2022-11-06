package ahas_openapi

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

// GetUserApplications invokes the ahas_openapi.GetUserApplications API synchronously
func (client *Client) GetUserApplications(request *GetUserApplicationsRequest) (response *GetUserApplicationsResponse, err error) {
	response = CreateGetUserApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserApplicationsWithChan invokes the ahas_openapi.GetUserApplications API asynchronously
func (client *Client) GetUserApplicationsWithChan(request *GetUserApplicationsRequest) (<-chan *GetUserApplicationsResponse, <-chan error) {
	responseChan := make(chan *GetUserApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserApplications(request)
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

// GetUserApplicationsWithCallback invokes the ahas_openapi.GetUserApplications API asynchronously
func (client *Client) GetUserApplicationsWithCallback(request *GetUserApplicationsRequest, callback func(response *GetUserApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserApplicationsResponse
		var err error
		defer close(result)
		response, err = client.GetUserApplications(request)
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

// GetUserApplicationsRequest is the request struct for api GetUserApplications
type GetUserApplicationsRequest struct {
	*requests.RpcRequest
	AhasRegionId string `position:"Query" name:"AhasRegionId"`
	Namespace    string `position:"Query" name:"Namespace"`
}

// GetUserApplicationsResponse is the response struct for api GetUserApplications
type GetUserApplicationsResponse struct {
	*responses.BaseResponse
	Message           string                  `json:"Message" xml:"Message"`
	RequestId         string                  `json:"RequestId" xml:"RequestId"`
	HttpStatusCode    int                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code              string                  `json:"Code" xml:"Code"`
	Success           bool                    `json:"Success" xml:"Success"`
	AppNameAndIdPairs []AppNameAndIdPairsItem `json:"AppNameAndIdPairs" xml:"AppNameAndIdPairs"`
}

// CreateGetUserApplicationsRequest creates a request to invoke GetUserApplications API
func CreateGetUserApplicationsRequest() (request *GetUserApplicationsRequest) {
	request = &GetUserApplicationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "GetUserApplications", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserApplicationsResponse creates a response to parse from GetUserApplications response
func CreateGetUserApplicationsResponse() (response *GetUserApplicationsResponse) {
	response = &GetUserApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
