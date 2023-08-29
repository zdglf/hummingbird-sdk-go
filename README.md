# hummingbird-sdk-go

#### 介绍
蜂鸟物联网平台官方Go语言SDK。

#### 方法

## Online
**功能说明** <br>

`设备连接物联网平台`

**方法原型** <br>
```go
func (d *DriverService) Online (deviceId string) error
```

## Offline
设备离线

- `deviceId` 设备Id标识

```go
func (d *DriverService) Offline (deviceId string) error 
```


## GetConnectStatus
**功能说明** <br>

`获取设备连接状态`

**云平台支持** <br>


## CreateDevice
创建设备

```go
func (d *DriverService) CreateDevice(device model.AddDevice) (model.Device, error) 
```


## GetDeviceList
获取所有的设备

```go
func (d *DriverService) GetDeviceList() map[string]model.Device 
```

## GetDeviceById
通过设备id获取设备详情
```go
func (d *DriverService) GetDeviceById(deviceId string) (model.Device, bool) 
```

## ProductList
获取当前实例下的所有产品
```go
func (d *DriverService) ProductList() map[string]model.Product 
```


## GetProductById
根据产品id获取产品信息
```go
func (d *DriverService) GetProductById(productId string) (model.Product, bool) 
```

## GetProductProperties
根据产品id获取产品所有属性信息
```go
func (d *DriverService) GetProductProperties(productId string) (map[string]model.Property, bool) 
```

## GetProductPropertyByCode
根据产品id与code获取属性信息
```go
func (d *DriverService) GetProductPropertyByCode(productId, code string) (model.Property, bool) 
```

## GetProductEvents
根据产品id获取产品所有事件信息
```go
func (d *DriverService) GetProductEvents(productId string) (map[string]model.Event, bool) 
```

## GetProductEventByCode
根据产品id与code获取事件信息
```go
func (d *DriverService) GetProductEventByCode(productId, code string) (model.Event, bool) 
```

## GetPropertyServices
根据产品id获取产品所有服务信息
```go
func (d *DriverService) GetPropertyServices(productId string) (map[string]model.Service, bool) 
```

## GetProductServiceByCode
根据产品id与code获取服务信息
```go
func (d *DriverService) GetProductServiceByCode(productId, code string) (model.Service, bool) 
```

## PropertyReport
物模型属性上报 如果data参数中的Sys.Ack设置为1，则该方法会同步阻塞等待云端返回结果。
```go
func (d *DriverService) PropertyReport(deviceId string, data model.PropertyReport) (model.CommonResponse, error) 
```

## EventReport
物模型事件上报
```go
func (d *DriverService) EventReport(deviceId string, data model.EventReport) (model.CommonResponse, error) 
```

## BatchReport
设备批量上报属性和事件
```go
func (d *DriverService) BatchReport(deviceId string, data model.BatchReport) (model.CommonResponse, error) 
```


## PropertySetResponse
设备属性下发响应
```go
func (d *DriverService) PropertySetResponse(deviceId string, data model.CommonResponse) error 
```

## PropertyGetResponse
设备属性查询响应
```go
func (d *DriverService) PropertyGetResponse(deviceId string, data model.PropertyGetResponse) error
```

## ServiceExecuteResponse
设备动作执行响应
```go
func (d *DriverService) ServiceExecuteResponse(deviceId string, data model.ServiceExecuteResponse) error 
```

## GetCustomStorage
根据key值获取驱动存储的自定义内容
```go
func (d *DriverService) GetCustomStorage(keys []string) (map[string][]byte, error)
```

## PutCustomStorage
存储驱动的自定义内容
```go
func (d *DriverService) PutCustomStorage(kvs map[string][]byte) error 
```

## DeleteCustomStorage
根据key值删除驱动存储的自定义内容
```go
func (d *DriverService) DeleteCustomStorage(keys []string) error 
```

## GetAllCustomStorage
获取所有驱动存储的自定义内容
```go
func (d *DriverService) GetAllCustomStorage() (map[string][]byte, error) 
```