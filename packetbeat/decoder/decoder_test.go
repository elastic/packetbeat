// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build !integration

package decoder

import (
	"strings"
	"testing"

	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/packetbeat/flows"
	"github.com/elastic/beats/v7/packetbeat/protos"

	"github.com/stretchr/testify/assert"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
)

type TestIcmp4Processor struct {
	icmp4 *layers.ICMPv4
	pkt   *protos.Packet
}

func (l *TestIcmp4Processor) ProcessICMPv4(id *flows.FlowID, icmp4 *layers.ICMPv4, pkt *protos.Packet) {
	l.icmp4 = icmp4
	l.pkt = pkt
}

type TestIcmp6Processor struct {
	icmp6 *layers.ICMPv6
	pkt   *protos.Packet
}

func (l *TestIcmp6Processor) ProcessICMPv6(id *flows.FlowID, icmp6 *layers.ICMPv6, pkt *protos.Packet) {
	l.icmp6 = icmp6
	l.pkt = pkt
}

type TestTCPProcessor struct {
	tcphdr *layers.TCP
	pkt    *protos.Packet
}

func (l *TestTCPProcessor) Process(id *flows.FlowID, tcphdr *layers.TCP, pkt *protos.Packet) {
	l.tcphdr = tcphdr
	l.pkt = pkt
}

type TestUDPProcessor struct {
	pkt *protos.Packet
}

func (l *TestUDPProcessor) Process(id *flows.FlowID, pkt *protos.Packet) {
	l.pkt = pkt
}

// 172.16.16.164:1108 172.16.16.139:53 DNS 87  Standard query 0x0007  AXFR contoso.local
var ipv4TcpDNS = []byte{
	0x00, 0x0c, 0x29, 0xce, 0xd1, 0x9e, 0x00, 0x0c, 0x29, 0x7e, 0xec, 0xa4, 0x08, 0x00, 0x45, 0x00,
	0x00, 0x49, 0x46, 0x54, 0x40, 0x00, 0x80, 0x06, 0x3b, 0x0b, 0xac, 0x10, 0x10, 0xa4, 0xac, 0x10,
	0x10, 0x8b, 0x04, 0x54, 0x00, 0x35, 0x5d, 0x9f, 0x0c, 0x90, 0x1a, 0xef, 0x6f, 0x43, 0x50, 0x18,
	0xfa, 0xf0, 0xbc, 0x3d, 0x00, 0x00, 0x00, 0x07, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x6f, 0x73, 0x6f, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x00, 0x00, 0xfc, 0x00, 0x01, 0x4d, 0x53,
}

// Test that DecodePacket decodes and IPv4/TCP packet and invokes the TCP processor.
func TestDecodePacketData_ipv4Tcp(t *testing.T) {
	logp.TestingSetup(logp.WithSelectors("decoder"))

	p := gopacket.NewPacket(ipv4TcpDNS, layers.LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	d, tcp, _ := newTestDecoder(t)
	d.OnPacket(p.Data(), &p.Metadata().CaptureInfo)

	assert.NotNil(t, tcp.pkt, "TCP packet not received")
	assert.Equal(t, "172.16.16.164", tcp.pkt.Tuple.SrcIP.String())
	assert.Equal(t, uint16(1108), tcp.pkt.Tuple.SrcPort)
	assert.Equal(t, "172.16.16.139", tcp.pkt.Tuple.DstIP.String())
	assert.Equal(t, uint16(53), tcp.pkt.Tuple.DstPort)
	assert.NotEqual(t, -1, strings.Index(string(p.Data()), string(tcp.pkt.Payload)))
}

// 192.168.170.8:32795 192.168.170.20:53  DNS 74  Standard query 0x75c0  A www.netbsd.org
var ipv4UdpDNS = []byte{
	0x00, 0xc0, 0x9f, 0x32, 0x41, 0x8c, 0x00, 0xe0, 0x18, 0xb1, 0x0c, 0xad, 0x08, 0x00, 0x45, 0x00,
	0x00, 0x3c, 0x00, 0x00, 0x40, 0x00, 0x40, 0x11, 0x65, 0x43, 0xc0, 0xa8, 0xaa, 0x08, 0xc0, 0xa8,
	0xaa, 0x14, 0x80, 0x1b, 0x00, 0x35, 0x00, 0x28, 0xaf, 0x61, 0x75, 0xc0, 0x01, 0x00, 0x00, 0x01,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x77, 0x77, 0x77, 0x06, 0x6e, 0x65, 0x74, 0x62, 0x73,
	0x64, 0x03, 0x6f, 0x72, 0x67, 0x00, 0x00, 0x01, 0x00, 0x01,
}

// Test that DecodePacket decodes and IPv4/UDP packet and invokes the UDP processor.
func TestDecodePacketData_ipv4Udp(t *testing.T) {
	p := gopacket.NewPacket(ipv4UdpDNS, layers.LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	d, _, udp := newTestDecoder(t)
	d.OnPacket(p.Data(), &p.Metadata().CaptureInfo)

	assert.NotNil(t, udp.pkt, "UDP packet not received")
	assert.Equal(t, "192.168.170.8", udp.pkt.Tuple.SrcIP.String())
	assert.Equal(t, uint16(32795), udp.pkt.Tuple.SrcPort)
	assert.Equal(t, "192.168.170.20", udp.pkt.Tuple.DstIP.String())
	assert.Equal(t, uint16(53), udp.pkt.Tuple.DstPort)
	assert.NotEqual(t, -1, strings.Index(string(p.Data()), string(udp.pkt.Payload)))
}

// IP6 2001:6f8:102d::2d0:9ff:fee3:e8de.59201 > 2001:6f8:900:7c0::2.80
var ipv6TcpHTTPGet = []byte{
	0x00, 0x11, 0x25, 0x82, 0x95, 0xb5, 0x00, 0xd0, 0x09, 0xe3, 0xe8, 0xde, 0x86, 0xdd, 0x60, 0x00,
	0x00, 0x00, 0x01, 0x04, 0x06, 0x40, 0x20, 0x01, 0x06, 0xf8, 0x10, 0x2d, 0x00, 0x00, 0x02, 0xd0,
	0x09, 0xff, 0xfe, 0xe3, 0xe8, 0xde, 0x20, 0x01, 0x06, 0xf8, 0x09, 0x00, 0x07, 0xc0, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xe7, 0x41, 0x00, 0x50, 0xab, 0xdc, 0xd6, 0x61, 0x01, 0x4a,
	0x73, 0x9f, 0x50, 0x18, 0x16, 0x80, 0xf4, 0x48, 0x00, 0x00, 0x47, 0x45, 0x54, 0x20, 0x2f, 0x20,
	0x48, 0x54, 0x54, 0x50, 0x2f, 0x31, 0x2e, 0x30, 0x0d, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x3a, 0x20,
	0x63, 0x6c, 0x2d, 0x31, 0x39, 0x38, 0x35, 0x2e, 0x68, 0x61, 0x6d, 0x2d, 0x30, 0x31, 0x2e, 0x64,
	0x65, 0x2e, 0x73, 0x69, 0x78, 0x78, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x0d, 0x0a, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x3a, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x2c, 0x20,
	0x74, 0x65, 0x78, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x2c, 0x20, 0x74, 0x65, 0x78, 0x74,
	0x2f, 0x63, 0x73, 0x73, 0x2c, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x73, 0x67, 0x6d, 0x6c, 0x2c,
	0x20, 0x2a, 0x2f, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x30, 0x31, 0x0d, 0x0a, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x2d, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x67, 0x7a,
	0x69, 0x70, 0x2c, 0x20, 0x62, 0x7a, 0x69, 0x70, 0x32, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x2d, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x65, 0x6e, 0x0d, 0x0a,
	0x55, 0x73, 0x65, 0x72, 0x2d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x4c, 0x79, 0x6e, 0x78,
	0x2f, 0x32, 0x2e, 0x38, 0x2e, 0x36, 0x72, 0x65, 0x6c, 0x2e, 0x32, 0x20, 0x6c, 0x69, 0x62, 0x77,
	0x77, 0x77, 0x2d, 0x46, 0x4d, 0x2f, 0x32, 0x2e, 0x31, 0x34, 0x20, 0x53, 0x53, 0x4c, 0x2d, 0x4d,
	0x4d, 0x2f, 0x31, 0x2e, 0x34, 0x2e, 0x31, 0x20, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x53, 0x4c, 0x2f,
	0x30, 0x2e, 0x39, 0x2e, 0x38, 0x62, 0x0d, 0x0a, 0x0d, 0x0a,
}

// Test that DecodePacket decodes and IPv6/TCP packet and invokes the TCP processor.
func TestDecodePacketData_ipv6Tcp(t *testing.T) {
	p := gopacket.NewPacket(ipv6TcpHTTPGet, layers.LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet: ", p.ErrorLayer().Error())
	}
	d, tcp, _ := newTestDecoder(t)
	d.OnPacket(p.Data(), &p.Metadata().CaptureInfo)

	assert.NotNil(t, tcp.pkt, "TCP packet not received")
	assert.Equal(t, "2001:6f8:102d:0:2d0:9ff:fee3:e8de", tcp.pkt.Tuple.SrcIP.String())
	assert.Equal(t, uint16(59201), tcp.pkt.Tuple.SrcPort)
	assert.Equal(t, "2001:6f8:900:7c0::2", tcp.pkt.Tuple.DstIP.String())
	assert.Equal(t, uint16(80), tcp.pkt.Tuple.DstPort)
	assert.NotEqual(t, -1, strings.Index(string(p.Data()), string(tcp.pkt.Payload)))
}

// 3ffe:507:0:1:200:86ff:fe05:80da.2415 > 3ffe:501:4819::42.53
var ipv6UdpDNS = []byte{
	0x00, 0x60, 0x97, 0x07, 0x69, 0xea, 0x00, 0x00, 0x86, 0x05, 0x80, 0xda, 0x86, 0xdd, 0x60, 0x00,
	0x00, 0x00, 0x00, 0x61, 0x11, 0x40, 0x3f, 0xfe, 0x05, 0x07, 0x00, 0x00, 0x00, 0x01, 0x02, 0x00,
	0x86, 0xff, 0xfe, 0x05, 0x80, 0xda, 0x3f, 0xfe, 0x05, 0x01, 0x48, 0x19, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0x09, 0x6f, 0x00, 0x35, 0x00, 0x61, 0xa3, 0x35, 0x5c, 0x78,
	0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x61, 0x01, 0x65, 0x01, 0x39,
	0x01, 0x36, 0x01, 0x37, 0x01, 0x30, 0x01, 0x65, 0x01, 0x66, 0x01, 0x66, 0x01, 0x66, 0x01, 0x37,
	0x01, 0x39, 0x01, 0x30, 0x01, 0x36, 0x01, 0x32, 0x01, 0x30, 0x01, 0x31, 0x01, 0x30, 0x01, 0x30,
	0x01, 0x30, 0x01, 0x30, 0x01, 0x30, 0x01, 0x30, 0x01, 0x30, 0x01, 0x37, 0x01, 0x30, 0x01, 0x35,
	0x01, 0x30, 0x01, 0x65, 0x01, 0x66, 0x01, 0x66, 0x01, 0x33, 0x03, 0x69, 0x70, 0x36, 0x03, 0x69,
	0x6e, 0x74, 0x00, 0x00, 0x0c, 0x00, 0x01,
}

// Test that DecodePacket decodes and IPv6/UDP packet and invokes the UDP processor.
func TestDecodePacketData_ipv6Udp(t *testing.T) {
	p := gopacket.NewPacket(ipv6UdpDNS, layers.LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	d, _, udp := newTestDecoder(t)
	d.OnPacket(p.Data(), &p.Metadata().CaptureInfo)

	assert.NotNil(t, udp.pkt, "UDP packet not received")
	assert.Equal(t, "3ffe:507:0:1:200:86ff:fe05:80da", udp.pkt.Tuple.SrcIP.String())
	assert.Equal(t, uint16(2415), udp.pkt.Tuple.SrcPort)
	assert.Equal(t, "3ffe:501:4819::42", udp.pkt.Tuple.DstIP.String())
	assert.Equal(t, uint16(53), udp.pkt.Tuple.DstPort)
	assert.NotEqual(t, -1, strings.Index(string(p.Data()), string(udp.pkt.Payload)))
}

// Creates a new TestDecoder that handles ethernet packets.
func newTestDecoder(t *testing.T) (*Decoder, *TestTCPProcessor, *TestUDPProcessor) {
	icmp4Layer := &TestIcmp4Processor{}
	icmp6Layer := &TestIcmp6Processor{}
	tcpLayer := &TestTCPProcessor{}
	udpLayer := &TestUDPProcessor{}
	d, err := New(nil, layers.LinkTypeEthernet, icmp4Layer, icmp6Layer, tcpLayer, udpLayer)
	if err != nil {
		t.Fatalf("Error creating decoder %v", err)
	}
	return d, tcpLayer, udpLayer
}
