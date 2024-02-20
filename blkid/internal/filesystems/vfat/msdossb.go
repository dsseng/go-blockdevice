// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "cstruct -pkg vfat -struct MSDOSSB -input msdos.h -endianness LittleEndian"; DO NOT EDIT.

package vfat

import "encoding/binary"

var _ = binary.LittleEndian

// MSDOSSB is a byte slice representing the msdos.h C header.
type MSDOSSB []byte

// Get_ms_ignored returns ms_ignored.
func (s MSDOSSB) Get_ms_ignored() []byte {
	return s[0:3]
}

// Get_ms_sysid returns ms_sysid.
func (s MSDOSSB) Get_ms_sysid() []byte {
	return s[3:11]
}

// Get_ms_sector_size returns ms_sector_size.
func (s MSDOSSB) Get_ms_sector_size() uint16 {
	return binary.LittleEndian.Uint16(s[11:13])
}

// Get_ms_cluster_size returns ms_cluster_size.
func (s MSDOSSB) Get_ms_cluster_size() byte {
	return s[13]
}

// Get_ms_reserved returns ms_reserved.
func (s MSDOSSB) Get_ms_reserved() uint16 {
	return binary.LittleEndian.Uint16(s[14:16])
}

// Get_ms_fats returns ms_fats.
func (s MSDOSSB) Get_ms_fats() byte {
	return s[16]
}

// Get_ms_dir_entries returns ms_dir_entries.
func (s MSDOSSB) Get_ms_dir_entries() uint16 {
	return binary.LittleEndian.Uint16(s[17:19])
}

// Get_ms_sectors returns =0 iff V3 or later.
func (s MSDOSSB) Get_ms_sectors() uint16 {
	return binary.LittleEndian.Uint16(s[19:21])
}

// Get_ms_media returns ms_media.
func (s MSDOSSB) Get_ms_media() byte {
	return s[21]
}

// Get_ms_fat_length returns Sectors per FAT.
func (s MSDOSSB) Get_ms_fat_length() uint16 {
	return binary.LittleEndian.Uint16(s[22:24])
}

// Get_ms_secs_track returns ms_secs_track.
func (s MSDOSSB) Get_ms_secs_track() uint16 {
	return binary.LittleEndian.Uint16(s[24:26])
}

// Get_ms_heads returns ms_heads.
func (s MSDOSSB) Get_ms_heads() uint16 {
	return binary.LittleEndian.Uint16(s[26:28])
}

// Get_ms_hidden returns ms_hidden.
func (s MSDOSSB) Get_ms_hidden() uint32 {
	return binary.LittleEndian.Uint32(s[28:32])
}

// Get_ms_total_sect returns iff ms_sectors == 0.
func (s MSDOSSB) Get_ms_total_sect() uint32 {
	return binary.LittleEndian.Uint32(s[32:36])
}

// Get_ms_drive_number returns ms_drive_number.
func (s MSDOSSB) Get_ms_drive_number() byte {
	return s[36]
}

// Get_ms_boot_flags returns ms_boot_flags.
func (s MSDOSSB) Get_ms_boot_flags() byte {
	return s[37]
}

// Get_ms_ext_boot_sign returns 0x28 - DOS 3.4 EBPB; 0x29 - DOS 4.0 EBPB.
func (s MSDOSSB) Get_ms_ext_boot_sign() byte {
	return s[38]
}

// Get_ms_serno returns ms_serno.
func (s MSDOSSB) Get_ms_serno() []byte {
	return s[39:43]
}

// Get_ms_label returns ms_label.
func (s MSDOSSB) Get_ms_label() []byte {
	return s[43:54]
}

// Get_ms_magic returns ms_magic.
func (s MSDOSSB) Get_ms_magic() []byte {
	return s[54:62]
}

// Get_ms_dummy2 returns ms_dummy2.
func (s MSDOSSB) Get_ms_dummy2() []byte {
	return s[62:510]
}

// Get_ms_pmagic returns ms_pmagic.
func (s MSDOSSB) Get_ms_pmagic() []byte {
	return s[510:512]
}

// MSDOSSB_SIZE is the size of the MSDOSSB struct.
const MSDOSSB_SIZE = 512
