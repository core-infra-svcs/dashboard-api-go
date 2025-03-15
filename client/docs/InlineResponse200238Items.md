# InlineResponse200238Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200238Totals**](InlineResponse200238Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200238ByAlertType**](InlineResponse200238ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200238Items

`func NewInlineResponse200238Items(segmentStart time.Time, totals InlineResponse200238Totals, byAlertType []InlineResponse200238ByAlertType, ) *InlineResponse200238Items`

NewInlineResponse200238Items instantiates a new InlineResponse200238Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200238ItemsWithDefaults

`func NewInlineResponse200238ItemsWithDefaults() *InlineResponse200238Items`

NewInlineResponse200238ItemsWithDefaults instantiates a new InlineResponse200238Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200238Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200238Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200238Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200238Items) GetTotals() InlineResponse200238Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200238Items) GetTotalsOk() (*InlineResponse200238Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200238Items) SetTotals(v InlineResponse200238Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200238Items) GetByAlertType() []InlineResponse200238ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200238Items) GetByAlertTypeOk() (*[]InlineResponse200238ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200238Items) SetByAlertType(v []InlineResponse200238ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


