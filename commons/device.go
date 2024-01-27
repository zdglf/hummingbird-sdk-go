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

import "github.com/zdglf/edge-driver-proto/driverdevice"

type DeviceConnectStatus string

const (
	Online  DeviceConnectStatus = "online"  //在线
	Offline DeviceConnectStatus = "offline" //离线
)

type DeviceStatus string

const (
	DeviceUnKnow   DeviceStatus = "unknow"
	DeviceOnline   DeviceStatus = "online"   //在线
	DeviceOffline  DeviceStatus = "offline"  //离线
	DeviceUnActive DeviceStatus = "unactive" //未激活
	DeviceDisable  DeviceStatus = "disable"  //禁用
)

func TransformRpcDeviceStatusToModel(deviceStatus driverdevice.DeviceStatus) DeviceStatus {
	switch deviceStatus {
	case driverdevice.DeviceStatus_OnLine:
		return DeviceOnline
	case driverdevice.DeviceStatus_OffLine:
		return DeviceOffline
	case driverdevice.DeviceStatus_UnActive:
		return DeviceUnActive
	case driverdevice.DeviceStatus_Disable:
		return DeviceDisable
	default:
		return DeviceUnKnow
	}
}
