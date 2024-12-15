# InlineResponse2013

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedsBlinkId** | Pointer to **string** | ID of led blink job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your leds blink request | [optional] 
**Status** | Pointer to **string** | Status of the leds blink request | [optional] 
**Request** | Pointer to [**InlineResponse2013Request**](InlineResponse2013Request.md) |  | [optional] 
**Error** | Pointer to **string** | An error message for a failed Blink LEDs execution, if present | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse2013

`func NewInlineResponse2013() *InlineResponse2013`

NewInlineResponse2013 instantiates a new InlineResponse2013 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2013WithDefaults

`func NewInlineResponse2013WithDefaults() *InlineResponse2013`

NewInlineResponse2013WithDefaults instantiates a new InlineResponse2013 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedsBlinkId

`func (o *InlineResponse2013) GetLedsBlinkId() string`

GetLedsBlinkId returns the LedsBlinkId field if non-nil, zero value otherwise.

### GetLedsBlinkIdOk

`func (o *InlineResponse2013) GetLedsBlinkIdOk() (*string, bool)`

GetLedsBlinkIdOk returns a tuple with the LedsBlinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedsBlinkId

`func (o *InlineResponse2013) SetLedsBlinkId(v string)`

SetLedsBlinkId sets LedsBlinkId field to given value.

### HasLedsBlinkId

`func (o *InlineResponse2013) HasLedsBlinkId() bool`

HasLedsBlinkId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2013) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2013) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2013) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2013) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2013) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2013) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2013) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2013) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2013) GetRequest() InlineResponse2013Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2013) GetRequestOk() (*InlineResponse2013Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2013) SetRequest(v InlineResponse2013Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2013) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse2013) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse2013) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse2013) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse2013) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse2013) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse2013) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse2013) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse2013) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


