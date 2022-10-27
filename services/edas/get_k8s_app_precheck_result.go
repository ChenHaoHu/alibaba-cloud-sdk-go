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

// GetK8sAppPrecheckResult invokes the edas.GetK8sAppPrecheckResult API synchronously
func (client *Client) GetK8sAppPrecheckResult(request *GetK8sAppPrecheckResultRequest) (response *GetK8sAppPrecheckResultResponse, err error) {
	response = CreateGetK8sAppPrecheckResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetK8sAppPrecheckResultWithChan invokes the edas.GetK8sAppPrecheckResult API asynchronously
func (client *Client) GetK8sAppPrecheckResultWithChan(request *GetK8sAppPrecheckResultRequest) (<-chan *GetK8sAppPrecheckResultResponse, <-chan error) {
	responseChan := make(chan *GetK8sAppPrecheckResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetK8sAppPrecheckResult(request)
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

// GetK8sAppPrecheckResultWithCallback invokes the edas.GetK8sAppPrecheckResult API asynchronously
func (client *Client) GetK8sAppPrecheckResultWithCallback(request *GetK8sAppPrecheckResultRequest, callback func(response *GetK8sAppPrecheckResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetK8sAppPrecheckResultResponse
		var err error
		defer close(result)
		response, err = client.GetK8sAppPrecheckResult(request)
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

// GetK8sAppPrecheckResultRequest is the request struct for api GetK8sAppPrecheckResult
type GetK8sAppPrecheckResultRequest struct {
	*requests.RoaRequest
	AppName   string `position:"Query" name:"AppName"`
	Namespace string `position:"Query" name:"Namespace"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

// GetK8sAppPrecheckResultResponse is the response struct for api GetK8sAppPrecheckResult
type GetK8sAppPrecheckResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetK8sAppPrecheckResultRequest creates a request to invoke GetK8sAppPrecheckResult API
func CreateGetK8sAppPrecheckResultRequest() (request *GetK8sAppPrecheckResultRequest) {
	request = &GetK8sAppPrecheckResultRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetK8sAppPrecheckResult", "/pop/v5/k8s/app_precheck", "edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetK8sAppPrecheckResultResponse creates a response to parse from GetK8sAppPrecheckResult response
func CreateGetK8sAppPrecheckResultResponse() (response *GetK8sAppPrecheckResultResponse) {
	response = &GetK8sAppPrecheckResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
