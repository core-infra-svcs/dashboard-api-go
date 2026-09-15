# InlineResponse200303Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200303Totals**](InlineResponse200303Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200303ByAlertType**](InlineResponse200303ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200303Items

`func NewInlineResponse200303Items(segmentStart time.Time, totals InlineResponse200303Totals, byAlertType []InlineResponse200303ByAlertType, ) *InlineResponse200303Items`

NewInlineResponse200303Items instantiates a new InlineResponse200303Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200303ItemsWithDefaults

`func NewInlineResponse200303ItemsWithDefaults() *InlineResponse200303Items`

NewInlineResponse200303ItemsWithDefaults instantiates a new InlineResponse200303Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200303Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200303Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200303Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200303Items) GetTotals() InlineResponse200303Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200303Items) GetTotalsOk() (*InlineResponse200303Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200303Items) SetTotals(v InlineResponse200303Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200303Items) GetByAlertType() []InlineResponse200303ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200303Items) GetByAlertTypeOk() (*[]InlineResponse200303ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200303Items) SetByAlertType(v []InlineResponse200303ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


