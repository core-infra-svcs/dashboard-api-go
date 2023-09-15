# InlineResponse20045Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Door** | Pointer to **int32** | Number of sensors that are currently alerting due to an open door | [optional] 
**Humidity** | Pointer to **int32** | Number of sensors that are currently alerting due to humidity readings | [optional] 
**IndoorAirQuality** | Pointer to **int32** | Number of sensors that are currently alerting due to indoor air quality readings | [optional] 
**Noise** | Pointer to [**InlineResponse20045CountsNoise**](InlineResponse20045CountsNoise.md) |  | [optional] 
**Pm25** | Pointer to **int32** | Number of sensors that are currently alerting due to PM2.5 readings | [optional] 
**Temperature** | Pointer to **int32** | Number of sensors that are currently alerting due to temperature readings | [optional] 
**Tvoc** | Pointer to **int32** | Number of sensors that are currently alerting due to TVOC readings | [optional] 
**Water** | Pointer to **int32** | Number of sensors that are currently alerting due to the presence of water | [optional] 

## Methods

### NewInlineResponse20045Counts

`func NewInlineResponse20045Counts() *InlineResponse20045Counts`

NewInlineResponse20045Counts instantiates a new InlineResponse20045Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20045CountsWithDefaults

`func NewInlineResponse20045CountsWithDefaults() *InlineResponse20045Counts`

NewInlineResponse20045CountsWithDefaults instantiates a new InlineResponse20045Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoor

`func (o *InlineResponse20045Counts) GetDoor() int32`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *InlineResponse20045Counts) GetDoorOk() (*int32, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *InlineResponse20045Counts) SetDoor(v int32)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *InlineResponse20045Counts) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetHumidity

`func (o *InlineResponse20045Counts) GetHumidity() int32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *InlineResponse20045Counts) GetHumidityOk() (*int32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *InlineResponse20045Counts) SetHumidity(v int32)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *InlineResponse20045Counts) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *InlineResponse20045Counts) GetIndoorAirQuality() int32`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *InlineResponse20045Counts) GetIndoorAirQualityOk() (*int32, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *InlineResponse20045Counts) SetIndoorAirQuality(v int32)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *InlineResponse20045Counts) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *InlineResponse20045Counts) GetNoise() InlineResponse20045CountsNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *InlineResponse20045Counts) GetNoiseOk() (*InlineResponse20045CountsNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *InlineResponse20045Counts) SetNoise(v InlineResponse20045CountsNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *InlineResponse20045Counts) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *InlineResponse20045Counts) GetPm25() int32`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *InlineResponse20045Counts) GetPm25Ok() (*int32, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *InlineResponse20045Counts) SetPm25(v int32)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *InlineResponse20045Counts) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetTemperature

`func (o *InlineResponse20045Counts) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *InlineResponse20045Counts) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *InlineResponse20045Counts) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *InlineResponse20045Counts) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *InlineResponse20045Counts) GetTvoc() int32`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *InlineResponse20045Counts) GetTvocOk() (*int32, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *InlineResponse20045Counts) SetTvoc(v int32)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *InlineResponse20045Counts) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetWater

`func (o *InlineResponse20045Counts) GetWater() int32`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *InlineResponse20045Counts) GetWaterOk() (*int32, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *InlineResponse20045Counts) SetWater(v int32)`

SetWater sets Water field to given value.

### HasWater

`func (o *InlineResponse20045Counts) HasWater() bool`

HasWater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


