# InlineResponse20024

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedsBlinkId** | Pointer to **string** | ID of led blink job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your leds blink request | [optional] 
**Status** | Pointer to **string** | Status of the leds blink request | [optional] 
**Request** | Pointer to [**InlineResponse2013Request**](InlineResponse2013Request.md) |  | [optional] 
**Error** | Pointer to **string** | An error message for a failed Blink LEDs execution, if present | [optional] 

## Methods

### NewInlineResponse20024

`func NewInlineResponse20024() *InlineResponse20024`

NewInlineResponse20024 instantiates a new InlineResponse20024 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20024WithDefaults

`func NewInlineResponse20024WithDefaults() *InlineResponse20024`

NewInlineResponse20024WithDefaults instantiates a new InlineResponse20024 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedsBlinkId

`func (o *InlineResponse20024) GetLedsBlinkId() string`

GetLedsBlinkId returns the LedsBlinkId field if non-nil, zero value otherwise.

### GetLedsBlinkIdOk

`func (o *InlineResponse20024) GetLedsBlinkIdOk() (*string, bool)`

GetLedsBlinkIdOk returns a tuple with the LedsBlinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedsBlinkId

`func (o *InlineResponse20024) SetLedsBlinkId(v string)`

SetLedsBlinkId sets LedsBlinkId field to given value.

### HasLedsBlinkId

`func (o *InlineResponse20024) HasLedsBlinkId() bool`

HasLedsBlinkId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20024) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20024) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20024) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20024) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20024) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20024) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20024) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20024) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20024) GetRequest() InlineResponse2013Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20024) GetRequestOk() (*InlineResponse2013Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20024) SetRequest(v InlineResponse2013Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20024) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20024) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20024) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20024) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20024) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


