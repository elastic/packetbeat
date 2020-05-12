// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type CustomProviderAction struct {
	// Specifies the endpoint of the action.
	Endpoint string `pulumi:"endpoint"`
	// Specifies the name of the action.
	Name string `pulumi:"name"`
}

// CustomProviderActionInput is an input type that accepts CustomProviderActionArgs and CustomProviderActionOutput values.
// You can construct a concrete instance of `CustomProviderActionInput` via:
//
// 		 CustomProviderActionArgs{...}
//
type CustomProviderActionInput interface {
	pulumi.Input

	ToCustomProviderActionOutput() CustomProviderActionOutput
	ToCustomProviderActionOutputWithContext(context.Context) CustomProviderActionOutput
}

type CustomProviderActionArgs struct {
	// Specifies the endpoint of the action.
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// Specifies the name of the action.
	Name pulumi.StringInput `pulumi:"name"`
}

func (CustomProviderActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProviderAction)(nil)).Elem()
}

func (i CustomProviderActionArgs) ToCustomProviderActionOutput() CustomProviderActionOutput {
	return i.ToCustomProviderActionOutputWithContext(context.Background())
}

func (i CustomProviderActionArgs) ToCustomProviderActionOutputWithContext(ctx context.Context) CustomProviderActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProviderActionOutput)
}

// CustomProviderActionArrayInput is an input type that accepts CustomProviderActionArray and CustomProviderActionArrayOutput values.
// You can construct a concrete instance of `CustomProviderActionArrayInput` via:
//
// 		 CustomProviderActionArray{ CustomProviderActionArgs{...} }
//
type CustomProviderActionArrayInput interface {
	pulumi.Input

	ToCustomProviderActionArrayOutput() CustomProviderActionArrayOutput
	ToCustomProviderActionArrayOutputWithContext(context.Context) CustomProviderActionArrayOutput
}

type CustomProviderActionArray []CustomProviderActionInput

func (CustomProviderActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomProviderAction)(nil)).Elem()
}

func (i CustomProviderActionArray) ToCustomProviderActionArrayOutput() CustomProviderActionArrayOutput {
	return i.ToCustomProviderActionArrayOutputWithContext(context.Background())
}

func (i CustomProviderActionArray) ToCustomProviderActionArrayOutputWithContext(ctx context.Context) CustomProviderActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProviderActionArrayOutput)
}

type CustomProviderActionOutput struct{ *pulumi.OutputState }

func (CustomProviderActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProviderAction)(nil)).Elem()
}

func (o CustomProviderActionOutput) ToCustomProviderActionOutput() CustomProviderActionOutput {
	return o
}

func (o CustomProviderActionOutput) ToCustomProviderActionOutputWithContext(ctx context.Context) CustomProviderActionOutput {
	return o
}

// Specifies the endpoint of the action.
func (o CustomProviderActionOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CustomProviderAction) string { return v.Endpoint }).(pulumi.StringOutput)
}

// Specifies the name of the action.
func (o CustomProviderActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomProviderAction) string { return v.Name }).(pulumi.StringOutput)
}

type CustomProviderActionArrayOutput struct{ *pulumi.OutputState }

func (CustomProviderActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomProviderAction)(nil)).Elem()
}

func (o CustomProviderActionArrayOutput) ToCustomProviderActionArrayOutput() CustomProviderActionArrayOutput {
	return o
}

func (o CustomProviderActionArrayOutput) ToCustomProviderActionArrayOutputWithContext(ctx context.Context) CustomProviderActionArrayOutput {
	return o
}

func (o CustomProviderActionArrayOutput) Index(i pulumi.IntInput) CustomProviderActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomProviderAction {
		return vs[0].([]CustomProviderAction)[vs[1].(int)]
	}).(CustomProviderActionOutput)
}

type CustomProviderResourceType struct {
	// Specifies the endpoint of the route definition.
	Endpoint string `pulumi:"endpoint"`
	// Specifies the name of the route definition.
	Name string `pulumi:"name"`
	// The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.
	RoutingType *string `pulumi:"routingType"`
}

// CustomProviderResourceTypeInput is an input type that accepts CustomProviderResourceTypeArgs and CustomProviderResourceTypeOutput values.
// You can construct a concrete instance of `CustomProviderResourceTypeInput` via:
//
// 		 CustomProviderResourceTypeArgs{...}
//
type CustomProviderResourceTypeInput interface {
	pulumi.Input

	ToCustomProviderResourceTypeOutput() CustomProviderResourceTypeOutput
	ToCustomProviderResourceTypeOutputWithContext(context.Context) CustomProviderResourceTypeOutput
}

type CustomProviderResourceTypeArgs struct {
	// Specifies the endpoint of the route definition.
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// Specifies the name of the route definition.
	Name pulumi.StringInput `pulumi:"name"`
	// The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.
	RoutingType pulumi.StringPtrInput `pulumi:"routingType"`
}

func (CustomProviderResourceTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProviderResourceType)(nil)).Elem()
}

func (i CustomProviderResourceTypeArgs) ToCustomProviderResourceTypeOutput() CustomProviderResourceTypeOutput {
	return i.ToCustomProviderResourceTypeOutputWithContext(context.Background())
}

func (i CustomProviderResourceTypeArgs) ToCustomProviderResourceTypeOutputWithContext(ctx context.Context) CustomProviderResourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProviderResourceTypeOutput)
}

// CustomProviderResourceTypeArrayInput is an input type that accepts CustomProviderResourceTypeArray and CustomProviderResourceTypeArrayOutput values.
// You can construct a concrete instance of `CustomProviderResourceTypeArrayInput` via:
//
// 		 CustomProviderResourceTypeArray{ CustomProviderResourceTypeArgs{...} }
//
type CustomProviderResourceTypeArrayInput interface {
	pulumi.Input

	ToCustomProviderResourceTypeArrayOutput() CustomProviderResourceTypeArrayOutput
	ToCustomProviderResourceTypeArrayOutputWithContext(context.Context) CustomProviderResourceTypeArrayOutput
}

type CustomProviderResourceTypeArray []CustomProviderResourceTypeInput

func (CustomProviderResourceTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomProviderResourceType)(nil)).Elem()
}

func (i CustomProviderResourceTypeArray) ToCustomProviderResourceTypeArrayOutput() CustomProviderResourceTypeArrayOutput {
	return i.ToCustomProviderResourceTypeArrayOutputWithContext(context.Background())
}

func (i CustomProviderResourceTypeArray) ToCustomProviderResourceTypeArrayOutputWithContext(ctx context.Context) CustomProviderResourceTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProviderResourceTypeArrayOutput)
}

type CustomProviderResourceTypeOutput struct{ *pulumi.OutputState }

func (CustomProviderResourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProviderResourceType)(nil)).Elem()
}

func (o CustomProviderResourceTypeOutput) ToCustomProviderResourceTypeOutput() CustomProviderResourceTypeOutput {
	return o
}

func (o CustomProviderResourceTypeOutput) ToCustomProviderResourceTypeOutputWithContext(ctx context.Context) CustomProviderResourceTypeOutput {
	return o
}

// Specifies the endpoint of the route definition.
func (o CustomProviderResourceTypeOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CustomProviderResourceType) string { return v.Endpoint }).(pulumi.StringOutput)
}

// Specifies the name of the route definition.
func (o CustomProviderResourceTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomProviderResourceType) string { return v.Name }).(pulumi.StringOutput)
}

// The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.
func (o CustomProviderResourceTypeOutput) RoutingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomProviderResourceType) *string { return v.RoutingType }).(pulumi.StringPtrOutput)
}

type CustomProviderResourceTypeArrayOutput struct{ *pulumi.OutputState }

func (CustomProviderResourceTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomProviderResourceType)(nil)).Elem()
}

func (o CustomProviderResourceTypeArrayOutput) ToCustomProviderResourceTypeArrayOutput() CustomProviderResourceTypeArrayOutput {
	return o
}

func (o CustomProviderResourceTypeArrayOutput) ToCustomProviderResourceTypeArrayOutputWithContext(ctx context.Context) CustomProviderResourceTypeArrayOutput {
	return o
}

func (o CustomProviderResourceTypeArrayOutput) Index(i pulumi.IntInput) CustomProviderResourceTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomProviderResourceType {
		return vs[0].([]CustomProviderResourceType)[vs[1].(int)]
	}).(CustomProviderResourceTypeOutput)
}

type CustomProviderValidation struct {
	// The endpoint where the validation specification is located.
	Specification string `pulumi:"specification"`
}

// CustomProviderValidationInput is an input type that accepts CustomProviderValidationArgs and CustomProviderValidationOutput values.
// You can construct a concrete instance of `CustomProviderValidationInput` via:
//
// 		 CustomProviderValidationArgs{...}
//
type CustomProviderValidationInput interface {
	pulumi.Input

	ToCustomProviderValidationOutput() CustomProviderValidationOutput
	ToCustomProviderValidationOutputWithContext(context.Context) CustomProviderValidationOutput
}

type CustomProviderValidationArgs struct {
	// The endpoint where the validation specification is located.
	Specification pulumi.StringInput `pulumi:"specification"`
}

func (CustomProviderValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProviderValidation)(nil)).Elem()
}

func (i CustomProviderValidationArgs) ToCustomProviderValidationOutput() CustomProviderValidationOutput {
	return i.ToCustomProviderValidationOutputWithContext(context.Background())
}

func (i CustomProviderValidationArgs) ToCustomProviderValidationOutputWithContext(ctx context.Context) CustomProviderValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProviderValidationOutput)
}

// CustomProviderValidationArrayInput is an input type that accepts CustomProviderValidationArray and CustomProviderValidationArrayOutput values.
// You can construct a concrete instance of `CustomProviderValidationArrayInput` via:
//
// 		 CustomProviderValidationArray{ CustomProviderValidationArgs{...} }
//
type CustomProviderValidationArrayInput interface {
	pulumi.Input

	ToCustomProviderValidationArrayOutput() CustomProviderValidationArrayOutput
	ToCustomProviderValidationArrayOutputWithContext(context.Context) CustomProviderValidationArrayOutput
}

type CustomProviderValidationArray []CustomProviderValidationInput

func (CustomProviderValidationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomProviderValidation)(nil)).Elem()
}

func (i CustomProviderValidationArray) ToCustomProviderValidationArrayOutput() CustomProviderValidationArrayOutput {
	return i.ToCustomProviderValidationArrayOutputWithContext(context.Background())
}

func (i CustomProviderValidationArray) ToCustomProviderValidationArrayOutputWithContext(ctx context.Context) CustomProviderValidationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProviderValidationArrayOutput)
}

type CustomProviderValidationOutput struct{ *pulumi.OutputState }

func (CustomProviderValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProviderValidation)(nil)).Elem()
}

func (o CustomProviderValidationOutput) ToCustomProviderValidationOutput() CustomProviderValidationOutput {
	return o
}

func (o CustomProviderValidationOutput) ToCustomProviderValidationOutputWithContext(ctx context.Context) CustomProviderValidationOutput {
	return o
}

// The endpoint where the validation specification is located.
func (o CustomProviderValidationOutput) Specification() pulumi.StringOutput {
	return o.ApplyT(func(v CustomProviderValidation) string { return v.Specification }).(pulumi.StringOutput)
}

type CustomProviderValidationArrayOutput struct{ *pulumi.OutputState }

func (CustomProviderValidationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomProviderValidation)(nil)).Elem()
}

func (o CustomProviderValidationArrayOutput) ToCustomProviderValidationArrayOutput() CustomProviderValidationArrayOutput {
	return o
}

func (o CustomProviderValidationArrayOutput) ToCustomProviderValidationArrayOutputWithContext(ctx context.Context) CustomProviderValidationArrayOutput {
	return o
}

func (o CustomProviderValidationArrayOutput) Index(i pulumi.IntInput) CustomProviderValidationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomProviderValidation {
		return vs[0].([]CustomProviderValidation)[vs[1].(int)]
	}).(CustomProviderValidationOutput)
}

type GetResourcesResource struct {
	// The ID of this Resource.
	Id string `pulumi:"id"`
	// The Azure Region in which this Resource exists.
	Location string `pulumi:"location"`
	// The name of the Resource.
	Name string `pulumi:"name"`
	// A map of tags assigned to this Resource.
	Tags map[string]string `pulumi:"tags"`
	// The Resource Type of the Resources you want to list (e.g. `Microsoft.Network/virtualNetworks`). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).
	Type string `pulumi:"type"`
}

// GetResourcesResourceInput is an input type that accepts GetResourcesResourceArgs and GetResourcesResourceOutput values.
// You can construct a concrete instance of `GetResourcesResourceInput` via:
//
// 		 GetResourcesResourceArgs{...}
//
type GetResourcesResourceInput interface {
	pulumi.Input

	ToGetResourcesResourceOutput() GetResourcesResourceOutput
	ToGetResourcesResourceOutputWithContext(context.Context) GetResourcesResourceOutput
}

type GetResourcesResourceArgs struct {
	// The ID of this Resource.
	Id pulumi.StringInput `pulumi:"id"`
	// The Azure Region in which this Resource exists.
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the Resource.
	Name pulumi.StringInput `pulumi:"name"`
	// A map of tags assigned to this Resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The Resource Type of the Resources you want to list (e.g. `Microsoft.Network/virtualNetworks`). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetResourcesResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcesResource)(nil)).Elem()
}

func (i GetResourcesResourceArgs) ToGetResourcesResourceOutput() GetResourcesResourceOutput {
	return i.ToGetResourcesResourceOutputWithContext(context.Background())
}

func (i GetResourcesResourceArgs) ToGetResourcesResourceOutputWithContext(ctx context.Context) GetResourcesResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetResourcesResourceOutput)
}

// GetResourcesResourceArrayInput is an input type that accepts GetResourcesResourceArray and GetResourcesResourceArrayOutput values.
// You can construct a concrete instance of `GetResourcesResourceArrayInput` via:
//
// 		 GetResourcesResourceArray{ GetResourcesResourceArgs{...} }
//
type GetResourcesResourceArrayInput interface {
	pulumi.Input

	ToGetResourcesResourceArrayOutput() GetResourcesResourceArrayOutput
	ToGetResourcesResourceArrayOutputWithContext(context.Context) GetResourcesResourceArrayOutput
}

type GetResourcesResourceArray []GetResourcesResourceInput

func (GetResourcesResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetResourcesResource)(nil)).Elem()
}

func (i GetResourcesResourceArray) ToGetResourcesResourceArrayOutput() GetResourcesResourceArrayOutput {
	return i.ToGetResourcesResourceArrayOutputWithContext(context.Background())
}

func (i GetResourcesResourceArray) ToGetResourcesResourceArrayOutputWithContext(ctx context.Context) GetResourcesResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetResourcesResourceArrayOutput)
}

type GetResourcesResourceOutput struct{ *pulumi.OutputState }

func (GetResourcesResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcesResource)(nil)).Elem()
}

func (o GetResourcesResourceOutput) ToGetResourcesResourceOutput() GetResourcesResourceOutput {
	return o
}

func (o GetResourcesResourceOutput) ToGetResourcesResourceOutputWithContext(ctx context.Context) GetResourcesResourceOutput {
	return o
}

// The ID of this Resource.
func (o GetResourcesResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcesResource) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Region in which this Resource exists.
func (o GetResourcesResourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcesResource) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the Resource.
func (o GetResourcesResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcesResource) string { return v.Name }).(pulumi.StringOutput)
}

// A map of tags assigned to this Resource.
func (o GetResourcesResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetResourcesResource) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The Resource Type of the Resources you want to list (e.g. `Microsoft.Network/virtualNetworks`). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).
func (o GetResourcesResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcesResource) string { return v.Type }).(pulumi.StringOutput)
}

type GetResourcesResourceArrayOutput struct{ *pulumi.OutputState }

func (GetResourcesResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetResourcesResource)(nil)).Elem()
}

func (o GetResourcesResourceArrayOutput) ToGetResourcesResourceArrayOutput() GetResourcesResourceArrayOutput {
	return o
}

func (o GetResourcesResourceArrayOutput) ToGetResourcesResourceArrayOutputWithContext(ctx context.Context) GetResourcesResourceArrayOutput {
	return o
}

func (o GetResourcesResourceArrayOutput) Index(i pulumi.IntInput) GetResourcesResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetResourcesResource {
		return vs[0].([]GetResourcesResource)[vs[1].(int)]
	}).(GetResourcesResourceOutput)
}

type GetSubscriptionsSubscription struct {
	// The subscription display name.
	DisplayName string `pulumi:"displayName"`
	// The subscription location placement ID.
	LocationPlacementId string `pulumi:"locationPlacementId"`
	// The subscription quota ID.
	QuotaId string `pulumi:"quotaId"`
	// The subscription spending limit.
	SpendingLimit string `pulumi:"spendingLimit"`
	// The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
	State string `pulumi:"state"`
	// The subscription GUID.
	SubscriptionId string `pulumi:"subscriptionId"`
	// The subscription tenant ID.
	TenantId string `pulumi:"tenantId"`
}

// GetSubscriptionsSubscriptionInput is an input type that accepts GetSubscriptionsSubscriptionArgs and GetSubscriptionsSubscriptionOutput values.
// You can construct a concrete instance of `GetSubscriptionsSubscriptionInput` via:
//
// 		 GetSubscriptionsSubscriptionArgs{...}
//
type GetSubscriptionsSubscriptionInput interface {
	pulumi.Input

	ToGetSubscriptionsSubscriptionOutput() GetSubscriptionsSubscriptionOutput
	ToGetSubscriptionsSubscriptionOutputWithContext(context.Context) GetSubscriptionsSubscriptionOutput
}

type GetSubscriptionsSubscriptionArgs struct {
	// The subscription display name.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// The subscription location placement ID.
	LocationPlacementId pulumi.StringInput `pulumi:"locationPlacementId"`
	// The subscription quota ID.
	QuotaId pulumi.StringInput `pulumi:"quotaId"`
	// The subscription spending limit.
	SpendingLimit pulumi.StringInput `pulumi:"spendingLimit"`
	// The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
	State pulumi.StringInput `pulumi:"state"`
	// The subscription GUID.
	SubscriptionId pulumi.StringInput `pulumi:"subscriptionId"`
	// The subscription tenant ID.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (GetSubscriptionsSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubscriptionsSubscription)(nil)).Elem()
}

func (i GetSubscriptionsSubscriptionArgs) ToGetSubscriptionsSubscriptionOutput() GetSubscriptionsSubscriptionOutput {
	return i.ToGetSubscriptionsSubscriptionOutputWithContext(context.Background())
}

func (i GetSubscriptionsSubscriptionArgs) ToGetSubscriptionsSubscriptionOutputWithContext(ctx context.Context) GetSubscriptionsSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSubscriptionsSubscriptionOutput)
}

// GetSubscriptionsSubscriptionArrayInput is an input type that accepts GetSubscriptionsSubscriptionArray and GetSubscriptionsSubscriptionArrayOutput values.
// You can construct a concrete instance of `GetSubscriptionsSubscriptionArrayInput` via:
//
// 		 GetSubscriptionsSubscriptionArray{ GetSubscriptionsSubscriptionArgs{...} }
//
type GetSubscriptionsSubscriptionArrayInput interface {
	pulumi.Input

	ToGetSubscriptionsSubscriptionArrayOutput() GetSubscriptionsSubscriptionArrayOutput
	ToGetSubscriptionsSubscriptionArrayOutputWithContext(context.Context) GetSubscriptionsSubscriptionArrayOutput
}

type GetSubscriptionsSubscriptionArray []GetSubscriptionsSubscriptionInput

func (GetSubscriptionsSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSubscriptionsSubscription)(nil)).Elem()
}

func (i GetSubscriptionsSubscriptionArray) ToGetSubscriptionsSubscriptionArrayOutput() GetSubscriptionsSubscriptionArrayOutput {
	return i.ToGetSubscriptionsSubscriptionArrayOutputWithContext(context.Background())
}

func (i GetSubscriptionsSubscriptionArray) ToGetSubscriptionsSubscriptionArrayOutputWithContext(ctx context.Context) GetSubscriptionsSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSubscriptionsSubscriptionArrayOutput)
}

type GetSubscriptionsSubscriptionOutput struct{ *pulumi.OutputState }

func (GetSubscriptionsSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubscriptionsSubscription)(nil)).Elem()
}

func (o GetSubscriptionsSubscriptionOutput) ToGetSubscriptionsSubscriptionOutput() GetSubscriptionsSubscriptionOutput {
	return o
}

func (o GetSubscriptionsSubscriptionOutput) ToGetSubscriptionsSubscriptionOutputWithContext(ctx context.Context) GetSubscriptionsSubscriptionOutput {
	return o
}

// The subscription display name.
func (o GetSubscriptionsSubscriptionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubscriptionsSubscription) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The subscription location placement ID.
func (o GetSubscriptionsSubscriptionOutput) LocationPlacementId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubscriptionsSubscription) string { return v.LocationPlacementId }).(pulumi.StringOutput)
}

// The subscription quota ID.
func (o GetSubscriptionsSubscriptionOutput) QuotaId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubscriptionsSubscription) string { return v.QuotaId }).(pulumi.StringOutput)
}

// The subscription spending limit.
func (o GetSubscriptionsSubscriptionOutput) SpendingLimit() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubscriptionsSubscription) string { return v.SpendingLimit }).(pulumi.StringOutput)
}

// The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
func (o GetSubscriptionsSubscriptionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubscriptionsSubscription) string { return v.State }).(pulumi.StringOutput)
}

// The subscription GUID.
func (o GetSubscriptionsSubscriptionOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubscriptionsSubscription) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// The subscription tenant ID.
func (o GetSubscriptionsSubscriptionOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubscriptionsSubscription) string { return v.TenantId }).(pulumi.StringOutput)
}

type GetSubscriptionsSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (GetSubscriptionsSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSubscriptionsSubscription)(nil)).Elem()
}

func (o GetSubscriptionsSubscriptionArrayOutput) ToGetSubscriptionsSubscriptionArrayOutput() GetSubscriptionsSubscriptionArrayOutput {
	return o
}

func (o GetSubscriptionsSubscriptionArrayOutput) ToGetSubscriptionsSubscriptionArrayOutputWithContext(ctx context.Context) GetSubscriptionsSubscriptionArrayOutput {
	return o
}

func (o GetSubscriptionsSubscriptionArrayOutput) Index(i pulumi.IntInput) GetSubscriptionsSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSubscriptionsSubscription {
		return vs[0].([]GetSubscriptionsSubscription)[vs[1].(int)]
	}).(GetSubscriptionsSubscriptionOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomProviderActionOutput{})
	pulumi.RegisterOutputType(CustomProviderActionArrayOutput{})
	pulumi.RegisterOutputType(CustomProviderResourceTypeOutput{})
	pulumi.RegisterOutputType(CustomProviderResourceTypeArrayOutput{})
	pulumi.RegisterOutputType(CustomProviderValidationOutput{})
	pulumi.RegisterOutputType(CustomProviderValidationArrayOutput{})
	pulumi.RegisterOutputType(GetResourcesResourceOutput{})
	pulumi.RegisterOutputType(GetResourcesResourceArrayOutput{})
	pulumi.RegisterOutputType(GetSubscriptionsSubscriptionOutput{})
	pulumi.RegisterOutputType(GetSubscriptionsSubscriptionArrayOutput{})
}
