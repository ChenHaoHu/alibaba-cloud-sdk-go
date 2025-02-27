package oceanbasepro

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

// CreateInstance invokes the oceanbasepro.CreateInstance API synchronously
func (client *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
	response = CreateCreateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateInstanceWithChan invokes the oceanbasepro.CreateInstance API asynchronously
func (client *Client) CreateInstanceWithChan(request *CreateInstanceRequest) (<-chan *CreateInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateInstance(request)
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

// CreateInstanceWithCallback invokes the oceanbasepro.CreateInstance API asynchronously
func (client *Client) CreateInstanceWithCallback(request *CreateInstanceRequest, callback func(response *CreateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateInstance(request)
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

// CreateInstanceRequest is the request struct for api CreateInstance
type CreateInstanceRequest struct {
	*requests.RpcRequest
	IsolationOptimization string           `position:"Body" name:"IsolationOptimization"`
	InstanceClass         string           `position:"Body" name:"InstanceClass"`
	ResourceGroupId       string           `position:"Body" name:"ResourceGroupId"`
	AutoRenewPeriod       requests.Integer `position:"Body" name:"AutoRenewPeriod"`
	Period                requests.Integer `position:"Body" name:"Period"`
	DryRun                requests.Boolean `position:"Body" name:"DryRun"`
	DiskSize              requests.Integer `position:"Body" name:"DiskSize"`
	Zones                 string           `position:"Body" name:"Zones"`
	DiskType              string           `position:"Body" name:"DiskType"`
	ObVersion             string           `position:"Body" name:"ObVersion"`
	PeriodUnit            string           `position:"Body" name:"PeriodUnit"`
	InstanceName          string           `position:"Body" name:"InstanceName"`
	ReplicaMode           string           `position:"Body" name:"ReplicaMode"`
	AutoRenew             requests.Boolean `position:"Body" name:"AutoRenew"`
	Series                string           `position:"Body" name:"Series"`
	ChargeType            string           `position:"Body" name:"ChargeType"`
	Bid                   string           `position:"Query" name:"Bid"`
}

// CreateInstanceResponse is the response struct for api CreateInstance
type CreateInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateInstanceRequest creates a request to invoke CreateInstance API
func CreateCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "CreateInstance", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateInstanceResponse creates a response to parse from CreateInstance response
func CreateCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
