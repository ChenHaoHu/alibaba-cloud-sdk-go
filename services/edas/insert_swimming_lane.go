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

// InsertSwimmingLane invokes the edas.InsertSwimmingLane API synchronously
func (client *Client) InsertSwimmingLane(request *InsertSwimmingLaneRequest) (response *InsertSwimmingLaneResponse, err error) {
	response = CreateInsertSwimmingLaneResponse()
	err = client.DoAction(request, response)
	return
}

// InsertSwimmingLaneWithChan invokes the edas.InsertSwimmingLane API asynchronously
func (client *Client) InsertSwimmingLaneWithChan(request *InsertSwimmingLaneRequest) (<-chan *InsertSwimmingLaneResponse, <-chan error) {
	responseChan := make(chan *InsertSwimmingLaneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertSwimmingLane(request)
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

// InsertSwimmingLaneWithCallback invokes the edas.InsertSwimmingLane API asynchronously
func (client *Client) InsertSwimmingLaneWithCallback(request *InsertSwimmingLaneRequest, callback func(response *InsertSwimmingLaneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertSwimmingLaneResponse
		var err error
		defer close(result)
		response, err = client.InsertSwimmingLane(request)
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

// InsertSwimmingLaneRequest is the request struct for api InsertSwimmingLane
type InsertSwimmingLaneRequest struct {
	*requests.RoaRequest
	AppInfos        string `position:"Query" name:"AppInfos"`
	EntryRules      string `position:"Query" name:"EntryRules"`
	LogicalRegionId string `position:"Query" name:"LogicalRegionId"`
	EnableRules     string `position:"Query" name:"EnableRules"`
	GroupId         string `position:"Query" name:"GroupId"`
	Name            string `position:"Query" name:"Name"`
	Tag             string `position:"Query" name:"Tag"`
}

// InsertSwimmingLaneResponse is the response struct for api InsertSwimmingLane
type InsertSwimmingLaneResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateInsertSwimmingLaneRequest creates a request to invoke InsertSwimmingLane API
func CreateInsertSwimmingLaneRequest() (request *InsertSwimmingLaneRequest) {
	request = &InsertSwimmingLaneRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "InsertSwimmingLane", "/pop/v5/trafficmgnt/swimming_lanes", "edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInsertSwimmingLaneResponse creates a response to parse from InsertSwimmingLane response
func CreateInsertSwimmingLaneResponse() (response *InsertSwimmingLaneResponse) {
	response = &InsertSwimmingLaneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
