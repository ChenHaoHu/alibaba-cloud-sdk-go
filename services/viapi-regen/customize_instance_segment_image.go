package viapi_regen

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

// CustomizeInstanceSegmentImage invokes the viapi_regen.CustomizeInstanceSegmentImage API synchronously
func (client *Client) CustomizeInstanceSegmentImage(request *CustomizeInstanceSegmentImageRequest) (response *CustomizeInstanceSegmentImageResponse, err error) {
	response = CreateCustomizeInstanceSegmentImageResponse()
	err = client.DoAction(request, response)
	return
}

// CustomizeInstanceSegmentImageWithChan invokes the viapi_regen.CustomizeInstanceSegmentImage API asynchronously
func (client *Client) CustomizeInstanceSegmentImageWithChan(request *CustomizeInstanceSegmentImageRequest) (<-chan *CustomizeInstanceSegmentImageResponse, <-chan error) {
	responseChan := make(chan *CustomizeInstanceSegmentImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CustomizeInstanceSegmentImage(request)
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

// CustomizeInstanceSegmentImageWithCallback invokes the viapi_regen.CustomizeInstanceSegmentImage API asynchronously
func (client *Client) CustomizeInstanceSegmentImageWithCallback(request *CustomizeInstanceSegmentImageRequest, callback func(response *CustomizeInstanceSegmentImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CustomizeInstanceSegmentImageResponse
		var err error
		defer close(result)
		response, err = client.CustomizeInstanceSegmentImage(request)
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

// CustomizeInstanceSegmentImageRequest is the request struct for api CustomizeInstanceSegmentImage
type CustomizeInstanceSegmentImageRequest struct {
	*requests.RpcRequest
	ImageUrl  string `position:"Body" name:"ImageUrl"`
	ServiceId string `position:"Body" name:"ServiceId"`
}

// CustomizeInstanceSegmentImageResponse is the response struct for api CustomizeInstanceSegmentImage
type CustomizeInstanceSegmentImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCustomizeInstanceSegmentImageRequest creates a request to invoke CustomizeInstanceSegmentImage API
func CreateCustomizeInstanceSegmentImageRequest() (request *CustomizeInstanceSegmentImageRequest) {
	request = &CustomizeInstanceSegmentImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "CustomizeInstanceSegmentImage", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCustomizeInstanceSegmentImageResponse creates a response to parse from CustomizeInstanceSegmentImage response
func CreateCustomizeInstanceSegmentImageResponse() (response *CustomizeInstanceSegmentImageResponse) {
	response = &CustomizeInstanceSegmentImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
