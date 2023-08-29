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

import (
	"github.com/winc-link/edge-driver-proto/drivercommon"
)

const Version = "1.0"

type ACK struct {
	Ack int8 `json:"ack"`
}

type CommonResponse struct {
	//RequestId    string
	ErrorMessage string
	Code         string
	Success      bool
}

type CommonRequest struct {
	Version string `json:"version"`
	MsgId   string `json:"msgId"`
	Time    int64  `json:"time"`
	Sys     ACK    `json:"sys"`
}

func NewCommonResponse(resp *drivercommon.CommonResponse) CommonResponse {
	return CommonResponse{
		ErrorMessage: resp.GetErrorMessage(),
		Code:         resp.GetCode(),
		Success:      resp.GetSuccess(),
	}
}
