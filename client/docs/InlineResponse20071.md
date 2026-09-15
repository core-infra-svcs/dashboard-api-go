# InlineResponse20071

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUrlPatterns** | Pointer to **[]string** | A list of URL patterns that are allowed | [optional] 
**BlockedUrlPatterns** | Pointer to **[]string** | A list of URL patterns that are blocked | [optional] 
**BlockedUrlCategories** | Pointer to [**[]InlineResponse20071BlockedUrlCategories**](InlineResponse20071BlockedUrlCategories.md) | A list of URL categories to block | [optional] 
**UrlCategoryListSize** | Pointer to **string** | URL category list size which is either &#39;topSites&#39; or &#39;fullList&#39; | [optional] 

## Methods

### NewInlineResponse20071

`func NewInlineResponse20071() *InlineResponse20071`

NewInlineResponse20071 instantiates a new InlineResponse20071 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20071WithDefaults

`func NewInlineResponse20071WithDefaults() *InlineResponse20071`

NewInlineResponse20071WithDefaults instantiates a new InlineResponse20071 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUrlPatterns

`func (o *InlineResponse20071) GetAllowedUrlPatterns() []string`

GetAllowedUrlPatterns returns the AllowedUrlPatterns field if non-nil, zero value otherwise.

### GetAllowedUrlPatternsOk

`func (o *InlineResponse20071) GetAllowedUrlPatternsOk() (*[]string, bool)`

GetAllowedUrlPatternsOk returns a tuple with the AllowedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrlPatterns

`func (o *InlineResponse20071) SetAllowedUrlPatterns(v []string)`

SetAllowedUrlPatterns sets AllowedUrlPatterns field to given value.

### HasAllowedUrlPatterns

`func (o *InlineResponse20071) HasAllowedUrlPatterns() bool`

HasAllowedUrlPatterns returns a boolean if a field has been set.

### GetBlockedUrlPatterns

`func (o *InlineResponse20071) GetBlockedUrlPatterns() []string`

GetBlockedUrlPatterns returns the BlockedUrlPatterns field if non-nil, zero value otherwise.

### GetBlockedUrlPatternsOk

`func (o *InlineResponse20071) GetBlockedUrlPatternsOk() (*[]string, bool)`

GetBlockedUrlPatternsOk returns a tuple with the BlockedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUrlPatterns

`func (o *InlineResponse20071) SetBlockedUrlPatterns(v []string)`

SetBlockedUrlPatterns sets BlockedUrlPatterns field to given value.

### HasBlockedUrlPatterns

`func (o *InlineResponse20071) HasBlockedUrlPatterns() bool`

HasBlockedUrlPatterns returns a boolean if a field has been set.

### GetBlockedUrlCategories

`func (o *InlineResponse20071) GetBlockedUrlCategories() []InlineResponse20071BlockedUrlCategories`

GetBlockedUrlCategories returns the BlockedUrlCategories field if non-nil, zero value otherwise.

### GetBlockedUrlCategoriesOk

`func (o *InlineResponse20071) GetBlockedUrlCategoriesOk() (*[]InlineResponse20071BlockedUrlCategories, bool)`

GetBlockedUrlCategoriesOk returns a tuple with the BlockedUrlCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUrlCategories

`func (o *InlineResponse20071) SetBlockedUrlCategories(v []InlineResponse20071BlockedUrlCategories)`

SetBlockedUrlCategories sets BlockedUrlCategories field to given value.

### HasBlockedUrlCategories

`func (o *InlineResponse20071) HasBlockedUrlCategories() bool`

HasBlockedUrlCategories returns a boolean if a field has been set.

### GetUrlCategoryListSize

`func (o *InlineResponse20071) GetUrlCategoryListSize() string`

GetUrlCategoryListSize returns the UrlCategoryListSize field if non-nil, zero value otherwise.

### GetUrlCategoryListSizeOk

`func (o *InlineResponse20071) GetUrlCategoryListSizeOk() (*string, bool)`

GetUrlCategoryListSizeOk returns a tuple with the UrlCategoryListSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCategoryListSize

`func (o *InlineResponse20071) SetUrlCategoryListSize(v string)`

SetUrlCategoryListSize sets UrlCategoryListSize field to given value.

### HasUrlCategoryListSize

`func (o *InlineResponse20071) HasUrlCategoryListSize() bool`

HasUrlCategoryListSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


