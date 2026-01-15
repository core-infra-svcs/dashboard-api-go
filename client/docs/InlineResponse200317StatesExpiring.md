# InlineResponse200317StatesExpiring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of expiring licenses | [optional] 
**Critical** | Pointer to [**InlineResponse200317StatesExpiringCritical**](InlineResponse200317StatesExpiringCritical.md) |  | [optional] 
**Warning** | Pointer to [**InlineResponse200317StatesExpiringWarning**](InlineResponse200317StatesExpiringWarning.md) |  | [optional] 

## Methods

### NewInlineResponse200317StatesExpiring

`func NewInlineResponse200317StatesExpiring() *InlineResponse200317StatesExpiring`

NewInlineResponse200317StatesExpiring instantiates a new InlineResponse200317StatesExpiring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200317StatesExpiringWithDefaults

`func NewInlineResponse200317StatesExpiringWithDefaults() *InlineResponse200317StatesExpiring`

NewInlineResponse200317StatesExpiringWithDefaults instantiates a new InlineResponse200317StatesExpiring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *InlineResponse200317StatesExpiring) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse200317StatesExpiring) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse200317StatesExpiring) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse200317StatesExpiring) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCritical

`func (o *InlineResponse200317StatesExpiring) GetCritical() InlineResponse200317StatesExpiringCritical`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *InlineResponse200317StatesExpiring) GetCriticalOk() (*InlineResponse200317StatesExpiringCritical, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *InlineResponse200317StatesExpiring) SetCritical(v InlineResponse200317StatesExpiringCritical)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *InlineResponse200317StatesExpiring) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetWarning

`func (o *InlineResponse200317StatesExpiring) GetWarning() InlineResponse200317StatesExpiringWarning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *InlineResponse200317StatesExpiring) GetWarningOk() (*InlineResponse200317StatesExpiringWarning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *InlineResponse200317StatesExpiring) SetWarning(v InlineResponse200317StatesExpiringWarning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *InlineResponse200317StatesExpiring) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


