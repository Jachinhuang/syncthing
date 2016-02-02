// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package discover

import (
	"github.com/calmh/xdr"
)

/*

Announce Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Magic                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                       Device Structure                        \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Number of Extra                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Device Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Announce {
	unsigned int Magic;
	Device This;
	Device Extra<16>;
}

*/

func (o Announce) XDRSize() int {
	return 4 +
		o.This.XDRSize() +
		4 + xdr.SizeOfSlice(o.Extra)
}

func (o Announce) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o Announce) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o Announce) MarshalXDRInto(m *xdr.Marshaller) error {
	m.MarshalUint32(o.Magic)
	if err := o.This.MarshalXDRInto(m); err != nil {
		return err
	}
	if l := len(o.Extra); l > 16 {
		return xdr.ElementSizeExceeded("Extra", l, 16)
	}
	m.MarshalUint32(uint32(len(o.Extra)))
	for i := range o.Extra {
		if err := o.Extra[i].MarshalXDRInto(m); err != nil {
			return err
		}
	}
	return m.Error
}

func (o *Announce) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *Announce) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.Magic = u.UnmarshalUint32()
	(&o.This).UnmarshalXDRFrom(u)
	_ExtraSize := int(u.UnmarshalUint32())
	if _ExtraSize < 0 {
		return xdr.ElementSizeExceeded("Extra", _ExtraSize, 16)
	} else if _ExtraSize == 0 {
		o.Extra = nil
	} else {
		if _ExtraSize > 16 {
			return xdr.ElementSizeExceeded("Extra", _ExtraSize, 16)
		}
		if _ExtraSize <= len(o.Extra) {
			o.Extra = o.Extra[:_ExtraSize]
		} else {
			o.Extra = make([]Device, _ExtraSize)
		}
		for i := range o.Extra {
			(&o.Extra[i]).UnmarshalXDRFrom(u)
		}
	}
	return u.Error
}

/*

Device Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                   ID (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                      Number of Addresses                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Address Structures                \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Relays                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                 Zero or more Relay Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Device {
	opaque ID<32>;
	Address Addresses<16>;
	Relay Relays<16>;
}

*/

func (o Device) XDRSize() int {
	return 4 + len(o.ID) + xdr.Padding(len(o.ID)) +
		4 + xdr.SizeOfSlice(o.Addresses) +
		4 + xdr.SizeOfSlice(o.Relays)
}

func (o Device) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o Device) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o Device) MarshalXDRInto(m *xdr.Marshaller) error {
	if l := len(o.ID); l > 32 {
		return xdr.ElementSizeExceeded("ID", l, 32)
	}
	m.MarshalBytes(o.ID)
	if l := len(o.Addresses); l > 16 {
		return xdr.ElementSizeExceeded("Addresses", l, 16)
	}
	m.MarshalUint32(uint32(len(o.Addresses)))
	for i := range o.Addresses {
		if err := o.Addresses[i].MarshalXDRInto(m); err != nil {
			return err
		}
	}
	if l := len(o.Relays); l > 16 {
		return xdr.ElementSizeExceeded("Relays", l, 16)
	}
	m.MarshalUint32(uint32(len(o.Relays)))
	for i := range o.Relays {
		if err := o.Relays[i].MarshalXDRInto(m); err != nil {
			return err
		}
	}
	return m.Error
}

func (o *Device) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *Device) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.ID = u.UnmarshalBytesMax(32)
	_AddressesSize := int(u.UnmarshalUint32())
	if _AddressesSize < 0 {
		return xdr.ElementSizeExceeded("Addresses", _AddressesSize, 16)
	} else if _AddressesSize == 0 {
		o.Addresses = nil
	} else {
		if _AddressesSize > 16 {
			return xdr.ElementSizeExceeded("Addresses", _AddressesSize, 16)
		}
		if _AddressesSize <= len(o.Addresses) {
			o.Addresses = o.Addresses[:_AddressesSize]
		} else {
			o.Addresses = make([]Address, _AddressesSize)
		}
		for i := range o.Addresses {
			(&o.Addresses[i]).UnmarshalXDRFrom(u)
		}
	}
	_RelaysSize := int(u.UnmarshalUint32())
	if _RelaysSize < 0 {
		return xdr.ElementSizeExceeded("Relays", _RelaysSize, 16)
	} else if _RelaysSize == 0 {
		o.Relays = nil
	} else {
		if _RelaysSize > 16 {
			return xdr.ElementSizeExceeded("Relays", _RelaysSize, 16)
		}
		if _RelaysSize <= len(o.Relays) {
			o.Relays = o.Relays[:_RelaysSize]
		} else {
			o.Relays = make([]Relay, _RelaysSize)
		}
		for i := range o.Relays {
			(&o.Relays[i]).UnmarshalXDRFrom(u)
		}
	}
	return u.Error
}

/*

Address Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                  URL (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Address {
	string URL<2083>;
}

*/

func (o Address) XDRSize() int {
	return 4 + len(o.URL) + xdr.Padding(len(o.URL))
}

func (o Address) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o Address) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o Address) MarshalXDRInto(m *xdr.Marshaller) error {
	if l := len(o.URL); l > 2083 {
		return xdr.ElementSizeExceeded("URL", l, 2083)
	}
	m.MarshalString(o.URL)
	return m.Error
}

func (o *Address) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *Address) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.URL = u.UnmarshalStringMax(2083)
	return u.Error
}

/*

Relay Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                  URL (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                            Latency                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Relay {
	string URL<2083>;
	int Latency;
}

*/

func (o Relay) XDRSize() int {
	return 4 + len(o.URL) + xdr.Padding(len(o.URL)) + 4
}

func (o Relay) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o Relay) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o Relay) MarshalXDRInto(m *xdr.Marshaller) error {
	if l := len(o.URL); l > 2083 {
		return xdr.ElementSizeExceeded("URL", l, 2083)
	}
	m.MarshalString(o.URL)
	m.MarshalUint32(uint32(o.Latency))
	return m.Error
}

func (o *Relay) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *Relay) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.URL = u.UnmarshalStringMax(2083)
	o.Latency = int32(u.UnmarshalUint32())
	return u.Error
}
