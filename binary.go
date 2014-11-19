// Package binary provides a uniform representation for the data of binary
// executables.
package binary

// A File contains the data of a binary executable.
type File struct {
	// Binary file format of the executable.
	Format Format
	// Machine architecture of the executable's assembly instructions.
	Arch Arch
	// Virtual address to the entry point of the executable, which indicates the
	// starting point of program execution.
	Entry uint64
	// Segments of data with associated access permissions.
	Segments []*Segment
	// Sections of data with associated access permissions.
	Sections []*Section
	// Imports maps from virtual addresses to import names.
	Imports map[uint64]string
	// Exports maps from export names to virtual addresses.
	Exports map[string]uint64
}

// Arch specifies the machine architecture of an executable's assembly
// instructions.
type Arch uint8

// Machine architectures.
const (
	// ArchX86 represents the 32-bit x86 machine architecture, as used by Intel
	// and AMD.
	ArchX86 Arch = 1 + iota
	// ArchX86_64 represents the 64-bit x86-64 machine architecture, as used by
	// Intel and AMD.
	ArchX86_64
)

// Format specifies the file format of a binary executable.
type Format uint8

// File formats.
const (
	// FormatELF represents the Executable and Linkable Format (ELF).
	FormatELF Format = 1 + iota
	// FormatPE represents the Portable Executable (PE) file format.
	FormatPE
)

// A Segment represent a continuous segment of memory.
type Segment struct {
	// Virtual address of the segment once loaded into memory.
	Addr uint64
	// Data contained within the segment.
	Data []byte
	// Access permissions of the segment in memory.
	Perm Perm
}

// A Section represent a continuous section of memory.
type Section struct {
	// Section name.
	Name string
	// Virtual address of the section once loaded into memory.
	Addr uint64
	// Data contained within the section.
	Data []byte
	// Access permissions of the section in memory.
	Perm Perm
}

// Perm specifies the access permissions of a segment or section in memory.
type Perm uint8

// Access permissions.
const (
	// PermExecute specifies that the memory is executable.
	PermExecute Perm = 1 << iota
	// PermWrite specifies that the memory is writeable.
	PermWrite
	// PermRead specifies that the memory is readable.
	PermRead
)
