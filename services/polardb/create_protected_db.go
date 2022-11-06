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

// CreateProtectedDB invokes the polardb.CreateProtectedDB API synchronously
func (client *Client) CreateProtectedDB(request *CreateProtectedDBRequest) (response *CreateProtectedDBResponse, err error) {
	response = CreateCreateProtectedDBResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProtectedDBWithChan invokes the polardb.CreateProtectedDB API asynchronously
func (client *Client) CreateProtectedDBWithChan(request *CreateProtectedDBRequest) (<-chan *CreateProtectedDBResponse, <-chan error) {
	responseChan := make(chan *CreateProtectedDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProtectedDB(request)
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

// CreateProtectedDBWithCallback invokes the polardb.CreateProtectedDB API asynchronously
func (client *Client) CreateProtectedDBWithCallback(request *CreateProtectedDBRequest, callback func(response *CreateProtectedDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProtectedDBResponse
		var err error
		defer close(result)
		response, err = client.CreateProtectedDB(request)
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

// CreateProtectedDBRequest is the request struct for api CreateProtectedDB
type CreateProtectedDBRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBDescription        string           `position:"Query" name:"DBDescription"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
	DBName               string           `position:"Query" name:"DBName"`
	CharacterSetName     string           `position:"Query" name:"CharacterSetName"`
}

// CreateProtectedDBResponse is the response struct for api CreateProtectedDB
type CreateProtectedDBResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateProtectedDBRequest creates a request to invoke CreateProtectedDB API
func CreateCreateProtectedDBRequest() (request *CreateProtectedDBRequest) {
	request = &CreateProtectedDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CreateProtectedDB", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProtectedDBResponse creates a response to parse from CreateProtectedDB response
func CreateCreateProtectedDBResponse() (response *CreateProtectedDBResponse) {
	response = &CreateProtectedDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
