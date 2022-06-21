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

// DeleteLiveStreamRecordIndexFiles invokes the live.DeleteLiveStreamRecordIndexFiles API synchronously
func (client *Client) DeleteLiveStreamRecordIndexFiles(request *DeleteLiveStreamRecordIndexFilesRequest) (response *DeleteLiveStreamRecordIndexFilesResponse, err error) {
	response = CreateDeleteLiveStreamRecordIndexFilesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveStreamRecordIndexFilesWithChan invokes the live.DeleteLiveStreamRecordIndexFiles API asynchronously
func (client *Client) DeleteLiveStreamRecordIndexFilesWithChan(request *DeleteLiveStreamRecordIndexFilesRequest) (<-chan *DeleteLiveStreamRecordIndexFilesResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveStreamRecordIndexFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveStreamRecordIndexFiles(request)
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

// DeleteLiveStreamRecordIndexFilesWithCallback invokes the live.DeleteLiveStreamRecordIndexFiles API asynchronously
func (client *Client) DeleteLiveStreamRecordIndexFilesWithCallback(request *DeleteLiveStreamRecordIndexFilesRequest, callback func(response *DeleteLiveStreamRecordIndexFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveStreamRecordIndexFilesResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveStreamRecordIndexFiles(request)
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

// DeleteLiveStreamRecordIndexFilesRequest is the request struct for api DeleteLiveStreamRecordIndexFiles
type DeleteLiveStreamRecordIndexFilesRequest struct {
	*requests.RpcRequest
	RemoveFile string           `position:"Query" name:"RemoveFile"`
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	RecordId   *[]string        `position:"Query" name:"RecordId"  type:"Repeated"`
}

// DeleteLiveStreamRecordIndexFilesResponse is the response struct for api DeleteLiveStreamRecordIndexFiles
type DeleteLiveStreamRecordIndexFilesResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Message              string               `json:"Message" xml:"Message"`
	Code                 string               `json:"Code" xml:"Code"`
	RecordDeleteInfoList RecordDeleteInfoList `json:"RecordDeleteInfoList" xml:"RecordDeleteInfoList"`
}

// CreateDeleteLiveStreamRecordIndexFilesRequest creates a request to invoke DeleteLiveStreamRecordIndexFiles API
func CreateDeleteLiveStreamRecordIndexFilesRequest() (request *DeleteLiveStreamRecordIndexFilesRequest) {
	request = &DeleteLiveStreamRecordIndexFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamRecordIndexFiles", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveStreamRecordIndexFilesResponse creates a response to parse from DeleteLiveStreamRecordIndexFiles response
func CreateDeleteLiveStreamRecordIndexFilesResponse() (response *DeleteLiveStreamRecordIndexFilesResponse) {
	response = &DeleteLiveStreamRecordIndexFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
