# RISC-V disasembler in long format

Disassemble RISC-V in a long format for easy understanding.

In the usual way
```
   11e76: 03 35 84 fc  	ld	a0, -56(s0)
   11e7a: 93 15 05 02  	slli	a1, a0, 32
   11e7e: 81 91        	srli	a1, a1, 32
   11e80: 23 34 b4 fa  	sd	a1, -88(s0)
   11e84: 9b 05 05 00  	sext.w	a1, a0
   11e88: 13 b5 05 04  	sltiu	a0, a1, 64
   11e8c: 23 38 a4 fa  	sd	a0, -80(s0)
```

In this tool
```
   11e76: 03 35 84 fc  	LoadI64		a0, -56(s0)
   11e7a: 93 15 05 02  	ShiftLeftLogicalImm	a1, a0, 32
   11e7e: 81 91        	ShiftRightLogicalImm	a1, a1, 32
   11e80: 23 34 b4 fa  	StoreI64	a1, -88(s0)
   11e84: 9b 05 05 00  	SignEXtend.I32	a1, a0
   11e88: 13 b5 05 04  	SetLessThanImmUnsigned	a0, a1, 64
   11e8c: 23 38 a4 fa  	StoreI64	a0, -80(s0)
```

## Install

This tool uses llvm-objdump. So you have to install llvm first.  
For example,
```
$ sudo apt install llvm
```

Then
```
$ go install github.com/tetsu-koba/riscv_disasm_long@latest
```

## Usage

```
$ riscv_disasm_long
2023/03/14 21:35:02 Usage:
2023/03/14 21:35:02 	riscv_disasm_long objfile  ("llvm-objdump -d objfile" is called internally)
2023/03/14 21:35:02 	riscv_disasm_long - < objdump_output
```

## Example

```
$ cat hello.c 
#include <stdio.h>

int main()
{
	printf("Hello, world!\n");
	return 0;
}
$ zig cc --target=riscv64-linux-musl -o hello hello.c
$ file hello
hello: ELF 64-bit LSB executable, UCB RISC-V, RVC, double-float ABI, version 1 (SYSV), statically linked, with debug_info, not stripped
$ llvm-objdump -d hello |head -20

hello:	file format elf64-littleriscv

Disassembly of section .text:

0000000000011d7c <_start>:
   11d7c: 97 61 00 00  	auipc	gp, 6
   11d80: 93 81 c1 4e  	addi	gp, gp, 1260
   11d84: 0a 85        	mv	a0, sp

0000000000011d86 <.Lpcrel_hi1>:
   11d86: 97 e5 fe ff  	auipc	a1, 1048558
   11d8a: 93 85 a5 27  	addi	a1, a1, 634
   11d8e: 13 71 01 ff  	andi	sp, sp, -16
   11d92: 17 03 00 00  	auipc	t1, 0
   11d96: 67 00 83 00  	jr	8(t1)

0000000000011d9a <_start_c>:
   11d9a: 0c 41        	lw	a1, 0(a0)

$ riscv_disasm_long hello |head -20

hello:	file format elf64-littleriscv

Disassembly of section .text:

0000000000011d7c <_start>:
   11d7c: 97 61 00 00  	AddUpperImmPC	gp, 6
   11d80: 93 81 c1 4e  	ADDImm		gp, gp, 1260
   11d84: 0a 85        	mv		a0, sp

0000000000011d86 <.Lpcrel_hi1>:
   11d86: 97 e5 fe ff  	AddUpperImmPC	a1, 1048558
   11d8a: 93 85 a5 27  	ADDImm		a1, a1, 634
   11d8e: 13 71 01 ff  	ANDImm		sp, sp, -16
   11d92: 17 03 00 00  	AddUpperImmPC	t1, 0
   11d96: 67 00 83 00  	JumpReg		8(t1)

0000000000011d9a <_start_c>:
   11d9a: 0c 41        	LoadI32		a1, 0(a0)


```





