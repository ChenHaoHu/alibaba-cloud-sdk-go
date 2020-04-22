package cassandra

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

// DataCenter is a nested struct in cassandra response
type DataCenter struct {
	ClusterId         string `json:"ClusterId" xml:"ClusterId"`
	DiskType          string `json:"DiskType" xml:"DiskType"`
	PayType           string `json:"PayType" xml:"PayType"`
	InstanceType      string `json:"InstanceType" xml:"InstanceType"`
	DataCenterName    string `json:"DataCenterName" xml:"DataCenterName"`
	NodeCount         int    `json:"NodeCount" xml:"NodeCount"`
	RegionId          string `json:"RegionId" xml:"RegionId"`
	CreatedTime       string `json:"CreatedTime" xml:"CreatedTime"`
	ZoneId            string `json:"ZoneId" xml:"ZoneId"`
	VswitchId         string `json:"VswitchId" xml:"VswitchId"`
	DiskSize          int    `json:"DiskSize" xml:"DiskSize"`
	ExpireTime        string `json:"ExpireTime" xml:"ExpireTime"`
	CommodityInstance string `json:"CommodityInstance" xml:"CommodityInstance"`
	Status            string `json:"Status" xml:"Status"`
	VpcId             string `json:"VpcId" xml:"VpcId"`
	LockMode          string `json:"LockMode" xml:"LockMode"`
	DataCenterId      string `json:"DataCenterId" xml:"DataCenterId"`
	Nodes             Nodes  `json:"Nodes" xml:"Nodes"`
}
