# InlineResponse200315Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID for integration | [optional] 
**Type** | Pointer to **string** | Identifier for integration type. One of Axis, Catalyst SD-WAN, Cisco Spaces, Genea, OAuth, PagerDuty, Secure Access, Secure Connect, Splunk, XDR | [optional] 
**Name** | Pointer to **string** | Name of integration | [optional] 
**Provider** | Pointer to **string** | Integration provider. One of Cisco, Genea, partner | [optional] 
**Tags** | Pointer to **[]string** | Integration types | [optional] 

## Methods

### NewInlineResponse200315Items

`func NewInlineResponse200315Items() *InlineResponse200315Items`

NewInlineResponse200315Items instantiates a new InlineResponse200315Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200315ItemsWithDefaults

`func NewInlineResponse200315ItemsWithDefaults() *InlineResponse200315Items`

NewInlineResponse200315ItemsWithDefaults instantiates a new InlineResponse200315Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200315Items) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200315Items) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200315Items) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200315Items) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200315Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200315Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200315Items) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200315Items) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200315Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200315Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200315Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200315Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *InlineResponse200315Items) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InlineResponse200315Items) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InlineResponse200315Items) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *InlineResponse200315Items) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200315Items) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200315Items) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200315Items) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200315Items) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


