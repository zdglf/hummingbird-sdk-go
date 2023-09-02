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

type (
	ServiceDataIn struct {
		Code        string                 `json:"code"`
		InputParams map[string]interface{} `json:"inputParams"`
	}
	ServiceDataOut struct {
		Code         string                 `json:"code"`
		OutputParams map[string]interface{} `json:"outputParams"`
	}

	ServiceExecuteRequest struct {
		CommonRequest `json:",inline"`
		Data          ServiceDataIn `json:"data"`
		Spec          Service       `json:"-"`
	}

	// ServiceExecuteResponse 执行设备动作响应
	ServiceExecuteResponse struct {
		MsgId string         `json:"msgId"`
		Data  ServiceDataOut `json:"data"`
	}
)

func NewServiceExecuteResponse(msgId string, data ServiceDataOut) ServiceExecuteResponse {
	return ServiceExecuteResponse{
		MsgId: msgId,
		Data:  data,
	}
}
