package vm

import (
	"fmt"
	"strings"
)

const (
	opcodeAdd  int64 = 1
	opcodeMul        = 2
	opcodeStop       = 99
)

const (
	// Simplest of vms. Day2 Part1
	TypeSimple = 0
)

type VM struct {
	backup []int64
	instr  []int64

	ip      int
	stopped bool
	// if need to distinct between days
	vmType int
	debug  bool
}

func New(instr []int64, vmType int, debug bool) *VM {
	backup := make([]int64, len(instr))
	insts := make([]int64, len(instr))

	copy(backup, instr)
	copy(insts, instr)

	return &VM{
		backup: backup,
		instr:  insts,
		vmType: vmType,
		debug:  debug,
	}
}

// Reset resets changes in vm
func (v *VM) Reset() {
	v.ip = 0
	v.stopped = false
	copy(v.instr, v.backup)
}

// SetAt sets opcode at index i. returns false if out of range
func (v *VM) SetAt(i int, opcode int64) bool {
	if i >= len(v.instr) {
		return false
	}

	v.instr[i] = opcode

	return true
}

// GetAt returns opcode at index i. returns false if out of range
func (v *VM) GetAt(i int) (int64, bool) {
	if i >= len(v.instr) {
		return 0, false
	}
	return v.instr[i], true
}

// Next sets step
func (v *VM) Next() (stepped bool) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		v.debugPrint("Stopped VM. Recovered %s\n", r)
		v.stopped = true

		stepped = false
	}()

	if v.stopped {
		return false
	}

	opcode := v.instr[v.ip]
	switch opcode {
	case opcodeAdd:
		v.debugPrint("IP: %d, opcode add\n", v.ip)
		arg1 := v.getArgAt(v.ip + 1)
		arg2 := v.getArgAt(v.ip + 2)
		v.placeAt(v.ip+3, arg1+arg2)
		v.ip += 4
		return true

	case opcodeMul:
		v.debugPrint("IP: %d, opcode mul\n", v.ip)
		arg1 := v.getArgAt(v.ip + 1)
		arg2 := v.getArgAt(v.ip + 2)
		v.placeAt(v.ip+3, arg1*arg2)
		v.ip += 4
		return true

	case opcodeStop:
		v.stopped = true
		return false
	}

	return true
}

// getArgAt returns operation argument with instructions at index i
func (v *VM) getArgAt(i int) int64 {
	// if unknown type - fallback on type Simple
	// opcode at i is position to get argument
	arg := v.instr[v.instr[i]]
	v.debugPrint("IP: %d, argument = %d\n", i, arg)
	return arg
}

// placeAt places opcode with instructions at index i
func (v *VM) placeAt(i int, opcode int64) {
	// if unknown type - fallback on type simple
	// opcode at i is position to place result
	pos := v.instr[i]
	v.debugPrint("IP: %d, position = %d\n", i, pos)
	v.instr[pos] = opcode
}

func (v *VM) debugPrint(format string, data ...interface{}) {
	if !v.debug {
		return
	}

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf(format, data...)
}
