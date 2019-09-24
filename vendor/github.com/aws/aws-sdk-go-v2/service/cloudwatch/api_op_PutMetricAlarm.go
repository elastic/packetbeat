// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/PutMetricAlarmInput
type PutMetricAlarmInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether actions should be executed during any changes to the alarm
	// state. The default is TRUE.
	ActionsEnabled *bool `type:"boolean"`

	// The actions to execute when this alarm transitions to the ALARM state from
	// any other state. Each action is specified as an Amazon Resource Name (ARN).
	//
	// Valid Values: arn:aws:automate:region:ec2:stop | arn:aws:automate:region:ec2:terminate
	// | arn:aws:automate:region:ec2:recover | arn:aws:automate:region:ec2:reboot
	// | arn:aws:sns:region:account-id:sns-topic-name | arn:aws:autoscaling:region:account-id:scalingPolicy:policy-idautoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// Valid Values (for use with IAM roles): arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	// | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	// | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	AlarmActions []string `type:"list"`

	// The description for the alarm.
	AlarmDescription *string `type:"string"`

	// The name for the alarm. This name must be unique within your AWS account.
	//
	// AlarmName is a required field
	AlarmName *string `min:"1" type:"string" required:"true"`

	// The arithmetic operation to use when comparing the specified statistic and
	// threshold. The specified statistic value is used as the first operand.
	//
	// ComparisonOperator is a required field
	ComparisonOperator ComparisonOperator `type:"string" required:"true" enum:"true"`

	// The number of datapoints that must be breaching to trigger the alarm. This
	// is used only if you are setting an "M out of N" alarm. In that case, this
	// value is the M. For more information, see Evaluating an Alarm (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation)
	// in the Amazon CloudWatch User Guide.
	DatapointsToAlarm *int64 `min:"1" type:"integer"`

	// The dimensions for the metric specified in MetricName.
	Dimensions []Dimension `type:"list"`

	// Used only for alarms based on percentiles. If you specify ignore, the alarm
	// state does not change during periods with too few data points to be statistically
	// significant. If you specify evaluate or omit this parameter, the alarm is
	// always evaluated and possibly changes state no matter how many data points
	// are available. For more information, see Percentile-Based CloudWatch Alarms
	// and Low Data Samples (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#percentiles-with-low-samples).
	//
	// Valid Values: evaluate | ignore
	EvaluateLowSampleCountPercentile *string `min:"1" type:"string"`

	// The number of periods over which data is compared to the specified threshold.
	// If you are setting an alarm that requires that a number of consecutive data
	// points be breaching to trigger the alarm, this value specifies that number.
	// If you are setting an "M out of N" alarm, this value is the N.
	//
	// An alarm's total current evaluation period can be no longer than one day,
	// so this number multiplied by Period cannot be more than 86,400 seconds.
	//
	// EvaluationPeriods is a required field
	EvaluationPeriods *int64 `min:"1" type:"integer" required:"true"`

	// The percentile statistic for the metric specified in MetricName. Specify
	// a value between p0.0 and p100. When you call PutMetricAlarm and specify a
	// MetricName, you must specify either Statistic or ExtendedStatistic, but not
	// both.
	ExtendedStatistic *string `type:"string"`

	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource
	// Name (ARN).
	//
	// Valid Values: arn:aws:automate:region:ec2:stop | arn:aws:automate:region:ec2:terminate
	// | arn:aws:automate:region:ec2:recover | arn:aws:automate:region:ec2:reboot
	// | arn:aws:sns:region:account-id:sns-topic-name | arn:aws:autoscaling:region:account-id:scalingPolicy:policy-idautoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// Valid Values (for use with IAM roles): >arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	// | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	// | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	InsufficientDataActions []string `type:"list"`

	// The name for the metric associated with the alarm.
	//
	// If you are creating an alarm based on a math expression, you cannot specify
	// this parameter, or any of the Dimensions, Period, Namespace, Statistic, or
	// ExtendedStatistic parameters. Instead, you specify all this information in
	// the Metrics array.
	MetricName *string `min:"1" type:"string"`

	// An array of MetricDataQuery structures that enable you to create an alarm
	// based on the result of a metric math expression. Each item in the Metrics
	// array either retrieves a metric or performs a math expression.
	//
	// One item in the Metrics array is the expression that the alarm watches. You
	// designate this expression by setting ReturnValue to true for this object
	// in the array. For more information, see MetricDataQuery.
	//
	// If you use the Metrics parameter, you cannot include the MetricName, Dimensions,
	// Period, Namespace, Statistic, or ExtendedStatistic parameters of PutMetricAlarm
	// in the same operation. Instead, you retrieve the metrics you are using in
	// your math expression as part of the Metrics array.
	Metrics []MetricDataQuery `type:"list"`

	// The namespace for the metric associated specified in MetricName.
	Namespace *string `min:"1" type:"string"`

	// The actions to execute when this alarm transitions to an OK state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN).
	//
	// Valid Values: arn:aws:automate:region:ec2:stop | arn:aws:automate:region:ec2:terminate
	// | arn:aws:automate:region:ec2:recover | arn:aws:automate:region:ec2:reboot
	// | arn:aws:sns:region:account-id:sns-topic-name | arn:aws:autoscaling:region:account-id:scalingPolicy:policy-idautoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// Valid Values (for use with IAM roles): arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	// | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	// | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	OKActions []string `type:"list"`

	// The length, in seconds, used each time the metric specified in MetricName
	// is evaluated. Valid values are 10, 30, and any multiple of 60.
	//
	// Be sure to specify 10 or 30 only for metrics that are stored by a PutMetricData
	// call with a StorageResolution of 1. If you specify a period of 10 or 30 for
	// a metric that does not have sub-minute resolution, the alarm still attempts
	// to gather data at the period rate that you specify. In this case, it does
	// not receive data for the attempts that do not correspond to a one-minute
	// data resolution, and the alarm may often lapse into INSUFFICENT_DATA status.
	// Specifying 10 or 30 also sets this alarm as a high-resolution alarm, which
	// has a higher charge than other alarms. For more information about pricing,
	// see Amazon CloudWatch Pricing (https://aws.amazon.com/cloudwatch/pricing/).
	//
	// An alarm's total current evaluation period can be no longer than one day,
	// so Period multiplied by EvaluationPeriods cannot be more than 86,400 seconds.
	Period *int64 `min:"1" type:"integer"`

	// The statistic for the metric specified in MetricName, other than percentile.
	// For percentile statistics, use ExtendedStatistic. When you call PutMetricAlarm
	// and specify a MetricName, you must specify either Statistic or ExtendedStatistic,
	// but not both.
	Statistic Statistic `type:"string" enum:"true"`

	// A list of key-value pairs to associate with the alarm. You can associate
	// as many as 50 tags with an alarm.
	//
	// Tags can help you organize and categorize your resources. You can also use
	// them to scope user permissions, by granting a user permission to access or
	// change only resources with certain tag values.
	Tags []Tag `type:"list"`

	// The value against which the specified statistic is compared.
	//
	// Threshold is a required field
	Threshold *float64 `type:"double" required:"true"`

	// Sets how this alarm is to handle missing data points. If TreatMissingData
	// is omitted, the default behavior of missing is used. For more information,
	// see Configuring How CloudWatch Alarms Treats Missing Data (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data).
	//
	// Valid Values: breaching | notBreaching | ignore | missing
	TreatMissingData *string `min:"1" type:"string"`

	// The unit of measure for the statistic. For example, the units for the Amazon
	// EC2 NetworkIn metric are Bytes because NetworkIn tracks the number of bytes
	// that an instance receives on all network interfaces. You can also specify
	// a unit when you create a custom metric. Units help provide conceptual meaning
	// to your data. Metric data points that specify a unit of measure, such as
	// Percent, are aggregated separately.
	//
	// If you specify a unit, you must use a unit that is appropriate for the metric.
	// Otherwise, the CloudWatch alarm can get stuck in the INSUFFICIENT DATA state.
	Unit StandardUnit `type:"string" enum:"true"`
}

// String returns the string representation
func (s PutMetricAlarmInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutMetricAlarmInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutMetricAlarmInput"}

	if s.AlarmName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AlarmName"))
	}
	if s.AlarmName != nil && len(*s.AlarmName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AlarmName", 1))
	}
	if len(s.ComparisonOperator) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ComparisonOperator"))
	}
	if s.DatapointsToAlarm != nil && *s.DatapointsToAlarm < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("DatapointsToAlarm", 1))
	}
	if s.EvaluateLowSampleCountPercentile != nil && len(*s.EvaluateLowSampleCountPercentile) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EvaluateLowSampleCountPercentile", 1))
	}

	if s.EvaluationPeriods == nil {
		invalidParams.Add(aws.NewErrParamRequired("EvaluationPeriods"))
	}
	if s.EvaluationPeriods != nil && *s.EvaluationPeriods < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("EvaluationPeriods", 1))
	}
	if s.MetricName != nil && len(*s.MetricName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricName", 1))
	}
	if s.Namespace != nil && len(*s.Namespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Namespace", 1))
	}
	if s.Period != nil && *s.Period < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Period", 1))
	}

	if s.Threshold == nil {
		invalidParams.Add(aws.NewErrParamRequired("Threshold"))
	}
	if s.TreatMissingData != nil && len(*s.TreatMissingData) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TreatMissingData", 1))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Metrics != nil {
		for i, v := range s.Metrics {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Metrics", i), err.(aws.ErrInvalidParams))
			}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/PutMetricAlarmOutput
type PutMetricAlarmOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutMetricAlarmOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutMetricAlarm = "PutMetricAlarm"

// PutMetricAlarmRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Creates or updates an alarm and associates it with the specified metric or
// metric math expression.
//
// When this operation creates an alarm, the alarm state is immediately set
// to INSUFFICIENT_DATA. The alarm is then evaluated and its state is set appropriately.
// Any actions associated with the new state are then executed.
//
// When you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm.
//
// If you are an IAM user, you must have Amazon EC2 permissions for some alarm
// operations:
//
//    * iam:CreateServiceLinkedRole for all alarms with EC2 actions
//
//    * ec2:DescribeInstanceStatus and ec2:DescribeInstances for all alarms
//    on EC2 instance status metrics
//
//    * ec2:StopInstances for alarms with stop actions
//
//    * ec2:TerminateInstances for alarms with terminate actions
//
//    * No specific permissions are needed for alarms with recover actions
//
// If you have read/write permissions for Amazon CloudWatch but not for Amazon
// EC2, you can still create an alarm, but the stop or terminate actions are
// not performed. However, if you are later granted the required permissions,
// the alarm actions that you created earlier are performed.
//
// If you are using an IAM role (for example, an EC2 instance profile), you
// cannot stop or terminate the instance using alarm actions. However, you can
// still see the alarm state and perform any other actions such as Amazon SNS
// notifications or Auto Scaling policies.
//
// If you are using temporary security credentials granted using AWS STS, you
// cannot stop or terminate an EC2 instance using alarm actions.
//
// The first time you create an alarm in the AWS Management Console, the CLI,
// or by using the PutMetricAlarm API, CloudWatch creates the necessary service-linked
// role for you. The service-linked role is called AWSServiceRoleForCloudWatchEvents.
// For more information, see AWS service-linked role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role).
//
//    // Example sending a request using PutMetricAlarmRequest.
//    req := client.PutMetricAlarmRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/PutMetricAlarm
func (c *Client) PutMetricAlarmRequest(input *PutMetricAlarmInput) PutMetricAlarmRequest {
	op := &aws.Operation{
		Name:       opPutMetricAlarm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutMetricAlarmInput{}
	}

	req := c.newRequest(op, input, &PutMetricAlarmOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutMetricAlarmRequest{Request: req, Input: input, Copy: c.PutMetricAlarmRequest}
}

// PutMetricAlarmRequest is the request type for the
// PutMetricAlarm API operation.
type PutMetricAlarmRequest struct {
	*aws.Request
	Input *PutMetricAlarmInput
	Copy  func(*PutMetricAlarmInput) PutMetricAlarmRequest
}

// Send marshals and sends the PutMetricAlarm API request.
func (r PutMetricAlarmRequest) Send(ctx context.Context) (*PutMetricAlarmResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutMetricAlarmResponse{
		PutMetricAlarmOutput: r.Request.Data.(*PutMetricAlarmOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutMetricAlarmResponse is the response type for the
// PutMetricAlarm API operation.
type PutMetricAlarmResponse struct {
	*PutMetricAlarmOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutMetricAlarm request.
func (r *PutMetricAlarmResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
