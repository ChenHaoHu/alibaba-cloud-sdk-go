package dataworks_public

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

// SetEntityTags invokes the dataworks_public.SetEntityTags API synchronously
func (client *Client) SetEntityTags(request *SetEntityTagsRequest) (response *SetEntityTagsResponse, err error) {
	response = CreateSetEntityTagsResponse()
	err = client.DoAction(request, response)
	return
}

// SetEntityTagsWithChan invokes the dataworks_public.SetEntityTags API asynchronously
func (client *Client) SetEntityTagsWithChan(request *SetEntityTagsRequest) (<-chan *SetEntityTagsResponse, <-chan error) {
	responseChan := make(chan *SetEntityTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetEntityTags(request)
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

// SetEntityTagsWithCallback invokes the dataworks_public.SetEntityTags API asynchronously
func (client *Client) SetEntityTagsWithCallback(request *SetEntityTagsRequest, callback func(response *SetEntityTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetEntityTagsResponse
		var err error
		defer close(result)
		response, err = client.SetEntityTags(request)
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

// SetEntityTagsRequest is the request struct for api SetEntityTags
type SetEntityTagsRequest struct {
	*requests.RpcRequest
	QualifiedName string               `position:"Query" name:"QualifiedName"`
	Tags          *[]SetEntityTagsTags `position:"Body" name:"Tags"  type:"Json"`
}

// SetEntityTagsTags is a repeated param struct in SetEntityTagsRequest
type SetEntityTagsTags struct {
	TagValue string `name:"TagValue"`
	TagKey   string `name:"TagKey"`
}

// SetEntityTagsResponse is the response struct for api SetEntityTags
type SetEntityTagsResponse struct {
	*responses.BaseResponse
	Data           bool   `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateSetEntityTagsRequest creates a request to invoke SetEntityTags API
func CreateSetEntityTagsRequest() (request *SetEntityTagsRequest) {
	request = &SetEntityTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "SetEntityTags", "", "")
	request.Method = requests.POST
	return
}

// CreateSetEntityTagsResponse creates a response to parse from SetEntityTags response
func CreateSetEntityTagsResponse() (response *SetEntityTagsResponse) {
	response = &SetEntityTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
