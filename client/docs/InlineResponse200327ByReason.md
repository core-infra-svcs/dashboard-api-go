# InlineResponse200327ByReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Reason source type. Null when the source type cannot be determined. | [optional] 
**Code** | Pointer to **NullableInt32** | Reason code from the reason source. Range: 0-999. Null when the source has invalid reason. | [optional] 
**Reason** | Pointer to **NullableString** | Human-readable failure reason. Null when the source has invalid reason. | [optional] 
**Count** | Pointer to **int32** | Failures for this reason. | [optional] 

## Methods

### NewInlineResponse200327ByReason

`func NewInlineResponse200327ByReason() *InlineResponse200327ByReason`

NewInlineResponse200327ByReason instantiates a new InlineResponse200327ByReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327ByReasonWithDefaults

`func NewInlineResponse200327ByReasonWithDefaults() *InlineResponse200327ByReason`

NewInlineResponse200327ByReasonWithDefaults instantiates a new InlineResponse200327ByReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200327ByReason) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200327ByReason) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200327ByReason) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200327ByReason) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *InlineResponse200327ByReason) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InlineResponse200327ByReason) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCode

`func (o *InlineResponse200327ByReason) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse200327ByReason) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse200327ByReason) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse200327ByReason) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *InlineResponse200327ByReason) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *InlineResponse200327ByReason) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetReason

`func (o *InlineResponse200327ByReason) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineResponse200327ByReason) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineResponse200327ByReason) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InlineResponse200327ByReason) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *InlineResponse200327ByReason) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *InlineResponse200327ByReason) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetCount

`func (o *InlineResponse200327ByReason) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse200327ByReason) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse200327ByReason) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse200327ByReason) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


