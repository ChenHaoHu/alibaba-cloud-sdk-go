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

// DeleteSnapshotFiles invokes the live.DeleteSnapshotFiles API synchronously
func (client *Client) DeleteSnapshotFiles(request *DeleteSnapshotFilesRequest) (response *DeleteSnapshotFilesResponse, err error) {
	response = CreateDeleteSnapshotFilesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSnapshotFilesWithChan invokes the live.DeleteSnapshotFiles API asynchronously
func (client *Client) DeleteSnapshotFilesWithChan(request *DeleteSnapshotFilesRequest) (<-chan *DeleteSnapshotFilesResponse, <-chan error) {
	responseChan := make(chan *DeleteSnapshotFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSnapshotFiles(request)
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

// DeleteSnapshotFilesWithCallback invokes the live.DeleteSnapshotFiles API asynchronously
func (client *Client) DeleteSnapshotFilesWithCallback(request *DeleteSnapshotFilesRequest, callback func(response *DeleteSnapshotFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSnapshotFilesResponse
		var err error
		defer close(result)
		response, err = client.DeleteSnapshotFiles(request)
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

// DeleteSnapshotFilesRequest is the request struct for api DeleteSnapshotFiles
type DeleteSnapshotFilesRequest struct {
	*requests.RpcRequest
	RemoveFile          requests.Boolean `position:"Query" name:"RemoveFile"`
	AppName             string           `position:"Query" name:"AppName"`
	StreamName          string           `position:"Query" name:"StreamName"`
	DomainName          string           `position:"Query" name:"DomainName"`
	OwnerId             requests.Integer `position:"Query" name:"OwnerId"`
	CreateTimestampList *[]string        `position:"Query" name:"CreateTimestampList"  type:"Repeated"`
}

// DeleteSnapshotFilesResponse is the response struct for api DeleteSnapshotFiles
type DeleteSnapshotFilesResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	SuccessCount           int                    `json:"SuccessCount" xml:"SuccessCount"`
	FailureCount           int                    `json:"FailureCount" xml:"FailureCount"`
	SnapshotDeleteInfoList SnapshotDeleteInfoList `json:"SnapshotDeleteInfoList" xml:"SnapshotDeleteInfoList"`
}

// CreateDeleteSnapshotFilesRequest creates a request to invoke DeleteSnapshotFiles API
func CreateDeleteSnapshotFilesRequest() (request *DeleteSnapshotFilesRequest) {
	request = &DeleteSnapshotFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteSnapshotFiles", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSnapshotFilesResponse creates a response to parse from DeleteSnapshotFiles response
func CreateDeleteSnapshotFilesResponse() (response *DeleteSnapshotFilesResponse) {
	response = &DeleteSnapshotFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
