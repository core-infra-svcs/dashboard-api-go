# InlineResponse20091AllowLists

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L3** | **[]string** | Layer 3 allow list of source IP addresses excluded from inspection. Maximum 600 entries. | 
**L7** | **[]string** | Layer 7 allow list of domains excluded from inspection. Maximum 600 entries. | 
**Applications** | [**[]InlineResponse20091AllowListsApplications**](InlineResponse20091AllowListsApplications.md) | Specific applications to be excluded from inspection. Maximum 600 entries. | 

## Methods

### NewInlineResponse20091AllowLists

`func NewInlineResponse20091AllowLists(l3 []string, l7 []string, applications []InlineResponse20091AllowListsApplications, ) *InlineResponse20091AllowLists`

NewInlineResponse20091AllowLists instantiates a new InlineResponse20091AllowLists object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20091AllowListsWithDefaults

`func NewInlineResponse20091AllowListsWithDefaults() *InlineResponse20091AllowLists`

NewInlineResponse20091AllowListsWithDefaults instantiates a new InlineResponse20091AllowLists object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL3

`func (o *InlineResponse20091AllowLists) GetL3() []string`

GetL3 returns the L3 field if non-nil, zero value otherwise.

### GetL3Ok

`func (o *InlineResponse20091AllowLists) GetL3Ok() (*[]string, bool)`

GetL3Ok returns a tuple with the L3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3

`func (o *InlineResponse20091AllowLists) SetL3(v []string)`

SetL3 sets L3 field to given value.


### GetL7

`func (o *InlineResponse20091AllowLists) GetL7() []string`

GetL7 returns the L7 field if non-nil, zero value otherwise.

### GetL7Ok

`func (o *InlineResponse20091AllowLists) GetL7Ok() (*[]string, bool)`

GetL7Ok returns a tuple with the L7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7

`func (o *InlineResponse20091AllowLists) SetL7(v []string)`

SetL7 sets L7 field to given value.


### GetApplications

`func (o *InlineResponse20091AllowLists) GetApplications() []InlineResponse20091AllowListsApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse20091AllowLists) GetApplicationsOk() (*[]InlineResponse20091AllowListsApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse20091AllowLists) SetApplications(v []InlineResponse20091AllowListsApplications)`

SetApplications sets Applications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


