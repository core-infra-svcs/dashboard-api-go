# InlineResponse200273StatesExpiring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of expiring licenses | [optional] 
**Critical** | Pointer to [**InlineResponse200273StatesExpiringCritical**](InlineResponse200273StatesExpiringCritical.md) |  | [optional] 
**Warning** | Pointer to [**InlineResponse200273StatesExpiringWarning**](InlineResponse200273StatesExpiringWarning.md) |  | [optional] 

## Methods

### NewInlineResponse200273StatesExpiring

`func NewInlineResponse200273StatesExpiring() *InlineResponse200273StatesExpiring`

NewInlineResponse200273StatesExpiring instantiates a new InlineResponse200273StatesExpiring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200273StatesExpiringWithDefaults

`func NewInlineResponse200273StatesExpiringWithDefaults() *InlineResponse200273StatesExpiring`

NewInlineResponse200273StatesExpiringWithDefaults instantiates a new InlineResponse200273StatesExpiring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *InlineResponse200273StatesExpiring) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse200273StatesExpiring) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse200273StatesExpiring) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse200273StatesExpiring) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCritical

`func (o *InlineResponse200273StatesExpiring) GetCritical() InlineResponse200273StatesExpiringCritical`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *InlineResponse200273StatesExpiring) GetCriticalOk() (*InlineResponse200273StatesExpiringCritical, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *InlineResponse200273StatesExpiring) SetCritical(v InlineResponse200273StatesExpiringCritical)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *InlineResponse200273StatesExpiring) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetWarning

`func (o *InlineResponse200273StatesExpiring) GetWarning() InlineResponse200273StatesExpiringWarning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *InlineResponse200273StatesExpiring) GetWarningOk() (*InlineResponse200273StatesExpiringWarning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *InlineResponse200273StatesExpiring) SetWarning(v InlineResponse200273StatesExpiringWarning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *InlineResponse200273StatesExpiring) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


