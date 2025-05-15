# InlineResponse2017

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThroughputTestId** | Pointer to **string** | ID of throughput test job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your throughput test request | [optional] 
**Status** | Pointer to **string** | Status of the throughput test request | [optional] 
**Result** | Pointer to [**InlineResponse2017Result**](InlineResponse2017Result.md) |  | [optional] 
**Request** | Pointer to [**InlineResponse2017Request**](InlineResponse2017Request.md) |  | [optional] 
**Error** | Pointer to **string** | Description of the error. | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse2017

`func NewInlineResponse2017() *InlineResponse2017`

NewInlineResponse2017 instantiates a new InlineResponse2017 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2017WithDefaults

`func NewInlineResponse2017WithDefaults() *InlineResponse2017`

NewInlineResponse2017WithDefaults instantiates a new InlineResponse2017 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThroughputTestId

`func (o *InlineResponse2017) GetThroughputTestId() string`

GetThroughputTestId returns the ThroughputTestId field if non-nil, zero value otherwise.

### GetThroughputTestIdOk

`func (o *InlineResponse2017) GetThroughputTestIdOk() (*string, bool)`

GetThroughputTestIdOk returns a tuple with the ThroughputTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputTestId

`func (o *InlineResponse2017) SetThroughputTestId(v string)`

SetThroughputTestId sets ThroughputTestId field to given value.

### HasThroughputTestId

`func (o *InlineResponse2017) HasThroughputTestId() bool`

HasThroughputTestId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2017) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2017) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2017) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2017) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2017) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2017) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2017) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2017) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse2017) GetResult() InlineResponse2017Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse2017) GetResultOk() (*InlineResponse2017Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse2017) SetResult(v InlineResponse2017Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse2017) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2017) GetRequest() InlineResponse2017Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2017) GetRequestOk() (*InlineResponse2017Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2017) SetRequest(v InlineResponse2017Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2017) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse2017) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse2017) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse2017) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse2017) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse2017) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse2017) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse2017) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse2017) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


