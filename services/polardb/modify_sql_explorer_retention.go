package polardb

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

// ModifySQLExplorerRetention invokes the polardb.ModifySQLExplorerRetention API synchronously
func (client *Client) ModifySQLExplorerRetention(request *ModifySQLExplorerRetentionRequest) (response *ModifySQLExplorerRetentionResponse, err error) {
	response = CreateModifySQLExplorerRetentionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySQLExplorerRetentionWithChan invokes the polardb.ModifySQLExplorerRetention API asynchronously
func (client *Client) ModifySQLExplorerRetentionWithChan(request *ModifySQLExplorerRetentionRequest) (<-chan *ModifySQLExplorerRetentionResponse, <-chan error) {
	responseChan := make(chan *ModifySQLExplorerRetentionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySQLExplorerRetention(request)
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

// ModifySQLExplorerRetentionWithCallback invokes the polardb.ModifySQLExplorerRetention API asynchronously
func (client *Client) ModifySQLExplorerRetentionWithCallback(request *ModifySQLExplorerRetentionRequest, callback func(response *ModifySQLExplorerRetentionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySQLExplorerRetentionResponse
		var err error
		defer close(result)
		response, err = client.ModifySQLExplorerRetention(request)
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

// ModifySQLExplorerRetentionRequest is the request struct for api ModifySQLExplorerRetention
type ModifySQLExplorerRetentionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ConfigValue          string           `position:"Query" name:"ConfigValue"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifySQLExplorerRetentionResponse is the response struct for api ModifySQLExplorerRetention
type ModifySQLExplorerRetentionResponse struct {
	*responses.BaseResponse
	TaskId         int    `json:"TaskId" xml:"TaskId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceID   int    `json:"DBInstanceID" xml:"DBInstanceID"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
}

// CreateModifySQLExplorerRetentionRequest creates a request to invoke ModifySQLExplorerRetention API
func CreateModifySQLExplorerRetentionRequest() (request *ModifySQLExplorerRetentionRequest) {
	request = &ModifySQLExplorerRetentionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifySQLExplorerRetention", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySQLExplorerRetentionResponse creates a response to parse from ModifySQLExplorerRetention response
func CreateModifySQLExplorerRetentionResponse() (response *ModifySQLExplorerRetentionResponse) {
	response = &ModifySQLExplorerRetentionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
