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

package cache

import (
	"context"
	"errors"
	"github.com/winc-link/edge-driver-proto/driverdevice"
	"github.com/winc-link/edge-driver-proto/driverproduct"
	"github.com/winc-link/hummingbird-sdk-go/commons"
	"github.com/winc-link/hummingbird-sdk-go/internal/client"
	"github.com/winc-link/hummingbird-sdk-go/internal/logger"
	"github.com/winc-link/hummingbird-sdk-go/model"
	"time"
)

func InitDeviceCache(baseMessage commons.BaseMessage, cli *client.ResourceClient, logger logger.Logger) (*DeviceCache, error) {
	var (
		err     error
		resp    *driverdevice.QueryDeviceListResponse
		devices []model.Device
	)
	c, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if resp, err = cli.RpcDeviceClient.QueryDeviceList(c, &driverdevice.QueryDeviceListRequest{
		BaseRequest: baseMessage.BuildBaseRequest(),
	}); err != nil {
		return nil, err
	}
	if !resp.BaseResponse.Success {
		return nil, errors.New(resp.BaseResponse.ErrorMessage)
	}
	if resp.Data != nil {
		for _, device := range resp.Data.Devices {
			devices = append(devices, model.TransformDeviceModel(device))
		}
	}
	return NewDeviceCache(devices), nil
}

func InitProductCache(baseMessage commons.BaseMessage, cli *client.ResourceClient, logger logger.Logger) (*ProductCache, error) {
	var (
		err   error
		ps    []model.Product
		dpMap = make(map[string]struct{})
	)
	c, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	resp, err := cli.RpcProductClient.QueryProductList(c, &driverproduct.QueryProductListRequest{
		BaseRequest: baseMessage.BuildBaseRequest(),
	})
	if err != nil {
		return nil, err
	}
	if !resp.BaseResponse.Success {
		return nil, errors.New(resp.BaseResponse.ErrorMessage)
	}

	for _, p := range resp.Data.Products {
		if _, ok := dpMap[p.Id]; ok {
			continue
		}
		dpMap[p.Id] = struct{}{}
		ps = append(ps, model.TransformProductModel(p))
	}
	return NewProductCache(ps), nil
}
