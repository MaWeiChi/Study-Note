# linkstatus

```[go]
type LinkStatus int

const (
	LinkUnknown LinkStatus = iota
	LinkUp
	LinkDown
	LinkError
)
```

To confirm the connection of each network interface, it effects the default route with this value.

We check the value by reading the following file.

[The period is 500 milliseconds (0.5 seconds).](https://gitlab.com/moxa/ibg/software/platform/thingspro/edge-device-mil/-/blob/release/iotedge/v2.2.1/pkg/ethernet/ethernet.go#L160)

Before 2.0.0, we define the network interface is connected as the carrier is "1".
In 2.1.0, we use the operstate as it is "up".
In some cases, the interface's operstate is dormant and it's carrier is 1 as the Wi-Fi client connects to the AP with a wrong password.
The physical link is up but the interface's handshake is failed.
If we still use the carrier to determine the link status, it will change the default route frequently.

## 2.0.0

```[go]
os.Open(fmt.Sprintf("/sys/class/net/%s/carrier", iface))
```

## 2.1.0 (due to the new feature as Wi-Fi client)

```[go]
os.Open(fmt.Sprintf("/sys/class/net/%s/operstate", iface))
```

the `iface` is the variable name of the network interface

## [documentation](https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-net)
```[bash]
# carrier
What:		/sys/class/net/<iface>/carrier
Date:		April 2005
KernelVersion:	2.6.12
Contact:	netdev@vger.kernel.org
Description:
		Indicates the current physical link state of the interface.
		Posssible values are:

		== =====================
		0  physical link is down
		1  physical link is up
		== =====================

		Note: some special devices, e.g: bonding and team drivers will
		allow this attribute to be written to force a link state for
		operating correctly and designating another fallback interface.

# operstate
What:		/sys/class/net/<iface>/operstate
Date:		March 2006
KernelVersion:	2.6.17
Contact:	netdev@vger.kernel.org
Description:
		Indicates the interface RFC2863 operational state as a string.

		Possible values are:

		"unknown", "notpresent", "down", "lowerlayerdown", "testing",
		"dormant", "up".
```

### [Operational States](https://www.kernel.org/doc/html/v5.12/networking/operstates.html)

#### **TLV IFLA_OPERSTATE**
contains RFC2863 state of the interface in numeric representation:

**IF_OPER_UNKNOWN (0):**
Interface is in unknown state, neither driver nor userspace has set operational state. Interface must be considered for user data as setting operational state has not been implemented in every driver.

**IF_OPER_NOTPRESENT (1):**
Unused in current kernel (notpresent interfaces normally disappear), just a numerical placeholder.

**IF_OPER_DOWN (2):**
Interface is unable to transfer data on L1, f.e. ethernet is not plugged or interface is ADMIN down.

**IF_OPER_LOWERLAYERDOWN (3):**
Interfaces stacked on an interface that is IF_OPER_DOWN show this state (f.e. VLAN).

**IF_OPER_TESTING (4):**
Unused in current kernel.

**IF_OPER_DORMANT (5):**
Interface is L1 up, but waiting for an external event, f.e. for a protocol to establish. (802.1X)

**IF_OPER_UP (6):**
Interface is operational up and can be used.

This TLV can also be queried via sysfs.
