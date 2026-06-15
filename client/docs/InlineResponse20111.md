# InlineResponse20111

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LookupId** | Pointer to **string** | Id of the route table lookup request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your routing table lookup request. | [optional] 
**Request** | Pointer to [**InlineResponse20111Request**](InlineResponse20111Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the routing table lookup request. | [optional] 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse20111

`func NewInlineResponse20111() *InlineResponse20111`

NewInlineResponse20111 instantiates a new InlineResponse20111 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20111WithDefaults

`func NewInlineResponse20111WithDefaults() *InlineResponse20111`

NewInlineResponse20111WithDefaults instantiates a new InlineResponse20111 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLookupId

`func (o *InlineResponse20111) GetLookupId() string`

GetLookupId returns the LookupId field if non-nil, zero value otherwise.

### GetLookupIdOk

`func (o *InlineResponse20111) GetLookupIdOk() (*string, bool)`

GetLookupIdOk returns a tuple with the LookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupId

`func (o *InlineResponse20111) SetLookupId(v string)`

SetLookupId sets LookupId field to given value.

### HasLookupId

`func (o *InlineResponse20111) HasLookupId() bool`

HasLookupId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20111) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20111) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20111) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20111) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20111) GetRequest() InlineResponse20111Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20111) GetRequestOk() (*InlineResponse20111Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20111) SetRequest(v InlineResponse20111Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20111) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20111) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20111) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20111) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20111) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCallback

`func (o *InlineResponse20111) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse20111) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse20111) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse20111) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


