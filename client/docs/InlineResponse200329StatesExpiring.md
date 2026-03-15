# InlineResponse200329StatesExpiring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of expiring licenses | [optional] 
**Critical** | Pointer to [**InlineResponse200329StatesExpiringCritical**](InlineResponse200329StatesExpiringCritical.md) |  | [optional] 
**Warning** | Pointer to [**InlineResponse200329StatesExpiringWarning**](InlineResponse200329StatesExpiringWarning.md) |  | [optional] 

## Methods

### NewInlineResponse200329StatesExpiring

`func NewInlineResponse200329StatesExpiring() *InlineResponse200329StatesExpiring`

NewInlineResponse200329StatesExpiring instantiates a new InlineResponse200329StatesExpiring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200329StatesExpiringWithDefaults

`func NewInlineResponse200329StatesExpiringWithDefaults() *InlineResponse200329StatesExpiring`

NewInlineResponse200329StatesExpiringWithDefaults instantiates a new InlineResponse200329StatesExpiring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *InlineResponse200329StatesExpiring) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse200329StatesExpiring) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse200329StatesExpiring) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse200329StatesExpiring) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCritical

`func (o *InlineResponse200329StatesExpiring) GetCritical() InlineResponse200329StatesExpiringCritical`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *InlineResponse200329StatesExpiring) GetCriticalOk() (*InlineResponse200329StatesExpiringCritical, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *InlineResponse200329StatesExpiring) SetCritical(v InlineResponse200329StatesExpiringCritical)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *InlineResponse200329StatesExpiring) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetWarning

`func (o *InlineResponse200329StatesExpiring) GetWarning() InlineResponse200329StatesExpiringWarning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *InlineResponse200329StatesExpiring) GetWarningOk() (*InlineResponse200329StatesExpiringWarning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *InlineResponse200329StatesExpiring) SetWarning(v InlineResponse200329StatesExpiringWarning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *InlineResponse200329StatesExpiring) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


