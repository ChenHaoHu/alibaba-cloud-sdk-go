package rds

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

// UpdateUserBackupFile invokes the rds.UpdateUserBackupFile API synchronously
func (client *Client) UpdateUserBackupFile(request *UpdateUserBackupFileRequest) (response *UpdateUserBackupFileResponse, err error) {
	response = CreateUpdateUserBackupFileResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserBackupFileWithChan invokes the rds.UpdateUserBackupFile API asynchronously
func (client *Client) UpdateUserBackupFileWithChan(request *UpdateUserBackupFileRequest) (<-chan *UpdateUserBackupFileResponse, <-chan error) {
	responseChan := make(chan *UpdateUserBackupFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUserBackupFile(request)
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

// UpdateUserBackupFileWithCallback invokes the rds.UpdateUserBackupFile API asynchronously
func (client *Client) UpdateUserBackupFileWithCallback(request *UpdateUserBackupFileRequest, callback func(response *UpdateUserBackupFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserBackupFileResponse
		var err error
		defer close(result)
		response, err = client.UpdateUserBackupFile(request)
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

// UpdateUserBackupFileRequest is the request struct for api UpdateUserBackupFile
type UpdateUserBackupFileRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Retention            requests.Integer `position:"Query" name:"Retention"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OpsServiceVersion    string           `position:"Query" name:"OpsServiceVersion"`
	Comment              string           `position:"Query" name:"Comment"`
}

// UpdateUserBackupFileResponse is the response struct for api UpdateUserBackupFile
type UpdateUserBackupFileResponse struct {
	*responses.BaseResponse
	BackupId  string `json:"BackupId" xml:"BackupId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateUserBackupFileRequest creates a request to invoke UpdateUserBackupFile API
func CreateUpdateUserBackupFileRequest() (request *UpdateUserBackupFileRequest) {
	request = &UpdateUserBackupFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "UpdateUserBackupFile", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateUserBackupFileResponse creates a response to parse from UpdateUserBackupFile response
func CreateUpdateUserBackupFileResponse() (response *UpdateUserBackupFileResponse) {
	response = &UpdateUserBackupFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
