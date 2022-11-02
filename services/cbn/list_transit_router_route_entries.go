package cbn

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

// ListTransitRouterRouteEntries invokes the cbn.ListTransitRouterRouteEntries API synchronously
func (client *Client) ListTransitRouterRouteEntries(request *ListTransitRouterRouteEntriesRequest) (response *ListTransitRouterRouteEntriesResponse, err error) {
	response = CreateListTransitRouterRouteEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransitRouterRouteEntriesWithChan invokes the cbn.ListTransitRouterRouteEntries API asynchronously
func (client *Client) ListTransitRouterRouteEntriesWithChan(request *ListTransitRouterRouteEntriesRequest) (<-chan *ListTransitRouterRouteEntriesResponse, <-chan error) {
	responseChan := make(chan *ListTransitRouterRouteEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransitRouterRouteEntries(request)
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

// ListTransitRouterRouteEntriesWithCallback invokes the cbn.ListTransitRouterRouteEntries API asynchronously
func (client *Client) ListTransitRouterRouteEntriesWithCallback(request *ListTransitRouterRouteEntriesRequest, callback func(response *ListTransitRouterRouteEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransitRouterRouteEntriesResponse
		var err error
		defer close(result)
		response, err = client.ListTransitRouterRouteEntries(request)
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

// ListTransitRouterRouteEntriesRequest is the request struct for api ListTransitRouterRouteEntries
type ListTransitRouterRouteEntriesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TransitRouterRouteEntryDestinationCidrBlock string           `position:"Query" name:"TransitRouterRouteEntryDestinationCidrBlock"`
	TransitRouterRouteTableId                   string           `position:"Query" name:"TransitRouterRouteTableId"`
	NextToken                                   string           `position:"Query" name:"NextToken"`
	TransitRouterRouteEntryStatus               string           `position:"Query" name:"TransitRouterRouteEntryStatus"`
	ResourceOwnerAccount                        string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                                string           `position:"Query" name:"OwnerAccount"`
	TransitRouterRouteEntryNames                *[]string        `position:"Query" name:"TransitRouterRouteEntryNames"  type:"Repeated"`
	TransitRouterRouteEntryIds                  *[]string        `position:"Query" name:"TransitRouterRouteEntryIds"  type:"Repeated"`
	OwnerId                                     requests.Integer `position:"Query" name:"OwnerId"`
	MaxResults                                  requests.Integer `position:"Query" name:"MaxResults"`
}

// ListTransitRouterRouteEntriesResponse is the response struct for api ListTransitRouterRouteEntries
type ListTransitRouterRouteEntriesResponse struct {
	*responses.BaseResponse
	NextToken                 string                    `json:"NextToken" xml:"NextToken"`
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	TotalCount                int                       `json:"TotalCount" xml:"TotalCount"`
	MaxResults                int                       `json:"MaxResults" xml:"MaxResults"`
	TransitRouterRouteEntries []TransitRouterRouteEntry `json:"TransitRouterRouteEntries" xml:"TransitRouterRouteEntries"`
}

// CreateListTransitRouterRouteEntriesRequest creates a request to invoke ListTransitRouterRouteEntries API
func CreateListTransitRouterRouteEntriesRequest() (request *ListTransitRouterRouteEntriesRequest) {
	request = &ListTransitRouterRouteEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListTransitRouterRouteEntries", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTransitRouterRouteEntriesResponse creates a response to parse from ListTransitRouterRouteEntries response
func CreateListTransitRouterRouteEntriesResponse() (response *ListTransitRouterRouteEntriesResponse) {
	response = &ListTransitRouterRouteEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
