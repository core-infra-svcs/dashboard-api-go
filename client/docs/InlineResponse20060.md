# InlineResponse20060

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUrlPatterns** | Pointer to **[]string** | A list of URL patterns that are allowed | [optional] 
**BlockedUrlPatterns** | Pointer to **[]string** | A list of URL patterns that are blocked | [optional] 
**BlockedUrlCategories** | Pointer to [**[]InlineResponse20060BlockedUrlCategories**](InlineResponse20060BlockedUrlCategories.md) | A list of URL categories to block | [optional] 
**UrlCategoryListSize** | Pointer to **string** | URL category list size which is either &#39;topSites&#39; or &#39;fullList&#39; | [optional] 

## Methods

### NewInlineResponse20060

`func NewInlineResponse20060() *InlineResponse20060`

NewInlineResponse20060 instantiates a new InlineResponse20060 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20060WithDefaults

`func NewInlineResponse20060WithDefaults() *InlineResponse20060`

NewInlineResponse20060WithDefaults instantiates a new InlineResponse20060 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUrlPatterns

`func (o *InlineResponse20060) GetAllowedUrlPatterns() []string`

GetAllowedUrlPatterns returns the AllowedUrlPatterns field if non-nil, zero value otherwise.

### GetAllowedUrlPatternsOk

`func (o *InlineResponse20060) GetAllowedUrlPatternsOk() (*[]string, bool)`

GetAllowedUrlPatternsOk returns a tuple with the AllowedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrlPatterns

`func (o *InlineResponse20060) SetAllowedUrlPatterns(v []string)`

SetAllowedUrlPatterns sets AllowedUrlPatterns field to given value.

### HasAllowedUrlPatterns

`func (o *InlineResponse20060) HasAllowedUrlPatterns() bool`

HasAllowedUrlPatterns returns a boolean if a field has been set.

### GetBlockedUrlPatterns

`func (o *InlineResponse20060) GetBlockedUrlPatterns() []string`

GetBlockedUrlPatterns returns the BlockedUrlPatterns field if non-nil, zero value otherwise.

### GetBlockedUrlPatternsOk

`func (o *InlineResponse20060) GetBlockedUrlPatternsOk() (*[]string, bool)`

GetBlockedUrlPatternsOk returns a tuple with the BlockedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUrlPatterns

`func (o *InlineResponse20060) SetBlockedUrlPatterns(v []string)`

SetBlockedUrlPatterns sets BlockedUrlPatterns field to given value.

### HasBlockedUrlPatterns

`func (o *InlineResponse20060) HasBlockedUrlPatterns() bool`

HasBlockedUrlPatterns returns a boolean if a field has been set.

### GetBlockedUrlCategories

`func (o *InlineResponse20060) GetBlockedUrlCategories() []InlineResponse20060BlockedUrlCategories`

GetBlockedUrlCategories returns the BlockedUrlCategories field if non-nil, zero value otherwise.

### GetBlockedUrlCategoriesOk

`func (o *InlineResponse20060) GetBlockedUrlCategoriesOk() (*[]InlineResponse20060BlockedUrlCategories, bool)`

GetBlockedUrlCategoriesOk returns a tuple with the BlockedUrlCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUrlCategories

`func (o *InlineResponse20060) SetBlockedUrlCategories(v []InlineResponse20060BlockedUrlCategories)`

SetBlockedUrlCategories sets BlockedUrlCategories field to given value.

### HasBlockedUrlCategories

`func (o *InlineResponse20060) HasBlockedUrlCategories() bool`

HasBlockedUrlCategories returns a boolean if a field has been set.

### GetUrlCategoryListSize

`func (o *InlineResponse20060) GetUrlCategoryListSize() string`

GetUrlCategoryListSize returns the UrlCategoryListSize field if non-nil, zero value otherwise.

### GetUrlCategoryListSizeOk

`func (o *InlineResponse20060) GetUrlCategoryListSizeOk() (*string, bool)`

GetUrlCategoryListSizeOk returns a tuple with the UrlCategoryListSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCategoryListSize

`func (o *InlineResponse20060) SetUrlCategoryListSize(v string)`

SetUrlCategoryListSize sets UrlCategoryListSize field to given value.

### HasUrlCategoryListSize

`func (o *InlineResponse20060) HasUrlCategoryListSize() bool`

HasUrlCategoryListSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


