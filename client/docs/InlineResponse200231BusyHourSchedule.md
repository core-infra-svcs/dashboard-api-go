# InlineResponse200231BusyHourSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The Busy Hour mode applied to the network when minimizeChanges is enabled. Must be one of &#39;automatic&#39; or &#39;manual&#39;. Automatic busy hour is only available on firmware versions &gt;&#x3D; MR 27.0 | [optional] 
**Automatic** | Pointer to [**InlineResponse200231BusyHourScheduleAutomatic**](InlineResponse200231BusyHourScheduleAutomatic.md) |  | [optional] 
**Manual** | Pointer to [**NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual**](NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual.md) |  | [optional] 

## Methods

### NewInlineResponse200231BusyHourSchedule

`func NewInlineResponse200231BusyHourSchedule() *InlineResponse200231BusyHourSchedule`

NewInlineResponse200231BusyHourSchedule instantiates a new InlineResponse200231BusyHourSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200231BusyHourScheduleWithDefaults

`func NewInlineResponse200231BusyHourScheduleWithDefaults() *InlineResponse200231BusyHourSchedule`

NewInlineResponse200231BusyHourScheduleWithDefaults instantiates a new InlineResponse200231BusyHourSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200231BusyHourSchedule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200231BusyHourSchedule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200231BusyHourSchedule) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200231BusyHourSchedule) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAutomatic

`func (o *InlineResponse200231BusyHourSchedule) GetAutomatic() InlineResponse200231BusyHourScheduleAutomatic`

GetAutomatic returns the Automatic field if non-nil, zero value otherwise.

### GetAutomaticOk

`func (o *InlineResponse200231BusyHourSchedule) GetAutomaticOk() (*InlineResponse200231BusyHourScheduleAutomatic, bool)`

GetAutomaticOk returns a tuple with the Automatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic

`func (o *InlineResponse200231BusyHourSchedule) SetAutomatic(v InlineResponse200231BusyHourScheduleAutomatic)`

SetAutomatic sets Automatic field to given value.

### HasAutomatic

`func (o *InlineResponse200231BusyHourSchedule) HasAutomatic() bool`

HasAutomatic returns a boolean if a field has been set.

### GetManual

`func (o *InlineResponse200231BusyHourSchedule) GetManual() NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *InlineResponse200231BusyHourSchedule) GetManualOk() (*NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *InlineResponse200231BusyHourSchedule) SetManual(v NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual)`

SetManual sets Manual field to given value.

### HasManual

`func (o *InlineResponse200231BusyHourSchedule) HasManual() bool`

HasManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


