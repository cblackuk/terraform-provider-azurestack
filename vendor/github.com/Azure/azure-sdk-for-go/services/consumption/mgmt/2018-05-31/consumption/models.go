package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/shopspring/decimal"
	"net/http"
)

// Bound enumerates the values for bound.
type Bound string

const (
	// Lower ...
	Lower Bound = "Lower"
	// Upper ...
	Upper Bound = "Upper"
)

// PossibleBoundValues returns an array of possible values for the Bound const type.
func PossibleBoundValues() []Bound {
	return []Bound{Lower, Upper}
}

// ChargeType enumerates the values for charge type.
type ChargeType string

const (
	// ChargeTypeActual ...
	ChargeTypeActual ChargeType = "Actual"
	// ChargeTypeForecast ...
	ChargeTypeForecast ChargeType = "Forecast"
)

// PossibleChargeTypeValues returns an array of possible values for the ChargeType const type.
func PossibleChargeTypeValues() []ChargeType {
	return []ChargeType{ChargeTypeActual, ChargeTypeForecast}
}

// Grain enumerates the values for grain.
type Grain string

const (
	// Daily ...
	Daily Grain = "Daily"
	// Monthly ...
	Monthly Grain = "Monthly"
	// Yearly ...
	Yearly Grain = "Yearly"
)

// PossibleGrainValues returns an array of possible values for the Grain const type.
func PossibleGrainValues() []Grain {
	return []Grain{Daily, Monthly, Yearly}
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The reason
// is provided in the error message.
type ErrorResponse struct {
	// Error - The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// Forecast a forecast resource.
type Forecast struct {
	*ForecastProperties `json:"properties,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Forecast.
func (f Forecast) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if f.ForecastProperties != nil {
		objectMap["properties"] = f.ForecastProperties
	}
	if f.ID != nil {
		objectMap["id"] = f.ID
	}
	if f.Name != nil {
		objectMap["name"] = f.Name
	}
	if f.Type != nil {
		objectMap["type"] = f.Type
	}
	if f.Tags != nil {
		objectMap["tags"] = f.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Forecast struct.
func (f *Forecast) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var forecastProperties ForecastProperties
				err = json.Unmarshal(*v, &forecastProperties)
				if err != nil {
					return err
				}
				f.ForecastProperties = &forecastProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				f.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				f.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				f.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				f.Tags = tags
			}
		}
	}

	return nil
}

// ForecastProperties the properties of the forecast charge.
type ForecastProperties struct {
	// UsageDate - The usage date of the forecast.
	UsageDate *string `json:"usageDate,omitempty"`
	// Grain - The granularity of forecast. Possible values include: 'Daily', 'Monthly', 'Yearly'
	Grain Grain `json:"grain,omitempty"`
	// Charge - The amount of charge
	Charge *decimal.Decimal `json:"charge,omitempty"`
	// Currency - The ISO currency in which the meter is charged, for example, USD.
	Currency *string `json:"currency,omitempty"`
	// ChargeType - The type of the charge. Could be actual or forecast. Possible values include: 'ChargeTypeActual', 'ChargeTypeForecast'
	ChargeType ChargeType `json:"chargeType,omitempty"`
	// ConfidenceLevels - The details about the forecast confidence levels. This is populated only when chargeType is Forecast.
	ConfidenceLevels *[]ForecastPropertiesConfidenceLevelsItem `json:"confidenceLevels,omitempty"`
}

// ForecastPropertiesConfidenceLevelsItem ...
type ForecastPropertiesConfidenceLevelsItem struct {
	// Percentage - The percentage level of the confidence
	Percentage *decimal.Decimal `json:"percentage,omitempty"`
	// Bound - The boundary of the percentage, values could be 'Upper' or 'Lower'. Possible values include: 'Upper', 'Lower'
	Bound Bound `json:"bound,omitempty"`
	// Value - The amount of forecast within the percentage level
	Value *decimal.Decimal `json:"value,omitempty"`
}

// ForecastsListResult result of listing forecasts. It contains a list of available forecasts.
type ForecastsListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of forecasts.
	Value *[]Forecast `json:"value,omitempty"`
}

// Operation a Consumption REST API operation.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Consumption.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: UsageDetail, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of listing consumption operations. It contains a list of operations and a URL link to
// get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of consumption operations supported by the Microsoft.Consumption resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer() (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) Next() error {
	next, err := page.fn(page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Resource the Resource model definition.
type Resource struct {
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}
