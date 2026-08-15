# InlineResponse200488SsidSplash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Splash mode | [optional] 
**CaptivePortal** | Pointer to [**InlineResponse200488SsidSplashCaptivePortal**](InlineResponse200488SsidSplashCaptivePortal.md) |  | [optional] 
**Redirect** | Pointer to [**InlineResponse200488SsidSplashRedirect**](InlineResponse200488SsidSplashRedirect.md) |  | [optional] 
**Timeout** | Pointer to **int32** | Splash timeout in minutes. How often users must complete the Splash | [optional] 
**Auth** | Pointer to [**InlineResponse200488SsidSplashAuth**](InlineResponse200488SsidSplashAuth.md) |  | [optional] 
**ControllerUnreachable** | Pointer to [**InlineResponse200488SsidSplashControllerUnreachable**](InlineResponse200488SsidSplashControllerUnreachable.md) |  | [optional] 
**PreAccess** | Pointer to [**InlineResponse200488SsidSplashPreAccess**](InlineResponse200488SsidSplashPreAccess.md) |  | [optional] 

## Methods

### NewInlineResponse200488SsidSplash

`func NewInlineResponse200488SsidSplash() *InlineResponse200488SsidSplash`

NewInlineResponse200488SsidSplash instantiates a new InlineResponse200488SsidSplash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200488SsidSplashWithDefaults

`func NewInlineResponse200488SsidSplashWithDefaults() *InlineResponse200488SsidSplash`

NewInlineResponse200488SsidSplashWithDefaults instantiates a new InlineResponse200488SsidSplash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200488SsidSplash) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200488SsidSplash) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200488SsidSplash) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200488SsidSplash) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCaptivePortal

`func (o *InlineResponse200488SsidSplash) GetCaptivePortal() InlineResponse200488SsidSplashCaptivePortal`

GetCaptivePortal returns the CaptivePortal field if non-nil, zero value otherwise.

### GetCaptivePortalOk

`func (o *InlineResponse200488SsidSplash) GetCaptivePortalOk() (*InlineResponse200488SsidSplashCaptivePortal, bool)`

GetCaptivePortalOk returns a tuple with the CaptivePortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptivePortal

`func (o *InlineResponse200488SsidSplash) SetCaptivePortal(v InlineResponse200488SsidSplashCaptivePortal)`

SetCaptivePortal sets CaptivePortal field to given value.

### HasCaptivePortal

`func (o *InlineResponse200488SsidSplash) HasCaptivePortal() bool`

HasCaptivePortal returns a boolean if a field has been set.

### GetRedirect

`func (o *InlineResponse200488SsidSplash) GetRedirect() InlineResponse200488SsidSplashRedirect`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *InlineResponse200488SsidSplash) GetRedirectOk() (*InlineResponse200488SsidSplashRedirect, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *InlineResponse200488SsidSplash) SetRedirect(v InlineResponse200488SsidSplashRedirect)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *InlineResponse200488SsidSplash) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineResponse200488SsidSplash) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse200488SsidSplash) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse200488SsidSplash) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse200488SsidSplash) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetAuth

`func (o *InlineResponse200488SsidSplash) GetAuth() InlineResponse200488SsidSplashAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *InlineResponse200488SsidSplash) GetAuthOk() (*InlineResponse200488SsidSplashAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *InlineResponse200488SsidSplash) SetAuth(v InlineResponse200488SsidSplashAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *InlineResponse200488SsidSplash) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetControllerUnreachable

`func (o *InlineResponse200488SsidSplash) GetControllerUnreachable() InlineResponse200488SsidSplashControllerUnreachable`

GetControllerUnreachable returns the ControllerUnreachable field if non-nil, zero value otherwise.

### GetControllerUnreachableOk

`func (o *InlineResponse200488SsidSplash) GetControllerUnreachableOk() (*InlineResponse200488SsidSplashControllerUnreachable, bool)`

GetControllerUnreachableOk returns a tuple with the ControllerUnreachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerUnreachable

`func (o *InlineResponse200488SsidSplash) SetControllerUnreachable(v InlineResponse200488SsidSplashControllerUnreachable)`

SetControllerUnreachable sets ControllerUnreachable field to given value.

### HasControllerUnreachable

`func (o *InlineResponse200488SsidSplash) HasControllerUnreachable() bool`

HasControllerUnreachable returns a boolean if a field has been set.

### GetPreAccess

`func (o *InlineResponse200488SsidSplash) GetPreAccess() InlineResponse200488SsidSplashPreAccess`

GetPreAccess returns the PreAccess field if non-nil, zero value otherwise.

### GetPreAccessOk

`func (o *InlineResponse200488SsidSplash) GetPreAccessOk() (*InlineResponse200488SsidSplashPreAccess, bool)`

GetPreAccessOk returns a tuple with the PreAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAccess

`func (o *InlineResponse200488SsidSplash) SetPreAccess(v InlineResponse200488SsidSplashPreAccess)`

SetPreAccess sets PreAccess field to given value.

### HasPreAccess

`func (o *InlineResponse200488SsidSplash) HasPreAccess() bool`

HasPreAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


