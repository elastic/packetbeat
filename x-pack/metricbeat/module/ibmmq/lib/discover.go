// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package ibmmqlib

/*
  Copyright (c) IBM Corporation 2016, 2018

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific

   Contributors:
     Mark Taylor - Initial Contribution
*/

/*
Functions in this file discover the data available from a queue manager
via the MQ V9 pub/sub monitoring feature. Each metric (element) is
found by discovering the types of metric, and the types are found by first
discovering the classes. Sample program amqsrua is shipped with MQ V9 to
give a good demonstration of the process, which is followed here.
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ibm-messaging/mq-golang/ibmmq"

	"github.com/elastic/beats/libbeat/logp"
)

// MonElement describes the real metric element generated by MQ
type MonElement struct {
	Parent      *MonType
	Description string // An English phrase describing the element
	MetricName  string // Reformatted description suitable as label
	Datatype    int32
	Values      map[string]int64
}

// MonType describes the "types" of data generated by MQ. Each class generates
// one or more type of data such as OPENCLOSE (from STATMQI class) or
// LOG (from DISK class)
type MonType struct {
	Parent       *MonClass
	Name         string
	Description  string
	ObjectTopic  string // topic for actual data responses
	elementTopic string // discovery of elements
	Elements     map[int]*MonElement
	subHobj      map[string]ibmmq.MQObject
}

// MonClass described the "classes" of data generated by MQ, such as DISK and CPU
type MonClass struct {
	Parent      *AllMetrics
	Name        string
	Description string
	typesTopic  string
	flags       int
	Types       map[int]*MonType
}

// The AllMetrics structure is the top of the tree, holding the set of classes.
type AllMetrics struct {
	Classes map[int]*MonClass
}

// QMgrMapKey can never be a real object name and is therefore useful in
// maps that may contain only this single entry
const QMgrMapKey = "@self"

// Metrics is the global variable for the tree of data
var Metrics AllMetrics

var qList []string

/*
DiscoverAndSubscribe does all the work of finding the
different resources available from a queue manager and
issuing the MQSUB calls to collect the data
*/
func DiscoverAndSubscribe(queueList string, checkQueueList bool, metaPrefix string) error {
	var err error
	// What metrics can the queue manager provide?
	if err == nil {
		err = discoverStats(metaPrefix)
	}

	// Which queues have we been asked to monitor? Expand wildcards
	// to explicit names so that subscriptions work.
	if err == nil {
		if checkQueueList {
			discoverQueues(queueList)
		} else {
			qList = strings.Split(queueList, ",")
		}
	}

	// Subscribe to all of the various topics
	if err == nil {
		createSubscriptions()
	}

	return err
}

func discoverClasses(metaPrefix string) error {
	var data []byte
	var sub ibmmq.MQObject
	var err error
	var rootTopic string

	// Have to know the starting point for the topic that tells about classes
	if metaPrefix == "" {
		rootTopic = "$SYS/MQ/INFO/QMGR/" + qMgr.Name + "/Monitor/METADATA/CLASSES"
	} else {
		rootTopic = metaPrefix + "/INFO/QMGR/" + qMgr.Name + "/Monitor/METADATA/CLASSES"
	}
	sub, err = subscribe(rootTopic)
	if err == nil {
		data, err = getMessage(true)
		sub.Close(0)

		elemList, _ := ParsePCFResponse(data)

		for i := 0; i < len(elemList); i++ {
			if elemList[i].Type != ibmmq.MQCFT_GROUP {
				continue
			}
			group := elemList[i]
			cl := new(MonClass)
			classIndex := 0
			cl.Types = make(map[int]*MonType)
			cl.Parent = &Metrics

			for j := 0; j < len(group.GroupList); j++ {
				elem := group.GroupList[j]
				switch elem.Parameter {
				case ibmmq.MQIAMO_MONITOR_CLASS:
					classIndex = int(elem.Int64Value[0])
				case ibmmq.MQIAMO_MONITOR_FLAGS:
					cl.flags = int(elem.Int64Value[0])
				case ibmmq.MQCAMO_MONITOR_CLASS:
					cl.Name = elem.String[0]
				case ibmmq.MQCAMO_MONITOR_DESC:
					cl.Description = elem.String[0]
				case ibmmq.MQCA_TOPIC_STRING:
					cl.typesTopic = elem.String[0]
				default:
					return fmt.Errorf("Unknown parameter %d in class discovery", elem.Parameter)
				}
			}
			Metrics.Classes[classIndex] = cl
		}
	}

	subsOpened = true
	return err
}

func discoverTypes(cl *MonClass) error {
	var data []byte
	var sub ibmmq.MQObject
	var err error

	sub, err = subscribe(cl.typesTopic)
	if err == nil {
		data, err = getMessage(true)
		sub.Close(0)

		elemList, _ := ParsePCFResponse(data)

		for i := 0; i < len(elemList); i++ {
			if elemList[i].Type != ibmmq.MQCFT_GROUP {
				continue
			}

			group := elemList[i]
			ty := new(MonType)
			ty.Elements = make(map[int]*MonElement)
			ty.subHobj = make(map[string]ibmmq.MQObject)

			typeIndex := 0
			ty.Parent = cl

			for j := 0; j < len(group.GroupList); j++ {
				elem := group.GroupList[j]
				switch elem.Parameter {

				case ibmmq.MQIAMO_MONITOR_TYPE:
					typeIndex = int(elem.Int64Value[0])
				case ibmmq.MQCAMO_MONITOR_TYPE:
					ty.Name = elem.String[0]
				case ibmmq.MQCAMO_MONITOR_DESC:
					ty.Description = elem.String[0]
				case ibmmq.MQCA_TOPIC_STRING:
					ty.elementTopic = elem.String[0]
				default:
					return fmt.Errorf("Unknown parameter %d in type discovery", elem.Parameter)
				}
			}
			cl.Types[typeIndex] = ty
		}
	}
	return err
}

func discoverElements(ty *MonType) error {
	var err error
	var data []byte
	var sub ibmmq.MQObject
	var elem *MonElement

	sub, err = subscribe(ty.elementTopic)
	if err == nil {
		data, err = getMessage(true)
		sub.Close(0)

		elemList, _ := ParsePCFResponse(data)

		for i := 0; i < len(elemList); i++ {

			if elemList[i].Type == ibmmq.MQCFT_STRING && elemList[i].Parameter == ibmmq.MQCA_TOPIC_STRING {
				ty.ObjectTopic = elemList[i].String[0]
				continue
			}

			if elemList[i].Type != ibmmq.MQCFT_GROUP {
				continue
			}

			group := elemList[i]

			elem = new(MonElement)
			elementIndex := 0
			elem.Parent = ty
			elem.Values = make(map[string]int64)

			for j := 0; j < len(group.GroupList); j++ {
				e := group.GroupList[j]
				switch e.Parameter {

				case ibmmq.MQIAMO_MONITOR_ELEMENT:
					elementIndex = int(e.Int64Value[0])
				case ibmmq.MQIAMO_MONITOR_DATATYPE:
					elem.Datatype = int32(e.Int64Value[0])
				case ibmmq.MQCAMO_MONITOR_DESC:
					elem.Description = e.String[0]
				default:
					return fmt.Errorf("Unknown parameter %d in type discovery", e.Parameter)
				}
			}

			elem.MetricName = formatDescription(elem)
			ty.Elements[elementIndex] = elem
		}
	}

	return err
}

/*
Discover the complete set of available statistics in the queue manager
by working through the classes, types and individual elements.

Then discover the list of individual queues we have been asked for.
*/
func discoverStats(metaPrefix string) error {
	var err error

	// Start with an empty set of information about the available stats
	Metrics.Classes = make(map[int]*MonClass)

	// Then get the list of CLASSES
	err = discoverClasses(metaPrefix)

	// For each CLASS, discover the TYPEs of data available
	if err == nil {
		for _, cl := range Metrics.Classes {
			err = discoverTypes(cl)
			// And for each CLASS, discover the actual statistics elements
			if err == nil {
				for _, ty := range cl.Types {
					err = discoverElements(ty)
				}
			}
		}
		//

	}

	return err
}

/*
discoverQueues lists the queues that match all of the configured
patterns.

The patterns must match the MQ rule - asterisk on the end of the
string only.

If a bad pattern is used, or no queues exist that match the pattern
then an error is reported but we continue processing other patterns.

An alternative would be to list ALL the queues (though that could be a
long list, and we would really have to worry about TRUNCATED message retrieval),
and then use a more general regexp match. Something for a later update
perhaps.
*/
func discoverQueues(monitoredQueues string) error {
	var err error
	var elem *ibmmq.PCFParameter
	var datalen int

	if monitoredQueues == "" {
		return err
	}

	queues := strings.Split(monitoredQueues, ",")
	for i := 0; i < len(queues) && err == nil; i++ {
		var buf []byte

		pattern := queues[i]

		if strings.Count(pattern, "*") > 1 ||
			(strings.Count(pattern, "*") == 1 && !strings.HasSuffix(pattern, "*")) {
			return fmt.Errorf("Queue pattern '%s' is not valid", pattern)
		}

		putmqmd := ibmmq.NewMQMD()
		pmo := ibmmq.NewMQPMO()

		pmo.Options = ibmmq.MQPMO_NO_SYNCPOINT
		pmo.Options |= ibmmq.MQPMO_NEW_MSG_ID
		pmo.Options |= ibmmq.MQPMO_NEW_CORREL_ID
		pmo.Options |= ibmmq.MQPMO_FAIL_IF_QUIESCING

		putmqmd.Format = "MQADMIN"
		putmqmd.ReplyToQ = replyQObj.Name
		putmqmd.MsgType = ibmmq.MQMT_REQUEST
		putmqmd.Report = ibmmq.MQRO_PASS_DISCARD_AND_EXPIRY

		cfh := ibmmq.NewMQCFH()

		// Can allow all the other fields to default
		cfh.Command = ibmmq.MQCMD_INQUIRE_Q_NAMES

		// Add the parameters one at a time into a buffer
		pcfparm := new(ibmmq.PCFParameter)
		pcfparm.Type = ibmmq.MQCFT_STRING
		pcfparm.Parameter = ibmmq.MQCA_Q_NAME
		pcfparm.String = []string{pattern}
		cfh.ParameterCount++
		buf = append(buf, pcfparm.Bytes()...)

		pcfparm = new(ibmmq.PCFParameter)
		pcfparm.Type = ibmmq.MQCFT_INTEGER
		pcfparm.Parameter = ibmmq.MQIA_Q_TYPE
		pcfparm.Int64Value = []int64{int64(ibmmq.MQQT_LOCAL)}
		cfh.ParameterCount++
		buf = append(buf, pcfparm.Bytes()...)

		// Once we know the total number of parameters, put the
		// CFH header on the front of the buffer.
		buf = append(cfh.Bytes(), buf...)

		// And put the command to the queue
		err = cmdQObj.Put(putmqmd, pmo, buf)

		if err != nil {
			return err
		}

		// Now get the response
		getmqmd := ibmmq.NewMQMD()
		gmo := ibmmq.NewMQGMO()
		gmo.Options = ibmmq.MQGMO_NO_SYNCPOINT
		gmo.Options |= ibmmq.MQGMO_FAIL_IF_QUIESCING
		gmo.Options |= ibmmq.MQGMO_WAIT
		gmo.Options |= ibmmq.MQGMO_CONVERT
		gmo.WaitInterval = 30 * 1000

		// Ought to add a loop here in case we get truncated data
		buf = make([]byte, 32768)

		datalen, err = replyQObj.Get(getmqmd, gmo, buf)
		if err == nil {
			cfh, offset := ibmmq.ReadPCFHeader(buf)
			if cfh.CompCode != ibmmq.MQCC_OK {
				return fmt.Errorf("PCF command failed with CC %d RC %d", cfh.CompCode, cfh.Reason)
			} else {
				parmAvail := true
				bytesRead := 0
				for parmAvail && cfh.CompCode != ibmmq.MQCC_FAILED {
					elem, bytesRead = ibmmq.ReadPCFParameter(buf[offset:])
					offset += bytesRead
					// Have we now reached the end of the message
					if offset >= datalen {
						parmAvail = false
					}

					switch elem.Parameter {
					case ibmmq.MQCACF_Q_NAMES:
						if len(elem.String) == 0 {
							return fmt.Errorf("No queues matching '%s' exist", pattern)
						}
						for i := 0; i < len(elem.String); i++ {
							qList = append(qList, strings.TrimSpace(elem.String[i]))
						}
					}
				}
			}
		} else {
			return err
		}
	}

	return err
}

/*
Now that we know which topics can return data, need to
create all the subscriptions.
*/
func createSubscriptions() error {
	var err error
	var sub ibmmq.MQObject

	for _, cl := range Metrics.Classes {
		for _, ty := range cl.Types {

			if strings.Contains(ty.ObjectTopic, "%s") {
				for i := 0; i < len(qList); i++ {
					topic := fmt.Sprintf(ty.ObjectTopic, qList[i])
					sub, err = subscribe(topic)
					ty.subHobj[qList[i]] = sub
				}
			} else {
				sub, err = subscribe(ty.ObjectTopic)
				ty.subHobj[QMgrMapKey] = sub
			}

			if err != nil {
				return fmt.Errorf("Error subscribing to %s: %v", ty.ObjectTopic, err)
			}
		}
	}

	return err
}

/*
ProcessPublications has to read all of the messages since the last scrape
and update the values for every relevant gauge.

Because the generation of the messages by the qmgr, and being told to
read them by the main loop, may not have identical frequencies, there may be
cases where multiple pieces of data have to be collated for the same
gauge. Conversely, there may be times when this is called but there
are no metrics to update.
*/
func ProcessPublications() error {
	var err error
	var data []byte

	var qName string
	var classidx int
	var typeidx int
	var elementidx int
	var value int64

	// Keep reading all available messages until queue is empty. Don't
	// do a GET-WAIT; just immediate removals.
	cnt := 0

	for err == nil {

		data, err = getMessage(false)

		// Most common error will be MQRC_NO_MESSAGE_AVAILABLE
		// which will end the loop.
		if err == nil {
			cnt++
			elemList, _ := ParsePCFResponse(data)
			// A typical publication contains some fixed
			// headers (qmgrName, objectName, class, type etc)
			// followed by a list of index/values.

			// This map contains those element indexes and values from each message
			values := make(map[int]int64)

			qName = ""

			for i := 0; i < len(elemList); i++ {
				switch elemList[i].Parameter {
				case ibmmq.MQCA_Q_MGR_NAME:
					_ = strings.TrimSpace(elemList[i].String[0])
				case ibmmq.MQCA_Q_NAME:
					qName = strings.TrimSpace(elemList[i].String[0])
				case ibmmq.MQCA_TOPIC_NAME:
					qName = strings.TrimSpace(elemList[i].String[0])
				case ibmmq.MQIACF_OBJECT_TYPE:
					// Will need to use this as part of the object key and
					// labelling if/when MQ starts to produce stats for other types
					// such as a topic. But for now we can ignore it.
					_ = ibmmq.MQItoString("OT", int(elemList[i].Int64Value[0]))
				case ibmmq.MQIAMO_MONITOR_CLASS:
					classidx = int(elemList[i].Int64Value[0])
				case ibmmq.MQIAMO_MONITOR_TYPE:
					typeidx = int(elemList[i].Int64Value[0])
				case ibmmq.MQIAMO64_MONITOR_INTERVAL:
					_ = elemList[i].Int64Value[0]
				case ibmmq.MQIAMO_MONITOR_FLAGS:
					_ = int(elemList[i].Int64Value[0])
				default:
					value = elemList[i].Int64Value[0]
					elementidx = int(elemList[i].Parameter)
					values[elementidx] = value
				}
			}
			// Now have all the values in this particular message
			// Have to incorporate them into any that already exist.
			//
			// Each element contains a map holding all the objects
			// touched by these messages. The map is referenced by
			// object name if it's a queue; for qmgr-level stats, the
			// map only needs to contain a single entry which I've
			// chosen to reference by "@self" which can never be a
			// real queue name.
			//
			// We have to know whether to need to add the values
			// contained from multiple publications that might
			// have arrived in the scrape interval
			// for the same resource, or whether we should just
			// overwrite with the latest. Although there are
			// several monitor Datatypes, all of them apart from
			// explicitly labelled "DELTA" are ones we should just
			// use the latest value.
			for key, newValue := range values {
				if elem, ok := Metrics.Classes[classidx].Types[typeidx].Elements[key]; ok {
					objectName := qName
					if objectName == "" {
						objectName = QMgrMapKey
					}

					if oldValue, ok := elem.Values[objectName]; ok {
						if elem.Datatype == ibmmq.MQIAMO_MONITOR_DELTA {
							value = oldValue + newValue
						} else {
							value = newValue
						}
					} else {
						value = newValue
					}
					elem.Values[objectName] = value
				}
			}
		} else {
			// err != nil
			mqreturn := err.(*ibmmq.MQReturn)

			if mqreturn.MQCC == ibmmq.MQCC_FAILED && mqreturn.MQRC != ibmmq.MQRC_NO_MSG_AVAILABLE {
				return mqreturn
			}
		}
	}
	return nil
}

/*
ParsePCFResponse message, returning the
elements. If an element represents a PCF group, that element
has the pieces of the group attached to itself. While
it is theoretically possible for groups to contain groups, MQ never
does that, so the code here does not need to recurse through multiple
levels.

Returns TRUE if this is the last response in a
set, based on the MQCFH.Control value.
*/
func ParsePCFResponse(buf []byte) ([]*ibmmq.PCFParameter, bool) {
	var elem *ibmmq.PCFParameter
	var elemList []*ibmmq.PCFParameter
	var bytesRead int

	rc := false

	// First get the MQCFH structure. This also returns
	// the number of bytes read so we know where to start
	// looking for the next element
	cfh, offset := ibmmq.ReadPCFHeader(buf)

	if cfh.CompCode != 0 {
		logp.Critical("There was a problem while reading response: CompCode: %v, Reason: %v", cfh.CompCode, cfh.Reason)
	}

	// If the command succeeded, loop through the remainder of the
	// message to decode each parameter.
	for i := 0; i < int(cfh.ParameterCount); i++ {
		// We don't know how long the parameter is, so we just
		// pass in "from here to the end" and let the parser
		// tell us how far it got.
		elem, bytesRead = ibmmq.ReadPCFParameter(buf[offset:])
		offset += bytesRead
		// Have we now reached the end of the message
		elemList = append(elemList, elem)
		if elem.Type == ibmmq.MQCFT_GROUP {
			groupElem := elem
			for j := 0; j < int(groupElem.ParameterCount); j++ {
				elem, bytesRead = ibmmq.ReadPCFParameter(buf[offset:])
				offset += bytesRead
				groupElem.GroupList = append(groupElem.GroupList, elem)
			}
		}

	}

	if cfh.Control == ibmmq.MQCFC_LAST {
		rc = true
	}
	return elemList, rc
}

/*
Need to turn the "friendly" name of each element into something
that is suitable for metric names.

Should also have consistency of units (always use seconds,
bytes etc), and organisation of the elements of the name (units last)

While we can't change the MQ-generated descriptions for its statistics,
we can reformat most of them heuristically here.
*/
func formatDescription(elem *MonElement) string {
	s := elem.Description
	s = strings.Replace(s, " ", "_", -1)
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, "-", "_", -1)

	/* Make sure we don't have multiple underscores */
	multiunder := regexp.MustCompile("__*")
	s = multiunder.ReplaceAllLiteralString(s, "_")

	/* make it all lowercase. Not essential, but looks better */
	s = strings.ToLower(s)

	/* Remove all cases of bytes, seconds, count or percentage (we add them back in later) */
	s = strings.Replace(s, "_count", "", -1)
	s = strings.Replace(s, "_bytes", "", -1)
	s = strings.Replace(s, "_byte", "", -1)
	s = strings.Replace(s, "_seconds", "", -1)
	s = strings.Replace(s, "_second", "", -1)
	s = strings.Replace(s, "_percentage", "", -1)

	// Switch round a couple of specific names
	s = strings.Replace(s, "messages_expired", "expired_messages", -1)

	// Add the unit at end
	switch elem.Datatype {
	case ibmmq.MQIAMO_MONITOR_PERCENT, ibmmq.MQIAMO_MONITOR_HUNDREDTHS:
		s = s + "_percentage"
	case ibmmq.MQIAMO_MONITOR_MB, ibmmq.MQIAMO_MONITOR_GB:
		s = s + "_bytes"
	case ibmmq.MQIAMO_MONITOR_MICROSEC:
		s = s + "_seconds"
	default:
		if strings.Contains(s, "_total") {
			/* If we specify it is a total in description put that at the end */
			s = strings.Replace(s, "_total", "", -1)
			s = s + "_total"
		} else if strings.Contains(s, "log_") {
			/* Weird case where the log datatype is not MB or GB but should be bytes */
			s = s + "_bytes"
		} else {
			s = s + "_count"
		}
	}

	return s
}

/*
ReadPatterns is called during the initial configuration step to read a file
containing object name patterns if they are not explicitly given
on the command line.
*/
func ReadPatterns(f string) (string, error) {
	var s string

	file, err := os.Open(f)
	if err != nil {
		return "", fmt.Errorf("Error Opening file %s: %v", f, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if s != "" {
			s += ","
		}
		s += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Error Reading from %s: %v", f, err)
	}

	return s, nil
}

/*
Normalise converts the value returned from MQ into the correct units
such as converting MB to bytes.
*/
func Normalise(elem *MonElement, key string, value int64) float64 {
	f := float64(value)
	// I've  seen negative numbers which are nonsense,
	// possibly 32-bit overflow or uninitialised values
	// in the qmgr. So force data to something sensible
	// just in case those were due to a bug.
	if f < 0 {
		f = 0
	}

	// Convert suitable metrics to base units
	if elem.Datatype == ibmmq.MQIAMO_MONITOR_PERCENT ||
		elem.Datatype == ibmmq.MQIAMO_MONITOR_HUNDREDTHS {
		f = f / 100
	} else if elem.Datatype == ibmmq.MQIAMO_MONITOR_MB {
		f = f * 1024 * 1024
	} else if elem.Datatype == ibmmq.MQIAMO_MONITOR_GB {
		f = f * 1024 * 1024 * 1024
	} else if elem.Datatype ==
		ibmmq.MQIAMO_MONITOR_MICROSEC {
		f = f / 1000000
	}

	return f
}
