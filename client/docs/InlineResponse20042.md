# InlineResponse20042

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LookupId** | Pointer to **string** | Id of the route table lookup request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your routing table lookup request. | [optional] 
**Request** | Pointer to [**InlineResponse20111Request**](InlineResponse20111Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the routing table lookup request. | [optional] 
**Entries** | Pointer to [**[]InlineResponse20042Entries**](InlineResponse20042Entries.md) | Routing table entries that match the search criteria | [optional] 
**Errors** | Pointer to **[]string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20042

`func NewInlineResponse20042() *InlineResponse20042`

NewInlineResponse20042 instantiates a new InlineResponse20042 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20042WithDefaults

`func NewInlineResponse20042WithDefaults() *InlineResponse20042`

NewInlineResponse20042WithDefaults instantiates a new InlineResponse20042 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLookupId

`func (o *InlineResponse20042) GetLookupId() string`

GetLookupId returns the LookupId field if non-nil, zero value otherwise.

### GetLookupIdOk

`func (o *InlineResponse20042) GetLookupIdOk() (*string, bool)`

GetLookupIdOk returns a tuple with the LookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupId

`func (o *InlineResponse20042) SetLookupId(v string)`

SetLookupId sets LookupId field to given value.

### HasLookupId

`func (o *InlineResponse20042) HasLookupId() bool`

HasLookupId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20042) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20042) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20042) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20042) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20042) GetRequest() InlineResponse20111Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20042) GetRequestOk() (*InlineResponse20111Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20042) SetRequest(v InlineResponse20111Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20042) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20042) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20042) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20042) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20042) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntries

`func (o *InlineResponse20042) GetEntries() []InlineResponse20042Entries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InlineResponse20042) GetEntriesOk() (*[]InlineResponse20042Entries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InlineResponse20042) SetEntries(v []InlineResponse20042Entries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InlineResponse20042) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20042) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20042) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20042) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20042) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


