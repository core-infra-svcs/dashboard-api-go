# InlineResponse20111Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**Type** | Pointer to **string** | The type of route defined | [optional] 
**Destination** | Pointer to [**InlineResponse20111RequestDestination**](InlineResponse20111RequestDestination.md) |  | [optional] 
**NextHop** | Pointer to [**DevicesSerialLiveToolsRoutingTableLookupsNextHop**](DevicesSerialLiveToolsRoutingTableLookupsNextHop.md) |  | [optional] 
**Vpn** | Pointer to [**DevicesSerialLiveToolsRoutingTableLookupsVpn**](DevicesSerialLiveToolsRoutingTableLookupsVpn.md) |  | [optional] 
**Vrf** | Pointer to [**DevicesSerialLiveToolsRoutingTableLookupsVrf**](DevicesSerialLiveToolsRoutingTableLookupsVrf.md) |  | [optional] 

## Methods

### NewInlineResponse20111Request

`func NewInlineResponse20111Request() *InlineResponse20111Request`

NewInlineResponse20111Request instantiates a new InlineResponse20111Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20111RequestWithDefaults

`func NewInlineResponse20111RequestWithDefaults() *InlineResponse20111Request`

NewInlineResponse20111RequestWithDefaults instantiates a new InlineResponse20111Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse20111Request) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20111Request) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20111Request) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20111Request) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20111Request) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20111Request) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20111Request) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20111Request) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDestination

`func (o *InlineResponse20111Request) GetDestination() InlineResponse20111RequestDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20111Request) GetDestinationOk() (*InlineResponse20111RequestDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20111Request) SetDestination(v InlineResponse20111RequestDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineResponse20111Request) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetNextHop

`func (o *InlineResponse20111Request) GetNextHop() DevicesSerialLiveToolsRoutingTableLookupsNextHop`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *InlineResponse20111Request) GetNextHopOk() (*DevicesSerialLiveToolsRoutingTableLookupsNextHop, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *InlineResponse20111Request) SetNextHop(v DevicesSerialLiveToolsRoutingTableLookupsNextHop)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *InlineResponse20111Request) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse20111Request) GetVpn() DevicesSerialLiveToolsRoutingTableLookupsVpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse20111Request) GetVpnOk() (*DevicesSerialLiveToolsRoutingTableLookupsVpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse20111Request) SetVpn(v DevicesSerialLiveToolsRoutingTableLookupsVpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse20111Request) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse20111Request) GetVrf() DevicesSerialLiveToolsRoutingTableLookupsVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20111Request) GetVrfOk() (*DevicesSerialLiveToolsRoutingTableLookupsVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20111Request) SetVrf(v DevicesSerialLiveToolsRoutingTableLookupsVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20111Request) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


