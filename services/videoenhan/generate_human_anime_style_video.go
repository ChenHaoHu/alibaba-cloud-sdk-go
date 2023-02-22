package videoenhan

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

// GenerateHumanAnimeStyleVideo invokes the videoenhan.GenerateHumanAnimeStyleVideo API synchronously
func (client *Client) GenerateHumanAnimeStyleVideo(request *GenerateHumanAnimeStyleVideoRequest) (response *GenerateHumanAnimeStyleVideoResponse, err error) {
	response = CreateGenerateHumanAnimeStyleVideoResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateHumanAnimeStyleVideoWithChan invokes the videoenhan.GenerateHumanAnimeStyleVideo API asynchronously
func (client *Client) GenerateHumanAnimeStyleVideoWithChan(request *GenerateHumanAnimeStyleVideoRequest) (<-chan *GenerateHumanAnimeStyleVideoResponse, <-chan error) {
	responseChan := make(chan *GenerateHumanAnimeStyleVideoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateHumanAnimeStyleVideo(request)
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

// GenerateHumanAnimeStyleVideoWithCallback invokes the videoenhan.GenerateHumanAnimeStyleVideo API asynchronously
func (client *Client) GenerateHumanAnimeStyleVideoWithCallback(request *GenerateHumanAnimeStyleVideoRequest, callback func(response *GenerateHumanAnimeStyleVideoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateHumanAnimeStyleVideoResponse
		var err error
		defer close(result)
		response, err = client.GenerateHumanAnimeStyleVideo(request)
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

// GenerateHumanAnimeStyleVideoRequest is the request struct for api GenerateHumanAnimeStyleVideo
type GenerateHumanAnimeStyleVideoRequest struct {
	*requests.RpcRequest
	CartoonStyle string           `position:"Body" name:"CartoonStyle"`
	Async        requests.Boolean `position:"Body" name:"Async"`
	VideoUrl     string           `position:"Body" name:"VideoUrl"`
}

// GenerateHumanAnimeStyleVideoResponse is the response struct for api GenerateHumanAnimeStyleVideo
type GenerateHumanAnimeStyleVideoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGenerateHumanAnimeStyleVideoRequest creates a request to invoke GenerateHumanAnimeStyleVideo API
func CreateGenerateHumanAnimeStyleVideoRequest() (request *GenerateHumanAnimeStyleVideoRequest) {
	request = &GenerateHumanAnimeStyleVideoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videoenhan", "2020-03-20", "GenerateHumanAnimeStyleVideo", "videoenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGenerateHumanAnimeStyleVideoResponse creates a response to parse from GenerateHumanAnimeStyleVideo response
func CreateGenerateHumanAnimeStyleVideoResponse() (response *GenerateHumanAnimeStyleVideoResponse) {
	response = &GenerateHumanAnimeStyleVideoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
