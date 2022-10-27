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

// UpdateConfigTemplate invokes the edas.UpdateConfigTemplate API synchronously
func (client *Client) UpdateConfigTemplate(request *UpdateConfigTemplateRequest) (response *UpdateConfigTemplateResponse, err error) {
	response = CreateUpdateConfigTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateConfigTemplateWithChan invokes the edas.UpdateConfigTemplate API asynchronously
func (client *Client) UpdateConfigTemplateWithChan(request *UpdateConfigTemplateRequest) (<-chan *UpdateConfigTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateConfigTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateConfigTemplate(request)
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

// UpdateConfigTemplateWithCallback invokes the edas.UpdateConfigTemplate API asynchronously
func (client *Client) UpdateConfigTemplateWithCallback(request *UpdateConfigTemplateRequest, callback func(response *UpdateConfigTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateConfigTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateConfigTemplate(request)
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

// UpdateConfigTemplateRequest is the request struct for api UpdateConfigTemplate
type UpdateConfigTemplateRequest struct {
	*requests.RoaRequest
	Name        string `position:"Body" name:"Name"`
	Format      string `position:"Body" name:"Format"`
	Description string `position:"Body" name:"Description"`
	Id          string `position:"Body" name:"Id"`
	Content     string `position:"Body" name:"Content"`
}

// UpdateConfigTemplateResponse is the response struct for api UpdateConfigTemplate
type UpdateConfigTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateConfigTemplateRequest creates a request to invoke UpdateConfigTemplate API
func CreateUpdateConfigTemplateRequest() (request *UpdateConfigTemplateRequest) {
	request = &UpdateConfigTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateConfigTemplate", "/pop/v5/config_template", "edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateConfigTemplateResponse creates a response to parse from UpdateConfigTemplate response
func CreateUpdateConfigTemplateResponse() (response *UpdateConfigTemplateResponse) {
	response = &UpdateConfigTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
