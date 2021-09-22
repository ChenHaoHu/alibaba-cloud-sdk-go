package green

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

// TextAsyncManualScan invokes the green.TextAsyncManualScan API synchronously
func (client *Client) TextAsyncManualScan(request *TextAsyncManualScanRequest) (response *TextAsyncManualScanResponse, err error) {
	response = CreateTextAsyncManualScanResponse()
	err = client.DoAction(request, response)
	return
}

// TextAsyncManualScanWithChan invokes the green.TextAsyncManualScan API asynchronously
func (client *Client) TextAsyncManualScanWithChan(request *TextAsyncManualScanRequest) (<-chan *TextAsyncManualScanResponse, <-chan error) {
	responseChan := make(chan *TextAsyncManualScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TextAsyncManualScan(request)
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

// TextAsyncManualScanWithCallback invokes the green.TextAsyncManualScan API asynchronously
func (client *Client) TextAsyncManualScanWithCallback(request *TextAsyncManualScanRequest, callback func(response *TextAsyncManualScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TextAsyncManualScanResponse
		var err error
		defer close(result)
		response, err = client.TextAsyncManualScan(request)
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

// TextAsyncManualScanRequest is the request struct for api TextAsyncManualScan
type TextAsyncManualScanRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// TextAsyncManualScanResponse is the response struct for api TextAsyncManualScan
type TextAsyncManualScanResponse struct {
	*responses.BaseResponse
}

// CreateTextAsyncManualScanRequest creates a request to invoke TextAsyncManualScan API
func CreateTextAsyncManualScanRequest() (request *TextAsyncManualScanRequest) {
	request = &TextAsyncManualScanRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "TextAsyncManualScan", "/green/text/manual/asyncScan", "", "")
	request.Method = requests.POST
	return
}

// CreateTextAsyncManualScanResponse creates a response to parse from TextAsyncManualScan response
func CreateTextAsyncManualScanResponse() (response *TextAsyncManualScanResponse) {
	response = &TextAsyncManualScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
