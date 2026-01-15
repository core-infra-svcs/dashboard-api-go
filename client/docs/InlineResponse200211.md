# InlineResponse200211

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200211Rules**](InlineResponse200211Rules.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). | [optional] 
**AllowLanAccess** | Pointer to **bool** | Allows wireless client access to local LAN (boolean value - true allows access and false denies access) | [optional] 

## Methods

### NewInlineResponse200211

`func NewInlineResponse200211() *InlineResponse200211`

NewInlineResponse200211 instantiates a new InlineResponse200211 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200211WithDefaults

`func NewInlineResponse200211WithDefaults() *InlineResponse200211`

NewInlineResponse200211WithDefaults instantiates a new InlineResponse200211 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineResponse200211) GetRules() []InlineResponse200211Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200211) GetRulesOk() (*[]InlineResponse200211Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200211) SetRules(v []InlineResponse200211Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200211) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetAllowLanAccess

`func (o *InlineResponse200211) GetAllowLanAccess() bool`

GetAllowLanAccess returns the AllowLanAccess field if non-nil, zero value otherwise.

### GetAllowLanAccessOk

`func (o *InlineResponse200211) GetAllowLanAccessOk() (*bool, bool)`

GetAllowLanAccessOk returns a tuple with the AllowLanAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLanAccess

`func (o *InlineResponse200211) SetAllowLanAccess(v bool)`

SetAllowLanAccess sets AllowLanAccess field to given value.

### HasAllowLanAccess

`func (o *InlineResponse200211) HasAllowLanAccess() bool`

HasAllowLanAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


