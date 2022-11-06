package polardb

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

// DescribeDBClusterBasicInfo invokes the polardb.DescribeDBClusterBasicInfo API synchronously
func (client *Client) DescribeDBClusterBasicInfo(request *DescribeDBClusterBasicInfoRequest) (response *DescribeDBClusterBasicInfoResponse, err error) {
	response = CreateDescribeDBClusterBasicInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterBasicInfoWithChan invokes the polardb.DescribeDBClusterBasicInfo API asynchronously
func (client *Client) DescribeDBClusterBasicInfoWithChan(request *DescribeDBClusterBasicInfoRequest) (<-chan *DescribeDBClusterBasicInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterBasicInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterBasicInfo(request)
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

// DescribeDBClusterBasicInfoWithCallback invokes the polardb.DescribeDBClusterBasicInfo API asynchronously
func (client *Client) DescribeDBClusterBasicInfoWithCallback(request *DescribeDBClusterBasicInfoRequest, callback func(response *DescribeDBClusterBasicInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterBasicInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterBasicInfo(request)
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

// DescribeDBClusterBasicInfoRequest is the request struct for api DescribeDBClusterBasicInfo
type DescribeDBClusterBasicInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterBasicInfoResponse is the response struct for api DescribeDBClusterBasicInfo
type DescribeDBClusterBasicInfoResponse struct {
	*responses.BaseResponse
	DeletionLock          int    `json:"DeletionLock" xml:"DeletionLock"`
	Category              string `json:"Category" xml:"Category"`
	ResourceGroupId       string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	DBClusterId           string `json:"DBClusterId" xml:"DBClusterId"`
	DBType                string `json:"DBType" xml:"DBType"`
	DBClusterNetworkType  string `json:"DBClusterNetworkType" xml:"DBClusterNetworkType"`
	StorageMax            int64  `json:"StorageMax" xml:"StorageMax"`
	DBVersion             string `json:"DBVersion" xml:"DBVersion"`
	MaintainTime          string `json:"MaintainTime" xml:"MaintainTime"`
	Engine                string `json:"Engine" xml:"Engine"`
	RequestId             string `json:"RequestId" xml:"RequestId"`
	VPCId                 string `json:"VPCId" xml:"VPCId"`
	DBClusterStatus       string `json:"DBClusterStatus" xml:"DBClusterStatus"`
	VSwitchId             string `json:"VSwitchId" xml:"VSwitchId"`
	DBClusterDescription  string `json:"DBClusterDescription" xml:"DBClusterDescription"`
	Expired               string `json:"Expired" xml:"Expired"`
	LockMode              string `json:"LockMode" xml:"LockMode"`
	DBVersionStatus       string `json:"DBVersionStatus" xml:"DBVersionStatus"`
	CreationTime          string `json:"CreationTime" xml:"CreationTime"`
	SQLSize               int64  `json:"SQLSize" xml:"SQLSize"`
	RegionId              string `json:"RegionId" xml:"RegionId"`
	ExpireTime            string `json:"ExpireTime" xml:"ExpireTime"`
	SubCategory           string `json:"SubCategory" xml:"SubCategory"`
	IsProxyLatestVersion  bool   `json:"IsProxyLatestVersion" xml:"IsProxyLatestVersion"`
	StorageType           string `json:"StorageType" xml:"StorageType"`
	ProxyCpuCores         string `json:"ProxyCpuCores" xml:"ProxyCpuCores"`
	ProxyStandardCpuCores string `json:"ProxyStandardCpuCores" xml:"ProxyStandardCpuCores"`
	ProxyType             string `json:"ProxyType" xml:"ProxyType"`
	ProxyStatus           string `json:"ProxyStatus" xml:"ProxyStatus"`
}

// CreateDescribeDBClusterBasicInfoRequest creates a request to invoke DescribeDBClusterBasicInfo API
func CreateDescribeDBClusterBasicInfoRequest() (request *DescribeDBClusterBasicInfoRequest) {
	request = &DescribeDBClusterBasicInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterBasicInfo", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterBasicInfoResponse creates a response to parse from DescribeDBClusterBasicInfo response
func CreateDescribeDBClusterBasicInfoResponse() (response *DescribeDBClusterBasicInfoResponse) {
	response = &DescribeDBClusterBasicInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
