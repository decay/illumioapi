package illumioapi

//ProtocolList returns a map for the IANA protocol numbers.
func ProtocolList() map[int]string {

	protocols := make(map[int]string)

	protocols[0] = "HOPOPT"
	protocols[1] = "ICMP"
	protocols[2] = "IGMP"
	protocols[3] = "GGP"
	protocols[4] = "IP-in-IP"
	protocols[5] = "ST"
	protocols[6] = "TCP"
	protocols[7] = "CBT"
	protocols[8] = "EGP"
	protocols[9] = "IGP"
	protocols[10] = "BBN-RCC-MON"
	protocols[11] = "NVP-II"
	protocols[12] = "PUP"
	protocols[13] = "ARGUS"
	protocols[14] = "EMCON"
	protocols[15] = "XNET"
	protocols[16] = "CHAOS"
	protocols[17] = "UDP"
	protocols[18] = "MUX"
	protocols[19] = "DCN-MEAS"
	protocols[20] = "HMP"
	protocols[21] = "PRM"
	protocols[22] = "XNS-IDP"
	protocols[23] = "TRUNK-1"
	protocols[24] = "TRUNK-2"
	protocols[25] = "LEAF-1"
	protocols[26] = "LEAF-2"
	protocols[27] = "RDP"
	protocols[28] = "IRTP"
	protocols[29] = "ISO-TP4"
	protocols[30] = "NETBLT"
	protocols[31] = "MFE-NSP"
	protocols[32] = "MERIT-INP"
	protocols[33] = "DCCP"
	protocols[34] = "3PC"
	protocols[35] = "IDPR"
	protocols[36] = "XTP"
	protocols[37] = "DDP"
	protocols[38] = "IDPR-CMTP"
	protocols[39] = "TP++"
	protocols[40] = "IL"
	protocols[41] = "IPv6"
	protocols[42] = "SDRP"
	protocols[43] = "IPv6-Route"
	protocols[44] = "IPv6-Frag"
	protocols[45] = "IDRP"
	protocols[46] = "RSVP"
	protocols[47] = "GREs"
	protocols[48] = "DSR"
	protocols[49] = "BNA"
	protocols[50] = "ESP"
	protocols[51] = "AH"
	protocols[52] = "I-NLSP"
	protocols[53] = "SWIPE"
	protocols[54] = "NARP"
	protocols[55] = "MOBILE"
	protocols[56] = "TLSP"
	protocols[57] = "SKIP"
	protocols[58] = "IPv6-ICMP"
	protocols[59] = "IPv6-NoNxt"
	protocols[60] = "IPv6-Opts"
	protocols[61] = ""
	protocols[62] = "CFTP"
	protocols[63] = ""
	protocols[64] = "SAT-EXPAK"
	protocols[65] = "KRYPTOLAN"
	protocols[66] = "RVD"
	protocols[67] = "IPPC"
	protocols[68] = ""
	protocols[69] = "SAT-MON"
	protocols[70] = "VISA"
	protocols[71] = "IPCU"
	protocols[72] = "CPNX"
	protocols[73] = "CPHB"
	protocols[74] = "WSN"
	protocols[75] = "PVP"
	protocols[76] = "BR-SAT-MON"
	protocols[77] = "SUN-ND"
	protocols[78] = "WB-MON"
	protocols[79] = "WB-EXPAK"
	protocols[80] = "ISO-IP"
	protocols[81] = "VMTP"
	protocols[82] = "SECURE-VMTP"
	protocols[83] = "VINES"
	protocols[84] = "TTP"
	protocols[84] = "IPTM"
	protocols[85] = "NSFNET-IGP"
	protocols[86] = "DGP"
	protocols[87] = "TCF"
	protocols[88] = "EIGRP"
	protocols[89] = "OSPF"
	protocols[90] = "Sprite-RPC"
	protocols[91] = "LARP"
	protocols[92] = "MTP"
	protocols[93] = "AX.25"
	protocols[94] = "OS"
	protocols[95] = "MICP"
	protocols[96] = "SCC-SP"
	protocols[97] = "ETHERIP"
	protocols[98] = "ENCAP"
	protocols[100] = "GMTP"
	protocols[101] = "IFMP"
	protocols[102] = "PNNI"
	protocols[103] = "PIM"
	protocols[104] = "ARIS"
	protocols[105] = "SCPS"
	protocols[106] = "QNX"
	protocols[107] = "A/N"
	protocols[108] = "IPComp"
	protocols[109] = "SNP"
	protocols[110] = "Compaq-Peer"
	protocols[111] = "IPX-in-IP"
	protocols[112] = "VRRP"
	protocols[113] = "PGM"
	protocols[115] = "L2TP"
	protocols[116] = "DDX"
	protocols[117] = "IATP"
	protocols[118] = "STP"
	protocols[119] = "SRP"
	protocols[120] = "UTI"
	protocols[121] = "SMP"
	protocols[122] = "SM"
	protocols[123] = "PTP"
	protocols[124] = "IS-IS over IPv4"
	protocols[125] = "FIRE"
	protocols[126] = "CRTP"
	protocols[127] = "CRUDP"
	protocols[128] = "SSCOPMCE"
	protocols[129] = "IPLT"
	protocols[130] = "SPS"
	protocols[131] = "PIPE"
	protocols[132] = "SCTP"
	protocols[133] = "FC"
	protocols[134] = "RSVP-E2E-IGNORE"
	protocols[135] = "Mobility Header"
	protocols[136] = "UDPLite"
	protocols[137] = "MPLS-in-IP"
	protocols[138] = "manet"
	protocols[139] = "HIP"
	protocols[140] = "Shim6"
	protocols[141] = "WESP"
	protocols[142] = "ROHC"

	return protocols
}
