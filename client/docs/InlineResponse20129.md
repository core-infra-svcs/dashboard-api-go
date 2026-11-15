# InlineResponse20129

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integrated** | Pointer to [**InlineResponse20129Integrated**](InlineResponse20129Integrated.md) |  | [optional] 
**LastUsedAt** | Pointer to **string** | ISO8601Z timestamp when integration was last used | [optional] 
**External** | Pointer to [**InlineResponse20129External**](InlineResponse20129External.md) |  | [optional] 
**Status** | Pointer to **string** | Integration status | [optional] 

## Methods

### NewInlineResponse20129

`func NewInlineResponse20129() *InlineResponse20129`

NewInlineResponse20129 instantiates a new InlineResponse20129 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20129WithDefaults

`func NewInlineResponse20129WithDefaults() *InlineResponse20129`

NewInlineResponse20129WithDefaults instantiates a new InlineResponse20129 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrated

`func (o *InlineResponse20129) GetIntegrated() InlineResponse20129Integrated`

GetIntegrated returns the Integrated field if non-nil, zero value otherwise.

### GetIntegratedOk

`func (o *InlineResponse20129) GetIntegratedOk() (*InlineResponse20129Integrated, bool)`

GetIntegratedOk returns a tuple with the Integrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrated

`func (o *InlineResponse20129) SetIntegrated(v InlineResponse20129Integrated)`

SetIntegrated sets Integrated field to given value.

### HasIntegrated

`func (o *InlineResponse20129) HasIntegrated() bool`

HasIntegrated returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *InlineResponse20129) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *InlineResponse20129) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *InlineResponse20129) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *InlineResponse20129) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetExternal

`func (o *InlineResponse20129) GetExternal() InlineResponse20129External`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InlineResponse20129) GetExternalOk() (*InlineResponse20129External, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InlineResponse20129) SetExternal(v InlineResponse20129External)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InlineResponse20129) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20129) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20129) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20129) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20129) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


