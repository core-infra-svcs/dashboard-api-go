# InlineResponse200240Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200240Totals**](InlineResponse200240Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200240ByAlertType**](InlineResponse200240ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200240Items

`func NewInlineResponse200240Items(segmentStart time.Time, totals InlineResponse200240Totals, byAlertType []InlineResponse200240ByAlertType, ) *InlineResponse200240Items`

NewInlineResponse200240Items instantiates a new InlineResponse200240Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200240ItemsWithDefaults

`func NewInlineResponse200240ItemsWithDefaults() *InlineResponse200240Items`

NewInlineResponse200240ItemsWithDefaults instantiates a new InlineResponse200240Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200240Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200240Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200240Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200240Items) GetTotals() InlineResponse200240Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200240Items) GetTotalsOk() (*InlineResponse200240Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200240Items) SetTotals(v InlineResponse200240Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200240Items) GetByAlertType() []InlineResponse200240ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200240Items) GetByAlertTypeOk() (*[]InlineResponse200240ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200240Items) SetByAlertType(v []InlineResponse200240ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


