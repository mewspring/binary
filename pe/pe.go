// Package pe implements parsing of Portable Executable (PE) files.
package pe

import (
	"github.com/mewrev/binary"
	"github.com/mewrev/pe"
)

// Parse parses the provided PE file.
func Parse(path string) (file *binary.File, err error) {
	// Open PE file.
	f, err := pe.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Parse file header.
	fileHdr, err := f.FileHeader()
	if err != nil {
		return nil, err
	}
	file = &binary.File{
		Format: binary.FormatPE,
	}
	switch fileHdr.Arch {
	case pe.ArchI386:
		file.Arch = binary.ArchX86
	case pe.ArchAMD64:
		file.Arch = binary.ArchX86_64
	}

	// Parse optional header.
	optHdr, err := f.OptHeader()
	if err != nil {
		return nil, err
	}
	base := optHdr.ImageBase
	file.Entry = uint64(base + optHdr.EntryRelAddr)

	// Parse sections.
	sectHdrs, err := f.SectHeaders()
	if err != nil {
		return nil, err
	}
	for _, sectHdr := range sectHdrs {
		section := new(binary.Section)
		section.Name = stringFromSZ(sectHdr.Name[:])
		section.Addr = uint64(base + sectHdr.RelAddr)
		if sectHdr.Flags&pe.SectFlagMemExec != 0 {
			section.Perm |= binary.PermExecute
		}
		if sectHdr.Flags&pe.SectFlagMemWrite != 0 {
			section.Perm |= binary.PermWrite
		}
		if sectHdr.Flags&pe.SectFlagMemRead != 0 {
			section.Perm |= binary.PermRead
		}
		section.Data, err = f.Section(sectHdr)
		if err != nil {
			return nil, err
		}
		file.Sections = append(file.Sections, section)
	}

	// TODO: Parse imports and exports.

	return file, nil
}

// stringFromSZ returns a Go string based on the provided NULL-terminated
// string.
func stringFromSZ(buf []byte) string {
	for i, b := range buf {
		if b == 0 {
			return string(buf[:i])
		}
	}
	return string(buf)
}
