# OrganizationsOrganizationIdSensorReadingsLatestReadings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewOrganizationsOrganizationIdSensorReadingsLatestReadings

`func NewOrganizationsOrganizationIdSensorReadingsLatestReadings() *OrganizationsOrganizationIdSensorReadingsLatestReadings`

NewOrganizationsOrganizationIdSensorReadingsLatestReadings instantiates a new OrganizationsOrganizationIdSensorReadingsLatestReadings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdSensorReadingsLatestReadingsWithDefaults

`func NewOrganizationsOrganizationIdSensorReadingsLatestReadingsWithDefaults() *OrganizationsOrganizationIdSensorReadingsLatestReadings`

NewOrganizationsOrganizationIdSensorReadingsLatestReadingsWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsLatestReadings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetMetric

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetApparentPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetApparentPower() OrganizationsOrganizationIdSensorReadingsHistoryApparentPower`

GetApparentPower returns the ApparentPower field if non-nil, zero value otherwise.

### GetApparentPowerOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetApparentPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryApparentPower, bool)`

GetApparentPowerOk returns a tuple with the ApparentPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApparentPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetApparentPower(v OrganizationsOrganizationIdSensorReadingsHistoryApparentPower)`

SetApparentPower sets ApparentPower field to given value.

### HasApparentPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasApparentPower() bool`

HasApparentPower returns a boolean if a field has been set.

### SetApparentPowerNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetApparentPowerNil(b bool)`

 SetApparentPowerNil sets the value for ApparentPower to be an explicit nil

### UnsetApparentPower
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetApparentPower()`

UnsetApparentPower ensures that no value is present for ApparentPower, not even an explicit nil
### GetBattery

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetBattery() OrganizationsOrganizationIdSensorReadingsHistoryBattery`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetBatteryOk() (*OrganizationsOrganizationIdSensorReadingsHistoryBattery, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetBattery(v OrganizationsOrganizationIdSensorReadingsHistoryBattery)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasBattery() bool`

HasBattery returns a boolean if a field has been set.

### SetBatteryNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetBatteryNil(b bool)`

 SetBatteryNil sets the value for Battery to be an explicit nil

### UnsetBattery
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetBattery()`

UnsetBattery ensures that no value is present for Battery, not even an explicit nil
### GetButton

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetButton() OrganizationsOrganizationIdSensorReadingsHistoryButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetButtonOk() (*OrganizationsOrganizationIdSensorReadingsHistoryButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetButton(v OrganizationsOrganizationIdSensorReadingsHistoryButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasButton() bool`

HasButton returns a boolean if a field has been set.

### SetButtonNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetButtonNil(b bool)`

 SetButtonNil sets the value for Button to be an explicit nil

### UnsetButton
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetButton()`

UnsetButton ensures that no value is present for Button, not even an explicit nil
### GetCo2

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetCo2() OrganizationsOrganizationIdSensorReadingsHistoryCo2`

GetCo2 returns the Co2 field if non-nil, zero value otherwise.

### GetCo2Ok

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetCo2Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryCo2, bool)`

GetCo2Ok returns a tuple with the Co2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetCo2(v OrganizationsOrganizationIdSensorReadingsHistoryCo2)`

SetCo2 sets Co2 field to given value.

### HasCo2

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasCo2() bool`

HasCo2 returns a boolean if a field has been set.

### SetCo2Nil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetCo2Nil(b bool)`

 SetCo2Nil sets the value for Co2 to be an explicit nil

### UnsetCo2
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetCo2()`

UnsetCo2 ensures that no value is present for Co2, not even an explicit nil
### GetCurrent

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetCurrent() OrganizationsOrganizationIdSensorReadingsHistoryCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetCurrentOk() (*OrganizationsOrganizationIdSensorReadingsHistoryCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetCurrent(v OrganizationsOrganizationIdSensorReadingsHistoryCurrent)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetDoor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetDoor() OrganizationsOrganizationIdSensorReadingsHistoryDoor`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetDoorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDoor, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetDoor(v OrganizationsOrganizationIdSensorReadingsHistoryDoor)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### SetDoorNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetDoorNil(b bool)`

 SetDoorNil sets the value for Door to be an explicit nil

### UnsetDoor
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetDoor()`

UnsetDoor ensures that no value is present for Door, not even an explicit nil
### GetDownstreamPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetDownstreamPower() OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower`

GetDownstreamPower returns the DownstreamPower field if non-nil, zero value otherwise.

### GetDownstreamPowerOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetDownstreamPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower, bool)`

GetDownstreamPowerOk returns a tuple with the DownstreamPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetDownstreamPower(v OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower)`

SetDownstreamPower sets DownstreamPower field to given value.

### HasDownstreamPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasDownstreamPower() bool`

HasDownstreamPower returns a boolean if a field has been set.

### SetDownstreamPowerNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetDownstreamPowerNil(b bool)`

 SetDownstreamPowerNil sets the value for DownstreamPower to be an explicit nil

### UnsetDownstreamPower
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetDownstreamPower()`

UnsetDownstreamPower ensures that no value is present for DownstreamPower, not even an explicit nil
### GetFrequency

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetFrequency() OrganizationsOrganizationIdSensorReadingsHistoryFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetFrequencyOk() (*OrganizationsOrganizationIdSensorReadingsHistoryFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetFrequency(v OrganizationsOrganizationIdSensorReadingsHistoryFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetHumidity

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetHumidity() OrganizationsOrganizationIdSensorReadingsHistoryHumidity`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetHumidityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryHumidity, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetHumidity(v OrganizationsOrganizationIdSensorReadingsHistoryHumidity)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### SetHumidityNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetHumidityNil(b bool)`

 SetHumidityNil sets the value for Humidity to be an explicit nil

### UnsetHumidity
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetHumidity()`

UnsetHumidity ensures that no value is present for Humidity, not even an explicit nil
### GetIndoorAirQuality

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetIndoorAirQuality() OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetIndoorAirQualityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetIndoorAirQuality(v OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### SetIndoorAirQualityNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetIndoorAirQualityNil(b bool)`

 SetIndoorAirQualityNil sets the value for IndoorAirQuality to be an explicit nil

### UnsetIndoorAirQuality
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetIndoorAirQuality()`

UnsetIndoorAirQuality ensures that no value is present for IndoorAirQuality, not even an explicit nil
### GetNoise

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetNoise() OrganizationsOrganizationIdSensorReadingsHistoryNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetNoiseOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetNoise(v OrganizationsOrganizationIdSensorReadingsHistoryNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### SetNoiseNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetNoiseNil(b bool)`

 SetNoiseNil sets the value for Noise to be an explicit nil

### UnsetNoise
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetNoise()`

UnsetNoise ensures that no value is present for Noise, not even an explicit nil
### GetNo2

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetNo2() OrganizationsOrganizationIdSensorReadingsHistoryNo2`

GetNo2 returns the No2 field if non-nil, zero value otherwise.

### GetNo2Ok

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetNo2Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryNo2, bool)`

GetNo2Ok returns a tuple with the No2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo2

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetNo2(v OrganizationsOrganizationIdSensorReadingsHistoryNo2)`

SetNo2 sets No2 field to given value.

### HasNo2

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasNo2() bool`

HasNo2 returns a boolean if a field has been set.

### SetNo2Nil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetNo2Nil(b bool)`

 SetNo2Nil sets the value for No2 to be an explicit nil

### UnsetNo2
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetNo2()`

UnsetNo2 ensures that no value is present for No2, not even an explicit nil
### GetO3

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetO3() OrganizationsOrganizationIdSensorReadingsHistoryO3`

GetO3 returns the O3 field if non-nil, zero value otherwise.

### GetO3Ok

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetO3Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryO3, bool)`

GetO3Ok returns a tuple with the O3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO3

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetO3(v OrganizationsOrganizationIdSensorReadingsHistoryO3)`

SetO3 sets O3 field to given value.

### HasO3

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasO3() bool`

HasO3 returns a boolean if a field has been set.

### SetO3Nil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetO3Nil(b bool)`

 SetO3Nil sets the value for O3 to be an explicit nil

### UnsetO3
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetO3()`

UnsetO3 ensures that no value is present for O3, not even an explicit nil
### GetPm10

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetPm10() OrganizationsOrganizationIdSensorReadingsHistoryPm10`

GetPm10 returns the Pm10 field if non-nil, zero value otherwise.

### GetPm10Ok

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetPm10Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm10, bool)`

GetPm10Ok returns a tuple with the Pm10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm10

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetPm10(v OrganizationsOrganizationIdSensorReadingsHistoryPm10)`

SetPm10 sets Pm10 field to given value.

### HasPm10

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasPm10() bool`

HasPm10 returns a boolean if a field has been set.

### SetPm10Nil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetPm10Nil(b bool)`

 SetPm10Nil sets the value for Pm10 to be an explicit nil

### UnsetPm10
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetPm10()`

UnsetPm10 ensures that no value is present for Pm10, not even an explicit nil
### GetPm25

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetPm25() OrganizationsOrganizationIdSensorReadingsHistoryPm25`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetPm25Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm25, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetPm25(v OrganizationsOrganizationIdSensorReadingsHistoryPm25)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### SetPm25Nil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetPm25Nil(b bool)`

 SetPm25Nil sets the value for Pm25 to be an explicit nil

### UnsetPm25
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetPm25()`

UnsetPm25 ensures that no value is present for Pm25, not even an explicit nil
### GetPowerFactor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetPowerFactor() OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetPowerFactorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetPowerFactor(v OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### SetPowerFactorNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetPowerFactorNil(b bool)`

 SetPowerFactorNil sets the value for PowerFactor to be an explicit nil

### UnsetPowerFactor
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetPowerFactor()`

UnsetPowerFactor ensures that no value is present for PowerFactor, not even an explicit nil
### GetRealPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetRealPower() OrganizationsOrganizationIdSensorReadingsHistoryRealPower`

GetRealPower returns the RealPower field if non-nil, zero value otherwise.

### GetRealPowerOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetRealPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRealPower, bool)`

GetRealPowerOk returns a tuple with the RealPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetRealPower(v OrganizationsOrganizationIdSensorReadingsHistoryRealPower)`

SetRealPower sets RealPower field to given value.

### HasRealPower

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasRealPower() bool`

HasRealPower returns a boolean if a field has been set.

### SetRealPowerNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetRealPowerNil(b bool)`

 SetRealPowerNil sets the value for RealPower to be an explicit nil

### UnsetRealPower
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetRealPower()`

UnsetRealPower ensures that no value is present for RealPower, not even an explicit nil
### GetRemoteLockoutSwitch

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetRemoteLockoutSwitch() OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch`

GetRemoteLockoutSwitch returns the RemoteLockoutSwitch field if non-nil, zero value otherwise.

### GetRemoteLockoutSwitchOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetRemoteLockoutSwitchOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch, bool)`

GetRemoteLockoutSwitchOk returns a tuple with the RemoteLockoutSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLockoutSwitch

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetRemoteLockoutSwitch(v OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch)`

SetRemoteLockoutSwitch sets RemoteLockoutSwitch field to given value.

### HasRemoteLockoutSwitch

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasRemoteLockoutSwitch() bool`

HasRemoteLockoutSwitch returns a boolean if a field has been set.

### SetRemoteLockoutSwitchNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetRemoteLockoutSwitchNil(b bool)`

 SetRemoteLockoutSwitchNil sets the value for RemoteLockoutSwitch to be an explicit nil

### UnsetRemoteLockoutSwitch
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetRemoteLockoutSwitch()`

UnsetRemoteLockoutSwitch ensures that no value is present for RemoteLockoutSwitch, not even an explicit nil
### GetTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTemperature() OrganizationsOrganizationIdSensorReadingsHistoryTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetTvoc

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTvoc() OrganizationsOrganizationIdSensorReadingsHistoryTvoc`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTvocOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTvoc, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetTvoc(v OrganizationsOrganizationIdSensorReadingsHistoryTvoc)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### SetTvocNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetTvocNil(b bool)`

 SetTvocNil sets the value for Tvoc to be an explicit nil

### UnsetTvoc
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetTvoc()`

UnsetTvoc ensures that no value is present for Tvoc, not even an explicit nil
### GetVoltage

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetVoltage() OrganizationsOrganizationIdSensorReadingsHistoryVoltage`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetVoltageOk() (*OrganizationsOrganizationIdSensorReadingsHistoryVoltage, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetVoltage(v OrganizationsOrganizationIdSensorReadingsHistoryVoltage)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### SetVoltageNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetVoltageNil(b bool)`

 SetVoltageNil sets the value for Voltage to be an explicit nil

### UnsetVoltage
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetVoltage()`

UnsetVoltage ensures that no value is present for Voltage, not even an explicit nil
### GetWater

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetWater() OrganizationsOrganizationIdSensorReadingsHistoryWater`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetWaterOk() (*OrganizationsOrganizationIdSensorReadingsHistoryWater, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetWater(v OrganizationsOrganizationIdSensorReadingsHistoryWater)`

SetWater sets Water field to given value.

### HasWater

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasWater() bool`

HasWater returns a boolean if a field has been set.

### SetWaterNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetWaterNil(b bool)`

 SetWaterNil sets the value for Water to be an explicit nil

### UnsetWater
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetWater()`

UnsetWater ensures that no value is present for Water, not even an explicit nil
### GetRawTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetRawTemperature() OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature`

GetRawTemperature returns the RawTemperature field if non-nil, zero value otherwise.

### GetRawTemperatureOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetRawTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature, bool)`

GetRawTemperatureOk returns a tuple with the RawTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetRawTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature)`

SetRawTemperature sets RawTemperature field to given value.

### HasRawTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasRawTemperature() bool`

HasRawTemperature returns a boolean if a field has been set.

### SetRawTemperatureNil

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetRawTemperatureNil(b bool)`

 SetRawTemperatureNil sets the value for RawTemperature to be an explicit nil

### UnsetRawTemperature
`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) UnsetRawTemperature()`

UnsetRawTemperature ensures that no value is present for RawTemperature, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


