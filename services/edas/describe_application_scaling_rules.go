package edas

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

// DescribeApplicationScalingRules invokes the edas.DescribeApplicationScalingRules API synchronously
func (client *Client) DescribeApplicationScalingRules(request *DescribeApplicationScalingRulesRequest) (response *DescribeApplicationScalingRulesResponse, err error) {
	response = CreateDescribeApplicationScalingRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApplicationScalingRulesWithChan invokes the edas.DescribeApplicationScalingRules API asynchronously
func (client *Client) DescribeApplicationScalingRulesWithChan(request *DescribeApplicationScalingRulesRequest) (<-chan *DescribeApplicationScalingRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeApplicationScalingRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApplicationScalingRules(request)
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

// DescribeApplicationScalingRulesWithCallback invokes the edas.DescribeApplicationScalingRules API asynchronously
func (client *Client) DescribeApplicationScalingRulesWithCallback(request *DescribeApplicationScalingRulesRequest, callback func(response *DescribeApplicationScalingRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApplicationScalingRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeApplicationScalingRules(request)
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

// DescribeApplicationScalingRulesRequest is the request struct for api DescribeApplicationScalingRules
type DescribeApplicationScalingRulesRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// DescribeApplicationScalingRulesResponse is the response struct for api DescribeApplicationScalingRules
type DescribeApplicationScalingRulesResponse struct {
	*responses.BaseResponse
	Code            int             `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	AppScalingRules AppScalingRules `json:"AppScalingRules" xml:"AppScalingRules"`
}

// CreateDescribeApplicationScalingRulesRequest creates a request to invoke DescribeApplicationScalingRules API
func CreateDescribeApplicationScalingRulesRequest() (request *DescribeApplicationScalingRulesRequest) {
	request = &DescribeApplicationScalingRulesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DescribeApplicationScalingRules", "/pop/v1/eam/scale/application_scaling_rules", "edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeApplicationScalingRulesResponse creates a response to parse from DescribeApplicationScalingRules response
func CreateDescribeApplicationScalingRulesResponse() (response *DescribeApplicationScalingRulesResponse) {
	response = &DescribeApplicationScalingRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
