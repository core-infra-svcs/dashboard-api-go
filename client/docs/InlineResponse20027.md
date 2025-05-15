# InlineResponse20027

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingId** | Pointer to **string** | Id to check the status of your ping request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request. | [optional] 
**Request** | Pointer to [**InlineResponse2016Request**](InlineResponse2016Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ping request. | [optional] 
**Results** | Pointer to [**InlineResponse20026Results**](InlineResponse20026Results.md) |  | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse20027

`func NewInlineResponse20027() *InlineResponse20027`

NewInlineResponse20027 instantiates a new InlineResponse20027 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027WithDefaults

`func NewInlineResponse20027WithDefaults() *InlineResponse20027`

NewInlineResponse20027WithDefaults instantiates a new InlineResponse20027 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPingId

`func (o *InlineResponse20027) GetPingId() string`

GetPingId returns the PingId field if non-nil, zero value otherwise.

### GetPingIdOk

`func (o *InlineResponse20027) GetPingIdOk() (*string, bool)`

GetPingIdOk returns a tuple with the PingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingId

`func (o *InlineResponse20027) SetPingId(v string)`

SetPingId sets PingId field to given value.

### HasPingId

`func (o *InlineResponse20027) HasPingId() bool`

HasPingId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20027) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20027) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20027) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20027) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20027) GetRequest() InlineResponse2016Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20027) GetRequestOk() (*InlineResponse2016Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20027) SetRequest(v InlineResponse2016Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20027) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20027) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20027) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20027) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20027) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse20027) GetResults() InlineResponse20026Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse20027) GetResultsOk() (*InlineResponse20026Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse20027) SetResults(v InlineResponse20026Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse20027) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse20027) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse20027) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse20027) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse20027) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


