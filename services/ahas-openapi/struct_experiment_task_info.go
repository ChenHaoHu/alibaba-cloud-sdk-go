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

// ExperimentTaskInfo is a nested struct in ahas_openapi response
type ExperimentTaskInfo struct {
	StartTime      int64   `json:"StartTime" xml:"StartTime"`
	Message        string  `json:"Message" xml:"Message"`
	State          string  `json:"State" xml:"State"`
	TaskId         string  `json:"TaskId" xml:"TaskId"`
	Result         string  `json:"Result" xml:"Result"`
	ExperimentId   string  `json:"ExperimentId" xml:"ExperimentId"`
	ExperimentName string  `json:"ExperimentName" xml:"ExperimentName"`
	Namespace      string  `json:"Namespace" xml:"Namespace"`
	CurrentPhase   string  `json:"CurrentPhase" xml:"CurrentPhase"`
	EndTime        int64   `json:"EndTime" xml:"EndTime"`
	Creator        Creator `json:"Creator" xml:"Creator"`
	ExtInfo        ExtInfo `json:"ExtInfo" xml:"ExtInfo"`
}
