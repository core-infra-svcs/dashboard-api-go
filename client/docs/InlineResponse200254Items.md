# InlineResponse200254Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200254Totals**](InlineResponse200254Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200254ByAlertType**](InlineResponse200254ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200254Items

`func NewInlineResponse200254Items(segmentStart time.Time, totals InlineResponse200254Totals, byAlertType []InlineResponse200254ByAlertType, ) *InlineResponse200254Items`

NewInlineResponse200254Items instantiates a new InlineResponse200254Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200254ItemsWithDefaults

`func NewInlineResponse200254ItemsWithDefaults() *InlineResponse200254Items`

NewInlineResponse200254ItemsWithDefaults instantiates a new InlineResponse200254Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200254Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200254Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200254Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200254Items) GetTotals() InlineResponse200254Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200254Items) GetTotalsOk() (*InlineResponse200254Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200254Items) SetTotals(v InlineResponse200254Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200254Items) GetByAlertType() []InlineResponse200254ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200254Items) GetByAlertTypeOk() (*[]InlineResponse200254ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200254Items) SetByAlertType(v []InlineResponse200254ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


