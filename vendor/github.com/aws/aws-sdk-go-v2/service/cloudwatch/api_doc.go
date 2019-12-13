// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudwatch provides the client and types for making API
// requests to CloudWatch.
//
// Amazon CloudWatch monitors your Amazon Web Services (AWS) resources and the
// applications you run on AWS in real time. You can use CloudWatch to collect
// and track metrics, which are the variables you want to measure for your resources
// and applications.
//
// CloudWatch alarms send notifications or automatically change the resources
// you are monitoring based on rules that you define. For example, you can monitor
// the CPU usage and disk reads and writes of your Amazon EC2 instances. Then,
// use this data to determine whether you should launch additional instances
// to handle increased load. You can also use this data to stop under-used instances
// to save money.
//
// In addition to monitoring the built-in metrics that come with AWS, you can
// monitor your own custom metrics. With CloudWatch, you gain system-wide visibility
// into resource utilization, application performance, and operational health.
//
// See https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01 for more information on this service.
//
// See cloudwatch package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatch/
//
// Using the Client
//
// To use CloudWatch with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the CloudWatch client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatch/#New
package cloudwatch
