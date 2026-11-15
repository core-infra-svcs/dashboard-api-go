# InlineResponse20044

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LookupId** | Pointer to **string** | Id of the route table lookup request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your routing table lookup request. | [optional] 
**Request** | Pointer to [**InlineResponse20112Request**](InlineResponse20112Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the routing table lookup request. | [optional] 
**Entries** | Pointer to [**[]InlineResponse20044Entries**](InlineResponse20044Entries.md) | Routing table entries that match the search criteria | [optional] 
**Errors** | Pointer to **[]string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20044

`func NewInlineResponse20044() *InlineResponse20044`

NewInlineResponse20044 instantiates a new InlineResponse20044 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20044WithDefaults

`func NewInlineResponse20044WithDefaults() *InlineResponse20044`

NewInlineResponse20044WithDefaults instantiates a new InlineResponse20044 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLookupId

`func (o *InlineResponse20044) GetLookupId() string`

GetLookupId returns the LookupId field if non-nil, zero value otherwise.

### GetLookupIdOk

`func (o *InlineResponse20044) GetLookupIdOk() (*string, bool)`

GetLookupIdOk returns a tuple with the LookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupId

`func (o *InlineResponse20044) SetLookupId(v string)`

SetLookupId sets LookupId field to given value.

### HasLookupId

`func (o *InlineResponse20044) HasLookupId() bool`

HasLookupId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20044) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20044) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20044) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20044) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20044) GetRequest() InlineResponse20112Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20044) GetRequestOk() (*InlineResponse20112Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20044) SetRequest(v InlineResponse20112Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20044) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20044) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20044) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20044) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20044) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntries

`func (o *InlineResponse20044) GetEntries() []InlineResponse20044Entries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InlineResponse20044) GetEntriesOk() (*[]InlineResponse20044Entries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InlineResponse20044) SetEntries(v []InlineResponse20044Entries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InlineResponse20044) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20044) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20044) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20044) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20044) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


