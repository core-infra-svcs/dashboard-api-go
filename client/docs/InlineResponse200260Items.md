# InlineResponse200260Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200260Totals**](InlineResponse200260Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200260ByAlertType**](InlineResponse200260ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200260Items

`func NewInlineResponse200260Items(segmentStart time.Time, totals InlineResponse200260Totals, byAlertType []InlineResponse200260ByAlertType, ) *InlineResponse200260Items`

NewInlineResponse200260Items instantiates a new InlineResponse200260Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200260ItemsWithDefaults

`func NewInlineResponse200260ItemsWithDefaults() *InlineResponse200260Items`

NewInlineResponse200260ItemsWithDefaults instantiates a new InlineResponse200260Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200260Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200260Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200260Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200260Items) GetTotals() InlineResponse200260Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200260Items) GetTotalsOk() (*InlineResponse200260Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200260Items) SetTotals(v InlineResponse200260Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200260Items) GetByAlertType() []InlineResponse200260ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200260Items) GetByAlertTypeOk() (*[]InlineResponse200260ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200260Items) SetByAlertType(v []InlineResponse200260ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


