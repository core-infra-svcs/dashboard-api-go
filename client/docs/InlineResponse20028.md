# InlineResponse20028

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThroughputTestId** | Pointer to **string** | ID of throughput test job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your throughput test request | [optional] 
**Status** | Pointer to **string** | Status of the throughput test request | [optional] 
**Result** | Pointer to [**InlineResponse2017Result**](InlineResponse2017Result.md) |  | [optional] 
**Request** | Pointer to [**InlineResponse2017Request**](InlineResponse2017Request.md) |  | [optional] 
**Error** | Pointer to **string** | Description of the error. | [optional] 

## Methods

### NewInlineResponse20028

`func NewInlineResponse20028() *InlineResponse20028`

NewInlineResponse20028 instantiates a new InlineResponse20028 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20028WithDefaults

`func NewInlineResponse20028WithDefaults() *InlineResponse20028`

NewInlineResponse20028WithDefaults instantiates a new InlineResponse20028 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThroughputTestId

`func (o *InlineResponse20028) GetThroughputTestId() string`

GetThroughputTestId returns the ThroughputTestId field if non-nil, zero value otherwise.

### GetThroughputTestIdOk

`func (o *InlineResponse20028) GetThroughputTestIdOk() (*string, bool)`

GetThroughputTestIdOk returns a tuple with the ThroughputTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputTestId

`func (o *InlineResponse20028) SetThroughputTestId(v string)`

SetThroughputTestId sets ThroughputTestId field to given value.

### HasThroughputTestId

`func (o *InlineResponse20028) HasThroughputTestId() bool`

HasThroughputTestId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20028) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20028) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20028) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20028) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20028) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20028) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20028) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20028) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse20028) GetResult() InlineResponse2017Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse20028) GetResultOk() (*InlineResponse2017Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse20028) SetResult(v InlineResponse2017Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse20028) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20028) GetRequest() InlineResponse2017Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20028) GetRequestOk() (*InlineResponse2017Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20028) SetRequest(v InlineResponse2017Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20028) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20028) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20028) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20028) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20028) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


