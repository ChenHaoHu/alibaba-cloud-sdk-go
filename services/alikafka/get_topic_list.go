package alikafka

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

// GetTopicList invokes the alikafka.GetTopicList API synchronously
func (client *Client) GetTopicList(request *GetTopicListRequest) (response *GetTopicListResponse, err error) {
	response = CreateGetTopicListResponse()
	err = client.DoAction(request, response)
	return
}

// GetTopicListWithChan invokes the alikafka.GetTopicList API asynchronously
func (client *Client) GetTopicListWithChan(request *GetTopicListRequest) (<-chan *GetTopicListResponse, <-chan error) {
	responseChan := make(chan *GetTopicListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTopicList(request)
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

// GetTopicListWithCallback invokes the alikafka.GetTopicList API asynchronously
func (client *Client) GetTopicListWithCallback(request *GetTopicListRequest, callback func(response *GetTopicListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTopicListResponse
		var err error
		defer close(result)
		response, err = client.GetTopicList(request)
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

// GetTopicListRequest is the request struct for api GetTopicList
type GetTopicListRequest struct {
	*requests.RpcRequest
	CurrentPage string `position:"Query" name:"CurrentPage"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	PageSize    string `position:"Query" name:"PageSize"`
	Topic       string `position:"Query" name:"Topic"`
}

// GetTopicListResponse is the response struct for api GetTopicList
type GetTopicListResponse struct {
	*responses.BaseResponse
	CurrentPage int                     `json:"CurrentPage" xml:"CurrentPage"`
	RequestId   string                  `json:"RequestId" xml:"RequestId"`
	Success     bool                    `json:"Success" xml:"Success"`
	Code        int                     `json:"Code" xml:"Code"`
	Message     string                  `json:"Message" xml:"Message"`
	PageSize    int                     `json:"PageSize" xml:"PageSize"`
	Total       int                     `json:"Total" xml:"Total"`
	TopicList   TopicListInGetTopicList `json:"TopicList" xml:"TopicList"`
}

// CreateGetTopicListRequest creates a request to invoke GetTopicList API
func CreateGetTopicListRequest() (request *GetTopicListRequest) {
	request = &GetTopicListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "GetTopicList", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTopicListResponse creates a response to parse from GetTopicList response
func CreateGetTopicListResponse() (response *GetTopicListResponse) {
	response = &GetTopicListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
