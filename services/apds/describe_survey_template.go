package apds

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

// DescribeSurveyTemplate invokes the apds.DescribeSurveyTemplate API synchronously
func (client *Client) DescribeSurveyTemplate(request *DescribeSurveyTemplateRequest) (response *DescribeSurveyTemplateResponse, err error) {
	response = CreateDescribeSurveyTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSurveyTemplateWithChan invokes the apds.DescribeSurveyTemplate API asynchronously
func (client *Client) DescribeSurveyTemplateWithChan(request *DescribeSurveyTemplateRequest) (<-chan *DescribeSurveyTemplateResponse, <-chan error) {
	responseChan := make(chan *DescribeSurveyTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSurveyTemplate(request)
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

// DescribeSurveyTemplateWithCallback invokes the apds.DescribeSurveyTemplate API asynchronously
func (client *Client) DescribeSurveyTemplateWithCallback(request *DescribeSurveyTemplateRequest, callback func(response *DescribeSurveyTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSurveyTemplateResponse
		var err error
		defer close(result)
		response, err = client.DescribeSurveyTemplate(request)
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

// DescribeSurveyTemplateRequest is the request struct for api DescribeSurveyTemplate
type DescribeSurveyTemplateRequest struct {
	*requests.RoaRequest
	ResourceType string `position:"Query" name:"resourceType"`
}

// DescribeSurveyTemplateResponse is the response struct for api DescribeSurveyTemplate
type DescribeSurveyTemplateResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateDescribeSurveyTemplateRequest creates a request to invoke DescribeSurveyTemplate API
func CreateDescribeSurveyTemplateRequest() (request *DescribeSurveyTemplateRequest) {
	request = &DescribeSurveyTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "DescribeSurveyTemplate", "/okss-services/survey-template/list", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSurveyTemplateResponse creates a response to parse from DescribeSurveyTemplate response
func CreateDescribeSurveyTemplateResponse() (response *DescribeSurveyTemplateResponse) {
	response = &DescribeSurveyTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
