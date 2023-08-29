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

package driver

import (
	"context"
	"github.com/winc-link/hummingbird-sdk-go/commons"
	"github.com/winc-link/hummingbird-sdk-go/example/internal/device"
	"github.com/winc-link/hummingbird-sdk-go/model"
	"github.com/winc-link/hummingbird-sdk-go/service"
)

type SimpleDriver struct {
	ctx context.Context
	sd  *service.DriverService
}

func (s *SimpleDriver) CloudPluginNotify(ctx context.Context, notifyType commons.CloudPluginNotifyType, plugName string) error {
	s.sd.GetLogger().Infof("CloudPluginNotify notifyType:%s plugName:%s ", notifyType, plugName)
	return nil
}

func (s *SimpleDriver) DeviceNotify(ctx context.Context, notifyType commons.DeviceNotifyType, deviceId string, device model.Device) error {
	s.sd.GetLogger().Infof("DeviceNotify notifyType:%s deviceId:%s device:%v", notifyType, deviceId, device)
	return nil
}

func (s *SimpleDriver) ProductNotify(ctx context.Context, notifyType commons.ProductNotifyType, productId string, product model.Product) error {
	s.sd.GetLogger().Infof("ProductNotify notifyType:%s productId:%s product:%v", notifyType, productId, product)
	return nil
}

func (s *SimpleDriver) Stop(ctx context.Context) error {
	return nil
}

func (s *SimpleDriver) HandlePropertySet(ctx context.Context, deviceId string, data model.PropertySet) error {
	s.sd.GetLogger().Infof("HandlePropertySet deviceId:%s data:%v", deviceId, data)
	return nil
}

func (s *SimpleDriver) HandlePropertyGet(ctx context.Context, deviceId string, data model.PropertyGet) error {
	s.sd.GetLogger().Infof("HandlePropertyGet deviceId:%s data:%v", deviceId, data)
	return nil
}

func (s *SimpleDriver) HandleServiceExecute(ctx context.Context, deviceId string, data model.ServiceExecuteRequest) error {
	s.sd.GetLogger().Infof("HandlePropertyGet deviceId:%s data:%v", deviceId, data)
	return nil
}

func NewSimpleDriver(ctx context.Context, sd *service.DriverService) *SimpleDriver {
	return &SimpleDriver{
		ctx: ctx,
		sd:  sd,
	}
}

func (s *SimpleDriver) Initialize() error {
	device.Initialize(s.ctx, s.sd)
	return nil
}
