# InlineResponse200244Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200244Totals**](InlineResponse200244Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200244ByAlertType**](InlineResponse200244ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200244Items

`func NewInlineResponse200244Items(segmentStart time.Time, totals InlineResponse200244Totals, byAlertType []InlineResponse200244ByAlertType, ) *InlineResponse200244Items`

NewInlineResponse200244Items instantiates a new InlineResponse200244Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200244ItemsWithDefaults

`func NewInlineResponse200244ItemsWithDefaults() *InlineResponse200244Items`

NewInlineResponse200244ItemsWithDefaults instantiates a new InlineResponse200244Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200244Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200244Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200244Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200244Items) GetTotals() InlineResponse200244Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200244Items) GetTotalsOk() (*InlineResponse200244Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200244Items) SetTotals(v InlineResponse200244Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200244Items) GetByAlertType() []InlineResponse200244ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200244Items) GetByAlertTypeOk() (*[]InlineResponse200244ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200244Items) SetByAlertType(v []InlineResponse200244ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


