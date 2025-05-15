# InlineResponse200114

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The network or organization identifier | [optional] 
**OrganizationWide** | Pointer to **bool** | If the data returned is organization-wide. False indicates the data is network-wide. | [optional] 
**NetworkId** | Pointer to **string** | The network identifier | [optional] 
**Type** | Pointer to **string** | The type of PII request | [optional] 
**Mac** | Pointer to **string** | The MAC address of the PII request | [optional] 
**Datasets** | Pointer to **string** | The stringified array of datasets related to the provided key that should be deleted. | [optional] 
**Status** | Pointer to **string** | The status of the PII request | [optional] 
**CreatedAt** | Pointer to **int32** | The request&#39;s creation time | [optional] 
**CompletedAt** | Pointer to **int32** | The request&#39;s completion time | [optional] 

## Methods

### NewInlineResponse200114

`func NewInlineResponse200114() *InlineResponse200114`

NewInlineResponse200114 instantiates a new InlineResponse200114 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200114WithDefaults

`func NewInlineResponse200114WithDefaults() *InlineResponse200114`

NewInlineResponse200114WithDefaults instantiates a new InlineResponse200114 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200114) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200114) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200114) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200114) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationWide

`func (o *InlineResponse200114) GetOrganizationWide() bool`

GetOrganizationWide returns the OrganizationWide field if non-nil, zero value otherwise.

### GetOrganizationWideOk

`func (o *InlineResponse200114) GetOrganizationWideOk() (*bool, bool)`

GetOrganizationWideOk returns a tuple with the OrganizationWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationWide

`func (o *InlineResponse200114) SetOrganizationWide(v bool)`

SetOrganizationWide sets OrganizationWide field to given value.

### HasOrganizationWide

`func (o *InlineResponse200114) HasOrganizationWide() bool`

HasOrganizationWide returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200114) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200114) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200114) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200114) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200114) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200114) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200114) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200114) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200114) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200114) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200114) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200114) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetDatasets

`func (o *InlineResponse200114) GetDatasets() string`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *InlineResponse200114) GetDatasetsOk() (*string, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *InlineResponse200114) SetDatasets(v string)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *InlineResponse200114) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200114) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200114) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200114) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200114) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200114) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200114) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200114) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200114) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *InlineResponse200114) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *InlineResponse200114) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *InlineResponse200114) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *InlineResponse200114) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


