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
	"github.com/winc-link/edge-driver-proto/drivercommon"
)

type IotPlatform string

const (
	HummingbirdIot IotPlatform = "蜂鸟物联网平台"     //蜂鸟平台
	WinCLinkIot    IotPlatform = "赢创物联网平台"     //赢创万联
	AliIot         IotPlatform = "阿里物联网平台"     //阿里
	HuaweiIot      IotPlatform = "华为物联网平台"     //华为
	TencentIot     IotPlatform = "腾讯物联网平台"     //腾讯
	OneNetIot      IotPlatform = "OneNET物联网平台" //中国移动
)

func (i IotPlatform) TransformModelToRpcPlatform() drivercommon.IotPlatform {
	switch i {
	case HummingbirdIot:
		return drivercommon.IotPlatform_LocalIot
	case WinCLinkIot:
		return drivercommon.IotPlatform_WinCLinkIot
	case AliIot:
		return drivercommon.IotPlatform_AliIot
	case HuaweiIot:
		return drivercommon.IotPlatform_HuaweiIot
	case TencentIot:
		return drivercommon.IotPlatform_TencentIot
	case OneNetIot:
		return drivercommon.IotPlatform_OneNetIot
	default:
		return drivercommon.IotPlatform_LocalIot
	}
}

func TransformRpcPlatformToModel(p drivercommon.IotPlatform) IotPlatform {
	switch p {
	case drivercommon.IotPlatform_LocalIot:
		return HummingbirdIot
	case drivercommon.IotPlatform_WinCLinkIot:
		return WinCLinkIot
	case drivercommon.IotPlatform_AliIot:
		return AliIot
	case drivercommon.IotPlatform_HuaweiIot:
		return HuaweiIot
	case drivercommon.IotPlatform_TencentIot:
		return TencentIot
	case drivercommon.IotPlatform_OneNetIot:
		return OneNetIot
	default:
		return HummingbirdIot
	}
}
