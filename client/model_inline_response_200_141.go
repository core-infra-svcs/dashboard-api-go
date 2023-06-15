/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200141 struct for InlineResponse200141
type InlineResponse200141 struct {
	// Serial number of the sensor that took the reading.
	Serial *string `json:"serial,omitempty"`
	Network *OrganizationsOrganizationIdSensorReadingsHistoryNetwork `json:"network,omitempty"`
	// Time at which the reading occurred, in ISO8601 format.
	Ts *string `json:"ts,omitempty"`
	// Type of sensor reading.
	Metric *string `json:"metric,omitempty"`
	Battery *OrganizationsOrganizationIdSensorReadingsHistoryBattery `json:"battery,omitempty"`
	Button *OrganizationsOrganizationIdSensorReadingsHistoryButton `json:"button,omitempty"`
	Door *OrganizationsOrganizationIdSensorReadingsHistoryDoor `json:"door,omitempty"`
	Humidity *OrganizationsOrganizationIdSensorReadingsHistoryHumidity `json:"humidity,omitempty"`
	IndoorAirQuality *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality `json:"indoorAirQuality,omitempty"`
	Noise *OrganizationsOrganizationIdSensorReadingsHistoryNoise `json:"noise,omitempty"`
	Pm25 *OrganizationsOrganizationIdSensorReadingsHistoryPm25 `json:"pm25,omitempty"`
	Temperature *OrganizationsOrganizationIdSensorReadingsHistoryTemperature `json:"temperature,omitempty"`
	Tvoc *OrganizationsOrganizationIdSensorReadingsHistoryTvoc `json:"tvoc,omitempty"`
	Water *OrganizationsOrganizationIdSensorReadingsHistoryWater `json:"water,omitempty"`
}

// NewInlineResponse200141 instantiates a new InlineResponse200141 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200141() *InlineResponse200141 {
	this := InlineResponse200141{}
	return &this
}

// NewInlineResponse200141WithDefaults instantiates a new InlineResponse200141 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200141WithDefaults() *InlineResponse200141 {
	this := InlineResponse200141{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200141) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200141) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200141) SetSerial(v string) {
	o.Serial = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200141) GetNetwork() OrganizationsOrganizationIdSensorReadingsHistoryNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetNetworkOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200141) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryNetwork and assigns it to the Network field.
func (o *InlineResponse200141) SetNetwork(v OrganizationsOrganizationIdSensorReadingsHistoryNetwork) {
	o.Network = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200141) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200141) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse200141) SetTs(v string) {
	o.Ts = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *InlineResponse200141) GetMetric() string {
	if o == nil || isNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetMetricOk() (*string, bool) {
	if o == nil || isNil(o.Metric) {
    return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *InlineResponse200141) HasMetric() bool {
	if o != nil && !isNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *InlineResponse200141) SetMetric(v string) {
	o.Metric = &v
}

// GetBattery returns the Battery field value if set, zero value otherwise.
func (o *InlineResponse200141) GetBattery() OrganizationsOrganizationIdSensorReadingsHistoryBattery {
	if o == nil || isNil(o.Battery) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryBattery
		return ret
	}
	return *o.Battery
}

// GetBatteryOk returns a tuple with the Battery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetBatteryOk() (*OrganizationsOrganizationIdSensorReadingsHistoryBattery, bool) {
	if o == nil || isNil(o.Battery) {
    return nil, false
	}
	return o.Battery, true
}

// HasBattery returns a boolean if a field has been set.
func (o *InlineResponse200141) HasBattery() bool {
	if o != nil && !isNil(o.Battery) {
		return true
	}

	return false
}

// SetBattery gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryBattery and assigns it to the Battery field.
func (o *InlineResponse200141) SetBattery(v OrganizationsOrganizationIdSensorReadingsHistoryBattery) {
	o.Battery = &v
}

// GetButton returns the Button field value if set, zero value otherwise.
func (o *InlineResponse200141) GetButton() OrganizationsOrganizationIdSensorReadingsHistoryButton {
	if o == nil || isNil(o.Button) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryButton
		return ret
	}
	return *o.Button
}

// GetButtonOk returns a tuple with the Button field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetButtonOk() (*OrganizationsOrganizationIdSensorReadingsHistoryButton, bool) {
	if o == nil || isNil(o.Button) {
    return nil, false
	}
	return o.Button, true
}

// HasButton returns a boolean if a field has been set.
func (o *InlineResponse200141) HasButton() bool {
	if o != nil && !isNil(o.Button) {
		return true
	}

	return false
}

// SetButton gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryButton and assigns it to the Button field.
func (o *InlineResponse200141) SetButton(v OrganizationsOrganizationIdSensorReadingsHistoryButton) {
	o.Button = &v
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *InlineResponse200141) GetDoor() OrganizationsOrganizationIdSensorReadingsHistoryDoor {
	if o == nil || isNil(o.Door) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryDoor
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetDoorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDoor, bool) {
	if o == nil || isNil(o.Door) {
    return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *InlineResponse200141) HasDoor() bool {
	if o != nil && !isNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryDoor and assigns it to the Door field.
func (o *InlineResponse200141) SetDoor(v OrganizationsOrganizationIdSensorReadingsHistoryDoor) {
	o.Door = &v
}

// GetHumidity returns the Humidity field value if set, zero value otherwise.
func (o *InlineResponse200141) GetHumidity() OrganizationsOrganizationIdSensorReadingsHistoryHumidity {
	if o == nil || isNil(o.Humidity) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryHumidity
		return ret
	}
	return *o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetHumidityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryHumidity, bool) {
	if o == nil || isNil(o.Humidity) {
    return nil, false
	}
	return o.Humidity, true
}

// HasHumidity returns a boolean if a field has been set.
func (o *InlineResponse200141) HasHumidity() bool {
	if o != nil && !isNil(o.Humidity) {
		return true
	}

	return false
}

// SetHumidity gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryHumidity and assigns it to the Humidity field.
func (o *InlineResponse200141) SetHumidity(v OrganizationsOrganizationIdSensorReadingsHistoryHumidity) {
	o.Humidity = &v
}

// GetIndoorAirQuality returns the IndoorAirQuality field value if set, zero value otherwise.
func (o *InlineResponse200141) GetIndoorAirQuality() OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality {
	if o == nil || isNil(o.IndoorAirQuality) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality
		return ret
	}
	return *o.IndoorAirQuality
}

// GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetIndoorAirQualityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality, bool) {
	if o == nil || isNil(o.IndoorAirQuality) {
    return nil, false
	}
	return o.IndoorAirQuality, true
}

// HasIndoorAirQuality returns a boolean if a field has been set.
func (o *InlineResponse200141) HasIndoorAirQuality() bool {
	if o != nil && !isNil(o.IndoorAirQuality) {
		return true
	}

	return false
}

// SetIndoorAirQuality gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality and assigns it to the IndoorAirQuality field.
func (o *InlineResponse200141) SetIndoorAirQuality(v OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) {
	o.IndoorAirQuality = &v
}

// GetNoise returns the Noise field value if set, zero value otherwise.
func (o *InlineResponse200141) GetNoise() OrganizationsOrganizationIdSensorReadingsHistoryNoise {
	if o == nil || isNil(o.Noise) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryNoise
		return ret
	}
	return *o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetNoiseOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNoise, bool) {
	if o == nil || isNil(o.Noise) {
    return nil, false
	}
	return o.Noise, true
}

// HasNoise returns a boolean if a field has been set.
func (o *InlineResponse200141) HasNoise() bool {
	if o != nil && !isNil(o.Noise) {
		return true
	}

	return false
}

// SetNoise gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryNoise and assigns it to the Noise field.
func (o *InlineResponse200141) SetNoise(v OrganizationsOrganizationIdSensorReadingsHistoryNoise) {
	o.Noise = &v
}

// GetPm25 returns the Pm25 field value if set, zero value otherwise.
func (o *InlineResponse200141) GetPm25() OrganizationsOrganizationIdSensorReadingsHistoryPm25 {
	if o == nil || isNil(o.Pm25) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryPm25
		return ret
	}
	return *o.Pm25
}

// GetPm25Ok returns a tuple with the Pm25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetPm25Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm25, bool) {
	if o == nil || isNil(o.Pm25) {
    return nil, false
	}
	return o.Pm25, true
}

// HasPm25 returns a boolean if a field has been set.
func (o *InlineResponse200141) HasPm25() bool {
	if o != nil && !isNil(o.Pm25) {
		return true
	}

	return false
}

// SetPm25 gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryPm25 and assigns it to the Pm25 field.
func (o *InlineResponse200141) SetPm25(v OrganizationsOrganizationIdSensorReadingsHistoryPm25) {
	o.Pm25 = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *InlineResponse200141) GetTemperature() OrganizationsOrganizationIdSensorReadingsHistoryTemperature {
	if o == nil || isNil(o.Temperature) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryTemperature
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTemperature, bool) {
	if o == nil || isNil(o.Temperature) {
    return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *InlineResponse200141) HasTemperature() bool {
	if o != nil && !isNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryTemperature and assigns it to the Temperature field.
func (o *InlineResponse200141) SetTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryTemperature) {
	o.Temperature = &v
}

// GetTvoc returns the Tvoc field value if set, zero value otherwise.
func (o *InlineResponse200141) GetTvoc() OrganizationsOrganizationIdSensorReadingsHistoryTvoc {
	if o == nil || isNil(o.Tvoc) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryTvoc
		return ret
	}
	return *o.Tvoc
}

// GetTvocOk returns a tuple with the Tvoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetTvocOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTvoc, bool) {
	if o == nil || isNil(o.Tvoc) {
    return nil, false
	}
	return o.Tvoc, true
}

// HasTvoc returns a boolean if a field has been set.
func (o *InlineResponse200141) HasTvoc() bool {
	if o != nil && !isNil(o.Tvoc) {
		return true
	}

	return false
}

// SetTvoc gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryTvoc and assigns it to the Tvoc field.
func (o *InlineResponse200141) SetTvoc(v OrganizationsOrganizationIdSensorReadingsHistoryTvoc) {
	o.Tvoc = &v
}

// GetWater returns the Water field value if set, zero value otherwise.
func (o *InlineResponse200141) GetWater() OrganizationsOrganizationIdSensorReadingsHistoryWater {
	if o == nil || isNil(o.Water) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryWater
		return ret
	}
	return *o.Water
}

// GetWaterOk returns a tuple with the Water field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetWaterOk() (*OrganizationsOrganizationIdSensorReadingsHistoryWater, bool) {
	if o == nil || isNil(o.Water) {
    return nil, false
	}
	return o.Water, true
}

// HasWater returns a boolean if a field has been set.
func (o *InlineResponse200141) HasWater() bool {
	if o != nil && !isNil(o.Water) {
		return true
	}

	return false
}

// SetWater gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryWater and assigns it to the Water field.
func (o *InlineResponse200141) SetWater(v OrganizationsOrganizationIdSensorReadingsHistoryWater) {
	o.Water = &v
}

func (o InlineResponse200141) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !isNil(o.Battery) {
		toSerialize["battery"] = o.Battery
	}
	if !isNil(o.Button) {
		toSerialize["button"] = o.Button
	}
	if !isNil(o.Door) {
		toSerialize["door"] = o.Door
	}
	if !isNil(o.Humidity) {
		toSerialize["humidity"] = o.Humidity
	}
	if !isNil(o.IndoorAirQuality) {
		toSerialize["indoorAirQuality"] = o.IndoorAirQuality
	}
	if !isNil(o.Noise) {
		toSerialize["noise"] = o.Noise
	}
	if !isNil(o.Pm25) {
		toSerialize["pm25"] = o.Pm25
	}
	if !isNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !isNil(o.Tvoc) {
		toSerialize["tvoc"] = o.Tvoc
	}
	if !isNil(o.Water) {
		toSerialize["water"] = o.Water
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200141 struct {
	value *InlineResponse200141
	isSet bool
}

func (v NullableInlineResponse200141) Get() *InlineResponse200141 {
	return v.value
}

func (v *NullableInlineResponse200141) Set(val *InlineResponse200141) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200141) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200141) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200141(val *InlineResponse200141) *NullableInlineResponse200141 {
	return &NullableInlineResponse200141{value: val, isSet: true}
}

func (v NullableInlineResponse200141) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200141) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

