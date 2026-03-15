# InlineResponse200193

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

## Methods

### NewInlineResponse200193

`func NewInlineResponse200193() *InlineResponse200193`

NewInlineResponse200193 instantiates a new InlineResponse200193 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200193WithDefaults

`func NewInlineResponse200193WithDefaults() *InlineResponse200193`

NewInlineResponse200193WithDefaults instantiates a new InlineResponse200193 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsid

`func (o *InlineResponse200193) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse200193) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse200193) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse200193) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetBssids

`func (o *InlineResponse200193) GetBssids() []NetworksNetworkIdWirelessAirMarshalBssids`

GetBssids returns the Bssids field if non-nil, zero value otherwise.

### GetBssidsOk

`func (o *InlineResponse200193) GetBssidsOk() (*[]NetworksNetworkIdWirelessAirMarshalBssids, bool)`

GetBssidsOk returns a tuple with the Bssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssids

`func (o *InlineResponse200193) SetBssids(v []NetworksNetworkIdWirelessAirMarshalBssids)`

SetBssids sets Bssids field to given value.

### HasBssids

`func (o *InlineResponse200193) HasBssids() bool`

HasBssids returns a boolean if a field has been set.

### GetChannels

`func (o *InlineResponse200193) GetChannels() []int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *InlineResponse200193) GetChannelsOk() (*[]int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *InlineResponse200193) SetChannels(v []int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *InlineResponse200193) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetFirstSeen

`func (o *InlineResponse200193) GetFirstSeen() int32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *InlineResponse200193) GetFirstSeenOk() (*int32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *InlineResponse200193) SetFirstSeen(v int32)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *InlineResponse200193) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *InlineResponse200193) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *InlineResponse200193) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *InlineResponse200193) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *InlineResponse200193) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetWiredMacs

`func (o *InlineResponse200193) GetWiredMacs() []string`

GetWiredMacs returns the WiredMacs field if non-nil, zero value otherwise.

### GetWiredMacsOk

`func (o *InlineResponse200193) GetWiredMacsOk() (*[]string, bool)`

GetWiredMacsOk returns a tuple with the WiredMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredMacs

`func (o *InlineResponse200193) SetWiredMacs(v []string)`

SetWiredMacs sets WiredMacs field to given value.

### HasWiredMacs

`func (o *InlineResponse200193) HasWiredMacs() bool`

HasWiredMacs returns a boolean if a field has been set.

### GetWiredVlans

`func (o *InlineResponse200193) GetWiredVlans() []int32`

GetWiredVlans returns the WiredVlans field if non-nil, zero value otherwise.

### GetWiredVlansOk

`func (o *InlineResponse200193) GetWiredVlansOk() (*[]int32, bool)`

GetWiredVlansOk returns a tuple with the WiredVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredVlans

`func (o *InlineResponse200193) SetWiredVlans(v []int32)`

SetWiredVlans sets WiredVlans field to given value.

### HasWiredVlans

`func (o *InlineResponse200193) HasWiredVlans() bool`

HasWiredVlans returns a boolean if a field has been set.

### GetWiredLastSeen

`func (o *InlineResponse200193) GetWiredLastSeen() int32`

GetWiredLastSeen returns the WiredLastSeen field if non-nil, zero value otherwise.

### GetWiredLastSeenOk

`func (o *InlineResponse200193) GetWiredLastSeenOk() (*int32, bool)`

GetWiredLastSeenOk returns a tuple with the WiredLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredLastSeen

`func (o *InlineResponse200193) SetWiredLastSeen(v int32)`

SetWiredLastSeen sets WiredLastSeen field to given value.

### HasWiredLastSeen

`func (o *InlineResponse200193) HasWiredLastSeen() bool`

HasWiredLastSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


