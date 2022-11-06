package edas

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

// GetJavaStartUpConfig invokes the edas.GetJavaStartUpConfig API synchronously
func (client *Client) GetJavaStartUpConfig(request *GetJavaStartUpConfigRequest) (response *GetJavaStartUpConfigResponse, err error) {
	response = CreateGetJavaStartUpConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetJavaStartUpConfigWithChan invokes the edas.GetJavaStartUpConfig API asynchronously
func (client *Client) GetJavaStartUpConfigWithChan(request *GetJavaStartUpConfigRequest) (<-chan *GetJavaStartUpConfigResponse, <-chan error) {
	responseChan := make(chan *GetJavaStartUpConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJavaStartUpConfig(request)
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

// GetJavaStartUpConfigWithCallback invokes the edas.GetJavaStartUpConfig API asynchronously
func (client *Client) GetJavaStartUpConfigWithCallback(request *GetJavaStartUpConfigRequest, callback func(response *GetJavaStartUpConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJavaStartUpConfigResponse
		var err error
		defer close(result)
		response, err = client.GetJavaStartUpConfig(request)
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

// GetJavaStartUpConfigRequest is the request struct for api GetJavaStartUpConfig
type GetJavaStartUpConfigRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// GetJavaStartUpConfigResponse is the response struct for api GetJavaStartUpConfig
type GetJavaStartUpConfigResponse struct {
	*responses.BaseResponse
	Code              int               `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	JavaStartUpConfig JavaStartUpConfig `json:"JavaStartUpConfig" xml:"JavaStartUpConfig"`
}

// CreateGetJavaStartUpConfigRequest creates a request to invoke GetJavaStartUpConfig API
func CreateGetJavaStartUpConfigRequest() (request *GetJavaStartUpConfigRequest) {
	request = &GetJavaStartUpConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetJavaStartUpConfig", "/pop/v5/oam/java_start_up_config", "edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetJavaStartUpConfigResponse creates a response to parse from GetJavaStartUpConfig response
func CreateGetJavaStartUpConfigResponse() (response *GetJavaStartUpConfigResponse) {
	response = &GetJavaStartUpConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
