# InlineResponse200199

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200199Rules**](InlineResponse200199Rules.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). | [optional] 
**AllowLanAccess** | Pointer to **bool** | Allows wireless client access to local LAN (boolean value - true allows access and false denies access) | [optional] 

## Methods

### NewInlineResponse200199

`func NewInlineResponse200199() *InlineResponse200199`

NewInlineResponse200199 instantiates a new InlineResponse200199 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200199WithDefaults

`func NewInlineResponse200199WithDefaults() *InlineResponse200199`

NewInlineResponse200199WithDefaults instantiates a new InlineResponse200199 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineResponse200199) GetRules() []InlineResponse200199Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200199) GetRulesOk() (*[]InlineResponse200199Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200199) SetRules(v []InlineResponse200199Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200199) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetAllowLanAccess

`func (o *InlineResponse200199) GetAllowLanAccess() bool`

GetAllowLanAccess returns the AllowLanAccess field if non-nil, zero value otherwise.

### GetAllowLanAccessOk

`func (o *InlineResponse200199) GetAllowLanAccessOk() (*bool, bool)`

GetAllowLanAccessOk returns a tuple with the AllowLanAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLanAccess

`func (o *InlineResponse200199) SetAllowLanAccess(v bool)`

SetAllowLanAccess sets AllowLanAccess field to given value.

### HasAllowLanAccess

`func (o *InlineResponse200199) HasAllowLanAccess() bool`

HasAllowLanAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


