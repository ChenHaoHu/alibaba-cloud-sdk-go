package oos

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

// ListStateConfigurations invokes the oos.ListStateConfigurations API synchronously
func (client *Client) ListStateConfigurations(request *ListStateConfigurationsRequest) (response *ListStateConfigurationsResponse, err error) {
	response = CreateListStateConfigurationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListStateConfigurationsWithChan invokes the oos.ListStateConfigurations API asynchronously
func (client *Client) ListStateConfigurationsWithChan(request *ListStateConfigurationsRequest) (<-chan *ListStateConfigurationsResponse, <-chan error) {
	responseChan := make(chan *ListStateConfigurationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListStateConfigurations(request)
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

// ListStateConfigurationsWithCallback invokes the oos.ListStateConfigurations API asynchronously
func (client *Client) ListStateConfigurationsWithCallback(request *ListStateConfigurationsRequest, callback func(response *ListStateConfigurationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListStateConfigurationsResponse
		var err error
		defer close(result)
		response, err = client.ListStateConfigurations(request)
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

// ListStateConfigurationsRequest is the request struct for api ListStateConfigurations
type ListStateConfigurationsRequest struct {
	*requests.RpcRequest
	StateConfigurationIds string           `position:"Query" name:"StateConfigurationIds"`
	Tags                  string           `position:"Query" name:"Tags"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	TemplateVersion       string           `position:"Query" name:"TemplateVersion"`
	NextToken             string           `position:"Query" name:"NextToken"`
	MaxResults            requests.Integer `position:"Query" name:"MaxResults"`
	TemplateName          string           `position:"Query" name:"TemplateName"`
}

// ListStateConfigurationsResponse is the response struct for api ListStateConfigurations
type ListStateConfigurationsResponse struct {
	*responses.BaseResponse
	RequestId           string                                        `json:"RequestId" xml:"RequestId"`
	StateConfigurations []StateConfigurationInListStateConfigurations `json:"StateConfigurations" xml:"StateConfigurations"`
}

// CreateListStateConfigurationsRequest creates a request to invoke ListStateConfigurations API
func CreateListStateConfigurationsRequest() (request *ListStateConfigurationsRequest) {
	request = &ListStateConfigurationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ListStateConfigurations", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListStateConfigurationsResponse creates a response to parse from ListStateConfigurations response
func CreateListStateConfigurationsResponse() (response *ListStateConfigurationsResponse) {
	response = &ListStateConfigurationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
