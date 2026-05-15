# InlineResponse200286Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200286Totals**](InlineResponse200286Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200286ByAlertType**](InlineResponse200286ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200286Items

`func NewInlineResponse200286Items(segmentStart time.Time, totals InlineResponse200286Totals, byAlertType []InlineResponse200286ByAlertType, ) *InlineResponse200286Items`

NewInlineResponse200286Items instantiates a new InlineResponse200286Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200286ItemsWithDefaults

`func NewInlineResponse200286ItemsWithDefaults() *InlineResponse200286Items`

NewInlineResponse200286ItemsWithDefaults instantiates a new InlineResponse200286Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200286Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200286Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200286Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200286Items) GetTotals() InlineResponse200286Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200286Items) GetTotalsOk() (*InlineResponse200286Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200286Items) SetTotals(v InlineResponse200286Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200286Items) GetByAlertType() []InlineResponse200286ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200286Items) GetByAlertTypeOk() (*[]InlineResponse200286ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200286Items) SetByAlertType(v []InlineResponse200286ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


