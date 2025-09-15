# InlineResponse200246Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200246Totals**](InlineResponse200246Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200246ByAlertType**](InlineResponse200246ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200246Items

`func NewInlineResponse200246Items(segmentStart time.Time, totals InlineResponse200246Totals, byAlertType []InlineResponse200246ByAlertType, ) *InlineResponse200246Items`

NewInlineResponse200246Items instantiates a new InlineResponse200246Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200246ItemsWithDefaults

`func NewInlineResponse200246ItemsWithDefaults() *InlineResponse200246Items`

NewInlineResponse200246ItemsWithDefaults instantiates a new InlineResponse200246Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200246Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200246Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200246Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200246Items) GetTotals() InlineResponse200246Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200246Items) GetTotalsOk() (*InlineResponse200246Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200246Items) SetTotals(v InlineResponse200246Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200246Items) GetByAlertType() []InlineResponse200246ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200246Items) GetByAlertTypeOk() (*[]InlineResponse200246ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200246Items) SetByAlertType(v []InlineResponse200246ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


