# InlineResponse200253Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200253Totals**](InlineResponse200253Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200253ByAlertType**](InlineResponse200253ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200253Items

`func NewInlineResponse200253Items(segmentStart time.Time, totals InlineResponse200253Totals, byAlertType []InlineResponse200253ByAlertType, ) *InlineResponse200253Items`

NewInlineResponse200253Items instantiates a new InlineResponse200253Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200253ItemsWithDefaults

`func NewInlineResponse200253ItemsWithDefaults() *InlineResponse200253Items`

NewInlineResponse200253ItemsWithDefaults instantiates a new InlineResponse200253Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200253Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200253Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200253Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200253Items) GetTotals() InlineResponse200253Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200253Items) GetTotalsOk() (*InlineResponse200253Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200253Items) SetTotals(v InlineResponse200253Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200253Items) GetByAlertType() []InlineResponse200253ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200253Items) GetByAlertTypeOk() (*[]InlineResponse200253ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200253Items) SetByAlertType(v []InlineResponse200253ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


