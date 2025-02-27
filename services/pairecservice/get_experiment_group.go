package pairecservice

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

// GetExperimentGroup invokes the pairecservice.GetExperimentGroup API synchronously
func (client *Client) GetExperimentGroup(request *GetExperimentGroupRequest) (response *GetExperimentGroupResponse, err error) {
	response = CreateGetExperimentGroupResponse()
	err = client.DoAction(request, response)
	return
}

// GetExperimentGroupWithChan invokes the pairecservice.GetExperimentGroup API asynchronously
func (client *Client) GetExperimentGroupWithChan(request *GetExperimentGroupRequest) (<-chan *GetExperimentGroupResponse, <-chan error) {
	responseChan := make(chan *GetExperimentGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetExperimentGroup(request)
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

// GetExperimentGroupWithCallback invokes the pairecservice.GetExperimentGroup API asynchronously
func (client *Client) GetExperimentGroupWithCallback(request *GetExperimentGroupRequest, callback func(response *GetExperimentGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetExperimentGroupResponse
		var err error
		defer close(result)
		response, err = client.GetExperimentGroup(request)
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

// GetExperimentGroupRequest is the request struct for api GetExperimentGroup
type GetExperimentGroupRequest struct {
	*requests.RoaRequest
	InstanceId        string `position:"Query" name:"InstanceId"`
	ExperimentGroupId string `position:"Path" name:"ExperimentGroupId"`
}

// GetExperimentGroupResponse is the response struct for api GetExperimentGroup
type GetExperimentGroupResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	LayerId                  string `json:"LayerId" xml:"LayerId"`
	LaboratoryId             string `json:"LaboratoryId" xml:"LaboratoryId"`
	SceneId                  string `json:"SceneId" xml:"SceneId"`
	Name                     string `json:"Name" xml:"Name"`
	Description              string `json:"Description" xml:"Description"`
	DebugUsers               string `json:"DebugUsers" xml:"DebugUsers"`
	DebugCrowdId             string `json:"DebugCrowdId" xml:"DebugCrowdId"`
	Owner                    string `json:"Owner" xml:"Owner"`
	NeedAA                   bool   `json:"NeedAA" xml:"NeedAA"`
	Filter                   string `json:"Filter" xml:"Filter"`
	DistributionType         string `json:"DistributionType" xml:"DistributionType"`
	DistributionTimeDuration int    `json:"DistributionTimeDuration" xml:"DistributionTimeDuration"`
	CrowdId                  string `json:"CrowdId" xml:"CrowdId"`
	Config                   string `json:"Config" xml:"Config"`
	ReservedBuckets          string `json:"ReservedBuckets" xml:"ReservedBuckets"`
	Status                   string `json:"Status" xml:"Status"`
}

// CreateGetExperimentGroupRequest creates a request to invoke GetExperimentGroup API
func CreateGetExperimentGroupRequest() (request *GetExperimentGroupRequest) {
	request = &GetExperimentGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "GetExperimentGroup", "/api/v1/experimentgroups/[ExperimentGroupId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetExperimentGroupResponse creates a response to parse from GetExperimentGroup response
func CreateGetExperimentGroupResponse() (response *GetExperimentGroupResponse) {
	response = &GetExperimentGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
