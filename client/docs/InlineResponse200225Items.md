# InlineResponse200225Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200225Totals**](InlineResponse200225Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200225ByAlertType**](InlineResponse200225ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200225Items

`func NewInlineResponse200225Items(segmentStart time.Time, totals InlineResponse200225Totals, byAlertType []InlineResponse200225ByAlertType, ) *InlineResponse200225Items`

NewInlineResponse200225Items instantiates a new InlineResponse200225Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200225ItemsWithDefaults

`func NewInlineResponse200225ItemsWithDefaults() *InlineResponse200225Items`

NewInlineResponse200225ItemsWithDefaults instantiates a new InlineResponse200225Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200225Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200225Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200225Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200225Items) GetTotals() InlineResponse200225Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200225Items) GetTotalsOk() (*InlineResponse200225Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200225Items) SetTotals(v InlineResponse200225Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200225Items) GetByAlertType() []InlineResponse200225ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200225Items) GetByAlertTypeOk() (*[]InlineResponse200225ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200225Items) SetByAlertType(v []InlineResponse200225ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


