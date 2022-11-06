package arms

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

// PageInfo is a nested struct in arms response
type PageInfo struct {
	NavigatePageNums  string         `json:"NavigatePageNums" xml:"NavigatePageNums"`
	NavigateFirstPage string         `json:"NavigateFirstPage" xml:"NavigateFirstPage"`
	HasNextPage       string         `json:"HasNextPage" xml:"HasNextPage"`
	Page              int64          `json:"Page" xml:"Page"`
	Pages             string         `json:"Pages" xml:"Pages"`
	IsFirstPage       bool           `json:"IsFirstPage" xml:"IsFirstPage"`
	NavigateLastPage  string         `json:"NavigateLastPage" xml:"NavigateLastPage"`
	HasPreviousPage   bool           `json:"HasPreviousPage" xml:"HasPreviousPage"`
	Size              int64          `json:"Size" xml:"Size"`
	Total             int64          `json:"Total" xml:"Total"`
	Prepage           string         `json:"Prepage" xml:"Prepage"`
	NextPage          string         `json:"NextPage" xml:"NextPage"`
	IsLastPage        bool           `json:"IsLastPage" xml:"IsLastPage"`
	Integrations      []Integrations `json:"Integrations" xml:"Integrations"`
	List              []List         `json:"List" xml:"List"`
}
