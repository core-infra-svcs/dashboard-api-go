# InlineResponse2018

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WakeOnLanId** | Pointer to **string** | ID of the Wake-on-LAN job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request | [optional] 
**Status** | Pointer to **string** | Status of the Wake-on-LAN request | [optional] 
**Request** | Pointer to [**InlineResponse2018Request**](InlineResponse2018Request.md) |  | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse2018

`func NewInlineResponse2018() *InlineResponse2018`

NewInlineResponse2018 instantiates a new InlineResponse2018 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2018WithDefaults

`func NewInlineResponse2018WithDefaults() *InlineResponse2018`

NewInlineResponse2018WithDefaults instantiates a new InlineResponse2018 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWakeOnLanId

`func (o *InlineResponse2018) GetWakeOnLanId() string`

GetWakeOnLanId returns the WakeOnLanId field if non-nil, zero value otherwise.

### GetWakeOnLanIdOk

`func (o *InlineResponse2018) GetWakeOnLanIdOk() (*string, bool)`

GetWakeOnLanIdOk returns a tuple with the WakeOnLanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanId

`func (o *InlineResponse2018) SetWakeOnLanId(v string)`

SetWakeOnLanId sets WakeOnLanId field to given value.

### HasWakeOnLanId

`func (o *InlineResponse2018) HasWakeOnLanId() bool`

HasWakeOnLanId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2018) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2018) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2018) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2018) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2018) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2018) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2018) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2018) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2018) GetRequest() InlineResponse2018Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2018) GetRequestOk() (*InlineResponse2018Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2018) SetRequest(v InlineResponse2018Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2018) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse2018) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse2018) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse2018) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse2018) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse2018) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse2018) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse2018) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse2018) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


