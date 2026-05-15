# InlineResponse20031

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpTableId** | Pointer to **string** | Id of the ARP table request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ARP table request. | [optional] 
**Request** | Pointer to [**InlineResponse2011Request**](InlineResponse2011Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ARP table request. | [optional] 
**Entries** | Pointer to [**[]InlineResponse20031Entries**](InlineResponse20031Entries.md) | The ARP table entries | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20031

`func NewInlineResponse20031() *InlineResponse20031`

NewInlineResponse20031 instantiates a new InlineResponse20031 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20031WithDefaults

`func NewInlineResponse20031WithDefaults() *InlineResponse20031`

NewInlineResponse20031WithDefaults instantiates a new InlineResponse20031 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpTableId

`func (o *InlineResponse20031) GetArpTableId() string`

GetArpTableId returns the ArpTableId field if non-nil, zero value otherwise.

### GetArpTableIdOk

`func (o *InlineResponse20031) GetArpTableIdOk() (*string, bool)`

GetArpTableIdOk returns a tuple with the ArpTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpTableId

`func (o *InlineResponse20031) SetArpTableId(v string)`

SetArpTableId sets ArpTableId field to given value.

### HasArpTableId

`func (o *InlineResponse20031) HasArpTableId() bool`

HasArpTableId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20031) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20031) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20031) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20031) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20031) GetRequest() InlineResponse2011Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20031) GetRequestOk() (*InlineResponse2011Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20031) SetRequest(v InlineResponse2011Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20031) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20031) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20031) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20031) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20031) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntries

`func (o *InlineResponse20031) GetEntries() []InlineResponse20031Entries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InlineResponse20031) GetEntriesOk() (*[]InlineResponse20031Entries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InlineResponse20031) SetEntries(v []InlineResponse20031Entries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InlineResponse20031) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20031) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20031) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20031) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20031) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


