# InlineResponse20026

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WakeOnLanId** | Pointer to **string** | ID of the Wake-on-LAN job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request | [optional] 
**Status** | Pointer to **string** | Status of the Wake-on-LAN request | [optional] 
**Request** | Pointer to [**InlineResponse2017Request**](InlineResponse2017Request.md) |  | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20026

`func NewInlineResponse20026() *InlineResponse20026`

NewInlineResponse20026 instantiates a new InlineResponse20026 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20026WithDefaults

`func NewInlineResponse20026WithDefaults() *InlineResponse20026`

NewInlineResponse20026WithDefaults instantiates a new InlineResponse20026 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWakeOnLanId

`func (o *InlineResponse20026) GetWakeOnLanId() string`

GetWakeOnLanId returns the WakeOnLanId field if non-nil, zero value otherwise.

### GetWakeOnLanIdOk

`func (o *InlineResponse20026) GetWakeOnLanIdOk() (*string, bool)`

GetWakeOnLanIdOk returns a tuple with the WakeOnLanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanId

`func (o *InlineResponse20026) SetWakeOnLanId(v string)`

SetWakeOnLanId sets WakeOnLanId field to given value.

### HasWakeOnLanId

`func (o *InlineResponse20026) HasWakeOnLanId() bool`

HasWakeOnLanId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20026) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20026) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20026) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20026) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20026) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20026) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20026) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20026) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20026) GetRequest() InlineResponse2017Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20026) GetRequestOk() (*InlineResponse2017Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20026) SetRequest(v InlineResponse2017Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20026) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20026) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20026) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20026) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20026) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


