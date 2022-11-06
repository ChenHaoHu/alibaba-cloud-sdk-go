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

// DescribeDcdnReport invokes the dcdn.DescribeDcdnReport API synchronously
func (client *Client) DescribeDcdnReport(request *DescribeDcdnReportRequest) (response *DescribeDcdnReportResponse, err error) {
	response = CreateDescribeDcdnReportResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnReportWithChan invokes the dcdn.DescribeDcdnReport API asynchronously
func (client *Client) DescribeDcdnReportWithChan(request *DescribeDcdnReportRequest) (<-chan *DescribeDcdnReportResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnReport(request)
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

// DescribeDcdnReportWithCallback invokes the dcdn.DescribeDcdnReport API asynchronously
func (client *Client) DescribeDcdnReportWithCallback(request *DescribeDcdnReportRequest, callback func(response *DescribeDcdnReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnReportResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnReport(request)
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

// DescribeDcdnReportRequest is the request struct for api DescribeDcdnReport
type DescribeDcdnReportRequest struct {
	*requests.RpcRequest
	ReportId   requests.Integer `position:"Query" name:"ReportId"`
	StartTime  string           `position:"Query" name:"StartTime"`
	Area       string           `position:"Query" name:"Area"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	HttpCode   string           `position:"Query" name:"HttpCode"`
	IsOverseas string           `position:"Query" name:"IsOverseas"`
}

// DescribeDcdnReportResponse is the response struct for api DescribeDcdnReport
type DescribeDcdnReportResponse struct {
	*responses.BaseResponse
	Content   map[string]interface{} `json:"Content" xml:"Content"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeDcdnReportRequest creates a request to invoke DescribeDcdnReport API
func CreateDescribeDcdnReportRequest() (request *DescribeDcdnReportRequest) {
	request = &DescribeDcdnReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnReport", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnReportResponse creates a response to parse from DescribeDcdnReport response
func CreateDescribeDcdnReportResponse() (response *DescribeDcdnReportResponse) {
	response = &DescribeDcdnReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
