package cdn

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

// DescribeCdnUserBillPrediction invokes the cdn.DescribeCdnUserBillPrediction API synchronously
func (client *Client) DescribeCdnUserBillPrediction(request *DescribeCdnUserBillPredictionRequest) (response *DescribeCdnUserBillPredictionResponse, err error) {
	response = CreateDescribeCdnUserBillPredictionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnUserBillPredictionWithChan invokes the cdn.DescribeCdnUserBillPrediction API asynchronously
func (client *Client) DescribeCdnUserBillPredictionWithChan(request *DescribeCdnUserBillPredictionRequest) (<-chan *DescribeCdnUserBillPredictionResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnUserBillPredictionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnUserBillPrediction(request)
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

// DescribeCdnUserBillPredictionWithCallback invokes the cdn.DescribeCdnUserBillPrediction API asynchronously
func (client *Client) DescribeCdnUserBillPredictionWithCallback(request *DescribeCdnUserBillPredictionRequest, callback func(response *DescribeCdnUserBillPredictionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnUserBillPredictionResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnUserBillPrediction(request)
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

// DescribeCdnUserBillPredictionRequest is the request struct for api DescribeCdnUserBillPrediction
type DescribeCdnUserBillPredictionRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Dimension string           `position:"Query" name:"Dimension"`
	Area      string           `position:"Query" name:"Area"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCdnUserBillPredictionResponse is the response struct for api DescribeCdnUserBillPrediction
type DescribeCdnUserBillPredictionResponse struct {
	*responses.BaseResponse
	EndTime            string             `json:"EndTime" xml:"EndTime"`
	StartTime          string             `json:"StartTime" xml:"StartTime"`
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	BillType           string             `json:"BillType" xml:"BillType"`
	BillPredictionData BillPredictionData `json:"BillPredictionData" xml:"BillPredictionData"`
}

// CreateDescribeCdnUserBillPredictionRequest creates a request to invoke DescribeCdnUserBillPrediction API
func CreateDescribeCdnUserBillPredictionRequest() (request *DescribeCdnUserBillPredictionRequest) {
	request = &DescribeCdnUserBillPredictionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnUserBillPrediction", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnUserBillPredictionResponse creates a response to parse from DescribeCdnUserBillPrediction response
func CreateDescribeCdnUserBillPredictionResponse() (response *DescribeCdnUserBillPredictionResponse) {
	response = &DescribeCdnUserBillPredictionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
