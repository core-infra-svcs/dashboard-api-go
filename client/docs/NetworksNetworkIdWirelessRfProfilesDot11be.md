# NetworksNetworkIdWirelessRfProfilesDot11be

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether 802.11be is enabled on Wi-Fi 7 capable APs | [optional] 
**Mode** | Pointer to **string** | 802.11be is enabled for all SSIDs or per SSID group | [optional] 
**Ssids** | Pointer to [**NetworksNetworkIdWirelessRfProfilesDot11beSsids**](NetworksNetworkIdWirelessRfProfilesDot11beSsids.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdWirelessRfProfilesDot11be

`func NewNetworksNetworkIdWirelessRfProfilesDot11be() *NetworksNetworkIdWirelessRfProfilesDot11be`

NewNetworksNetworkIdWirelessRfProfilesDot11be instantiates a new NetworksNetworkIdWirelessRfProfilesDot11be object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessRfProfilesDot11beWithDefaults

`func NewNetworksNetworkIdWirelessRfProfilesDot11beWithDefaults() *NetworksNetworkIdWirelessRfProfilesDot11be`

NewNetworksNetworkIdWirelessRfProfilesDot11beWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesDot11be object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMode

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSsids

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) GetSsids() NetworksNetworkIdWirelessRfProfilesDot11beSsids`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) GetSsidsOk() (*NetworksNetworkIdWirelessRfProfilesDot11beSsids, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) SetSsids(v NetworksNetworkIdWirelessRfProfilesDot11beSsids)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *NetworksNetworkIdWirelessRfProfilesDot11be) HasSsids() bool`

HasSsids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


