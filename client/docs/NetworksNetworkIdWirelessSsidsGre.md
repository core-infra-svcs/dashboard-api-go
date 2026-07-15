# NetworksNetworkIdWirelessSsidsGre

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concentrator** | Pointer to [**NetworksNetworkIdWirelessSsidsGreConcentrator**](NetworksNetworkIdWirelessSsidsGreConcentrator.md) |  | [optional] 
**Key** | Pointer to **int32** | Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel. | [optional] 
**ClientIsolation** | Pointer to **bool** | Whether or not client isolation is enabled. | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsGre

`func NewNetworksNetworkIdWirelessSsidsGre() *NetworksNetworkIdWirelessSsidsGre`

NewNetworksNetworkIdWirelessSsidsGre instantiates a new NetworksNetworkIdWirelessSsidsGre object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsGreWithDefaults

`func NewNetworksNetworkIdWirelessSsidsGreWithDefaults() *NetworksNetworkIdWirelessSsidsGre`

NewNetworksNetworkIdWirelessSsidsGreWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsGre object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcentrator

`func (o *NetworksNetworkIdWirelessSsidsGre) GetConcentrator() NetworksNetworkIdWirelessSsidsGreConcentrator`

GetConcentrator returns the Concentrator field if non-nil, zero value otherwise.

### GetConcentratorOk

`func (o *NetworksNetworkIdWirelessSsidsGre) GetConcentratorOk() (*NetworksNetworkIdWirelessSsidsGreConcentrator, bool)`

GetConcentratorOk returns a tuple with the Concentrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcentrator

`func (o *NetworksNetworkIdWirelessSsidsGre) SetConcentrator(v NetworksNetworkIdWirelessSsidsGreConcentrator)`

SetConcentrator sets Concentrator field to given value.

### HasConcentrator

`func (o *NetworksNetworkIdWirelessSsidsGre) HasConcentrator() bool`

HasConcentrator returns a boolean if a field has been set.

### GetKey

`func (o *NetworksNetworkIdWirelessSsidsGre) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NetworksNetworkIdWirelessSsidsGre) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NetworksNetworkIdWirelessSsidsGre) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *NetworksNetworkIdWirelessSsidsGre) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetClientIsolation

`func (o *NetworksNetworkIdWirelessSsidsGre) GetClientIsolation() bool`

GetClientIsolation returns the ClientIsolation field if non-nil, zero value otherwise.

### GetClientIsolationOk

`func (o *NetworksNetworkIdWirelessSsidsGre) GetClientIsolationOk() (*bool, bool)`

GetClientIsolationOk returns a tuple with the ClientIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIsolation

`func (o *NetworksNetworkIdWirelessSsidsGre) SetClientIsolation(v bool)`

SetClientIsolation sets ClientIsolation field to given value.

### HasClientIsolation

`func (o *NetworksNetworkIdWirelessSsidsGre) HasClientIsolation() bool`

HasClientIsolation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


