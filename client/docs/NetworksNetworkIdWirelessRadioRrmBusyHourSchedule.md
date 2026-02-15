# NetworksNetworkIdWirelessRadioRrmBusyHourSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The Busy Hour mode applied to the network when minimizeChanges is enabled. Must be one of &#39;automatic&#39; or &#39;manual&#39;. Automatic busy hour is only available on firmware versions &gt;&#x3D; MR 27.0 | [optional] 
**Manual** | Pointer to [**NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual**](NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdWirelessRadioRrmBusyHourSchedule

`func NewNetworksNetworkIdWirelessRadioRrmBusyHourSchedule() *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule`

NewNetworksNetworkIdWirelessRadioRrmBusyHourSchedule instantiates a new NetworksNetworkIdWirelessRadioRrmBusyHourSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessRadioRrmBusyHourScheduleWithDefaults

`func NewNetworksNetworkIdWirelessRadioRrmBusyHourScheduleWithDefaults() *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule`

NewNetworksNetworkIdWirelessRadioRrmBusyHourScheduleWithDefaults instantiates a new NetworksNetworkIdWirelessRadioRrmBusyHourSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetManual

`func (o *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule) GetManual() NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule) GetManualOk() (*NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule) SetManual(v NetworksNetworkIdWirelessRadioRrmBusyHourScheduleManual)`

SetManual sets Manual field to given value.

### HasManual

`func (o *NetworksNetworkIdWirelessRadioRrmBusyHourSchedule) HasManual() bool`

HasManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


