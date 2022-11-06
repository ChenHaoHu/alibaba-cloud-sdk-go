package live

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

// StreamData is a nested struct in live response
type StreamData struct {
	P2pTraffic  int64   `json:"P2pTraffic" xml:"P2pTraffic"`
	RtmpBps     float64 `json:"RtmpBps" xml:"RtmpBps"`
	HlsBps      float64 `json:"HlsBps" xml:"HlsBps"`
	P2pBps      float64 `json:"P2pBps" xml:"P2pBps"`
	StreamName  string  `json:"StreamName" xml:"StreamName"`
	P2pCount    int64   `json:"P2pCount" xml:"P2pCount"`
	Traffic     int64   `json:"Traffic" xml:"Traffic"`
	RtmpCount   int64   `json:"RtmpCount" xml:"RtmpCount"`
	RtsCount    int64   `json:"RtsCount" xml:"RtsCount"`
	RtsBps      float64 `json:"RtsBps" xml:"RtsBps"`
	Bps         float64 `json:"Bps" xml:"Bps"`
	HlsTraffic  int64   `json:"HlsTraffic" xml:"HlsTraffic"`
	FlvBps      float64 `json:"FlvBps" xml:"FlvBps"`
	Count       int64   `json:"Count" xml:"Count"`
	RtsTraffic  int64   `json:"RtsTraffic" xml:"RtsTraffic"`
	RtmpTraffic int64   `json:"RtmpTraffic" xml:"RtmpTraffic"`
	HlsCount    int64   `json:"HlsCount" xml:"HlsCount"`
	AppName     string  `json:"AppName" xml:"AppName"`
	FlvCount    int64   `json:"FlvCount" xml:"FlvCount"`
	FlvTraffic  int64   `json:"FlvTraffic" xml:"FlvTraffic"`
	TimeStamp   string  `json:"TimeStamp" xml:"TimeStamp"`
}
