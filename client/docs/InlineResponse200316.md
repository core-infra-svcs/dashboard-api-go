# InlineResponse200316

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number of the sensor that took the reading. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryNetwork**](OrganizationsOrganizationIdSensorReadingsHistoryNetwork.md) |  | [optional] 
**Ts** | Pointer to **string** | Time at which the reading occurred, in ISO8601 format. | [optional] 
**Metric** | Pointer to **string** | Type of sensor reading. | [optional] 
**ApparentPower** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryApparentPower**](OrganizationsOrganizationIdSensorReadingsHistoryApparentPower.md) |  | [optional] 
**Battery** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryBattery**](OrganizationsOrganizationIdSensorReadingsHistoryBattery.md) |  | [optional] 
**Button** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryButton**](OrganizationsOrganizationIdSensorReadingsHistoryButton.md) |  | [optional] 
**Co2** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryCo2**](OrganizationsOrganizationIdSensorReadingsHistoryCo2.md) |  | [optional] 
**Current** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryCurrent**](OrganizationsOrganizationIdSensorReadingsHistoryCurrent.md) |  | [optional] 
**Door** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryDoor**](OrganizationsOrganizationIdSensorReadingsHistoryDoor.md) |  | [optional] 
**DownstreamPower** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower**](OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower.md) |  | [optional] 
**Frequency** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryFrequency**](OrganizationsOrganizationIdSensorReadingsHistoryFrequency.md) |  | [optional] 
**Humidity** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryHumidity**](OrganizationsOrganizationIdSensorReadingsHistoryHumidity.md) |  | [optional] 
**IndoorAirQuality** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality**](OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality.md) |  | [optional] 
**Noise** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryNoise**](OrganizationsOrganizationIdSensorReadingsHistoryNoise.md) |  | [optional] 
**Pm25** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryPm25**](OrganizationsOrganizationIdSensorReadingsHistoryPm25.md) |  | [optional] 
**PowerFactor** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor**](OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor.md) |  | [optional] 
**RealPower** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryRealPower**](OrganizationsOrganizationIdSensorReadingsHistoryRealPower.md) |  | [optional] 
**RemoteLockoutSwitch** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch**](OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch.md) |  | [optional] 
**Temperature** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryTemperature**](OrganizationsOrganizationIdSensorReadingsHistoryTemperature.md) |  | [optional] 
**Tvoc** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryTvoc**](OrganizationsOrganizationIdSensorReadingsHistoryTvoc.md) |  | [optional] 
**Voltage** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryVoltage**](OrganizationsOrganizationIdSensorReadingsHistoryVoltage.md) |  | [optional] 
**Water** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryWater**](OrganizationsOrganizationIdSensorReadingsHistoryWater.md) |  | [optional] 
**RawTemperature** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature**](OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature.md) |  | [optional] 

## Methods

### NewInlineResponse200316

`func NewInlineResponse200316() *InlineResponse200316`

NewInlineResponse200316 instantiates a new InlineResponse200316 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200316WithDefaults

`func NewInlineResponse200316WithDefaults() *InlineResponse200316`

NewInlineResponse200316WithDefaults instantiates a new InlineResponse200316 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200316) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200316) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200316) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200316) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200316) GetNetwork() OrganizationsOrganizationIdSensorReadingsHistoryNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200316) GetNetworkOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200316) SetNetwork(v OrganizationsOrganizationIdSensorReadingsHistoryNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200316) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse200316) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200316) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200316) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200316) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetMetric

`func (o *InlineResponse200316) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *InlineResponse200316) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *InlineResponse200316) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *InlineResponse200316) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetApparentPower

`func (o *InlineResponse200316) GetApparentPower() OrganizationsOrganizationIdSensorReadingsHistoryApparentPower`

GetApparentPower returns the ApparentPower field if non-nil, zero value otherwise.

### GetApparentPowerOk

`func (o *InlineResponse200316) GetApparentPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryApparentPower, bool)`

GetApparentPowerOk returns a tuple with the ApparentPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApparentPower

`func (o *InlineResponse200316) SetApparentPower(v OrganizationsOrganizationIdSensorReadingsHistoryApparentPower)`

SetApparentPower sets ApparentPower field to given value.

### HasApparentPower

`func (o *InlineResponse200316) HasApparentPower() bool`

HasApparentPower returns a boolean if a field has been set.

### GetBattery

`func (o *InlineResponse200316) GetBattery() OrganizationsOrganizationIdSensorReadingsHistoryBattery`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *InlineResponse200316) GetBatteryOk() (*OrganizationsOrganizationIdSensorReadingsHistoryBattery, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *InlineResponse200316) SetBattery(v OrganizationsOrganizationIdSensorReadingsHistoryBattery)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *InlineResponse200316) HasBattery() bool`

HasBattery returns a boolean if a field has been set.

### GetButton

`func (o *InlineResponse200316) GetButton() OrganizationsOrganizationIdSensorReadingsHistoryButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *InlineResponse200316) GetButtonOk() (*OrganizationsOrganizationIdSensorReadingsHistoryButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *InlineResponse200316) SetButton(v OrganizationsOrganizationIdSensorReadingsHistoryButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *InlineResponse200316) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetCo2

`func (o *InlineResponse200316) GetCo2() OrganizationsOrganizationIdSensorReadingsHistoryCo2`

GetCo2 returns the Co2 field if non-nil, zero value otherwise.

### GetCo2Ok

`func (o *InlineResponse200316) GetCo2Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryCo2, bool)`

GetCo2Ok returns a tuple with the Co2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2

`func (o *InlineResponse200316) SetCo2(v OrganizationsOrganizationIdSensorReadingsHistoryCo2)`

SetCo2 sets Co2 field to given value.

### HasCo2

`func (o *InlineResponse200316) HasCo2() bool`

HasCo2 returns a boolean if a field has been set.

### GetCurrent

`func (o *InlineResponse200316) GetCurrent() OrganizationsOrganizationIdSensorReadingsHistoryCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *InlineResponse200316) GetCurrentOk() (*OrganizationsOrganizationIdSensorReadingsHistoryCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *InlineResponse200316) SetCurrent(v OrganizationsOrganizationIdSensorReadingsHistoryCurrent)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *InlineResponse200316) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetDoor

`func (o *InlineResponse200316) GetDoor() OrganizationsOrganizationIdSensorReadingsHistoryDoor`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *InlineResponse200316) GetDoorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDoor, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *InlineResponse200316) SetDoor(v OrganizationsOrganizationIdSensorReadingsHistoryDoor)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *InlineResponse200316) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetDownstreamPower

`func (o *InlineResponse200316) GetDownstreamPower() OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower`

GetDownstreamPower returns the DownstreamPower field if non-nil, zero value otherwise.

### GetDownstreamPowerOk

`func (o *InlineResponse200316) GetDownstreamPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower, bool)`

GetDownstreamPowerOk returns a tuple with the DownstreamPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamPower

`func (o *InlineResponse200316) SetDownstreamPower(v OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower)`

SetDownstreamPower sets DownstreamPower field to given value.

### HasDownstreamPower

`func (o *InlineResponse200316) HasDownstreamPower() bool`

HasDownstreamPower returns a boolean if a field has been set.

### GetFrequency

`func (o *InlineResponse200316) GetFrequency() OrganizationsOrganizationIdSensorReadingsHistoryFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *InlineResponse200316) GetFrequencyOk() (*OrganizationsOrganizationIdSensorReadingsHistoryFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *InlineResponse200316) SetFrequency(v OrganizationsOrganizationIdSensorReadingsHistoryFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *InlineResponse200316) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHumidity

`func (o *InlineResponse200316) GetHumidity() OrganizationsOrganizationIdSensorReadingsHistoryHumidity`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *InlineResponse200316) GetHumidityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryHumidity, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *InlineResponse200316) SetHumidity(v OrganizationsOrganizationIdSensorReadingsHistoryHumidity)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *InlineResponse200316) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *InlineResponse200316) GetIndoorAirQuality() OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *InlineResponse200316) GetIndoorAirQualityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *InlineResponse200316) SetIndoorAirQuality(v OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *InlineResponse200316) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *InlineResponse200316) GetNoise() OrganizationsOrganizationIdSensorReadingsHistoryNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *InlineResponse200316) GetNoiseOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *InlineResponse200316) SetNoise(v OrganizationsOrganizationIdSensorReadingsHistoryNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *InlineResponse200316) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *InlineResponse200316) GetPm25() OrganizationsOrganizationIdSensorReadingsHistoryPm25`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *InlineResponse200316) GetPm25Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm25, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *InlineResponse200316) SetPm25(v OrganizationsOrganizationIdSensorReadingsHistoryPm25)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *InlineResponse200316) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetPowerFactor

`func (o *InlineResponse200316) GetPowerFactor() OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *InlineResponse200316) GetPowerFactorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *InlineResponse200316) SetPowerFactor(v OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *InlineResponse200316) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### GetRealPower

`func (o *InlineResponse200316) GetRealPower() OrganizationsOrganizationIdSensorReadingsHistoryRealPower`

GetRealPower returns the RealPower field if non-nil, zero value otherwise.

### GetRealPowerOk

`func (o *InlineResponse200316) GetRealPowerOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRealPower, bool)`

GetRealPowerOk returns a tuple with the RealPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealPower

`func (o *InlineResponse200316) SetRealPower(v OrganizationsOrganizationIdSensorReadingsHistoryRealPower)`

SetRealPower sets RealPower field to given value.

### HasRealPower

`func (o *InlineResponse200316) HasRealPower() bool`

HasRealPower returns a boolean if a field has been set.

### GetRemoteLockoutSwitch

`func (o *InlineResponse200316) GetRemoteLockoutSwitch() OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch`

GetRemoteLockoutSwitch returns the RemoteLockoutSwitch field if non-nil, zero value otherwise.

### GetRemoteLockoutSwitchOk

`func (o *InlineResponse200316) GetRemoteLockoutSwitchOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch, bool)`

GetRemoteLockoutSwitchOk returns a tuple with the RemoteLockoutSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLockoutSwitch

`func (o *InlineResponse200316) SetRemoteLockoutSwitch(v OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch)`

SetRemoteLockoutSwitch sets RemoteLockoutSwitch field to given value.

### HasRemoteLockoutSwitch

`func (o *InlineResponse200316) HasRemoteLockoutSwitch() bool`

HasRemoteLockoutSwitch returns a boolean if a field has been set.

### GetTemperature

`func (o *InlineResponse200316) GetTemperature() OrganizationsOrganizationIdSensorReadingsHistoryTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *InlineResponse200316) GetTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *InlineResponse200316) SetTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *InlineResponse200316) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *InlineResponse200316) GetTvoc() OrganizationsOrganizationIdSensorReadingsHistoryTvoc`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *InlineResponse200316) GetTvocOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTvoc, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *InlineResponse200316) SetTvoc(v OrganizationsOrganizationIdSensorReadingsHistoryTvoc)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *InlineResponse200316) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetVoltage

`func (o *InlineResponse200316) GetVoltage() OrganizationsOrganizationIdSensorReadingsHistoryVoltage`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *InlineResponse200316) GetVoltageOk() (*OrganizationsOrganizationIdSensorReadingsHistoryVoltage, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *InlineResponse200316) SetVoltage(v OrganizationsOrganizationIdSensorReadingsHistoryVoltage)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *InlineResponse200316) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetWater

`func (o *InlineResponse200316) GetWater() OrganizationsOrganizationIdSensorReadingsHistoryWater`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *InlineResponse200316) GetWaterOk() (*OrganizationsOrganizationIdSensorReadingsHistoryWater, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *InlineResponse200316) SetWater(v OrganizationsOrganizationIdSensorReadingsHistoryWater)`

SetWater sets Water field to given value.

### HasWater

`func (o *InlineResponse200316) HasWater() bool`

HasWater returns a boolean if a field has been set.

### GetRawTemperature

`func (o *InlineResponse200316) GetRawTemperature() OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature`

GetRawTemperature returns the RawTemperature field if non-nil, zero value otherwise.

### GetRawTemperatureOk

`func (o *InlineResponse200316) GetRawTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature, bool)`

GetRawTemperatureOk returns a tuple with the RawTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTemperature

`func (o *InlineResponse200316) SetRawTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryRawTemperature)`

SetRawTemperature sets RawTemperature field to given value.

### HasRawTemperature

`func (o *InlineResponse200316) HasRawTemperature() bool`

HasRawTemperature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


