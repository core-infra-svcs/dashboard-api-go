# InlineResponse200216Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200216Totals**](InlineResponse200216Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200216ByAlertType**](InlineResponse200216ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200216Items

`func NewInlineResponse200216Items(segmentStart time.Time, totals InlineResponse200216Totals, byAlertType []InlineResponse200216ByAlertType, ) *InlineResponse200216Items`

NewInlineResponse200216Items instantiates a new InlineResponse200216Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200216ItemsWithDefaults

`func NewInlineResponse200216ItemsWithDefaults() *InlineResponse200216Items`

NewInlineResponse200216ItemsWithDefaults instantiates a new InlineResponse200216Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200216Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200216Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200216Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200216Items) GetTotals() InlineResponse200216Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200216Items) GetTotalsOk() (*InlineResponse200216Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200216Items) SetTotals(v InlineResponse200216Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200216Items) GetByAlertType() []InlineResponse200216ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200216Items) GetByAlertTypeOk() (*[]InlineResponse200216ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200216Items) SetByAlertType(v []InlineResponse200216ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


