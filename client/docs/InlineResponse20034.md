# InlineResponse20034

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacTableId** | Pointer to **string** | ID of the MAC table request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your MAC table request. | [optional] 
**Request** | Pointer to [**InlineResponse2014Request**](InlineResponse2014Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the MAC table request. | [optional] 
**Entries** | Pointer to [**[]InlineResponse20034Entries**](InlineResponse20034Entries.md) | MAC address table entries | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20034

`func NewInlineResponse20034() *InlineResponse20034`

NewInlineResponse20034 instantiates a new InlineResponse20034 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20034WithDefaults

`func NewInlineResponse20034WithDefaults() *InlineResponse20034`

NewInlineResponse20034WithDefaults instantiates a new InlineResponse20034 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacTableId

`func (o *InlineResponse20034) GetMacTableId() string`

GetMacTableId returns the MacTableId field if non-nil, zero value otherwise.

### GetMacTableIdOk

`func (o *InlineResponse20034) GetMacTableIdOk() (*string, bool)`

GetMacTableIdOk returns a tuple with the MacTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacTableId

`func (o *InlineResponse20034) SetMacTableId(v string)`

SetMacTableId sets MacTableId field to given value.

### HasMacTableId

`func (o *InlineResponse20034) HasMacTableId() bool`

HasMacTableId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20034) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20034) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20034) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20034) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20034) GetRequest() InlineResponse2014Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20034) GetRequestOk() (*InlineResponse2014Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20034) SetRequest(v InlineResponse2014Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20034) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20034) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20034) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20034) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20034) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntries

`func (o *InlineResponse20034) GetEntries() []InlineResponse20034Entries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InlineResponse20034) GetEntriesOk() (*[]InlineResponse20034Entries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InlineResponse20034) SetEntries(v []InlineResponse20034Entries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InlineResponse20034) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20034) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20034) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20034) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20034) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


