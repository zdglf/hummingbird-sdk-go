/*******************************************************************************
 * Copyright 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package model

import "time"

type (
	PropertyData struct {
		Value interface{} `json:"value"` // 上报的属性值
		Time  int64       `json:"time"`  // 属性变更时间戳
	}

	// PropertyReport 属性上报 属性查询响应
	PropertyReport struct {
		CommonRequest `json:",inline"`
		Data          map[string]PropertyData `json:"data"`
	}

	// PropertySet 属性下发
	PropertySet struct {
		CommonRequest `json:",inline"`
		Data          map[string]interface{} `json:"data"`
		Spec          map[string]Property    `json:"-"`
	}

	// PropertyGet 属性查询
	PropertyGet struct {
		CommonRequest `json:",inline"`
		Data          []string            `json:"data"`
		Spec          map[string]Property `json:"-"`
	}
	// PropertyGetResponse 属性查询设备响应
	PropertyGetResponse struct {
		CommonResponse `json:",inline"`
		Data           map[string]Property `json:"data"`
	}
	// PropertyDesiredGet 设备拉取属性期望值
	PropertyDesiredGet struct {
		CommonRequest `json:",inline"`
		Data          []string `json:"data"`
	}

	// PropertyDesiredGetResponse 设备拉取属性期望值响应
	PropertyDesiredGetResponse struct {
		CommonResponse `json:",inline"`
		Data           map[string]PropertyDesiredGetValue `json:"data"`
	}

	// PropertyDesiredDelete 设备清除属性期望值
	PropertyDesiredDelete struct {
		CommonRequest `json:",inline"`
		Data          map[string]PropertyDesiredDeleteValue `json:"data"`
	}

	// PropertyDesiredDeleteResponse 设备清除属性期望值响应
	PropertyDesiredDeleteResponse struct {
		CommonResponse `json:",inline"`
		Data           map[string]PropertyDesiredGetValue `json:"data"`
	}

	PropertyDesiredGetValue struct {
		Value   interface{} `json:"value"`
		Version int64       `json:"version"`
	}

	PropertyDesiredDeleteValue struct {
		Version int64 `json:"version"`
	}
)

func NewPropertyData(value interface{}) PropertyData {
	return PropertyData{
		Value: value,
		Time:  time.Now().UnixMilli(),
	}
}

func NewPropertyReport(ack bool, data map[string]PropertyData) PropertyReport {
	var needAck int8
	if ack {
		needAck = 1
	}

	return PropertyReport{
		CommonRequest: CommonRequest{
			Version: Version,
			//MsgId: node.GetId().String()
			//Time:    time.Now().UnixMilli(),
			Sys: ACK{
				Ack: needAck,
			},
		},
		Data: data,
	}
}
