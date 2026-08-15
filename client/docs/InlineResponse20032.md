# InlineResponse20032

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpTableId** | Pointer to **string** | Id of the ARP table request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ARP table request. | [optional] 
**Request** | Pointer to [**InlineResponse2011Request**](InlineResponse2011Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ARP table request. | [optional] 
**Entries** | Pointer to [**[]InlineResponse20032Entries**](InlineResponse20032Entries.md) | The ARP table entries | [optional] 
**Error** | Pointer to **NullableString** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20032

`func NewInlineResponse20032() *InlineResponse20032`

NewInlineResponse20032 instantiates a new InlineResponse20032 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20032WithDefaults

`func NewInlineResponse20032WithDefaults() *InlineResponse20032`

NewInlineResponse20032WithDefaults instantiates a new InlineResponse20032 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpTableId

`func (o *InlineResponse20032) GetArpTableId() string`

GetArpTableId returns the ArpTableId field if non-nil, zero value otherwise.

### GetArpTableIdOk

`func (o *InlineResponse20032) GetArpTableIdOk() (*string, bool)`

GetArpTableIdOk returns a tuple with the ArpTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpTableId

`func (o *InlineResponse20032) SetArpTableId(v string)`

SetArpTableId sets ArpTableId field to given value.

### HasArpTableId

`func (o *InlineResponse20032) HasArpTableId() bool`

HasArpTableId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20032) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20032) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20032) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20032) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20032) GetRequest() InlineResponse2011Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20032) GetRequestOk() (*InlineResponse2011Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20032) SetRequest(v InlineResponse2011Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20032) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20032) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20032) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20032) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20032) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntries

`func (o *InlineResponse20032) GetEntries() []InlineResponse20032Entries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InlineResponse20032) GetEntriesOk() (*[]InlineResponse20032Entries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InlineResponse20032) SetEntries(v []InlineResponse20032Entries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InlineResponse20032) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20032) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20032) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20032) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20032) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *InlineResponse20032) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *InlineResponse20032) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


