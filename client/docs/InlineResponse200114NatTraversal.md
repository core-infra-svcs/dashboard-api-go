# InlineResponse200114NatTraversal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Whether NAT traversal is automatic or manually configured. | [optional] 
**PublicIp** | Pointer to **string** | The public IPv4 address used for manual NAT traversal. | [optional] 
**PublicPort** | Pointer to **int32** | The public port used for manual NAT traversal. | [optional] 

## Methods

### NewInlineResponse200114NatTraversal

`func NewInlineResponse200114NatTraversal() *InlineResponse200114NatTraversal`

NewInlineResponse200114NatTraversal instantiates a new InlineResponse200114NatTraversal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200114NatTraversalWithDefaults

`func NewInlineResponse200114NatTraversalWithDefaults() *InlineResponse200114NatTraversal`

NewInlineResponse200114NatTraversalWithDefaults instantiates a new InlineResponse200114NatTraversal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200114NatTraversal) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200114NatTraversal) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200114NatTraversal) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200114NatTraversal) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse200114NatTraversal) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse200114NatTraversal) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse200114NatTraversal) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse200114NatTraversal) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicPort

`func (o *InlineResponse200114NatTraversal) GetPublicPort() int32`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *InlineResponse200114NatTraversal) GetPublicPortOk() (*int32, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *InlineResponse200114NatTraversal) SetPublicPort(v int32)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *InlineResponse200114NatTraversal) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


