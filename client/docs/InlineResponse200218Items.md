# InlineResponse200218Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200218Totals**](InlineResponse200218Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200218ByAlertType**](InlineResponse200218ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200218Items

`func NewInlineResponse200218Items(segmentStart time.Time, totals InlineResponse200218Totals, byAlertType []InlineResponse200218ByAlertType, ) *InlineResponse200218Items`

NewInlineResponse200218Items instantiates a new InlineResponse200218Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200218ItemsWithDefaults

`func NewInlineResponse200218ItemsWithDefaults() *InlineResponse200218Items`

NewInlineResponse200218ItemsWithDefaults instantiates a new InlineResponse200218Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200218Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200218Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200218Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200218Items) GetTotals() InlineResponse200218Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200218Items) GetTotalsOk() (*InlineResponse200218Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200218Items) SetTotals(v InlineResponse200218Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200218Items) GetByAlertType() []InlineResponse200218ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200218Items) GetByAlertTypeOk() (*[]InlineResponse200218ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200218Items) SetByAlertType(v []InlineResponse200218ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


