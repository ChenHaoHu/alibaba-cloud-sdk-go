package dg

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

// ConnectDatabase invokes the dg.ConnectDatabase API synchronously
func (client *Client) ConnectDatabase(request *ConnectDatabaseRequest) (response *ConnectDatabaseResponse, err error) {
	response = CreateConnectDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// ConnectDatabaseWithChan invokes the dg.ConnectDatabase API asynchronously
func (client *Client) ConnectDatabaseWithChan(request *ConnectDatabaseRequest) (<-chan *ConnectDatabaseResponse, <-chan error) {
	responseChan := make(chan *ConnectDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConnectDatabase(request)
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

// ConnectDatabaseWithCallback invokes the dg.ConnectDatabase API asynchronously
func (client *Client) ConnectDatabaseWithCallback(request *ConnectDatabaseRequest, callback func(response *ConnectDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConnectDatabaseResponse
		var err error
		defer close(result)
		response, err = client.ConnectDatabase(request)
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

// ConnectDatabaseRequest is the request struct for api ConnectDatabase
type ConnectDatabaseRequest struct {
	*requests.RpcRequest
	DbName     string           `position:"Body" name:"DbName"`
	Port       requests.Integer `position:"Body" name:"Port"`
	DbPassword string           `position:"Body" name:"DbPassword"`
	Host       string           `position:"Body" name:"Host"`
	DbType     string           `position:"Body" name:"DbType"`
	DbUserName string           `position:"Body" name:"DbUserName"`
	GatewayId  string           `position:"Body" name:"GatewayId"`
}

// ConnectDatabaseResponse is the response struct for api ConnectDatabase
type ConnectDatabaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateConnectDatabaseRequest creates a request to invoke ConnectDatabase API
func CreateConnectDatabaseRequest() (request *ConnectDatabaseRequest) {
	request = &ConnectDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dg", "2019-03-27", "ConnectDatabase", "dg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConnectDatabaseResponse creates a response to parse from ConnectDatabase response
func CreateConnectDatabaseResponse() (response *ConnectDatabaseResponse) {
	response = &ConnectDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
