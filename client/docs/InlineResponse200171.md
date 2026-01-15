# InlineResponse200171

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default. | [optional] 
**HelloTimerInSeconds** | Pointer to **int32** | Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds. | [optional] 
**DeadTimerInSeconds** | Pointer to **int32** | Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535 | [optional] 
**Areas** | Pointer to [**[]InlineResponse200171Areas**](InlineResponse200171Areas.md) | OSPF areas | [optional] 
**V3** | Pointer to [**InlineResponse200171V3**](InlineResponse200171V3.md) |  | [optional] 
**Md5AuthenticationEnabled** | Pointer to **bool** | Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default. | [optional] 
**Md5AuthenticationKey** | Pointer to [**InlineResponse200171Md5AuthenticationKey**](InlineResponse200171Md5AuthenticationKey.md) |  | [optional] 
**Vrf** | Pointer to [**InlineResponse200171Vrf**](InlineResponse200171Vrf.md) |  | [optional] 

## Methods

### NewInlineResponse200171

`func NewInlineResponse200171() *InlineResponse200171`

NewInlineResponse200171 instantiates a new InlineResponse200171 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200171WithDefaults

`func NewInlineResponse200171WithDefaults() *InlineResponse200171`

NewInlineResponse200171WithDefaults instantiates a new InlineResponse200171 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200171) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200171) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200171) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200171) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHelloTimerInSeconds

`func (o *InlineResponse200171) GetHelloTimerInSeconds() int32`

GetHelloTimerInSeconds returns the HelloTimerInSeconds field if non-nil, zero value otherwise.

### GetHelloTimerInSecondsOk

`func (o *InlineResponse200171) GetHelloTimerInSecondsOk() (*int32, bool)`

GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTimerInSeconds

`func (o *InlineResponse200171) SetHelloTimerInSeconds(v int32)`

SetHelloTimerInSeconds sets HelloTimerInSeconds field to given value.

### HasHelloTimerInSeconds

`func (o *InlineResponse200171) HasHelloTimerInSeconds() bool`

HasHelloTimerInSeconds returns a boolean if a field has been set.

### GetDeadTimerInSeconds

`func (o *InlineResponse200171) GetDeadTimerInSeconds() int32`

GetDeadTimerInSeconds returns the DeadTimerInSeconds field if non-nil, zero value otherwise.

### GetDeadTimerInSecondsOk

`func (o *InlineResponse200171) GetDeadTimerInSecondsOk() (*int32, bool)`

GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadTimerInSeconds

`func (o *InlineResponse200171) SetDeadTimerInSeconds(v int32)`

SetDeadTimerInSeconds sets DeadTimerInSeconds field to given value.

### HasDeadTimerInSeconds

`func (o *InlineResponse200171) HasDeadTimerInSeconds() bool`

HasDeadTimerInSeconds returns a boolean if a field has been set.

### GetAreas

`func (o *InlineResponse200171) GetAreas() []InlineResponse200171Areas`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *InlineResponse200171) GetAreasOk() (*[]InlineResponse200171Areas, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *InlineResponse200171) SetAreas(v []InlineResponse200171Areas)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *InlineResponse200171) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetV3

`func (o *InlineResponse200171) GetV3() InlineResponse200171V3`

GetV3 returns the V3 field if non-nil, zero value otherwise.

### GetV3Ok

`func (o *InlineResponse200171) GetV3Ok() (*InlineResponse200171V3, bool)`

GetV3Ok returns a tuple with the V3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3

`func (o *InlineResponse200171) SetV3(v InlineResponse200171V3)`

SetV3 sets V3 field to given value.

### HasV3

`func (o *InlineResponse200171) HasV3() bool`

HasV3 returns a boolean if a field has been set.

### GetMd5AuthenticationEnabled

`func (o *InlineResponse200171) GetMd5AuthenticationEnabled() bool`

GetMd5AuthenticationEnabled returns the Md5AuthenticationEnabled field if non-nil, zero value otherwise.

### GetMd5AuthenticationEnabledOk

`func (o *InlineResponse200171) GetMd5AuthenticationEnabledOk() (*bool, bool)`

GetMd5AuthenticationEnabledOk returns a tuple with the Md5AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5AuthenticationEnabled

`func (o *InlineResponse200171) SetMd5AuthenticationEnabled(v bool)`

SetMd5AuthenticationEnabled sets Md5AuthenticationEnabled field to given value.

### HasMd5AuthenticationEnabled

`func (o *InlineResponse200171) HasMd5AuthenticationEnabled() bool`

HasMd5AuthenticationEnabled returns a boolean if a field has been set.

### GetMd5AuthenticationKey

`func (o *InlineResponse200171) GetMd5AuthenticationKey() InlineResponse200171Md5AuthenticationKey`

GetMd5AuthenticationKey returns the Md5AuthenticationKey field if non-nil, zero value otherwise.

### GetMd5AuthenticationKeyOk

`func (o *InlineResponse200171) GetMd5AuthenticationKeyOk() (*InlineResponse200171Md5AuthenticationKey, bool)`

GetMd5AuthenticationKeyOk returns a tuple with the Md5AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5AuthenticationKey

`func (o *InlineResponse200171) SetMd5AuthenticationKey(v InlineResponse200171Md5AuthenticationKey)`

SetMd5AuthenticationKey sets Md5AuthenticationKey field to given value.

### HasMd5AuthenticationKey

`func (o *InlineResponse200171) HasMd5AuthenticationKey() bool`

HasMd5AuthenticationKey returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse200171) GetVrf() InlineResponse200171Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200171) GetVrfOk() (*InlineResponse200171Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200171) SetVrf(v InlineResponse200171Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200171) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


