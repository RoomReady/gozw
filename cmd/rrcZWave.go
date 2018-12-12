package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/comail/colog"
	"github.com/davecgh/go-spew/spew"
	"github.com/kardianos/service"
	"github.com/peterh/liner"
	"github.com/rmacster/gozw/application"
	doorlock "github.com/rmacster/gozw/cc/door-lock"
	switchall "github.com/rmacster/gozw/cc/switch-all"
	switchbinary "github.com/rmacster/gozw/cc/switch-binary"
	switchbinaryv2 "github.com/rmacster/gozw/cc/switch-binary-v2"
	"github.com/rmacster/gozw/frame"
	"github.com/rmacster/gozw/serial-api"
	"github.com/rmacster/gozw/session"
	"github.com/rmacster/gozw/transport"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	doIt()
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	<-time.After(time.Second * 13)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "rrcZwave",
		DisplayName: "rrcZwave",
		Description: "rrcZwave - Service that connects to Z-Wave devices.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

}

func init() {
	colog.Register()
	colog.ParseFields(true)
}

func doIt() {
	fmt.Println("Starting rrcZwave...")

	jsn, ok := loadJsonFile("zwave.json")
	if !ok {
		time.Sleep(time.Second * 3)
		panic("zwave.json not found")
	}

	if !(jsn.Exists("port")) {
		time.Sleep(time.Second * 3)
		panic("key: 'port' not found in zwave.json")
	}
	port := jsn.S("port").Data().(string)

	fmt.Println("transport layer...")
	transport, err := transport.NewSerialPortTransport(port, 115200)
	if err != nil {
		panic(err)
	}

	fmt.Println("frame layer...")
	frameLayer := frame.NewFrameLayer(transport)
	fmt.Println("session layer...")
	sessionLayer := session.NewSessionLayer(frameLayer)
	fmt.Println("api layer...")
	apiLayer := serialapi.NewLayer(sessionLayer)
	fmt.Println("app layer...")
	appLayer, err := application.NewLayer(apiLayer)
	if err != nil {
		panic(err)
	}

	defer appLayer.Shutdown()

	line := liner.NewLiner()
	defer line.Close()

	commands := strings.Join([]string{
		"(a)dd node",
		"(r)emove node",
		"(v) load command class versions for node",
		"(m) load manufacturer-specific data for node",
		"(pv) print the result of the above",
		"(nif) request node information frame from node",
		"(f)ailed node removal",
		"(p)rint network info",
		"(on) turn light on",
		"(off) turn light off",
		"(on2) turn light on - v2",
		"(off2) turn light off - v2",
		"(allon) Switch All On",
		"(alloff) Switch All Off",
		"(lock) lock door",
		"(unlock) lock door",
		"(q)uit",
	}, "\n")

	fmt.Println(commands)

	for {
		cmd, _ := line.Prompt("> ")
		switch cmd {
		case "a":
			spew.Dump(appLayer.AddNode())
		case "r":
			spew.Dump(appLayer.RemoveNode())
		case "V":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			spew.Dump(node.LoadCommandClassVersions())
		case "m":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			spew.Dump(node.LoadManufacturerInfo())

		case "pv":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			for id, cc := range node.CommandClasses {
				fmt.Printf(
					"%s: %d\n",
					id,
					cc.Version,
				)
			}

		case "on":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			spew.Dump(node.SendCommand(&switchbinary.Set{
				SwitchValue: 0x01,
			}))
		case "off":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			spew.Dump(node.SendCommand(&switchbinary.Set{
				SwitchValue: 0x00,
			}))

		case "on2":
			input, _ := line.Prompt("node id: ")
			delay, _ := line.Prompt("delay  : ")

			nodeId, _ := strconv.Atoi(input)
			del, _ := strconv.Atoi(delay)

			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			spew.Dump(node.SendCommand(&switchbinaryv2.Set{
				TargetValue: 0x01,
				Duration:    byte(del),
			}))
		case "off2":
			input, _ := line.Prompt("node id: ")
			delay, _ := line.Prompt("delay  : ")

			nodeId, _ := strconv.Atoi(input)
			del, _ := strconv.Atoi(delay)

			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			spew.Dump(node.SendCommand(&switchbinaryv2.Set{
				TargetValue: 0x00,
				Duration:    byte(del),
			}))

		case "allon": // switch all on
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}
			spew.Dump(node.SendCommand(&switchall.On{}))
			// spew.Dump(node.SendCommand(&switchall.Set{
			// 	Mode: 0x01,
			// }))

		case "alloff": // switch all off
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			// spew.Dump(node.SendCommand(&switchbinary.Set{
			// 	SwitchValue: 0x00,
			// }))
			spew.Dump(node.SendCommand(&switchall.Off{}))

		case "unlock":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			spew.Dump(node.SendCommand(&doorlock.OperationSet{
				DoorLockMode: 0x00,
			}))

		case "lock":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, err := appLayer.Node(byte(nodeId))
			if err != nil {
				spew.Dump(err)
				continue
			}

			spew.Dump(node.SendCommand(&doorlock.OperationSet{
				DoorLockMode: 0xFF,
			}))

		case "nif":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			node, _ := appLayer.Node(byte(nodeId))
			spew.Dump(node.RequestNodeInformationFrame())
		case "f":
			input, _ := line.Prompt("node id: ")
			nodeId, _ := strconv.Atoi(input)
			spew.Dump(appLayer.RemoveFailedNode(byte(nodeId)))
		case "p":
			fmt.Printf("Home ID: 0x%x; Node ID: %d\n", appLayer.Controller.HomeID, appLayer.Controller.NodeID)
			fmt.Println("API Version:", appLayer.Controller.APIVersion)
			fmt.Println("Library:", appLayer.Controller.APILibraryType)
			fmt.Println("Version:", appLayer.Controller.Version)
			fmt.Println("API Type:", appLayer.Controller.APIType)
			fmt.Println("Is Primary Controller:", appLayer.Controller.IsPrimaryController)
			fmt.Println("Node count:", len(appLayer.Nodes()))

			for _, node := range appLayer.Nodes() {
				fmt.Println(node.String())
			}
		case "q":
			return
		default:
			fmt.Printf("invalid selection\n\n")
			fmt.Println(commands)
		}
	}
}
