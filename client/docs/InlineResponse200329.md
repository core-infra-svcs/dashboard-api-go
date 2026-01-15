# InlineResponse200329

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number of the sensor that took the reading. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryNetwork**](OrganizationsOrganizationIdSensorReadingsHistoryNetwork.md) |  | [optional] 
**Ts** | Pointer to **string** | Time at which the reading occurred, in ISO8601 format. | [optional] 
**Metric** | Pointer to **string** | Type of sensor reading. | [optional] 
**ApparentPower** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower**](OrganizationsOrganizationIdSensorReadingsHistoryApparentPower.md) |  | [optional] 
**Battery** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery**](OrganizationsOrganizationIdSensorReadingsHistoryBattery.md) |  | [optional] 
**Button** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryButton**](OrganizationsOrganizationIdSensorReadingsHistoryButton.md) |  | [optional] 
**Co2** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2**](OrganizationsOrganizationIdSensorReadingsHistoryCo2.md) |  | [optional] 
**Current** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent**](OrganizationsOrganizationIdSensorReadingsHistoryCurrent.md) |  | [optional] 
**Door** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor**](OrganizationsOrganizationIdSensorReadingsHistoryDoor.md) |  | [optional] 
**DownstreamPower** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower**](OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower.md) |  | [optional] 
**Frequency** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency**](OrganizationsOrganizationIdSensorReadingsHistoryFrequency.md) |  | [optional] 
**Humidity** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity**](OrganizationsOrganizationIdSensorReadingsHistoryHumidity.md) |  | [optional] 
**IndoorAirQuality** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality**](OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality.md) |  | [optional] 
**Noise** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise**](OrganizationsOrganizationIdSensorReadingsHistoryNoise.md) |  | [optional] 
**No2** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryNo2**](OrganizationsOrganizationIdSensorReadingsHistoryNo2.md) |  | [optional] 
**O3** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryO3**](OrganizationsOrganizationIdSensorReadingsHistoryO3.md) |  | [optional] 
**Pm10** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryPm10**](OrganizationsOrganizationIdSensorReadingsHistoryPm10.md) |  | [optional] 
**Pm25** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25**](OrganizationsOrganizationIdSensorReadingsHistoryPm25.md) |  | [optional] 
**PowerFactor** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor**](OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor.md) |  | [optional] 
**RealPower** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryRealPower**](OrganizationsOrganizationIdSensorReadingsHistoryRealPower.md) |  | [optional] 
**RemoteLockoutSwitch** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch**](OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch.md) |  | [optional] 
**Temperature** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryTemperature**](OrganizationsOrganizationIdSensorReadingsHistoryTemperature.md) |  | [optional] 
**Tvoc** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc**](OrganizationsOrganizationIdSensorReadingsHistoryTvoc.md) |  | [optional] 
**Voltage** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage**](OrganizationsOrganizationIdSensorReadingsHistoryVoltage.md) |  | [optional] 
**Water** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryWater**](OrganizationsOrganizationIdSensorReadingsHistoryWater.md) |  | [optional] 
**RawTemperature** | Pointer to [**NullableOrganizationsOrganizationIdSensorReadingsHistoryRawTemperature**](OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature.md) |  | [optional] 

## Methods

### NewInlineResponse200329

`func NewInlineResponse200329() *InlineResponse200329`

NewInlineResponse200329 instantiates a new InlineResponse200329 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200329WithDefaults

`func NewInlineResponse200329WithDefaults() *InlineResponse200329`

NewInlineResponse200329WithDefaults instantiates a new InlineResponse200329 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200329) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200329) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200329) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200329) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200329) GetNetwork() OrganizationsOrganizationIdSensorReadingsHistoryNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200329) GetNetworkOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200329) SetNetwork(v OrganizationsOrganizationIdSensorReadingsHistoryNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200329) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse200329) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200329) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200329) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200329) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetMetric

`func (o *InlineResponse200329) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *InlineResponse200329) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *InlineResponse200329) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *InlineResponse200329) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetApparentPower

`func (o *InlineResponse200329) GetApparentPower() OrganizationsOrganizationIdSensorReadingsHistoryApparentPower`

GetApparentPower returns the ApparentPower field if non-nil, zero value otherwise.

### GetApparentPowerOk

`func (o *InlineResponse200329) GetApparentPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryApparentPower, bool)`

GetApparentPowerOk returns a tuple with the ApparentPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApparentPower

`func (o *InlineResponse200329) SetApparentPower(v OrganizationsOrganizationIdSensorReadingsHistoryApparentPower)`

SetApparentPower sets ApparentPower field to given value.

### HasApparentPower

`func (o *InlineResponse200329) HasApparentPower() bool`

HasApparentPower returns a boolean if a field has been set.

### SetApparentPowerNil

`func (o *InlineResponse200329) SetApparentPowerNil(b bool)`

 SetApparentPowerNil sets the value for ApparentPower to be an explicit nil

### UnsetApparentPower
`func (o *InlineResponse200329) UnsetApparentPower()`

UnsetApparentPower ensures that no value is present for ApparentPower, not even an explicit nil
### GetBattery

`func (o *InlineResponse200329) GetBattery() OrganizationsOrganizationIdSensorReadingsHistoryBattery`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *InlineResponse200329) GetBatteryOk() (*OrganizationsOrganizationIdSensorReadingsHistoryBattery, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *InlineResponse200329) SetBattery(v OrganizationsOrganizationIdSensorReadingsHistoryBattery)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *InlineResponse200329) HasBattery() bool`

HasBattery returns a boolean if a field has been set.

### SetBatteryNil

`func (o *InlineResponse200329) SetBatteryNil(b bool)`

 SetBatteryNil sets the value for Battery to be an explicit nil

### UnsetBattery
`func (o *InlineResponse200329) UnsetBattery()`

UnsetBattery ensures that no value is present for Battery, not even an explicit nil
### GetButton

`func (o *InlineResponse200329) GetButton() OrganizationsOrganizationIdSensorReadingsHistoryButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *InlineResponse200329) GetButtonOk() (*OrganizationsOrganizationIdSensorReadingsHistoryButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *InlineResponse200329) SetButton(v OrganizationsOrganizationIdSensorReadingsHistoryButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *InlineResponse200329) HasButton() bool`

HasButton returns a boolean if a field has been set.

### SetButtonNil

`func (o *InlineResponse200329) SetButtonNil(b bool)`

 SetButtonNil sets the value for Button to be an explicit nil

### UnsetButton
`func (o *InlineResponse200329) UnsetButton()`

UnsetButton ensures that no value is present for Button, not even an explicit nil
### GetCo2

`func (o *InlineResponse200329) GetCo2() OrganizationsOrganizationIdSensorReadingsHistoryCo2`

GetCo2 returns the Co2 field if non-nil, zero value otherwise.

### GetCo2Ok

`func (o *InlineResponse200329) GetCo2Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryCo2, bool)`

GetCo2Ok returns a tuple with the Co2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2

`func (o *InlineResponse200329) SetCo2(v OrganizationsOrganizationIdSensorReadingsHistoryCo2)`

SetCo2 sets Co2 field to given value.

### HasCo2

`func (o *InlineResponse200329) HasCo2() bool`

HasCo2 returns a boolean if a field has been set.

### SetCo2Nil

`func (o *InlineResponse200329) SetCo2Nil(b bool)`

 SetCo2Nil sets the value for Co2 to be an explicit nil

### UnsetCo2
`func (o *InlineResponse200329) UnsetCo2()`

UnsetCo2 ensures that no value is present for Co2, not even an explicit nil
### GetCurrent

`func (o *InlineResponse200329) GetCurrent() OrganizationsOrganizationIdSensorReadingsHistoryCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *InlineResponse200329) GetCurrentOk() (*OrganizationsOrganizationIdSensorReadingsHistoryCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *InlineResponse200329) SetCurrent(v OrganizationsOrganizationIdSensorReadingsHistoryCurrent)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *InlineResponse200329) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *InlineResponse200329) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *InlineResponse200329) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetDoor

`func (o *InlineResponse200329) GetDoor() OrganizationsOrganizationIdSensorReadingsHistoryDoor`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *InlineResponse200329) GetDoorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDoor, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *InlineResponse200329) SetDoor(v OrganizationsOrganizationIdSensorReadingsHistoryDoor)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *InlineResponse200329) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### SetDoorNil

`func (o *InlineResponse200329) SetDoorNil(b bool)`

 SetDoorNil sets the value for Door to be an explicit nil

### UnsetDoor
`func (o *InlineResponse200329) UnsetDoor()`

UnsetDoor ensures that no value is present for Door, not even an explicit nil
### GetDownstreamPower

`func (o *InlineResponse200329) GetDownstreamPower() OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower`

GetDownstreamPower returns the DownstreamPower field if non-nil, zero value otherwise.

### GetDownstreamPowerOk

`func (o *InlineResponse200329) GetDownstreamPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower, bool)`

GetDownstreamPowerOk returns a tuple with the DownstreamPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamPower

`func (o *InlineResponse200329) SetDownstreamPower(v OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower)`

SetDownstreamPower sets DownstreamPower field to given value.

### HasDownstreamPower

`func (o *InlineResponse200329) HasDownstreamPower() bool`

HasDownstreamPower returns a boolean if a field has been set.

### SetDownstreamPowerNil

`func (o *InlineResponse200329) SetDownstreamPowerNil(b bool)`

 SetDownstreamPowerNil sets the value for DownstreamPower to be an explicit nil

### UnsetDownstreamPower
`func (o *InlineResponse200329) UnsetDownstreamPower()`

UnsetDownstreamPower ensures that no value is present for DownstreamPower, not even an explicit nil
### GetFrequency

`func (o *InlineResponse200329) GetFrequency() OrganizationsOrganizationIdSensorReadingsHistoryFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *InlineResponse200329) GetFrequencyOk() (*OrganizationsOrganizationIdSensorReadingsHistoryFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *InlineResponse200329) SetFrequency(v OrganizationsOrganizationIdSensorReadingsHistoryFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *InlineResponse200329) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *InlineResponse200329) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *InlineResponse200329) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetHumidity

`func (o *InlineResponse200329) GetHumidity() OrganizationsOrganizationIdSensorReadingsHistoryHumidity`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *InlineResponse200329) GetHumidityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryHumidity, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *InlineResponse200329) SetHumidity(v OrganizationsOrganizationIdSensorReadingsHistoryHumidity)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *InlineResponse200329) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### SetHumidityNil

`func (o *InlineResponse200329) SetHumidityNil(b bool)`

 SetHumidityNil sets the value for Humidity to be an explicit nil

### UnsetHumidity
`func (o *InlineResponse200329) UnsetHumidity()`

UnsetHumidity ensures that no value is present for Humidity, not even an explicit nil
### GetIndoorAirQuality

`func (o *InlineResponse200329) GetIndoorAirQuality() OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *InlineResponse200329) GetIndoorAirQualityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *InlineResponse200329) SetIndoorAirQuality(v OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *InlineResponse200329) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### SetIndoorAirQualityNil

`func (o *InlineResponse200329) SetIndoorAirQualityNil(b bool)`

 SetIndoorAirQualityNil sets the value for IndoorAirQuality to be an explicit nil

### UnsetIndoorAirQuality
`func (o *InlineResponse200329) UnsetIndoorAirQuality()`

UnsetIndoorAirQuality ensures that no value is present for IndoorAirQuality, not even an explicit nil
### GetNoise

`func (o *InlineResponse200329) GetNoise() OrganizationsOrganizationIdSensorReadingsHistoryNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *InlineResponse200329) GetNoiseOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *InlineResponse200329) SetNoise(v OrganizationsOrganizationIdSensorReadingsHistoryNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *InlineResponse200329) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### SetNoiseNil

`func (o *InlineResponse200329) SetNoiseNil(b bool)`

 SetNoiseNil sets the value for Noise to be an explicit nil

### UnsetNoise
`func (o *InlineResponse200329) UnsetNoise()`

UnsetNoise ensures that no value is present for Noise, not even an explicit nil
### GetNo2

`func (o *InlineResponse200329) GetNo2() OrganizationsOrganizationIdSensorReadingsHistoryNo2`

GetNo2 returns the No2 field if non-nil, zero value otherwise.

### GetNo2Ok

`func (o *InlineResponse200329) GetNo2Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryNo2, bool)`

GetNo2Ok returns a tuple with the No2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo2

`func (o *InlineResponse200329) SetNo2(v OrganizationsOrganizationIdSensorReadingsHistoryNo2)`

SetNo2 sets No2 field to given value.

### HasNo2

`func (o *InlineResponse200329) HasNo2() bool`

HasNo2 returns a boolean if a field has been set.

### SetNo2Nil

`func (o *InlineResponse200329) SetNo2Nil(b bool)`

 SetNo2Nil sets the value for No2 to be an explicit nil

### UnsetNo2
`func (o *InlineResponse200329) UnsetNo2()`

UnsetNo2 ensures that no value is present for No2, not even an explicit nil
### GetO3

`func (o *InlineResponse200329) GetO3() OrganizationsOrganizationIdSensorReadingsHistoryO3`

GetO3 returns the O3 field if non-nil, zero value otherwise.

### GetO3Ok

`func (o *InlineResponse200329) GetO3Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryO3, bool)`

GetO3Ok returns a tuple with the O3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO3

`func (o *InlineResponse200329) SetO3(v OrganizationsOrganizationIdSensorReadingsHistoryO3)`

SetO3 sets O3 field to given value.

### HasO3

`func (o *InlineResponse200329) HasO3() bool`

HasO3 returns a boolean if a field has been set.

### SetO3Nil

`func (o *InlineResponse200329) SetO3Nil(b bool)`

 SetO3Nil sets the value for O3 to be an explicit nil

### UnsetO3
`func (o *InlineResponse200329) UnsetO3()`

UnsetO3 ensures that no value is present for O3, not even an explicit nil
### GetPm10

`func (o *InlineResponse200329) GetPm10() OrganizationsOrganizationIdSensorReadingsHistoryPm10`

GetPm10 returns the Pm10 field if non-nil, zero value otherwise.

### GetPm10Ok

`func (o *InlineResponse200329) GetPm10Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm10, bool)`

GetPm10Ok returns a tuple with the Pm10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm10

`func (o *InlineResponse200329) SetPm10(v OrganizationsOrganizationIdSensorReadingsHistoryPm10)`

SetPm10 sets Pm10 field to given value.

### HasPm10

`func (o *InlineResponse200329) HasPm10() bool`

HasPm10 returns a boolean if a field has been set.

### SetPm10Nil

`func (o *InlineResponse200329) SetPm10Nil(b bool)`

 SetPm10Nil sets the value for Pm10 to be an explicit nil

### UnsetPm10
`func (o *InlineResponse200329) UnsetPm10()`

UnsetPm10 ensures that no value is present for Pm10, not even an explicit nil
### GetPm25

`func (o *InlineResponse200329) GetPm25() OrganizationsOrganizationIdSensorReadingsHistoryPm25`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *InlineResponse200329) GetPm25Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm25, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *InlineResponse200329) SetPm25(v OrganizationsOrganizationIdSensorReadingsHistoryPm25)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *InlineResponse200329) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### SetPm25Nil

`func (o *InlineResponse200329) SetPm25Nil(b bool)`

 SetPm25Nil sets the value for Pm25 to be an explicit nil

### UnsetPm25
`func (o *InlineResponse200329) UnsetPm25()`

UnsetPm25 ensures that no value is present for Pm25, not even an explicit nil
### GetPowerFactor

`func (o *InlineResponse200329) GetPowerFactor() OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *InlineResponse200329) GetPowerFactorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *InlineResponse200329) SetPowerFactor(v OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *InlineResponse200329) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### SetPowerFactorNil

`func (o *InlineResponse200329) SetPowerFactorNil(b bool)`

 SetPowerFactorNil sets the value for PowerFactor to be an explicit nil

### UnsetPowerFactor
`func (o *InlineResponse200329) UnsetPowerFactor()`

UnsetPowerFactor ensures that no value is present for PowerFactor, not even an explicit nil
### GetRealPower

`func (o *InlineResponse200329) GetRealPower() OrganizationsOrganizationIdSensorReadingsHistoryRealPower`

GetRealPower returns the RealPower field if non-nil, zero value otherwise.

### GetRealPowerOk

`func (o *InlineResponse200329) GetRealPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRealPower, bool)`

GetRealPowerOk returns a tuple with the RealPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealPower

`func (o *InlineResponse200329) SetRealPower(v OrganizationsOrganizationIdSensorReadingsHistoryRealPower)`

SetRealPower sets RealPower field to given value.

### HasRealPower

`func (o *InlineResponse200329) HasRealPower() bool`

HasRealPower returns a boolean if a field has been set.

### SetRealPowerNil

`func (o *InlineResponse200329) SetRealPowerNil(b bool)`

 SetRealPowerNil sets the value for RealPower to be an explicit nil

### UnsetRealPower
`func (o *InlineResponse200329) UnsetRealPower()`

UnsetRealPower ensures that no value is present for RealPower, not even an explicit nil
### GetRemoteLockoutSwitch

`func (o *InlineResponse200329) GetRemoteLockoutSwitch() OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch`

GetRemoteLockoutSwitch returns the RemoteLockoutSwitch field if non-nil, zero value otherwise.

### GetRemoteLockoutSwitchOk

`func (o *InlineResponse200329) GetRemoteLockoutSwitchOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch, bool)`

GetRemoteLockoutSwitchOk returns a tuple with the RemoteLockoutSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLockoutSwitch

`func (o *InlineResponse200329) SetRemoteLockoutSwitch(v OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch)`

SetRemoteLockoutSwitch sets RemoteLockoutSwitch field to given value.

### HasRemoteLockoutSwitch

`func (o *InlineResponse200329) HasRemoteLockoutSwitch() bool`

HasRemoteLockoutSwitch returns a boolean if a field has been set.

### SetRemoteLockoutSwitchNil

`func (o *InlineResponse200329) SetRemoteLockoutSwitchNil(b bool)`

 SetRemoteLockoutSwitchNil sets the value for RemoteLockoutSwitch to be an explicit nil

### UnsetRemoteLockoutSwitch
`func (o *InlineResponse200329) UnsetRemoteLockoutSwitch()`

UnsetRemoteLockoutSwitch ensures that no value is present for RemoteLockoutSwitch, not even an explicit nil
### GetTemperature

`func (o *InlineResponse200329) GetTemperature() OrganizationsOrganizationIdSensorReadingsHistoryTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *InlineResponse200329) GetTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *InlineResponse200329) SetTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *InlineResponse200329) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *InlineResponse200329) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *InlineResponse200329) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetTvoc

`func (o *InlineResponse200329) GetTvoc() OrganizationsOrganizationIdSensorReadingsHistoryTvoc`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *InlineResponse200329) GetTvocOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTvoc, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *InlineResponse200329) SetTvoc(v OrganizationsOrganizationIdSensorReadingsHistoryTvoc)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *InlineResponse200329) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### SetTvocNil

`func (o *InlineResponse200329) SetTvocNil(b bool)`

 SetTvocNil sets the value for Tvoc to be an explicit nil

### UnsetTvoc
`func (o *InlineResponse200329) UnsetTvoc()`

UnsetTvoc ensures that no value is present for Tvoc, not even an explicit nil
### GetVoltage

`func (o *InlineResponse200329) GetVoltage() OrganizationsOrganizationIdSensorReadingsHistoryVoltage`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *InlineResponse200329) GetVoltageOk() (*OrganizationsOrganizationIdSensorReadingsHistoryVoltage, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *InlineResponse200329) SetVoltage(v OrganizationsOrganizationIdSensorReadingsHistoryVoltage)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *InlineResponse200329) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### SetVoltageNil

`func (o *InlineResponse200329) SetVoltageNil(b bool)`

 SetVoltageNil sets the value for Voltage to be an explicit nil

### UnsetVoltage
`func (o *InlineResponse200329) UnsetVoltage()`

UnsetVoltage ensures that no value is present for Voltage, not even an explicit nil
### GetWater

`func (o *InlineResponse200329) GetWater() OrganizationsOrganizationIdSensorReadingsHistoryWater`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *InlineResponse200329) GetWaterOk() (*OrganizationsOrganizationIdSensorReadingsHistoryWater, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *InlineResponse200329) SetWater(v OrganizationsOrganizationIdSensorReadingsHistoryWater)`

SetWater sets Water field to given value.

### HasWater

`func (o *InlineResponse200329) HasWater() bool`

HasWater returns a boolean if a field has been set.

### SetWaterNil

`func (o *InlineResponse200329) SetWaterNil(b bool)`

 SetWaterNil sets the value for Water to be an explicit nil

### UnsetWater
`func (o *InlineResponse200329) UnsetWater()`

UnsetWater ensures that no value is present for Water, not even an explicit nil
### GetRawTemperature

`func (o *InlineResponse200329) GetRawTemperature() OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature`

GetRawTemperature returns the RawTemperature field if non-nil, zero value otherwise.

### GetRawTemperatureOk

`func (o *InlineResponse200329) GetRawTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature, bool)`

GetRawTemperatureOk returns a tuple with the RawTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTemperature

`func (o *InlineResponse200329) SetRawTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature)`

SetRawTemperature sets RawTemperature field to given value.

### HasRawTemperature

`func (o *InlineResponse200329) HasRawTemperature() bool`

HasRawTemperature returns a boolean if a field has been set.

### SetRawTemperatureNil

`func (o *InlineResponse200329) SetRawTemperatureNil(b bool)`

 SetRawTemperatureNil sets the value for RawTemperature to be an explicit nil

### UnsetRawTemperature
`func (o *InlineResponse200329) UnsetRawTemperature()`

UnsetRawTemperature ensures that no value is present for RawTemperature, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


