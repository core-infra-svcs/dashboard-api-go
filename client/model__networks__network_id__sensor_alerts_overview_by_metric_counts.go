/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsOverviewByMetricCounts Counts of sensor alerts over the timespan, by reading metric
type NetworksNetworkIdSensorAlertsOverviewByMetricCounts struct {
	// Number of sensor alerts that occurred due to apparent power readings
	ApparentPower *int32 `json:"apparentPower,omitempty"`
	// Number of sensors that are currently alerting due to CO2 readings
	Co2 *int32 `json:"co2,omitempty"`
	// Number of sensor alerts that occurred due to electrical current readings
	Current *int32 `json:"current,omitempty"`
	// Number of sensor alerts that occurred due to an open door
	Door *int32 `json:"door,omitempty"`
	// Number of sensor alerts that occurred due to frequency readings
	Frequency *int32 `json:"frequency,omitempty"`
	// Number of sensor alerts that occurred due to humidity readings
	Humidity *int32 `json:"humidity,omitempty"`
	// Number of sensor alerts that occurred due to indoor air quality readings
	IndoorAirQuality *int32 `json:"indoorAirQuality,omitempty"`
	Noise *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise `json:"noise,omitempty"`
	// Number of sensor alerts that occurred due to PM2.5 readings
	Pm25 *int32 `json:"pm25,omitempty"`
	// Number of sensor alerts that occurred due to power factor readings
	PowerFactor *int32 `json:"powerFactor,omitempty"`
	// Number of sensor alerts that occurred due to real power readings
	RealPower *int32 `json:"realPower,omitempty"`
	// Number of sensor alerts that occurred due to temperature readings
	Temperature *int32 `json:"temperature,omitempty"`
	// Number of sensor alerts that occurred due to TVOC readings
	Tvoc *int32 `json:"tvoc,omitempty"`
	// Number of sensor alerts that occurred due to upstream power outages
	UpstreamPower *int32 `json:"upstreamPower,omitempty"`
	// Number of sensor alerts that occurred due to voltage readings
	Voltage *int32 `json:"voltage,omitempty"`
	// Number of sensor alerts that occurred due to the presence of water
	Water *int32 `json:"water,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsOverviewByMetricCounts instantiates a new NetworksNetworkIdSensorAlertsOverviewByMetricCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsOverviewByMetricCounts() *NetworksNetworkIdSensorAlertsOverviewByMetricCounts {
	this := NetworksNetworkIdSensorAlertsOverviewByMetricCounts{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsOverviewByMetricCountsWithDefaults instantiates a new NetworksNetworkIdSensorAlertsOverviewByMetricCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsOverviewByMetricCountsWithDefaults() *NetworksNetworkIdSensorAlertsOverviewByMetricCounts {
	this := NetworksNetworkIdSensorAlertsOverviewByMetricCounts{}
	return &this
}

// GetApparentPower returns the ApparentPower field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetApparentPower() int32 {
	if o == nil || isNil(o.ApparentPower) {
		var ret int32
		return ret
	}
	return *o.ApparentPower
}

// GetApparentPowerOk returns a tuple with the ApparentPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetApparentPowerOk() (*int32, bool) {
	if o == nil || isNil(o.ApparentPower) {
    return nil, false
	}
	return o.ApparentPower, true
}

// HasApparentPower returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasApparentPower() bool {
	if o != nil && !isNil(o.ApparentPower) {
		return true
	}

	return false
}

// SetApparentPower gets a reference to the given int32 and assigns it to the ApparentPower field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetApparentPower(v int32) {
	o.ApparentPower = &v
}

// GetCo2 returns the Co2 field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetCo2() int32 {
	if o == nil || isNil(o.Co2) {
		var ret int32
		return ret
	}
	return *o.Co2
}

// GetCo2Ok returns a tuple with the Co2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetCo2Ok() (*int32, bool) {
	if o == nil || isNil(o.Co2) {
    return nil, false
	}
	return o.Co2, true
}

// HasCo2 returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasCo2() bool {
	if o != nil && !isNil(o.Co2) {
		return true
	}

	return false
}

// SetCo2 gets a reference to the given int32 and assigns it to the Co2 field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetCo2(v int32) {
	o.Co2 = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetCurrent() int32 {
	if o == nil || isNil(o.Current) {
		var ret int32
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetCurrentOk() (*int32, bool) {
	if o == nil || isNil(o.Current) {
    return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasCurrent() bool {
	if o != nil && !isNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given int32 and assigns it to the Current field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetCurrent(v int32) {
	o.Current = &v
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetDoor() int32 {
	if o == nil || isNil(o.Door) {
		var ret int32
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetDoorOk() (*int32, bool) {
	if o == nil || isNil(o.Door) {
    return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasDoor() bool {
	if o != nil && !isNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given int32 and assigns it to the Door field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetDoor(v int32) {
	o.Door = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetFrequency() int32 {
	if o == nil || isNil(o.Frequency) {
		var ret int32
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetFrequencyOk() (*int32, bool) {
	if o == nil || isNil(o.Frequency) {
    return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasFrequency() bool {
	if o != nil && !isNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given int32 and assigns it to the Frequency field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetFrequency(v int32) {
	o.Frequency = &v
}

// GetHumidity returns the Humidity field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetHumidity() int32 {
	if o == nil || isNil(o.Humidity) {
		var ret int32
		return ret
	}
	return *o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetHumidityOk() (*int32, bool) {
	if o == nil || isNil(o.Humidity) {
    return nil, false
	}
	return o.Humidity, true
}

// HasHumidity returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasHumidity() bool {
	if o != nil && !isNil(o.Humidity) {
		return true
	}

	return false
}

// SetHumidity gets a reference to the given int32 and assigns it to the Humidity field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetHumidity(v int32) {
	o.Humidity = &v
}

// GetIndoorAirQuality returns the IndoorAirQuality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetIndoorAirQuality() int32 {
	if o == nil || isNil(o.IndoorAirQuality) {
		var ret int32
		return ret
	}
	return *o.IndoorAirQuality
}

// GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetIndoorAirQualityOk() (*int32, bool) {
	if o == nil || isNil(o.IndoorAirQuality) {
    return nil, false
	}
	return o.IndoorAirQuality, true
}

// HasIndoorAirQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasIndoorAirQuality() bool {
	if o != nil && !isNil(o.IndoorAirQuality) {
		return true
	}

	return false
}

// SetIndoorAirQuality gets a reference to the given int32 and assigns it to the IndoorAirQuality field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetIndoorAirQuality(v int32) {
	o.IndoorAirQuality = &v
}

// GetNoise returns the Noise field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetNoise() NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise {
	if o == nil || isNil(o.Noise) {
		var ret NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise
		return ret
	}
	return *o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetNoiseOk() (*NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise, bool) {
	if o == nil || isNil(o.Noise) {
    return nil, false
	}
	return o.Noise, true
}

// HasNoise returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasNoise() bool {
	if o != nil && !isNil(o.Noise) {
		return true
	}

	return false
}

// SetNoise gets a reference to the given NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise and assigns it to the Noise field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetNoise(v NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) {
	o.Noise = &v
}

// GetPm25 returns the Pm25 field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetPm25() int32 {
	if o == nil || isNil(o.Pm25) {
		var ret int32
		return ret
	}
	return *o.Pm25
}

// GetPm25Ok returns a tuple with the Pm25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetPm25Ok() (*int32, bool) {
	if o == nil || isNil(o.Pm25) {
    return nil, false
	}
	return o.Pm25, true
}

// HasPm25 returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasPm25() bool {
	if o != nil && !isNil(o.Pm25) {
		return true
	}

	return false
}

// SetPm25 gets a reference to the given int32 and assigns it to the Pm25 field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetPm25(v int32) {
	o.Pm25 = &v
}

// GetPowerFactor returns the PowerFactor field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetPowerFactor() int32 {
	if o == nil || isNil(o.PowerFactor) {
		var ret int32
		return ret
	}
	return *o.PowerFactor
}

// GetPowerFactorOk returns a tuple with the PowerFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetPowerFactorOk() (*int32, bool) {
	if o == nil || isNil(o.PowerFactor) {
    return nil, false
	}
	return o.PowerFactor, true
}

// HasPowerFactor returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasPowerFactor() bool {
	if o != nil && !isNil(o.PowerFactor) {
		return true
	}

	return false
}

// SetPowerFactor gets a reference to the given int32 and assigns it to the PowerFactor field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetPowerFactor(v int32) {
	o.PowerFactor = &v
}

// GetRealPower returns the RealPower field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetRealPower() int32 {
	if o == nil || isNil(o.RealPower) {
		var ret int32
		return ret
	}
	return *o.RealPower
}

// GetRealPowerOk returns a tuple with the RealPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetRealPowerOk() (*int32, bool) {
	if o == nil || isNil(o.RealPower) {
    return nil, false
	}
	return o.RealPower, true
}

// HasRealPower returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasRealPower() bool {
	if o != nil && !isNil(o.RealPower) {
		return true
	}

	return false
}

// SetRealPower gets a reference to the given int32 and assigns it to the RealPower field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetRealPower(v int32) {
	o.RealPower = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetTemperature() int32 {
	if o == nil || isNil(o.Temperature) {
		var ret int32
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetTemperatureOk() (*int32, bool) {
	if o == nil || isNil(o.Temperature) {
    return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasTemperature() bool {
	if o != nil && !isNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given int32 and assigns it to the Temperature field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetTemperature(v int32) {
	o.Temperature = &v
}

// GetTvoc returns the Tvoc field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetTvoc() int32 {
	if o == nil || isNil(o.Tvoc) {
		var ret int32
		return ret
	}
	return *o.Tvoc
}

// GetTvocOk returns a tuple with the Tvoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetTvocOk() (*int32, bool) {
	if o == nil || isNil(o.Tvoc) {
    return nil, false
	}
	return o.Tvoc, true
}

// HasTvoc returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasTvoc() bool {
	if o != nil && !isNil(o.Tvoc) {
		return true
	}

	return false
}

// SetTvoc gets a reference to the given int32 and assigns it to the Tvoc field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetTvoc(v int32) {
	o.Tvoc = &v
}

// GetUpstreamPower returns the UpstreamPower field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetUpstreamPower() int32 {
	if o == nil || isNil(o.UpstreamPower) {
		var ret int32
		return ret
	}
	return *o.UpstreamPower
}

// GetUpstreamPowerOk returns a tuple with the UpstreamPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetUpstreamPowerOk() (*int32, bool) {
	if o == nil || isNil(o.UpstreamPower) {
    return nil, false
	}
	return o.UpstreamPower, true
}

// HasUpstreamPower returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasUpstreamPower() bool {
	if o != nil && !isNil(o.UpstreamPower) {
		return true
	}

	return false
}

// SetUpstreamPower gets a reference to the given int32 and assigns it to the UpstreamPower field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetUpstreamPower(v int32) {
	o.UpstreamPower = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetVoltage() int32 {
	if o == nil || isNil(o.Voltage) {
		var ret int32
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetVoltageOk() (*int32, bool) {
	if o == nil || isNil(o.Voltage) {
    return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasVoltage() bool {
	if o != nil && !isNil(o.Voltage) {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given int32 and assigns it to the Voltage field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetVoltage(v int32) {
	o.Voltage = &v
}

// GetWater returns the Water field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetWater() int32 {
	if o == nil || isNil(o.Water) {
		var ret int32
		return ret
	}
	return *o.Water
}

// GetWaterOk returns a tuple with the Water field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) GetWaterOk() (*int32, bool) {
	if o == nil || isNil(o.Water) {
    return nil, false
	}
	return o.Water, true
}

// HasWater returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) HasWater() bool {
	if o != nil && !isNil(o.Water) {
		return true
	}

	return false
}

// SetWater gets a reference to the given int32 and assigns it to the Water field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) SetWater(v int32) {
	o.Water = &v
}

func (o NetworksNetworkIdSensorAlertsOverviewByMetricCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApparentPower) {
		toSerialize["apparentPower"] = o.ApparentPower
	}
	if !isNil(o.Co2) {
		toSerialize["co2"] = o.Co2
	}
	if !isNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !isNil(o.Door) {
		toSerialize["door"] = o.Door
	}
	if !isNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
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
	if !isNil(o.PowerFactor) {
		toSerialize["powerFactor"] = o.PowerFactor
	}
	if !isNil(o.RealPower) {
		toSerialize["realPower"] = o.RealPower
	}
	if !isNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !isNil(o.Tvoc) {
		toSerialize["tvoc"] = o.Tvoc
	}
	if !isNil(o.UpstreamPower) {
		toSerialize["upstreamPower"] = o.UpstreamPower
	}
	if !isNil(o.Voltage) {
		toSerialize["voltage"] = o.Voltage
	}
	if !isNil(o.Water) {
		toSerialize["water"] = o.Water
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts struct {
	value *NetworksNetworkIdSensorAlertsOverviewByMetricCounts
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts) Get() *NetworksNetworkIdSensorAlertsOverviewByMetricCounts {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts) Set(val *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts(val *NetworksNetworkIdSensorAlertsOverviewByMetricCounts) *NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts {
	return &NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsOverviewByMetricCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


