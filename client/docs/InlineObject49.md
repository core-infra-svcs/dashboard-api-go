# InlineObject49

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** | A string indicating the rule for which IPs are allowed to use the specified service. Acceptable values are \&quot;blocked\&quot; (no remote IPs can access the service), \&quot;restricted\&quot; (only allowed IPs can access the service), and \&quot;unrestriced\&quot; (any remote IP can access the service). This field is required | 
**AllowedIps** | Pointer to **[]string** | An array of allowed CIDRs that can access the service. This field is required if \&quot;access\&quot; is set to \&quot;restricted\&quot;. Otherwise this field is ignored | [optional] 

## Methods

### NewInlineObject49

`func NewInlineObject49(access string, ) *InlineObject49`

NewInlineObject49 instantiates a new InlineObject49 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject49WithDefaults

`func NewInlineObject49WithDefaults() *InlineObject49`

NewInlineObject49WithDefaults instantiates a new InlineObject49 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *InlineObject49) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineObject49) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineObject49) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllowedIps

`func (o *InlineObject49) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *InlineObject49) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *InlineObject49) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *InlineObject49) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


