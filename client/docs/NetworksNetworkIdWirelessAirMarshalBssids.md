# NetworksNetworkIdWirelessAirMarshalBssids

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bssid** | Pointer to **string** | BSSID of the observed SSID | [optional] 
**Contained** | Pointer to **bool** | Whether this BSSID is contained | [optional] 
**DetectedBy** | Pointer to [**[]NetworksNetworkIdWirelessAirMarshalDetectedBy**](NetworksNetworkIdWirelessAirMarshalDetectedBy.md) | Access points that detected the BSSID | [optional] 

## Methods

### NewNetworksNetworkIdWirelessAirMarshalBssids

`func NewNetworksNetworkIdWirelessAirMarshalBssids() *NetworksNetworkIdWirelessAirMarshalBssids`

NewNetworksNetworkIdWirelessAirMarshalBssids instantiates a new NetworksNetworkIdWirelessAirMarshalBssids object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessAirMarshalBssidsWithDefaults

`func NewNetworksNetworkIdWirelessAirMarshalBssidsWithDefaults() *NetworksNetworkIdWirelessAirMarshalBssids`

NewNetworksNetworkIdWirelessAirMarshalBssidsWithDefaults instantiates a new NetworksNetworkIdWirelessAirMarshalBssids object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBssid

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetContained

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) GetContained() bool`

GetContained returns the Contained field if non-nil, zero value otherwise.

### GetContainedOk

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) GetContainedOk() (*bool, bool)`

GetContainedOk returns a tuple with the Contained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContained

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) SetContained(v bool)`

SetContained sets Contained field to given value.

### HasContained

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) HasContained() bool`

HasContained returns a boolean if a field has been set.

### GetDetectedBy

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) GetDetectedBy() []NetworksNetworkIdWirelessAirMarshalDetectedBy`

GetDetectedBy returns the DetectedBy field if non-nil, zero value otherwise.

### GetDetectedByOk

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) GetDetectedByOk() (*[]NetworksNetworkIdWirelessAirMarshalDetectedBy, bool)`

GetDetectedByOk returns a tuple with the DetectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBy

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) SetDetectedBy(v []NetworksNetworkIdWirelessAirMarshalDetectedBy)`

SetDetectedBy sets DetectedBy field to given value.

### HasDetectedBy

`func (o *NetworksNetworkIdWirelessAirMarshalBssids) HasDetectedBy() bool`

HasDetectedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


