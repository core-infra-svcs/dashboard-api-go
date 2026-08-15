# InlineResponse200301Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200301Totals**](InlineResponse200301Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200301ByAlertType**](InlineResponse200301ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200301Items

`func NewInlineResponse200301Items(segmentStart time.Time, totals InlineResponse200301Totals, byAlertType []InlineResponse200301ByAlertType, ) *InlineResponse200301Items`

NewInlineResponse200301Items instantiates a new InlineResponse200301Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200301ItemsWithDefaults

`func NewInlineResponse200301ItemsWithDefaults() *InlineResponse200301Items`

NewInlineResponse200301ItemsWithDefaults instantiates a new InlineResponse200301Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200301Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200301Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200301Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200301Items) GetTotals() InlineResponse200301Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200301Items) GetTotalsOk() (*InlineResponse200301Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200301Items) SetTotals(v InlineResponse200301Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200301Items) GetByAlertType() []InlineResponse200301ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200301Items) GetByAlertTypeOk() (*[]InlineResponse200301ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200301Items) SetByAlertType(v []InlineResponse200301ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


