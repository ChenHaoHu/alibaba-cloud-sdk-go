package oos

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

// OpsItem is a nested struct in oos response
type OpsItem struct {
	Category        string                 `json:"Category" xml:"Category"`
	Resources       string                 `json:"Resources" xml:"Resources"`
	Title           string                 `json:"Title" xml:"Title"`
	Attributes      string                 `json:"Attributes" xml:"Attributes"`
	Priority        string                 `json:"Priority" xml:"Priority"`
	Source          string                 `json:"Source" xml:"Source"`
	CreatedBy       string                 `json:"CreatedBy" xml:"CreatedBy"`
	Solutions       string                 `json:"Solutions" xml:"Solutions"`
	CreateDate      string                 `json:"CreateDate" xml:"CreateDate"`
	ResourceGroupId string                 `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Severity        string                 `json:"Severity" xml:"Severity"`
	UpdateDate      string                 `json:"UpdateDate" xml:"UpdateDate"`
	LastModifiedBy  string                 `json:"LastModifiedBy" xml:"LastModifiedBy"`
	Status          string                 `json:"Status" xml:"Status"`
	OpsItemId       string                 `json:"OpsItemId" xml:"OpsItemId"`
	Description     string                 `json:"Description" xml:"Description"`
	Tags            map[string]interface{} `json:"Tags" xml:"Tags"`
}
