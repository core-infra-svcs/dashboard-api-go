# InlineResponse20040

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WakeOnLanId** | Pointer to **string** | ID of the Wake-on-LAN job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request | [optional] 
**Status** | Pointer to **string** | Status of the Wake-on-LAN request | [optional] 
**Request** | Pointer to [**InlineResponse20110Request**](InlineResponse20110Request.md) |  | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20040

`func NewInlineResponse20040() *InlineResponse20040`

NewInlineResponse20040 instantiates a new InlineResponse20040 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20040WithDefaults

`func NewInlineResponse20040WithDefaults() *InlineResponse20040`

NewInlineResponse20040WithDefaults instantiates a new InlineResponse20040 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWakeOnLanId

`func (o *InlineResponse20040) GetWakeOnLanId() string`

GetWakeOnLanId returns the WakeOnLanId field if non-nil, zero value otherwise.

### GetWakeOnLanIdOk

`func (o *InlineResponse20040) GetWakeOnLanIdOk() (*string, bool)`

GetWakeOnLanIdOk returns a tuple with the WakeOnLanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanId

`func (o *InlineResponse20040) SetWakeOnLanId(v string)`

SetWakeOnLanId sets WakeOnLanId field to given value.

### HasWakeOnLanId

`func (o *InlineResponse20040) HasWakeOnLanId() bool`

HasWakeOnLanId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20040) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20040) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20040) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20040) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20040) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20040) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20040) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20040) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20040) GetRequest() InlineResponse20110Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20040) GetRequestOk() (*InlineResponse20110Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20040) SetRequest(v InlineResponse20110Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20040) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20040) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20040) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20040) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20040) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


