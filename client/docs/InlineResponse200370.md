# InlineResponse200370

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** | Short name of the early access feature | [optional] 
**Name** | Pointer to **string** | Name of the early access feature | [optional] 
**Descriptions** | Pointer to [**OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions**](OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions.md) |  | [optional] 
**Topic** | Pointer to **string** | Topic of the early access feature | [optional] 
**IsOrgScopedOnly** | Pointer to **bool** | If this early access feature can only be opted in for the entire organization | [optional] 
**DocumentationLink** | Pointer to **string** | Link to the documentation of this early access feature | [optional] 
**SupportLink** | Pointer to **string** | Link to get support for this early access feature | [optional] 
**PrivacyLink** | Pointer to **string** | Link to the privacy policy for this early access feature | [optional] 
**Advantage** | Pointer to **bool** | If this early access feature requires an Advantage license | [optional] 
**AdvantageTrial** | Pointer to **bool** | If this Advantage early access feature offers an introductory trial | [optional] 

## Methods

### NewInlineResponse200370

`func NewInlineResponse200370() *InlineResponse200370`

NewInlineResponse200370 instantiates a new InlineResponse200370 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200370WithDefaults

`func NewInlineResponse200370WithDefaults() *InlineResponse200370`

NewInlineResponse200370WithDefaults instantiates a new InlineResponse200370 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *InlineResponse200370) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *InlineResponse200370) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *InlineResponse200370) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *InlineResponse200370) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200370) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200370) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200370) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200370) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescriptions

`func (o *InlineResponse200370) GetDescriptions() OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *InlineResponse200370) GetDescriptionsOk() (*OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *InlineResponse200370) SetDescriptions(v OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *InlineResponse200370) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetTopic

`func (o *InlineResponse200370) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *InlineResponse200370) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *InlineResponse200370) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *InlineResponse200370) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetIsOrgScopedOnly

`func (o *InlineResponse200370) GetIsOrgScopedOnly() bool`

GetIsOrgScopedOnly returns the IsOrgScopedOnly field if non-nil, zero value otherwise.

### GetIsOrgScopedOnlyOk

`func (o *InlineResponse200370) GetIsOrgScopedOnlyOk() (*bool, bool)`

GetIsOrgScopedOnlyOk returns a tuple with the IsOrgScopedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgScopedOnly

`func (o *InlineResponse200370) SetIsOrgScopedOnly(v bool)`

SetIsOrgScopedOnly sets IsOrgScopedOnly field to given value.

### HasIsOrgScopedOnly

`func (o *InlineResponse200370) HasIsOrgScopedOnly() bool`

HasIsOrgScopedOnly returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *InlineResponse200370) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *InlineResponse200370) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *InlineResponse200370) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *InlineResponse200370) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetSupportLink

`func (o *InlineResponse200370) GetSupportLink() string`

GetSupportLink returns the SupportLink field if non-nil, zero value otherwise.

### GetSupportLinkOk

`func (o *InlineResponse200370) GetSupportLinkOk() (*string, bool)`

GetSupportLinkOk returns a tuple with the SupportLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLink

`func (o *InlineResponse200370) SetSupportLink(v string)`

SetSupportLink sets SupportLink field to given value.

### HasSupportLink

`func (o *InlineResponse200370) HasSupportLink() bool`

HasSupportLink returns a boolean if a field has been set.

### GetPrivacyLink

`func (o *InlineResponse200370) GetPrivacyLink() string`

GetPrivacyLink returns the PrivacyLink field if non-nil, zero value otherwise.

### GetPrivacyLinkOk

`func (o *InlineResponse200370) GetPrivacyLinkOk() (*string, bool)`

GetPrivacyLinkOk returns a tuple with the PrivacyLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLink

`func (o *InlineResponse200370) SetPrivacyLink(v string)`

SetPrivacyLink sets PrivacyLink field to given value.

### HasPrivacyLink

`func (o *InlineResponse200370) HasPrivacyLink() bool`

HasPrivacyLink returns a boolean if a field has been set.

### GetAdvantage

`func (o *InlineResponse200370) GetAdvantage() bool`

GetAdvantage returns the Advantage field if non-nil, zero value otherwise.

### GetAdvantageOk

`func (o *InlineResponse200370) GetAdvantageOk() (*bool, bool)`

GetAdvantageOk returns a tuple with the Advantage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvantage

`func (o *InlineResponse200370) SetAdvantage(v bool)`

SetAdvantage sets Advantage field to given value.

### HasAdvantage

`func (o *InlineResponse200370) HasAdvantage() bool`

HasAdvantage returns a boolean if a field has been set.

### GetAdvantageTrial

`func (o *InlineResponse200370) GetAdvantageTrial() bool`

GetAdvantageTrial returns the AdvantageTrial field if non-nil, zero value otherwise.

### GetAdvantageTrialOk

`func (o *InlineResponse200370) GetAdvantageTrialOk() (*bool, bool)`

GetAdvantageTrialOk returns a tuple with the AdvantageTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvantageTrial

`func (o *InlineResponse200370) SetAdvantageTrial(v bool)`

SetAdvantageTrial sets AdvantageTrial field to given value.

### HasAdvantageTrial

`func (o *InlineResponse200370) HasAdvantageTrial() bool`

HasAdvantageTrial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


