package nlb

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

// HealthCheck is a nested struct in nlb response
type HealthCheck struct {
	HealthCheckEnabled        bool     `json:"HealthCheckEnabled" xml:"HealthCheckEnabled"`
	HealthCheckType           string   `json:"HealthCheckType" xml:"HealthCheckType"`
	HealthCheckConnectPort    int      `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthyThreshold          int      `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold        int      `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int      `json:"HealthCheckConnectTimeout" xml:"HealthCheckConnectTimeout"`
	HealthCheckInterval       int      `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthCheckDomain         string   `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	HealthCheckUrl            string   `json:"HealthCheckUrl" xml:"HealthCheckUrl"`
	HttpCheckMethod           string   `json:"HttpCheckMethod" xml:"HttpCheckMethod"`
	HealthCheckHttpCode       []string `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
}
