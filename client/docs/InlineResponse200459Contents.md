# InlineResponse200459Contents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceCategory** | Pointer to **string** | Device category | [optional] 
**StartTime** | Pointer to **string** | Start time (ISO8601) | [optional] 
**EndTime** | Pointer to **string** | End time (ISO8601) | [optional] 
**PowerConsumption** | Pointer to **float32** | Total power consumption in kWh | [optional] 
**CarbonEmissions** | Pointer to **float32** | Total carbon emissions in KgCO2 | [optional] 
**PowerCost** | Pointer to **float32** | Total power cost | [optional] 
**PercentageOfTotal** | Pointer to **float32** | Percentage of total power consumption | [optional] 
**IsEstimated** | Pointer to **int32** | Estimation flag (0 or 1) | [optional] 

## Methods

### NewInlineResponse200459Contents

`func NewInlineResponse200459Contents() *InlineResponse200459Contents`

NewInlineResponse200459Contents instantiates a new InlineResponse200459Contents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200459ContentsWithDefaults

`func NewInlineResponse200459ContentsWithDefaults() *InlineResponse200459Contents`

NewInlineResponse200459ContentsWithDefaults instantiates a new InlineResponse200459Contents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceCategory

`func (o *InlineResponse200459Contents) GetDeviceCategory() string`

GetDeviceCategory returns the DeviceCategory field if non-nil, zero value otherwise.

### GetDeviceCategoryOk

`func (o *InlineResponse200459Contents) GetDeviceCategoryOk() (*string, bool)`

GetDeviceCategoryOk returns a tuple with the DeviceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCategory

`func (o *InlineResponse200459Contents) SetDeviceCategory(v string)`

SetDeviceCategory sets DeviceCategory field to given value.

### HasDeviceCategory

`func (o *InlineResponse200459Contents) HasDeviceCategory() bool`

HasDeviceCategory returns a boolean if a field has been set.

### GetStartTime

`func (o *InlineResponse200459Contents) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InlineResponse200459Contents) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InlineResponse200459Contents) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InlineResponse200459Contents) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *InlineResponse200459Contents) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InlineResponse200459Contents) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InlineResponse200459Contents) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InlineResponse200459Contents) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPowerConsumption

`func (o *InlineResponse200459Contents) GetPowerConsumption() float32`

GetPowerConsumption returns the PowerConsumption field if non-nil, zero value otherwise.

### GetPowerConsumptionOk

`func (o *InlineResponse200459Contents) GetPowerConsumptionOk() (*float32, bool)`

GetPowerConsumptionOk returns a tuple with the PowerConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConsumption

`func (o *InlineResponse200459Contents) SetPowerConsumption(v float32)`

SetPowerConsumption sets PowerConsumption field to given value.

### HasPowerConsumption

`func (o *InlineResponse200459Contents) HasPowerConsumption() bool`

HasPowerConsumption returns a boolean if a field has been set.

### GetCarbonEmissions

`func (o *InlineResponse200459Contents) GetCarbonEmissions() float32`

GetCarbonEmissions returns the CarbonEmissions field if non-nil, zero value otherwise.

### GetCarbonEmissionsOk

`func (o *InlineResponse200459Contents) GetCarbonEmissionsOk() (*float32, bool)`

GetCarbonEmissionsOk returns a tuple with the CarbonEmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarbonEmissions

`func (o *InlineResponse200459Contents) SetCarbonEmissions(v float32)`

SetCarbonEmissions sets CarbonEmissions field to given value.

### HasCarbonEmissions

`func (o *InlineResponse200459Contents) HasCarbonEmissions() bool`

HasCarbonEmissions returns a boolean if a field has been set.

### GetPowerCost

`func (o *InlineResponse200459Contents) GetPowerCost() float32`

GetPowerCost returns the PowerCost field if non-nil, zero value otherwise.

### GetPowerCostOk

`func (o *InlineResponse200459Contents) GetPowerCostOk() (*float32, bool)`

GetPowerCostOk returns a tuple with the PowerCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCost

`func (o *InlineResponse200459Contents) SetPowerCost(v float32)`

SetPowerCost sets PowerCost field to given value.

### HasPowerCost

`func (o *InlineResponse200459Contents) HasPowerCost() bool`

HasPowerCost returns a boolean if a field has been set.

### GetPercentageOfTotal

`func (o *InlineResponse200459Contents) GetPercentageOfTotal() float32`

GetPercentageOfTotal returns the PercentageOfTotal field if non-nil, zero value otherwise.

### GetPercentageOfTotalOk

`func (o *InlineResponse200459Contents) GetPercentageOfTotalOk() (*float32, bool)`

GetPercentageOfTotalOk returns a tuple with the PercentageOfTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfTotal

`func (o *InlineResponse200459Contents) SetPercentageOfTotal(v float32)`

SetPercentageOfTotal sets PercentageOfTotal field to given value.

### HasPercentageOfTotal

`func (o *InlineResponse200459Contents) HasPercentageOfTotal() bool`

HasPercentageOfTotal returns a boolean if a field has been set.

### GetIsEstimated

`func (o *InlineResponse200459Contents) GetIsEstimated() int32`

GetIsEstimated returns the IsEstimated field if non-nil, zero value otherwise.

### GetIsEstimatedOk

`func (o *InlineResponse200459Contents) GetIsEstimatedOk() (*int32, bool)`

GetIsEstimatedOk returns a tuple with the IsEstimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEstimated

`func (o *InlineResponse200459Contents) SetIsEstimated(v int32)`

SetIsEstimated sets IsEstimated field to given value.

### HasIsEstimated

`func (o *InlineResponse200459Contents) HasIsEstimated() bool`

HasIsEstimated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


