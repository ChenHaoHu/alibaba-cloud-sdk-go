package cbn

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

// CreateTrafficMarkingPolicy invokes the cbn.CreateTrafficMarkingPolicy API synchronously
func (client *Client) CreateTrafficMarkingPolicy(request *CreateTrafficMarkingPolicyRequest) (response *CreateTrafficMarkingPolicyResponse, err error) {
	response = CreateCreateTrafficMarkingPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTrafficMarkingPolicyWithChan invokes the cbn.CreateTrafficMarkingPolicy API asynchronously
func (client *Client) CreateTrafficMarkingPolicyWithChan(request *CreateTrafficMarkingPolicyRequest) (<-chan *CreateTrafficMarkingPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateTrafficMarkingPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTrafficMarkingPolicy(request)
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

// CreateTrafficMarkingPolicyWithCallback invokes the cbn.CreateTrafficMarkingPolicy API asynchronously
func (client *Client) CreateTrafficMarkingPolicyWithCallback(request *CreateTrafficMarkingPolicyRequest, callback func(response *CreateTrafficMarkingPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTrafficMarkingPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateTrafficMarkingPolicy(request)
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

// CreateTrafficMarkingPolicyRequest is the request struct for api CreateTrafficMarkingPolicy
type CreateTrafficMarkingPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                 requests.Integer                               `position:"Query" name:"ResourceOwnerId"`
	ClientToken                     string                                         `position:"Query" name:"ClientToken"`
	TrafficMarkingPolicyDescription string                                         `position:"Query" name:"TrafficMarkingPolicyDescription"`
	TrafficMarkingPolicyName        string                                         `position:"Query" name:"TrafficMarkingPolicyName"`
	DryRun                          requests.Boolean                               `position:"Query" name:"DryRun"`
	TrafficMatchRules               *[]CreateTrafficMarkingPolicyTrafficMatchRules `position:"Query" name:"TrafficMatchRules"  type:"Repeated"`
	ResourceOwnerAccount            string                                         `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                    string                                         `position:"Query" name:"OwnerAccount"`
	OwnerId                         requests.Integer                               `position:"Query" name:"OwnerId"`
	TransitRouterId                 string                                         `position:"Query" name:"TransitRouterId"`
	Priority                        requests.Integer                               `position:"Query" name:"Priority"`
	MarkingDscp                     requests.Integer                               `position:"Query" name:"MarkingDscp"`
}

// CreateTrafficMarkingPolicyTrafficMatchRules is a repeated param struct in CreateTrafficMarkingPolicyRequest
type CreateTrafficMarkingPolicyTrafficMatchRules struct {
	DstPortRange                *[]string `name:"DstPortRange" type:"Repeated"`
	MatchDscp                   string    `name:"MatchDscp"`
	Protocol                    string    `name:"Protocol"`
	TrafficMatchRuleDescription string    `name:"TrafficMatchRuleDescription"`
	SrcPortRange                *[]string `name:"SrcPortRange" type:"Repeated"`
	DstCidr                     string    `name:"DstCidr"`
	TrafficMatchRuleName        string    `name:"TrafficMatchRuleName"`
	SrcCidr                     string    `name:"SrcCidr"`
}

// CreateTrafficMarkingPolicyResponse is the response struct for api CreateTrafficMarkingPolicy
type CreateTrafficMarkingPolicyResponse struct {
	*responses.BaseResponse
	TrafficMarkingPolicyId string `json:"TrafficMarkingPolicyId" xml:"TrafficMarkingPolicyId"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTrafficMarkingPolicyRequest creates a request to invoke CreateTrafficMarkingPolicy API
func CreateCreateTrafficMarkingPolicyRequest() (request *CreateTrafficMarkingPolicyRequest) {
	request = &CreateTrafficMarkingPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateTrafficMarkingPolicy", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTrafficMarkingPolicyResponse creates a response to parse from CreateTrafficMarkingPolicy response
func CreateCreateTrafficMarkingPolicyResponse() (response *CreateTrafficMarkingPolicyResponse) {
	response = &CreateTrafficMarkingPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
