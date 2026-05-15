# InlineResponse20033

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedsBlinkId** | Pointer to **string** | ID of led blink job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your leds blink request | [optional] 
**Status** | Pointer to **string** | Status of the leds blink request | [optional] 
**Request** | Pointer to [**InlineResponse2013Request**](InlineResponse2013Request.md) |  | [optional] 
**Error** | Pointer to **string** | An error message for a failed Blink LEDs execution, if present | [optional] 

## Methods

### NewInlineResponse20033

`func NewInlineResponse20033() *InlineResponse20033`

NewInlineResponse20033 instantiates a new InlineResponse20033 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20033WithDefaults

`func NewInlineResponse20033WithDefaults() *InlineResponse20033`

NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedsBlinkId

`func (o *InlineResponse20033) GetLedsBlinkId() string`

GetLedsBlinkId returns the LedsBlinkId field if non-nil, zero value otherwise.

### GetLedsBlinkIdOk

`func (o *InlineResponse20033) GetLedsBlinkIdOk() (*string, bool)`

GetLedsBlinkIdOk returns a tuple with the LedsBlinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedsBlinkId

`func (o *InlineResponse20033) SetLedsBlinkId(v string)`

SetLedsBlinkId sets LedsBlinkId field to given value.

### HasLedsBlinkId

`func (o *InlineResponse20033) HasLedsBlinkId() bool`

HasLedsBlinkId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20033) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20033) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20033) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20033) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20033) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20033) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20033) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20033) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20033) GetRequest() InlineResponse2013Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20033) GetRequestOk() (*InlineResponse2013Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20033) SetRequest(v InlineResponse2013Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20033) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20033) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20033) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20033) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20033) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


