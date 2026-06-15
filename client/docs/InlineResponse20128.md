# InlineResponse20128

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integrated** | Pointer to [**InlineResponse20128Integrated**](InlineResponse20128Integrated.md) |  | [optional] 
**LastUsedAt** | Pointer to **string** | ISO8601Z timestamp when integration was last used | [optional] 
**External** | Pointer to [**InlineResponse20128External**](InlineResponse20128External.md) |  | [optional] 
**Status** | Pointer to **string** | Integration status | [optional] 

## Methods

### NewInlineResponse20128

`func NewInlineResponse20128() *InlineResponse20128`

NewInlineResponse20128 instantiates a new InlineResponse20128 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20128WithDefaults

`func NewInlineResponse20128WithDefaults() *InlineResponse20128`

NewInlineResponse20128WithDefaults instantiates a new InlineResponse20128 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrated

`func (o *InlineResponse20128) GetIntegrated() InlineResponse20128Integrated`

GetIntegrated returns the Integrated field if non-nil, zero value otherwise.

### GetIntegratedOk

`func (o *InlineResponse20128) GetIntegratedOk() (*InlineResponse20128Integrated, bool)`

GetIntegratedOk returns a tuple with the Integrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrated

`func (o *InlineResponse20128) SetIntegrated(v InlineResponse20128Integrated)`

SetIntegrated sets Integrated field to given value.

### HasIntegrated

`func (o *InlineResponse20128) HasIntegrated() bool`

HasIntegrated returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *InlineResponse20128) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *InlineResponse20128) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *InlineResponse20128) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *InlineResponse20128) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetExternal

`func (o *InlineResponse20128) GetExternal() InlineResponse20128External`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InlineResponse20128) GetExternalOk() (*InlineResponse20128External, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InlineResponse20128) SetExternal(v InlineResponse20128External)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InlineResponse20128) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20128) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20128) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20128) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20128) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


