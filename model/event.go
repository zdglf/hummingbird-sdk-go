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
	// EventReport 设备向云端上报事件
	EventReport struct {
		CommonRequest `json:",inline"`
		Data          EventData `json:"data"`
	}
	EventData struct {
		EventCode    string                 `json:"eventCode"`
		EventTime    int64                  `json:"eventTime"`
		OutputParams map[string]interface{} `json:"outputParams"`
	}
)

func NewEventData(code string, outputParams map[string]interface{}) EventData {
	return EventData{
		EventCode:    code,
		OutputParams: outputParams,
		EventTime:    time.Now().UnixMilli(),
	}
}

func NewEventReport(needACK bool, data EventData) EventReport {
	var ack int8
	if needACK {
		ack = 1
	}
	return EventReport{
		CommonRequest: CommonRequest{
			Version: Version,
			//Time:    time.Now().UnixMilli(),
			Sys: ACK{
				Ack: ack,
			},
		},
		Data: data,
	}
}
