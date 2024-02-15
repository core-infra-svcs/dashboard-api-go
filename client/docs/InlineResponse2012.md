# InlineResponse2012

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CableTestId** | Pointer to **string** | Id of the cable test request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your cable test request. | [optional] 
**Request** | Pointer to [**InlineResponse2012Request**](InlineResponse2012Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the cable test request. | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse2012

`func NewInlineResponse2012() *InlineResponse2012`

NewInlineResponse2012 instantiates a new InlineResponse2012 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2012WithDefaults

`func NewInlineResponse2012WithDefaults() *InlineResponse2012`

NewInlineResponse2012WithDefaults instantiates a new InlineResponse2012 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCableTestId

`func (o *InlineResponse2012) GetCableTestId() string`

GetCableTestId returns the CableTestId field if non-nil, zero value otherwise.

### GetCableTestIdOk

`func (o *InlineResponse2012) GetCableTestIdOk() (*string, bool)`

GetCableTestIdOk returns a tuple with the CableTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableTestId

`func (o *InlineResponse2012) SetCableTestId(v string)`

SetCableTestId sets CableTestId field to given value.

### HasCableTestId

`func (o *InlineResponse2012) HasCableTestId() bool`

HasCableTestId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2012) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2012) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2012) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2012) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2012) GetRequest() InlineResponse2012Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2012) GetRequestOk() (*InlineResponse2012Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2012) SetRequest(v InlineResponse2012Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2012) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2012) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2012) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2012) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2012) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse2012) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse2012) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse2012) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse2012) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


