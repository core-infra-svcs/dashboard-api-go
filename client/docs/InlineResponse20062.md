# InlineResponse20062

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | Start of the timespan over which sensor alerts are counted | [optional] 
**EndTs** | Pointer to **time.Time** | End of the timespan over which sensor alerts are counted | [optional] 
**Counts** | Pointer to [**NetworksNetworkIdSensorAlertsOverviewByMetricCounts**](NetworksNetworkIdSensorAlertsOverviewByMetricCounts.md) |  | [optional] 

## Methods

### NewInlineResponse20062

`func NewInlineResponse20062() *InlineResponse20062`

NewInlineResponse20062 instantiates a new InlineResponse20062 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20062WithDefaults

`func NewInlineResponse20062WithDefaults() *InlineResponse20062`

NewInlineResponse20062WithDefaults instantiates a new InlineResponse20062 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse20062) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse20062) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse20062) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse20062) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse20062) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse20062) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse20062) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse20062) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse20062) GetCounts() NetworksNetworkIdSensorAlertsOverviewByMetricCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse20062) GetCountsOk() (*NetworksNetworkIdSensorAlertsOverviewByMetricCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse20062) SetCounts(v NetworksNetworkIdSensorAlertsOverviewByMetricCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse20062) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


