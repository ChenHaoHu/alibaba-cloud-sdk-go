package swas_open

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

// CreateCustomImage invokes the swas_open.CreateCustomImage API synchronously
func (client *Client) CreateCustomImage(request *CreateCustomImageRequest) (response *CreateCustomImageResponse, err error) {
	response = CreateCreateCustomImageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCustomImageWithChan invokes the swas_open.CreateCustomImage API asynchronously
func (client *Client) CreateCustomImageWithChan(request *CreateCustomImageRequest) (<-chan *CreateCustomImageResponse, <-chan error) {
	responseChan := make(chan *CreateCustomImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCustomImage(request)
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

// CreateCustomImageWithCallback invokes the swas_open.CreateCustomImage API asynchronously
func (client *Client) CreateCustomImageWithCallback(request *CreateCustomImageRequest, callback func(response *CreateCustomImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCustomImageResponse
		var err error
		defer close(result)
		response, err = client.CreateCustomImage(request)
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

// CreateCustomImageRequest is the request struct for api CreateCustomImage
type CreateCustomImageRequest struct {
	*requests.RpcRequest
	SystemSnapshotId string `position:"Query" name:"SystemSnapshotId"`
	DataSnapshotId   string `position:"Query" name:"DataSnapshotId"`
	ClientToken      string `position:"Query" name:"ClientToken"`
	Description      string `position:"Query" name:"Description"`
	Platform         string `position:"Query" name:"Platform"`
	ImageName        string `position:"Query" name:"ImageName"`
	InstanceId       string `position:"Query" name:"InstanceId"`
}

// CreateCustomImageResponse is the response struct for api CreateCustomImage
type CreateCustomImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ImageId   string `json:"ImageId" xml:"ImageId"`
}

// CreateCreateCustomImageRequest creates a request to invoke CreateCustomImage API
func CreateCreateCustomImageRequest() (request *CreateCustomImageRequest) {
	request = &CreateCustomImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "CreateCustomImage", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCustomImageResponse creates a response to parse from CreateCustomImage response
func CreateCreateCustomImageResponse() (response *CreateCustomImageResponse) {
	response = &CreateCustomImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
