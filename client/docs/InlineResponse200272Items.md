# InlineResponse200272Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200272Totals**](InlineResponse200272Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200272ByAlertType**](InlineResponse200272ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200272Items

`func NewInlineResponse200272Items(segmentStart time.Time, totals InlineResponse200272Totals, byAlertType []InlineResponse200272ByAlertType, ) *InlineResponse200272Items`

NewInlineResponse200272Items instantiates a new InlineResponse200272Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200272ItemsWithDefaults

`func NewInlineResponse200272ItemsWithDefaults() *InlineResponse200272Items`

NewInlineResponse200272ItemsWithDefaults instantiates a new InlineResponse200272Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200272Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200272Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200272Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200272Items) GetTotals() InlineResponse200272Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200272Items) GetTotalsOk() (*InlineResponse200272Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200272Items) SetTotals(v InlineResponse200272Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200272Items) GetByAlertType() []InlineResponse200272ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200272Items) GetByAlertTypeOk() (*[]InlineResponse200272ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200272Items) SetByAlertType(v []InlineResponse200272ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


