package radius

// RFC required
const (
	User_Name                uint8 = 1
	User_Password                  = 2
	CHAP_Password                  = 3
	NAS_IP_Address                 = 4
	NAS_Port                       = 5
	Service_Type                   = 6
	Framed_Protocol                = 7
	Framed_IP_Address              = 8
	Framed_IP_Netmask              = 9
	Framed_Routing                 = 10
	Filter_Id                      = 11
	Framed_MTU                     = 12
	Framed_Compression             = 13
	Login_IP_Host                  = 14
	Login_Service                  = 15
	Login_TCP_Port                 = 16
	Reply_Message                  = 18
	Callback_Number                = 19
	Callback_Id                    = 20
	Framed_Route                   = 22
	Framed_IPX_Network             = 23
	State                          = 24
	Class                          = 25
	Vendor_Specific                = 26
	Session_Timeout                = 27
	Idle_Timeout                   = 28
	Termination_Action             = 29
	Called_Station_Id              = 30
	Calling_Station_Id             = 31
	NAS_Identifier                 = 32
	Proxy_State                    = 33
	Login_LAT_Service              = 34
	Login_LAT_Node                 = 35
	Login_LAT_Group                = 36
	Framed_AppleTalk_Link          = 37
	Framed_AppleTalk_Network       = 38
	Framed_AppleTalk_Zone          = 39
	CHAP_Challenge                 = 60
	NAS_Port_Type                  = 61
	Port_Limit                     = 62
	Login_LAT_Port                 = 63
)

// Value constants used to represent the attribute's Value type
const (
	//1-253 octets containing UTF-8 encoded 10646 [7]
	//characters.  Text of length zero (0) MUST NOT be sent;
	//omit the entire attribute instead.
	Text byte = iota

	//1-253 octets containing binary data (values 0 through
	//255 decimal, inclusive).  Strings of length zero (0)
	//MUST NOT be sent; omit the entire attribute instead.
	String

	// 32 bit value, most significant octet first.
	Address

	//32 bit unsigned value, most significant octet first.
	Integer

	//32 bit unsigned value, most significant octet first --
	//seconds since 00:00:00 UTC, January 1, 1970.  The
	//standard Attributes do not use this data type but it is
	//presented here for possible use in future attributes.
	Time
)

type Attribute struct {
	VendorId uint16
	TypeId   uint8
	Value    byte
	Name     string
}

func (a *Attribute) Len() byte {

}
