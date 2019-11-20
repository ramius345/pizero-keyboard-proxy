package evthid

type EvtHid struct {
	Name string
	Hidcode byte
	Hidmodbits byte
}

func BuildEvtToHidMap() map[int32]EvtHid {
	//map event codes to hidcodes
	m := make(map[int32]EvtHid)
	m[0] = EvtHid{ "KEY_RESERVED",0, 0 }
	m[1] = EvtHid{ "KEY_ESC", 0x29, 0 }
	m[2] = EvtHid{ "KEY_1", 0x1e, 0 }
	m[3] = EvtHid{ "KEY_2", 0x1f, 0 }
	m[4] = EvtHid{ "KEY_3", 0x20, 0 }
	m[5] = EvtHid{ "KEY_4", 0x21, 0 }
	m[6] = EvtHid{ "KEY_5", 0x22, 0 }
	m[7] = EvtHid{ "KEY_6", 0x23, 0 }
	m[8] = EvtHid{ "KEY_7", 0x24, 0 }
	m[9] = EvtHid{ "KEY_8", 0x25, 0 }
	m[10] = EvtHid{ "KEY_9", 0x26, 0 }
	m[11] = EvtHid{ "KEY_0", 0x27, 0 }
	m[12] = EvtHid{ "KEY_MINUS", 0x2d, 0 }
	m[13] = EvtHid{ "KEY_EQUAL", 0x2e, 0 }
	m[14] = EvtHid{ "KEY_BACKSPACE", 0x2a, 0 }
	m[15] = EvtHid{ "KEY_TAB", 0x2b, 0 }
	m[16] = EvtHid{ "KEY_Q", 0x14, 0 }
	m[17] = EvtHid{ "KEY_W", 0x1a, 0 }
	m[18] = EvtHid{ "KEY_E", 0x08, 0 }
	m[19] = EvtHid{ "KEY_R", 0x15, 0 }
	m[20] = EvtHid{ "KEY_T", 0x17, 0 }
	m[21] = EvtHid{ "KEY_Y", 0x1c, 0 }
	m[22] = EvtHid{ "KEY_U", 0x18, 0 }
	m[23] = EvtHid{ "KEY_I", 0x0c, 0 }
	m[24] = EvtHid{ "KEY_O", 0x12, 0 }
	m[25] = EvtHid{ "KEY_P", 0x13, 0 }
	m[26] = EvtHid{ "KEY_LEFTBRACE", 0x2f, 0 }
	m[27] = EvtHid{ "KEY_RIGHTBRACE", 0x30, 0 }
	m[28] = EvtHid{ "KEY_ENTER", 0x28, 0 }
	m[29] = EvtHid{ "KEY_LEFTCTRL", 0, 0x01 }
	m[30] = EvtHid{ "KEY_A", 0x04, 0 }
	m[31] = EvtHid{ "KEY_S", 0x16, 0 }
	m[32] = EvtHid{ "KEY_D", 0x07, 0 }
	m[33] = EvtHid{ "KEY_F", 0x09, 0 }
	m[34] = EvtHid{ "KEY_G", 0x0a, 0 }
	m[35] = EvtHid{ "KEY_H", 0x0b, 0 }
	m[36] = EvtHid{ "KEY_J", 0x0d, 0 }
	m[37] = EvtHid{ "KEY_K", 0x0e, 0 }
	m[38] = EvtHid{ "KEY_L", 0x0f, 0 }
	m[39] = EvtHid{ "KEY_SEMICOLON", 0x33, 0 }
	m[40] = EvtHid{ "KEY_APOSTROPHE", 0x34, 0 }
	m[41] = EvtHid{ "KEY_GRAVE", 0x35, 0 }
	m[42] = EvtHid{ "KEY_LEFTSHIFT", 0, 0x02 }
	m[43] = EvtHid{ "KEY_BACKSLASH", 0x31, 0 }
	m[44] = EvtHid{ "KEY_Z", 0x1d, 0 }
	m[45] = EvtHid{ "KEY_X", 0x1b, 0 }
	m[46] = EvtHid{ "KEY_C", 0x06, 0 }
	m[47] = EvtHid{ "KEY_V", 0x19, 0 }
	m[48] = EvtHid{ "KEY_B", 0x05, 0 }
	m[49] = EvtHid{ "KEY_N", 0x11, 0 }
	m[50] = EvtHid{ "KEY_M", 0x10, 0 }
	m[51] = EvtHid{ "KEY_COMMA", 0x36, 0 }
	m[52] = EvtHid{ "KEY_DOT", 0x37, 0 }
	m[53] = EvtHid{ "KEY_SLASH", 0x38, 0 }
	m[54] = EvtHid{ "KEY_RIGHTSHIFT", 0, 0x20 }
	m[55] = EvtHid{ "KEY_KPASTERISK", 0x55, 0 }
	m[56] = EvtHid{ "KEY_LEFTALT", 0 ,0x04 }
	m[57] = EvtHid{ "KEY_SPACE", 0x2c, 0 }
	m[58] = EvtHid{ "KEY_CAPSLOCK", 0x39, 0 }
	m[59] = EvtHid{ "KEY_F1", 0x3a, 0 }
	m[60] = EvtHid{ "KEY_F2", 0x3b, 0 }
	m[61] = EvtHid{ "KEY_F3", 0x3c, 0 }
	m[62] = EvtHid{ "KEY_F4", 0x3d, 0 }
	m[63] = EvtHid{ "KEY_F5", 0x3e, 0 }
	m[64] = EvtHid{ "KEY_F6", 0x3f, 0 }
	m[65] = EvtHid{ "KEY_F7", 0x40, 0 }
	m[66] = EvtHid{ "KEY_F8", 0x41, 0 }
	m[67] = EvtHid{ "KEY_F9", 0x42, 0 }
	m[68] = EvtHid{ "KEY_F10", 0x43, 0 }
	m[69] = EvtHid{ "KEY_NUMLOCK", 0x53, 0 }
	m[70] = EvtHid{ "KEY_SCROLLLOCK", 0x54, 0 }
	m[71] = EvtHid{ "KEY_KP7", 0x5f, 0 }
	m[72] = EvtHid{ "KEY_KP8", 0x60, 0 }
	m[73] = EvtHid{ "KEY_KP9", 0x61, 0 }
	m[74] = EvtHid{ "KEY_KPMINUS", 0x56, 0 }
	m[75] = EvtHid{ "KEY_KP4", 0x5c,  0 }
	m[76] = EvtHid{ "KEY_KP5", 0x5d, 0 }
	m[77] = EvtHid{ "KEY_KP6", 0x5e, 0 }
	m[78] = EvtHid{ "KEY_KPPLUS", 0x57, 0 }

	m[87] = EvtHid{ "KEY_F11", 0x44, 0 }
	m[88] = EvtHid{ "KEY_F12", 0x45, 0 }

	m[97] = EvtHid{ "KEY_RIGHTCTRL", 0, 0x10 }

	m[100] = EvtHid{ "KEY_RIGHTALT", 0, 0x40 }

	m[102] = EvtHid{ "KEY_HOME", 0x4a, 0 }
	m[103] = EvtHid{ "KEY_UP", 0x52, 0 }
	m[104] = EvtHid{ "KEY_PAGEUP", 0x4b, 0 }
	m[105] = EvtHid{ "KEY_LEFT", 0x50, 0 }
	m[106] = EvtHid{ "KEY_RIGHT", 0x4f, 0 }
	m[107] = EvtHid{ "KEY_END", 0x4d, 0 }
	m[108] = EvtHid{ "KEY_DOWN", 0x51, 0 }
	m[109] = EvtHid{ "KEY_PAGEDOWN", 0x4e, 0 }
	m[110] = EvtHid{ "KEY_INSERT", 0x49, 0 }
	m[111] = EvtHid{ "KEY_DELETE", 0x4c, 0 }

	m[125] = EvtHid{ "KEY_LEFTMETA", 0, 0x08 }
	m[126] = EvtHid{ "KEY_RIGHTMETA", 0, 0x80 }
	m[127] = EvtHid{ "KEY_COMPOSE", 0x65, 0 }
	
	return m
}

func (e EvtHid) GetName() string {
	return e.Name
}

func (e EvtHid) IsNormalKey() bool {
	return e.Hidcode != 0
}

