package ibmmqlib

import (
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

type RequestObject struct {
	Commands []struct {
		Cmd    string
		Params map[string]interface{}
	}
}

/**
Config contains all configuration objects
*/
type Config struct {
	QueueManager       string           `config:"bindingQueueManager"`
	RemoteQueueManager []string         `config:"targetQueueManager"`
	QMgrStat           bool             `config:"queueManagerStatus"`
	PubSub             bool             `config:"pubSub"`
	ConnectionConfig   ConnectionConfig `config:"cc"`
}

type ConnectionConfig struct {
	ClientMode bool   `config:"clientMode"`
	MqServer   string `config:"mqServer"`
	UserID     string `config:"user"`
	Password   string `config:"password"`
}

var (
	//DefaultConfig contains the default configuration for this module
	DefaultConfig = Config{
		PubSub:             false,
		QMgrStat:           true,
		RemoteQueueManager: []string{""},
		ConnectionConfig: ConnectionConfig{
			ClientMode: false,
			UserID:     "",
			Password:   "",
		},
	}
)

var (
	first      = true
	errorCount = 0
)

func connectPubSub(qmgrName string, queuePattern string, cc *ConnectionConfig) error {
	var err error

	// Connect to MQ
	logp.Info("Connect to QM %v start", qmgrName)
	err = InitConnection(qmgrName, "SYSTEM.DEFAULT.MODEL.QUEUE", cc)
	if err == nil {
		logp.Info("Connected to queue manager %v", qmgrName)
	}

	logp.Info("Connect to QM done")

	// What metrics can the queue manager provide? Find out, and
	// subscribe.
	if err == nil {
		logp.Info("DiscoverAndSubscribe start")
		err = DiscoverAndSubscribe(queuePattern, true, "")
	}
	logp.Info("DiscoverAndSubscribe done")

	return err
}

func collectPubSub(qmgrName string, eventType string) {
	// #####Code for collecting the MQ metrics
	// Clear out everything we know so far. In particular, replace
	// the map of values for each object so the collection starts
	// clean.
	logp.Info("Start MQ Metric collection")

	for _, cl := range Metrics.Classes {
		//logp.Info("Define class %v", cl.Name)
		for _, ty := range cl.Types {
			//logp.Info("Define type %v", ty.ObjectTopic)
			for _, elem := range ty.Elements {
				//logp.Info("Define elem %v", elem.Values)
				//logp.Info("test: ",elem.Values)
				elem.Values = make(map[string]int64)
			}
		}
	}

	//if (cl.length > 0) {
	// Process all the publications that have arrived
	logp.Info("ProcessPublications start")
	ProcessPublications()
	logp.Info("ProcessPublications done")

	if !first {

		for _, cl := range Metrics.Classes {
			for _, ty := range cl.Types {
				event := beat.Event{
					Timestamp: time.Now(),
					Fields: common.MapStr{
						"metrictype":  cl.Name,
						"objecttopic": ty.ObjectTopic,
						"type":        eventType,
						"qmgr":        qmgrName,
					},
				}
				for _, elem := range ty.Elements {
					for key, value := range elem.Values {
						f := Normalise(elem, key, value)

						//Add some metadata information based on type
						if key != QMgrMapKey {
							event.Fields["queue"] = key
							event.Fields["metricset"] = "queue"
						} else {
							event.Fields["metricset"] = "queuemanager"
						}
						event.Fields[elem.MetricName] = float32(f)
					}
				}
			}
		}

		//}
	}
	// ###### END Code for collecting the MQ metrics

}

func connectLegacyMode(qmgrName string, cc ConnectionConfig) error {

	logp.Info("Connect in legacy mode")

	err = InitConnection(qmgrName, "SYSTEM.DEFAULT.MODEL.QUEUE", &cc)

	if err != nil {
		return err
	}

	logp.Info("Connection successfull")
	return err

	return nil
}

func createEvents(eventType string, qmgrName string, responseObj map[string]*Response) []beat.Event {
	var events []beat.Event

	for id, elem := range responseObj {
		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":         eventType,
				"qmgr":         qmgrName,
				"metricset":    elem.Metricset,
				"metrictype":   elem.Metrictype,
				"targetObject": elem.TargetObject,
			},
		}
		for key, value := range responseObj[id].Values {
			event.Fields[key] = value
		}

		events = append(events, event)
	}
	return events
}

func mergeEventsWithResponseObj(events []beat.Event, responseObj map[string]*Response) []beat.Event {

	for id, _ := range responseObj {
		for _, event := range events {
			if id == event.Fields["targetObject"] {
				for key, value := range responseObj[id].Values {
					event.Fields[key] = value
				}
			}
		}
	}
	return events
}

func generateConnectedObjectsField(events []beat.Event) []beat.Event {
	for _, event := range events {
		var connections []string
		connections = make([]string, 0)
		connections = append(connections, event.Fields["qmgr"].(string))
		connections = append(connections, event.Fields["targetObject"].(string))
		if event.Fields["mqcach_xmit_q_name"] != nil {
			connections = append(connections, event.Fields["mqcach_xmit_q_name"].(string))
		}
		if event.Fields["mqca_remote_q_mgr_name"] != nil {
			connections = append(connections, event.Fields["mqca_remote_q_mgr_name"].(string))
		}
		if event.Fields["mqca_remote_q_name"] != nil {
			connections = append(connections, event.Fields["mqca_remote_q_name"].(string))
		}
		//Remove whitespaces
		for _, connection := range connections {
			connection = strings.TrimSpace(connection)
		}
		event.Fields["Conntections"] = connections
	}
	return events
}

/**
CollectQmgrMetricset is responsible for collecting data from Queue Manager
*/
func CollectQmgrMetricset(eventType string, qmgrName string, cc ConnectionConfig) ([]beat.Event, error) {
	//Collect queue statistics
	var err error
	var events []beat.Event

	connectLegacyMode(qmgrName, cc)

	qMgrMetadata, err := getQManagerMetadata(qmgrName)
	if err != nil {
		return nil, err
	}

	qMgrStatus, err := getQManagerStatus(qmgrName)
	if err != nil {
		return nil, err
	}
	tmpEvents := createEvents(eventType, qmgrName, qMgrMetadata)
	events = append(events, mergeEventsWithResponseObj(tmpEvents, qMgrStatus)...)

	return events, err
}
