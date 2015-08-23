package application

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"github.com/helioslabs/gozw/zwave/command-class"
	"github.com/helioslabs/gozw/zwave/command-class/alarm"
	"github.com/helioslabs/gozw/zwave/command-class/association"
	"github.com/helioslabs/gozw/zwave/command-class/battery"
	"github.com/helioslabs/gozw/zwave/command-class/door-lock"
	"github.com/helioslabs/gozw/zwave/command-class/manufacturer-specific"
	"github.com/helioslabs/gozw/zwave/command-class/security"
	"github.com/helioslabs/gozw/zwave/command-class/thermostat-mode"
	"github.com/helioslabs/gozw/zwave/command-class/thermostat-operating-state"
	"github.com/helioslabs/gozw/zwave/command-class/thermostat-setpoint"
	"github.com/helioslabs/gozw/zwave/command-class/user-code"
	"github.com/helioslabs/gozw/zwave/command-class/version"
	"github.com/helioslabs/gozw/zwave/protocol"
	"github.com/helioslabs/gozw/zwave/serial-api"
	"github.com/helioslabs/proto"
)

// CommandClassSupport defines a node's support level for a command class
type CommandClassSupport int

const (
	// CommandClassNotSupported indicates that the command class is not supported
	// at all
	CommandClassNotSupported CommandClassSupport = iota

	// CommandClassSupportedInsecure indicates that the command class is supported
	// regardless of the security environment. The node may or may not be incldued
	// securely, and the command may or may not be sent securely.
	CommandClassSupportedInsecure

	// CommandClassSupportedSecure indicates that the command class is only
	// supported through Z-Wave security. The node *MUST* be securely included
	// in order to use this command class.
	CommandClassSupportedSecure
)

// Node is an in-memory representation of a Z-Wave node
type Node struct {
	NodeID byte

	Capability          byte
	BasicDeviceClass    byte
	GenericDeviceClass  byte
	SpecificDeviceClass byte

	Failing bool

	SupportedCommandClasses        map[commandclass.ID]bool
	SecureSupportedCommandClasses  map[commandclass.ID]bool
	SecureControlledCommandClasses map[commandclass.ID]bool

	CommandClassVersions map[commandclass.ID]byte

	NetworkKeySent bool

	ManufacturerID uint16
	ProductTypeID  uint16
	ProductID      uint16

	application          *Layer
	receivedUpdate       chan bool
	receivedSecurityInfo chan bool
}

func NewNode(application *Layer, nodeID byte) (*Node, error) {
	node := &Node{
		NodeID: nodeID,

		SupportedCommandClasses:        map[commandclass.ID]bool{},
		SecureSupportedCommandClasses:  map[commandclass.ID]bool{},
		SecureControlledCommandClasses: map[commandclass.ID]bool{},

		CommandClassVersions: map[commandclass.ID]byte{},

		application:          application,
		receivedUpdate:       make(chan bool),
		receivedSecurityInfo: make(chan bool),
	}

	err := node.loadFromDb()
	if err != nil {
		initErr := node.initialize()
		if initErr != nil {
			return nil, initErr
		}

		node.saveToDb()
	}

	return node, nil
}

func (n *Node) loadFromDb() error {
	var data []byte
	err := n.application.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		data = bucket.Get([]byte{n.NodeID})

		if len(data) == 0 {
			return errors.New("Node not found")
		}

		return nil
	})

	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(data, n)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) initialize() error {
	nodeInfo, err := n.application.serialAPI.GetNodeProtocolInfo(n.NodeID)
	if err != nil {
		fmt.Println(err)
	} else {
		n.setFromNodeProtocolInfo(nodeInfo)
	}

	if n.NodeID == 1 {
		// self is never failing
		n.Failing = false
	} else {
		failing, err := n.application.serialAPI.IsFailedNode(n.NodeID)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		n.Failing = failing
	}

	return n.saveToDb()
}

func (n *Node) saveToDb() error {
	data, err := msgpack.Marshal(n)
	if err != nil {
		return err
	}

	return n.application.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		return bucket.Put([]byte{n.NodeID}, data)
	})
}

func (n *Node) IsSecure() bool {
	_, found := n.SupportedCommandClasses[commandclass.Security]
	return found
}

func (n *Node) IsListening() bool {
	return n.Capability&0x80 == 0x80
}

func (n *Node) GetBasicDeviceClassName() string {
	return protocol.GetBasicDeviceTypeName(n.BasicDeviceClass)
}

func (n *Node) GetGenericDeviceClassName() string {
	return protocol.GetGenericDeviceTypeName(n.GenericDeviceClass)
}

func (n *Node) GetSpecificDeviceClassName() string {
	return protocol.GetSpecificDeviceTypeName(n.GenericDeviceClass, n.SpecificDeviceClass)
}

func (n *Node) SendCommand(command commandclass.Command) error {
	commandClass := commandclass.ID(command.CommandClassID())
	supportType := n.SupportsCommandClass(commandClass)

	if commandClass == commandclass.Security {
		switch command.(type) {
		case *security.CommandsSupportedGet, *security.CommandsSupportedReport:
			supportType = CommandClassSupportedSecure
		}
	}

	switch supportType {
	case CommandClassSupportedSecure:
		return n.application.SendDataSecure(n.NodeID, command)
	case CommandClassSupportedInsecure:
		return n.application.SendData(n.NodeID, command)
	case CommandClassNotSupported:
		return errors.New("Command class not supported")
	default:
		return errors.New("Command class not supported")
	}
}

func (n *Node) SupportsCommandClass(commandClass commandclass.ID) CommandClassSupport {
	if supported, ok := n.SupportedCommandClasses[commandClass]; ok && supported {
		return CommandClassSupportedInsecure
	}

	if supported, ok := n.SecureSupportedCommandClasses[commandClass]; ok && supported {
		return CommandClassSupportedSecure
	}

	return CommandClassNotSupported
}

func (n *Node) AddAssociation(groupID byte, nodeIDs ...byte) error {
	// sort of an arbitrary limit for now, but I'm not sure what it should be
	if len(nodeIDs) > 20 {
		return errors.New("Too many associated nodes")
	}

	fmt.Println("Associating")

	return n.SendCommand(&association.Set{
		GroupingIdentifier: groupID,
		NodeId:             nodeIDs,
	})
}

func (n *Node) RequestSupportedSecurityCommands() error {
	return n.application.SendDataSecure(n.NodeID, &security.CommandsSupportedGet{})
}

func (n *Node) RequestNodeInformationFrame() error {
	return n.application.serialAPI.RequestNodeInfo(n.NodeID)
}

func (n *Node) LoadCommandClassVersions() error {
	for cc := range n.SupportedCommandClasses {
		time.Sleep(1 * time.Second)

		cmd := &version.CommandClassGet{RequestedCommandClass: byte(cc)}
		if err := n.application.SendData(n.NodeID, cmd); err != nil {
			return err
		}
	}

	for cc := range n.SecureSupportedCommandClasses {
		time.Sleep(1 * time.Second)

		cmd := &version.CommandClassGet{RequestedCommandClass: byte(cc)}
		if err := n.application.SendDataSecure(n.NodeID, cmd); err != nil {
			return err
		}
	}

	return nil
}

func (n *Node) LoadManufacturerInfo() error {
	return n.SendCommand(&manufacturerspecific.Get{})
}

func (n *Node) emitNodeEvent(event interface{}) {
	n.application.EventBus.Publish("event", proto.Event{
		Payload: proto.NodeEvent{
			NodeId: n.NodeID,
			Event:  event,
		},
	})
}

func (n *Node) receiveControllerUpdate(update serialapi.ControllerUpdate) {
	select {
	case n.receivedUpdate <- true:
	default:
	}

	n.setFromApplicationControllerUpdate(update)
	n.saveToDb()
}

// func (n *Node) sendNoOp() {
// 	n.manager.session.SendData(n.NodeId, []byte{
// 		commandclass.CommandClassNoOperation,
// 	})
// }

func (n *Node) setFromAddNodeCallback(nodeInfo *serialapi.AddRemoveNodeCallback) {
	n.NodeID = nodeInfo.Source
	n.BasicDeviceClass = nodeInfo.Basic
	n.GenericDeviceClass = nodeInfo.Generic
	n.SpecificDeviceClass = nodeInfo.Specific

	for _, cc := range nodeInfo.CommandClasses {
		n.SupportedCommandClasses[commandclass.ID(cc)] = true
	}

	n.saveToDb()
}

func (n *Node) setFromApplicationControllerUpdate(nodeInfo serialapi.ControllerUpdate) {
	n.BasicDeviceClass = nodeInfo.Basic
	n.GenericDeviceClass = nodeInfo.Generic
	n.SpecificDeviceClass = nodeInfo.Specific

	for _, cc := range nodeInfo.CommandClasses {
		n.SupportedCommandClasses[commandclass.ID(cc)] = true
	}

	n.saveToDb()
}

func (n *Node) setFromNodeProtocolInfo(nodeInfo *serialapi.NodeProtocolInfo) {
	n.Capability = nodeInfo.Capability
	n.BasicDeviceClass = nodeInfo.BasicDeviceClass
	n.GenericDeviceClass = nodeInfo.GenericDeviceClass
	n.SpecificDeviceClass = nodeInfo.SpecificDeviceClass

	n.saveToDb()
}

func (n *Node) receiveSecurityCommandsSupportedReport(cc security.CommandsSupportedReport) {
	for _, cc := range cc.CommandClassSupport {
		n.SecureSupportedCommandClasses[commandclass.ID(cc)] = true
	}

	for _, cc := range cc.CommandClassControl {
		n.SecureControlledCommandClasses[commandclass.ID(cc)] = true
	}

	select {
	case n.receivedSecurityInfo <- true:
	default:
	}

	n.saveToDb()
}

func (n *Node) receiveApplicationCommand(cmd serialapi.ApplicationCommand) {
	cc := commandclass.ID(cmd.CommandData[0])
	ver, ok := n.CommandClassVersions[cc]

	if !ok {
		if cc == commandclass.Version || cc == commandclass.Security {
			ver = 1
		} else {
			fmt.Printf("error: no version loaded for %s\n", cc)
		}
	}

	command, err := commandclass.Parse(ver, cmd.CommandData)
	if err != nil {
		fmt.Println("error parsing command class", err)
		return
	}

	switch command.(type) {

	case battery.Report:
		if cmd.CommandData[2] == 0xFF {
			fmt.Printf("Node %d: low battery alert\n", n.NodeID)
		} else {
			fmt.Printf("Node %d: battery level is %d\n", n.NodeID, command.(battery.Report))
		}

	case security.CommandsSupportedReport:
		fmt.Println("security commands supported report")
		n.receiveSecurityCommandsSupportedReport(command.(security.CommandsSupportedReport))
		fmt.Println(n.GetSupportedSecureCommandClassStrings())

	case alarm.Report:
		spew.Dump(command.(alarm.Report))

	case usercode.Report:
		spew.Dump(command.(usercode.Report))

	case doorlock.OperationReport:
		spew.Dump(command.(doorlock.OperationReport))

	case thermostatmode.Report:
		spew.Dump(command.(thermostatmode.Report))

	case thermostatoperatingstate.Report:
		spew.Dump(command.(thermostatoperatingstate.Report))

	case thermostatsetpoint.Report:
		spew.Dump(command.(thermostatsetpoint.Report))

	case version.CommandClassReport:
		spew.Dump(command.(version.CommandClassReport))
		report := command.(version.CommandClassReport)
		n.CommandClassVersions[commandclass.ID(report.RequestedCommandClass)] = report.CommandClassVersion
		n.saveToDb()

	case manufacturerspecific.Report:
		spew.Dump(command.(manufacturerspecific.Report))
		// mfgInfo := commandclass.ParseManufacturerSpecificReport(cmd.CommandData)
		// n.ManufacturerID = mfgInfo.ManufacturerID
		// n.ProductTypeID = mfgInfo.ProductTypeID
		// n.ProductID = mfgInfo.ProductID
	default:
		spew.Dump(command)
	}
}

func (n *Node) String() string {
	str := fmt.Sprintf("Node %d: \n", n.NodeID)
	str += fmt.Sprintf("  Failing? %t\n", n.Failing)
	str += fmt.Sprintf("  Is listening? %t\n", n.IsListening())
	str += fmt.Sprintf("  Is secure? %t\n", n.IsSecure())
	str += fmt.Sprintf("  Basic device class: %s\n", n.GetBasicDeviceClassName())
	str += fmt.Sprintf("  Generic device class: %s\n", n.GetGenericDeviceClassName())
	str += fmt.Sprintf("  Specific device class: %s\n", n.GetSpecificDeviceClassName())
	str += fmt.Sprintf("  Manufacturer ID: %#x\n", n.ManufacturerID)
	str += fmt.Sprintf("  Product Type ID: %#x\n", n.ProductTypeID)
	str += fmt.Sprintf("  Product ID: %#x\n", n.ProductID)
	str += fmt.Sprintf("  Supported command classes:\n")
	for _, cc := range n.GetSupportedCommandClassStrings() {
		str += fmt.Sprintf("    - %s\n", cc)
	}

	if len(n.SecureSupportedCommandClasses) > 0 {
		secureCommands := commandClassSetToStrings(n.SecureSupportedCommandClasses)
		str += fmt.Sprintf("  Supported command classes (secure):\n")
		for _, cc := range secureCommands {
			str += fmt.Sprintf("    - %s\n", cc)
		}
	}

	if len(n.SecureControlledCommandClasses) > 0 {
		secureCommands := commandClassSetToStrings(n.SecureControlledCommandClasses)
		str += fmt.Sprintf("  Controlled command classes (secure):\n")
		for _, cc := range secureCommands {
			str += fmt.Sprintf("    - %s\n", cc)
		}
	}

	return str
}

func (n *Node) GetSupportedCommandClassStrings() []string {
	strings := commandClassSetToStrings(n.SupportedCommandClasses)
	if len(strings) == 0 {
		return []string{
			"None (probably not loaded; need to request a NIF)",
		}
	}

	return strings
}

func (n *Node) GetSupportedSecureCommandClassStrings() []string {
	strings := commandClassSetToStrings(n.SecureSupportedCommandClasses)
	return strings
}

func commandClassSetToStrings(commandClasses map[commandclass.ID]bool) []string {
	if len(commandClasses) == 0 {
		return []string{}
	}

	ccStrings := []string{}

	for cc := range commandClasses {
		ccStrings = append(ccStrings, cc.String())
	}

	return ccStrings
}
