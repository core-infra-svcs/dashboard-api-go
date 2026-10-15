# InlineResponse200460Contents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Node group ID | [optional] 
**DeviceType** | Pointer to **string** | Device type | [optional] 
**DeviceCategory** | Pointer to **string** | Device category | [optional] 
**Serial** | Pointer to **string** | Device ID | [optional] 
**Location** | Pointer to [**InlineResponse200460Location**](InlineResponse200460Location.md) |  | [optional] 
**StartTime** | Pointer to **string** | Start time (ISO8601) | [optional] 
**EndTime** | Pointer to **string** | End time (ISO8601) | [optional] 
**PowerConsumption** | Pointer to **float32** | Total power consumption in kWh | [optional] 
**CarbonEmissions** | Pointer to **float32** | Total carbon emissions in KgCO2 | [optional] 
**PowerCost** | Pointer to **float32** | Total power cost | [optional] 
**PercentageOfTotal** | Pointer to **float32** | Power consumption as a percentage of total | [optional] 
**IsEstimated** | Pointer to **int32** | Estimation flag (0 or 1) | [optional] 

## Methods

### NewInlineResponse200460Contents

`func NewInlineResponse200460Contents() *InlineResponse200460Contents`

NewInlineResponse200460Contents instantiates a new InlineResponse200460Contents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200460ContentsWithDefaults

`func NewInlineResponse200460ContentsWithDefaults() *InlineResponse200460Contents`

NewInlineResponse200460ContentsWithDefaults instantiates a new InlineResponse200460Contents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200460Contents) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200460Contents) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200460Contents) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200460Contents) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetDeviceType

`func (o *InlineResponse200460Contents) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InlineResponse200460Contents) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InlineResponse200460Contents) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *InlineResponse200460Contents) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceCategory

`func (o *InlineResponse200460Contents) GetDeviceCategory() string`

GetDeviceCategory returns the DeviceCategory field if non-nil, zero value otherwise.

### GetDeviceCategoryOk

`func (o *InlineResponse200460Contents) GetDeviceCategoryOk() (*string, bool)`

GetDeviceCategoryOk returns a tuple with the DeviceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCategory

`func (o *InlineResponse200460Contents) SetDeviceCategory(v string)`

SetDeviceCategory sets DeviceCategory field to given value.

### HasDeviceCategory

`func (o *InlineResponse200460Contents) HasDeviceCategory() bool`

HasDeviceCategory returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200460Contents) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200460Contents) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200460Contents) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200460Contents) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetLocation

`func (o *InlineResponse200460Contents) GetLocation() InlineResponse200460Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InlineResponse200460Contents) GetLocationOk() (*InlineResponse200460Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InlineResponse200460Contents) SetLocation(v InlineResponse200460Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InlineResponse200460Contents) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStartTime

`func (o *InlineResponse200460Contents) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InlineResponse200460Contents) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InlineResponse200460Contents) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InlineResponse200460Contents) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *InlineResponse200460Contents) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InlineResponse200460Contents) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InlineResponse200460Contents) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InlineResponse200460Contents) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPowerConsumption

`func (o *InlineResponse200460Contents) GetPowerConsumption() float32`

GetPowerConsumption returns the PowerConsumption field if non-nil, zero value otherwise.

### GetPowerConsumptionOk

`func (o *InlineResponse200460Contents) GetPowerConsumptionOk() (*float32, bool)`

GetPowerConsumptionOk returns a tuple with the PowerConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConsumption

`func (o *InlineResponse200460Contents) SetPowerConsumption(v float32)`

SetPowerConsumption sets PowerConsumption field to given value.

### HasPowerConsumption

`func (o *InlineResponse200460Contents) HasPowerConsumption() bool`

HasPowerConsumption returns a boolean if a field has been set.

### GetCarbonEmissions

`func (o *InlineResponse200460Contents) GetCarbonEmissions() float32`

GetCarbonEmissions returns the CarbonEmissions field if non-nil, zero value otherwise.

### GetCarbonEmissionsOk

`func (o *InlineResponse200460Contents) GetCarbonEmissionsOk() (*float32, bool)`

GetCarbonEmissionsOk returns a tuple with the CarbonEmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarbonEmissions

`func (o *InlineResponse200460Contents) SetCarbonEmissions(v float32)`

SetCarbonEmissions sets CarbonEmissions field to given value.

### HasCarbonEmissions

`func (o *InlineResponse200460Contents) HasCarbonEmissions() bool`

HasCarbonEmissions returns a boolean if a field has been set.

### GetPowerCost

`func (o *InlineResponse200460Contents) GetPowerCost() float32`

GetPowerCost returns the PowerCost field if non-nil, zero value otherwise.

### GetPowerCostOk

`func (o *InlineResponse200460Contents) GetPowerCostOk() (*float32, bool)`

GetPowerCostOk returns a tuple with the PowerCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCost

`func (o *InlineResponse200460Contents) SetPowerCost(v float32)`

SetPowerCost sets PowerCost field to given value.

### HasPowerCost

`func (o *InlineResponse200460Contents) HasPowerCost() bool`

HasPowerCost returns a boolean if a field has been set.

### GetPercentageOfTotal

`func (o *InlineResponse200460Contents) GetPercentageOfTotal() float32`

GetPercentageOfTotal returns the PercentageOfTotal field if non-nil, zero value otherwise.

### GetPercentageOfTotalOk

`func (o *InlineResponse200460Contents) GetPercentageOfTotalOk() (*float32, bool)`

GetPercentageOfTotalOk returns a tuple with the PercentageOfTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfTotal

`func (o *InlineResponse200460Contents) SetPercentageOfTotal(v float32)`

SetPercentageOfTotal sets PercentageOfTotal field to given value.

### HasPercentageOfTotal

`func (o *InlineResponse200460Contents) HasPercentageOfTotal() bool`

HasPercentageOfTotal returns a boolean if a field has been set.

### GetIsEstimated

`func (o *InlineResponse200460Contents) GetIsEstimated() int32`

GetIsEstimated returns the IsEstimated field if non-nil, zero value otherwise.

### GetIsEstimatedOk

`func (o *InlineResponse200460Contents) GetIsEstimatedOk() (*int32, bool)`

GetIsEstimatedOk returns a tuple with the IsEstimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEstimated

`func (o *InlineResponse200460Contents) SetIsEstimated(v int32)`

SetIsEstimated sets IsEstimated field to given value.

### HasIsEstimated

`func (o *InlineResponse200460Contents) HasIsEstimated() bool`

HasIsEstimated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


