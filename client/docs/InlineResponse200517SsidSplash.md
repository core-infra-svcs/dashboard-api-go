# InlineResponse200517SsidSplash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Splash mode | [optional] 
**CaptivePortal** | Pointer to [**InlineResponse200517SsidSplashCaptivePortal**](InlineResponse200517SsidSplashCaptivePortal.md) |  | [optional] 
**Redirect** | Pointer to [**InlineResponse200517SsidSplashRedirect**](InlineResponse200517SsidSplashRedirect.md) |  | [optional] 
**Timeout** | Pointer to **int32** | Splash timeout in minutes. How often users must complete the Splash | [optional] 
**Auth** | Pointer to [**InlineResponse200517SsidSplashAuth**](InlineResponse200517SsidSplashAuth.md) |  | [optional] 
**ControllerUnreachable** | Pointer to [**InlineResponse200517SsidSplashControllerUnreachable**](InlineResponse200517SsidSplashControllerUnreachable.md) |  | [optional] 
**PreAccess** | Pointer to [**InlineResponse200517SsidSplashPreAccess**](InlineResponse200517SsidSplashPreAccess.md) |  | [optional] 

## Methods

### NewInlineResponse200517SsidSplash

`func NewInlineResponse200517SsidSplash() *InlineResponse200517SsidSplash`

NewInlineResponse200517SsidSplash instantiates a new InlineResponse200517SsidSplash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200517SsidSplashWithDefaults

`func NewInlineResponse200517SsidSplashWithDefaults() *InlineResponse200517SsidSplash`

NewInlineResponse200517SsidSplashWithDefaults instantiates a new InlineResponse200517SsidSplash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200517SsidSplash) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200517SsidSplash) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200517SsidSplash) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200517SsidSplash) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCaptivePortal

`func (o *InlineResponse200517SsidSplash) GetCaptivePortal() InlineResponse200517SsidSplashCaptivePortal`

GetCaptivePortal returns the CaptivePortal field if non-nil, zero value otherwise.

### GetCaptivePortalOk

`func (o *InlineResponse200517SsidSplash) GetCaptivePortalOk() (*InlineResponse200517SsidSplashCaptivePortal, bool)`

GetCaptivePortalOk returns a tuple with the CaptivePortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptivePortal

`func (o *InlineResponse200517SsidSplash) SetCaptivePortal(v InlineResponse200517SsidSplashCaptivePortal)`

SetCaptivePortal sets CaptivePortal field to given value.

### HasCaptivePortal

`func (o *InlineResponse200517SsidSplash) HasCaptivePortal() bool`

HasCaptivePortal returns a boolean if a field has been set.

### GetRedirect

`func (o *InlineResponse200517SsidSplash) GetRedirect() InlineResponse200517SsidSplashRedirect`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *InlineResponse200517SsidSplash) GetRedirectOk() (*InlineResponse200517SsidSplashRedirect, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *InlineResponse200517SsidSplash) SetRedirect(v InlineResponse200517SsidSplashRedirect)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *InlineResponse200517SsidSplash) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineResponse200517SsidSplash) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse200517SsidSplash) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse200517SsidSplash) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse200517SsidSplash) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetAuth

`func (o *InlineResponse200517SsidSplash) GetAuth() InlineResponse200517SsidSplashAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *InlineResponse200517SsidSplash) GetAuthOk() (*InlineResponse200517SsidSplashAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *InlineResponse200517SsidSplash) SetAuth(v InlineResponse200517SsidSplashAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *InlineResponse200517SsidSplash) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetControllerUnreachable

`func (o *InlineResponse200517SsidSplash) GetControllerUnreachable() InlineResponse200517SsidSplashControllerUnreachable`

GetControllerUnreachable returns the ControllerUnreachable field if non-nil, zero value otherwise.

### GetControllerUnreachableOk

`func (o *InlineResponse200517SsidSplash) GetControllerUnreachableOk() (*InlineResponse200517SsidSplashControllerUnreachable, bool)`

GetControllerUnreachableOk returns a tuple with the ControllerUnreachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerUnreachable

`func (o *InlineResponse200517SsidSplash) SetControllerUnreachable(v InlineResponse200517SsidSplashControllerUnreachable)`

SetControllerUnreachable sets ControllerUnreachable field to given value.

### HasControllerUnreachable

`func (o *InlineResponse200517SsidSplash) HasControllerUnreachable() bool`

HasControllerUnreachable returns a boolean if a field has been set.

### GetPreAccess

`func (o *InlineResponse200517SsidSplash) GetPreAccess() InlineResponse200517SsidSplashPreAccess`

GetPreAccess returns the PreAccess field if non-nil, zero value otherwise.

### GetPreAccessOk

`func (o *InlineResponse200517SsidSplash) GetPreAccessOk() (*InlineResponse200517SsidSplashPreAccess, bool)`

GetPreAccessOk returns a tuple with the PreAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAccess

`func (o *InlineResponse200517SsidSplash) SetPreAccess(v InlineResponse200517SsidSplashPreAccess)`

SetPreAccess sets PreAccess field to given value.

### HasPreAccess

`func (o *InlineResponse200517SsidSplash) HasPreAccess() bool`

HasPreAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


