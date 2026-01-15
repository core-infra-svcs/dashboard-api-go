# InlineResponse200252Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert Type | 
**LastResolvedAt** | Pointer to **time.Time** | Last time an alert of this type was resolved | [optional] 
**Count** | **int32** | Total count of the given alert type | 

## Methods

### NewInlineResponse200252Items

`func NewInlineResponse200252Items(type_ string, count int32, ) *InlineResponse200252Items`

NewInlineResponse200252Items instantiates a new InlineResponse200252Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200252ItemsWithDefaults

`func NewInlineResponse200252ItemsWithDefaults() *InlineResponse200252Items`

NewInlineResponse200252ItemsWithDefaults instantiates a new InlineResponse200252Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200252Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200252Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200252Items) SetType(v string)`

SetType sets Type field to given value.


### GetLastResolvedAt

`func (o *InlineResponse200252Items) GetLastResolvedAt() time.Time`

GetLastResolvedAt returns the LastResolvedAt field if non-nil, zero value otherwise.

### GetLastResolvedAtOk

`func (o *InlineResponse200252Items) GetLastResolvedAtOk() (*time.Time, bool)`

GetLastResolvedAtOk returns a tuple with the LastResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResolvedAt

`func (o *InlineResponse200252Items) SetLastResolvedAt(v time.Time)`

SetLastResolvedAt sets LastResolvedAt field to given value.

### HasLastResolvedAt

`func (o *InlineResponse200252Items) HasLastResolvedAt() bool`

HasLastResolvedAt returns a boolean if a field has been set.

### GetCount

`func (o *InlineResponse200252Items) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse200252Items) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse200252Items) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


