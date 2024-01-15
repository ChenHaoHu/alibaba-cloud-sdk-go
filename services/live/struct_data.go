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

// Data is a nested struct in live response
type Data struct {
	Roles               int      `json:"Roles" xml:"Roles"`
	ModifyTime          string   `json:"ModifyTime" xml:"ModifyTime"`
	V20012              string   `json:"V20012" xml:"V20012"`
	VvSuccess           string   `json:"VvSuccess" xml:"VvSuccess"`
	CreateTime          string   `json:"CreateTime" xml:"CreateTime"`
	CallbackUrl         string   `json:"CallbackUrl" xml:"CallbackUrl"`
	FirstFrameComplete  string   `json:"FirstFrameComplete" xml:"FirstFrameComplete"`
	ChannelId           string   `json:"ChannelId" xml:"ChannelId"`
	TimeStamp           string   `json:"TimeStamp" xml:"TimeStamp"`
	FrameDelay          string   `json:"FrameDelay" xml:"FrameDelay"`
	PlayTime            string   `json:"PlayTime" xml:"PlayTime"`
	FinishGetStreamInfo string   `json:"FinishGetStreamInfo" xml:"FinishGetStreamInfo"`
	Connected           string   `json:"Connected" xml:"Connected"`
	V20002              string   `json:"V20002" xml:"V20002"`
	FirstPacket         string   `json:"FirstPacket" xml:"FirstPacket"`
	V20001              string   `json:"V20001" xml:"V20001"`
	Initialized         string   `json:"Initialized" xml:"Initialized"`
	V20011              string   `json:"V20011" xml:"V20011"`
	StallTime           string   `json:"StallTime" xml:"StallTime"`
	Message             string   `json:"message" xml:"message"`
	SubId               string   `json:"SubId" xml:"SubId"`
	Status              string   `json:"status" xml:"status"`
	V20052              string   `json:"V20052" xml:"V20052"`
	VvTotal             string   `json:"VvTotal" xml:"VvTotal"`
	V20013              string   `json:"V20013" xml:"V20013"`
	Users               []string `json:"Users" xml:"Users"`
	Events              []string `json:"Events" xml:"Events"`
}
