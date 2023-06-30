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

// ReleaseInstanceConnection invokes the rds.ReleaseInstanceConnection API synchronously
func (client *Client) ReleaseInstanceConnection(request *ReleaseInstanceConnectionRequest) (response *ReleaseInstanceConnectionResponse, err error) {
	response = CreateReleaseInstanceConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseInstanceConnectionWithChan invokes the rds.ReleaseInstanceConnection API asynchronously
func (client *Client) ReleaseInstanceConnectionWithChan(request *ReleaseInstanceConnectionRequest) (<-chan *ReleaseInstanceConnectionResponse, <-chan error) {
	responseChan := make(chan *ReleaseInstanceConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseInstanceConnection(request)
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

// ReleaseInstanceConnectionWithCallback invokes the rds.ReleaseInstanceConnection API asynchronously
func (client *Client) ReleaseInstanceConnectionWithCallback(request *ReleaseInstanceConnectionRequest, callback func(response *ReleaseInstanceConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseInstanceConnectionResponse
		var err error
		defer close(result)
		response, err = client.ReleaseInstanceConnection(request)
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

// ReleaseInstanceConnectionRequest is the request struct for api ReleaseInstanceConnection
type ReleaseInstanceConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId            string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	CurrentConnectionString string           `position:"Query" name:"CurrentConnectionString"`
	InstanceNetworkType     string           `position:"Query" name:"InstanceNetworkType"`
}

// ReleaseInstanceConnectionResponse is the response struct for api ReleaseInstanceConnection
type ReleaseInstanceConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseInstanceConnectionRequest creates a request to invoke ReleaseInstanceConnection API
func CreateReleaseInstanceConnectionRequest() (request *ReleaseInstanceConnectionRequest) {
	request = &ReleaseInstanceConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ReleaseInstanceConnection", "", "")
	request.Method = requests.POST
	return
}

// CreateReleaseInstanceConnectionResponse creates a response to parse from ReleaseInstanceConnection response
func CreateReleaseInstanceConnectionResponse() (response *ReleaseInstanceConnectionResponse) {
	response = &ReleaseInstanceConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
