# InlineResponse200205BusyHourSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The Busy Hour mode applied to the network when minimizeChanges is enabled. Must be one of &#39;automatic&#39; or &#39;manual&#39;. Automatic busy hour is only available on firmware versions &gt;&#x3D; MR 27.0 | [optional] 
**Automatic** | Pointer to [**InlineResponse200205BusyHourScheduleAutomatic**](InlineResponse200205BusyHourScheduleAutomatic.md) |  | [optional] 
**Manual** | Pointer to [**NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual**](NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual.md) |  | [optional] 

## Methods

### NewInlineResponse200205BusyHourSchedule

`func NewInlineResponse200205BusyHourSchedule() *InlineResponse200205BusyHourSchedule`

NewInlineResponse200205BusyHourSchedule instantiates a new InlineResponse200205BusyHourSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200205BusyHourScheduleWithDefaults

`func NewInlineResponse200205BusyHourScheduleWithDefaults() *InlineResponse200205BusyHourSchedule`

NewInlineResponse200205BusyHourScheduleWithDefaults instantiates a new InlineResponse200205BusyHourSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200205BusyHourSchedule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200205BusyHourSchedule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200205BusyHourSchedule) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200205BusyHourSchedule) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAutomatic

`func (o *InlineResponse200205BusyHourSchedule) GetAutomatic() InlineResponse200205BusyHourScheduleAutomatic`

GetAutomatic returns the Automatic field if non-nil, zero value otherwise.

### GetAutomaticOk

`func (o *InlineResponse200205BusyHourSchedule) GetAutomaticOk() (*InlineResponse200205BusyHourScheduleAutomatic, bool)`

GetAutomaticOk returns a tuple with the Automatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic

`func (o *InlineResponse200205BusyHourSchedule) SetAutomatic(v InlineResponse200205BusyHourScheduleAutomatic)`

SetAutomatic sets Automatic field to given value.

### HasAutomatic

`func (o *InlineResponse200205BusyHourSchedule) HasAutomatic() bool`

HasAutomatic returns a boolean if a field has been set.

### GetManual

`func (o *InlineResponse200205BusyHourSchedule) GetManual() NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *InlineResponse200205BusyHourSchedule) GetManualOk() (*NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *InlineResponse200205BusyHourSchedule) SetManual(v NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual)`

SetManual sets Manual field to given value.

### HasManual

`func (o *InlineResponse200205BusyHourSchedule) HasManual() bool`

HasManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


