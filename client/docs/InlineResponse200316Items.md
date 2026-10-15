# InlineResponse200316Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200316Totals**](InlineResponse200316Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200316ByAlertType**](InlineResponse200316ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200316Items

`func NewInlineResponse200316Items(segmentStart time.Time, totals InlineResponse200316Totals, byAlertType []InlineResponse200316ByAlertType, ) *InlineResponse200316Items`

NewInlineResponse200316Items instantiates a new InlineResponse200316Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200316ItemsWithDefaults

`func NewInlineResponse200316ItemsWithDefaults() *InlineResponse200316Items`

NewInlineResponse200316ItemsWithDefaults instantiates a new InlineResponse200316Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200316Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200316Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200316Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200316Items) GetTotals() InlineResponse200316Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200316Items) GetTotalsOk() (*InlineResponse200316Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200316Items) SetTotals(v InlineResponse200316Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200316Items) GetByAlertType() []InlineResponse200316ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200316Items) GetByAlertTypeOk() (*[]InlineResponse200316ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200316Items) SetByAlertType(v []InlineResponse200316ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


