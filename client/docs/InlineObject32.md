# InlineObject32

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of route defined | [optional] 
**Destination** | Pointer to [**DevicesSerialLiveToolsRoutingTableLookupsDestination**](DevicesSerialLiveToolsRoutingTableLookupsDestination.md) |  | [optional] 
**NextHop** | Pointer to [**DevicesSerialLiveToolsRoutingTableLookupsNextHop**](DevicesSerialLiveToolsRoutingTableLookupsNextHop.md) |  | [optional] 
**Vpn** | Pointer to [**DevicesSerialLiveToolsRoutingTableLookupsVpn**](DevicesSerialLiveToolsRoutingTableLookupsVpn.md) |  | [optional] 
**Vrf** | Pointer to [**DevicesSerialLiveToolsRoutingTableLookupsVrf**](DevicesSerialLiveToolsRoutingTableLookupsVrf.md) |  | [optional] 
**Callback** | Pointer to [**DevicesSerialLiveToolsArpTableCallback**](DevicesSerialLiveToolsArpTableCallback.md) |  | [optional] 

## Methods

### NewInlineObject32

`func NewInlineObject32() *InlineObject32`

NewInlineObject32 instantiates a new InlineObject32 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject32WithDefaults

`func NewInlineObject32WithDefaults() *InlineObject32`

NewInlineObject32WithDefaults instantiates a new InlineObject32 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject32) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject32) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject32) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject32) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDestination

`func (o *InlineObject32) GetDestination() DevicesSerialLiveToolsRoutingTableLookupsDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineObject32) GetDestinationOk() (*DevicesSerialLiveToolsRoutingTableLookupsDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineObject32) SetDestination(v DevicesSerialLiveToolsRoutingTableLookupsDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineObject32) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetNextHop

`func (o *InlineObject32) GetNextHop() DevicesSerialLiveToolsRoutingTableLookupsNextHop`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *InlineObject32) GetNextHopOk() (*DevicesSerialLiveToolsRoutingTableLookupsNextHop, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *InlineObject32) SetNextHop(v DevicesSerialLiveToolsRoutingTableLookupsNextHop)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *InlineObject32) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetVpn

`func (o *InlineObject32) GetVpn() DevicesSerialLiveToolsRoutingTableLookupsVpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineObject32) GetVpnOk() (*DevicesSerialLiveToolsRoutingTableLookupsVpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineObject32) SetVpn(v DevicesSerialLiveToolsRoutingTableLookupsVpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineObject32) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetVrf

`func (o *InlineObject32) GetVrf() DevicesSerialLiveToolsRoutingTableLookupsVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject32) GetVrfOk() (*DevicesSerialLiveToolsRoutingTableLookupsVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject32) SetVrf(v DevicesSerialLiveToolsRoutingTableLookupsVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject32) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetCallback

`func (o *InlineObject32) GetCallback() DevicesSerialLiveToolsArpTableCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineObject32) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineObject32) SetCallback(v DevicesSerialLiveToolsArpTableCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineObject32) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


