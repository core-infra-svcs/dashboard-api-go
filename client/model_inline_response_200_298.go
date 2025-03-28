/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200298 struct for InlineResponse200298
type InlineResponse200298 struct {
	// Serial number of the sensor that took the reading.
	Serial *string `json:"serial,omitempty"`
	Network *OrganizationsOrganizationIdSensorReadingsHistoryNetwork `json:"network,omitempty"`
	// Time at which the reading occurred, in ISO8601 format.
	Ts *string `json:"ts,omitempty"`
	// Type of sensor reading.
	Metric *string `json:"metric,omitempty"`
	ApparentPower *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower `json:"apparentPower,omitempty"`
	Battery *OrganizationsOrganizationIdSensorReadingsHistoryBattery `json:"battery,omitempty"`
	Button *OrganizationsOrganizationIdSensorReadingsHistoryButton `json:"button,omitempty"`
	Co2 *OrganizationsOrganizationIdSensorReadingsHistoryCo2 `json:"co2,omitempty"`
	Current *OrganizationsOrganizationIdSensorReadingsHistoryCurrent `json:"current,omitempty"`
	Door *OrganizationsOrganizationIdSensorReadingsHistoryDoor `json:"door,omitempty"`
	DownstreamPower *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower `json:"downstreamPower,omitempty"`
	Frequency *OrganizationsOrganizationIdSensorReadingsHistoryFrequency `json:"frequency,omitempty"`
	Humidity *OrganizationsOrganizationIdSensorReadingsHistoryHumidity `json:"humidity,omitempty"`
	IndoorAirQuality *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality `json:"indoorAirQuality,omitempty"`
	Noise *OrganizationsOrganizationIdSensorReadingsHistoryNoise `json:"noise,omitempty"`
	Pm25 *OrganizationsOrganizationIdSensorReadingsHistoryPm25 `json:"pm25,omitempty"`
	PowerFactor *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor `json:"powerFactor,omitempty"`
	RealPower *OrganizationsOrganizationIdSensorReadingsHistoryRealPower `json:"realPower,omitempty"`
	RemoteLockoutSwitch *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch `json:"remoteLockoutSwitch,omitempty"`
	Temperature *OrganizationsOrganizationIdSensorReadingsHistoryTemperature `json:"temperature,omitempty"`
	Tvoc *OrganizationsOrganizationIdSensorReadingsHistoryTvoc `json:"tvoc,omitempty"`
	Voltage *OrganizationsOrganizationIdSensorReadingsHistoryVoltage `json:"voltage,omitempty"`
	Water *OrganizationsOrganizationIdSensorReadingsHistoryWater `json:"water,omitempty"`
}

// NewInlineResponse200298 instantiates a new InlineResponse200298 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200298() *InlineResponse200298 {
	this := InlineResponse200298{}
	return &this
}

// NewInlineResponse200298WithDefaults instantiates a new InlineResponse200298 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200298WithDefaults() *InlineResponse200298 {
	this := InlineResponse200298{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200298) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200298) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200298) SetSerial(v string) {
	o.Serial = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200298) GetNetwork() OrganizationsOrganizationIdSensorReadingsHistoryNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetNetworkOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200298) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryNetwork and assigns it to the Network field.
func (o *InlineResponse200298) SetNetwork(v OrganizationsOrganizationIdSensorReadingsHistoryNetwork) {
	o.Network = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200298) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200298) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse200298) SetTs(v string) {
	o.Ts = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *InlineResponse200298) GetMetric() string {
	if o == nil || isNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetMetricOk() (*string, bool) {
	if o == nil || isNil(o.Metric) {
    return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *InlineResponse200298) HasMetric() bool {
	if o != nil && !isNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *InlineResponse200298) SetMetric(v string) {
	o.Metric = &v
}

// GetApparentPower returns the ApparentPower field value if set, zero value otherwise.
func (o *InlineResponse200298) GetApparentPower() OrganizationsOrganizationIdSensorReadingsHistoryApparentPower {
	if o == nil || isNil(o.ApparentPower) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryApparentPower
		return ret
	}
	return *o.ApparentPower
}

// GetApparentPowerOk returns a tuple with the ApparentPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetApparentPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryApparentPower, bool) {
	if o == nil || isNil(o.ApparentPower) {
    return nil, false
	}
	return o.ApparentPower, true
}

// HasApparentPower returns a boolean if a field has been set.
func (o *InlineResponse200298) HasApparentPower() bool {
	if o != nil && !isNil(o.ApparentPower) {
		return true
	}

	return false
}

// SetApparentPower gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryApparentPower and assigns it to the ApparentPower field.
func (o *InlineResponse200298) SetApparentPower(v OrganizationsOrganizationIdSensorReadingsHistoryApparentPower) {
	o.ApparentPower = &v
}

// GetBattery returns the Battery field value if set, zero value otherwise.
func (o *InlineResponse200298) GetBattery() OrganizationsOrganizationIdSensorReadingsHistoryBattery {
	if o == nil || isNil(o.Battery) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryBattery
		return ret
	}
	return *o.Battery
}

// GetBatteryOk returns a tuple with the Battery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetBatteryOk() (*OrganizationsOrganizationIdSensorReadingsHistoryBattery, bool) {
	if o == nil || isNil(o.Battery) {
    return nil, false
	}
	return o.Battery, true
}

// HasBattery returns a boolean if a field has been set.
func (o *InlineResponse200298) HasBattery() bool {
	if o != nil && !isNil(o.Battery) {
		return true
	}

	return false
}

// SetBattery gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryBattery and assigns it to the Battery field.
func (o *InlineResponse200298) SetBattery(v OrganizationsOrganizationIdSensorReadingsHistoryBattery) {
	o.Battery = &v
}

// GetButton returns the Button field value if set, zero value otherwise.
func (o *InlineResponse200298) GetButton() OrganizationsOrganizationIdSensorReadingsHistoryButton {
	if o == nil || isNil(o.Button) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryButton
		return ret
	}
	return *o.Button
}

// GetButtonOk returns a tuple with the Button field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetButtonOk() (*OrganizationsOrganizationIdSensorReadingsHistoryButton, bool) {
	if o == nil || isNil(o.Button) {
    return nil, false
	}
	return o.Button, true
}

// HasButton returns a boolean if a field has been set.
func (o *InlineResponse200298) HasButton() bool {
	if o != nil && !isNil(o.Button) {
		return true
	}

	return false
}

// SetButton gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryButton and assigns it to the Button field.
func (o *InlineResponse200298) SetButton(v OrganizationsOrganizationIdSensorReadingsHistoryButton) {
	o.Button = &v
}

// GetCo2 returns the Co2 field value if set, zero value otherwise.
func (o *InlineResponse200298) GetCo2() OrganizationsOrganizationIdSensorReadingsHistoryCo2 {
	if o == nil || isNil(o.Co2) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryCo2
		return ret
	}
	return *o.Co2
}

// GetCo2Ok returns a tuple with the Co2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetCo2Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryCo2, bool) {
	if o == nil || isNil(o.Co2) {
    return nil, false
	}
	return o.Co2, true
}

// HasCo2 returns a boolean if a field has been set.
func (o *InlineResponse200298) HasCo2() bool {
	if o != nil && !isNil(o.Co2) {
		return true
	}

	return false
}

// SetCo2 gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryCo2 and assigns it to the Co2 field.
func (o *InlineResponse200298) SetCo2(v OrganizationsOrganizationIdSensorReadingsHistoryCo2) {
	o.Co2 = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *InlineResponse200298) GetCurrent() OrganizationsOrganizationIdSensorReadingsHistoryCurrent {
	if o == nil || isNil(o.Current) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryCurrent
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetCurrentOk() (*OrganizationsOrganizationIdSensorReadingsHistoryCurrent, bool) {
	if o == nil || isNil(o.Current) {
    return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *InlineResponse200298) HasCurrent() bool {
	if o != nil && !isNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryCurrent and assigns it to the Current field.
func (o *InlineResponse200298) SetCurrent(v OrganizationsOrganizationIdSensorReadingsHistoryCurrent) {
	o.Current = &v
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *InlineResponse200298) GetDoor() OrganizationsOrganizationIdSensorReadingsHistoryDoor {
	if o == nil || isNil(o.Door) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryDoor
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetDoorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDoor, bool) {
	if o == nil || isNil(o.Door) {
    return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *InlineResponse200298) HasDoor() bool {
	if o != nil && !isNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryDoor and assigns it to the Door field.
func (o *InlineResponse200298) SetDoor(v OrganizationsOrganizationIdSensorReadingsHistoryDoor) {
	o.Door = &v
}

// GetDownstreamPower returns the DownstreamPower field value if set, zero value otherwise.
func (o *InlineResponse200298) GetDownstreamPower() OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower {
	if o == nil || isNil(o.DownstreamPower) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower
		return ret
	}
	return *o.DownstreamPower
}

// GetDownstreamPowerOk returns a tuple with the DownstreamPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetDownstreamPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower, bool) {
	if o == nil || isNil(o.DownstreamPower) {
    return nil, false
	}
	return o.DownstreamPower, true
}

// HasDownstreamPower returns a boolean if a field has been set.
func (o *InlineResponse200298) HasDownstreamPower() bool {
	if o != nil && !isNil(o.DownstreamPower) {
		return true
	}

	return false
}

// SetDownstreamPower gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower and assigns it to the DownstreamPower field.
func (o *InlineResponse200298) SetDownstreamPower(v OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) {
	o.DownstreamPower = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *InlineResponse200298) GetFrequency() OrganizationsOrganizationIdSensorReadingsHistoryFrequency {
	if o == nil || isNil(o.Frequency) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryFrequency
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetFrequencyOk() (*OrganizationsOrganizationIdSensorReadingsHistoryFrequency, bool) {
	if o == nil || isNil(o.Frequency) {
    return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *InlineResponse200298) HasFrequency() bool {
	if o != nil && !isNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryFrequency and assigns it to the Frequency field.
func (o *InlineResponse200298) SetFrequency(v OrganizationsOrganizationIdSensorReadingsHistoryFrequency) {
	o.Frequency = &v
}

// GetHumidity returns the Humidity field value if set, zero value otherwise.
func (o *InlineResponse200298) GetHumidity() OrganizationsOrganizationIdSensorReadingsHistoryHumidity {
	if o == nil || isNil(o.Humidity) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryHumidity
		return ret
	}
	return *o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetHumidityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryHumidity, bool) {
	if o == nil || isNil(o.Humidity) {
    return nil, false
	}
	return o.Humidity, true
}

// HasHumidity returns a boolean if a field has been set.
func (o *InlineResponse200298) HasHumidity() bool {
	if o != nil && !isNil(o.Humidity) {
		return true
	}

	return false
}

// SetHumidity gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryHumidity and assigns it to the Humidity field.
func (o *InlineResponse200298) SetHumidity(v OrganizationsOrganizationIdSensorReadingsHistoryHumidity) {
	o.Humidity = &v
}

// GetIndoorAirQuality returns the IndoorAirQuality field value if set, zero value otherwise.
func (o *InlineResponse200298) GetIndoorAirQuality() OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality {
	if o == nil || isNil(o.IndoorAirQuality) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality
		return ret
	}
	return *o.IndoorAirQuality
}

// GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetIndoorAirQualityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality, bool) {
	if o == nil || isNil(o.IndoorAirQuality) {
    return nil, false
	}
	return o.IndoorAirQuality, true
}

// HasIndoorAirQuality returns a boolean if a field has been set.
func (o *InlineResponse200298) HasIndoorAirQuality() bool {
	if o != nil && !isNil(o.IndoorAirQuality) {
		return true
	}

	return false
}

// SetIndoorAirQuality gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality and assigns it to the IndoorAirQuality field.
func (o *InlineResponse200298) SetIndoorAirQuality(v OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) {
	o.IndoorAirQuality = &v
}

// GetNoise returns the Noise field value if set, zero value otherwise.
func (o *InlineResponse200298) GetNoise() OrganizationsOrganizationIdSensorReadingsHistoryNoise {
	if o == nil || isNil(o.Noise) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryNoise
		return ret
	}
	return *o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetNoiseOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNoise, bool) {
	if o == nil || isNil(o.Noise) {
    return nil, false
	}
	return o.Noise, true
}

// HasNoise returns a boolean if a field has been set.
func (o *InlineResponse200298) HasNoise() bool {
	if o != nil && !isNil(o.Noise) {
		return true
	}

	return false
}

// SetNoise gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryNoise and assigns it to the Noise field.
func (o *InlineResponse200298) SetNoise(v OrganizationsOrganizationIdSensorReadingsHistoryNoise) {
	o.Noise = &v
}

// GetPm25 returns the Pm25 field value if set, zero value otherwise.
func (o *InlineResponse200298) GetPm25() OrganizationsOrganizationIdSensorReadingsHistoryPm25 {
	if o == nil || isNil(o.Pm25) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryPm25
		return ret
	}
	return *o.Pm25
}

// GetPm25Ok returns a tuple with the Pm25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetPm25Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm25, bool) {
	if o == nil || isNil(o.Pm25) {
    return nil, false
	}
	return o.Pm25, true
}

// HasPm25 returns a boolean if a field has been set.
func (o *InlineResponse200298) HasPm25() bool {
	if o != nil && !isNil(o.Pm25) {
		return true
	}

	return false
}

// SetPm25 gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryPm25 and assigns it to the Pm25 field.
func (o *InlineResponse200298) SetPm25(v OrganizationsOrganizationIdSensorReadingsHistoryPm25) {
	o.Pm25 = &v
}

// GetPowerFactor returns the PowerFactor field value if set, zero value otherwise.
func (o *InlineResponse200298) GetPowerFactor() OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor {
	if o == nil || isNil(o.PowerFactor) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor
		return ret
	}
	return *o.PowerFactor
}

// GetPowerFactorOk returns a tuple with the PowerFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetPowerFactorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor, bool) {
	if o == nil || isNil(o.PowerFactor) {
    return nil, false
	}
	return o.PowerFactor, true
}

// HasPowerFactor returns a boolean if a field has been set.
func (o *InlineResponse200298) HasPowerFactor() bool {
	if o != nil && !isNil(o.PowerFactor) {
		return true
	}

	return false
}

// SetPowerFactor gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor and assigns it to the PowerFactor field.
func (o *InlineResponse200298) SetPowerFactor(v OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) {
	o.PowerFactor = &v
}

// GetRealPower returns the RealPower field value if set, zero value otherwise.
func (o *InlineResponse200298) GetRealPower() OrganizationsOrganizationIdSensorReadingsHistoryRealPower {
	if o == nil || isNil(o.RealPower) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryRealPower
		return ret
	}
	return *o.RealPower
}

// GetRealPowerOk returns a tuple with the RealPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetRealPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRealPower, bool) {
	if o == nil || isNil(o.RealPower) {
    return nil, false
	}
	return o.RealPower, true
}

// HasRealPower returns a boolean if a field has been set.
func (o *InlineResponse200298) HasRealPower() bool {
	if o != nil && !isNil(o.RealPower) {
		return true
	}

	return false
}

// SetRealPower gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryRealPower and assigns it to the RealPower field.
func (o *InlineResponse200298) SetRealPower(v OrganizationsOrganizationIdSensorReadingsHistoryRealPower) {
	o.RealPower = &v
}

// GetRemoteLockoutSwitch returns the RemoteLockoutSwitch field value if set, zero value otherwise.
func (o *InlineResponse200298) GetRemoteLockoutSwitch() OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch {
	if o == nil || isNil(o.RemoteLockoutSwitch) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch
		return ret
	}
	return *o.RemoteLockoutSwitch
}

// GetRemoteLockoutSwitchOk returns a tuple with the RemoteLockoutSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetRemoteLockoutSwitchOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch, bool) {
	if o == nil || isNil(o.RemoteLockoutSwitch) {
    return nil, false
	}
	return o.RemoteLockoutSwitch, true
}

// HasRemoteLockoutSwitch returns a boolean if a field has been set.
func (o *InlineResponse200298) HasRemoteLockoutSwitch() bool {
	if o != nil && !isNil(o.RemoteLockoutSwitch) {
		return true
	}

	return false
}

// SetRemoteLockoutSwitch gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch and assigns it to the RemoteLockoutSwitch field.
func (o *InlineResponse200298) SetRemoteLockoutSwitch(v OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) {
	o.RemoteLockoutSwitch = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *InlineResponse200298) GetTemperature() OrganizationsOrganizationIdSensorReadingsHistoryTemperature {
	if o == nil || isNil(o.Temperature) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryTemperature
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTemperature, bool) {
	if o == nil || isNil(o.Temperature) {
    return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *InlineResponse200298) HasTemperature() bool {
	if o != nil && !isNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryTemperature and assigns it to the Temperature field.
func (o *InlineResponse200298) SetTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryTemperature) {
	o.Temperature = &v
}

// GetTvoc returns the Tvoc field value if set, zero value otherwise.
func (o *InlineResponse200298) GetTvoc() OrganizationsOrganizationIdSensorReadingsHistoryTvoc {
	if o == nil || isNil(o.Tvoc) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryTvoc
		return ret
	}
	return *o.Tvoc
}

// GetTvocOk returns a tuple with the Tvoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetTvocOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTvoc, bool) {
	if o == nil || isNil(o.Tvoc) {
    return nil, false
	}
	return o.Tvoc, true
}

// HasTvoc returns a boolean if a field has been set.
func (o *InlineResponse200298) HasTvoc() bool {
	if o != nil && !isNil(o.Tvoc) {
		return true
	}

	return false
}

// SetTvoc gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryTvoc and assigns it to the Tvoc field.
func (o *InlineResponse200298) SetTvoc(v OrganizationsOrganizationIdSensorReadingsHistoryTvoc) {
	o.Tvoc = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *InlineResponse200298) GetVoltage() OrganizationsOrganizationIdSensorReadingsHistoryVoltage {
	if o == nil || isNil(o.Voltage) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryVoltage
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetVoltageOk() (*OrganizationsOrganizationIdSensorReadingsHistoryVoltage, bool) {
	if o == nil || isNil(o.Voltage) {
    return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *InlineResponse200298) HasVoltage() bool {
	if o != nil && !isNil(o.Voltage) {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryVoltage and assigns it to the Voltage field.
func (o *InlineResponse200298) SetVoltage(v OrganizationsOrganizationIdSensorReadingsHistoryVoltage) {
	o.Voltage = &v
}

// GetWater returns the Water field value if set, zero value otherwise.
func (o *InlineResponse200298) GetWater() OrganizationsOrganizationIdSensorReadingsHistoryWater {
	if o == nil || isNil(o.Water) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryWater
		return ret
	}
	return *o.Water
}

// GetWaterOk returns a tuple with the Water field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetWaterOk() (*OrganizationsOrganizationIdSensorReadingsHistoryWater, bool) {
	if o == nil || isNil(o.Water) {
    return nil, false
	}
	return o.Water, true
}

// HasWater returns a boolean if a field has been set.
func (o *InlineResponse200298) HasWater() bool {
	if o != nil && !isNil(o.Water) {
		return true
	}

	return false
}

// SetWater gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryWater and assigns it to the Water field.
func (o *InlineResponse200298) SetWater(v OrganizationsOrganizationIdSensorReadingsHistoryWater) {
	o.Water = &v
}

func (o InlineResponse200298) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ApparentPower) {
		toSerialize["apparentPower"] = o.ApparentPower
	}
	if !isNil(o.Battery) {
		toSerialize["battery"] = o.Battery
	}
	if !isNil(o.Button) {
		toSerialize["button"] = o.Button
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
	if !isNil(o.DownstreamPower) {
		toSerialize["downstreamPower"] = o.DownstreamPower
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
	if !isNil(o.RemoteLockoutSwitch) {
		toSerialize["remoteLockoutSwitch"] = o.RemoteLockoutSwitch
	}
	if !isNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !isNil(o.Tvoc) {
		toSerialize["tvoc"] = o.Tvoc
	}
	if !isNil(o.Voltage) {
		toSerialize["voltage"] = o.Voltage
	}
	if !isNil(o.Water) {
		toSerialize["water"] = o.Water
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200298 struct {
	value *InlineResponse200298
	isSet bool
}

func (v NullableInlineResponse200298) Get() *InlineResponse200298 {
	return v.value
}

func (v *NullableInlineResponse200298) Set(val *InlineResponse200298) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200298) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200298) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200298(val *InlineResponse200298) *NullableInlineResponse200298 {
	return &NullableInlineResponse200298{value: val, isSet: true}
}

func (v NullableInlineResponse200298) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200298) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


