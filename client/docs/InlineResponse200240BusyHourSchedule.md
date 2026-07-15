# InlineResponse200240BusyHourSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The Busy Hour mode applied to the network when minimizeChanges is enabled. Must be one of &#39;automatic&#39; or &#39;manual&#39;. Automatic busy hour is only available on firmware versions &gt;&#x3D; MR 27.0 | [optional] 
**Automatic** | Pointer to [**InlineResponse200240BusyHourScheduleAutomatic**](InlineResponse200240BusyHourScheduleAutomatic.md) |  | [optional] 
**Manual** | Pointer to [**NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual**](NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual.md) |  | [optional] 

## Methods

### NewInlineResponse200240BusyHourSchedule

`func NewInlineResponse200240BusyHourSchedule() *InlineResponse200240BusyHourSchedule`

NewInlineResponse200240BusyHourSchedule instantiates a new InlineResponse200240BusyHourSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200240BusyHourScheduleWithDefaults

`func NewInlineResponse200240BusyHourScheduleWithDefaults() *InlineResponse200240BusyHourSchedule`

NewInlineResponse200240BusyHourScheduleWithDefaults instantiates a new InlineResponse200240BusyHourSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200240BusyHourSchedule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200240BusyHourSchedule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200240BusyHourSchedule) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200240BusyHourSchedule) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAutomatic

`func (o *InlineResponse200240BusyHourSchedule) GetAutomatic() InlineResponse200240BusyHourScheduleAutomatic`

GetAutomatic returns the Automatic field if non-nil, zero value otherwise.

### GetAutomaticOk

`func (o *InlineResponse200240BusyHourSchedule) GetAutomaticOk() (*InlineResponse200240BusyHourScheduleAutomatic, bool)`

GetAutomaticOk returns a tuple with the Automatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic

`func (o *InlineResponse200240BusyHourSchedule) SetAutomatic(v InlineResponse200240BusyHourScheduleAutomatic)`

SetAutomatic sets Automatic field to given value.

### HasAutomatic

`func (o *InlineResponse200240BusyHourSchedule) HasAutomatic() bool`

HasAutomatic returns a boolean if a field has been set.

### GetManual

`func (o *InlineResponse200240BusyHourSchedule) GetManual() NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *InlineResponse200240BusyHourSchedule) GetManualOk() (*NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *InlineResponse200240BusyHourSchedule) SetManual(v NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual)`

SetManual sets Manual field to given value.

### HasManual

`func (o *InlineResponse200240BusyHourSchedule) HasManual() bool`

HasManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


