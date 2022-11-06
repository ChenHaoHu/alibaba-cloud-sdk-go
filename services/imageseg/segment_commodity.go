package imageseg

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

// SegmentCommodity invokes the imageseg.SegmentCommodity API synchronously
func (client *Client) SegmentCommodity(request *SegmentCommodityRequest) (response *SegmentCommodityResponse, err error) {
	response = CreateSegmentCommodityResponse()
	err = client.DoAction(request, response)
	return
}

// SegmentCommodityWithChan invokes the imageseg.SegmentCommodity API asynchronously
func (client *Client) SegmentCommodityWithChan(request *SegmentCommodityRequest) (<-chan *SegmentCommodityResponse, <-chan error) {
	responseChan := make(chan *SegmentCommodityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SegmentCommodity(request)
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

// SegmentCommodityWithCallback invokes the imageseg.SegmentCommodity API asynchronously
func (client *Client) SegmentCommodityWithCallback(request *SegmentCommodityRequest, callback func(response *SegmentCommodityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SegmentCommodityResponse
		var err error
		defer close(result)
		response, err = client.SegmentCommodity(request)
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

// SegmentCommodityRequest is the request struct for api SegmentCommodity
type SegmentCommodityRequest struct {
	*requests.RpcRequest
	ReturnForm     string `position:"Query" name:"ReturnForm"`
	OssFile        string `position:"Query" name:"OssFile"`
	RequestProxyBy string `position:"Query" name:"RequestProxyBy"`
	ImageURL       string `position:"Query" name:"ImageURL"`
}

// SegmentCommodityResponse is the response struct for api SegmentCommodity
type SegmentCommodityResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSegmentCommodityRequest creates a request to invoke SegmentCommodity API
func CreateSegmentCommodityRequest() (request *SegmentCommodityRequest) {
	request = &SegmentCommodityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageseg", "2019-12-30", "SegmentCommodity", "", "")
	request.Method = requests.POST
	return
}

// CreateSegmentCommodityResponse creates a response to parse from SegmentCommodity response
func CreateSegmentCommodityResponse() (response *SegmentCommodityResponse) {
	response = &SegmentCommodityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
