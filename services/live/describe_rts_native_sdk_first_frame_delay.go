package live

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

// DescribeRTSNativeSDKFirstFrameDelay invokes the live.DescribeRTSNativeSDKFirstFrameDelay API synchronously
func (client *Client) DescribeRTSNativeSDKFirstFrameDelay(request *DescribeRTSNativeSDKFirstFrameDelayRequest) (response *DescribeRTSNativeSDKFirstFrameDelayResponse, err error) {
	response = CreateDescribeRTSNativeSDKFirstFrameDelayResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRTSNativeSDKFirstFrameDelayWithChan invokes the live.DescribeRTSNativeSDKFirstFrameDelay API asynchronously
func (client *Client) DescribeRTSNativeSDKFirstFrameDelayWithChan(request *DescribeRTSNativeSDKFirstFrameDelayRequest) (<-chan *DescribeRTSNativeSDKFirstFrameDelayResponse, <-chan error) {
	responseChan := make(chan *DescribeRTSNativeSDKFirstFrameDelayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRTSNativeSDKFirstFrameDelay(request)
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

// DescribeRTSNativeSDKFirstFrameDelayWithCallback invokes the live.DescribeRTSNativeSDKFirstFrameDelay API asynchronously
func (client *Client) DescribeRTSNativeSDKFirstFrameDelayWithCallback(request *DescribeRTSNativeSDKFirstFrameDelayRequest, callback func(response *DescribeRTSNativeSDKFirstFrameDelayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRTSNativeSDKFirstFrameDelayResponse
		var err error
		defer close(result)
		response, err = client.DescribeRTSNativeSDKFirstFrameDelay(request)
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

// DescribeRTSNativeSDKFirstFrameDelayRequest is the request struct for api DescribeRTSNativeSDKFirstFrameDelay
type DescribeRTSNativeSDKFirstFrameDelayRequest struct {
	*requests.RpcRequest
	EndTime        string    `position:"Query" name:"EndTime"`
	DomainNameList *[]string `position:"Query" name:"DomainNameList"  type:"Json"`
	StartTime      string    `position:"Query" name:"StartTime"`
	DataInterval   string    `position:"Query" name:"DataInterval"`
}

// DescribeRTSNativeSDKFirstFrameDelayResponse is the response struct for api DescribeRTSNativeSDKFirstFrameDelay
type DescribeRTSNativeSDKFirstFrameDelayResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DataInterval   string `json:"DataInterval" xml:"DataInterval"`
	StartTime      string `json:"StartTime" xml:"StartTime"`
	EndTime        string `json:"EndTime" xml:"EndTime"`
	FrameDelayData []Data `json:"FrameDelayData" xml:"FrameDelayData"`
}

// CreateDescribeRTSNativeSDKFirstFrameDelayRequest creates a request to invoke DescribeRTSNativeSDKFirstFrameDelay API
func CreateDescribeRTSNativeSDKFirstFrameDelayRequest() (request *DescribeRTSNativeSDKFirstFrameDelayRequest) {
	request = &DescribeRTSNativeSDKFirstFrameDelayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeRTSNativeSDKFirstFrameDelay", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRTSNativeSDKFirstFrameDelayResponse creates a response to parse from DescribeRTSNativeSDKFirstFrameDelay response
func CreateDescribeRTSNativeSDKFirstFrameDelayResponse() (response *DescribeRTSNativeSDKFirstFrameDelayResponse) {
	response = &DescribeRTSNativeSDKFirstFrameDelayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
