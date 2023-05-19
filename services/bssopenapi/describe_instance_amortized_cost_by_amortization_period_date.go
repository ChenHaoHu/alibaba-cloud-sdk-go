package bssopenapi

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

// DescribeInstanceAmortizedCostByAmortizationPeriodDate invokes the bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodDate API synchronously
func (client *Client) DescribeInstanceAmortizedCostByAmortizationPeriodDate(request *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) (response *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse, err error) {
	response = CreateDescribeInstanceAmortizedCostByAmortizationPeriodDateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceAmortizedCostByAmortizationPeriodDateWithChan invokes the bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodDate API asynchronously
func (client *Client) DescribeInstanceAmortizedCostByAmortizationPeriodDateWithChan(request *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) (<-chan *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAmortizedCostByAmortizationPeriodDate(request)
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

// DescribeInstanceAmortizedCostByAmortizationPeriodDateWithCallback invokes the bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodDate API asynchronously
func (client *Client) DescribeInstanceAmortizedCostByAmortizationPeriodDateWithCallback(request *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest, callback func(response *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAmortizedCostByAmortizationPeriodDate(request)
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

// DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest is the request struct for api DescribeInstanceAmortizedCostByAmortizationPeriodDate
type DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest struct {
	*requests.RpcRequest
	ProductCode           string           `position:"Body" name:"ProductCode"`
	AmortizationDateStart string           `position:"Body" name:"AmortizationDateStart"`
	SubscriptionType      string           `position:"Body" name:"SubscriptionType"`
	CostUnitCode          string           `position:"Body" name:"CostUnitCode"`
	NextToken             string           `position:"Body" name:"NextToken"`
	BillUserIdList        *[]string        `position:"Body" name:"BillUserIdList"  type:"Repeated"`
	ProductDetail         string           `position:"Body" name:"ProductDetail"`
	BillOwnerIdList       *[]string        `position:"Body" name:"BillOwnerIdList"  type:"Repeated"`
	BillingCycle          string           `position:"Body" name:"BillingCycle"`
	AmortizationDateEnd   string           `position:"Body" name:"AmortizationDateEnd"`
	InstanceIdList        *[]string        `position:"Body" name:"InstanceIdList"  type:"Repeated"`
	MaxResults            requests.Integer `position:"Body" name:"MaxResults"`
}

// DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse is the response struct for api DescribeInstanceAmortizedCostByAmortizationPeriodDate
type DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeInstanceAmortizedCostByAmortizationPeriodDateRequest creates a request to invoke DescribeInstanceAmortizedCostByAmortizationPeriodDate API
func CreateDescribeInstanceAmortizedCostByAmortizationPeriodDateRequest() (request *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) {
	request = &DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "DescribeInstanceAmortizedCostByAmortizationPeriodDate", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceAmortizedCostByAmortizationPeriodDateResponse creates a response to parse from DescribeInstanceAmortizedCostByAmortizationPeriodDate response
func CreateDescribeInstanceAmortizedCostByAmortizationPeriodDateResponse() (response *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) {
	response = &DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
