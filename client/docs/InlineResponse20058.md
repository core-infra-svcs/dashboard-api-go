# InlineResponse20058

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Current status of malware prevention | [optional] 
**AllowedUrls** | Pointer to [**[]InlineResponse20058AllowedUrls**](InlineResponse20058AllowedUrls.md) | URLs permitted by the malware detection engine | [optional] 
**AllowedFiles** | Pointer to [**[]InlineResponse20058AllowedFiles**](InlineResponse20058AllowedFiles.md) | Sha256 digests of files permitted by the malware detection engine | [optional] 

## Methods

### NewInlineResponse20058

`func NewInlineResponse20058() *InlineResponse20058`

NewInlineResponse20058 instantiates a new InlineResponse20058 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20058WithDefaults

`func NewInlineResponse20058WithDefaults() *InlineResponse20058`

NewInlineResponse20058WithDefaults instantiates a new InlineResponse20058 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20058) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20058) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20058) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20058) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAllowedUrls

`func (o *InlineResponse20058) GetAllowedUrls() []InlineResponse20058AllowedUrls`

GetAllowedUrls returns the AllowedUrls field if non-nil, zero value otherwise.

### GetAllowedUrlsOk

`func (o *InlineResponse20058) GetAllowedUrlsOk() (*[]InlineResponse20058AllowedUrls, bool)`

GetAllowedUrlsOk returns a tuple with the AllowedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrls

`func (o *InlineResponse20058) SetAllowedUrls(v []InlineResponse20058AllowedUrls)`

SetAllowedUrls sets AllowedUrls field to given value.

### HasAllowedUrls

`func (o *InlineResponse20058) HasAllowedUrls() bool`

HasAllowedUrls returns a boolean if a field has been set.

### GetAllowedFiles

`func (o *InlineResponse20058) GetAllowedFiles() []InlineResponse20058AllowedFiles`

GetAllowedFiles returns the AllowedFiles field if non-nil, zero value otherwise.

### GetAllowedFilesOk

`func (o *InlineResponse20058) GetAllowedFilesOk() (*[]InlineResponse20058AllowedFiles, bool)`

GetAllowedFilesOk returns a tuple with the AllowedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFiles

`func (o *InlineResponse20058) SetAllowedFiles(v []InlineResponse20058AllowedFiles)`

SetAllowedFiles sets AllowedFiles field to given value.

### HasAllowedFiles

`func (o *InlineResponse20058) HasAllowedFiles() bool`

HasAllowedFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


