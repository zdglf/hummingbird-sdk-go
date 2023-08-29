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

package device

import (
	"context"
	"crypto/rand"
	"gitee.com/winc-link/hummingbird-sdk-go/commons"
	"gitee.com/winc-link/hummingbird-sdk-go/internal/logger"
	"gitee.com/winc-link/hummingbird-sdk-go/model"
	"gitee.com/winc-link/hummingbird-sdk-go/service"
	"math/big"
	"time"
)

type Device struct {
	sd     *service.DriverService
	ctx    context.Context
	logger logger.Logger
}

func RandomNum() int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	return n.Int64()
}

func (d *Device) run(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 5)
	devices := d.sd.GetDeviceList()

	for {
		select {
		case <-ticker.C:
			for _, device := range devices {
				d.logger.Info("Print current time: ", time.Now().Format("2006-01-02 15:04:05"))
				deviceStatus, _ := d.sd.GetConnectStatus(device.Id)
				if deviceStatus != commons.Online {
					_ = d.sd.Online(device.Id)
				}
				//构造数据
				reportData := model.BatchData{
					Properties: map[string]model.BatchProperty{
						"electric_fr":  {Value: RandomNum()},
						"electric_fra": {Value: RandomNum()},
						"electric_frb": {Value: RandomNum()},
						"electric_frc": {Value: RandomNum()},
						"electric_pfa": {Value: RandomNum()},
						"electric_pfb": {Value: RandomNum()},
						"electric_pfc": {Value: RandomNum()},
						"electric_pqa": {Value: RandomNum()},
						"electric_pqb": {Value: RandomNum()},
						"electric_pqc": {Value: RandomNum()},
					},
				}
				d.logger.Infof("device id %s , report data %v", device.Id, reportData)
				_, err := d.sd.BatchReport(device.Id, model.NewBatchReport(false, reportData))
				if err != nil {
					d.logger.Error("report data error:", err.Error())
				}
			}
		case <-ctx.Done():
			d.logger.Info("simple driver stop")
			for _, device := range devices {
				_ = d.sd.Offline(device.Id)
			}
			return
		}
	}
}

func Initialize(ctx context.Context, sd *service.DriverService) {

	devices := sd.GetDeviceList()
	sd.GetLogger().Info("device count:", len(devices))

	for _, device := range devices {
		dev := device
		go deviceReportData(dev.Id, sd)
	}
}

func deviceReportData(deviceId string, sd *service.DriverService) {
	deviceStatus, _ := sd.GetConnectStatus(deviceId)
	if deviceStatus != commons.Online {
		_ = sd.Online(deviceId)
	}
	for {
		time.Sleep(1 * time.Second)
		//构造数据
		reportData := model.BatchData{
			Properties: map[string]model.BatchProperty{
				"electric_fr":  {Value: RandomNum()},
				"electric_fra": {Value: RandomNum()},
				"electric_frb": {Value: RandomNum()},
				"electric_frc": {Value: RandomNum()},
				"electric_pfa": {Value: RandomNum()},
				"electric_pfb": {Value: RandomNum()},
				"electric_pfc": {Value: RandomNum()},
				"electric_pqa": {Value: RandomNum()},
				"electric_pqb": {Value: RandomNum()},
				"electric_pqc": {Value: RandomNum()},
			},
		}
		sd.GetLogger().Infof("device id %s , report data %v", deviceId, reportData)
		_, err := sd.BatchReport(deviceId, model.NewBatchReport(false, reportData))
		if err != nil {
			sd.GetLogger().Error("report data error:", err.Error())
		}
	}
}
