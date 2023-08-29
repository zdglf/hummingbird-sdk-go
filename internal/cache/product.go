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

	"gitee.com/winc-link/hummingbird-sdk-go/model"
)

type ProductProvider interface {
	All() map[string]model.Product
	SearchById(id string) (model.Product, bool)
	Add(p model.Product)
	Update(p model.Product)
	RemoveById(id string)
	GetPropertySpecByCode(productId, code string) (model.Property, bool)
	GetServiceSpecByCode(productId, code string) (model.Service, bool)
	GetEventSpecByCode(productId, code string) (model.Event, bool)
	GetProductProperties(productId string) (map[string]model.Property, bool)
	GetProductEvents(productId string) (map[string]model.Event, bool)
	GetProductServices(productId string) (map[string]model.Service, bool)
}

type ProductCache struct {
	mu          sync.RWMutex
	productMap  map[string]*model.Product
	propertyMap map[string]map[string]model.Property
	serviceMap  map[string]map[string]model.Service
	eventMap    map[string]map[string]model.Event
}

func (t *ProductCache) GetPropertySpecByCode(productId, code string) (model.Property, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	pm, ok := t.propertyMap[productId]
	if !ok {
		return model.Property{}, false
	}
	ps, ok := pm[code]
	if !ok {
		return model.Property{}, false
	}
	return ps, ok
}

func NewProductCache(products []model.Product) *ProductCache {
	defaultSize := len(products)
	pm := make(map[string]*model.Product, defaultSize)
	ppm := make(map[string]map[string]model.Property)
	am := make(map[string]map[string]model.Service)
	em := make(map[string]map[string]model.Event)
	for i, p := range products {
		pm[p.Id] = &products[i]
		ppm[p.Id] = propertyTransformToMap(products[i].Properties)
		am[p.Id] = serviceTransformToMap(products[i].Services)
		em[p.Id] = eventTransformToMap(products[i].Events)
	}
	return &ProductCache{
		productMap:  pm,
		propertyMap: ppm,
		serviceMap:  am,
		eventMap:    em,
	}
}

func (t *ProductCache) SearchById(id string) (model.Product, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	p, ok := t.productMap[id]
	if !ok {
		return model.Product{}, false
	}
	return *p, ok
}
func (t *ProductCache) All() map[string]model.Product {
	t.mu.RLock()
	defer t.mu.RUnlock()

	ps := make(map[string]model.Product, len(t.productMap))
	for k, p := range t.productMap {
		ps[k] = *p
	}
	return ps
}
func (t *ProductCache) Add(p model.Product) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.productMap[p.Id] = &p
	t.propertyMap[p.Id] = propertyTransformToMap(p.Properties)
	t.serviceMap[p.Id] = serviceTransformToMap(p.Services)
	t.eventMap[p.Id] = eventTransformToMap(p.Events)
}

func (t *ProductCache) Update(p model.Product) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.removeById(p.Id)
	t.productMap[p.Id] = &p
	t.propertyMap[p.Id] = propertyTransformToMap(p.Properties)
	t.serviceMap[p.Id] = serviceTransformToMap(p.Services)
	t.eventMap[p.Id] = eventTransformToMap(p.Events)

}
func (t *ProductCache) RemoveById(id string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.removeById(id)
}

func (t *ProductCache) removeById(id string) {
	delete(t.productMap, id)
	delete(t.propertyMap, id)
	delete(t.serviceMap, id)
	delete(t.eventMap, id)
}

func (t *ProductCache) GetPropertyByCode(pid, code string) (model.Property, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	pm, ok := t.propertyMap[pid]
	if !ok {
		return model.Property{}, false
	}
	ps, ok := pm[code]
	if !ok {
		return model.Property{}, false
	}
	return ps, ok
}

func (t *ProductCache) GetEventSpecByCode(pid, code string) (model.Event, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	pm, ok := t.eventMap[pid]
	if !ok {
		return model.Event{}, false
	}
	ps, ok := pm[code]
	if !ok {
		return model.Event{}, false
	}
	return ps, ok
}

func (t *ProductCache) GetServiceSpecByCode(pid, code string) (model.Service, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	am, ok := t.serviceMap[pid]
	if !ok {
		return model.Service{}, false
	}
	a, ok := am[code]
	if !ok {
		return model.Service{}, false
	}
	return a, ok
}

func (t *ProductCache) GetProductEvents(productId string) (map[string]model.Event, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	e, ok := t.eventMap[productId]
	if !ok {
		return map[string]model.Event{}, false
	}
	return e, ok
}
func (t *ProductCache) GetProductProperties(productId string) (map[string]model.Property, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	p, ok := t.propertyMap[productId]
	if !ok {
		return map[string]model.Property{}, false
	}
	return p, ok
}

func (t *ProductCache) GetProductServices(productId string) (map[string]model.Service, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	s, ok := t.serviceMap[productId]
	if !ok {
		return map[string]model.Service{}, false
	}
	return s, ok
}

func propertyTransformToMap(properties []model.Property) map[string]model.Property {
	pMap := make(map[string]model.Property, len(properties))
	for _, p := range properties {
		pMap[p.Code] = p
	}
	return pMap
}

func serviceTransformToMap(services []model.Service) map[string]model.Service {
	aMap := make(map[string]model.Service, len(services))
	for _, a := range services {
		aMap[a.Code] = a
	}
	return aMap
}

func eventTransformToMap(events []model.Event) map[string]model.Event {
	eMap := make(map[string]model.Event, len(events))
	for _, e := range events {
		eMap[e.Code] = e
	}
	return eMap
}
