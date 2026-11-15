# InlineResponse200322Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200322Totals**](InlineResponse200322Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200322ByAlertType**](InlineResponse200322ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200322Items

`func NewInlineResponse200322Items(segmentStart time.Time, totals InlineResponse200322Totals, byAlertType []InlineResponse200322ByAlertType, ) *InlineResponse200322Items`

NewInlineResponse200322Items instantiates a new InlineResponse200322Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200322ItemsWithDefaults

`func NewInlineResponse200322ItemsWithDefaults() *InlineResponse200322Items`

NewInlineResponse200322ItemsWithDefaults instantiates a new InlineResponse200322Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200322Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200322Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200322Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200322Items) GetTotals() InlineResponse200322Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200322Items) GetTotalsOk() (*InlineResponse200322Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200322Items) SetTotals(v InlineResponse200322Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200322Items) GetByAlertType() []InlineResponse200322ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200322Items) GetByAlertTypeOk() (*[]InlineResponse200322ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200322Items) SetByAlertType(v []InlineResponse200322ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


