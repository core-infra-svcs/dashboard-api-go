# InlineResponse200224Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200224Totals**](InlineResponse200224Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200224ByAlertType**](InlineResponse200224ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200224Items

`func NewInlineResponse200224Items(segmentStart time.Time, totals InlineResponse200224Totals, byAlertType []InlineResponse200224ByAlertType, ) *InlineResponse200224Items`

NewInlineResponse200224Items instantiates a new InlineResponse200224Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200224ItemsWithDefaults

`func NewInlineResponse200224ItemsWithDefaults() *InlineResponse200224Items`

NewInlineResponse200224ItemsWithDefaults instantiates a new InlineResponse200224Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200224Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200224Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200224Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200224Items) GetTotals() InlineResponse200224Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200224Items) GetTotalsOk() (*InlineResponse200224Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200224Items) SetTotals(v InlineResponse200224Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200224Items) GetByAlertType() []InlineResponse200224ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200224Items) GetByAlertTypeOk() (*[]InlineResponse200224ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200224Items) SetByAlertType(v []InlineResponse200224ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


