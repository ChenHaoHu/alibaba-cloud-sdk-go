package antiddos_public

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

// DescribeDdosCredit invokes the antiddos_public.DescribeDdosCredit API synchronously
func (client *Client) DescribeDdosCredit(request *DescribeDdosCreditRequest) (response *DescribeDdosCreditResponse, err error) {
	response = CreateDescribeDdosCreditResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDdosCreditWithChan invokes the antiddos_public.DescribeDdosCredit API asynchronously
func (client *Client) DescribeDdosCreditWithChan(request *DescribeDdosCreditRequest) (<-chan *DescribeDdosCreditResponse, <-chan error) {
	responseChan := make(chan *DescribeDdosCreditResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDdosCredit(request)
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

// DescribeDdosCreditWithCallback invokes the antiddos_public.DescribeDdosCredit API asynchronously
func (client *Client) DescribeDdosCreditWithCallback(request *DescribeDdosCreditRequest, callback func(response *DescribeDdosCreditResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDdosCreditResponse
		var err error
		defer close(result)
		response, err = client.DescribeDdosCredit(request)
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

// DescribeDdosCreditRequest is the request struct for api DescribeDdosCredit
type DescribeDdosCreditRequest struct {
	*requests.RpcRequest
	SourceIp     string `position:"Query" name:"SourceIp"`
	DdosRegionId string `position:"Query" name:"DdosRegionId"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeDdosCreditResponse is the response struct for api DescribeDdosCredit
type DescribeDdosCreditResponse struct {
	*responses.BaseResponse
	Success    bool       `json:"Success" xml:"Success"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	DdosCredit DdosCredit `json:"DdosCredit" xml:"DdosCredit"`
}

// CreateDescribeDdosCreditRequest creates a request to invoke DescribeDdosCredit API
func CreateDescribeDdosCreditRequest() (request *DescribeDdosCreditRequest) {
	request = &DescribeDdosCreditRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("antiddos-public", "2017-05-18", "DescribeDdosCredit", "ddosbasic", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDdosCreditResponse creates a response to parse from DescribeDdosCredit response
func CreateDescribeDdosCreditResponse() (response *DescribeDdosCreditResponse) {
	response = &DescribeDdosCreditResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
