package dts

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

// DescribeSynchronizationJobAlert invokes the dts.DescribeSynchronizationJobAlert API synchronously
func (client *Client) DescribeSynchronizationJobAlert(request *DescribeSynchronizationJobAlertRequest) (response *DescribeSynchronizationJobAlertResponse, err error) {
	response = CreateDescribeSynchronizationJobAlertResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSynchronizationJobAlertWithChan invokes the dts.DescribeSynchronizationJobAlert API asynchronously
func (client *Client) DescribeSynchronizationJobAlertWithChan(request *DescribeSynchronizationJobAlertRequest) (<-chan *DescribeSynchronizationJobAlertResponse, <-chan error) {
	responseChan := make(chan *DescribeSynchronizationJobAlertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSynchronizationJobAlert(request)
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

// DescribeSynchronizationJobAlertWithCallback invokes the dts.DescribeSynchronizationJobAlert API asynchronously
func (client *Client) DescribeSynchronizationJobAlertWithCallback(request *DescribeSynchronizationJobAlertRequest, callback func(response *DescribeSynchronizationJobAlertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSynchronizationJobAlertResponse
		var err error
		defer close(result)
		response, err = client.DescribeSynchronizationJobAlert(request)
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

// DescribeSynchronizationJobAlertRequest is the request struct for api DescribeSynchronizationJobAlert
type DescribeSynchronizationJobAlertRequest struct {
	*requests.RpcRequest
	ClientToken              string `position:"Query" name:"ClientToken"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
	SynchronizationJobId     string `position:"Query" name:"SynchronizationJobId"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
}

// DescribeSynchronizationJobAlertResponse is the response struct for api DescribeSynchronizationJobAlert
type DescribeSynchronizationJobAlertResponse struct {
	*responses.BaseResponse
	SynchronizationJobName   string `json:"SynchronizationJobName" xml:"SynchronizationJobName"`
	ErrorAlertStatus         string `json:"ErrorAlertStatus" xml:"ErrorAlertStatus"`
	ErrCode                  string `json:"ErrCode" xml:"ErrCode"`
	Success                  string `json:"Success" xml:"Success"`
	ErrorAlertPhone          string `json:"ErrorAlertPhone" xml:"ErrorAlertPhone"`
	ErrMessage               string `json:"ErrMessage" xml:"ErrMessage"`
	DelayAlertStatus         string `json:"DelayAlertStatus" xml:"DelayAlertStatus"`
	DelayAlertPhone          string `json:"DelayAlertPhone" xml:"DelayAlertPhone"`
	DelayOverSeconds         string `json:"DelayOverSeconds" xml:"DelayOverSeconds"`
	RequestId                string `json:"RequestId" xml:"RequestId"`
	SynchronizationJobId     string `json:"SynchronizationJobId" xml:"SynchronizationJobId"`
	SynchronizationDirection string `json:"SynchronizationDirection" xml:"SynchronizationDirection"`
}

// CreateDescribeSynchronizationJobAlertRequest creates a request to invoke DescribeSynchronizationJobAlert API
func CreateDescribeSynchronizationJobAlertRequest() (request *DescribeSynchronizationJobAlertRequest) {
	request = &DescribeSynchronizationJobAlertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2019-09-01", "DescribeSynchronizationJobAlert", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSynchronizationJobAlertResponse creates a response to parse from DescribeSynchronizationJobAlert response
func CreateDescribeSynchronizationJobAlertResponse() (response *DescribeSynchronizationJobAlertResponse) {
	response = &DescribeSynchronizationJobAlertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
