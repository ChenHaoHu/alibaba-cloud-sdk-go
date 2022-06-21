package live

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

// DescribeRoomStatus invokes the live.DescribeRoomStatus API synchronously
func (client *Client) DescribeRoomStatus(request *DescribeRoomStatusRequest) (response *DescribeRoomStatusResponse, err error) {
	response = CreateDescribeRoomStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRoomStatusWithChan invokes the live.DescribeRoomStatus API asynchronously
func (client *Client) DescribeRoomStatusWithChan(request *DescribeRoomStatusRequest) (<-chan *DescribeRoomStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeRoomStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRoomStatus(request)
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

// DescribeRoomStatusWithCallback invokes the live.DescribeRoomStatus API asynchronously
func (client *Client) DescribeRoomStatusWithCallback(request *DescribeRoomStatusRequest, callback func(response *DescribeRoomStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRoomStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeRoomStatus(request)
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

// DescribeRoomStatusRequest is the request struct for api DescribeRoomStatus
type DescribeRoomStatusRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	RoomId  string           `position:"Query" name:"RoomId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// DescribeRoomStatusResponse is the response struct for api DescribeRoomStatus
type DescribeRoomStatusResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	RoomStatus int    `json:"RoomStatus" xml:"RoomStatus"`
}

// CreateDescribeRoomStatusRequest creates a request to invoke DescribeRoomStatus API
func CreateDescribeRoomStatusRequest() (request *DescribeRoomStatusRequest) {
	request = &DescribeRoomStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeRoomStatus", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRoomStatusResponse creates a response to parse from DescribeRoomStatus response
func CreateDescribeRoomStatusResponse() (response *DescribeRoomStatusResponse) {
	response = &DescribeRoomStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
