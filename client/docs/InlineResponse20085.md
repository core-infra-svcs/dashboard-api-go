# InlineResponse20085

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Current status of malware prevention | [optional] 
**AllowedUrls** | Pointer to [**[]InlineResponse20085AllowedUrls**](InlineResponse20085AllowedUrls.md) | URLs permitted by the malware detection engine | [optional] 
**AllowedFiles** | Pointer to [**[]InlineResponse20085AllowedFiles**](InlineResponse20085AllowedFiles.md) | Sha256 digests of files permitted by the malware detection engine | [optional] 

## Methods

### NewInlineResponse20085

`func NewInlineResponse20085() *InlineResponse20085`

NewInlineResponse20085 instantiates a new InlineResponse20085 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20085WithDefaults

`func NewInlineResponse20085WithDefaults() *InlineResponse20085`

NewInlineResponse20085WithDefaults instantiates a new InlineResponse20085 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20085) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20085) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20085) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20085) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAllowedUrls

`func (o *InlineResponse20085) GetAllowedUrls() []InlineResponse20085AllowedUrls`

GetAllowedUrls returns the AllowedUrls field if non-nil, zero value otherwise.

### GetAllowedUrlsOk

`func (o *InlineResponse20085) GetAllowedUrlsOk() (*[]InlineResponse20085AllowedUrls, bool)`

GetAllowedUrlsOk returns a tuple with the AllowedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrls

`func (o *InlineResponse20085) SetAllowedUrls(v []InlineResponse20085AllowedUrls)`

SetAllowedUrls sets AllowedUrls field to given value.

### HasAllowedUrls

`func (o *InlineResponse20085) HasAllowedUrls() bool`

HasAllowedUrls returns a boolean if a field has been set.

### GetAllowedFiles

`func (o *InlineResponse20085) GetAllowedFiles() []InlineResponse20085AllowedFiles`

GetAllowedFiles returns the AllowedFiles field if non-nil, zero value otherwise.

### GetAllowedFilesOk

`func (o *InlineResponse20085) GetAllowedFilesOk() (*[]InlineResponse20085AllowedFiles, bool)`

GetAllowedFilesOk returns a tuple with the AllowedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFiles

`func (o *InlineResponse20085) SetAllowedFiles(v []InlineResponse20085AllowedFiles)`

SetAllowedFiles sets AllowedFiles field to given value.

### HasAllowedFiles

`func (o *InlineResponse20085) HasAllowedFiles() bool`

HasAllowedFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


