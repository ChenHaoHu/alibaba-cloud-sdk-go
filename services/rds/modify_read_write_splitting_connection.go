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

// ModifyReadWriteSplittingConnection invokes the rds.ModifyReadWriteSplittingConnection API synchronously
func (client *Client) ModifyReadWriteSplittingConnection(request *ModifyReadWriteSplittingConnectionRequest) (response *ModifyReadWriteSplittingConnectionResponse, err error) {
	response = CreateModifyReadWriteSplittingConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReadWriteSplittingConnectionWithChan invokes the rds.ModifyReadWriteSplittingConnection API asynchronously
func (client *Client) ModifyReadWriteSplittingConnectionWithChan(request *ModifyReadWriteSplittingConnectionRequest) (<-chan *ModifyReadWriteSplittingConnectionResponse, <-chan error) {
	responseChan := make(chan *ModifyReadWriteSplittingConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReadWriteSplittingConnection(request)
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

// ModifyReadWriteSplittingConnectionWithCallback invokes the rds.ModifyReadWriteSplittingConnection API asynchronously
func (client *Client) ModifyReadWriteSplittingConnectionWithCallback(request *ModifyReadWriteSplittingConnectionRequest, callback func(response *ModifyReadWriteSplittingConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReadWriteSplittingConnectionResponse
		var err error
		defer close(result)
		response, err = client.ModifyReadWriteSplittingConnection(request)
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

// ModifyReadWriteSplittingConnectionRequest is the request struct for api ModifyReadWriteSplittingConnection
type ModifyReadWriteSplittingConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string           `position:"Query" name:"ConnectionStringPrefix"`
	DistributionType       string           `position:"Query" name:"DistributionType"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	Weight                 string           `position:"Query" name:"Weight"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	Port                   string           `position:"Query" name:"Port"`
	MaxDelayTime           string           `position:"Query" name:"MaxDelayTime"`
}

// ModifyReadWriteSplittingConnectionResponse is the response struct for api ModifyReadWriteSplittingConnection
type ModifyReadWriteSplittingConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyReadWriteSplittingConnectionRequest creates a request to invoke ModifyReadWriteSplittingConnection API
func CreateModifyReadWriteSplittingConnectionRequest() (request *ModifyReadWriteSplittingConnectionRequest) {
	request = &ModifyReadWriteSplittingConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyReadWriteSplittingConnection", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyReadWriteSplittingConnectionResponse creates a response to parse from ModifyReadWriteSplittingConnection response
func CreateModifyReadWriteSplittingConnectionResponse() (response *ModifyReadWriteSplittingConnectionResponse) {
	response = &ModifyReadWriteSplittingConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
