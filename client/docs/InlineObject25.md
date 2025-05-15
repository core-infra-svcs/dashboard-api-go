# InlineObject25

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Operation to run on the sensor. &#39;enableDownstreamPower&#39;, &#39;disableDownstreamPower&#39;, and &#39;cycleDownstreamPower&#39; turn power on/off to the device that is connected downstream of an MT40 power monitor. &#39;refreshData&#39; causes an MT15 or MT40 device to upload its latest readings so that they are immediately available in the Dashboard API. | 

## Methods

### NewInlineObject25

`func NewInlineObject25(operation string, ) *InlineObject25`

NewInlineObject25 instantiates a new InlineObject25 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject25WithDefaults

`func NewInlineObject25WithDefaults() *InlineObject25`

NewInlineObject25WithDefaults instantiates a new InlineObject25 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *InlineObject25) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineObject25) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineObject25) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


