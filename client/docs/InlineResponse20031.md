# InlineResponse20031

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceMac** | Pointer to **string** | Source MAC address | [optional] 
**Ports** | Pointer to [**map[string]InlineResponse20031Ports**](InlineResponse20031Ports.md) | Mapping of ports to lldp and/or cdp information | [optional] 

## Methods

### NewInlineResponse20031

`func NewInlineResponse20031() *InlineResponse20031`

NewInlineResponse20031 instantiates a new InlineResponse20031 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20031WithDefaults

`func NewInlineResponse20031WithDefaults() *InlineResponse20031`

NewInlineResponse20031WithDefaults instantiates a new InlineResponse20031 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceMac

`func (o *InlineResponse20031) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *InlineResponse20031) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *InlineResponse20031) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *InlineResponse20031) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse20031) GetPorts() map[string]InlineResponse20031Ports`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse20031) GetPortsOk() (*map[string]InlineResponse20031Ports, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse20031) SetPorts(v map[string]InlineResponse20031Ports)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse20031) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


