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

package commons

import (
	"github.com/zdglf/edge-driver-proto/drivercommon"
)

type BaseMessage struct {
	UsePlatform      bool
	CloudServiceInfo *CloudServiceInfo
	DriverInstanceId string
}

func (c BaseMessage) UseCloudPlatform() bool {
	return c.UsePlatform
}

func (c BaseMessage) BuildBaseRequest() *drivercommon.BaseRequestMessage {
	baseRpcRequest := new(drivercommon.BaseRequestMessage)
	baseRpcRequest.UseCloudPlatform = c.UsePlatform
	baseRpcRequest.DriverInstanceId = c.DriverInstanceId
	if baseRpcRequest.UseCloudPlatform {
		baseRpcRequest.CloudInstanceInfo = new(drivercommon.CloudInstanceInfo)
		baseRpcRequest.CloudInstanceInfo.CloudInstanceName = c.CloudServiceInfo.CloudInstanceName
		baseRpcRequest.CloudInstanceInfo.CloudInstanceId = c.CloudServiceInfo.CloudInstanceId
		baseRpcRequest.CloudInstanceInfo.BaseAddress = c.CloudServiceInfo.Address
		baseRpcRequest.CloudInstanceInfo.IotPlatform = c.CloudServiceInfo.Platform.TransformModelToRpcPlatform()
	}
	return baseRpcRequest
}
