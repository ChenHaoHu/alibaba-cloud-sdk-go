package scdn

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

// CertInfo is a nested struct in scdn response
type CertInfo struct {
	Status         string `json:"Status" xml:"Status"`
	CertLife       string `json:"CertLife" xml:"CertLife"`
	CertExpireTime string `json:"CertExpireTime" xml:"CertExpireTime"`
	SSLPub         string `json:"SSLPub" xml:"SSLPub"`
	SSLProtocol    string `json:"SSLProtocol" xml:"SSLProtocol"`
	CertType       string `json:"CertType" xml:"CertType"`
	CertDomainName string `json:"CertDomainName" xml:"CertDomainName"`
	CertName       string `json:"CertName" xml:"CertName"`
	CertOrg        string `json:"CertOrg" xml:"CertOrg"`
	DomainName     string `json:"DomainName" xml:"DomainName"`
}
