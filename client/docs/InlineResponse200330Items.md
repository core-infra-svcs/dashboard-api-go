# InlineResponse200330Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse20111Network**](InlineResponse20111Network.md) |  | [optional] 
**RuleId** | Pointer to **string** | Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default) | [optional] 
**Type** | Pointer to **string** | Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default) | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Updated at timestamp | [optional] 
**CreatedAt** | Pointer to **time.Time** | Created at timestamp | [optional] 
**Match** | Pointer to [**InlineResponse200330Match**](InlineResponse200330Match.md) |  | [optional] 

## Methods

### NewInlineResponse200330Items

`func NewInlineResponse200330Items() *InlineResponse200330Items`

NewInlineResponse200330Items instantiates a new InlineResponse200330Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200330ItemsWithDefaults

`func NewInlineResponse200330ItemsWithDefaults() *InlineResponse200330Items`

NewInlineResponse200330ItemsWithDefaults instantiates a new InlineResponse200330Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200330Items) GetNetwork() InlineResponse20111Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200330Items) GetNetworkOk() (*InlineResponse20111Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200330Items) SetNetwork(v InlineResponse20111Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200330Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRuleId

`func (o *InlineResponse200330Items) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200330Items) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200330Items) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200330Items) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200330Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200330Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200330Items) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200330Items) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200330Items) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200330Items) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200330Items) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200330Items) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200330Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200330Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200330Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200330Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMatch

`func (o *InlineResponse200330Items) GetMatch() InlineResponse200330Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *InlineResponse200330Items) GetMatchOk() (*InlineResponse200330Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *InlineResponse200330Items) SetMatch(v InlineResponse200330Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *InlineResponse200330Items) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


