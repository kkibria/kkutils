package kkutils

import (
	"log"
	"time"
	"unsafe"
)

// TimeTrack prints elapsed time
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// IsLittleEndian detects endianness
func IsLittleEndian() bool {
	n := uint32(0x01020304)
	return *(*byte)(unsafe.Pointer(&n)) == 0x04
}
