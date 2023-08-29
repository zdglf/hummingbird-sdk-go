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
	"encoding/json"
	"errors"
	"github.com/winc-link/edge-driver-proto/thingmodel"
)

const (
	PropertyReport = iota + 1
	PropertySetResponse
	PropertyGetResponse
	ServiceExecuteResponse
	EventReport
	BatchReport
	LogReport
	PropertyDesiredGet
	PropertyDesiredDelete
)

func TransformToProtoMsg(deviceId string, t int, data interface{}, baseMsg BaseMessage) (*thingmodel.ThingModelMsg, error) {
	var opt thingmodel.OperationType
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	switch t {
	case PropertyReport:
		opt = thingmodel.OperationType_PROPERTY_REPORT
	case PropertySetResponse:
		opt = thingmodel.OperationType_PROPERTY_SET_RESPONSE
	case PropertyGetResponse:
		opt = thingmodel.OperationType_PROPERTY_GET_RESPONSE
	case ServiceExecuteResponse:
		opt = thingmodel.OperationType_SERVICE_EXECUTE_RESPONSE
	case EventReport:
		opt = thingmodel.OperationType_EVENT_REPORT
	case BatchReport:
		opt = thingmodel.OperationType_DATA_BATCH_REPORT
	case PropertyDesiredGet:
		opt = thingmodel.OperationType_PROPERTY_DESIRED_GET
	case PropertyDesiredDelete:
		opt = thingmodel.OperationType_PROPERTY_DESIRED_DELETE
	case LogReport:

	default:
		return nil, errors.New("unsupported")
	}
	return &thingmodel.ThingModelMsg{
		BaseRequest:   baseMsg.BuildBaseRequest(),
		DeviceId:      deviceId,
		OperationType: opt,
		Data:          string(payload),
	}, nil
}
