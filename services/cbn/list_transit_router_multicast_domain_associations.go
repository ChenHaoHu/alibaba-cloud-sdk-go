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

// ListTransitRouterMulticastDomainAssociations invokes the cbn.ListTransitRouterMulticastDomainAssociations API synchronously
func (client *Client) ListTransitRouterMulticastDomainAssociations(request *ListTransitRouterMulticastDomainAssociationsRequest) (response *ListTransitRouterMulticastDomainAssociationsResponse, err error) {
	response = CreateListTransitRouterMulticastDomainAssociationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransitRouterMulticastDomainAssociationsWithChan invokes the cbn.ListTransitRouterMulticastDomainAssociations API asynchronously
func (client *Client) ListTransitRouterMulticastDomainAssociationsWithChan(request *ListTransitRouterMulticastDomainAssociationsRequest) (<-chan *ListTransitRouterMulticastDomainAssociationsResponse, <-chan error) {
	responseChan := make(chan *ListTransitRouterMulticastDomainAssociationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransitRouterMulticastDomainAssociations(request)
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

// ListTransitRouterMulticastDomainAssociationsWithCallback invokes the cbn.ListTransitRouterMulticastDomainAssociations API asynchronously
func (client *Client) ListTransitRouterMulticastDomainAssociationsWithCallback(request *ListTransitRouterMulticastDomainAssociationsRequest, callback func(response *ListTransitRouterMulticastDomainAssociationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransitRouterMulticastDomainAssociationsResponse
		var err error
		defer close(result)
		response, err = client.ListTransitRouterMulticastDomainAssociations(request)
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

// ListTransitRouterMulticastDomainAssociationsRequest is the request struct for api ListTransitRouterMulticastDomainAssociations
type ListTransitRouterMulticastDomainAssociationsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                    string           `position:"Query" name:"ClientToken"`
	VSwitchIds                     *[]string        `position:"Query" name:"VSwitchIds"  type:"Repeated"`
	TransitRouterMulticastDomainId string           `position:"Query" name:"TransitRouterMulticastDomainId"`
	NextToken                      string           `position:"Query" name:"NextToken"`
	ResourceId                     string           `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount           string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                   string           `position:"Query" name:"OwnerAccount"`
	OwnerId                        requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType                   string           `position:"Query" name:"ResourceType"`
	TransitRouterAttachmentId      string           `position:"Query" name:"TransitRouterAttachmentId"`
	MaxResults                     requests.Integer `position:"Query" name:"MaxResults"`
}

// ListTransitRouterMulticastDomainAssociationsResponse is the response struct for api ListTransitRouterMulticastDomainAssociations
type ListTransitRouterMulticastDomainAssociationsResponse struct {
	*responses.BaseResponse
	RequestId                          string                              `json:"RequestId" xml:"RequestId"`
	TotalCount                         int                                 `json:"TotalCount" xml:"TotalCount"`
	MaxResults                         int                                 `json:"MaxResults" xml:"MaxResults"`
	NextToken                          string                              `json:"NextToken" xml:"NextToken"`
	TransitRouterMulticastAssociations []TransitRouterMulticastAssociation `json:"TransitRouterMulticastAssociations" xml:"TransitRouterMulticastAssociations"`
}

// CreateListTransitRouterMulticastDomainAssociationsRequest creates a request to invoke ListTransitRouterMulticastDomainAssociations API
func CreateListTransitRouterMulticastDomainAssociationsRequest() (request *ListTransitRouterMulticastDomainAssociationsRequest) {
	request = &ListTransitRouterMulticastDomainAssociationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListTransitRouterMulticastDomainAssociations", "", "")
	request.Method = requests.POST
	return
}

// CreateListTransitRouterMulticastDomainAssociationsResponse creates a response to parse from ListTransitRouterMulticastDomainAssociations response
func CreateListTransitRouterMulticastDomainAssociationsResponse() (response *ListTransitRouterMulticastDomainAssociationsResponse) {
	response = &ListTransitRouterMulticastDomainAssociationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
