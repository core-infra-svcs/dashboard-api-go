# InlineResponse200218

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the access period | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the access period | [optional] 
**Counts** | Pointer to [**[]OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts**](OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts.md) | list of response codes and a count of how many requests had that code in the given time period | [optional] 

## Methods

### NewInlineResponse200218

`func NewInlineResponse200218() *InlineResponse200218`

NewInlineResponse200218 instantiates a new InlineResponse200218 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200218WithDefaults

`func NewInlineResponse200218WithDefaults() *InlineResponse200218`

NewInlineResponse200218WithDefaults instantiates a new InlineResponse200218 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200218) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200218) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200218) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200218) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200218) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200218) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200218) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200218) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200218) GetCounts() []OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200218) GetCountsOk() (*[]OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200218) SetCounts(v []OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200218) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


