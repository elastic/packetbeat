// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

function Audit(keep_original_message) {
    var processor = require("processor");

    // The pub/sub input writes the Stackdriver LogEntry object into the message
    // field. The message needs decoded as JSON.
    var decodeJson = new processor.DecodeJSONFields({
        fields: ["message"],
        target: "json",
    });

    // Set @timetamp the LogEntry's timestamp.
    var parseTimestamp = new processor.Timestamp({
        field: "json.timestamp",
        timezone: "UTC",
        layouts: ["2006-01-02T15:04:05.999999999Z07:00"],
        tests: ["2019-06-14T03:50:10.845445834Z"],
        ignore_missing: true,
    });

    var saveOriginalMessage = function(evt) {};
    if (keep_original_message) {
        saveOriginalMessage = new processor.Convert({
            fields: [
                {from: "message", to: "event.original"}
            ],
            mode: "rename"
        });
    }

    var dropPubSubFields = function(evt) {
        evt.Delete("message");
    };

    var categorizeEvent = new processor.AddFields({
        target: "event",
        fields: {
            category: "audit",
            type: "audit-log",
        },
    });


    var saveMetadata = new processor.Convert({
        fields: [
            {from: "json.logName", to: "log.logger"},
        ],
        ignore_missing: true
    });

    var setCloudMetadata = new processor.Convert({
        fields: [
            {from: "json.resource.labels.project_id", to: "cloud.project.id"},
        ],
        ignore_missing: true
    });

    // The log includes a protoPayload field.
    // https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
    var convertLogEntry = new processor.Convert({
        fields: [
            {from: "json.protoPayload", to: "json"},
        ],
        mode: "rename",
    });

    // The LogEntry's protoPayload is moved to the json field. The protoPayload
    // contains the structured audit log fields.
    var convertProtoPayload = new processor.Convert({
        fields: [
            {from: "json.@type", to: "json.type"},

            {from: "json.authenticationInfo", to: "json.authentication_info"},
            {from: "json.authentication_info.principalEmail", to: "json.authentication_info.principal_email"},
            {from: "json.authentication_info.authoritySelector", to: "json.authentication_info.authority_selector"},

            {from: "json.authorizationInfo", to: "json.authorization_info"},

            {from: "json.methodName", to: "json.method_name"},

            {from: "json.numResponseItems", to: "json.num_response_items", type: "long"},

            {from: "json.request.@type", to: "json.request.type"},

            {from: "json.requestMetaData", to: "json.request_meta_data"},
            {from: "json.request_meta_data.callerIp", to: "json.request_meta_data.caller_ip"},
            {from: "json.request_meta_data.callerSuppliedUserAgent", to: "json.request_meta_data.caller_supplied_user_agent"},

            {from: "json.resourceName", to: "json.resource_name"},

            {from: "json.serviceName", to: "json.service_name"},

            {from: "json", to: "googlecloud.audit"},
        ],
        mode: "rename",
        ignore_missing: true,
    });

    // Copy the caller.ip  to source.ip.
    var copyAddressFields = new processor.Convert({
        fields: [
            {from: "json.request_meta_data.caller_ip", to: "source.ip"},
        ],
        fail_on_error: false,
    });

    // Copy caller_supplied_user_agent to user_agent.original.
    var copyCallerUserAgent = new processor.Convert({
        fields: [
            {from: "json.request_meta_data.caller_supplied_user_agent", to: "user_agent.original"},
        ],
        fail_on_error: false,
    });

    var pipeline = new processor.Chain()
        .Add(decodeJson)
        .Add(parseTimestamp)
        .Add(saveOriginalMessage)
        .Add(dropPubSubFields)
        .Add(categorizeEvent)
        .Add(saveMetadata)
        .Add(setCloudMetadata)
        .Add(convertLogEntry)
        .Add(convertProtoPayload)
        .Add(copyAddressFields)
        .Add(copyCallerUserAgent)
        .Build();

    return {
        process: pipeline.Run,
    };
}

var audit;

// Register params from configuration.
function register(params) {
    audit = new Audit(params.keep_original_message);
}

function process(evt) {
    return audit.process(evt);
}
