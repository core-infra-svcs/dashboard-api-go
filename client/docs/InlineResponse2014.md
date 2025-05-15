# InlineResponse2014

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacTableId** | Pointer to **string** | ID of the MAC table request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your MAC table request. | [optional] 
**Request** | Pointer to [**InlineResponse2014Request**](InlineResponse2014Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the MAC table request. | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse2014

`func NewInlineResponse2014() *InlineResponse2014`

NewInlineResponse2014 instantiates a new InlineResponse2014 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2014WithDefaults

`func NewInlineResponse2014WithDefaults() *InlineResponse2014`

NewInlineResponse2014WithDefaults instantiates a new InlineResponse2014 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacTableId

`func (o *InlineResponse2014) GetMacTableId() string`

GetMacTableId returns the MacTableId field if non-nil, zero value otherwise.

### GetMacTableIdOk

`func (o *InlineResponse2014) GetMacTableIdOk() (*string, bool)`

GetMacTableIdOk returns a tuple with the MacTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacTableId

`func (o *InlineResponse2014) SetMacTableId(v string)`

SetMacTableId sets MacTableId field to given value.

### HasMacTableId

`func (o *InlineResponse2014) HasMacTableId() bool`

HasMacTableId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2014) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2014) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2014) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2014) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2014) GetRequest() InlineResponse2014Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2014) GetRequestOk() (*InlineResponse2014Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2014) SetRequest(v InlineResponse2014Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2014) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2014) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2014) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2014) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2014) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse2014) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse2014) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse2014) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse2014) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


