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

// ListEntitiesByTags invokes the dataworks_public.ListEntitiesByTags API synchronously
func (client *Client) ListEntitiesByTags(request *ListEntitiesByTagsRequest) (response *ListEntitiesByTagsResponse, err error) {
	response = CreateListEntitiesByTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ListEntitiesByTagsWithChan invokes the dataworks_public.ListEntitiesByTags API asynchronously
func (client *Client) ListEntitiesByTagsWithChan(request *ListEntitiesByTagsRequest) (<-chan *ListEntitiesByTagsResponse, <-chan error) {
	responseChan := make(chan *ListEntitiesByTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEntitiesByTags(request)
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

// ListEntitiesByTagsWithCallback invokes the dataworks_public.ListEntitiesByTags API asynchronously
func (client *Client) ListEntitiesByTagsWithCallback(request *ListEntitiesByTagsRequest, callback func(response *ListEntitiesByTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEntitiesByTagsResponse
		var err error
		defer close(result)
		response, err = client.ListEntitiesByTags(request)
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

// ListEntitiesByTagsRequest is the request struct for api ListEntitiesByTags
type ListEntitiesByTagsRequest struct {
	*requests.RpcRequest
	EntityType string                    `position:"Query" name:"EntityType"`
	NextToken  string                    `position:"Query" name:"NextToken"`
	PageSize   requests.Integer          `position:"Query" name:"PageSize"`
	Tags       *[]ListEntitiesByTagsTags `position:"Query" name:"Tags"  type:"Json"`
}

// ListEntitiesByTagsTags is a repeated param struct in ListEntitiesByTagsRequest
type ListEntitiesByTagsTags struct {
	TagValue string `name:"TagValue"`
	TagKey   string `name:"TagKey"`
}

// ListEntitiesByTagsResponse is the response struct for api ListEntitiesByTags
type ListEntitiesByTagsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListEntitiesByTagsRequest creates a request to invoke ListEntitiesByTags API
func CreateListEntitiesByTagsRequest() (request *ListEntitiesByTagsRequest) {
	request = &ListEntitiesByTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListEntitiesByTags", "", "")
	request.Method = requests.GET
	return
}

// CreateListEntitiesByTagsResponse creates a response to parse from ListEntitiesByTags response
func CreateListEntitiesByTagsResponse() (response *ListEntitiesByTagsResponse) {
	response = &ListEntitiesByTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
