# InlineResponse2019

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WakeOnLanId** | Pointer to **string** | ID of the Wake-on-LAN job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request | [optional] 
**Status** | Pointer to **string** | Status of the Wake-on-LAN request | [optional] 
**Request** | Pointer to [**InlineResponse2019Request**](InlineResponse2019Request.md) |  | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse2019

`func NewInlineResponse2019() *InlineResponse2019`

NewInlineResponse2019 instantiates a new InlineResponse2019 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2019WithDefaults

`func NewInlineResponse2019WithDefaults() *InlineResponse2019`

NewInlineResponse2019WithDefaults instantiates a new InlineResponse2019 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWakeOnLanId

`func (o *InlineResponse2019) GetWakeOnLanId() string`

GetWakeOnLanId returns the WakeOnLanId field if non-nil, zero value otherwise.

### GetWakeOnLanIdOk

`func (o *InlineResponse2019) GetWakeOnLanIdOk() (*string, bool)`

GetWakeOnLanIdOk returns a tuple with the WakeOnLanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanId

`func (o *InlineResponse2019) SetWakeOnLanId(v string)`

SetWakeOnLanId sets WakeOnLanId field to given value.

### HasWakeOnLanId

`func (o *InlineResponse2019) HasWakeOnLanId() bool`

HasWakeOnLanId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2019) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2019) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2019) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2019) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2019) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2019) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2019) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2019) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2019) GetRequest() InlineResponse2019Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2019) GetRequestOk() (*InlineResponse2019Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2019) SetRequest(v InlineResponse2019Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2019) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse2019) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse2019) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse2019) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse2019) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse2019) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse2019) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse2019) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse2019) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


