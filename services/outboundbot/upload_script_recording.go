package outboundbot

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

// UploadScriptRecording invokes the outboundbot.UploadScriptRecording API synchronously
func (client *Client) UploadScriptRecording(request *UploadScriptRecordingRequest) (response *UploadScriptRecordingResponse, err error) {
	response = CreateUploadScriptRecordingResponse()
	err = client.DoAction(request, response)
	return
}

// UploadScriptRecordingWithChan invokes the outboundbot.UploadScriptRecording API asynchronously
func (client *Client) UploadScriptRecordingWithChan(request *UploadScriptRecordingRequest) (<-chan *UploadScriptRecordingResponse, <-chan error) {
	responseChan := make(chan *UploadScriptRecordingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadScriptRecording(request)
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

// UploadScriptRecordingWithCallback invokes the outboundbot.UploadScriptRecording API asynchronously
func (client *Client) UploadScriptRecordingWithCallback(request *UploadScriptRecordingRequest, callback func(response *UploadScriptRecordingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadScriptRecordingResponse
		var err error
		defer close(result)
		response, err = client.UploadScriptRecording(request)
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

// UploadScriptRecordingRequest is the request struct for api UploadScriptRecording
type UploadScriptRecordingRequest struct {
	*requests.RpcRequest
	Content    string `position:"Query" name:"Content"`
	ScriptId   string `position:"Query" name:"ScriptId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	FileName   string `position:"Query" name:"FileName"`
	FileId     string `position:"Query" name:"FileId"`
}

// UploadScriptRecordingResponse is the response struct for api UploadScriptRecording
type UploadScriptRecordingResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Uuid           string `json:"Uuid" xml:"Uuid"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
}

// CreateUploadScriptRecordingRequest creates a request to invoke UploadScriptRecording API
func CreateUploadScriptRecordingRequest() (request *UploadScriptRecordingRequest) {
	request = &UploadScriptRecordingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "UploadScriptRecording", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadScriptRecordingResponse creates a response to parse from UploadScriptRecording response
func CreateUploadScriptRecordingResponse() (response *UploadScriptRecordingResponse) {
	response = &UploadScriptRecordingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
