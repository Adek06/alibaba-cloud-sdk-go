package amqp_open

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

// Data is a nested struct in amqp_open response
type Data struct {
	CurrentVirtualHosts int               `json:"CurrentVirtualHosts" xml:"CurrentVirtualHosts"`
	MaxQueues           int               `json:"MaxQueues" xml:"MaxQueues"`
	MaxVirtualHosts     int               `json:"MaxVirtualHosts" xml:"MaxVirtualHosts"`
	MaxExchanges        int               `json:"MaxExchanges" xml:"MaxExchanges"`
	NextToken           string            `json:"NextToken" xml:"NextToken"`
	CurrentExchanges    int               `json:"CurrentExchanges" xml:"CurrentExchanges"`
	CurrentQueues       int               `json:"CurrentQueues" xml:"CurrentQueues"`
	MaxResults          int               `json:"MaxResults" xml:"MaxResults"`
	Instances           []InstanceVO      `json:"Instances" xml:"Instances"`
	Queues              []QueueVO         `json:"Queues" xml:"Queues"`
	Exchanges           []ExchangeVO      `json:"Exchanges" xml:"Exchanges"`
	Consumers           []QueueConsumerVO `json:"Consumers" xml:"Consumers"`
	Bindings            []BindingDO       `json:"Bindings" xml:"Bindings"`
	VirtualHosts        []VhostVO         `json:"VirtualHosts" xml:"VirtualHosts"`
}
