# InlineResponse200246Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert Type | 
**LastResolvedAt** | Pointer to **time.Time** | Last time an alert of this type was resolved | [optional] 
**Count** | **int32** | Total count of the given alert type | 

## Methods

### NewInlineResponse200246Items

`func NewInlineResponse200246Items(type_ string, count int32, ) *InlineResponse200246Items`

NewInlineResponse200246Items instantiates a new InlineResponse200246Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200246ItemsWithDefaults

`func NewInlineResponse200246ItemsWithDefaults() *InlineResponse200246Items`

NewInlineResponse200246ItemsWithDefaults instantiates a new InlineResponse200246Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200246Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200246Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200246Items) SetType(v string)`

SetType sets Type field to given value.


### GetLastResolvedAt

`func (o *InlineResponse200246Items) GetLastResolvedAt() time.Time`

GetLastResolvedAt returns the LastResolvedAt field if non-nil, zero value otherwise.

### GetLastResolvedAtOk

`func (o *InlineResponse200246Items) GetLastResolvedAtOk() (*time.Time, bool)`

GetLastResolvedAtOk returns a tuple with the LastResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResolvedAt

`func (o *InlineResponse200246Items) SetLastResolvedAt(v time.Time)`

SetLastResolvedAt sets LastResolvedAt field to given value.

### HasLastResolvedAt

`func (o *InlineResponse200246Items) HasLastResolvedAt() bool`

HasLastResolvedAt returns a boolean if a field has been set.

### GetCount

`func (o *InlineResponse200246Items) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse200246Items) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse200246Items) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


