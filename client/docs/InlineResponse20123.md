# InlineResponse20123

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integrated** | Pointer to [**InlineResponse20123Integrated**](InlineResponse20123Integrated.md) |  | [optional] 
**LastUsedAt** | Pointer to **string** | ISO8601Z timestamp when integration was last used | [optional] 
**External** | Pointer to [**InlineResponse20123External**](InlineResponse20123External.md) |  | [optional] 
**Status** | Pointer to **string** | Integration status | [optional] 

## Methods

### NewInlineResponse20123

`func NewInlineResponse20123() *InlineResponse20123`

NewInlineResponse20123 instantiates a new InlineResponse20123 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20123WithDefaults

`func NewInlineResponse20123WithDefaults() *InlineResponse20123`

NewInlineResponse20123WithDefaults instantiates a new InlineResponse20123 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrated

`func (o *InlineResponse20123) GetIntegrated() InlineResponse20123Integrated`

GetIntegrated returns the Integrated field if non-nil, zero value otherwise.

### GetIntegratedOk

`func (o *InlineResponse20123) GetIntegratedOk() (*InlineResponse20123Integrated, bool)`

GetIntegratedOk returns a tuple with the Integrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrated

`func (o *InlineResponse20123) SetIntegrated(v InlineResponse20123Integrated)`

SetIntegrated sets Integrated field to given value.

### HasIntegrated

`func (o *InlineResponse20123) HasIntegrated() bool`

HasIntegrated returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *InlineResponse20123) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *InlineResponse20123) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *InlineResponse20123) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *InlineResponse20123) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetExternal

`func (o *InlineResponse20123) GetExternal() InlineResponse20123External`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InlineResponse20123) GetExternalOk() (*InlineResponse20123External, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InlineResponse20123) SetExternal(v InlineResponse20123External)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InlineResponse20123) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20123) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20123) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20123) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20123) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


