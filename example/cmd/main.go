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

package main

import (
	"github.com/winc-link/hummingbird-sdk-go/commons"
	"github.com/winc-link/hummingbird-sdk-go/service"

	"context"
	"github.com/winc-link/hummingbird-sdk-go/example/internal/driver"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sd := service.NewDriverService("simple-driver", commons.HummingbirdIot)
	simpleDriver := driver.NewSimpleDriver(ctx, sd)
	if err := simpleDriver.Initialize(); err != nil {
		sd.GetLogger().Error("init simple driver error: %s", err)
		return
	}
	go func() {
		if err := sd.Start(simpleDriver); err != nil {
			sd.GetLogger().Error("driver service start error: %s", err)
			return
		}
	}()
	waitForSignal(cancel, sd)
}

func waitForSignal(cancel context.CancelFunc, sd *service.DriverService) os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	//signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	s := <-signalChan
	sd.GetLogger().Info("docker stop")
	cancel()
	signal.Stop(signalChan)
	return s
}
