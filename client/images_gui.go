// +build !nogui

package main

func (i Indicator) pngBytes() []byte {
	return indicatorPNGBytes[i]
}

var indicatorPNGBytes = [][]byte{
	{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x08, 0x06, 0x00, 0x00, 0x00, 0x1f, 0x15, 0xc4,
		0x89, 0x00, 0x00, 0x00, 0x01, 0x73, 0x52, 0x47, 0x42, 0x00, 0xae, 0xce, 0x1c, 0xe9, 0x00, 0x00,
		0x00, 0x0b, 0x49, 0x44, 0x41, 0x54, 0x08, 0xd7, 0x63, 0x60, 0x00, 0x02, 0x00, 0x00, 0x05, 0x00,
		0x01, 0xe2, 0x26, 0x05, 0x9b, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60,
		0x82,
	},
	{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x08, 0x08, 0x06, 0x00, 0x00, 0x00, 0xc4, 0x0f, 0xbe,
		0x8b, 0x00, 0x00, 0x00, 0x04, 0x73, 0x42, 0x49, 0x54, 0x08, 0x08, 0x08, 0x08, 0x7c, 0x08, 0x64,
		0x88, 0x00, 0x00, 0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00, 0x00, 0x00, 0x6b, 0x00, 0x00, 0x00,
		0x6b, 0x01, 0x0e, 0x5f, 0x5b, 0x51, 0x00, 0x00, 0x00, 0x19, 0x74, 0x45, 0x58, 0x74, 0x53, 0x6f,
		0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x00, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x6e, 0x6b, 0x73, 0x63,
		0x61, 0x70, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x9b, 0xee, 0x3c, 0x1a, 0x00, 0x00, 0x00, 0x58, 0x49,
		0x44, 0x41, 0x54, 0x18, 0x95, 0x85, 0xcf, 0xb1, 0x11, 0x83, 0x40, 0x0c, 0x04, 0xc0, 0x95, 0x52,
		0x17, 0xf0, 0x65, 0xd0, 0x93, 0x13, 0x6a, 0xfa, 0xc4, 0x3d, 0x7d, 0x19, 0x14, 0x40, 0x2c, 0x07,
		0x7c, 0x80, 0x19, 0xdb, 0x6c, 0xa8, 0x1b, 0x69, 0x4e, 0x51, 0x55, 0x44, 0x34, 0xac, 0x58, 0x1c,
		0x06, 0xba, 0xaa, 0x2d, 0x8a, 0x86, 0x17, 0x1e, 0x3e, 0xed, 0x78, 0xe6, 0xdc, 0xbc, 0x86, 0xe6,
		0x6c, 0xcd, 0xd3, 0xd9, 0x6f, 0x96, 0xfc, 0x13, 0x82, 0x9c, 0x85, 0x7e, 0x19, 0x89, 0x3e, 0x0b,
		0x5d, 0xed, 0xe8, 0x71, 0xf7, 0xe6, 0x1b, 0xbe, 0xd0, 0x19, 0x1e, 0x20, 0xed, 0x4e, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	},
	{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x08, 0x08, 0x06, 0x00, 0x00, 0x00, 0xc4, 0x0f, 0xbe,
		0x8b, 0x00, 0x00, 0x00, 0x04, 0x73, 0x42, 0x49, 0x54, 0x08, 0x08, 0x08, 0x08, 0x7c, 0x08, 0x64,
		0x88, 0x00, 0x00, 0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00, 0x00, 0x00, 0x6b, 0x00, 0x00, 0x00,
		0x6b, 0x01, 0x0e, 0x5f, 0x5b, 0x51, 0x00, 0x00, 0x00, 0x19, 0x74, 0x45, 0x58, 0x74, 0x53, 0x6f,
		0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x00, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x6e, 0x6b, 0x73, 0x63,
		0x61, 0x70, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x9b, 0xee, 0x3c, 0x1a, 0x00, 0x00, 0x00, 0x5a, 0x49,
		0x44, 0x41, 0x54, 0x18, 0x95, 0x85, 0xcf, 0xb1, 0x0d, 0x83, 0x40, 0x0c, 0x05, 0xd0, 0x77, 0x6e,
		0xa3, 0xb4, 0x8c, 0xc1, 0x52, 0x34, 0x19, 0x89, 0xe6, 0x96, 0x62, 0x0c, 0xca, 0x44, 0x57, 0x9b,
		0x22, 0x57, 0x04, 0x14, 0xe0, 0x95, 0xfe, 0xb2, 0xf5, 0x5d, 0x32, 0x93, 0x4f, 0x19, 0x30, 0x61,
		0xf4, 0xb5, 0xa0, 0x7a, 0xe6, 0x5a, 0xf2, 0x6d, 0xc0, 0x8c, 0x87, 0xbd, 0x86, 0x57, 0xf4, 0xcd,
		0x63, 0xa8, 0xcf, 0xa6, 0xf8, 0x39, 0xfb, 0xcf, 0x18, 0x17, 0x21, 0x88, 0x5e, 0xe8, 0xcc, 0x12,
		0xa8, 0xbd, 0xd0, 0x51, 0x43, 0x2d, 0x77, 0x6f, 0x6e, 0xbd, 0xe2, 0x1a, 0x0f, 0x2f, 0x31, 0x73,
		0x73, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	},
	{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x08, 0x08, 0x06, 0x00, 0x00, 0x00, 0xc4, 0x0f, 0xbe,
		0x8b, 0x00, 0x00, 0x00, 0x04, 0x73, 0x42, 0x49, 0x54, 0x08, 0x08, 0x08, 0x08, 0x7c, 0x08, 0x64,
		0x88, 0x00, 0x00, 0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00, 0x00, 0x00, 0x6b, 0x00, 0x00, 0x00,
		0x6b, 0x01, 0x0e, 0x5f, 0x5b, 0x51, 0x00, 0x00, 0x00, 0x19, 0x74, 0x45, 0x58, 0x74, 0x53, 0x6f,
		0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x00, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x6e, 0x6b, 0x73, 0x63,
		0x61, 0x70, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x9b, 0xee, 0x3c, 0x1a, 0x00, 0x00, 0x00, 0x57, 0x49,
		0x44, 0x41, 0x54, 0x18, 0x95, 0x85, 0xcf, 0xa1, 0x11, 0xc3, 0x40, 0x0c, 0x04, 0xc0, 0x95, 0x68,
		0xb8, 0xcb, 0xf8, 0xa6, 0x4c, 0x52, 0x52, 0xc8, 0x37, 0xf5, 0x65, 0x98, 0x1b, 0xcb, 0x20, 0x0f,
		0x12, 0x8f, 0x9d, 0x2c, 0xd4, 0x8d, 0x34, 0xa7, 0xa8, 0x2a, 0x21, 0x16, 0xac, 0x68, 0xde, 0x06,
		0x7a, 0xa9, 0x2d, 0x94, 0x05, 0x2f, 0x3c, 0x7c, 0xdb, 0xf1, 0xcc, 0xb9, 0x79, 0x0e, 0xcd, 0xd9,
		0x9a, 0x1f, 0x67, 0xaf, 0xb4, 0xfc, 0x11, 0x82, 0x9c, 0x85, 0xee, 0x8c, 0x44, 0x9f, 0x85, 0xce,
		0x76, 0xf4, 0xf8, 0xf7, 0xe6, 0x01, 0xb8, 0xfa, 0x19, 0x1d, 0x1e, 0x33, 0x74, 0x3e, 0x00, 0x00,
		0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	},
	{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x08, 0x08, 0x06, 0x00, 0x00, 0x00, 0xc4, 0x0f, 0xbe,
		0x8b, 0x00, 0x00, 0x00, 0x01, 0x73, 0x52, 0x47, 0x42, 0x00, 0xae, 0xce, 0x1c, 0xe9, 0x00, 0x00,
		0x00, 0x3e, 0x49, 0x44, 0x41, 0x54, 0x18, 0x57, 0x63, 0xf8, 0xff, 0xff, 0x3f, 0x03, 0x3b, 0xe3,
		0x7f, 0x71, 0x20, 0xae, 0x03, 0xe2, 0xb5, 0x50, 0x0c, 0x62, 0x8b, 0x83, 0xe4, 0x60, 0x92, 0x5b,
		0x81, 0xf8, 0x00, 0x1a, 0x06, 0x89, 0x89, 0x33, 0x40, 0x55, 0xa3, 0x4b, 0xc2, 0x70, 0x1d, 0x03,
		0xd4, 0x48, 0x5c, 0x0a, 0xd6, 0x12, 0xa5, 0x80, 0xa0, 0x15, 0xf8, 0x1d, 0x49, 0xc8, 0x9b, 0x00,
		0xaf, 0x96, 0x6e, 0x7d, 0x64, 0x1b, 0xd8, 0x8b, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44,
		0xae, 0x42, 0x60, 0x82,
	},
	{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x08, 0x08, 0x06, 0x00, 0x00, 0x00, 0xc4, 0x0f, 0xbe,
		0x8b, 0x00, 0x00, 0x00, 0x06, 0x62, 0x4b, 0x47, 0x44, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0xa0,
		0xbd, 0xa7, 0x93, 0x00, 0x00, 0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00, 0x00, 0x00, 0x6b, 0x00,
		0x00, 0x00, 0x6b, 0x01, 0x0e, 0x5f, 0x5b, 0x51, 0x00, 0x00, 0x00, 0x07, 0x74, 0x49, 0x4d, 0x45,
		0x07, 0xdd, 0x05, 0x05, 0x13, 0x30, 0x18, 0x4f, 0x00, 0x69, 0x9f, 0x00, 0x00, 0x00, 0x3b, 0x49,
		0x44, 0x41, 0x54, 0x28, 0xcf, 0x63, 0xf8, 0xff, 0xff, 0x3f, 0x03, 0x10, 0x88, 0x03, 0x71, 0x1d,
		0x10, 0xaf, 0x85, 0x62, 0x10, 0x5b, 0x1c, 0x2a, 0x07, 0x96, 0xdc, 0x0a, 0xc4, 0x07, 0xd0, 0xf0,
		0x56, 0xa8, 0x1c, 0x58, 0xf5, 0x01, 0x1c, 0x18, 0x24, 0x07, 0x36, 0x12, 0x97, 0x82, 0xb5, 0x44,
		0x29, 0x20, 0x68, 0x05, 0x7e, 0x47, 0x12, 0xf2, 0x26, 0x00, 0x00, 0x83, 0x30, 0xd9, 0xc1, 0x52,
		0x3f, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	},
	{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x09, 0x08, 0x06, 0x00, 0x00, 0x00, 0x0f, 0x53, 0x6d,
		0x2e, 0x00, 0x00, 0x00, 0x04, 0x73, 0x42, 0x49, 0x54, 0x08, 0x08, 0x08, 0x08, 0x7c, 0x08, 0x64,
		0x88, 0x00, 0x00, 0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00, 0x00, 0x00, 0x6e, 0x00, 0x00, 0x00,
		0x6e, 0x01, 0x23, 0xe5, 0x3e, 0xa7, 0x00, 0x00, 0x00, 0x62, 0x49, 0x44, 0x41, 0x54, 0x18, 0xd3,
		0x63, 0x50, 0x50, 0x50, 0x58, 0x0e, 0xc4, 0x99, 0xff, 0xff, 0xff, 0x67, 0x40, 0xc6, 0x20, 0x31,
		0x20, 0x5e, 0x00, 0x62, 0x54, 0x00, 0xf1, 0x1f, 0x20, 0xf6, 0x45, 0x92, 0xf4, 0x05, 0xe2, 0xbf,
		0x40, 0x5c, 0x07, 0x13, 0x98, 0x0d, 0xc4, 0x5f, 0x80, 0xd8, 0x18, 0x88, 0x8d, 0xa0, 0xec, 0x05,
		0x20, 0x39, 0x98, 0x02, 0x16, 0x20, 0xde, 0x09, 0xc4, 0xaf, 0x80, 0xf8, 0x25, 0x10, 0xef, 0x05,
		0x62, 0x56, 0xb8, 0x02, 0xa8, 0x22, 0x69, 0x20, 0xfe, 0x0e, 0xc4, 0x3f, 0x80, 0x58, 0x16, 0x26,
		0x4e, 0x9c, 0x09, 0x78, 0xdd, 0x40, 0xd0, 0x17, 0x84, 0xc2, 0x01, 0x00, 0xcf, 0x55, 0x74, 0xd8,
		0x8c, 0x0e, 0x23, 0xe0, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	},
	{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x0a, 0x08, 0x06, 0x00, 0x00, 0x00, 0x8d, 0x32, 0xcf,
		0xbd, 0x00, 0x00, 0x00, 0x04, 0x73, 0x42, 0x49, 0x54, 0x08, 0x08, 0x08, 0x08, 0x7c, 0x08, 0x64,
		0x88, 0x00, 0x00, 0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00, 0x00, 0x00, 0x61, 0x00, 0x00, 0x00,
		0x61, 0x01, 0x55, 0x2b, 0x90, 0xbd, 0x00, 0x00, 0x00, 0x6d, 0x49, 0x44, 0x41, 0x54, 0x18, 0xd3,
		0x63, 0xf8, 0xff, 0xff, 0x3f, 0x03, 0x0c, 0x2b, 0x28, 0x28, 0x70, 0x00, 0xf1, 0x45, 0x20, 0x36,
		0x46, 0x16, 0x07, 0x61, 0x06, 0x34, 0x85, 0x02, 0x40, 0xfc, 0x1f, 0x88, 0x03, 0x08, 0x29, 0x14,
		0xc4, 0xa9, 0x10, 0x28, 0xa8, 0x06, 0xc4, 0x0e, 0x50, 0xec, 0x0b, 0x55, 0x58, 0x83, 0x24, 0x06,
		0xc2, 0x6a, 0x20, 0x85, 0x37, 0xa1, 0x92, 0xf8, 0xf0, 0x2d, 0x90, 0x42, 0x23, 0x90, 0x55, 0x40,
		0x1c, 0x08, 0xc4, 0x09, 0x50, 0x89, 0x2e, 0x20, 0x0e, 0x05, 0xe2, 0x30, 0x20, 0x0e, 0x07, 0x62,
		0x73, 0x74, 0x37, 0x8a, 0x12, 0xeb, 0x19, 0x31, 0x62, 0x15, 0x8a, 0x13, 0xab, 0x90, 0x17, 0x88,
		0x2f, 0x83, 0xdc, 0x8d, 0xae, 0x10, 0x00, 0xb6, 0xdd, 0xbf, 0x8e, 0x33, 0x76, 0x4d, 0x53, 0x00,
		0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	},
}
