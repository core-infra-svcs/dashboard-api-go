# InlineResponse20018ValueDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address (e.g.\&quot;192.168.10.1\&quot;, which is the same as \&quot;192.168.10.1/32\&quot;), or \&quot;any\&quot; | [optional] 

## Methods

### NewInlineResponse20018ValueDestination

`func NewInlineResponse20018ValueDestination() *InlineResponse20018ValueDestination`

NewInlineResponse20018ValueDestination instantiates a new InlineResponse20018ValueDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20018ValueDestinationWithDefaults

`func NewInlineResponse20018ValueDestinationWithDefaults() *InlineResponse20018ValueDestination`

NewInlineResponse20018ValueDestinationWithDefaults instantiates a new InlineResponse20018ValueDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *InlineResponse20018ValueDestination) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20018ValueDestination) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20018ValueDestination) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20018ValueDestination) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse20018ValueDestination) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse20018ValueDestination) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse20018ValueDestination) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse20018ValueDestination) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


