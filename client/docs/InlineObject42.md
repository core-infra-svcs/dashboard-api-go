# InlineObject42

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Set mode to &#39;enabled&#39; to enable malware prevention, otherwise &#39;disabled&#39; | 
**AllowedUrls** | Pointer to [**[]NetworksNetworkIdApplianceSecurityMalwareAllowedUrls**](NetworksNetworkIdApplianceSecurityMalwareAllowedUrls.md) | The urls that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing | [optional] 
**AllowedFiles** | Pointer to [**[]NetworksNetworkIdApplianceSecurityMalwareAllowedFiles**](NetworksNetworkIdApplianceSecurityMalwareAllowedFiles.md) | The sha256 digests of files that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing | [optional] 

## Methods

### NewInlineObject42

`func NewInlineObject42(mode string, ) *InlineObject42`

NewInlineObject42 instantiates a new InlineObject42 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject42WithDefaults

`func NewInlineObject42WithDefaults() *InlineObject42`

NewInlineObject42WithDefaults instantiates a new InlineObject42 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject42) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject42) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject42) SetMode(v string)`

SetMode sets Mode field to given value.


### GetAllowedUrls

`func (o *InlineObject42) GetAllowedUrls() []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls`

GetAllowedUrls returns the AllowedUrls field if non-nil, zero value otherwise.

### GetAllowedUrlsOk

`func (o *InlineObject42) GetAllowedUrlsOk() (*[]NetworksNetworkIdApplianceSecurityMalwareAllowedUrls, bool)`

GetAllowedUrlsOk returns a tuple with the AllowedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrls

`func (o *InlineObject42) SetAllowedUrls(v []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls)`

SetAllowedUrls sets AllowedUrls field to given value.

### HasAllowedUrls

`func (o *InlineObject42) HasAllowedUrls() bool`

HasAllowedUrls returns a boolean if a field has been set.

### GetAllowedFiles

`func (o *InlineObject42) GetAllowedFiles() []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles`

GetAllowedFiles returns the AllowedFiles field if non-nil, zero value otherwise.

### GetAllowedFilesOk

`func (o *InlineObject42) GetAllowedFilesOk() (*[]NetworksNetworkIdApplianceSecurityMalwareAllowedFiles, bool)`

GetAllowedFilesOk returns a tuple with the AllowedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFiles

`func (o *InlineObject42) SetAllowedFiles(v []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles)`

SetAllowedFiles sets AllowedFiles field to given value.

### HasAllowedFiles

`func (o *InlineObject42) HasAllowedFiles() bool`

HasAllowedFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


