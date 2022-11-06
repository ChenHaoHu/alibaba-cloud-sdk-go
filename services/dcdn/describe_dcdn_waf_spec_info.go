package dcdn

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

// DescribeDcdnWafSpecInfo invokes the dcdn.DescribeDcdnWafSpecInfo API synchronously
func (client *Client) DescribeDcdnWafSpecInfo(request *DescribeDcdnWafSpecInfoRequest) (response *DescribeDcdnWafSpecInfoResponse, err error) {
	response = CreateDescribeDcdnWafSpecInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnWafSpecInfoWithChan invokes the dcdn.DescribeDcdnWafSpecInfo API asynchronously
func (client *Client) DescribeDcdnWafSpecInfoWithChan(request *DescribeDcdnWafSpecInfoRequest) (<-chan *DescribeDcdnWafSpecInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnWafSpecInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnWafSpecInfo(request)
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

// DescribeDcdnWafSpecInfoWithCallback invokes the dcdn.DescribeDcdnWafSpecInfo API asynchronously
func (client *Client) DescribeDcdnWafSpecInfoWithCallback(request *DescribeDcdnWafSpecInfoRequest, callback func(response *DescribeDcdnWafSpecInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnWafSpecInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnWafSpecInfo(request)
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

// DescribeDcdnWafSpecInfoRequest is the request struct for api DescribeDcdnWafSpecInfo
type DescribeDcdnWafSpecInfoRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnWafSpecInfoResponse is the response struct for api DescribeDcdnWafSpecInfo
type DescribeDcdnWafSpecInfoResponse struct {
	*responses.BaseResponse
	Edition   string         `json:"Edition" xml:"Edition"`
	RequestId string         `json:"RequestId" xml:"RequestId"`
	SpecInfos []RuleInfoItem `json:"SpecInfos" xml:"SpecInfos"`
}

// CreateDescribeDcdnWafSpecInfoRequest creates a request to invoke DescribeDcdnWafSpecInfo API
func CreateDescribeDcdnWafSpecInfoRequest() (request *DescribeDcdnWafSpecInfoRequest) {
	request = &DescribeDcdnWafSpecInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnWafSpecInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnWafSpecInfoResponse creates a response to parse from DescribeDcdnWafSpecInfo response
func CreateDescribeDcdnWafSpecInfoResponse() (response *DescribeDcdnWafSpecInfoResponse) {
	response = &DescribeDcdnWafSpecInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
