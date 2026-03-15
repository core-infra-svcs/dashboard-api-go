# InlineResponse200314Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Identifier for integration type. One of Axis, Catalyst SD-WAN, Cisco Spaces, Genea, OAuth, PagerDuty, Secure Access, Secure Connect, Splunk, XDR | [optional] 
**Name** | Pointer to **string** | Integration name | [optional] 
**Provider** | Pointer to **string** | Integration provider. One of Cisco, Genea, partner | [optional] 
**Tags** | Pointer to **[]string** | Integration categories | [optional] 
**ShortDescription** | Pointer to **string** | Short description for integration | [optional] 
**IsDeployable** | Pointer to **bool** | Whether the integration is deployable for this organization | [optional] 
**ReleaseType** | Pointer to **string** | Release type of the integration. One of Alpha, Beta, or GA | [optional] 
**LogoUrl** | Pointer to **string** | Url for integration logo | [optional] 
**RedirectUrl** | Pointer to **string** | Integration&#39;s redirect url | [optional] 
**IsCiscoProduct** | Pointer to **bool** | Whether the integration is a Cisco product | [optional] 

## Methods

### NewInlineResponse200314Items

`func NewInlineResponse200314Items() *InlineResponse200314Items`

NewInlineResponse200314Items instantiates a new InlineResponse200314Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200314ItemsWithDefaults

`func NewInlineResponse200314ItemsWithDefaults() *InlineResponse200314Items`

NewInlineResponse200314ItemsWithDefaults instantiates a new InlineResponse200314Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200314Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200314Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200314Items) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200314Items) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200314Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200314Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200314Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200314Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *InlineResponse200314Items) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InlineResponse200314Items) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InlineResponse200314Items) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *InlineResponse200314Items) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200314Items) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200314Items) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200314Items) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200314Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetShortDescription

`func (o *InlineResponse200314Items) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *InlineResponse200314Items) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *InlineResponse200314Items) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *InlineResponse200314Items) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetIsDeployable

`func (o *InlineResponse200314Items) GetIsDeployable() bool`

GetIsDeployable returns the IsDeployable field if non-nil, zero value otherwise.

### GetIsDeployableOk

`func (o *InlineResponse200314Items) GetIsDeployableOk() (*bool, bool)`

GetIsDeployableOk returns a tuple with the IsDeployable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeployable

`func (o *InlineResponse200314Items) SetIsDeployable(v bool)`

SetIsDeployable sets IsDeployable field to given value.

### HasIsDeployable

`func (o *InlineResponse200314Items) HasIsDeployable() bool`

HasIsDeployable returns a boolean if a field has been set.

### GetReleaseType

`func (o *InlineResponse200314Items) GetReleaseType() string`

GetReleaseType returns the ReleaseType field if non-nil, zero value otherwise.

### GetReleaseTypeOk

`func (o *InlineResponse200314Items) GetReleaseTypeOk() (*string, bool)`

GetReleaseTypeOk returns a tuple with the ReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseType

`func (o *InlineResponse200314Items) SetReleaseType(v string)`

SetReleaseType sets ReleaseType field to given value.

### HasReleaseType

`func (o *InlineResponse200314Items) HasReleaseType() bool`

HasReleaseType returns a boolean if a field has been set.

### GetLogoUrl

`func (o *InlineResponse200314Items) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InlineResponse200314Items) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InlineResponse200314Items) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InlineResponse200314Items) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InlineResponse200314Items) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InlineResponse200314Items) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InlineResponse200314Items) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InlineResponse200314Items) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetIsCiscoProduct

`func (o *InlineResponse200314Items) GetIsCiscoProduct() bool`

GetIsCiscoProduct returns the IsCiscoProduct field if non-nil, zero value otherwise.

### GetIsCiscoProductOk

`func (o *InlineResponse200314Items) GetIsCiscoProductOk() (*bool, bool)`

GetIsCiscoProductOk returns a tuple with the IsCiscoProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCiscoProduct

`func (o *InlineResponse200314Items) SetIsCiscoProduct(v bool)`

SetIsCiscoProduct sets IsCiscoProduct field to given value.

### HasIsCiscoProduct

`func (o *InlineResponse200314Items) HasIsCiscoProduct() bool`

HasIsCiscoProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


