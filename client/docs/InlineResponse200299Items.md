# InlineResponse200299Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200299Totals**](InlineResponse200299Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200299ByAlertType**](InlineResponse200299ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200299Items

`func NewInlineResponse200299Items(segmentStart time.Time, totals InlineResponse200299Totals, byAlertType []InlineResponse200299ByAlertType, ) *InlineResponse200299Items`

NewInlineResponse200299Items instantiates a new InlineResponse200299Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200299ItemsWithDefaults

`func NewInlineResponse200299ItemsWithDefaults() *InlineResponse200299Items`

NewInlineResponse200299ItemsWithDefaults instantiates a new InlineResponse200299Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200299Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200299Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200299Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200299Items) GetTotals() InlineResponse200299Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200299Items) GetTotalsOk() (*InlineResponse200299Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200299Items) SetTotals(v InlineResponse200299Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200299Items) GetByAlertType() []InlineResponse200299ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200299Items) GetByAlertTypeOk() (*[]InlineResponse200299ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200299Items) SetByAlertType(v []InlineResponse200299ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


