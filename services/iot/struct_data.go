package iot

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

// Data is a nested struct in iot response
type Data struct {
	IsEnable                  string                                          `json:"IsEnable" xml:"IsEnable"`
	LatestDeploymentStatus    int                                             `json:"LatestDeploymentStatus" xml:"LatestDeploymentStatus"`
	SourceData                string                                          `json:"SourceData" xml:"SourceData"`
	RoleAttachTime            string                                          `json:"RoleAttachTime" xml:"RoleAttachTime"`
	RequestProtocol           string                                          `json:"RequestProtocol" xml:"RequestProtocol"`
	ContainerConfig           string                                          `json:"ContainerConfig" xml:"ContainerConfig"`
	SuccessSum                int64                                           `json:"SuccessSum" xml:"SuccessSum"`
	RoleName                  string                                          `json:"RoleName" xml:"RoleName"`
	Spec                      int                                             `json:"Spec" xml:"Spec"`
	RequestMethod             string                                          `json:"RequestMethod" xml:"RequestMethod"`
	Nickname                  string                                          `json:"Nickname" xml:"Nickname"`
	DevEui                    string                                          `json:"DevEui" xml:"DevEui"`
	ProjectId                 string                                          `json:"ProjectId" xml:"ProjectId"`
	IsBeian                   string                                          `json:"IsBeian" xml:"IsBeian"`
	GroupId                   string                                          `json:"GroupId" xml:"GroupId"`
	TunnelId                  string                                          `json:"TunnelId" xml:"TunnelId"`
	LatestDeploymentType      string                                          `json:"LatestDeploymentType" xml:"LatestDeploymentType"`
	CoordinateSystem          int                                             `json:"CoordinateSystem" xml:"CoordinateSystem"`
	OssPreSignedAddress       string                                          `json:"OssPreSignedAddress" xml:"OssPreSignedAddress"`
	Type                      string                                          `json:"Type" xml:"Type"`
	DeviceConnState           string                                          `json:"DeviceConnState" xml:"DeviceConnState"`
	FileId                    string                                          `json:"FileId" xml:"FileId"`
	ThingModelJson            string                                          `json:"ThingModelJson" xml:"ThingModelJson"`
	LastUpdateTime            int64                                           `json:"LastUpdateTime" xml:"LastUpdateTime"`
	Versions                  string                                          `json:"Versions" xml:"Versions"`
	TotalCount                int64                                           `json:"TotalCount" xml:"TotalCount"`
	TslUri                    string                                          `json:"TslUri" xml:"TslUri"`
	Code                      string                                          `json:"Code" xml:"Code"`
	CsvUrl                    string                                          `json:"CsvUrl" xml:"CsvUrl"`
	OssAddress                string                                          `json:"OssAddress" xml:"OssAddress"`
	FirmwareId                string                                          `json:"FirmwareId" xml:"FirmwareId"`
	AliyunCommodityCode       string                                          `json:"AliyunCommodityCode" xml:"AliyunCommodityCode"`
	OssAccessKeyId            string                                          `json:"OssAccessKeyId" xml:"OssAccessKeyId"`
	ApplyId                   int64                                           `json:"ApplyId" xml:"ApplyId"`
	UtcCreated                string                                          `json:"UtcCreated" xml:"UtcCreated"`
	LongJobId                 string                                          `json:"LongJobId" xml:"LongJobId"`
	Host                      string                                          `json:"Host" xml:"Host"`
	DeviceName                string                                          `json:"DeviceName" xml:"DeviceName"`
	PageCount                 int64                                           `json:"PageCount" xml:"PageCount"`
	Count                     int64                                           `json:"Count" xml:"Count"`
	Size                      string                                          `json:"Size" xml:"Size"`
	Udi                       string                                          `json:"Udi" xml:"Udi"`
	CsvFileName               string                                          `json:"CsvFileName" xml:"CsvFileName"`
	IsOpen                    bool                                            `json:"IsOpen" xml:"IsOpen"`
	SourceURI                 string                                          `json:"SourceURI" xml:"SourceURI"`
	TotalSize                 int                                             `json:"TotalSize" xml:"TotalSize"`
	ProductName               string                                          `json:"ProductName" xml:"ProductName"`
	Name                      string                                          `json:"Name" xml:"Name"`
	DatasetId                 string                                          `json:"DatasetId" xml:"DatasetId"`
	GmtCreateTimestamp        int64                                           `json:"GmtCreateTimestamp" xml:"GmtCreateTimestamp"`
	SpeechCode                string                                          `json:"SpeechCode" xml:"SpeechCode"`
	SpeechType                string                                          `json:"SpeechType" xml:"SpeechType"`
	DownloadUrl               string                                          `json:"DownloadUrl" xml:"DownloadUrl"`
	PageSize                  int                                             `json:"PageSize" xml:"PageSize"`
	Key                       string                                          `json:"Key" xml:"Key"`
	ResultDataInString        string                                          `json:"ResultDataInString" xml:"ResultDataInString"`
	GmtCreate                 string                                          `json:"GmtCreate" xml:"GmtCreate"`
	SourceType                string                                          `json:"SourceType" xml:"SourceType"`
	EnableSoundCode           bool                                            `json:"EnableSoundCode" xml:"EnableSoundCode"`
	ResultJson                string                                          `json:"ResultJson" xml:"ResultJson"`
	InstanceId                string                                          `json:"InstanceId" xml:"InstanceId"`
	SourceConnState           string                                          `json:"SourceConnState" xml:"SourceConnState"`
	UtcClosed                 string                                          `json:"UtcClosed" xml:"UtcClosed"`
	Voice                     string                                          `json:"Voice" xml:"Voice"`
	Policy                    string                                          `json:"Policy" xml:"Policy"`
	DateFormat                string                                          `json:"DateFormat" xml:"DateFormat"`
	Volume                    int                                             `json:"Volume" xml:"Volume"`
	TenantId                  string                                          `json:"TenantId" xml:"TenantId"`
	DeviceOnline              int                                             `json:"DeviceOnline" xml:"DeviceOnline"`
	JobId                     string                                          `json:"JobId" xml:"JobId"`
	Id                        int                                             `json:"Id" xml:"Id"`
	ProductSecret             string                                          `json:"ProductSecret" xml:"ProductSecret"`
	Ip                        string                                          `json:"Ip" xml:"Ip"`
	ScriptType                string                                          `json:"ScriptType" xml:"ScriptType"`
	PageNum                   int                                             `json:"PageNum" xml:"PageNum"`
	AvailableQuota            int                                             `json:"AvailableQuota" xml:"AvailableQuota"`
	AccessKeyId               string                                          `json:"AccessKeyId" xml:"AccessKeyId"`
	EdgeVersion               string                                          `json:"EdgeVersion" xml:"EdgeVersion"`
	JoinEui                   string                                          `json:"JoinEui" xml:"JoinEui"`
	DriverId                  string                                          `json:"DriverId" xml:"DriverId"`
	CurrentPage               int                                             `json:"CurrentPage" xml:"CurrentPage"`
	Signature                 string                                          `json:"Signature" xml:"Signature"`
	Adcode                    int64                                           `json:"Adcode" xml:"Adcode"`
	ObjectStorage             string                                          `json:"ObjectStorage" xml:"ObjectStorage"`
	DeviceCount               int                                             `json:"DeviceCount" xml:"DeviceCount"`
	ProtocolType              string                                          `json:"ProtocolType" xml:"ProtocolType"`
	AuthType                  string                                          `json:"AuthType" xml:"AuthType"`
	SourceAccessToken         string                                          `json:"SourceAccessToken" xml:"SourceAccessToken"`
	Total                     int64                                           `json:"Total" xml:"Total"`
	DataFormat                int                                             `json:"DataFormat" xml:"DataFormat"`
	BizEnable                 bool                                            `json:"BizEnable" xml:"BizEnable"`
	GmtCompleted              string                                          `json:"GmtCompleted" xml:"GmtCompleted"`
	ProgressId                string                                          `json:"ProgressId" xml:"ProgressId"`
	UtcCreate                 string                                          `json:"UtcCreate" xml:"UtcCreate"`
	Latitude                  float64                                         `json:"Latitude" xml:"Latitude"`
	EndTime                   int64                                           `json:"EndTime" xml:"EndTime"`
	HasNext                   bool                                            `json:"HasNext" xml:"HasNext"`
	TargetType                string                                          `json:"TargetType" xml:"TargetType"`
	DeviceActive              int                                             `json:"DeviceActive" xml:"DeviceActive"`
	DriverVersion             string                                          `json:"DriverVersion" xml:"DriverVersion"`
	BizType                   string                                          `json:"BizType" xml:"BizType"`
	Password                  string                                          `json:"Password" xml:"Password"`
	PageNo                    int                                             `json:"PageNo" xml:"PageNo"`
	GmtCompletedTimestamp     int64                                           `json:"GmtCompletedTimestamp" xml:"GmtCompletedTimestamp"`
	Progress                  int                                             `json:"Progress" xml:"Progress"`
	PreviewSize               int                                             `json:"PreviewSize" xml:"PreviewSize"`
	DriverConfig              string                                          `json:"DriverConfig" xml:"DriverConfig"`
	BizCode                   string                                          `json:"BizCode" xml:"BizCode"`
	RoleArn                   string                                          `json:"RoleArn" xml:"RoleArn"`
	Token                     string                                          `json:"Token" xml:"Token"`
	Tags                      string                                          `json:"Tags" xml:"Tags"`
	Configuration             string                                          `json:"Configuration" xml:"Configuration"`
	ResultCsvFile             string                                          `json:"ResultCsvFile" xml:"ResultCsvFile"`
	AuthMode                  int                                             `json:"AuthMode" xml:"AuthMode"`
	UtcCreatedOn              string                                          `json:"UtcCreatedOn" xml:"UtcCreatedOn"`
	ExpiredQuota              int                                             `json:"ExpiredQuota" xml:"ExpiredQuota"`
	Longitude                 float64                                         `json:"Longitude" xml:"Longitude"`
	Province                  string                                          `json:"Province" xml:"Province"`
	AppId                     string                                          `json:"AppId" xml:"AppId"`
	MessageId                 string                                          `json:"MessageId" xml:"MessageId"`
	FileUrl                   string                                          `json:"FileUrl" xml:"FileUrl"`
	BeginTime                 int64                                           `json:"BeginTime" xml:"BeginTime"`
	Id2                       bool                                            `json:"Id2" xml:"Id2"`
	NodeType                  int                                             `json:"NodeType" xml:"NodeType"`
	ConfigCheckRule           string                                          `json:"ConfigCheckRule" xml:"ConfigCheckRule"`
	TslStr                    string                                          `json:"TslStr" xml:"TslStr"`
	ApiSrn                    string                                          `json:"ApiSrn" xml:"ApiSrn"`
	ScriptUrl                 string                                          `json:"ScriptUrl" xml:"ScriptUrl"`
	OSSAccessKeyId            string                                          `json:"OSSAccessKeyId" xml:"OSSAccessKeyId"`
	Text                      string                                          `json:"Text" xml:"Text"`
	ShareId                   string                                          `json:"ShareId" xml:"ShareId"`
	AudioFormat               string                                          `json:"AudioFormat" xml:"AudioFormat"`
	TargetData                string                                          `json:"TargetData" xml:"TargetData"`
	DynamicGroupExpression    string                                          `json:"DynamicGroupExpression" xml:"DynamicGroupExpression"`
	SourceConfig              string                                          `json:"SourceConfig" xml:"SourceConfig"`
	GroupName                 string                                          `json:"GroupName" xml:"GroupName"`
	CreateTime                int64                                           `json:"CreateTime" xml:"CreateTime"`
	DeploymentId              string                                          `json:"DeploymentId" xml:"DeploymentId"`
	BizId                     string                                          `json:"BizId" xml:"BizId"`
	FirmwareUrl               string                                          `json:"FirmwareUrl" xml:"FirmwareUrl"`
	RoleAttachTimestamp       int64                                           `json:"RoleAttachTimestamp" xml:"RoleAttachTimestamp"`
	GmtOpened                 int64                                           `json:"GmtOpened" xml:"GmtOpened"`
	CheckProgressId           string                                          `json:"CheckProgressId" xml:"CheckProgressId"`
	Description               string                                          `json:"Description" xml:"Description"`
	GmtModifiedTimestamp      int64                                           `json:"GmtModifiedTimestamp" xml:"GmtModifiedTimestamp"`
	Sn                        string                                          `json:"Sn" xml:"Sn"`
	ApiPath                   string                                          `json:"ApiPath" xml:"ApiPath"`
	Country                   string                                          `json:"Country" xml:"Country"`
	Status                    int                                             `json:"Status" xml:"Status"`
	Protocol                  string                                          `json:"Protocol" xml:"Protocol"`
	TopicFilter               string                                          `json:"TopicFilter" xml:"TopicFilter"`
	DeviceSecret              string                                          `json:"DeviceSecret" xml:"DeviceSecret"`
	AsyncExecute              bool                                            `json:"AsyncExecute" xml:"AsyncExecute"`
	ProductKey                string                                          `json:"ProductKey" xml:"ProductKey"`
	GmtModified               string                                          `json:"GmtModified" xml:"GmtModified"`
	City                      string                                          `json:"City" xml:"City"`
	DisplayName               string                                          `json:"DisplayName" xml:"DisplayName"`
	VersionState              string                                          `json:"VersionState" xml:"VersionState"`
	IotId                     string                                          `json:"IotId" xml:"IotId"`
	FailSum                   int64                                           `json:"FailSum" xml:"FailSum"`
	GroupDesc                 string                                          `json:"GroupDesc" xml:"GroupDesc"`
	Argument                  string                                          `json:"Argument" xml:"Argument"`
	SpeechRate                int                                             `json:"SpeechRate" xml:"SpeechRate"`
	ExpiringQuota             int                                             `json:"ExpiringQuota" xml:"ExpiringQuota"`
	TunnelState               string                                          `json:"TunnelState" xml:"TunnelState"`
	InvalidDeviceNameList     InvalidDeviceNameListInBatchCheckDeviceNames    `json:"InvalidDeviceNameList" xml:"InvalidDeviceNameList"`
	InvalidSnList             InvalidSnListInBatchCheckImportDevice           `json:"InvalidSnList" xml:"InvalidSnList"`
	ResultList                ResultList                                      `json:"ResultList" xml:"ResultList"`
	FieldNameList             FieldNameList                                   `json:"FieldNameList" xml:"FieldNameList"`
	Result                    Result                                          `json:"Result" xml:"Result"`
	InvalidDeviceSecretList   InvalidDeviceSecretListInBatchCheckImportDevice `json:"InvalidDeviceSecretList" xml:"InvalidDeviceSecretList"`
	InvalidDeviceNicknameList InvalidDeviceNicknameList                       `json:"InvalidDeviceNicknameList" xml:"InvalidDeviceNicknameList"`
	RepeatedDeviceNameList    RepeatedDeviceNameListInBatchCheckImportDevice  `json:"RepeatedDeviceNameList" xml:"RepeatedDeviceNameList"`
	RouteContext              RouteContext                                    `json:"RouteContext" xml:"RouteContext"`
	QuerySetting              QuerySetting                                    `json:"QuerySetting" xml:"QuerySetting"`
	SqlTemplateDTO            SqlTemplateDTO                                  `json:"SqlTemplateDTO" xml:"SqlTemplateDTO"`
	TokenInfo                 TokenInfo                                       `json:"TokenInfo" xml:"TokenInfo"`
	SoundCodeConfig           SoundCodeConfig                                 `json:"SoundCodeConfig" xml:"SoundCodeConfig"`
	Header                    []HeaderItem                                    `json:"Header" xml:"Header"`
	TaskList                  []Task                                          `json:"TaskList" xml:"TaskList"`
	ModelVersions             []ModelVersion                                  `json:"ModelVersions" xml:"ModelVersions"`
	Points                    []PointsItem                                    `json:"Points" xml:"Points"`
	List                      ListInGetThingTopo                              `json:"List" xml:"List"`
	DynamicRegClientIds       []DynamicRegClientId                            `json:"DynamicRegClientIds" xml:"DynamicRegClientIds"`
}
