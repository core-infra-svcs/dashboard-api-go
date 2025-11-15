# InlineResponse200247Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200247Totals**](InlineResponse200247Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200247ByAlertType**](InlineResponse200247ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200247Items

`func NewInlineResponse200247Items(segmentStart time.Time, totals InlineResponse200247Totals, byAlertType []InlineResponse200247ByAlertType, ) *InlineResponse200247Items`

NewInlineResponse200247Items instantiates a new InlineResponse200247Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200247ItemsWithDefaults

`func NewInlineResponse200247ItemsWithDefaults() *InlineResponse200247Items`

NewInlineResponse200247ItemsWithDefaults instantiates a new InlineResponse200247Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200247Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200247Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200247Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200247Items) GetTotals() InlineResponse200247Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200247Items) GetTotalsOk() (*InlineResponse200247Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200247Items) SetTotals(v InlineResponse200247Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200247Items) GetByAlertType() []InlineResponse200247ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200247Items) GetByAlertTypeOk() (*[]InlineResponse200247ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200247Items) SetByAlertType(v []InlineResponse200247ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


