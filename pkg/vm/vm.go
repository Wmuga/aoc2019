package vm

import (
	"fmt"
	"strings"
)

const (
	opcodeAdd  int64 = 1
	opcodeMul        = 2
	opcodeInp        = 3
	opcodeOut        = 4
	opcodeJT         = 5
	opcodeJF         = 6
	opcodeLT         = 7
	opcodeEQ         = 8
	opcodeStop       = 99
)

const (
	// Simplest of vms. Day 2 Part 1
	TypeSimple = iota
	// Upgraded VM. Day 5 Part 1
	TypeInOut
	// Upgraded with logical instructions. Day 6 Part 2
	TypeLogical
	// Can await new input
	TypeAwaiter
	// Type beyond our imagination
	TypeUnknown
)

const (
	StatusOK = iota
	StatusAwaitInput
	StatusHalt
	StatusError
)

type VM struct {
	backup []int64
	instr  []int64
	// instruction pointer
	ip      int
	stopped bool
	status  int
	// if need to distinct between days
	vmType int
	debug  bool
	// in, out buffers
	inp []int64
	out []int64
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
	v.status = StatusOK
	v.inp = nil
	v.out = nil
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

func (v *VM) Input(in []int64) {
	v.inp = in
}

func (v *VM) Output() []int64 {
	return v.out
}

func (v *VM) Run() []int64 {
	// Partial reset. IP = 0. Clear output
	v.ip = 0
	v.out = nil
	v.status = StatusOK
	for v.Next() {
	}

	return v.out
}

func (v *VM) Continue() []int64 {
	if v.vmType <= TypeLogical || v.vmType >= TypeUnknown {
		v.debugPrint("Unsupported feature")
		return nil
	}

	v.stopped = false
	v.out = nil

	for v.Next() {
	}

	return v.out
}

func (v *VM) Status() int {
	return v.status
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
		v.status = StatusError

		stepped = false
	}()

	if v.ip >= len(v.instr) {
		v.debugPrint("IP out of range\n")
		v.status = StatusError
		v.stopped = false
	}

	if v.stopped {
		return false
	}

	opcode := v.instr[v.ip]
	opcodeParsed := opcode
	if v.vmType != TypeSimple {
		opcodeParsed = opcode % 100
	}

	switch opcodeParsed {
	case opcodeAdd:
		v.debugPrint("IP: %d, opcode add\n", v.ip)
		modes := parseModes(opcode)
		arg1 := v.getArgAt(modes[2], v.ip+1)
		arg2 := v.getArgAt(modes[1], v.ip+2)
		v.placeAt(modes[0], v.ip+3, arg1+arg2)
		v.ip += 4
		return true

	case opcodeMul:
		v.debugPrint("IP: %d, opcode mul\n", v.ip)
		modes := parseModes(opcode)
		arg1 := v.getArgAt(modes[2], v.ip+1)
		arg2 := v.getArgAt(modes[1], v.ip+2)
		v.placeAt(modes[0], v.ip+3, arg1*arg2)
		v.ip += 4
		return true

	case opcodeInp:
		v.debugPrint("IP: %d, opcode Inp\n", v.ip)
		// if feature not supported - return, empty input - return
		if v.vmType <= TypeSimple || v.vmType >= TypeUnknown {
			v.debugPrint("feature not supported\n", v.ip)
			v.status = StatusError
			v.stopped = true
			return false
		}

		if len(v.inp) == 0 {
			v.debugPrint("Insufficient input\n")
			v.status = StatusAwaitInput
			v.stopped = true
			return false
		}

		modes := parseModes(opcode)
		// get from input. place
		arg := v.inp[0]
		v.inp = v.inp[1:]
		v.placeAt(modes[2], v.ip+1, arg)
		v.ip += 2
		return true

	case opcodeOut:
		v.debugPrint("IP: %d, opcode Out\n", v.ip)
		// if feature not supported - return, empty input - return
		if v.vmType <= TypeSimple || v.vmType >= TypeUnknown {
			v.debugPrint("feature not supported\n")
			v.status = StatusError
			v.stopped = true
			return false
		}
		modes := parseModes(opcode)
		arg := v.getArgAt(modes[2], v.ip+1)
		v.out = append(v.out, arg)
		v.ip += 2
		return true

	case opcodeJT:
		v.debugPrint("IP: %d, opcode JT\n", v.ip)
		// if feature not supported - return, empty input - return
		if v.vmType <= TypeInOut || v.vmType >= TypeUnknown {
			v.debugPrint("feature not supported\n")
			v.status = StatusError
			v.stopped = true
			return false
		}
		modes := parseModes(opcode)
		arg1 := v.getArgAt(modes[2], v.ip+1)
		arg2 := v.getArgAt(modes[1], v.ip+2)
		if arg1 != 0 {
			v.ip = int(arg2)
			return true
		}
		v.ip += 3
		return true

	case opcodeJF:
		v.debugPrint("IP: %d, opcode JF\n", v.ip)
		// if feature not supported - return, empty input - return
		if v.vmType <= TypeInOut || v.vmType >= TypeUnknown {
			v.debugPrint("feature not supported\n")
			v.status = StatusError
			v.stopped = true
			return false
		}
		modes := parseModes(opcode)
		arg1 := v.getArgAt(modes[2], v.ip+1)
		arg2 := v.getArgAt(modes[1], v.ip+2)
		if arg1 == 0 {
			v.ip = int(arg2)
			return true
		}
		v.ip += 3
		return true

	case opcodeLT:
		v.debugPrint("IP: %d, opcode LT\n", v.ip)
		// if feature not supported - return, empty input - return
		if v.vmType <= TypeInOut || v.vmType >= TypeUnknown {
			v.status = StatusError
			v.debugPrint("feature not supported\n")
			v.stopped = true
			return false
		}
		modes := parseModes(opcode)
		arg1 := v.getArgAt(modes[2], v.ip+1)
		arg2 := v.getArgAt(modes[1], v.ip+2)
		var res int64
		if arg1 < arg2 {
			res = 1
		}
		v.placeAt(modes[0], v.ip+3, res)
		v.ip += 4
		return true

	case opcodeEQ:
		v.debugPrint("IP: %d, opcode EQ\n", v.ip)
		// if feature not supported - return, empty input - return
		if v.vmType <= TypeInOut || v.vmType >= TypeUnknown {
			v.debugPrint("feature not supported\n")
			v.status = StatusError
			v.stopped = true
			return false
		}
		modes := parseModes(opcode)
		arg1 := v.getArgAt(modes[2], v.ip+1)
		arg2 := v.getArgAt(modes[1], v.ip+2)
		var res int64
		if arg1 == arg2 {
			res = 1
		}
		v.placeAt(modes[0], v.ip+3, res)
		v.ip += 4
		return true

	// On opcodeStop and unknown opcodes - return
	case opcodeStop:
		v.debugPrint("IP: %d, opcode HALT\n", v.ip)
		v.status = StatusHalt
		v.stopped = true
		return false

	default:
		v.debugPrint("IP: %d, opcode ??\n", v.ip)
		v.status = StatusError
		v.stopped = true
		return false
	}
}

// getArgAt returns operation argument with instructions at index i
func (v *VM) getArgAt(mode int, i int) int64 {
	// if unknown type - fallback on type Simple
	if v.vmType <= TypeSimple || v.vmType >= TypeUnknown {
		mode = 0
	}
	var arg int64
	switch mode {
	// immidiate mode
	case 1:
		arg = v.instr[i]
	// fallback to mode 0
	// opcode at i is position to get argument
	default:
		arg = v.instr[v.instr[i]]
	}
	v.debugPrint("ARG; IP: %d, Mode: %d, opcode: %d, argument = %d\n", i, mode, v.instr[i], arg)
	return arg
}

// placeAt places opcode with instructions at index i
func (v *VM) placeAt(mode int, i int, opcode int64) {
	// if unknown type - fallback on type simple
	if v.vmType <= TypeSimple || v.vmType >= TypeUnknown {
		mode = 0
	}
	var pos int
	switch mode {
	// cannot be immidiate mode
	case 1:
		v.debugPrint("Immidiate mode not supported\n")
		v.status = StatusError
		v.stopped = true
		return
	// fallback mode 0
	// opcode at i is position to place result
	default:
		pos = int(v.instr[i])
	}

	v.debugPrint("OUT; IP: %d, Mode: %d, opcode: %d, position = %d\n", i, mode, v.instr[i], pos)
	v.instr[pos] = opcode
}

func parseModes(opcode int64) [3]int {
	return [3]int{
		int((opcode / 10000) % 10),
		int((opcode / 1000) % 10),
		int((opcode / 100) % 10),
	}
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
