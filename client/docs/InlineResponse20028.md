# InlineResponse20028

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingId** | Pointer to **string** | Id to check the status of your ping request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request. | [optional] 
**Request** | Pointer to [**InlineResponse2017Request**](InlineResponse2017Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ping request. | [optional] 
**Results** | Pointer to [**InlineResponse20027Results**](InlineResponse20027Results.md) |  | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

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

### GetPingId

`func (o *InlineResponse20028) GetPingId() string`

GetPingId returns the PingId field if non-nil, zero value otherwise.

### GetPingIdOk

`func (o *InlineResponse20028) GetPingIdOk() (*string, bool)`

GetPingIdOk returns a tuple with the PingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingId

`func (o *InlineResponse20028) SetPingId(v string)`

SetPingId sets PingId field to given value.

### HasPingId

`func (o *InlineResponse20028) HasPingId() bool`

HasPingId returns a boolean if a field has been set.

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

### GetResults

`func (o *InlineResponse20028) GetResults() InlineResponse20027Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse20028) GetResultsOk() (*InlineResponse20027Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse20028) SetResults(v InlineResponse20027Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse20028) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse20028) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse20028) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse20028) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse20028) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


