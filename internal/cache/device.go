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
	"sync"

	"github.com/winc-link/hummingbird-sdk-go/model"
)

type DeviceProvider interface {
	SearchById(id string) (model.Device, bool)
	All() map[string]model.Device
	Add(d model.Device)
	Update(d model.Device)
	RemoveById(id string)
}

type DeviceCache struct {
	mutex     sync.RWMutex
	deviceMap map[string]*model.Device
}

func NewDeviceCache(devices []model.Device) *DeviceCache {
	defaultSize := len(devices)
	dm := make(map[string]*model.Device, defaultSize)
	for i, d := range devices {
		dm[d.Id] = &devices[i]
	}

	return &DeviceCache{
		deviceMap: dm,
	}
}

func (dc *DeviceCache) SearchById(id string) (model.Device, bool) {
	dc.mutex.RLock()
	defer dc.mutex.RUnlock()

	d, ok := dc.deviceMap[id]
	if !ok {
		return model.Device{}, ok
	}
	return *d, ok
}

func (dc *DeviceCache) All() map[string]model.Device {
	dc.mutex.RLock()
	defer dc.mutex.RUnlock()

	dMap := make(map[string]model.Device, len(dc.deviceMap))
	for k, d := range dc.deviceMap {
		dMap[k] = *d
	}
	return dMap
}

func (dc *DeviceCache) Add(d model.Device) {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	dc.deviceMap[d.Id] = &d
}

func (dc *DeviceCache) Update(d model.Device) {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	dc.deviceMap[d.Id] = &d
}

func (dc *DeviceCache) RemoveById(id string) {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	delete(dc.deviceMap, id)
}
