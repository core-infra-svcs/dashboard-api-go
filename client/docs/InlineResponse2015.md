# InlineResponse2015

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MulticastRoutingId** | Pointer to **string** | ID of the Multicast routing request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this URL to check the status of your Multicast routing request. | [optional] 
**Request** | Pointer to [**InlineResponse2015Request**](InlineResponse2015Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the Multicast routing request. | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse2015

`func NewInlineResponse2015() *InlineResponse2015`

NewInlineResponse2015 instantiates a new InlineResponse2015 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2015WithDefaults

`func NewInlineResponse2015WithDefaults() *InlineResponse2015`

NewInlineResponse2015WithDefaults instantiates a new InlineResponse2015 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMulticastRoutingId

`func (o *InlineResponse2015) GetMulticastRoutingId() string`

GetMulticastRoutingId returns the MulticastRoutingId field if non-nil, zero value otherwise.

### GetMulticastRoutingIdOk

`func (o *InlineResponse2015) GetMulticastRoutingIdOk() (*string, bool)`

GetMulticastRoutingIdOk returns a tuple with the MulticastRoutingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRoutingId

`func (o *InlineResponse2015) SetMulticastRoutingId(v string)`

SetMulticastRoutingId sets MulticastRoutingId field to given value.

### HasMulticastRoutingId

`func (o *InlineResponse2015) HasMulticastRoutingId() bool`

HasMulticastRoutingId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2015) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2015) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2015) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2015) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2015) GetRequest() InlineResponse2015Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2015) GetRequestOk() (*InlineResponse2015Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2015) SetRequest(v InlineResponse2015Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2015) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2015) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2015) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2015) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2015) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse2015) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse2015) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse2015) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse2015) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


