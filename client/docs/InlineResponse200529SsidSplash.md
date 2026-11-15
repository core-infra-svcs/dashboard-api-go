# InlineResponse200529SsidSplash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Splash mode | [optional] 
**CaptivePortal** | Pointer to [**InlineResponse200529SsidSplashCaptivePortal**](InlineResponse200529SsidSplashCaptivePortal.md) |  | [optional] 
**Redirect** | Pointer to [**InlineResponse200529SsidSplashRedirect**](InlineResponse200529SsidSplashRedirect.md) |  | [optional] 
**Timeout** | Pointer to **int32** | Splash timeout in minutes. How often users must complete the Splash | [optional] 
**Auth** | Pointer to [**InlineResponse200529SsidSplashAuth**](InlineResponse200529SsidSplashAuth.md) |  | [optional] 
**ControllerUnreachable** | Pointer to [**InlineResponse200529SsidSplashControllerUnreachable**](InlineResponse200529SsidSplashControllerUnreachable.md) |  | [optional] 
**PreAccess** | Pointer to [**InlineResponse200529SsidSplashPreAccess**](InlineResponse200529SsidSplashPreAccess.md) |  | [optional] 

## Methods

### NewInlineResponse200529SsidSplash

`func NewInlineResponse200529SsidSplash() *InlineResponse200529SsidSplash`

NewInlineResponse200529SsidSplash instantiates a new InlineResponse200529SsidSplash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200529SsidSplashWithDefaults

`func NewInlineResponse200529SsidSplashWithDefaults() *InlineResponse200529SsidSplash`

NewInlineResponse200529SsidSplashWithDefaults instantiates a new InlineResponse200529SsidSplash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200529SsidSplash) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200529SsidSplash) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200529SsidSplash) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200529SsidSplash) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCaptivePortal

`func (o *InlineResponse200529SsidSplash) GetCaptivePortal() InlineResponse200529SsidSplashCaptivePortal`

GetCaptivePortal returns the CaptivePortal field if non-nil, zero value otherwise.

### GetCaptivePortalOk

`func (o *InlineResponse200529SsidSplash) GetCaptivePortalOk() (*InlineResponse200529SsidSplashCaptivePortal, bool)`

GetCaptivePortalOk returns a tuple with the CaptivePortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptivePortal

`func (o *InlineResponse200529SsidSplash) SetCaptivePortal(v InlineResponse200529SsidSplashCaptivePortal)`

SetCaptivePortal sets CaptivePortal field to given value.

### HasCaptivePortal

`func (o *InlineResponse200529SsidSplash) HasCaptivePortal() bool`

HasCaptivePortal returns a boolean if a field has been set.

### GetRedirect

`func (o *InlineResponse200529SsidSplash) GetRedirect() InlineResponse200529SsidSplashRedirect`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *InlineResponse200529SsidSplash) GetRedirectOk() (*InlineResponse200529SsidSplashRedirect, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *InlineResponse200529SsidSplash) SetRedirect(v InlineResponse200529SsidSplashRedirect)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *InlineResponse200529SsidSplash) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineResponse200529SsidSplash) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse200529SsidSplash) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse200529SsidSplash) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse200529SsidSplash) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetAuth

`func (o *InlineResponse200529SsidSplash) GetAuth() InlineResponse200529SsidSplashAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *InlineResponse200529SsidSplash) GetAuthOk() (*InlineResponse200529SsidSplashAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *InlineResponse200529SsidSplash) SetAuth(v InlineResponse200529SsidSplashAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *InlineResponse200529SsidSplash) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetControllerUnreachable

`func (o *InlineResponse200529SsidSplash) GetControllerUnreachable() InlineResponse200529SsidSplashControllerUnreachable`

GetControllerUnreachable returns the ControllerUnreachable field if non-nil, zero value otherwise.

### GetControllerUnreachableOk

`func (o *InlineResponse200529SsidSplash) GetControllerUnreachableOk() (*InlineResponse200529SsidSplashControllerUnreachable, bool)`

GetControllerUnreachableOk returns a tuple with the ControllerUnreachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerUnreachable

`func (o *InlineResponse200529SsidSplash) SetControllerUnreachable(v InlineResponse200529SsidSplashControllerUnreachable)`

SetControllerUnreachable sets ControllerUnreachable field to given value.

### HasControllerUnreachable

`func (o *InlineResponse200529SsidSplash) HasControllerUnreachable() bool`

HasControllerUnreachable returns a boolean if a field has been set.

### GetPreAccess

`func (o *InlineResponse200529SsidSplash) GetPreAccess() InlineResponse200529SsidSplashPreAccess`

GetPreAccess returns the PreAccess field if non-nil, zero value otherwise.

### GetPreAccessOk

`func (o *InlineResponse200529SsidSplash) GetPreAccessOk() (*InlineResponse200529SsidSplashPreAccess, bool)`

GetPreAccessOk returns a tuple with the PreAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAccess

`func (o *InlineResponse200529SsidSplash) SetPreAccess(v InlineResponse200529SsidSplashPreAccess)`

SetPreAccess sets PreAccess field to given value.

### HasPreAccess

`func (o *InlineResponse200529SsidSplash) HasPreAccess() bool`

HasPreAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


