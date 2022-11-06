package ahas_openapi

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

// GetExperimentMeta invokes the ahas_openapi.GetExperimentMeta API synchronously
func (client *Client) GetExperimentMeta(request *GetExperimentMetaRequest) (response *GetExperimentMetaResponse, err error) {
	response = CreateGetExperimentMetaResponse()
	err = client.DoAction(request, response)
	return
}

// GetExperimentMetaWithChan invokes the ahas_openapi.GetExperimentMeta API asynchronously
func (client *Client) GetExperimentMetaWithChan(request *GetExperimentMetaRequest) (<-chan *GetExperimentMetaResponse, <-chan error) {
	responseChan := make(chan *GetExperimentMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetExperimentMeta(request)
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

// GetExperimentMetaWithCallback invokes the ahas_openapi.GetExperimentMeta API asynchronously
func (client *Client) GetExperimentMetaWithCallback(request *GetExperimentMetaRequest, callback func(response *GetExperimentMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetExperimentMetaResponse
		var err error
		defer close(result)
		response, err = client.GetExperimentMeta(request)
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

// GetExperimentMetaRequest is the request struct for api GetExperimentMeta
type GetExperimentMetaRequest struct {
	*requests.RpcRequest
	NameSpace    string `position:"Query" name:"NameSpace"`
	ExperimentId string `position:"Query" name:"ExperimentId"`
}

// GetExperimentMetaResponse is the response struct for api GetExperimentMeta
type GetExperimentMetaResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	Message      string   `json:"Message" xml:"Message"`
	State        string   `json:"State" xml:"State"`
	ExperimentId string   `json:"ExperimentId" xml:"ExperimentId"`
	CreateTime   string   `json:"CreateTime" xml:"CreateTime"`
	Code         string   `json:"Code" xml:"Code"`
	Success      bool     `json:"Success" xml:"Success"`
	Name         string   `json:"Name" xml:"Name"`
	Tags         []string `json:"Tags" xml:"Tags"`
}

// CreateGetExperimentMetaRequest creates a request to invoke GetExperimentMeta API
func CreateGetExperimentMetaRequest() (request *GetExperimentMetaRequest) {
	request = &GetExperimentMetaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "GetExperimentMeta", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetExperimentMetaResponse creates a response to parse from GetExperimentMeta response
func CreateGetExperimentMetaResponse() (response *GetExperimentMetaResponse) {
	response = &GetExperimentMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
