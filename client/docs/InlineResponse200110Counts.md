# InlineResponse200110Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApparentPower** | Pointer to **int32** | Number of sensors that are currently alerting due to apparent power readings | [optional] 
**Co2** | Pointer to **int32** | Number of sensors that are currently alerting due to CO2 readings | [optional] 
**Current** | Pointer to **int32** | Number of sensors that are currently alerting due to electrical current readings | [optional] 
**Door** | Pointer to **int32** | Number of sensors that are currently alerting due to an open door | [optional] 
**Frequency** | Pointer to **int32** | Number of sensors that are currently alerting due to frequency readings | [optional] 
**Humidity** | Pointer to **int32** | Number of sensors that are currently alerting due to humidity readings | [optional] 
**IndoorAirQuality** | Pointer to **int32** | Number of sensors that are currently alerting due to indoor air quality readings | [optional] 
**Noise** | Pointer to [**InlineResponse200110CountsNoise**](InlineResponse200110CountsNoise.md) |  | [optional] 
**Pm25** | Pointer to **int32** | Number of sensors that are currently alerting due to PM2.5 readings | [optional] 
**PowerFactor** | Pointer to **int32** | Number of sensors that are currently alerting due to power factor readings | [optional] 
**RealPower** | Pointer to **int32** | Number of sensors that are currently alerting due to real power readings | [optional] 
**Temperature** | Pointer to **int32** | Number of sensors that are currently alerting due to temperature readings | [optional] 
**Tvoc** | Pointer to **int32** | Number of sensors that are currently alerting due to TVOC readings | [optional] 
**UpstreamPower** | Pointer to **int32** | Number of sensors that are currently alerting due to an upstream power outage | [optional] 
**Voltage** | Pointer to **int32** | Number of sensors that are currently alerting due to voltage readings | [optional] 
**Water** | Pointer to **int32** | Number of sensors that are currently alerting due to the presence of water | [optional] 

## Methods

### NewInlineResponse200110Counts

`func NewInlineResponse200110Counts() *InlineResponse200110Counts`

NewInlineResponse200110Counts instantiates a new InlineResponse200110Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110CountsWithDefaults

`func NewInlineResponse200110CountsWithDefaults() *InlineResponse200110Counts`

NewInlineResponse200110CountsWithDefaults instantiates a new InlineResponse200110Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApparentPower

`func (o *InlineResponse200110Counts) GetApparentPower() int32`

GetApparentPower returns the ApparentPower field if non-nil, zero value otherwise.

### GetApparentPowerOk

`func (o *InlineResponse200110Counts) GetApparentPowerOk() (*int32, bool)`

GetApparentPowerOk returns a tuple with the ApparentPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApparentPower

`func (o *InlineResponse200110Counts) SetApparentPower(v int32)`

SetApparentPower sets ApparentPower field to given value.

### HasApparentPower

`func (o *InlineResponse200110Counts) HasApparentPower() bool`

HasApparentPower returns a boolean if a field has been set.

### GetCo2

`func (o *InlineResponse200110Counts) GetCo2() int32`

GetCo2 returns the Co2 field if non-nil, zero value otherwise.

### GetCo2Ok

`func (o *InlineResponse200110Counts) GetCo2Ok() (*int32, bool)`

GetCo2Ok returns a tuple with the Co2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2

`func (o *InlineResponse200110Counts) SetCo2(v int32)`

SetCo2 sets Co2 field to given value.

### HasCo2

`func (o *InlineResponse200110Counts) HasCo2() bool`

HasCo2 returns a boolean if a field has been set.

### GetCurrent

`func (o *InlineResponse200110Counts) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *InlineResponse200110Counts) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *InlineResponse200110Counts) SetCurrent(v int32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *InlineResponse200110Counts) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetDoor

`func (o *InlineResponse200110Counts) GetDoor() int32`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *InlineResponse200110Counts) GetDoorOk() (*int32, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *InlineResponse200110Counts) SetDoor(v int32)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *InlineResponse200110Counts) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetFrequency

`func (o *InlineResponse200110Counts) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *InlineResponse200110Counts) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *InlineResponse200110Counts) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *InlineResponse200110Counts) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHumidity

`func (o *InlineResponse200110Counts) GetHumidity() int32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *InlineResponse200110Counts) GetHumidityOk() (*int32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *InlineResponse200110Counts) SetHumidity(v int32)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *InlineResponse200110Counts) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *InlineResponse200110Counts) GetIndoorAirQuality() int32`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *InlineResponse200110Counts) GetIndoorAirQualityOk() (*int32, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *InlineResponse200110Counts) SetIndoorAirQuality(v int32)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *InlineResponse200110Counts) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *InlineResponse200110Counts) GetNoise() InlineResponse200110CountsNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *InlineResponse200110Counts) GetNoiseOk() (*InlineResponse200110CountsNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *InlineResponse200110Counts) SetNoise(v InlineResponse200110CountsNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *InlineResponse200110Counts) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *InlineResponse200110Counts) GetPm25() int32`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *InlineResponse200110Counts) GetPm25Ok() (*int32, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *InlineResponse200110Counts) SetPm25(v int32)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *InlineResponse200110Counts) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetPowerFactor

`func (o *InlineResponse200110Counts) GetPowerFactor() int32`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *InlineResponse200110Counts) GetPowerFactorOk() (*int32, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *InlineResponse200110Counts) SetPowerFactor(v int32)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *InlineResponse200110Counts) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### GetRealPower

`func (o *InlineResponse200110Counts) GetRealPower() int32`

GetRealPower returns the RealPower field if non-nil, zero value otherwise.

### GetRealPowerOk

`func (o *InlineResponse200110Counts) GetRealPowerOk() (*int32, bool)`

GetRealPowerOk returns a tuple with the RealPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealPower

`func (o *InlineResponse200110Counts) SetRealPower(v int32)`

SetRealPower sets RealPower field to given value.

### HasRealPower

`func (o *InlineResponse200110Counts) HasRealPower() bool`

HasRealPower returns a boolean if a field has been set.

### GetTemperature

`func (o *InlineResponse200110Counts) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *InlineResponse200110Counts) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *InlineResponse200110Counts) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *InlineResponse200110Counts) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *InlineResponse200110Counts) GetTvoc() int32`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *InlineResponse200110Counts) GetTvocOk() (*int32, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *InlineResponse200110Counts) SetTvoc(v int32)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *InlineResponse200110Counts) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetUpstreamPower

`func (o *InlineResponse200110Counts) GetUpstreamPower() int32`

GetUpstreamPower returns the UpstreamPower field if non-nil, zero value otherwise.

### GetUpstreamPowerOk

`func (o *InlineResponse200110Counts) GetUpstreamPowerOk() (*int32, bool)`

GetUpstreamPowerOk returns a tuple with the UpstreamPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamPower

`func (o *InlineResponse200110Counts) SetUpstreamPower(v int32)`

SetUpstreamPower sets UpstreamPower field to given value.

### HasUpstreamPower

`func (o *InlineResponse200110Counts) HasUpstreamPower() bool`

HasUpstreamPower returns a boolean if a field has been set.

### GetVoltage

`func (o *InlineResponse200110Counts) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *InlineResponse200110Counts) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *InlineResponse200110Counts) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *InlineResponse200110Counts) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetWater

`func (o *InlineResponse200110Counts) GetWater() int32`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *InlineResponse200110Counts) GetWaterOk() (*int32, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *InlineResponse200110Counts) SetWater(v int32)`

SetWater sets Water field to given value.

### HasWater

`func (o *InlineResponse200110Counts) HasWater() bool`

HasWater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

