# InlineResponse2016

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingId** | Pointer to **string** | Id to check the status of your ping request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request. | [optional] 
**Request** | Pointer to [**InlineResponse2016Request**](InlineResponse2016Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ping request. | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse2016

`func NewInlineResponse2016() *InlineResponse2016`

NewInlineResponse2016 instantiates a new InlineResponse2016 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2016WithDefaults

`func NewInlineResponse2016WithDefaults() *InlineResponse2016`

NewInlineResponse2016WithDefaults instantiates a new InlineResponse2016 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPingId

`func (o *InlineResponse2016) GetPingId() string`

GetPingId returns the PingId field if non-nil, zero value otherwise.

### GetPingIdOk

`func (o *InlineResponse2016) GetPingIdOk() (*string, bool)`

GetPingIdOk returns a tuple with the PingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingId

`func (o *InlineResponse2016) SetPingId(v string)`

SetPingId sets PingId field to given value.

### HasPingId

`func (o *InlineResponse2016) HasPingId() bool`

HasPingId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2016) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2016) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2016) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2016) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2016) GetRequest() InlineResponse2016Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2016) GetRequestOk() (*InlineResponse2016Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2016) SetRequest(v InlineResponse2016Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2016) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2016) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2016) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2016) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2016) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse2016) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse2016) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse2016) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse2016) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


