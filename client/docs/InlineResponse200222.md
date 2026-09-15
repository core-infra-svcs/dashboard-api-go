# InlineResponse200222

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ssid** | Pointer to **string** | SSID name | [optional] 
**Bssids** | Pointer to [**[]NetworksNetworkIdWirelessAirMarshalBssids**](NetworksNetworkIdWirelessAirMarshalBssids.md) | BSSIDs broadcasting the SSID | [optional] 
**Channels** | Pointer to **[]int32** | Channels where the SSID was observed | [optional] 
**FirstSeen** | Pointer to **int32** | First time the SSID was observed (epoch seconds) | [optional] 
**LastSeen** | Pointer to **int32** | Most recent time the SSID was observed (epoch seconds) | [optional] 
**WiredMacs** | Pointer to **[]string** | MAC addresses observed on the SSID | [optional] 
**WiredVlans** | Pointer to **[]int32** | VLAN IDs observed on the SSID | [optional] 
**WiredLastSeen** | Pointer to **int32** | Last time observed on the SSID (epoch seconds) | [optional] 
**Manufacturers** | Pointer to **[]string** | Vendor names resolved from the OUI of each broadcasting BSSID (not the OUI hex prefix itself). Strings come from the IEEE OUI registry and are not casing-normalized -- callers should compare case-insensitively. Example values: \&quot;Linksys\&quot;, \&quot;Cisco Systems\&quot;, \&quot;OMRON TATEISI Electronics\&quot;. Deduplicated and capped at 10 entries; empty when no OUI can be resolved. | [optional] 
**Encryption** | Pointer to **NullableString** | High-level encryption mode observed for the SSID; null when it cannot be inferred | [optional] 
**Types** | Pointer to **[]string** | Threat classifications applied to this SSID. \&quot;rogue\&quot;: seen on LAN and not allow-listed by an is_known containment rule. \&quot;spoof\&quot;: SSID name matches an enabled SSID on this network. Empty when no classification applies. | [optional] 

## Methods

### NewInlineResponse200222

`func NewInlineResponse200222() *InlineResponse200222`

NewInlineResponse200222 instantiates a new InlineResponse200222 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200222WithDefaults

`func NewInlineResponse200222WithDefaults() *InlineResponse200222`

NewInlineResponse200222WithDefaults instantiates a new InlineResponse200222 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsid

`func (o *InlineResponse200222) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse200222) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse200222) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse200222) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetBssids

`func (o *InlineResponse200222) GetBssids() []NetworksNetworkIdWirelessAirMarshalBssids`

GetBssids returns the Bssids field if non-nil, zero value otherwise.

### GetBssidsOk

`func (o *InlineResponse200222) GetBssidsOk() (*[]NetworksNetworkIdWirelessAirMarshalBssids, bool)`

GetBssidsOk returns a tuple with the Bssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssids

`func (o *InlineResponse200222) SetBssids(v []NetworksNetworkIdWirelessAirMarshalBssids)`

SetBssids sets Bssids field to given value.

### HasBssids

`func (o *InlineResponse200222) HasBssids() bool`

HasBssids returns a boolean if a field has been set.

### GetChannels

`func (o *InlineResponse200222) GetChannels() []int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *InlineResponse200222) GetChannelsOk() (*[]int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *InlineResponse200222) SetChannels(v []int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *InlineResponse200222) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetFirstSeen

`func (o *InlineResponse200222) GetFirstSeen() int32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *InlineResponse200222) GetFirstSeenOk() (*int32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *InlineResponse200222) SetFirstSeen(v int32)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *InlineResponse200222) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *InlineResponse200222) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *InlineResponse200222) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *InlineResponse200222) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *InlineResponse200222) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetWiredMacs

`func (o *InlineResponse200222) GetWiredMacs() []string`

GetWiredMacs returns the WiredMacs field if non-nil, zero value otherwise.

### GetWiredMacsOk

`func (o *InlineResponse200222) GetWiredMacsOk() (*[]string, bool)`

GetWiredMacsOk returns a tuple with the WiredMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredMacs

`func (o *InlineResponse200222) SetWiredMacs(v []string)`

SetWiredMacs sets WiredMacs field to given value.

### HasWiredMacs

`func (o *InlineResponse200222) HasWiredMacs() bool`

HasWiredMacs returns a boolean if a field has been set.

### GetWiredVlans

`func (o *InlineResponse200222) GetWiredVlans() []int32`

GetWiredVlans returns the WiredVlans field if non-nil, zero value otherwise.

### GetWiredVlansOk

`func (o *InlineResponse200222) GetWiredVlansOk() (*[]int32, bool)`

GetWiredVlansOk returns a tuple with the WiredVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredVlans

`func (o *InlineResponse200222) SetWiredVlans(v []int32)`

SetWiredVlans sets WiredVlans field to given value.

### HasWiredVlans

`func (o *InlineResponse200222) HasWiredVlans() bool`

HasWiredVlans returns a boolean if a field has been set.

### GetWiredLastSeen

`func (o *InlineResponse200222) GetWiredLastSeen() int32`

GetWiredLastSeen returns the WiredLastSeen field if non-nil, zero value otherwise.

### GetWiredLastSeenOk

`func (o *InlineResponse200222) GetWiredLastSeenOk() (*int32, bool)`

GetWiredLastSeenOk returns a tuple with the WiredLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredLastSeen

`func (o *InlineResponse200222) SetWiredLastSeen(v int32)`

SetWiredLastSeen sets WiredLastSeen field to given value.

### HasWiredLastSeen

`func (o *InlineResponse200222) HasWiredLastSeen() bool`

HasWiredLastSeen returns a boolean if a field has been set.

### GetManufacturers

`func (o *InlineResponse200222) GetManufacturers() []string`

GetManufacturers returns the Manufacturers field if non-nil, zero value otherwise.

### GetManufacturersOk

`func (o *InlineResponse200222) GetManufacturersOk() (*[]string, bool)`

GetManufacturersOk returns a tuple with the Manufacturers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturers

`func (o *InlineResponse200222) SetManufacturers(v []string)`

SetManufacturers sets Manufacturers field to given value.

### HasManufacturers

`func (o *InlineResponse200222) HasManufacturers() bool`

HasManufacturers returns a boolean if a field has been set.

### GetEncryption

`func (o *InlineResponse200222) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *InlineResponse200222) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *InlineResponse200222) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *InlineResponse200222) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### SetEncryptionNil

`func (o *InlineResponse200222) SetEncryptionNil(b bool)`

 SetEncryptionNil sets the value for Encryption to be an explicit nil

### UnsetEncryption
`func (o *InlineResponse200222) UnsetEncryption()`

UnsetEncryption ensures that no value is present for Encryption, not even an explicit nil
### GetTypes

`func (o *InlineResponse200222) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *InlineResponse200222) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *InlineResponse200222) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *InlineResponse200222) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


