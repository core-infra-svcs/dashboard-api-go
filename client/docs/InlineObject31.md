# InlineObject31

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the route provided. The only supported value for now is &#39;BGP&#39;. | 
**Vrf** | [**DevicesSerialLiveToolsRouteDetailsVrf**](DevicesSerialLiveToolsRouteDetailsVrf.md) |  | 
**Destination** | [**DevicesSerialLiveToolsRouteDetailsDestination**](DevicesSerialLiveToolsRouteDetailsDestination.md) |  | 
**NextHop** | [**DevicesSerialLiveToolsRouteDetailsNextHop**](DevicesSerialLiveToolsRouteDetailsNextHop.md) |  | 
**Callback** | Pointer to [**DevicesSerialLiveToolsArpTableCallback**](DevicesSerialLiveToolsArpTableCallback.md) |  | [optional] 

## Methods

### NewInlineObject31

`func NewInlineObject31(type_ string, vrf DevicesSerialLiveToolsRouteDetailsVrf, destination DevicesSerialLiveToolsRouteDetailsDestination, nextHop DevicesSerialLiveToolsRouteDetailsNextHop, ) *InlineObject31`

NewInlineObject31 instantiates a new InlineObject31 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject31WithDefaults

`func NewInlineObject31WithDefaults() *InlineObject31`

NewInlineObject31WithDefaults instantiates a new InlineObject31 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject31) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject31) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject31) SetType(v string)`

SetType sets Type field to given value.


### GetVrf

`func (o *InlineObject31) GetVrf() DevicesSerialLiveToolsRouteDetailsVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject31) GetVrfOk() (*DevicesSerialLiveToolsRouteDetailsVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject31) SetVrf(v DevicesSerialLiveToolsRouteDetailsVrf)`

SetVrf sets Vrf field to given value.


### GetDestination

`func (o *InlineObject31) GetDestination() DevicesSerialLiveToolsRouteDetailsDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineObject31) GetDestinationOk() (*DevicesSerialLiveToolsRouteDetailsDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineObject31) SetDestination(v DevicesSerialLiveToolsRouteDetailsDestination)`

SetDestination sets Destination field to given value.


### GetNextHop

`func (o *InlineObject31) GetNextHop() DevicesSerialLiveToolsRouteDetailsNextHop`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *InlineObject31) GetNextHopOk() (*DevicesSerialLiveToolsRouteDetailsNextHop, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *InlineObject31) SetNextHop(v DevicesSerialLiveToolsRouteDetailsNextHop)`

SetNextHop sets NextHop field to given value.


### GetCallback

`func (o *InlineObject31) GetCallback() DevicesSerialLiveToolsArpTableCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineObject31) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineObject31) SetCallback(v DevicesSerialLiveToolsArpTableCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineObject31) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


