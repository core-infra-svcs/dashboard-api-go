# InlineResponse2011Callback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id} | [optional] 
**Url** | Pointer to **string** | The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver | [optional] 
**Status** | Pointer to **string** | The status of the callback | [optional] 

## Methods

### NewInlineResponse2011Callback

`func NewInlineResponse2011Callback() *InlineResponse2011Callback`

NewInlineResponse2011Callback instantiates a new InlineResponse2011Callback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2011CallbackWithDefaults

`func NewInlineResponse2011CallbackWithDefaults() *InlineResponse2011Callback`

NewInlineResponse2011CallbackWithDefaults instantiates a new InlineResponse2011Callback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2011Callback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2011Callback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2011Callback) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2011Callback) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2011Callback) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2011Callback) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2011Callback) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2011Callback) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2011Callback) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2011Callback) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2011Callback) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2011Callback) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


