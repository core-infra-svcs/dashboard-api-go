# InlineResponse20029

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThroughputTestId** | Pointer to **string** | ID of throughput test job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your throughput test request | [optional] 
**Status** | Pointer to **string** | Status of the throughput test request | [optional] 
**Result** | Pointer to [**InlineResponse2018Result**](InlineResponse2018Result.md) |  | [optional] 
**Request** | Pointer to [**InlineResponse2018Request**](InlineResponse2018Request.md) |  | [optional] 
**Error** | Pointer to **string** | Description of the error. | [optional] 

## Methods

### NewInlineResponse20029

`func NewInlineResponse20029() *InlineResponse20029`

NewInlineResponse20029 instantiates a new InlineResponse20029 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029WithDefaults

`func NewInlineResponse20029WithDefaults() *InlineResponse20029`

NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThroughputTestId

`func (o *InlineResponse20029) GetThroughputTestId() string`

GetThroughputTestId returns the ThroughputTestId field if non-nil, zero value otherwise.

### GetThroughputTestIdOk

`func (o *InlineResponse20029) GetThroughputTestIdOk() (*string, bool)`

GetThroughputTestIdOk returns a tuple with the ThroughputTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputTestId

`func (o *InlineResponse20029) SetThroughputTestId(v string)`

SetThroughputTestId sets ThroughputTestId field to given value.

### HasThroughputTestId

`func (o *InlineResponse20029) HasThroughputTestId() bool`

HasThroughputTestId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20029) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20029) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20029) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20029) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20029) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20029) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20029) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20029) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse20029) GetResult() InlineResponse2018Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse20029) GetResultOk() (*InlineResponse2018Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse20029) SetResult(v InlineResponse2018Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse20029) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20029) GetRequest() InlineResponse2018Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20029) GetRequestOk() (*InlineResponse2018Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20029) SetRequest(v InlineResponse2018Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20029) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20029) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20029) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20029) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20029) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


