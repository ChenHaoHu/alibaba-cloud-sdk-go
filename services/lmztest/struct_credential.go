package lmztest

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

// Credential is a nested struct in lmztest response
type Credential struct {
	KeyPrefix           string `json:"KeyPrefix" xml:"KeyPrefix"`
	SecurityToken       string `json:"SecurityToken" xml:"SecurityToken"`
	AccessKeySecret     string `json:"AccessKeySecret" xml:"AccessKeySecret"`
	Expiration          string `json:"Expiration" xml:"Expiration"`
	AccessKeyId         string `json:"AccessKeyId" xml:"AccessKeyId"`
	Bucket              string `json:"Bucket" xml:"Bucket"`
	RegionId            string `json:"RegionId" xml:"RegionId"`
	OssPublicEndpoint   string `json:"OssPublicEndpoint" xml:"OssPublicEndpoint"`
	OssInternalEndpoint string `json:"OssInternalEndpoint" xml:"OssInternalEndpoint"`
	OssVpcEndpoint      string `json:"OssVpcEndpoint" xml:"OssVpcEndpoint"`
}
