# InlineResponse2009

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to [**DevicesSerialAppliancePrefixesDelegatedOrigin**](DevicesSerialAppliancePrefixesDelegatedOrigin.md) |  | [optional] 
**Prefix** | **string** | Delegated IPv6 prefix in CIDR notation. | 
**Counts** | Pointer to [**DevicesSerialAppliancePrefixesDelegatedCounts**](DevicesSerialAppliancePrefixesDelegatedCounts.md) |  | [optional] 
**Method** | **string** | Delegation method. | 
**StaticDelegatedPrefixId** | Pointer to **string** | Static delegated prefix id. Present when method is manual. | [optional] 
**Description** | Pointer to **string** | Description for a manual delegated prefix. | [optional] 
**IsPreferred** | **bool** | Whether this delegated prefix is preferred. | 
**ExpiresAt** | Pointer to **NullableTime** | Expiration time of the delegated prefix. | [optional] 

## Methods

### NewInlineResponse2009

`func NewInlineResponse2009(prefix string, method string, isPreferred bool, ) *InlineResponse2009`

NewInlineResponse2009 instantiates a new InlineResponse2009 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2009WithDefaults

`func NewInlineResponse2009WithDefaults() *InlineResponse2009`

NewInlineResponse2009WithDefaults instantiates a new InlineResponse2009 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *InlineResponse2009) GetOrigin() DevicesSerialAppliancePrefixesDelegatedOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *InlineResponse2009) GetOriginOk() (*DevicesSerialAppliancePrefixesDelegatedOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *InlineResponse2009) SetOrigin(v DevicesSerialAppliancePrefixesDelegatedOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *InlineResponse2009) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPrefix

`func (o *InlineResponse2009) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *InlineResponse2009) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *InlineResponse2009) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetCounts

`func (o *InlineResponse2009) GetCounts() DevicesSerialAppliancePrefixesDelegatedCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse2009) GetCountsOk() (*DevicesSerialAppliancePrefixesDelegatedCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse2009) SetCounts(v DevicesSerialAppliancePrefixesDelegatedCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse2009) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetMethod

`func (o *InlineResponse2009) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineResponse2009) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineResponse2009) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetStaticDelegatedPrefixId

`func (o *InlineResponse2009) GetStaticDelegatedPrefixId() string`

GetStaticDelegatedPrefixId returns the StaticDelegatedPrefixId field if non-nil, zero value otherwise.

### GetStaticDelegatedPrefixIdOk

`func (o *InlineResponse2009) GetStaticDelegatedPrefixIdOk() (*string, bool)`

GetStaticDelegatedPrefixIdOk returns a tuple with the StaticDelegatedPrefixId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDelegatedPrefixId

`func (o *InlineResponse2009) SetStaticDelegatedPrefixId(v string)`

SetStaticDelegatedPrefixId sets StaticDelegatedPrefixId field to given value.

### HasStaticDelegatedPrefixId

`func (o *InlineResponse2009) HasStaticDelegatedPrefixId() bool`

HasStaticDelegatedPrefixId returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse2009) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse2009) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse2009) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse2009) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPreferred

`func (o *InlineResponse2009) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *InlineResponse2009) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *InlineResponse2009) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.


### GetExpiresAt

`func (o *InlineResponse2009) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InlineResponse2009) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InlineResponse2009) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InlineResponse2009) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *InlineResponse2009) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *InlineResponse2009) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


