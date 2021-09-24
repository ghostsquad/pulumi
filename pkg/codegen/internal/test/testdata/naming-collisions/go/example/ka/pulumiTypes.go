// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ka

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationInputType struct {
	Bar *string `pulumi:"bar"`
}

// ApplicationInputTypeInput is an input type that accepts ApplicationInputTypeArgs and ApplicationInputTypeOutput values.
// You can construct a concrete instance of `ApplicationInputTypeInput` via:
//
//          ApplicationInputTypeArgs{...}
type ApplicationInputTypeInput interface {
	pulumi.Input

	ToApplicationInputTypeOutput() ApplicationInputTypeOutput
	ToApplicationInputTypeOutputWithContext(context.Context) ApplicationInputTypeOutput
}

type ApplicationInputTypeArgs struct {
	Bar pulumi.StringPtrInput `pulumi:"bar"`
}

func (ApplicationInputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInputType)(nil)).Elem()
}

func (i ApplicationInputTypeArgs) ToApplicationInputTypeOutput() ApplicationInputTypeOutput {
	return i.ToApplicationInputTypeOutputWithContext(context.Background())
}

func (i ApplicationInputTypeArgs) ToApplicationInputTypeOutputWithContext(ctx context.Context) ApplicationInputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInputTypeOutput)
}

// ApplicationInputTypeArrayInput is an input type that accepts ApplicationInputTypeArray and ApplicationInputTypeArrayOutput values.
// You can construct a concrete instance of `ApplicationInputTypeArrayInput` via:
//
//          ApplicationInputTypeArray{ ApplicationInputTypeArgs{...} }
type ApplicationInputTypeArrayInput interface {
	pulumi.Input

	ToApplicationInputTypeArrayOutput() ApplicationInputTypeArrayOutput
	ToApplicationInputTypeArrayOutputWithContext(context.Context) ApplicationInputTypeArrayOutput
}

type ApplicationInputTypeArray []ApplicationInputTypeInput

func (ApplicationInputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationInputType)(nil)).Elem()
}

func (i ApplicationInputTypeArray) ToApplicationInputTypeArrayOutput() ApplicationInputTypeArrayOutput {
	return i.ToApplicationInputTypeArrayOutputWithContext(context.Background())
}

func (i ApplicationInputTypeArray) ToApplicationInputTypeArrayOutputWithContext(ctx context.Context) ApplicationInputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInputTypeArrayOutput)
}

type ApplicationInputTypeOutput struct{ *pulumi.OutputState }

func (ApplicationInputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInputType)(nil)).Elem()
}

func (o ApplicationInputTypeOutput) ToApplicationInputTypeOutput() ApplicationInputTypeOutput {
	return o
}

func (o ApplicationInputTypeOutput) ToApplicationInputTypeOutputWithContext(ctx context.Context) ApplicationInputTypeOutput {
	return o
}

func (o ApplicationInputTypeOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInputType) *string { return v.Bar }).(pulumi.StringPtrOutput)
}

type ApplicationInputTypeArrayOutput struct{ *pulumi.OutputState }

func (ApplicationInputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationInputType)(nil)).Elem()
}

func (o ApplicationInputTypeArrayOutput) ToApplicationInputTypeArrayOutput() ApplicationInputTypeArrayOutput {
	return o
}

func (o ApplicationInputTypeArrayOutput) ToApplicationInputTypeArrayOutputWithContext(ctx context.Context) ApplicationInputTypeArrayOutput {
	return o
}

func (o ApplicationInputTypeArrayOutput) Index(i pulumi.IntInput) ApplicationInputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationInputType {
		return vs[0].([]ApplicationInputType)[vs[1].(int)]
	}).(ApplicationInputTypeOutput)
}

type ApplicationOutputResourceOutput struct {
	Bar *string `pulumi:"bar"`
}

// ApplicationOutputResourceOutputInput is an input type that accepts ApplicationOutputResourceOutputArgs and ApplicationOutputResourceOutputOutput values.
// You can construct a concrete instance of `ApplicationOutputResourceOutputInput` via:
//
//          ApplicationOutputResourceOutputArgs{...}
type ApplicationOutputResourceOutputInput interface {
	pulumi.Input

	ToApplicationOutputResourceOutputOutput() ApplicationOutputResourceOutputOutput
	ToApplicationOutputResourceOutputOutputWithContext(context.Context) ApplicationOutputResourceOutputOutput
}

type ApplicationOutputResourceOutputArgs struct {
	Bar pulumi.StringPtrInput `pulumi:"bar"`
}

func (ApplicationOutputResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationOutputResourceOutput)(nil)).Elem()
}

func (i ApplicationOutputResourceOutputArgs) ToApplicationOutputResourceOutputOutput() ApplicationOutputResourceOutputOutput {
	return i.ToApplicationOutputResourceOutputOutputWithContext(context.Background())
}

func (i ApplicationOutputResourceOutputArgs) ToApplicationOutputResourceOutputOutputWithContext(ctx context.Context) ApplicationOutputResourceOutputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutputResourceOutputOutput)
}

func (i ApplicationOutputResourceOutputArgs) ToApplicationOutputResourceOutputPtrOutput() ApplicationOutputResourceOutputPtrOutput {
	return i.ToApplicationOutputResourceOutputPtrOutputWithContext(context.Background())
}

func (i ApplicationOutputResourceOutputArgs) ToApplicationOutputResourceOutputPtrOutputWithContext(ctx context.Context) ApplicationOutputResourceOutputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutputResourceOutputOutput).ToApplicationOutputResourceOutputPtrOutputWithContext(ctx)
}

// ApplicationOutputResourceOutputPtrInput is an input type that accepts ApplicationOutputResourceOutputArgs, ApplicationOutputResourceOutputPtr and ApplicationOutputResourceOutputPtrOutput values.
// You can construct a concrete instance of `ApplicationOutputResourceOutputPtrInput` via:
//
//          ApplicationOutputResourceOutputArgs{...}
//
//  or:
//
//          nil
type ApplicationOutputResourceOutputPtrInput interface {
	pulumi.Input

	ToApplicationOutputResourceOutputPtrOutput() ApplicationOutputResourceOutputPtrOutput
	ToApplicationOutputResourceOutputPtrOutputWithContext(context.Context) ApplicationOutputResourceOutputPtrOutput
}

type applicationOutputResourceOutputPtrType ApplicationOutputResourceOutputArgs

func ApplicationOutputResourceOutputPtr(v *ApplicationOutputResourceOutputArgs) ApplicationOutputResourceOutputPtrInput {
	return (*applicationOutputResourceOutputPtrType)(v)
}

func (*applicationOutputResourceOutputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOutputResourceOutput)(nil)).Elem()
}

func (i *applicationOutputResourceOutputPtrType) ToApplicationOutputResourceOutputPtrOutput() ApplicationOutputResourceOutputPtrOutput {
	return i.ToApplicationOutputResourceOutputPtrOutputWithContext(context.Background())
}

func (i *applicationOutputResourceOutputPtrType) ToApplicationOutputResourceOutputPtrOutputWithContext(ctx context.Context) ApplicationOutputResourceOutputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutputResourceOutputPtrOutput)
}

type ApplicationOutputResourceOutputOutput struct{ *pulumi.OutputState }

func (ApplicationOutputResourceOutputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationOutputResourceOutput)(nil)).Elem()
}

func (o ApplicationOutputResourceOutputOutput) ToApplicationOutputResourceOutputOutput() ApplicationOutputResourceOutputOutput {
	return o
}

func (o ApplicationOutputResourceOutputOutput) ToApplicationOutputResourceOutputOutputWithContext(ctx context.Context) ApplicationOutputResourceOutputOutput {
	return o
}

func (o ApplicationOutputResourceOutputOutput) ToApplicationOutputResourceOutputPtrOutput() ApplicationOutputResourceOutputPtrOutput {
	return o.ToApplicationOutputResourceOutputPtrOutputWithContext(context.Background())
}

func (o ApplicationOutputResourceOutputOutput) ToApplicationOutputResourceOutputPtrOutputWithContext(ctx context.Context) ApplicationOutputResourceOutputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationOutputResourceOutput) *ApplicationOutputResourceOutput {
		return &v
	}).(ApplicationOutputResourceOutputPtrOutput)
}

func (o ApplicationOutputResourceOutputOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationOutputResourceOutput) *string { return v.Bar }).(pulumi.StringPtrOutput)
}

type ApplicationOutputResourceOutputPtrOutput struct{ *pulumi.OutputState }

func (ApplicationOutputResourceOutputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOutputResourceOutput)(nil)).Elem()
}

func (o ApplicationOutputResourceOutputPtrOutput) ToApplicationOutputResourceOutputPtrOutput() ApplicationOutputResourceOutputPtrOutput {
	return o
}

func (o ApplicationOutputResourceOutputPtrOutput) ToApplicationOutputResourceOutputPtrOutputWithContext(ctx context.Context) ApplicationOutputResourceOutputPtrOutput {
	return o
}

func (o ApplicationOutputResourceOutputPtrOutput) Elem() ApplicationOutputResourceOutputOutput {
	return o.ApplyT(func(v *ApplicationOutputResourceOutput) ApplicationOutputResourceOutput {
		if v != nil {
			return *v
		}
		var ret ApplicationOutputResourceOutput
		return ret
	}).(ApplicationOutputResourceOutputOutput)
}

func (o ApplicationOutputResourceOutputPtrOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationOutputResourceOutput) *string {
		if v == nil {
			return nil
		}
		return v.Bar
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationInputTypeOutput{})
	pulumi.RegisterOutputType(ApplicationInputTypeArrayOutput{})
	pulumi.RegisterOutputType(ApplicationOutputResourceOutputOutput{})
	pulumi.RegisterOutputType(ApplicationOutputResourceOutputPtrOutput{})
}
