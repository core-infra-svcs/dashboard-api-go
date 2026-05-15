# InlineResponse20038

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CyclePortId** | Pointer to **string** | ID of the cycle port request | [optional] 
**Url** | Pointer to **string** | GET this URL to check the status of your cycle port request. | [optional] 
**Request** | Pointer to [**InlineResponse2018Request**](InlineResponse2018Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the cycle port request | [optional] 
**Error** | Pointer to **string** | An error message for a failed cycle port execution, if present | [optional] 

## Methods

### NewInlineResponse20038

`func NewInlineResponse20038() *InlineResponse20038`

NewInlineResponse20038 instantiates a new InlineResponse20038 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20038WithDefaults

`func NewInlineResponse20038WithDefaults() *InlineResponse20038`

NewInlineResponse20038WithDefaults instantiates a new InlineResponse20038 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCyclePortId

`func (o *InlineResponse20038) GetCyclePortId() string`

GetCyclePortId returns the CyclePortId field if non-nil, zero value otherwise.

### GetCyclePortIdOk

`func (o *InlineResponse20038) GetCyclePortIdOk() (*string, bool)`

GetCyclePortIdOk returns a tuple with the CyclePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyclePortId

`func (o *InlineResponse20038) SetCyclePortId(v string)`

SetCyclePortId sets CyclePortId field to given value.

### HasCyclePortId

`func (o *InlineResponse20038) HasCyclePortId() bool`

HasCyclePortId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20038) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20038) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20038) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20038) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20038) GetRequest() InlineResponse2018Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20038) GetRequestOk() (*InlineResponse2018Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20038) SetRequest(v InlineResponse2018Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20038) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20038) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20038) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20038) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20038) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20038) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20038) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20038) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20038) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


