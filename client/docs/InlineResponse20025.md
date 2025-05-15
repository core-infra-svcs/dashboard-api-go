# InlineResponse20025

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacTableId** | Pointer to **string** | ID of the MAC table request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your MAC table request. | [optional] 
**Request** | Pointer to [**InlineResponse2014Request**](InlineResponse2014Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the MAC table request. | [optional] 
**Entries** | Pointer to [**[]InlineResponse20025Entries**](InlineResponse20025Entries.md) | MAC address table entries | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20025

`func NewInlineResponse20025() *InlineResponse20025`

NewInlineResponse20025 instantiates a new InlineResponse20025 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20025WithDefaults

`func NewInlineResponse20025WithDefaults() *InlineResponse20025`

NewInlineResponse20025WithDefaults instantiates a new InlineResponse20025 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacTableId

`func (o *InlineResponse20025) GetMacTableId() string`

GetMacTableId returns the MacTableId field if non-nil, zero value otherwise.

### GetMacTableIdOk

`func (o *InlineResponse20025) GetMacTableIdOk() (*string, bool)`

GetMacTableIdOk returns a tuple with the MacTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacTableId

`func (o *InlineResponse20025) SetMacTableId(v string)`

SetMacTableId sets MacTableId field to given value.

### HasMacTableId

`func (o *InlineResponse20025) HasMacTableId() bool`

HasMacTableId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20025) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20025) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20025) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20025) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20025) GetRequest() InlineResponse2014Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20025) GetRequestOk() (*InlineResponse2014Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20025) SetRequest(v InlineResponse2014Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20025) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20025) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20025) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20025) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20025) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntries

`func (o *InlineResponse20025) GetEntries() []InlineResponse20025Entries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InlineResponse20025) GetEntriesOk() (*[]InlineResponse20025Entries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InlineResponse20025) SetEntries(v []InlineResponse20025Entries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InlineResponse20025) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20025) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20025) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20025) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20025) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


