package config

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

// DiscoveredResourceProfile is a nested struct in config response
type DiscoveredResourceProfile struct {
	Region               string `json:"Region" xml:"Region"`
	ResourceType         string `json:"ResourceType" xml:"ResourceType"`
	ResourceId           string `json:"ResourceId" xml:"ResourceId"`
	ResourceStatus       string `json:"ResourceStatus" xml:"ResourceStatus"`
	Tags                 string `json:"Tags" xml:"Tags"`
	Version              int64  `json:"Version" xml:"Version"`
	AccountId            int64  `json:"AccountId" xml:"AccountId"`
	ResourceName         string `json:"ResourceName" xml:"ResourceName"`
	ResourceOwnerId      int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	ResourceCreationTime int64  `json:"ResourceCreationTime" xml:"ResourceCreationTime"`
	ResourceDeleted      int    `json:"ResourceDeleted" xml:"ResourceDeleted"`
}
