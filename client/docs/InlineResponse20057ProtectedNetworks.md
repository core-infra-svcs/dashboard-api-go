# InlineResponse20057ProtectedNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseDefault** | Pointer to **bool** | Whether special IPv4 addresses should be used (see: https://tools.ietf.org/html/rfc5735) | [optional] 
**IncludedCidr** | Pointer to **[]string** | List of IP addresses or subnets being protected | [optional] 
**ExcludedCidr** | Pointer to **[]string** | List of IP addresses or subnets being excluded from protection | [optional] 

## Methods

### NewInlineResponse20057ProtectedNetworks

`func NewInlineResponse20057ProtectedNetworks() *InlineResponse20057ProtectedNetworks`

NewInlineResponse20057ProtectedNetworks instantiates a new InlineResponse20057ProtectedNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20057ProtectedNetworksWithDefaults

`func NewInlineResponse20057ProtectedNetworksWithDefaults() *InlineResponse20057ProtectedNetworks`

NewInlineResponse20057ProtectedNetworksWithDefaults instantiates a new InlineResponse20057ProtectedNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseDefault

`func (o *InlineResponse20057ProtectedNetworks) GetUseDefault() bool`

GetUseDefault returns the UseDefault field if non-nil, zero value otherwise.

### GetUseDefaultOk

`func (o *InlineResponse20057ProtectedNetworks) GetUseDefaultOk() (*bool, bool)`

GetUseDefaultOk returns a tuple with the UseDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefault

`func (o *InlineResponse20057ProtectedNetworks) SetUseDefault(v bool)`

SetUseDefault sets UseDefault field to given value.

### HasUseDefault

`func (o *InlineResponse20057ProtectedNetworks) HasUseDefault() bool`

HasUseDefault returns a boolean if a field has been set.

### GetIncludedCidr

`func (o *InlineResponse20057ProtectedNetworks) GetIncludedCidr() []string`

GetIncludedCidr returns the IncludedCidr field if non-nil, zero value otherwise.

### GetIncludedCidrOk

`func (o *InlineResponse20057ProtectedNetworks) GetIncludedCidrOk() (*[]string, bool)`

GetIncludedCidrOk returns a tuple with the IncludedCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedCidr

`func (o *InlineResponse20057ProtectedNetworks) SetIncludedCidr(v []string)`

SetIncludedCidr sets IncludedCidr field to given value.

### HasIncludedCidr

`func (o *InlineResponse20057ProtectedNetworks) HasIncludedCidr() bool`

HasIncludedCidr returns a boolean if a field has been set.

### GetExcludedCidr

`func (o *InlineResponse20057ProtectedNetworks) GetExcludedCidr() []string`

GetExcludedCidr returns the ExcludedCidr field if non-nil, zero value otherwise.

### GetExcludedCidrOk

`func (o *InlineResponse20057ProtectedNetworks) GetExcludedCidrOk() (*[]string, bool)`

GetExcludedCidrOk returns a tuple with the ExcludedCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCidr

`func (o *InlineResponse20057ProtectedNetworks) SetExcludedCidr(v []string)`

SetExcludedCidr sets ExcludedCidr field to given value.

### HasExcludedCidr

`func (o *InlineResponse20057ProtectedNetworks) HasExcludedCidr() bool`

HasExcludedCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


