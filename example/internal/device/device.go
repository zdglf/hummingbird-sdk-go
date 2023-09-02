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
	"github.com/winc-link/hummingbird-sdk-go/commons"
	"github.com/winc-link/hummingbird-sdk-go/internal/logger"
	"github.com/winc-link/hummingbird-sdk-go/model"
	"github.com/winc-link/hummingbird-sdk-go/service"
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

func run(sd *service.DriverService) {
	ticker := time.NewTicker(time.Second * 5)
	devices := sd.GetDeviceList()

	for {
		select {
		case <-ticker.C:
			for _, device := range devices {
				sd.GetLogger().Info("Print current time: ", time.Now().Format("2006-01-02 15:04:05"))
				deviceStatus, _ := sd.GetConnectStatus(device.Id)
				if deviceStatus != commons.Online {
					_ = sd.Online(device.Id)
				}
				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_fr": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_fra": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_frb": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_frc": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_pfa": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_pfb": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_pfc": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_pqa": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_pqb": model.NewPropertyData(RandomNum()),
				}))

				_, _ = sd.PropertyReport(device.Id, model.NewPropertyReport(false, map[string]model.PropertyData{
					"electric_pqc": model.NewPropertyData(RandomNum()),
				}))
			}
		}
	}
}

func Initialize(ctx context.Context, sd *service.DriverService) {
	go run(sd)
}
