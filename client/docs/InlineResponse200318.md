# InlineResponse200318

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VppAccountId** | Pointer to **string** | The id of the VPP Account | [optional] 
**ContentToken** | Pointer to **string** | The VPP service token | [optional] 
**Email** | Pointer to **string** | The email address associated with the VPP account | [optional] 
**Name** | Pointer to **string** | The name of the VPP account | [optional] 
**AllowedAdmins** | Pointer to **string** | The allowed admins for the VPP account | [optional] 
**NetworkIdAdmins** | Pointer to **string** | The network IDs of the admins for the VPP account | [optional] 
**AssignableNetworks** | Pointer to **string** | The assignable networks for the VPP account | [optional] 
**AssignableNetworkIds** | Pointer to **[]string** | The network IDs of the assignable networks for the VPP account | [optional] 
**VppLocationId** | Pointer to **string** | The VPP location ID | [optional] 
**VppLocationName** | Pointer to **string** | The VPP location name | [optional] 
**LastSyncedAt** | Pointer to **string** | The last time the VPP account was synced | [optional] 
**LastForceSyncedAt** | Pointer to **string** | The last time the VPP account was force synced | [optional] 
**ParsedToken** | Pointer to [**OrganizationsOrganizationIdSmVppAccountsParsedToken**](OrganizationsOrganizationIdSmVppAccountsParsedToken.md) |  | [optional] 
**Id** | Pointer to **string** | The id of the VPP Account | [optional] 
**VppServiceToken** | Pointer to **string** | The VPP Account&#39;s Service Token | [optional] 

## Methods

### NewInlineResponse200318

`func NewInlineResponse200318() *InlineResponse200318`

NewInlineResponse200318 instantiates a new InlineResponse200318 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200318WithDefaults

`func NewInlineResponse200318WithDefaults() *InlineResponse200318`

NewInlineResponse200318WithDefaults instantiates a new InlineResponse200318 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVppAccountId

`func (o *InlineResponse200318) GetVppAccountId() string`

GetVppAccountId returns the VppAccountId field if non-nil, zero value otherwise.

### GetVppAccountIdOk

`func (o *InlineResponse200318) GetVppAccountIdOk() (*string, bool)`

GetVppAccountIdOk returns a tuple with the VppAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppAccountId

`func (o *InlineResponse200318) SetVppAccountId(v string)`

SetVppAccountId sets VppAccountId field to given value.

### HasVppAccountId

`func (o *InlineResponse200318) HasVppAccountId() bool`

HasVppAccountId returns a boolean if a field has been set.

### GetContentToken

`func (o *InlineResponse200318) GetContentToken() string`

GetContentToken returns the ContentToken field if non-nil, zero value otherwise.

### GetContentTokenOk

`func (o *InlineResponse200318) GetContentTokenOk() (*string, bool)`

GetContentTokenOk returns a tuple with the ContentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentToken

`func (o *InlineResponse200318) SetContentToken(v string)`

SetContentToken sets ContentToken field to given value.

### HasContentToken

`func (o *InlineResponse200318) HasContentToken() bool`

HasContentToken returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse200318) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse200318) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse200318) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse200318) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200318) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200318) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200318) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200318) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllowedAdmins

`func (o *InlineResponse200318) GetAllowedAdmins() string`

GetAllowedAdmins returns the AllowedAdmins field if non-nil, zero value otherwise.

### GetAllowedAdminsOk

`func (o *InlineResponse200318) GetAllowedAdminsOk() (*string, bool)`

GetAllowedAdminsOk returns a tuple with the AllowedAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAdmins

`func (o *InlineResponse200318) SetAllowedAdmins(v string)`

SetAllowedAdmins sets AllowedAdmins field to given value.

### HasAllowedAdmins

`func (o *InlineResponse200318) HasAllowedAdmins() bool`

HasAllowedAdmins returns a boolean if a field has been set.

### GetNetworkIdAdmins

`func (o *InlineResponse200318) GetNetworkIdAdmins() string`

GetNetworkIdAdmins returns the NetworkIdAdmins field if non-nil, zero value otherwise.

### GetNetworkIdAdminsOk

`func (o *InlineResponse200318) GetNetworkIdAdminsOk() (*string, bool)`

GetNetworkIdAdminsOk returns a tuple with the NetworkIdAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdAdmins

`func (o *InlineResponse200318) SetNetworkIdAdmins(v string)`

SetNetworkIdAdmins sets NetworkIdAdmins field to given value.

### HasNetworkIdAdmins

`func (o *InlineResponse200318) HasNetworkIdAdmins() bool`

HasNetworkIdAdmins returns a boolean if a field has been set.

### GetAssignableNetworks

`func (o *InlineResponse200318) GetAssignableNetworks() string`

GetAssignableNetworks returns the AssignableNetworks field if non-nil, zero value otherwise.

### GetAssignableNetworksOk

`func (o *InlineResponse200318) GetAssignableNetworksOk() (*string, bool)`

GetAssignableNetworksOk returns a tuple with the AssignableNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableNetworks

`func (o *InlineResponse200318) SetAssignableNetworks(v string)`

SetAssignableNetworks sets AssignableNetworks field to given value.

### HasAssignableNetworks

`func (o *InlineResponse200318) HasAssignableNetworks() bool`

HasAssignableNetworks returns a boolean if a field has been set.

### GetAssignableNetworkIds

`func (o *InlineResponse200318) GetAssignableNetworkIds() []string`

GetAssignableNetworkIds returns the AssignableNetworkIds field if non-nil, zero value otherwise.

### GetAssignableNetworkIdsOk

`func (o *InlineResponse200318) GetAssignableNetworkIdsOk() (*[]string, bool)`

GetAssignableNetworkIdsOk returns a tuple with the AssignableNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableNetworkIds

`func (o *InlineResponse200318) SetAssignableNetworkIds(v []string)`

SetAssignableNetworkIds sets AssignableNetworkIds field to given value.

### HasAssignableNetworkIds

`func (o *InlineResponse200318) HasAssignableNetworkIds() bool`

HasAssignableNetworkIds returns a boolean if a field has been set.

### GetVppLocationId

`func (o *InlineResponse200318) GetVppLocationId() string`

GetVppLocationId returns the VppLocationId field if non-nil, zero value otherwise.

### GetVppLocationIdOk

`func (o *InlineResponse200318) GetVppLocationIdOk() (*string, bool)`

GetVppLocationIdOk returns a tuple with the VppLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppLocationId

`func (o *InlineResponse200318) SetVppLocationId(v string)`

SetVppLocationId sets VppLocationId field to given value.

### HasVppLocationId

`func (o *InlineResponse200318) HasVppLocationId() bool`

HasVppLocationId returns a boolean if a field has been set.

### GetVppLocationName

`func (o *InlineResponse200318) GetVppLocationName() string`

GetVppLocationName returns the VppLocationName field if non-nil, zero value otherwise.

### GetVppLocationNameOk

`func (o *InlineResponse200318) GetVppLocationNameOk() (*string, bool)`

GetVppLocationNameOk returns a tuple with the VppLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppLocationName

`func (o *InlineResponse200318) SetVppLocationName(v string)`

SetVppLocationName sets VppLocationName field to given value.

### HasVppLocationName

`func (o *InlineResponse200318) HasVppLocationName() bool`

HasVppLocationName returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *InlineResponse200318) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *InlineResponse200318) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *InlineResponse200318) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *InlineResponse200318) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetLastForceSyncedAt

`func (o *InlineResponse200318) GetLastForceSyncedAt() string`

GetLastForceSyncedAt returns the LastForceSyncedAt field if non-nil, zero value otherwise.

### GetLastForceSyncedAtOk

`func (o *InlineResponse200318) GetLastForceSyncedAtOk() (*string, bool)`

GetLastForceSyncedAtOk returns a tuple with the LastForceSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastForceSyncedAt

`func (o *InlineResponse200318) SetLastForceSyncedAt(v string)`

SetLastForceSyncedAt sets LastForceSyncedAt field to given value.

### HasLastForceSyncedAt

`func (o *InlineResponse200318) HasLastForceSyncedAt() bool`

HasLastForceSyncedAt returns a boolean if a field has been set.

### GetParsedToken

`func (o *InlineResponse200318) GetParsedToken() OrganizationsOrganizationIdSmVppAccountsParsedToken`

GetParsedToken returns the ParsedToken field if non-nil, zero value otherwise.

### GetParsedTokenOk

`func (o *InlineResponse200318) GetParsedTokenOk() (*OrganizationsOrganizationIdSmVppAccountsParsedToken, bool)`

GetParsedTokenOk returns a tuple with the ParsedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedToken

`func (o *InlineResponse200318) SetParsedToken(v OrganizationsOrganizationIdSmVppAccountsParsedToken)`

SetParsedToken sets ParsedToken field to given value.

### HasParsedToken

`func (o *InlineResponse200318) HasParsedToken() bool`

HasParsedToken returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200318) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200318) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200318) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200318) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVppServiceToken

`func (o *InlineResponse200318) GetVppServiceToken() string`

GetVppServiceToken returns the VppServiceToken field if non-nil, zero value otherwise.

### GetVppServiceTokenOk

`func (o *InlineResponse200318) GetVppServiceTokenOk() (*string, bool)`

GetVppServiceTokenOk returns a tuple with the VppServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppServiceToken

`func (o *InlineResponse200318) SetVppServiceToken(v string)`

SetVppServiceToken sets VppServiceToken field to given value.

### HasVppServiceToken

`func (o *InlineResponse200318) HasVppServiceToken() bool`

HasVppServiceToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


