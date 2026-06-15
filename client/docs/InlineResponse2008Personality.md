# InlineResponse2008Personality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The type of interface, &#39;wan&#39; or &#39;lan&#39;, the port is configured as | [optional] 
**IsFlexible** | Pointer to **bool** | Indicates whether the port interface type can convert between WAN and LAN | [optional] 
**Layer** | Pointer to [**InlineResponse2008PersonalityLayer**](InlineResponse2008PersonalityLayer.md) |  | [optional] 

## Methods

### NewInlineResponse2008Personality

`func NewInlineResponse2008Personality() *InlineResponse2008Personality`

NewInlineResponse2008Personality instantiates a new InlineResponse2008Personality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008PersonalityWithDefaults

`func NewInlineResponse2008PersonalityWithDefaults() *InlineResponse2008Personality`

NewInlineResponse2008PersonalityWithDefaults instantiates a new InlineResponse2008Personality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse2008Personality) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse2008Personality) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse2008Personality) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse2008Personality) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIsFlexible

`func (o *InlineResponse2008Personality) GetIsFlexible() bool`

GetIsFlexible returns the IsFlexible field if non-nil, zero value otherwise.

### GetIsFlexibleOk

`func (o *InlineResponse2008Personality) GetIsFlexibleOk() (*bool, bool)`

GetIsFlexibleOk returns a tuple with the IsFlexible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlexible

`func (o *InlineResponse2008Personality) SetIsFlexible(v bool)`

SetIsFlexible sets IsFlexible field to given value.

### HasIsFlexible

`func (o *InlineResponse2008Personality) HasIsFlexible() bool`

HasIsFlexible returns a boolean if a field has been set.

### GetLayer

`func (o *InlineResponse2008Personality) GetLayer() InlineResponse2008PersonalityLayer`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *InlineResponse2008Personality) GetLayerOk() (*InlineResponse2008PersonalityLayer, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *InlineResponse2008Personality) SetLayer(v InlineResponse2008PersonalityLayer)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *InlineResponse2008Personality) HasLayer() bool`

HasLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


