# InlineResponse200493SsidSplash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Splash mode | [optional] 
**CaptivePortal** | Pointer to [**InlineResponse200493SsidSplashCaptivePortal**](InlineResponse200493SsidSplashCaptivePortal.md) |  | [optional] 
**Redirect** | Pointer to [**InlineResponse200493SsidSplashRedirect**](InlineResponse200493SsidSplashRedirect.md) |  | [optional] 
**Timeout** | Pointer to **int32** | Splash timeout in minutes. How often users must complete the Splash | [optional] 
**Auth** | Pointer to [**InlineResponse200493SsidSplashAuth**](InlineResponse200493SsidSplashAuth.md) |  | [optional] 
**ControllerUnreachable** | Pointer to [**InlineResponse200493SsidSplashControllerUnreachable**](InlineResponse200493SsidSplashControllerUnreachable.md) |  | [optional] 
**PreAccess** | Pointer to [**InlineResponse200493SsidSplashPreAccess**](InlineResponse200493SsidSplashPreAccess.md) |  | [optional] 

## Methods

### NewInlineResponse200493SsidSplash

`func NewInlineResponse200493SsidSplash() *InlineResponse200493SsidSplash`

NewInlineResponse200493SsidSplash instantiates a new InlineResponse200493SsidSplash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200493SsidSplashWithDefaults

`func NewInlineResponse200493SsidSplashWithDefaults() *InlineResponse200493SsidSplash`

NewInlineResponse200493SsidSplashWithDefaults instantiates a new InlineResponse200493SsidSplash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200493SsidSplash) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200493SsidSplash) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200493SsidSplash) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200493SsidSplash) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCaptivePortal

`func (o *InlineResponse200493SsidSplash) GetCaptivePortal() InlineResponse200493SsidSplashCaptivePortal`

GetCaptivePortal returns the CaptivePortal field if non-nil, zero value otherwise.

### GetCaptivePortalOk

`func (o *InlineResponse200493SsidSplash) GetCaptivePortalOk() (*InlineResponse200493SsidSplashCaptivePortal, bool)`

GetCaptivePortalOk returns a tuple with the CaptivePortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptivePortal

`func (o *InlineResponse200493SsidSplash) SetCaptivePortal(v InlineResponse200493SsidSplashCaptivePortal)`

SetCaptivePortal sets CaptivePortal field to given value.

### HasCaptivePortal

`func (o *InlineResponse200493SsidSplash) HasCaptivePortal() bool`

HasCaptivePortal returns a boolean if a field has been set.

### GetRedirect

`func (o *InlineResponse200493SsidSplash) GetRedirect() InlineResponse200493SsidSplashRedirect`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *InlineResponse200493SsidSplash) GetRedirectOk() (*InlineResponse200493SsidSplashRedirect, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *InlineResponse200493SsidSplash) SetRedirect(v InlineResponse200493SsidSplashRedirect)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *InlineResponse200493SsidSplash) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineResponse200493SsidSplash) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse200493SsidSplash) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse200493SsidSplash) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse200493SsidSplash) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetAuth

`func (o *InlineResponse200493SsidSplash) GetAuth() InlineResponse200493SsidSplashAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *InlineResponse200493SsidSplash) GetAuthOk() (*InlineResponse200493SsidSplashAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *InlineResponse200493SsidSplash) SetAuth(v InlineResponse200493SsidSplashAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *InlineResponse200493SsidSplash) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetControllerUnreachable

`func (o *InlineResponse200493SsidSplash) GetControllerUnreachable() InlineResponse200493SsidSplashControllerUnreachable`

GetControllerUnreachable returns the ControllerUnreachable field if non-nil, zero value otherwise.

### GetControllerUnreachableOk

`func (o *InlineResponse200493SsidSplash) GetControllerUnreachableOk() (*InlineResponse200493SsidSplashControllerUnreachable, bool)`

GetControllerUnreachableOk returns a tuple with the ControllerUnreachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerUnreachable

`func (o *InlineResponse200493SsidSplash) SetControllerUnreachable(v InlineResponse200493SsidSplashControllerUnreachable)`

SetControllerUnreachable sets ControllerUnreachable field to given value.

### HasControllerUnreachable

`func (o *InlineResponse200493SsidSplash) HasControllerUnreachable() bool`

HasControllerUnreachable returns a boolean if a field has been set.

### GetPreAccess

`func (o *InlineResponse200493SsidSplash) GetPreAccess() InlineResponse200493SsidSplashPreAccess`

GetPreAccess returns the PreAccess field if non-nil, zero value otherwise.

### GetPreAccessOk

`func (o *InlineResponse200493SsidSplash) GetPreAccessOk() (*InlineResponse200493SsidSplashPreAccess, bool)`

GetPreAccessOk returns a tuple with the PreAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAccess

`func (o *InlineResponse200493SsidSplash) SetPreAccess(v InlineResponse200493SsidSplashPreAccess)`

SetPreAccess sets PreAccess field to given value.

### HasPreAccess

`func (o *InlineResponse200493SsidSplash) HasPreAccess() bool`

HasPreAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


