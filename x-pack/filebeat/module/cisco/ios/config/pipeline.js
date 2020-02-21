// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

var ciscoIOS = (function() {
    var processor = require("processor");

    var newDissect = function(pattern) {
        return new processor.Dissect({
            "tokenizer": pattern,
            "field": "message",
            "target_prefix": "",
            "ignore_failure" : true,
            "ignore_missing": true
        }).Run;
    };

    var ciscoLogPatterns = {
        //A
        //B
        "BGP_SESSION" : {
            "ADJCHANGE": newDissect("neighbor %{destination.address} %{} vpn vrf %{cisco.vrf} topology base removed from session %{}"),
        },
        "BGP" : {
            "ADJCHANGE": [
                        newDissect("neighbor %{destination.address} vpn vrf %{cisco.vrf} %{cisco.bgp.state} %{}"),
                        newDissect("neighbor %{destination.address} %{cisco.bgp.state->}"),
                ],
            "NBR_RESET": newDissect("Neighbor %{destination.address} reset (%{})"),
        },
        //c
        "CDP": {
            "DUPLEX_MISMATCH": newDissect("duplex mismatch discovered on %{source.interface.name} (%{source.interface.name}), with %{destination.address} %{destination.interface.name} (%{destination.interface.duplex})."),
            "NATIVE_VLAN_MISMATCH": newDissect("Native VLAN mismatch discovered on %{source.interface.name} (%{source.interface.vlan}), with %{destination.address} %{destination.interface.name} (%{destination.interface.vlan})."),
        },
        "CLEAR" : {
            "COUNTERS": newDissect("Clear counter on interface %{source.interface.name} by %{user.name} on %{cisco.terminal} (%{source.address})"),
        },
        "CRYPTO" : {
            "RECVD_PKT_INV_SPI": newDissect("decaps: rec'd IPSEC packet has invalid spi for destaddr=%{destination.address}, prot=%{cisco.prot}, spi=%{}, srcaddr=%{source.address}, input interface=%{source.interface.name}"),
        },
        //D
        "DUAL" : {
            "NBRCHANGE": newDissect("EIGRP-IPv4 %{}: Neighbor %{destination.address} (%{destination.interface.name}) is %{destination.interface.state}: %{cisco.reason}"),
        },
        //E
        //F
        // "FMANFP": == "SEC" Late Bind it
        //G
        //H
        "HSRP" : {
            "DIFFVIP1": newDissect("%{source.interface.name} Grp %{cisco.hsrp.group} active routers virtual IP address %{cisco.hsrp.vip} is different to the locally configured address %{source.address}"),
            "STATECHANGE": newDissect("%{source.interface.name} Grp %{cisco.hsrp.group} state %{cisco.hsrp.state}"),
        },
        //I
        "IP": {
            "DUPADDR": newDissect("Duplicate address %{source.address} on %{source.interface.name}, sourced by %{source.mac}"),
        },
        //J
        //K
        //L
        "LINEPROTO": {
            "UPDOWN": newDissect("Line protocol on Interface %{source.interface.name}, changed state to %{source.interface.state}"),
        },
        "LINK":  {
            "UPDOWN": newDissect("Interface %{interface.name}, changed state to %{interface.state}"),
        },
        //M
        //N
        //O
        //P
        //Q
        //R
        //S
        "SEC": {
            "IPACCESSLOGP": newDissect("list %{cisco.ios.access_list} %{event.outcome} " +
                "%{network.transport} %{source.address}(%{source.port}) -> " +
                "%{destination.address}(%{destination.port}), %{source.packets} packet"),
            "IPACCESSLOGDP": newDissect("list %{cisco.ios.access_list} %{event.outcome} " +
                "%{network.transport} %{source.address} -> " +
                "%{destination.address} (%{icmp.type}/%{icmp.code}), %{source.packets} packet"),
            "IPACCESSLOGRP": newDissect("list %{cisco.ios.access_list} %{event.outcome} " +
                "%{network.transport} %{source.address} -> " +
                "%{destination.address}, %{source.packets} packet"),
            "IPACCESSLOGSP": newDissect("list %{cisco.ios.access_list} %{event.outcome} " +
                "%{network.transport} %{source.address} -> " +
                "%{destination.address} (%{igmp.type}), %{source.packets} packet"),
            "IPACCESSLOGNP": newDissect("list %{cisco.ios.access_list} %{event.outcome} " +
                "%{network.iana_number} %{source.address} -> " +
                "%{destination.address}, %{source.packets} packet"),
        },
        "SEC_LOGIN": {
            "LOGIN_SUCCESS": newDissect("Login Success [user: %{user.name}] [Source: %{source.address}] [localport: %{source.port}] at %{}"),
        },
        "SSH" : {
            "SSH2_UNEXPECTED_MSG": newDissect("Unexpected message type has arrived. Terminating the connection from %{source.address}"),
        },
        "SYS": {
            "CONFIG_I": [
                        newDissect("Configured from %{cisco.terminal} by %{user.name} on %{cisco.terminal} (%{source.address})"),
                        newDissect("Configured from %{cisco.terminal} by %{user.name} on %{cisco.terminal}"),
                        newDissect("Configured from %{cisco.terminal} by console (%{source.address})"),
            ],
            "LOGOUT": newDissect("User %{user.name} has exited tty session %{cisco.tty_session_id}(%{source.address})"),
        },
        //T
        //U
        //V
        //W
        //X
        //Y
        //Z
    };

    ciscoLogPatterns.SEC.ACCESSLOGP = ciscoLogPatterns.SEC.IPACCESSLOGP;
    ciscoLogPatterns.SEC.ACCESSLOGSP = ciscoLogPatterns.SEC.IPACCESSLOGSP;
    ciscoLogPatterns.SEC.ACCESSLOGDP = ciscoLogPatterns.SEC.IPACCESSLOGDP;
    ciscoLogPatterns.SEC.ACCESSLOGNP = ciscoLogPatterns.SEC.IPACCESSLOGNP;
    ciscoLogPatterns.FMANFP = ciscoLogPatterns.SEC;

    var setLogLevel = function(evt) {
        var severity = evt.Get("event.severity");

        var levelKeyword = "";
        switch (severity) {
            case 0:
                levelKeyword = "emergencies";
                break;
            case 1:
                levelKeyword = "alerts";
                break;
            case 2:
                levelKeyword = "critical";
                break;
            case 3:
                levelKeyword = "errors";
                break;
            case 4:
                levelKeyword = "warnings";
                break;
            case 5:
                levelKeyword = "notifications";
                break;
            case 6:
                levelKeyword = "informational";
                break;
            case 7:
                levelKeyword = "debugging";
                break;
            default:
                return;
        }

        evt.Put("log.level", levelKeyword);
    };

    var copyOriginalMessage = new processor.Convert({
        fields: [
            {from: "message", to: "log.original"},
        ],
        mode: "copy",
    });

    var parseSyslogFileHeader = new processor.Chain()
        .Dissect({
            tokenizer: "%{_tmp.ts->} %{+_tmp.ts} %{+_tmp.ts->} %{log.source.address} %{event.sequence}: %{_tmp.timestamp}: %{_tmp.message}",
            field: "message",
            target_prefix: "",
        })
        .Convert({
            fields: [
                {from: "_tmp.message", to: "message"},
            ],
            mode: "rename",
        })
        .Convert({
            fields: [
                {from: "event.sequence", type: "long"},
            ],
            ignore_missing: true,
        })
        .Add(function(evt) {
            processor.Timestamp({
                field: "_tmp.timestamp",
                target_field: "@timestamp",
                timezone: evt.Get("event.timezone"),
                layouts: [
                    'Jan _2 15:04:05.999',
                    'Jan _2 15:04:05.999 MST',
                    'Jan _2 2020 15:04:05.999 MST',
                ],
                ignore_missing: true,
            }).Run(evt);
        })
        .Add(function(evt) {
            evt.Delete("_tmp");
        })
        .Build();

    var processMessage = new processor.Chain()
        // Parse the header of the message that is common to all messages.
        .Dissect({
            "tokenizer": "%{}%%{cisco.ios.facility}-%{_event_severity}-%{event.code}: %{_message}",
            "field": "message",
            "target_prefix": "",
        })
        .Add(function(evt) {
            evt.Delete("message");
            evt.Rename("_message", "message");
            evt.Delete("event.severity");
            evt.Rename("_event_severity", "event.severity");
        })
        .Convert({
            fields: [
                {from: "event.severity", type: "long"},
            ],
        })
        .Add(setLogLevel)

        .Add(function(evt) {

            var facility = evt.Get("cisco.ios.facility");
            if (!facility) {
                return;
            }
            var facility_patterns = ciscoLogPatterns[facility]
            // If no patterns for this facility then return
            if (!facility_patterns) {
                return;
            }else{
                var eventCode = evt.Get("event.code");
                if (!eventCode) {
                    return;
                }

                //Odd case for IPACCESSLOGP:fman_fp_image needs handling
                //if (eventCode =~ "\:fman_fp_image") {
                //    eventCode = eventCode.replace('\:fman_fp_image','')
                //}
                // Use a specific dissect pattern based on the event.code.
                var dissect = facility_patterns[eventCode];
                if (dissect) {
                    if (Array.isArray(dissect)) {
                        var pattern_count = dissect.length;
                        for (var i = 0; i < pattern_count; i++) {
                                var dissect_fn = dissect[i];
                                dissect_fn(evt);
                                var err = evt.Get("error.message");
                                // It has Worked so Breakout
                                if (!err) {
                                  break;
                                }
                        }
                    } else {
                           dissect(evt);
                    }
                    coerceNumbers(evt);
                    coerceIPs(evt);
                    normalizeEventOutcome(evt);
                    setNetworkType(evt);
                    setRelatedIP(evt);
                    evt.AppendTo("tags", "cisco-ios-dissect");
                    return;
                }

                // Add Special Case for Access Logs so they can be treated like firewall rules
                if (eventCode =~ "ACCESS") {
                    evt.Put("event.category", "network_traffic");
                    evt.Put("event.type", "firewall");
                }
            }
        })
        .CommunityID()
        .Build();

    var coerceNumbers = new processor.Convert({
        fields: [
            {from: "destination.port", type: "long"},
            {from: "source.port", type: "long"},
            {from: "source.packets", type: "long"},
            {from: "source.packets", to: "network.packets", type: "long"},
            {from: "icmp.type", type: "long"},
            {from: "icmp.code", type: "long"},
            {from: "igmp.type", type: "long"},
        ],
        ignore_missing: true,
    }).Run;

    //address != ip so this is allowed to fail
    var coerceIPs = new processor.Convert({
        fields: [
            {from: "destination.address", to: "destination.ip", type: "ip"},
            {from: "source.address", to: "source.ip", type: "ip"},
        ],
        ignore_missing: true,
        fail_on_error: false,
    }).Run;

    var normalizeEventOutcome = function(evt) {
        var outcome = evt.Get("event.outcome");
        if (!outcome) {
            return;
        }
        switch (outcome) {
            case "denied":
                evt.Put("event.outcome", "deny");
                break;
            case "permitted":
                evt.Put("event.outcome", "allow");
                break;
        }
    };

    var setNetworkType = function(event) {
        var ip = event.Get("source.ip");
        if (!ip) {
            return;
        }

        if (ip.indexOf(".") !== -1) {
            event.Put("network.type", "ipv4");
        } else {
            event.Put("network.type", "ipv6");
        }
    };

    var setRelatedIP = function(event) {
        var src_ip = event.Get("source.ip");
        if (src_ip) {
            event.AppendTo("related.ip", src_ip);
        }

        var dst_ip = event.Get("destination.ip");
        if (dst_ip) {
            event.AppendTo("related.ip", dst_ip);
        }
    };

    return {
        process: function(evt) {
            copyOriginalMessage.Run(evt);

            if (evt.Get("input.type") === "log") {
                parseSyslogFileHeader.Run(evt);
            }

            processMessage.Run(evt);
        },
    };
})();

function process(evt) {
    ciscoIOS.process(evt);
}
