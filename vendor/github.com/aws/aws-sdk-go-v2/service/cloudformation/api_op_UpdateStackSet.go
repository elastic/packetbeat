// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateStackSetInput struct {
	_ struct{} `type:"structure"`

	// [Self-managed permissions] The accounts in which to update associated stack
	// instances. If you specify accounts, you must also specify the Regions in
	// which to update stack set instances.
	//
	// To update all the stack instances associated with this stack set, do not
	// specify the Accounts or Regions properties.
	//
	// If the stack set update includes changes to the template (that is, if the
	// TemplateBody or TemplateURL properties are specified), or the Parameters
	// property, AWS CloudFormation marks all stack instances with a status of OUTDATED
	// prior to updating the stack instances in the specified accounts and Regions.
	// If the stack set update does not include changes to the template or parameters,
	// AWS CloudFormation updates the stack instances in the specified accounts
	// and Regions, while leaving all other stack instances with their existing
	// stack instance status.
	Accounts []string `type:"list"`

	// The Amazon Resource Number (ARN) of the IAM role to use to update this stack
	// set.
	//
	// Specify an IAM role only if you are using customized administrator roles
	// to control which users or groups can manage specific stack sets within the
	// same administrator account. For more information, see Granting Permissions
	// for Stack Set Operations (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html)
	// in the AWS CloudFormation User Guide.
	//
	// If you specified a customized administrator role when you created the stack
	// set, you must specify a customized administrator role, even if it is the
	// same customized administrator role used with this stack set previously.
	AdministrationRoleARN *string `min:"20" type:"string"`

	// [Service-managed permissions] Describes whether StackSets automatically deploys
	// to AWS Organizations accounts that are added to a target organization or
	// organizational unit (OU).
	//
	// If you specify AutoDeployment, do not specify DeploymentTargets or Regions.
	AutoDeployment *AutoDeployment `type:"structure"`

	// In some cases, you must explicitly acknowledge that your stack template contains
	// certain capabilities in order for AWS CloudFormation to update the stack
	// set and its associated stack instances.
	//
	//    * CAPABILITY_IAM and CAPABILITY_NAMED_IAM Some stack templates might include
	//    resources that can affect permissions in your AWS account; for example,
	//    by creating new AWS Identity and Access Management (IAM) users. For those
	//    stacks sets, you must explicitly acknowledge this by specifying one of
	//    these capabilities. The following IAM resources require you to specify
	//    either the CAPABILITY_IAM or CAPABILITY_NAMED_IAM capability. If you have
	//    IAM resources, you can specify either capability. If you have IAM resources
	//    with custom names, you must specify CAPABILITY_NAMED_IAM. If you don't
	//    specify either of these capabilities, AWS CloudFormation returns an InsufficientCapabilities
	//    error. If your stack template contains these resources, we recommend that
	//    you review all permissions associated with them and edit their permissions
	//    if necessary. AWS::IAM::AccessKey (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html)
	//    AWS::IAM::Group (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html)
	//    AWS::IAM::InstanceProfile (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html)
	//    AWS::IAM::Policy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html)
	//    AWS::IAM::Role (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html)
	//    AWS::IAM::User (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html)
	//    AWS::IAM::UserToGroupAddition (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html)
	//    For more information, see Acknowledging IAM Resources in AWS CloudFormation
	//    Templates (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	//
	//    * CAPABILITY_AUTO_EXPAND Some templates contain macros. If your stack
	//    template contains one or more macros, and you choose to update a stack
	//    directly from the processed template, without first reviewing the resulting
	//    changes in a change set, you must acknowledge this capability. For more
	//    information, see Using AWS CloudFormation Macros to Perform Custom Processing
	//    on Templates (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html).
	//    Stack sets do not currently support macros in stack templates. (This includes
	//    the AWS::Include (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html)
	//    and AWS::Serverless (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html)
	//    transforms, which are macros hosted by AWS CloudFormation.) Even if you
	//    specify this capability, if you include a macro in your template the stack
	//    set operation will fail.
	Capabilities []Capability `type:"list"`

	// [Service-managed permissions] The AWS Organizations accounts in which to
	// update associated stack instances.
	//
	// To update all the stack instances associated with this stack set, do not
	// specify DeploymentTargets or Regions.
	//
	// If the stack set update includes changes to the template (that is, if TemplateBody
	// or TemplateURL is specified), or the Parameters, AWS CloudFormation marks
	// all stack instances with a status of OUTDATED prior to updating the stack
	// instances in the specified accounts and Regions. If the stack set update
	// does not include changes to the template or parameters, AWS CloudFormation
	// updates the stack instances in the specified accounts and Regions, while
	// leaving all other stack instances with their existing stack instance status.
	DeploymentTargets *DeploymentTargets `type:"structure"`

	// A brief description of updates that you are making.
	Description *string `min:"1" type:"string"`

	// The name of the IAM execution role to use to update the stack set. If you
	// do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole
	// role for the stack set operation.
	//
	// Specify an IAM role only if you are using customized execution roles to control
	// which stack resources users and groups can include in their stack sets.
	//
	// If you specify a customized execution role, AWS CloudFormation uses that
	// role to update the stack. If you do not specify a customized execution role,
	// AWS CloudFormation performs the update using the role previously associated
	// with the stack set, so long as you have permissions to perform operations
	// on the stack set.
	ExecutionRoleName *string `min:"1" type:"string"`

	// The unique ID for this stack set operation.
	//
	// The operation ID also functions as an idempotency token, to ensure that AWS
	// CloudFormation performs the stack set operation only once, even if you retry
	// the request multiple times. You might retry stack set operation requests
	// to ensure that AWS CloudFormation successfully received them.
	//
	// If you don't specify an operation ID, AWS CloudFormation generates one automatically.
	//
	// Repeating this stack set operation with a new operation ID retries all stack
	// instances whose status is OUTDATED.
	OperationId *string `min:"1" type:"string" idempotencyToken:"true"`

	// Preferences for how AWS CloudFormation performs this stack set operation.
	OperationPreferences *StackSetOperationPreferences `type:"structure"`

	// A list of input parameters for the stack set template.
	Parameters []Parameter `type:"list"`

	// Describes how the IAM roles required for stack set operations are created.
	// You cannot modify PermissionModel if there are stack instances associated
	// with your stack set.
	//
	//    * With self-managed permissions, you must create the administrator and
	//    execution roles required to deploy to target accounts. For more information,
	//    see Grant Self-Managed Stack Set Permissions (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html).
	//
	//    * With service-managed permissions, StackSets automatically creates the
	//    IAM roles required to deploy to accounts managed by AWS Organizations.
	//    For more information, see Grant Service-Managed Stack Set Permissions
	//    (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-service-managed.html).
	PermissionModel PermissionModels `type:"string" enum:"true"`

	// The Regions in which to update associated stack instances. If you specify
	// Regions, you must also specify accounts in which to update stack set instances.
	//
	// To update all the stack instances associated with this stack set, do not
	// specify the Accounts or Regions properties.
	//
	// If the stack set update includes changes to the template (that is, if the
	// TemplateBody or TemplateURL properties are specified), or the Parameters
	// property, AWS CloudFormation marks all stack instances with a status of OUTDATED
	// prior to updating the stack instances in the specified accounts and Regions.
	// If the stack set update does not include changes to the template or parameters,
	// AWS CloudFormation updates the stack instances in the specified accounts
	// and Regions, while leaving all other stack instances with their existing
	// stack instance status.
	Regions []string `type:"list"`

	// The name or unique ID of the stack set that you want to update.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`

	// The key-value pairs to associate with this stack set and the stacks created
	// from it. AWS CloudFormation also propagates these tags to supported resources
	// that are created in the stacks. You can specify a maximum number of 50 tags.
	//
	// If you specify tags for this parameter, those tags replace any list of tags
	// that are currently associated with this stack set. This means:
	//
	//    * If you don't specify this parameter, AWS CloudFormation doesn't modify
	//    the stack's tags.
	//
	//    * If you specify any tags using this parameter, you must specify all the
	//    tags that you want associated with this stack set, even tags you've specifed
	//    before (for example, when creating the stack set or during a previous
	//    update of the stack set.). Any tags that you don't include in the updated
	//    list of tags are removed from the stack set, and therefore from the stacks
	//    and resources as well.
	//
	//    * If you specify an empty value, AWS CloudFormation removes all currently
	//    associated tags.
	//
	// If you specify new tags as part of an UpdateStackSet action, AWS CloudFormation
	// checks to see if you have the required IAM permission to tag resources. If
	// you omit tags that are currently associated with the stack set from the list
	// of tags you specify, AWS CloudFormation assumes that you want to remove those
	// tags from the stack set, and checks to see if you have permission to untag
	// resources. If you don't have the necessary permission(s), the entire UpdateStackSet
	// action fails with an access denied error, and the stack set is not updated.
	Tags []Tag `type:"list"`

	// The structure that contains the template body, with a minimum length of 1
	// byte and a maximum length of 51,200 bytes. For more information, see Template
	// Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must specify only one of the following parameters: TemplateBody
	// or TemplateURL—or set UsePreviousTemplate to true.
	TemplateBody *string `min:"1" type:"string"`

	// The location of the file that contains the template body. The URL must point
	// to a template (maximum size: 460,800 bytes) that is located in an Amazon
	// S3 bucket. For more information, see Template Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must specify only one of the following parameters: TemplateBody
	// or TemplateURL—or set UsePreviousTemplate to true.
	TemplateURL *string `min:"1" type:"string"`

	// Use the existing template that's associated with the stack set that you're
	// updating.
	//
	// Conditional: You must specify only one of the following parameters: TemplateBody
	// or TemplateURL—or set UsePreviousTemplate to true.
	UsePreviousTemplate *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateStackSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateStackSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateStackSetInput"}
	if s.AdministrationRoleARN != nil && len(*s.AdministrationRoleARN) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("AdministrationRoleARN", 20))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.ExecutionRoleName != nil && len(*s.ExecutionRoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExecutionRoleName", 1))
	}
	if s.OperationId != nil && len(*s.OperationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OperationId", 1))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}
	if s.TemplateBody != nil && len(*s.TemplateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateBody", 1))
	}
	if s.TemplateURL != nil && len(*s.TemplateURL) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateURL", 1))
	}
	if s.OperationPreferences != nil {
		if err := s.OperationPreferences.Validate(); err != nil {
			invalidParams.AddNested("OperationPreferences", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateStackSetOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID for this stack set operation.
	OperationId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateStackSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateStackSet = "UpdateStackSet"

// UpdateStackSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Updates the stack set, and associated stack instances in the specified accounts
// and Regions.
//
// Even if the stack set operation created by updating the stack set fails (completely
// or partially, below or above a specified failure tolerance), the stack set
// is updated with your changes. Subsequent CreateStackInstances calls on the
// specified stack set use the updated stack set.
//
//    // Example sending a request using UpdateStackSetRequest.
//    req := client.UpdateStackSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/UpdateStackSet
func (c *Client) UpdateStackSetRequest(input *UpdateStackSetInput) UpdateStackSetRequest {
	op := &aws.Operation{
		Name:       opUpdateStackSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateStackSetInput{}
	}

	req := c.newRequest(op, input, &UpdateStackSetOutput{})

	return UpdateStackSetRequest{Request: req, Input: input, Copy: c.UpdateStackSetRequest}
}

// UpdateStackSetRequest is the request type for the
// UpdateStackSet API operation.
type UpdateStackSetRequest struct {
	*aws.Request
	Input *UpdateStackSetInput
	Copy  func(*UpdateStackSetInput) UpdateStackSetRequest
}

// Send marshals and sends the UpdateStackSet API request.
func (r UpdateStackSetRequest) Send(ctx context.Context) (*UpdateStackSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateStackSetResponse{
		UpdateStackSetOutput: r.Request.Data.(*UpdateStackSetOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateStackSetResponse is the response type for the
// UpdateStackSet API operation.
type UpdateStackSetResponse struct {
	*UpdateStackSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateStackSet request.
func (r *UpdateStackSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
