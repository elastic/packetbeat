// Copyright 2012, Google, Inc. All rights reserved.
// Copyright 2009-2011 Andreas Krennmair. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"github.com/tsg/gopacket"
	"net"
	"reflect"
	"testing"
)

// testPacketICMPv6 is the packet:
//   10:48:30.088384 IP6 2620:0:1005:0:26be:5ff:fe27:b17 > fe80::21f:caff:feb3:7640: ICMP6, neighbor advertisement, tgt is 2620:0:1005:0:26be:5ff:fe27:b17, length 24
//      0x0000:  001f cab3 7640 24be 0527 0b17 86dd 6000  ....v@$..'....`.
//      0x0010:  0000 0018 3aff 2620 0000 1005 0000 26be  ....:.&.......&.
//      0x0020:  05ff fe27 0b17 fe80 0000 0000 0000 021f  ...'............
//      0x0030:  caff feb3 7640 8800 1ed6 4000 0000 2620  ....v@....@...&.
//      0x0040:  0000 1005 0000 26be 05ff fe27 0b17       ......&....'..
var testPacketICMPv6 = []byte{
	0x00, 0x1f, 0xca, 0xb3, 0x76, 0x40, 0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x86, 0xdd, 0x60, 0x00,
	0x00, 0x00, 0x00, 0x18, 0x3a, 0xff, 0x26, 0x20, 0x00, 0x00, 0x10, 0x05, 0x00, 0x00, 0x26, 0xbe,
	0x05, 0xff, 0xfe, 0x27, 0x0b, 0x17, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x1f,
	0xca, 0xff, 0xfe, 0xb3, 0x76, 0x40, 0x88, 0x00, 0x1e, 0xd6, 0x40, 0x00, 0x00, 0x00, 0x26, 0x20,
	0x00, 0x00, 0x10, 0x05, 0x00, 0x00, 0x26, 0xbe, 0x05, 0xff, 0xfe, 0x27, 0x0b, 0x17,
}

func TestPacketICMPv6(t *testing.T) {
	p := gopacket.NewPacket(testPacketICMPv6, LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv6, LayerTypeICMPv6, gopacket.LayerTypePayload}, t)
	if got, ok := p.Layer(LayerTypeIPv6).(*IPv6); ok {
		want := &IPv6{
			BaseLayer: BaseLayer{
				Contents: []byte{0x60, 0x0, 0x0, 0x0, 0x0, 0x18,
					0x3a, 0xff, 0x26, 0x20, 0x0, 0x0, 0x10, 0x5, 0x0, 0x0, 0x26, 0xbe, 0x5,
					0xff, 0xfe, 0x27, 0xb, 0x17, 0xfe, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x2, 0x1f, 0xca, 0xff, 0xfe, 0xb3, 0x76, 0x40},
				Payload: []byte{0x88, 0x0, 0x1e, 0xd6, 0x40, 0x0, 0x0, 0x0, 0x26, 0x20,
					0x0, 0x0, 0x10, 0x5, 0x0, 0x0, 0x26, 0xbe, 0x5, 0xff, 0xfe, 0x27, 0xb,
					0x17},
			},
			Version:      6,
			TrafficClass: 0,
			FlowLabel:    0,
			Length:       24,
			NextHeader:   IPProtocolICMPv6,
			HopLimit:     255,
			SrcIP:        net.IP{0x26, 0x20, 0x0, 0x0, 0x10, 0x5, 0x0, 0x0, 0x26, 0xbe, 0x5, 0xff, 0xfe, 0x27, 0xb, 0x17},
			DstIP:        net.IP{0xfe, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x1f, 0xca, 0xff, 0xfe, 0xb3, 0x76, 0x40},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IPv6 packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
	} else {
		t.Error("No IPv6 layer type found in packet")
	}
	if got, ok := p.Layer(LayerTypeICMPv6).(*ICMPv6); ok {
		want := &ICMPv6{
			BaseLayer: BaseLayer{
				Contents: []byte{0x88, 0x0, 0x1e, 0xd6, 0x40, 0x0, 0x0, 0x0},
				Payload: []byte{0x26, 0x20, 0x0, 0x0, 0x10,
					0x5, 0x0, 0x0, 0x26, 0xbe, 0x5, 0xff, 0xfe, 0x27, 0xb, 0x17},
			},
			TypeCode:  0x8800,
			Checksum:  0x1ed6,
			TypeBytes: []byte{0x40, 0x0, 0x0, 0x0},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ICMPv6 packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
		if got.TypeCode.String() != "NeighborAdvertisement(0)" {
			t.Errorf("ICMPv6 type code, got %q want 'NeighborAdvertisement(0)'", got.TypeCode.String())
		}
	} else {
		t.Error("No ICMPv6 layer type found in packet")
	}
}
