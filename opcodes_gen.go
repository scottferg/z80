/*

Copyright (c) 2010 Andrea Fazzi

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.
8
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

*/

//
// Automatically generated file -- DO NOT EDIT
//

package z80

func initOpcodes() {
	// BEGIN of non shifted opcodes
	/* NOP */
	OpcodesMap[0x00] = instr__NOP
	/* LD BC,nnnn */
	OpcodesMap[0x01] = instr__LD_BC_NNNN
	/* LD (BC),A */
	OpcodesMap[0x02] = instr__LD_iBC_A
	/* INC BC */
	OpcodesMap[0x03] = instr__INC_BC
	/* INC B */
	OpcodesMap[0x04] = instr__INC_B
	/* DEC B */
	OpcodesMap[0x05] = instr__DEC_B
	/* LD B,nn */
	OpcodesMap[0x06] = instr__LD_B_NN
	/* RLCA */
	OpcodesMap[0x07] = instr__RLCA
	/* LD (nn),SP */
	OpcodesMap[0x08] = instr__LD_iNNNN_SP
	/* ADD HL,BC */
	OpcodesMap[0x09] = instr__ADD_HL_BC
	/* LD A,(BC) */
	OpcodesMap[0x0a] = instr__LD_A_iBC
	/* DEC BC */
	OpcodesMap[0x0b] = instr__DEC_BC
	/* INC C */
	OpcodesMap[0x0c] = instr__INC_C
	/* DEC C */
	OpcodesMap[0x0d] = instr__DEC_C
	/* LD C,nn */
	OpcodesMap[0x0e] = instr__LD_C_NN
	/* RRCA */
	OpcodesMap[0x0f] = instr__RRCA
	/* STOP */
	OpcodesMap[0x10] = instr__STOP
	/* LD DE,nnnn */
	OpcodesMap[0x11] = instr__LD_DE_NNNN
	/* LD (DE),A */
	OpcodesMap[0x12] = instr__LD_iDE_A
	/* INC DE */
	OpcodesMap[0x13] = instr__INC_DE
	/* INC D */
	OpcodesMap[0x14] = instr__INC_D
	/* DEC D */
	OpcodesMap[0x15] = instr__DEC_D
	/* LD D,nn */
	OpcodesMap[0x16] = instr__LD_D_NN
	/* RLA */
	OpcodesMap[0x17] = instr__RLA
	/* JR offset */
	OpcodesMap[0x18] = instr__JR_OFFSET
	/* ADD HL,DE */
	OpcodesMap[0x19] = instr__ADD_HL_DE
	/* LD A,(DE) */
	OpcodesMap[0x1a] = instr__LD_A_iDE
	/* DEC DE */
	OpcodesMap[0x1b] = instr__DEC_DE
	/* INC E */
	OpcodesMap[0x1c] = instr__INC_E
	/* DEC E */
	OpcodesMap[0x1d] = instr__DEC_E
	/* LD E,nn */
	OpcodesMap[0x1e] = instr__LD_E_NN
	/* RRA */
	OpcodesMap[0x1f] = instr__RRA
	/* JR NZ,offset */
	OpcodesMap[0x20] = instr__JR_NZ_OFFSET
	/* LD HL,nnnn */
	OpcodesMap[0x21] = instr__LD_HL_NNNN
	/* LDI (HL),A */
	OpcodesMap[0x22] = instr__LDI_iHL_A
	/* INC HL */
	OpcodesMap[0x23] = instr__INC_HL
	/* INC H */
	OpcodesMap[0x24] = instr__INC_H
	/* DEC H */
	OpcodesMap[0x25] = instr__DEC_H
	/* LD H,nn */
	OpcodesMap[0x26] = instr__LD_H_NN
	/* DAA */
	OpcodesMap[0x27] = instr__DAA
	/* JR Z,offset */
	OpcodesMap[0x28] = instr__JR_Z_OFFSET
	/* ADD HL,HL */
	OpcodesMap[0x29] = instr__ADD_HL_HL
	/* LDI A,(HL) */
	OpcodesMap[0x2a] = instr__LDI_A_iHL
	/* DEC HL */
	OpcodesMap[0x2b] = instr__DEC_HL
	/* INC L */
	OpcodesMap[0x2c] = instr__INC_L
	/* DEC L */
	OpcodesMap[0x2d] = instr__DEC_L
	/* LD L,nn */
	OpcodesMap[0x2e] = instr__LD_L_NN
	/* CPL */
	OpcodesMap[0x2f] = instr__CPL
	/* JR NC,offset */
	OpcodesMap[0x30] = instr__JR_NC_OFFSET
	/* LD SP,nnnn */
	OpcodesMap[0x31] = instr__LD_SP_NNNN
	/* LDD (HL),A */
	OpcodesMap[0x32] = instr__LDD_iHL_A
	/* INC SP */
	OpcodesMap[0x33] = instr__INC_SP
	/* INC (HL) */
	OpcodesMap[0x34] = instr__INC_iHL
	/* DEC (HL) */
	OpcodesMap[0x35] = instr__DEC_iHL
	/* LD (HL),nn */
	OpcodesMap[0x36] = instr__LD_iHL_NN
	/* SCF */
	OpcodesMap[0x37] = instr__SCF
	/* JR C,offset */
	OpcodesMap[0x38] = instr__JR_C_OFFSET
	/* ADD HL,SP */
	OpcodesMap[0x39] = instr__ADD_HL_SP
	/* LDD A,(HL) */
	OpcodesMap[0x3a] = instr__LDD_A_iHL
	/* DEC SP */
	OpcodesMap[0x3b] = instr__DEC_SP
	/* INC A */
	OpcodesMap[0x3c] = instr__INC_A
	/* DEC A */
	OpcodesMap[0x3d] = instr__DEC_A
	/* LD A,nn */
	OpcodesMap[0x3e] = instr__LD_A_NN
	/* CCF */
	OpcodesMap[0x3f] = instr__CCF
	/* LD B,B */
	OpcodesMap[0x40] = instr__LD_B_B
	/* LD B,C */
	OpcodesMap[0x41] = instr__LD_B_C
	/* LD B,D */
	OpcodesMap[0x42] = instr__LD_B_D
	/* LD B,E */
	OpcodesMap[0x43] = instr__LD_B_E
	/* LD B,H */
	OpcodesMap[0x44] = instr__LD_B_H
	/* LD B,L */
	OpcodesMap[0x45] = instr__LD_B_L
	/* LD B,(HL) */
	OpcodesMap[0x46] = instr__LD_B_iHL
	/* LD B,A */
	OpcodesMap[0x47] = instr__LD_B_A
	/* LD C,B */
	OpcodesMap[0x48] = instr__LD_C_B
	/* LD C,C */
	OpcodesMap[0x49] = instr__LD_C_C
	/* LD C,D */
	OpcodesMap[0x4a] = instr__LD_C_D
	/* LD C,E */
	OpcodesMap[0x4b] = instr__LD_C_E
	/* LD C,H */
	OpcodesMap[0x4c] = instr__LD_C_H
	/* LD C,L */
	OpcodesMap[0x4d] = instr__LD_C_L
	/* LD C,(HL) */
	OpcodesMap[0x4e] = instr__LD_C_iHL
	/* LD C,A */
	OpcodesMap[0x4f] = instr__LD_C_A
	/* LD D,B */
	OpcodesMap[0x50] = instr__LD_D_B
	/* LD D,C */
	OpcodesMap[0x51] = instr__LD_D_C
	/* LD D,D */
	OpcodesMap[0x52] = instr__LD_D_D
	/* LD D,E */
	OpcodesMap[0x53] = instr__LD_D_E
	/* LD D,H */
	OpcodesMap[0x54] = instr__LD_D_H
	/* LD D,L */
	OpcodesMap[0x55] = instr__LD_D_L
	/* LD D,(HL) */
	OpcodesMap[0x56] = instr__LD_D_iHL
	/* LD D,A */
	OpcodesMap[0x57] = instr__LD_D_A
	/* LD E,B */
	OpcodesMap[0x58] = instr__LD_E_B
	/* LD E,C */
	OpcodesMap[0x59] = instr__LD_E_C
	/* LD E,D */
	OpcodesMap[0x5a] = instr__LD_E_D
	/* LD E,E */
	OpcodesMap[0x5b] = instr__LD_E_E
	/* LD E,H */
	OpcodesMap[0x5c] = instr__LD_E_H
	/* LD E,L */
	OpcodesMap[0x5d] = instr__LD_E_L
	/* LD E,(HL) */
	OpcodesMap[0x5e] = instr__LD_E_iHL
	/* LD E,A */
	OpcodesMap[0x5f] = instr__LD_E_A
	/* LD H,B */
	OpcodesMap[0x60] = instr__LD_H_B
	/* LD H,C */
	OpcodesMap[0x61] = instr__LD_H_C
	/* LD H,D */
	OpcodesMap[0x62] = instr__LD_H_D
	/* LD H,E */
	OpcodesMap[0x63] = instr__LD_H_E
	/* LD H,H */
	OpcodesMap[0x64] = instr__LD_H_H
	/* LD H,L */
	OpcodesMap[0x65] = instr__LD_H_L
	/* LD H,(HL) */
	OpcodesMap[0x66] = instr__LD_H_iHL
	/* LD H,A */
	OpcodesMap[0x67] = instr__LD_H_A
	/* LD L,B */
	OpcodesMap[0x68] = instr__LD_L_B
	/* LD L,C */
	OpcodesMap[0x69] = instr__LD_L_C
	/* LD L,D */
	OpcodesMap[0x6a] = instr__LD_L_D
	/* LD L,E */
	OpcodesMap[0x6b] = instr__LD_L_E
	/* LD L,H */
	OpcodesMap[0x6c] = instr__LD_L_H
	/* LD L,L */
	OpcodesMap[0x6d] = instr__LD_L_L
	/* LD L,(HL) */
	OpcodesMap[0x6e] = instr__LD_L_iHL
	/* LD L,A */
	OpcodesMap[0x6f] = instr__LD_L_A
	/* LD (HL),B */
	OpcodesMap[0x70] = instr__LD_iHL_B
	/* LD (HL),C */
	OpcodesMap[0x71] = instr__LD_iHL_C
	/* LD (HL),D */
	OpcodesMap[0x72] = instr__LD_iHL_D
	/* LD (HL),E */
	OpcodesMap[0x73] = instr__LD_iHL_E
	/* LD (HL),H */
	OpcodesMap[0x74] = instr__LD_iHL_H
	/* LD (HL),L */
	OpcodesMap[0x75] = instr__LD_iHL_L
	/* HALT */
	OpcodesMap[0x76] = instr__HALT
	/* LD (HL),A */
	OpcodesMap[0x77] = instr__LD_iHL_A
	/* LD A,B */
	OpcodesMap[0x78] = instr__LD_A_B
	/* LD A,C */
	OpcodesMap[0x79] = instr__LD_A_C
	/* LD A,D */
	OpcodesMap[0x7a] = instr__LD_A_D
	/* LD A,E */
	OpcodesMap[0x7b] = instr__LD_A_E
	/* LD A,H */
	OpcodesMap[0x7c] = instr__LD_A_H
	/* LD A,L */
	OpcodesMap[0x7d] = instr__LD_A_L
	/* LD A,(HL) */
	OpcodesMap[0x7e] = instr__LD_A_iHL
	/* LD A,A */
	OpcodesMap[0x7f] = instr__LD_A_A
	/* ADD A,B */
	OpcodesMap[0x80] = instr__ADD_A_B
	/* ADD A,C */
	OpcodesMap[0x81] = instr__ADD_A_C
	/* ADD A,D */
	OpcodesMap[0x82] = instr__ADD_A_D
	/* ADD A,E */
	OpcodesMap[0x83] = instr__ADD_A_E
	/* ADD A,H */
	OpcodesMap[0x84] = instr__ADD_A_H
	/* ADD A,L */
	OpcodesMap[0x85] = instr__ADD_A_L
	/* ADD A,(HL) */
	OpcodesMap[0x86] = instr__ADD_A_iHL
	/* ADD A,A */
	OpcodesMap[0x87] = instr__ADD_A_A
	/* ADC A,B */
	OpcodesMap[0x88] = instr__ADC_A_B
	/* ADC A,C */
	OpcodesMap[0x89] = instr__ADC_A_C
	/* ADC A,D */
	OpcodesMap[0x8a] = instr__ADC_A_D
	/* ADC A,E */
	OpcodesMap[0x8b] = instr__ADC_A_E
	/* ADC A,H */
	OpcodesMap[0x8c] = instr__ADC_A_H
	/* ADC A,L */
	OpcodesMap[0x8d] = instr__ADC_A_L
	/* ADC A,(HL) */
	OpcodesMap[0x8e] = instr__ADC_A_iHL
	/* ADC A,A */
	OpcodesMap[0x8f] = instr__ADC_A_A
	/* SUB A,B */
	OpcodesMap[0x90] = instr__SUB_A_B
	/* SUB A,C */
	OpcodesMap[0x91] = instr__SUB_A_C
	/* SUB A,D */
	OpcodesMap[0x92] = instr__SUB_A_D
	/* SUB A,E */
	OpcodesMap[0x93] = instr__SUB_A_E
	/* SUB A,H */
	OpcodesMap[0x94] = instr__SUB_A_H
	/* SUB A,L */
	OpcodesMap[0x95] = instr__SUB_A_L
	/* SUB A,(HL) */
	OpcodesMap[0x96] = instr__SUB_A_iHL
	/* SUB A,A */
	OpcodesMap[0x97] = instr__SUB_A_A
	/* SBC A,B */
	OpcodesMap[0x98] = instr__SBC_A_B
	/* SBC A,C */
	OpcodesMap[0x99] = instr__SBC_A_C
	/* SBC A,D */
	OpcodesMap[0x9a] = instr__SBC_A_D
	/* SBC A,E */
	OpcodesMap[0x9b] = instr__SBC_A_E
	/* SBC A,H */
	OpcodesMap[0x9c] = instr__SBC_A_H
	/* SBC A,L */
	OpcodesMap[0x9d] = instr__SBC_A_L
	/* SBC A,(HL) */
	OpcodesMap[0x9e] = instr__SBC_A_iHL
	/* SBC A,A */
	OpcodesMap[0x9f] = instr__SBC_A_A
	/* AND A,B */
	OpcodesMap[0xa0] = instr__AND_A_B
	/* AND A,C */
	OpcodesMap[0xa1] = instr__AND_A_C
	/* AND A,D */
	OpcodesMap[0xa2] = instr__AND_A_D
	/* AND A,E */
	OpcodesMap[0xa3] = instr__AND_A_E
	/* AND A,H */
	OpcodesMap[0xa4] = instr__AND_A_H
	/* AND A,L */
	OpcodesMap[0xa5] = instr__AND_A_L
	/* AND A,(HL) */
	OpcodesMap[0xa6] = instr__AND_A_iHL
	/* AND A,A */
	OpcodesMap[0xa7] = instr__AND_A_A
	/* XOR A,B */
	OpcodesMap[0xa8] = instr__XOR_A_B
	/* XOR A,C */
	OpcodesMap[0xa9] = instr__XOR_A_C
	/* XOR A,D */
	OpcodesMap[0xaa] = instr__XOR_A_D
	/* XOR A,E */
	OpcodesMap[0xab] = instr__XOR_A_E
	/* XOR A,H */
	OpcodesMap[0xac] = instr__XOR_A_H
	/* XOR A,L */
	OpcodesMap[0xad] = instr__XOR_A_L
	/* XOR A,(HL) */
	OpcodesMap[0xae] = instr__XOR_A_iHL
	/* XOR A,A */
	OpcodesMap[0xaf] = instr__XOR_A_A
	/* OR A,B */
	OpcodesMap[0xb0] = instr__OR_A_B
	/* OR A,C */
	OpcodesMap[0xb1] = instr__OR_A_C
	/* OR A,D */
	OpcodesMap[0xb2] = instr__OR_A_D
	/* OR A,E */
	OpcodesMap[0xb3] = instr__OR_A_E
	/* OR A,H */
	OpcodesMap[0xb4] = instr__OR_A_H
	/* OR A,L */
	OpcodesMap[0xb5] = instr__OR_A_L
	/* OR A,(HL) */
	OpcodesMap[0xb6] = instr__OR_A_iHL
	/* OR A,A */
	OpcodesMap[0xb7] = instr__OR_A_A
	/* CP B */
	OpcodesMap[0xb8] = instr__CP_B
	/* CP C */
	OpcodesMap[0xb9] = instr__CP_C
	/* CP D */
	OpcodesMap[0xba] = instr__CP_D
	/* CP E */
	OpcodesMap[0xbb] = instr__CP_E
	/* CP H */
	OpcodesMap[0xbc] = instr__CP_H
	/* CP L */
	OpcodesMap[0xbd] = instr__CP_L
	/* CP (HL) */
	OpcodesMap[0xbe] = instr__CP_iHL
	/* CP A */
	OpcodesMap[0xbf] = instr__CP_A
	/* RET NZ */
	OpcodesMap[0xc0] = instr__RET_NZ
	/* POP BC */
	OpcodesMap[0xc1] = instr__POP_BC
	/* JP NZ,nnnn */
	OpcodesMap[0xc2] = instr__JP_NZ_NNNN
	/* JP nnnn */
	OpcodesMap[0xc3] = instr__JP_NNNN
	/* CALL NZ,nnnn */
	OpcodesMap[0xc4] = instr__CALL_NZ_NNNN
	/* PUSH BC */
	OpcodesMap[0xc5] = instr__PUSH_BC
	/* ADD A,nn */
	OpcodesMap[0xc6] = instr__ADD_A_NN
	/* RST 00 */
	OpcodesMap[0xc7] = instr__RST_00
	/* RET Z */
	OpcodesMap[0xc8] = instr__RET_Z
	/* RET */
	OpcodesMap[0xc9] = instr__RET
	/* JP Z,nnnn */
	OpcodesMap[0xca] = instr__JP_Z_NNNN
	/* shift CB */
	OpcodesMap[0xcb] = instr__SHIFT_CB
	/* CALL Z,nnnn */
	OpcodesMap[0xcc] = instr__CALL_Z_NNNN
	/* CALL nnnn */
	OpcodesMap[0xcd] = instr__CALL_NNNN
	/* ADC A,nn */
	OpcodesMap[0xce] = instr__ADC_A_NN
	/* RST 8 */
	OpcodesMap[0xcf] = instr__RST_8
	/* RET NC */
	OpcodesMap[0xd0] = instr__RET_NC
	/* POP DE */
	OpcodesMap[0xd1] = instr__POP_DE
	/* JP NC,nnnn */
	OpcodesMap[0xd2] = instr__JP_NC_NNNN
	/* OUT (nn),A */
	OpcodesMap[0xd3] = instr__OUT_iNN_A
	/* CALL NC,nnnn */
	OpcodesMap[0xd4] = instr__CALL_NC_NNNN
	/* PUSH DE */
	OpcodesMap[0xd5] = instr__PUSH_DE
	/* SUB nn */
	OpcodesMap[0xd6] = instr__SUB_NN
	/* RST 10 */
	OpcodesMap[0xd7] = instr__RST_10
	/* RET C */
	OpcodesMap[0xd8] = instr__RET_C
	/* RETI */
	OpcodesMap[0xd9] = instr__RETI
	/* JP C,nnnn */
	OpcodesMap[0xda] = instr__JP_C_NNNN
	/* CALL C,nnnn */
	OpcodesMap[0xdc] = instr__CALL_C_NNNN
	/* SBC A,nn */
	OpcodesMap[0xde] = instr__SBC_A_NN
	/* RST 18 */
	OpcodesMap[0xdf] = instr__RST_18
	/* LD (FF00+n),A */
	OpcodesMap[0xe0] = instr__LD_iFF00N_A
	/* POP HL */
	OpcodesMap[0xe1] = instr__POP_HL
	/* LD (FF00+C),A */
	OpcodesMap[0xe2] = instr__LD_iFF00C_A
	/* PUSH HL */
	OpcodesMap[0xe5] = instr__PUSH_HL
	/* AND nn */
	OpcodesMap[0xe6] = instr__AND_NN
	/* RST 20 */
	OpcodesMap[0xe7] = instr__RST_20
	/* ADD SP,dd */
	OpcodesMap[0xe8] = instr__ADD_SP_DD
	/* JP HL */
	OpcodesMap[0xe9] = instr__JP_HL
	/* LD (nn),A */
	OpcodesMap[0xea] = instr__LD_iNNNN_A
	/* XOR A,nn */
	OpcodesMap[0xee] = instr__XOR_A_NN
	/* RST 28 */
	OpcodesMap[0xef] = instr__RST_28
	/* LD A,(FF00+n) */
	OpcodesMap[0xf0] = instr__LD_A_iFF00N
	/* POP AF */
	OpcodesMap[0xf1] = instr__POP_AF
	/* LD A,(FF00+C) */
	OpcodesMap[0xf2] = instr__LD_A_iFF00C
	/* DI */
	OpcodesMap[0xf3] = instr__DI
	/* PUSH AF */
	OpcodesMap[0xf5] = instr__PUSH_AF
	/* OR nn */
	OpcodesMap[0xf6] = instr__OR_NN
	/* RST 30 */
	OpcodesMap[0xf7] = instr__RST_30
	/* LD HL,SP+dd */
	OpcodesMap[0xf8] = instr__LD_HL_SPDD
	/* LD SP,HL */
	OpcodesMap[0xf9] = instr__LD_SP_HL
	/* LD A,(nn) */
	OpcodesMap[0xfa] = instr__LD_A_iNNNN
	/* EI */
	OpcodesMap[0xfb] = instr__EI
	/* CP nn */
	OpcodesMap[0xfe] = instr__CP_NN
	/* RST 38 */
	OpcodesMap[0xff] = instr__RST_38

	// END of non shifted opcodes

	// BEGIN of 0xcb shifted opcodes
	/* RLC B */
	OpcodesMap[SHIFT_0xCB+0x00] = instrCB__RLC_B
	/* RLC C */
	OpcodesMap[SHIFT_0xCB+0x01] = instrCB__RLC_C
	/* RLC D */
	OpcodesMap[SHIFT_0xCB+0x02] = instrCB__RLC_D
	/* RLC E */
	OpcodesMap[SHIFT_0xCB+0x03] = instrCB__RLC_E
	/* RLC H */
	OpcodesMap[SHIFT_0xCB+0x04] = instrCB__RLC_H
	/* RLC L */
	OpcodesMap[SHIFT_0xCB+0x05] = instrCB__RLC_L
	/* RLC (HL) */
	OpcodesMap[SHIFT_0xCB+0x06] = instrCB__RLC_iHL
	/* RLC A */
	OpcodesMap[SHIFT_0xCB+0x07] = instrCB__RLC_A
	/* RRC B */
	OpcodesMap[SHIFT_0xCB+0x08] = instrCB__RRC_B
	/* RRC C */
	OpcodesMap[SHIFT_0xCB+0x09] = instrCB__RRC_C
	/* RRC D */
	OpcodesMap[SHIFT_0xCB+0x0a] = instrCB__RRC_D
	/* RRC E */
	OpcodesMap[SHIFT_0xCB+0x0b] = instrCB__RRC_E
	/* RRC H */
	OpcodesMap[SHIFT_0xCB+0x0c] = instrCB__RRC_H
	/* RRC L */
	OpcodesMap[SHIFT_0xCB+0x0d] = instrCB__RRC_L
	/* RRC (HL) */
	OpcodesMap[SHIFT_0xCB+0x0e] = instrCB__RRC_iHL
	/* RRC A */
	OpcodesMap[SHIFT_0xCB+0x0f] = instrCB__RRC_A
	/* RL B */
	OpcodesMap[SHIFT_0xCB+0x10] = instrCB__RL_B
	/* RL C */
	OpcodesMap[SHIFT_0xCB+0x11] = instrCB__RL_C
	/* RL D */
	OpcodesMap[SHIFT_0xCB+0x12] = instrCB__RL_D
	/* RL E */
	OpcodesMap[SHIFT_0xCB+0x13] = instrCB__RL_E
	/* RL H */
	OpcodesMap[SHIFT_0xCB+0x14] = instrCB__RL_H
	/* RL L */
	OpcodesMap[SHIFT_0xCB+0x15] = instrCB__RL_L
	/* RL (HL) */
	OpcodesMap[SHIFT_0xCB+0x16] = instrCB__RL_iHL
	/* RL A */
	OpcodesMap[SHIFT_0xCB+0x17] = instrCB__RL_A
	/* RR B */
	OpcodesMap[SHIFT_0xCB+0x18] = instrCB__RR_B
	/* RR C */
	OpcodesMap[SHIFT_0xCB+0x19] = instrCB__RR_C
	/* RR D */
	OpcodesMap[SHIFT_0xCB+0x1a] = instrCB__RR_D
	/* RR E */
	OpcodesMap[SHIFT_0xCB+0x1b] = instrCB__RR_E
	/* RR H */
	OpcodesMap[SHIFT_0xCB+0x1c] = instrCB__RR_H
	/* RR L */
	OpcodesMap[SHIFT_0xCB+0x1d] = instrCB__RR_L
	/* RR (HL) */
	OpcodesMap[SHIFT_0xCB+0x1e] = instrCB__RR_iHL
	/* RR A */
	OpcodesMap[SHIFT_0xCB+0x1f] = instrCB__RR_A
	/* SLA B */
	OpcodesMap[SHIFT_0xCB+0x20] = instrCB__SLA_B
	/* SLA C */
	OpcodesMap[SHIFT_0xCB+0x21] = instrCB__SLA_C
	/* SLA D */
	OpcodesMap[SHIFT_0xCB+0x22] = instrCB__SLA_D
	/* SLA E */
	OpcodesMap[SHIFT_0xCB+0x23] = instrCB__SLA_E
	/* SLA H */
	OpcodesMap[SHIFT_0xCB+0x24] = instrCB__SLA_H
	/* SLA L */
	OpcodesMap[SHIFT_0xCB+0x25] = instrCB__SLA_L
	/* SLA (HL) */
	OpcodesMap[SHIFT_0xCB+0x26] = instrCB__SLA_iHL
	/* SLA A */
	OpcodesMap[SHIFT_0xCB+0x27] = instrCB__SLA_A
	/* SRA B */
	OpcodesMap[SHIFT_0xCB+0x28] = instrCB__SRA_B
	/* SRA C */
	OpcodesMap[SHIFT_0xCB+0x29] = instrCB__SRA_C
	/* SRA D */
	OpcodesMap[SHIFT_0xCB+0x2a] = instrCB__SRA_D
	/* SRA E */
	OpcodesMap[SHIFT_0xCB+0x2b] = instrCB__SRA_E
	/* SRA H */
	OpcodesMap[SHIFT_0xCB+0x2c] = instrCB__SRA_H
	/* SRA L */
	OpcodesMap[SHIFT_0xCB+0x2d] = instrCB__SRA_L
	/* SRA (HL) */
	OpcodesMap[SHIFT_0xCB+0x2e] = instrCB__SRA_iHL
	/* SRA A */
	OpcodesMap[SHIFT_0xCB+0x2f] = instrCB__SRA_A
	/* SWAP B */
	OpcodesMap[SHIFT_0xCB+0x30] = instrCB__SWAP_B
	/* SWAP C */
	OpcodesMap[SHIFT_0xCB+0x31] = instrCB__SWAP_C
	/* SWAP D */
	OpcodesMap[SHIFT_0xCB+0x32] = instrCB__SWAP_D
	/* SWAP E */
	OpcodesMap[SHIFT_0xCB+0x33] = instrCB__SWAP_E
	/* SWAP H */
	OpcodesMap[SHIFT_0xCB+0x34] = instrCB__SWAP_H
	/* SWAP L */
	OpcodesMap[SHIFT_0xCB+0x35] = instrCB__SWAP_L
	/* SWAP (HL) */
	OpcodesMap[SHIFT_0xCB+0x36] = instrCB__SWAP_iHL
	/* SWAP A */
	OpcodesMap[SHIFT_0xCB+0x37] = instrCB__SWAP_A
	/* SRL B */
	OpcodesMap[SHIFT_0xCB+0x38] = instrCB__SRL_B
	/* SRL C */
	OpcodesMap[SHIFT_0xCB+0x39] = instrCB__SRL_C
	/* SRL D */
	OpcodesMap[SHIFT_0xCB+0x3a] = instrCB__SRL_D
	/* SRL E */
	OpcodesMap[SHIFT_0xCB+0x3b] = instrCB__SRL_E
	/* SRL H */
	OpcodesMap[SHIFT_0xCB+0x3c] = instrCB__SRL_H
	/* SRL L */
	OpcodesMap[SHIFT_0xCB+0x3d] = instrCB__SRL_L
	/* SRL (HL) */
	OpcodesMap[SHIFT_0xCB+0x3e] = instrCB__SRL_iHL
	/* SRL A */
	OpcodesMap[SHIFT_0xCB+0x3f] = instrCB__SRL_A
	/* BIT 0,B */
	OpcodesMap[SHIFT_0xCB+0x40] = instrCB__BIT_0_B
	/* BIT 0,C */
	OpcodesMap[SHIFT_0xCB+0x41] = instrCB__BIT_0_C
	/* BIT 0,D */
	OpcodesMap[SHIFT_0xCB+0x42] = instrCB__BIT_0_D
	/* BIT 0,E */
	OpcodesMap[SHIFT_0xCB+0x43] = instrCB__BIT_0_E
	/* BIT 0,H */
	OpcodesMap[SHIFT_0xCB+0x44] = instrCB__BIT_0_H
	/* BIT 0,L */
	OpcodesMap[SHIFT_0xCB+0x45] = instrCB__BIT_0_L
	/* BIT 0,(HL) */
	OpcodesMap[SHIFT_0xCB+0x46] = instrCB__BIT_0_iHL
	/* BIT 0,A */
	OpcodesMap[SHIFT_0xCB+0x47] = instrCB__BIT_0_A
	/* BIT 1,B */
	OpcodesMap[SHIFT_0xCB+0x48] = instrCB__BIT_1_B
	/* BIT 1,C */
	OpcodesMap[SHIFT_0xCB+0x49] = instrCB__BIT_1_C
	/* BIT 1,D */
	OpcodesMap[SHIFT_0xCB+0x4a] = instrCB__BIT_1_D
	/* BIT 1,E */
	OpcodesMap[SHIFT_0xCB+0x4b] = instrCB__BIT_1_E
	/* BIT 1,H */
	OpcodesMap[SHIFT_0xCB+0x4c] = instrCB__BIT_1_H
	/* BIT 1,L */
	OpcodesMap[SHIFT_0xCB+0x4d] = instrCB__BIT_1_L
	/* BIT 1,(HL) */
	OpcodesMap[SHIFT_0xCB+0x4e] = instrCB__BIT_1_iHL
	/* BIT 1,A */
	OpcodesMap[SHIFT_0xCB+0x4f] = instrCB__BIT_1_A
	/* BIT 2,B */
	OpcodesMap[SHIFT_0xCB+0x50] = instrCB__BIT_2_B
	/* BIT 2,C */
	OpcodesMap[SHIFT_0xCB+0x51] = instrCB__BIT_2_C
	/* BIT 2,D */
	OpcodesMap[SHIFT_0xCB+0x52] = instrCB__BIT_2_D
	/* BIT 2,E */
	OpcodesMap[SHIFT_0xCB+0x53] = instrCB__BIT_2_E
	/* BIT 2,H */
	OpcodesMap[SHIFT_0xCB+0x54] = instrCB__BIT_2_H
	/* BIT 2,L */
	OpcodesMap[SHIFT_0xCB+0x55] = instrCB__BIT_2_L
	/* BIT 2,(HL) */
	OpcodesMap[SHIFT_0xCB+0x56] = instrCB__BIT_2_iHL
	/* BIT 2,A */
	OpcodesMap[SHIFT_0xCB+0x57] = instrCB__BIT_2_A
	/* BIT 3,B */
	OpcodesMap[SHIFT_0xCB+0x58] = instrCB__BIT_3_B
	/* BIT 3,C */
	OpcodesMap[SHIFT_0xCB+0x59] = instrCB__BIT_3_C
	/* BIT 3,D */
	OpcodesMap[SHIFT_0xCB+0x5a] = instrCB__BIT_3_D
	/* BIT 3,E */
	OpcodesMap[SHIFT_0xCB+0x5b] = instrCB__BIT_3_E
	/* BIT 3,H */
	OpcodesMap[SHIFT_0xCB+0x5c] = instrCB__BIT_3_H
	/* BIT 3,L */
	OpcodesMap[SHIFT_0xCB+0x5d] = instrCB__BIT_3_L
	/* BIT 3,(HL) */
	OpcodesMap[SHIFT_0xCB+0x5e] = instrCB__BIT_3_iHL
	/* BIT 3,A */
	OpcodesMap[SHIFT_0xCB+0x5f] = instrCB__BIT_3_A
	/* BIT 4,B */
	OpcodesMap[SHIFT_0xCB+0x60] = instrCB__BIT_4_B
	/* BIT 4,C */
	OpcodesMap[SHIFT_0xCB+0x61] = instrCB__BIT_4_C
	/* BIT 4,D */
	OpcodesMap[SHIFT_0xCB+0x62] = instrCB__BIT_4_D
	/* BIT 4,E */
	OpcodesMap[SHIFT_0xCB+0x63] = instrCB__BIT_4_E
	/* BIT 4,H */
	OpcodesMap[SHIFT_0xCB+0x64] = instrCB__BIT_4_H
	/* BIT 4,L */
	OpcodesMap[SHIFT_0xCB+0x65] = instrCB__BIT_4_L
	/* BIT 4,(HL) */
	OpcodesMap[SHIFT_0xCB+0x66] = instrCB__BIT_4_iHL
	/* BIT 4,A */
	OpcodesMap[SHIFT_0xCB+0x67] = instrCB__BIT_4_A
	/* BIT 5,B */
	OpcodesMap[SHIFT_0xCB+0x68] = instrCB__BIT_5_B
	/* BIT 5,C */
	OpcodesMap[SHIFT_0xCB+0x69] = instrCB__BIT_5_C
	/* BIT 5,D */
	OpcodesMap[SHIFT_0xCB+0x6a] = instrCB__BIT_5_D
	/* BIT 5,E */
	OpcodesMap[SHIFT_0xCB+0x6b] = instrCB__BIT_5_E
	/* BIT 5,H */
	OpcodesMap[SHIFT_0xCB+0x6c] = instrCB__BIT_5_H
	/* BIT 5,L */
	OpcodesMap[SHIFT_0xCB+0x6d] = instrCB__BIT_5_L
	/* BIT 5,(HL) */
	OpcodesMap[SHIFT_0xCB+0x6e] = instrCB__BIT_5_iHL
	/* BIT 5,A */
	OpcodesMap[SHIFT_0xCB+0x6f] = instrCB__BIT_5_A
	/* BIT 6,B */
	OpcodesMap[SHIFT_0xCB+0x70] = instrCB__BIT_6_B
	/* BIT 6,C */
	OpcodesMap[SHIFT_0xCB+0x71] = instrCB__BIT_6_C
	/* BIT 6,D */
	OpcodesMap[SHIFT_0xCB+0x72] = instrCB__BIT_6_D
	/* BIT 6,E */
	OpcodesMap[SHIFT_0xCB+0x73] = instrCB__BIT_6_E
	/* BIT 6,H */
	OpcodesMap[SHIFT_0xCB+0x74] = instrCB__BIT_6_H
	/* BIT 6,L */
	OpcodesMap[SHIFT_0xCB+0x75] = instrCB__BIT_6_L
	/* BIT 6,(HL) */
	OpcodesMap[SHIFT_0xCB+0x76] = instrCB__BIT_6_iHL
	/* BIT 6,A */
	OpcodesMap[SHIFT_0xCB+0x77] = instrCB__BIT_6_A
	/* BIT 7,B */
	OpcodesMap[SHIFT_0xCB+0x78] = instrCB__BIT_7_B
	/* BIT 7,C */
	OpcodesMap[SHIFT_0xCB+0x79] = instrCB__BIT_7_C
	/* BIT 7,D */
	OpcodesMap[SHIFT_0xCB+0x7a] = instrCB__BIT_7_D
	/* BIT 7,E */
	OpcodesMap[SHIFT_0xCB+0x7b] = instrCB__BIT_7_E
	/* BIT 7,H */
	OpcodesMap[SHIFT_0xCB+0x7c] = instrCB__BIT_7_H
	/* BIT 7,L */
	OpcodesMap[SHIFT_0xCB+0x7d] = instrCB__BIT_7_L
	/* BIT 7,(HL) */
	OpcodesMap[SHIFT_0xCB+0x7e] = instrCB__BIT_7_iHL
	/* BIT 7,A */
	OpcodesMap[SHIFT_0xCB+0x7f] = instrCB__BIT_7_A
	/* RES 0,B */
	OpcodesMap[SHIFT_0xCB+0x80] = instrCB__RES_0_B
	/* RES 0,C */
	OpcodesMap[SHIFT_0xCB+0x81] = instrCB__RES_0_C
	/* RES 0,D */
	OpcodesMap[SHIFT_0xCB+0x82] = instrCB__RES_0_D
	/* RES 0,E */
	OpcodesMap[SHIFT_0xCB+0x83] = instrCB__RES_0_E
	/* RES 0,H */
	OpcodesMap[SHIFT_0xCB+0x84] = instrCB__RES_0_H
	/* RES 0,L */
	OpcodesMap[SHIFT_0xCB+0x85] = instrCB__RES_0_L
	/* RES 0,(HL) */
	OpcodesMap[SHIFT_0xCB+0x86] = instrCB__RES_0_iHL
	/* RES 0,A */
	OpcodesMap[SHIFT_0xCB+0x87] = instrCB__RES_0_A
	/* RES 1,B */
	OpcodesMap[SHIFT_0xCB+0x88] = instrCB__RES_1_B
	/* RES 1,C */
	OpcodesMap[SHIFT_0xCB+0x89] = instrCB__RES_1_C
	/* RES 1,D */
	OpcodesMap[SHIFT_0xCB+0x8a] = instrCB__RES_1_D
	/* RES 1,E */
	OpcodesMap[SHIFT_0xCB+0x8b] = instrCB__RES_1_E
	/* RES 1,H */
	OpcodesMap[SHIFT_0xCB+0x8c] = instrCB__RES_1_H
	/* RES 1,L */
	OpcodesMap[SHIFT_0xCB+0x8d] = instrCB__RES_1_L
	/* RES 1,(HL) */
	OpcodesMap[SHIFT_0xCB+0x8e] = instrCB__RES_1_iHL
	/* RES 1,A */
	OpcodesMap[SHIFT_0xCB+0x8f] = instrCB__RES_1_A
	/* RES 2,B */
	OpcodesMap[SHIFT_0xCB+0x90] = instrCB__RES_2_B
	/* RES 2,C */
	OpcodesMap[SHIFT_0xCB+0x91] = instrCB__RES_2_C
	/* RES 2,D */
	OpcodesMap[SHIFT_0xCB+0x92] = instrCB__RES_2_D
	/* RES 2,E */
	OpcodesMap[SHIFT_0xCB+0x93] = instrCB__RES_2_E
	/* RES 2,H */
	OpcodesMap[SHIFT_0xCB+0x94] = instrCB__RES_2_H
	/* RES 2,L */
	OpcodesMap[SHIFT_0xCB+0x95] = instrCB__RES_2_L
	/* RES 2,(HL) */
	OpcodesMap[SHIFT_0xCB+0x96] = instrCB__RES_2_iHL
	/* RES 2,A */
	OpcodesMap[SHIFT_0xCB+0x97] = instrCB__RES_2_A
	/* RES 3,B */
	OpcodesMap[SHIFT_0xCB+0x98] = instrCB__RES_3_B
	/* RES 3,C */
	OpcodesMap[SHIFT_0xCB+0x99] = instrCB__RES_3_C
	/* RES 3,D */
	OpcodesMap[SHIFT_0xCB+0x9a] = instrCB__RES_3_D
	/* RES 3,E */
	OpcodesMap[SHIFT_0xCB+0x9b] = instrCB__RES_3_E
	/* RES 3,H */
	OpcodesMap[SHIFT_0xCB+0x9c] = instrCB__RES_3_H
	/* RES 3,L */
	OpcodesMap[SHIFT_0xCB+0x9d] = instrCB__RES_3_L
	/* RES 3,(HL) */
	OpcodesMap[SHIFT_0xCB+0x9e] = instrCB__RES_3_iHL
	/* RES 3,A */
	OpcodesMap[SHIFT_0xCB+0x9f] = instrCB__RES_3_A
	/* RES 4,B */
	OpcodesMap[SHIFT_0xCB+0xa0] = instrCB__RES_4_B
	/* RES 4,C */
	OpcodesMap[SHIFT_0xCB+0xa1] = instrCB__RES_4_C
	/* RES 4,D */
	OpcodesMap[SHIFT_0xCB+0xa2] = instrCB__RES_4_D
	/* RES 4,E */
	OpcodesMap[SHIFT_0xCB+0xa3] = instrCB__RES_4_E
	/* RES 4,H */
	OpcodesMap[SHIFT_0xCB+0xa4] = instrCB__RES_4_H
	/* RES 4,L */
	OpcodesMap[SHIFT_0xCB+0xa5] = instrCB__RES_4_L
	/* RES 4,(HL) */
	OpcodesMap[SHIFT_0xCB+0xa6] = instrCB__RES_4_iHL
	/* RES 4,A */
	OpcodesMap[SHIFT_0xCB+0xa7] = instrCB__RES_4_A
	/* RES 5,B */
	OpcodesMap[SHIFT_0xCB+0xa8] = instrCB__RES_5_B
	/* RES 5,C */
	OpcodesMap[SHIFT_0xCB+0xa9] = instrCB__RES_5_C
	/* RES 5,D */
	OpcodesMap[SHIFT_0xCB+0xaa] = instrCB__RES_5_D
	/* RES 5,E */
	OpcodesMap[SHIFT_0xCB+0xab] = instrCB__RES_5_E
	/* RES 5,H */
	OpcodesMap[SHIFT_0xCB+0xac] = instrCB__RES_5_H
	/* RES 5,L */
	OpcodesMap[SHIFT_0xCB+0xad] = instrCB__RES_5_L
	/* RES 5,(HL) */
	OpcodesMap[SHIFT_0xCB+0xae] = instrCB__RES_5_iHL
	/* RES 5,A */
	OpcodesMap[SHIFT_0xCB+0xaf] = instrCB__RES_5_A
	/* RES 6,B */
	OpcodesMap[SHIFT_0xCB+0xb0] = instrCB__RES_6_B
	/* RES 6,C */
	OpcodesMap[SHIFT_0xCB+0xb1] = instrCB__RES_6_C
	/* RES 6,D */
	OpcodesMap[SHIFT_0xCB+0xb2] = instrCB__RES_6_D
	/* RES 6,E */
	OpcodesMap[SHIFT_0xCB+0xb3] = instrCB__RES_6_E
	/* RES 6,H */
	OpcodesMap[SHIFT_0xCB+0xb4] = instrCB__RES_6_H
	/* RES 6,L */
	OpcodesMap[SHIFT_0xCB+0xb5] = instrCB__RES_6_L
	/* RES 6,(HL) */
	OpcodesMap[SHIFT_0xCB+0xb6] = instrCB__RES_6_iHL
	/* RES 6,A */
	OpcodesMap[SHIFT_0xCB+0xb7] = instrCB__RES_6_A
	/* RES 7,B */
	OpcodesMap[SHIFT_0xCB+0xb8] = instrCB__RES_7_B
	/* RES 7,C */
	OpcodesMap[SHIFT_0xCB+0xb9] = instrCB__RES_7_C
	/* RES 7,D */
	OpcodesMap[SHIFT_0xCB+0xba] = instrCB__RES_7_D
	/* RES 7,E */
	OpcodesMap[SHIFT_0xCB+0xbb] = instrCB__RES_7_E
	/* RES 7,H */
	OpcodesMap[SHIFT_0xCB+0xbc] = instrCB__RES_7_H
	/* RES 7,L */
	OpcodesMap[SHIFT_0xCB+0xbd] = instrCB__RES_7_L
	/* RES 7,(HL) */
	OpcodesMap[SHIFT_0xCB+0xbe] = instrCB__RES_7_iHL
	/* RES 7,A */
	OpcodesMap[SHIFT_0xCB+0xbf] = instrCB__RES_7_A
	/* SET 0,B */
	OpcodesMap[SHIFT_0xCB+0xc0] = instrCB__SET_0_B
	/* SET 0,C */
	OpcodesMap[SHIFT_0xCB+0xc1] = instrCB__SET_0_C
	/* SET 0,D */
	OpcodesMap[SHIFT_0xCB+0xc2] = instrCB__SET_0_D
	/* SET 0,E */
	OpcodesMap[SHIFT_0xCB+0xc3] = instrCB__SET_0_E
	/* SET 0,H */
	OpcodesMap[SHIFT_0xCB+0xc4] = instrCB__SET_0_H
	/* SET 0,L */
	OpcodesMap[SHIFT_0xCB+0xc5] = instrCB__SET_0_L
	/* SET 0,(HL) */
	OpcodesMap[SHIFT_0xCB+0xc6] = instrCB__SET_0_iHL
	/* SET 0,A */
	OpcodesMap[SHIFT_0xCB+0xc7] = instrCB__SET_0_A
	/* SET 1,B */
	OpcodesMap[SHIFT_0xCB+0xc8] = instrCB__SET_1_B
	/* SET 1,C */
	OpcodesMap[SHIFT_0xCB+0xc9] = instrCB__SET_1_C
	/* SET 1,D */
	OpcodesMap[SHIFT_0xCB+0xca] = instrCB__SET_1_D
	/* SET 1,E */
	OpcodesMap[SHIFT_0xCB+0xcb] = instrCB__SET_1_E
	/* SET 1,H */
	OpcodesMap[SHIFT_0xCB+0xcc] = instrCB__SET_1_H
	/* SET 1,L */
	OpcodesMap[SHIFT_0xCB+0xcd] = instrCB__SET_1_L
	/* SET 1,(HL) */
	OpcodesMap[SHIFT_0xCB+0xce] = instrCB__SET_1_iHL
	/* SET 1,A */
	OpcodesMap[SHIFT_0xCB+0xcf] = instrCB__SET_1_A
	/* SET 2,B */
	OpcodesMap[SHIFT_0xCB+0xd0] = instrCB__SET_2_B
	/* SET 2,C */
	OpcodesMap[SHIFT_0xCB+0xd1] = instrCB__SET_2_C
	/* SET 2,D */
	OpcodesMap[SHIFT_0xCB+0xd2] = instrCB__SET_2_D
	/* SET 2,E */
	OpcodesMap[SHIFT_0xCB+0xd3] = instrCB__SET_2_E
	/* SET 2,H */
	OpcodesMap[SHIFT_0xCB+0xd4] = instrCB__SET_2_H
	/* SET 2,L */
	OpcodesMap[SHIFT_0xCB+0xd5] = instrCB__SET_2_L
	/* SET 2,(HL) */
	OpcodesMap[SHIFT_0xCB+0xd6] = instrCB__SET_2_iHL
	/* SET 2,A */
	OpcodesMap[SHIFT_0xCB+0xd7] = instrCB__SET_2_A
	/* SET 3,B */
	OpcodesMap[SHIFT_0xCB+0xd8] = instrCB__SET_3_B
	/* SET 3,C */
	OpcodesMap[SHIFT_0xCB+0xd9] = instrCB__SET_3_C
	/* SET 3,D */
	OpcodesMap[SHIFT_0xCB+0xda] = instrCB__SET_3_D
	/* SET 3,E */
	OpcodesMap[SHIFT_0xCB+0xdb] = instrCB__SET_3_E
	/* SET 3,H */
	OpcodesMap[SHIFT_0xCB+0xdc] = instrCB__SET_3_H
	/* SET 3,L */
	OpcodesMap[SHIFT_0xCB+0xdd] = instrCB__SET_3_L
	/* SET 3,(HL) */
	OpcodesMap[SHIFT_0xCB+0xde] = instrCB__SET_3_iHL
	/* SET 3,A */
	OpcodesMap[SHIFT_0xCB+0xdf] = instrCB__SET_3_A
	/* SET 4,B */
	OpcodesMap[SHIFT_0xCB+0xe0] = instrCB__SET_4_B
	/* SET 4,C */
	OpcodesMap[SHIFT_0xCB+0xe1] = instrCB__SET_4_C
	/* SET 4,D */
	OpcodesMap[SHIFT_0xCB+0xe2] = instrCB__SET_4_D
	/* SET 4,E */
	OpcodesMap[SHIFT_0xCB+0xe3] = instrCB__SET_4_E
	/* SET 4,H */
	OpcodesMap[SHIFT_0xCB+0xe4] = instrCB__SET_4_H
	/* SET 4,L */
	OpcodesMap[SHIFT_0xCB+0xe5] = instrCB__SET_4_L
	/* SET 4,(HL) */
	OpcodesMap[SHIFT_0xCB+0xe6] = instrCB__SET_4_iHL
	/* SET 4,A */
	OpcodesMap[SHIFT_0xCB+0xe7] = instrCB__SET_4_A
	/* SET 5,B */
	OpcodesMap[SHIFT_0xCB+0xe8] = instrCB__SET_5_B
	/* SET 5,C */
	OpcodesMap[SHIFT_0xCB+0xe9] = instrCB__SET_5_C
	/* SET 5,D */
	OpcodesMap[SHIFT_0xCB+0xea] = instrCB__SET_5_D
	/* SET 5,E */
	OpcodesMap[SHIFT_0xCB+0xeb] = instrCB__SET_5_E
	/* SET 5,H */
	OpcodesMap[SHIFT_0xCB+0xec] = instrCB__SET_5_H
	/* SET 5,L */
	OpcodesMap[SHIFT_0xCB+0xed] = instrCB__SET_5_L
	/* SET 5,(HL) */
	OpcodesMap[SHIFT_0xCB+0xee] = instrCB__SET_5_iHL
	/* SET 5,A */
	OpcodesMap[SHIFT_0xCB+0xef] = instrCB__SET_5_A
	/* SET 6,B */
	OpcodesMap[SHIFT_0xCB+0xf0] = instrCB__SET_6_B
	/* SET 6,C */
	OpcodesMap[SHIFT_0xCB+0xf1] = instrCB__SET_6_C
	/* SET 6,D */
	OpcodesMap[SHIFT_0xCB+0xf2] = instrCB__SET_6_D
	/* SET 6,E */
	OpcodesMap[SHIFT_0xCB+0xf3] = instrCB__SET_6_E
	/* SET 6,H */
	OpcodesMap[SHIFT_0xCB+0xf4] = instrCB__SET_6_H
	/* SET 6,L */
	OpcodesMap[SHIFT_0xCB+0xf5] = instrCB__SET_6_L
	/* SET 6,(HL) */
	OpcodesMap[SHIFT_0xCB+0xf6] = instrCB__SET_6_iHL
	/* SET 6,A */
	OpcodesMap[SHIFT_0xCB+0xf7] = instrCB__SET_6_A
	/* SET 7,B */
	OpcodesMap[SHIFT_0xCB+0xf8] = instrCB__SET_7_B
	/* SET 7,C */
	OpcodesMap[SHIFT_0xCB+0xf9] = instrCB__SET_7_C
	/* SET 7,D */
	OpcodesMap[SHIFT_0xCB+0xfa] = instrCB__SET_7_D
	/* SET 7,E */
	OpcodesMap[SHIFT_0xCB+0xfb] = instrCB__SET_7_E
	/* SET 7,H */
	OpcodesMap[SHIFT_0xCB+0xfc] = instrCB__SET_7_H
	/* SET 7,L */
	OpcodesMap[SHIFT_0xCB+0xfd] = instrCB__SET_7_L
	/* SET 7,(HL) */
	OpcodesMap[SHIFT_0xCB+0xfe] = instrCB__SET_7_iHL
	/* SET 7,A */
	OpcodesMap[SHIFT_0xCB+0xff] = instrCB__SET_7_A

	// END of 0xcb shifted opcodes
}

/* NOP */
func instr__NOP(z80 *Z80) {
}

/* LD BC,nnnn */
func instr__LD_BC_NNNN(z80 *Z80) {
	b1 := z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	b2 := z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.SetBC(joinBytes(b2, b1))
}

/* LD (BC),A */
func instr__LD_iBC_A(z80 *Z80) {
	z80.memory.WriteByte(z80.BC(), z80.A)
}

/* INC BC */
func instr__INC_BC(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.IncBC()
}

/* INC B */
func instr__INC_B(z80 *Z80) {
	z80.incB()
}

/* DEC B */
func instr__DEC_B(z80 *Z80) {
	z80.decB()
}

/* LD B,nn */
func instr__LD_B_NN(z80 *Z80) {
	z80.B = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
}

/* RLCA */
func instr__RLCA(z80 *Z80) {
	z80.A = (z80.A << 1) | (z80.A >> 7)
	z80.F = (z80.F & (FLAG_P | FLAG_Z | FLAG_S)) |
		(z80.A & (FLAG_C | FLAG_3 | FLAG_5))
}

/* ADD HL,BC */
func instr__ADD_HL_BC(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
	z80.add16(z80.hl, z80.BC())
}

/* LD A,(BC) */
func instr__LD_A_iBC(z80 *Z80) {
	z80.A = z80.memory.ReadByte(z80.BC())
}

/* DEC BC */
func instr__DEC_BC(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.DecBC()
}

/* INC C */
func instr__INC_C(z80 *Z80) {
	z80.incC()
}

/* DEC C */
func instr__DEC_C(z80 *Z80) {
	z80.decC()
}

/* LD C,nn */
func instr__LD_C_NN(z80 *Z80) {
	z80.C = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
}

/* RRCA */
func instr__RRCA(z80 *Z80) {
	z80.F = (z80.F & (FLAG_P | FLAG_Z | FLAG_S)) | (z80.A & FLAG_C)
	z80.A = (z80.A >> 1) | (z80.A << 7)
	z80.F |= (z80.A & (FLAG_3 | FLAG_5))
}

/* LD DE,nnnn */
func instr__LD_DE_NNNN(z80 *Z80) {
	b1 := z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	b2 := z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.SetDE(joinBytes(b2, b1))
}

/* LD (DE),A */
func instr__LD_iDE_A(z80 *Z80) {
	z80.memory.WriteByte(z80.DE(), z80.A)
}

/* INC DE */
func instr__INC_DE(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.IncDE()
}

/* INC D */
func instr__INC_D(z80 *Z80) {
	z80.incD()
}

/* DEC D */
func instr__DEC_D(z80 *Z80) {
	z80.decD()
}

/* LD D,nn */
func instr__LD_D_NN(z80 *Z80) {
	z80.D = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
}

/* RLA */
func instr__RLA(z80 *Z80) {
	var bytetemp byte = z80.A
	z80.A = (z80.A << 1) | (z80.F & FLAG_C)
	z80.F = (z80.F & (FLAG_P | FLAG_Z | FLAG_S)) | (z80.A & (FLAG_3 | FLAG_5)) | (bytetemp >> 7)
}

/* JR offset */
func instr__JR_OFFSET(z80 *Z80) {
	z80.jr()
	z80.IncPC(1)
}

/* ADD HL,DE */
func instr__ADD_HL_DE(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
	z80.add16(z80.hl, z80.DE())
}

/* LD A,(DE) */
func instr__LD_A_iDE(z80 *Z80) {
	z80.A = z80.memory.ReadByte(z80.DE())
}

/* DEC DE */
func instr__DEC_DE(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.DecDE()
}

/* INC E */
func instr__INC_E(z80 *Z80) {
	z80.incE()
}

/* DEC E */
func instr__DEC_E(z80 *Z80) {
	z80.decE()
}

/* LD E,nn */
func instr__LD_E_NN(z80 *Z80) {
	z80.E = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
}

/* RRA */
func instr__RRA(z80 *Z80) {
	var bytetemp byte = z80.A
	z80.A = (z80.A >> 1) | (z80.F << 7)
	z80.F = (z80.F & (FLAG_P | FLAG_Z | FLAG_S)) | (z80.A & (FLAG_3 | FLAG_5)) | (bytetemp & FLAG_C)
}

/* JR NZ,offset */
func instr__JR_NZ_OFFSET(z80 *Z80) {
	if (z80.F & FLAG_Z) == 0 {
		z80.jr()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
	}
	z80.IncPC(1)
}

/* LD HL,nnnn */
func instr__LD_HL_NNNN(z80 *Z80) {
	b1 := z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	b2 := z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.SetHL(joinBytes(b2, b1))
}

/* LDI (HL),A */
func instr__LDI_iHL_A(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.A)
	z80.IncHL()
}

/* INC HL */
func instr__INC_HL(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.IncHL()
}

/* INC H */
func instr__INC_H(z80 *Z80) {
	z80.incH()
}

/* DEC H */
func instr__DEC_H(z80 *Z80) {
	z80.decH()
}

/* LD H,nn */
func instr__LD_H_NN(z80 *Z80) {
	z80.H = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
}

/* DAA */
func instr__DAA(z80 *Z80) {
	var add, carry byte = 0, (z80.F & FLAG_C)
	if ((z80.F & FLAG_H) != 0) || ((z80.A & 0x0f) > 9) {
		add = 6
	}
	if (carry != 0) || (z80.A > 0x99) {
		add |= 0x60
	}
	if z80.A > 0x99 {
		carry = FLAG_C
	}
	if (z80.F & FLAG_N) != 0 {
		z80.sub(add)
	} else {
		z80.add(add)
	}
	var temp byte = byte(int(z80.F) & ^(FLAG_C|FLAG_P)) | carry | parityTable[z80.A]
	z80.F = temp
}

/* JR Z,offset */
func instr__JR_Z_OFFSET(z80 *Z80) {
	if (z80.F & FLAG_Z) != 0 {
		z80.jr()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
	}
	z80.IncPC(1)
}

/* ADD HL,HL */
func instr__ADD_HL_HL(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
	z80.add16(z80.hl, z80.HL())
}

/* LDI A,(HL) */
func instr__LDI_A_iHL(z80 *Z80) {
	val := z80.memory.ReadByte(z80.HL())
	z80.memory.WriteByte(uint16(val&0xFF), z80.A)
	z80.DecHL()
}

/* DEC HL */
func instr__DEC_HL(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.DecHL()
}

/* INC L */
func instr__INC_L(z80 *Z80) {
	z80.incL()
}

/* DEC L */
func instr__DEC_L(z80 *Z80) {
	z80.decL()
}

/* LD L,nn */
func instr__LD_L_NN(z80 *Z80) {
	z80.L = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
}

/* CPL */
func instr__CPL(z80 *Z80) {
	z80.A ^= 0xff
	z80.F = (z80.F & (FLAG_C | FLAG_P | FLAG_Z | FLAG_S)) |
		(z80.A & (FLAG_3 | FLAG_5)) |
		(FLAG_N | FLAG_H)
}

/* JR NC,offset */
func instr__JR_NC_OFFSET(z80 *Z80) {
	if (z80.F & FLAG_C) == 0 {
		z80.jr()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
	}
	z80.IncPC(1)
}

/* LD SP,nnnn */
func instr__LD_SP_NNNN(z80 *Z80) {
	b1 := z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	b2 := z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.SetSP(joinBytes(b2, b1))
}

/* LDD (HL),A */
func instr__LDD_iHL_A(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.A)
	z80.DecHL()
}

/* LD (nnnn),A */
func instr__LD_iNNNN_A(z80 *Z80) {
	var wordtemp uint16 = uint16(z80.memory.ReadByte(z80.PC()))
	z80.IncPC(1)
	wordtemp |= uint16(z80.memory.ReadByte(z80.PC())) << 8
	z80.IncPC(1)
	z80.memory.WriteByte(wordtemp, z80.A)
}

/* LD (nnnn),SP */
func instr__LD_iNNNN_SP(z80 *Z80) {
	var wordtemp uint16 = uint16(z80.memory.ReadByte(z80.PC()))
	z80.IncPC(1)
	wordtemp |= uint16(z80.memory.ReadByte(z80.PC())) << 8
	z80.IncPC(1)
	z80.memory.WriteByte(wordtemp, byte(z80.SP()&0xFF))
	z80.memory.WriteByte(wordtemp+1, byte((z80.SP()>>8)&0xFF))
}

/* INC SP */
func instr__INC_SP(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.IncSP()
}

/* INC (HL) */
func instr__INC_iHL(z80 *Z80) {
	{
		var bytetemp byte = z80.memory.ReadByte(z80.HL())
		z80.memory.ContendReadNoMreq(z80.HL(), 1)
		z80.inc(&bytetemp)
		z80.memory.WriteByte(z80.HL(), bytetemp)
	}
}

/* DEC (HL) */
func instr__DEC_iHL(z80 *Z80) {
	{
		var bytetemp byte = z80.memory.ReadByte(z80.HL())
		z80.memory.ContendReadNoMreq(z80.HL(), 1)
		z80.dec(&bytetemp)
		z80.memory.WriteByte(z80.HL(), bytetemp)
	}
}

/* LD (HL),nn */
func instr__LD_iHL_NN(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.memory.ReadByte(z80.PC()))
	z80.IncPC(1)
}

/* SCF */
func instr__SCF(z80 *Z80) {
	z80.F = (z80.F & (FLAG_P | FLAG_Z | FLAG_S)) |
		(z80.A & (FLAG_3 | FLAG_5)) |
		FLAG_C
}

/* JR C,offset */
func instr__JR_C_OFFSET(z80 *Z80) {
	if (z80.F & FLAG_C) != 0 {
		z80.jr()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
	}
	z80.IncPC(1)
}

/* ADD HL,SP */
func instr__ADD_HL_SP(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
	z80.add16(z80.hl, z80.SP())
}

/* LDD A,(HL) */
func instr__LDD_A_iHL(z80 *Z80) {
	val := z80.memory.ReadByte(z80.HL())
	z80.memory.WriteByte(uint16(val&0xFF), z80.A)
	z80.DecHL()
}

/* LD A,(nnnn) */
func instr__LD_A_iNNNN(z80 *Z80) {
	var wordtemp uint16 = uint16(z80.memory.ReadByte(z80.PC()))
	z80.IncPC(1)
	wordtemp |= uint16(z80.memory.ReadByte(z80.PC())) << 8
	z80.IncPC(1)
	z80.A = z80.memory.ReadByte(wordtemp)
}

/* DEC SP */
func instr__DEC_SP(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.DecSP()
}

/* INC A */
func instr__INC_A(z80 *Z80) {
	z80.incA()
}

/* DEC A */
func instr__DEC_A(z80 *Z80) {
	z80.decA()
}

/* LD A,nn */
func instr__LD_A_NN(z80 *Z80) {
	z80.A = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
}

/* CCF */
func instr__CCF(z80 *Z80) {
	z80.F = (z80.F & (FLAG_P | FLAG_Z | FLAG_S)) |
		ternOpB((z80.F&FLAG_C) != 0, FLAG_H, FLAG_C) |
		(z80.A & (FLAG_3 | FLAG_5))
}

/* LD B,B */
func instr__LD_B_B(z80 *Z80) {
}

/* LD B,C */
func instr__LD_B_C(z80 *Z80) {
	z80.B = z80.C
}

/* LD B,D */
func instr__LD_B_D(z80 *Z80) {
	z80.B = z80.D
}

/* LD B,E */
func instr__LD_B_E(z80 *Z80) {
	z80.B = z80.E
}

/* LD B,H */
func instr__LD_B_H(z80 *Z80) {
	z80.B = z80.H
}

/* LD B,L */
func instr__LD_B_L(z80 *Z80) {
	z80.B = z80.L
}

/* LD B,(HL) */
func instr__LD_B_iHL(z80 *Z80) {
	z80.B = z80.memory.ReadByte(z80.HL())
}

/* LD B,A */
func instr__LD_B_A(z80 *Z80) {
	z80.B = z80.A
}

/* LD C,B */
func instr__LD_C_B(z80 *Z80) {
	z80.C = z80.B
}

/* LD C,C */
func instr__LD_C_C(z80 *Z80) {
}

/* LD C,D */
func instr__LD_C_D(z80 *Z80) {
	z80.C = z80.D
}

/* LD C,E */
func instr__LD_C_E(z80 *Z80) {
	z80.C = z80.E
}

/* LD C,H */
func instr__LD_C_H(z80 *Z80) {
	z80.C = z80.H
}

/* LD C,L */
func instr__LD_C_L(z80 *Z80) {
	z80.C = z80.L
}

/* LD C,(HL) */
func instr__LD_C_iHL(z80 *Z80) {
	z80.C = z80.memory.ReadByte(z80.HL())
}

/* LD C,A */
func instr__LD_C_A(z80 *Z80) {
	z80.C = z80.A
}

/* LD D,B */
func instr__LD_D_B(z80 *Z80) {
	z80.D = z80.B
}

/* LD D,C */
func instr__LD_D_C(z80 *Z80) {
	z80.D = z80.C
}

/* LD D,D */
func instr__LD_D_D(z80 *Z80) {
}

/* LD D,E */
func instr__LD_D_E(z80 *Z80) {
	z80.D = z80.E
}

/* LD D,H */
func instr__LD_D_H(z80 *Z80) {
	z80.D = z80.H
}

/* LD D,L */
func instr__LD_D_L(z80 *Z80) {
	z80.D = z80.L
}

/* LD D,(HL) */
func instr__LD_D_iHL(z80 *Z80) {
	z80.D = z80.memory.ReadByte(z80.HL())
}

/* LD D,A */
func instr__LD_D_A(z80 *Z80) {
	z80.D = z80.A
}

/* LD E,B */
func instr__LD_E_B(z80 *Z80) {
	z80.E = z80.B
}

/* LD E,C */
func instr__LD_E_C(z80 *Z80) {
	z80.E = z80.C
}

/* LD E,D */
func instr__LD_E_D(z80 *Z80) {
	z80.E = z80.D
}

/* LD E,E */
func instr__LD_E_E(z80 *Z80) {
}

/* LD E,H */
func instr__LD_E_H(z80 *Z80) {
	z80.E = z80.H
}

/* LD E,L */
func instr__LD_E_L(z80 *Z80) {
	z80.E = z80.L
}

/* LD E,(HL) */
func instr__LD_E_iHL(z80 *Z80) {
	z80.E = z80.memory.ReadByte(z80.HL())
}

/* LD E,A */
func instr__LD_E_A(z80 *Z80) {
	z80.E = z80.A
}

/* LD H,B */
func instr__LD_H_B(z80 *Z80) {
	z80.H = z80.B
}

/* LD H,C */
func instr__LD_H_C(z80 *Z80) {
	z80.H = z80.C
}

/* LD H,D */
func instr__LD_H_D(z80 *Z80) {
	z80.H = z80.D
}

/* LD H,E */
func instr__LD_H_E(z80 *Z80) {
	z80.H = z80.E
}

/* LD H,H */
func instr__LD_H_H(z80 *Z80) {
}

/* LD H,L */
func instr__LD_H_L(z80 *Z80) {
	z80.H = z80.L
}

/* LD H,(HL) */
func instr__LD_H_iHL(z80 *Z80) {
	z80.H = z80.memory.ReadByte(z80.HL())
}

/* LD H,A */
func instr__LD_H_A(z80 *Z80) {
	z80.H = z80.A
}

/* LD L,B */
func instr__LD_L_B(z80 *Z80) {
	z80.L = z80.B
}

/* LD L,C */
func instr__LD_L_C(z80 *Z80) {
	z80.L = z80.C
}

/* LD L,D */
func instr__LD_L_D(z80 *Z80) {
	z80.L = z80.D
}

/* LD L,E */
func instr__LD_L_E(z80 *Z80) {
	z80.L = z80.E
}

/* LD L,H */
func instr__LD_L_H(z80 *Z80) {
	z80.L = z80.H
}

/* LD L,L */
func instr__LD_L_L(z80 *Z80) {
}

/* LD L,(HL) */
func instr__LD_L_iHL(z80 *Z80) {
	z80.L = z80.memory.ReadByte(z80.HL())
}

/* LD L,A */
func instr__LD_L_A(z80 *Z80) {
	z80.L = z80.A
}

/* LD (HL),B */
func instr__LD_iHL_B(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.B)
}

/* LD (HL),C */
func instr__LD_iHL_C(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.C)
}

/* LD (HL),D */
func instr__LD_iHL_D(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.D)
}

/* LD (HL),E */
func instr__LD_iHL_E(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.E)
}

/* LD (HL),H */
func instr__LD_iHL_H(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.H)
}

/* LD (HL),L */
func instr__LD_iHL_L(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.L)
}

/* HALT */
func instr__HALT(z80 *Z80) {
	z80.Halted = true
	z80.DecPC(1)
	return
}

/* STOP */
func instr__STOP(z80 *Z80) {
	// TODO
	panic("STOP: Unfinished opcode!")
}

/* LD (HL),A */
func instr__LD_iHL_A(z80 *Z80) {
	z80.memory.WriteByte(z80.HL(), z80.A)
}

/* LD A,B */
func instr__LD_A_B(z80 *Z80) {
	z80.A = z80.B
}

/* LD A,C */
func instr__LD_A_C(z80 *Z80) {
	z80.A = z80.C
}

/* LD A,D */
func instr__LD_A_D(z80 *Z80) {
	z80.A = z80.D
}

/* LD A,E */
func instr__LD_A_E(z80 *Z80) {
	z80.A = z80.E
}

/* LD A,H */
func instr__LD_A_H(z80 *Z80) {
	z80.A = z80.H
}

/* LD A,L */
func instr__LD_A_L(z80 *Z80) {
	z80.A = z80.L
}

/* LD A,(HL) */
func instr__LD_A_iHL(z80 *Z80) {
	z80.A = z80.memory.ReadByte(z80.HL())
}

/* LD A,A */
func instr__LD_A_A(z80 *Z80) {
}

/* ADD A,B */
func instr__ADD_A_B(z80 *Z80) {
	z80.add(z80.B)
}

/* ADD A,C */
func instr__ADD_A_C(z80 *Z80) {
	z80.add(z80.C)
}

/* ADD A,D */
func instr__ADD_A_D(z80 *Z80) {
	z80.add(z80.D)
}

/* ADD A,E */
func instr__ADD_A_E(z80 *Z80) {
	z80.add(z80.E)
}

/* ADD A,H */
func instr__ADD_A_H(z80 *Z80) {
	z80.add(z80.H)
}

/* ADD A,L */
func instr__ADD_A_L(z80 *Z80) {
	z80.add(z80.L)
}

/* ADD A,(HL) */
func instr__ADD_A_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())

	z80.add(bytetemp)
}

/* ADD A,A */
func instr__ADD_A_A(z80 *Z80) {
	z80.add(z80.A)
}

/* ADC A,B */
func instr__ADC_A_B(z80 *Z80) {
	z80.adc(z80.B)
}

/* ADC A,C */
func instr__ADC_A_C(z80 *Z80) {
	z80.adc(z80.C)
}

/* ADC A,D */
func instr__ADC_A_D(z80 *Z80) {
	z80.adc(z80.D)
}

/* ADC A,E */
func instr__ADC_A_E(z80 *Z80) {
	z80.adc(z80.E)
}

/* ADC A,H */
func instr__ADC_A_H(z80 *Z80) {
	z80.adc(z80.H)
}

/* ADC A,L */
func instr__ADC_A_L(z80 *Z80) {
	z80.adc(z80.L)
}

/* ADC A,(HL) */
func instr__ADC_A_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())

	z80.adc(bytetemp)
}

/* ADC A,A */
func instr__ADC_A_A(z80 *Z80) {
	z80.adc(z80.A)
}

/* SUB A,B */
func instr__SUB_A_B(z80 *Z80) {
	z80.sub(z80.B)
}

/* SUB A,C */
func instr__SUB_A_C(z80 *Z80) {
	z80.sub(z80.C)
}

/* SUB A,D */
func instr__SUB_A_D(z80 *Z80) {
	z80.sub(z80.D)
}

/* SUB A,E */
func instr__SUB_A_E(z80 *Z80) {
	z80.sub(z80.E)
}

/* SUB A,H */
func instr__SUB_A_H(z80 *Z80) {
	z80.sub(z80.H)
}

/* SUB A,L */
func instr__SUB_A_L(z80 *Z80) {
	z80.sub(z80.L)
}

/* SUB A,(HL) */
func instr__SUB_A_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())

	z80.sub(bytetemp)
}

/* SUB A,A */
func instr__SUB_A_A(z80 *Z80) {
	z80.sub(z80.A)
}

/* SBC A,B */
func instr__SBC_A_B(z80 *Z80) {
	z80.sbc(z80.B)
}

/* SBC A,C */
func instr__SBC_A_C(z80 *Z80) {
	z80.sbc(z80.C)
}

/* SBC A,D */
func instr__SBC_A_D(z80 *Z80) {
	z80.sbc(z80.D)
}

/* SBC A,E */
func instr__SBC_A_E(z80 *Z80) {
	z80.sbc(z80.E)
}

/* SBC A,H */
func instr__SBC_A_H(z80 *Z80) {
	z80.sbc(z80.H)
}

/* SBC A,L */
func instr__SBC_A_L(z80 *Z80) {
	z80.sbc(z80.L)
}

/* SBC A,(HL) */
func instr__SBC_A_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())

	z80.sbc(bytetemp)
}

/* SBC A,A */
func instr__SBC_A_A(z80 *Z80) {
	z80.sbc(z80.A)
}

/* AND A,B */
func instr__AND_A_B(z80 *Z80) {
	z80.and(z80.B)
}

/* AND A,C */
func instr__AND_A_C(z80 *Z80) {
	z80.and(z80.C)
}

/* AND A,D */
func instr__AND_A_D(z80 *Z80) {
	z80.and(z80.D)
}

/* AND A,E */
func instr__AND_A_E(z80 *Z80) {
	z80.and(z80.E)
}

/* AND A,H */
func instr__AND_A_H(z80 *Z80) {
	z80.and(z80.H)
}

/* AND A,L */
func instr__AND_A_L(z80 *Z80) {
	z80.and(z80.L)
}

/* AND A,(HL) */
func instr__AND_A_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())

	z80.and(bytetemp)
}

/* AND A,A */
func instr__AND_A_A(z80 *Z80) {
	z80.and(z80.A)
}

/* XOR A,B */
func instr__XOR_A_B(z80 *Z80) {
	z80.xor(z80.B)
}

/* XOR A,C */
func instr__XOR_A_C(z80 *Z80) {
	z80.xor(z80.C)
}

/* XOR A,D */
func instr__XOR_A_D(z80 *Z80) {
	z80.xor(z80.D)
}

/* XOR A,E */
func instr__XOR_A_E(z80 *Z80) {
	z80.xor(z80.E)
}

/* XOR A,H */
func instr__XOR_A_H(z80 *Z80) {
	z80.xor(z80.H)
}

/* XOR A,L */
func instr__XOR_A_L(z80 *Z80) {
	z80.xor(z80.L)
}

/* XOR A,(HL) */
func instr__XOR_A_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())

	z80.xor(bytetemp)
}

/* XOR A,A */
func instr__XOR_A_A(z80 *Z80) {
	z80.xor(z80.A)
}

/* OR A,B */
func instr__OR_A_B(z80 *Z80) {
	z80.or(z80.B)
}

/* OR A,C */
func instr__OR_A_C(z80 *Z80) {
	z80.or(z80.C)
}

/* OR A,D */
func instr__OR_A_D(z80 *Z80) {
	z80.or(z80.D)
}

/* OR A,E */
func instr__OR_A_E(z80 *Z80) {
	z80.or(z80.E)
}

/* OR A,H */
func instr__OR_A_H(z80 *Z80) {
	z80.or(z80.H)
}

/* OR A,L */
func instr__OR_A_L(z80 *Z80) {
	z80.or(z80.L)
}

/* OR A,(HL) */
func instr__OR_A_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())

	z80.or(bytetemp)
}

/* OR A,A */
func instr__OR_A_A(z80 *Z80) {
	z80.or(z80.A)
}

/* CP B */
func instr__CP_B(z80 *Z80) {
	z80.cp(z80.B)
}

/* CP C */
func instr__CP_C(z80 *Z80) {
	z80.cp(z80.C)
}

/* CP D */
func instr__CP_D(z80 *Z80) {
	z80.cp(z80.D)
}

/* CP E */
func instr__CP_E(z80 *Z80) {
	z80.cp(z80.E)
}

/* CP H */
func instr__CP_H(z80 *Z80) {
	z80.cp(z80.H)
}

/* CP L */
func instr__CP_L(z80 *Z80) {
	z80.cp(z80.L)
}

/* CP (HL) */
func instr__CP_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())

	z80.cp(bytetemp)
}

/* CP A */
func instr__CP_A(z80 *Z80) {
	z80.cp(z80.A)
}

/* RET NZ */
func instr__RET_NZ(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	if !((z80.F & FLAG_Z) != 0) {
		z80.ret()
	}
}

/* POP BC */
func instr__POP_BC(z80 *Z80) {
	z80.C, z80.B = z80.pop16()
}

/* JP NZ,nnnn */
func instr__JP_NZ_NNNN(z80 *Z80) {
	if (z80.F & FLAG_Z) == 0 {
		z80.jp()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
		z80.memory.ContendRead(z80.PC()+1, 3)
		z80.IncPC(2)
	}
}

/* JP nnnn */
func instr__JP_NNNN(z80 *Z80) {
	z80.jp()
}

/* CALL NZ,nnnn */
func instr__CALL_NZ_NNNN(z80 *Z80) {
	if (z80.F & FLAG_Z) == 0 {
		z80.call()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
		z80.memory.ContendRead(z80.PC()+1, 3)
		z80.IncPC(2)
	}
}

/* PUSH BC */
func instr__PUSH_BC(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.push16(z80.C, z80.B)
}

/* ADD A,nn */
func instr__ADD_A_NN(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.add(bytetemp)
}

/* RST 00 */
func instr__RST_00(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.rst(0x00)
}

/* RET Z */
func instr__RET_Z(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	if (z80.F & FLAG_Z) != 0 {
		z80.ret()
	}
}

/* RET */
func instr__RET(z80 *Z80) {
	z80.ret()
}

/* JP Z,nnnn */
func instr__JP_Z_NNNN(z80 *Z80) {
	if (z80.F & FLAG_Z) != 0 {
		z80.jp()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
		z80.memory.ContendRead(z80.PC()+1, 3)
		z80.IncPC(2)
	}
}

/* shift CB */
func instr__SHIFT_CB(z80 *Z80) {
}

/* CALL Z,nnnn */
func instr__CALL_Z_NNNN(z80 *Z80) {
	if (z80.F & FLAG_Z) != 0 {
		z80.call()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
		z80.memory.ContendRead(z80.PC()+1, 3)
		z80.IncPC(2)
	}
}

/* CALL nnnn */
func instr__CALL_NNNN(z80 *Z80) {
	z80.call()
}

/* ADC A,nn */
func instr__ADC_A_NN(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.adc(bytetemp)
}

/* RST 8 */
func instr__RST_8(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.rst(0x8)
}

/* RET NC */
func instr__RET_NC(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	if !((z80.F & FLAG_C) != 0) {
		z80.ret()
	}
}

/* POP DE */
func instr__POP_DE(z80 *Z80) {
	z80.E, z80.D = z80.pop16()
}

/* JP NC,nnnn */
func instr__JP_NC_NNNN(z80 *Z80) {
	if (z80.F & FLAG_C) == 0 {
		z80.jp()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
		z80.memory.ContendRead(z80.PC()+1, 3)
		z80.IncPC(2)
	}
}

/* OUT (nn),A */
func instr__OUT_iNN_A(z80 *Z80) {
	var outtemp uint16 = uint16(z80.memory.ReadByte(z80.PC())) + (uint16(z80.A) << 8)
	z80.IncPC(1)
	z80.writePort(outtemp, z80.A)
}

/* CALL NC,nnnn */
func instr__CALL_NC_NNNN(z80 *Z80) {
	if (z80.F & FLAG_C) == 0 {
		z80.call()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
		z80.memory.ContendRead(z80.PC()+1, 3)
		z80.IncPC(2)
	}
}

/* PUSH DE */
func instr__PUSH_DE(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.push16(z80.E, z80.D)
}

/* SUB nn */
func instr__SUB_NN(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.sub(bytetemp)
}

/* RST 10 */
func instr__RST_10(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.rst(0x10)
}

/* RET C */
func instr__RET_C(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	if (z80.F & FLAG_C) != 0 {
		z80.ret()
	}
}

/* RETI */
func instr__RETI(z80 *Z80) {
    z80.IME = 1
	z80.interruptsEnabledAt = int(z80.Tstates)
	z80.ret()
}

/* JP C,nnnn */
func instr__JP_C_NNNN(z80 *Z80) {
	if (z80.F & FLAG_C) != 0 {
		z80.jp()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
		z80.memory.ContendRead(z80.PC()+1, 3)
		z80.IncPC(2)
	}
}

/* CALL C,nnnn */
func instr__CALL_C_NNNN(z80 *Z80) {
	if (z80.F & FLAG_C) != 0 {
		z80.call()
	} else {
		z80.memory.ContendRead(z80.PC(), 3)
		z80.memory.ContendRead(z80.PC()+1, 3)
		z80.IncPC(2)
	}
}

/* SBC A,nn */
func instr__SBC_A_NN(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.sbc(bytetemp)
}

/* RST 18 */
func instr__RST_18(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.rst(0x18)
}

/* LD (FF00+n),A */
func instr__LD_iFF00N_A(z80 *Z80) {
	addr := z80.memory.ReadByte(z80.PC())
	z80.memory.WriteByte(0xFF00+uint16(addr), z80.A)
	z80.IncPC(1)
}

/* POP HL */
func instr__POP_HL(z80 *Z80) {
	z80.L, z80.H = z80.pop16()
}

/* LD (FF00+C),A */
func instr__LD_iFF00C_A(z80 *Z80) {
	z80.memory.WriteByte(0xFF00+uint16(z80.C)&0xFFFF, z80.A)
}

/* PUSH HL */
func instr__PUSH_HL(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.push16(z80.L, z80.H)
}

/* AND nn */
func instr__AND_NN(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.and(bytetemp)
}

/* RST 20 */
func instr__RST_20(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.rst(0x20)
}

/* ADD SP,dd */
func instr__ADD_SP_DD(z80 *Z80) {
	v := z80.memory.ReadByte(z80.PC())

	if v > 0x7F {
		v -= (^v + 1) & 0xFF
	}

	z80.IncPC(1)
	z80.SetSP(z80.SP() + uint16(v))
}

/* JP HL */
func instr__JP_HL(z80 *Z80) {
	z80.SetPC(z80.HL()) /* NB: NOT INDIRECT! */
}

/* XOR A,nn */
func instr__XOR_A_NN(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.xor(bytetemp)
}

/* RST 28 */
func instr__RST_28(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.rst(0x28)
}

/* LD A,(FF00+n) */
func instr__LD_A_iFF00N(z80 *Z80) {
	offset := z80.memory.ReadByte(z80.PC())
	z80.A = z80.memory.ReadByte(0xFF00 + uint16(offset))
	z80.IncPC(1)
}

/* POP AF */
func instr__POP_AF(z80 *Z80) {
	z80.F, z80.A = z80.pop16()
}

/* LD, A(FF00+C) */
func instr__LD_A_iFF00C(z80 *Z80) {
	z80.A = z80.memory.ReadByte(0xFF00 + uint16(z80.C))
}

/* DI */
func instr__DI(z80 *Z80) {
	z80.IME = 0
}

/* PUSH AF */
func instr__PUSH_AF(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.push16(z80.F, z80.A)
}

/* OR nn */
func instr__OR_NN(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.or(bytetemp)
}

/* RST 30 */
func instr__RST_30(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.rst(0x30)
}

/* LD HL,SP+dd */
func instr__LD_HL_SPDD(z80 *Z80) {
	// TODO
	panic("LD_HL_SPDD: Unfinished opcode!")
}

/* LD SP,HL */
func instr__LD_SP_HL(z80 *Z80) {
	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
	z80.SetSP(z80.HL())
}

/* EI */
func instr__EI(z80 *Z80) {
	/* Interrupts are not accepted immediately after an EI, but are
	   accepted after the next instruction */
	z80.IME = 1
	z80.interruptsEnabledAt = int(z80.Tstates)
	// eventAdd(z80.Tstates + 1, z80InterruptEvent)
}

/* CP nn */
func instr__CP_NN(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.PC())
	z80.IncPC(1)
	z80.cp(bytetemp)
}

/* RST 38 */
func instr__RST_38(z80 *Z80) {
	z80.memory.ContendReadNoMreq(z80.IR(), 1)
	z80.rst(0x38)
}

/* RLC B */
func instrCB__RLC_B(z80 *Z80) {
	z80.B = z80.rlc(z80.B)
}

/* RLC C */
func instrCB__RLC_C(z80 *Z80) {
	z80.C = z80.rlc(z80.C)
}

/* RLC D */
func instrCB__RLC_D(z80 *Z80) {
	z80.D = z80.rlc(z80.D)
}

/* RLC E */
func instrCB__RLC_E(z80 *Z80) {
	z80.E = z80.rlc(z80.E)
}

/* RLC H */
func instrCB__RLC_H(z80 *Z80) {
	z80.H = z80.rlc(z80.H)
}

/* RLC L */
func instrCB__RLC_L(z80 *Z80) {
	z80.L = z80.rlc(z80.L)
}

/* RLC (HL) */
func instrCB__RLC_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	bytetemp = z80.rlc(bytetemp)
	z80.memory.WriteByte(z80.HL(), bytetemp)
}

/* RLC A */
func instrCB__RLC_A(z80 *Z80) {
	z80.A = z80.rlc(z80.A)
}

/* RRC B */
func instrCB__RRC_B(z80 *Z80) {
	z80.B = z80.rrc(z80.B)
}

/* RRC C */
func instrCB__RRC_C(z80 *Z80) {
	z80.C = z80.rrc(z80.C)
}

/* RRC D */
func instrCB__RRC_D(z80 *Z80) {
	z80.D = z80.rrc(z80.D)
}

/* RRC E */
func instrCB__RRC_E(z80 *Z80) {
	z80.E = z80.rrc(z80.E)
}

/* RRC H */
func instrCB__RRC_H(z80 *Z80) {
	z80.H = z80.rrc(z80.H)
}

/* RRC L */
func instrCB__RRC_L(z80 *Z80) {
	z80.L = z80.rrc(z80.L)
}

/* RRC (HL) */
func instrCB__RRC_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	bytetemp = z80.rrc(bytetemp)
	z80.memory.WriteByte(z80.HL(), bytetemp)
}

/* RRC A */
func instrCB__RRC_A(z80 *Z80) {
	z80.A = z80.rrc(z80.A)
}

/* RL B */
func instrCB__RL_B(z80 *Z80) {
	z80.B = z80.rl(z80.B)
}

/* RL C */
func instrCB__RL_C(z80 *Z80) {
	z80.C = z80.rl(z80.C)
}

/* RL D */
func instrCB__RL_D(z80 *Z80) {
	z80.D = z80.rl(z80.D)
}

/* RL E */
func instrCB__RL_E(z80 *Z80) {
	z80.E = z80.rl(z80.E)
}

/* RL H */
func instrCB__RL_H(z80 *Z80) {
	z80.H = z80.rl(z80.H)
}

/* RL L */
func instrCB__RL_L(z80 *Z80) {
	z80.L = z80.rl(z80.L)
}

/* RL (HL) */
func instrCB__RL_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	bytetemp = z80.rl(bytetemp)
	z80.memory.WriteByte(z80.HL(), bytetemp)
}

/* RL A */
func instrCB__RL_A(z80 *Z80) {
	z80.A = z80.rl(z80.A)
}

/* RR B */
func instrCB__RR_B(z80 *Z80) {
	z80.B = z80.rr(z80.B)
}

/* RR C */
func instrCB__RR_C(z80 *Z80) {
	z80.C = z80.rr(z80.C)
}

/* RR D */
func instrCB__RR_D(z80 *Z80) {
	z80.D = z80.rr(z80.D)
}

/* RR E */
func instrCB__RR_E(z80 *Z80) {
	z80.E = z80.rr(z80.E)
}

/* RR H */
func instrCB__RR_H(z80 *Z80) {
	z80.H = z80.rr(z80.H)
}

/* RR L */
func instrCB__RR_L(z80 *Z80) {
	z80.L = z80.rr(z80.L)
}

/* RR (HL) */
func instrCB__RR_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	bytetemp = z80.rr(bytetemp)
	z80.memory.WriteByte(z80.HL(), bytetemp)
}

/* RR A */
func instrCB__RR_A(z80 *Z80) {
	z80.A = z80.rr(z80.A)
}

/* SLA B */
func instrCB__SLA_B(z80 *Z80) {
	z80.B = z80.sla(z80.B)
}

/* SLA C */
func instrCB__SLA_C(z80 *Z80) {
	z80.C = z80.sla(z80.C)
}

/* SLA D */
func instrCB__SLA_D(z80 *Z80) {
	z80.D = z80.sla(z80.D)
}

/* SLA E */
func instrCB__SLA_E(z80 *Z80) {
	z80.E = z80.sla(z80.E)
}

/* SLA H */
func instrCB__SLA_H(z80 *Z80) {
	z80.H = z80.sla(z80.H)
}

/* SLA L */
func instrCB__SLA_L(z80 *Z80) {
	z80.L = z80.sla(z80.L)
}

/* SLA (HL) */
func instrCB__SLA_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	bytetemp = z80.sla(bytetemp)
	z80.memory.WriteByte(z80.HL(), bytetemp)
}

/* SLA A */
func instrCB__SLA_A(z80 *Z80) {
	z80.A = z80.sla(z80.A)
}

/* SRA B */
func instrCB__SRA_B(z80 *Z80) {
	z80.B = z80.sra(z80.B)
}

/* SRA C */
func instrCB__SRA_C(z80 *Z80) {
	z80.C = z80.sra(z80.C)
}

/* SRA D */
func instrCB__SRA_D(z80 *Z80) {
	z80.D = z80.sra(z80.D)
}

/* SRA E */
func instrCB__SRA_E(z80 *Z80) {
	z80.E = z80.sra(z80.E)
}

/* SRA H */
func instrCB__SRA_H(z80 *Z80) {
	z80.H = z80.sra(z80.H)
}

/* SRA L */
func instrCB__SRA_L(z80 *Z80) {
	z80.L = z80.sra(z80.L)
}

/* SRA (HL) */
func instrCB__SRA_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	bytetemp = z80.sra(bytetemp)
	z80.memory.WriteByte(z80.HL(), bytetemp)
}

/* SRA A */
func instrCB__SRA_A(z80 *Z80) {
	z80.A = z80.sra(z80.A)
}

/* SWAP B */
func instrCB__SWAP_B(z80 *Z80) {
	z80.B = ((z80.B & 0xF) << 4) | ((z80.B & 0xF0) >> 4)
	if z80.B > 0 {
		z80.F = 0
	} else {
		z80.F = 0x80
	}
}

/* SWAP C */
func instrCB__SWAP_C(z80 *Z80) {
	z80.C = ((z80.C & 0xF) << 4) | ((z80.C & 0xF0) >> 4)
	if z80.C > 0 {
		z80.F = 0
	} else {
		z80.F = 0x80
	}
}

/* SWAP D */
func instrCB__SWAP_D(z80 *Z80) {
	z80.D = ((z80.D & 0xF) << 4) | ((z80.D & 0xF0) >> 4)
	if z80.D > 0 {
		z80.F = 0
	} else {
		z80.F = 0x80
	}
}

/* SWAP E */
func instrCB__SWAP_E(z80 *Z80) {
	z80.E = ((z80.E & 0xF) << 4) | ((z80.E & 0xF0) >> 4)
	if z80.E > 0 {
		z80.F = 0
	} else {
		z80.F = 0x80
	}
}

/* SWAP H */
func instrCB__SWAP_H(z80 *Z80) {
	z80.H = ((z80.H & 0xF) << 4) | ((z80.H & 0xF0) >> 4)
	if z80.H > 0 {
		z80.F = 0
	} else {
		z80.F = 0x80
	}
}

/* SWAP L */
func instrCB__SWAP_L(z80 *Z80) {
	z80.L = ((z80.L & 0xF) << 4) | ((z80.L & 0xF0) >> 4)
	if z80.L > 0 {
		z80.F = 0
	} else {
		z80.F = 0x80
	}
}

/* SWAP (HL) */
func instrCB__SWAP_iHL(z80 *Z80) {
	z80.SetHL(((z80.HL() & 0xF) << 4) | ((z80.HL() & 0xF0) >> 4))
	if z80.HL() > 0 {
		z80.F = 0
	} else {
		z80.F = 0x80
	}
}

/* SWAP A */
func instrCB__SWAP_A(z80 *Z80) {
	z80.A = ((z80.A & 0xF) << 4) | ((z80.A & 0xF0) >> 4)
	if z80.A > 0 {
		z80.F = 0
	} else {
		z80.F = 0x80
	}
}

/* SRL B */
func instrCB__SRL_B(z80 *Z80) {
	z80.B = z80.srl(z80.B)
}

/* SRL C */
func instrCB__SRL_C(z80 *Z80) {
	z80.C = z80.srl(z80.C)
}

/* SRL D */
func instrCB__SRL_D(z80 *Z80) {
	z80.D = z80.srl(z80.D)
}

/* SRL E */
func instrCB__SRL_E(z80 *Z80) {
	z80.E = z80.srl(z80.E)
}

/* SRL H */
func instrCB__SRL_H(z80 *Z80) {
	z80.H = z80.srl(z80.H)
}

/* SRL L */
func instrCB__SRL_L(z80 *Z80) {
	z80.L = z80.srl(z80.L)
}

/* SRL (HL) */
func instrCB__SRL_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	bytetemp = z80.srl(bytetemp)
	z80.memory.WriteByte(z80.HL(), bytetemp)
}

/* SRL A */
func instrCB__SRL_A(z80 *Z80) {
	z80.A = z80.srl(z80.A)
}

/* BIT 0,B */
func instrCB__BIT_0_B(z80 *Z80) {
	z80.bit(0, z80.B)
}

/* BIT 0,C */
func instrCB__BIT_0_C(z80 *Z80) {
	z80.bit(0, z80.C)
}

/* BIT 0,D */
func instrCB__BIT_0_D(z80 *Z80) {
	z80.bit(0, z80.D)
}

/* BIT 0,E */
func instrCB__BIT_0_E(z80 *Z80) {
	z80.bit(0, z80.E)
}

/* BIT 0,H */
func instrCB__BIT_0_H(z80 *Z80) {
	z80.bit(0, z80.H)
}

/* BIT 0,L */
func instrCB__BIT_0_L(z80 *Z80) {
	z80.bit(0, z80.L)
}

/* BIT 0,(HL) */
func instrCB__BIT_0_iHL(z80 *Z80) {
	bytetemp := z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.bit(0, bytetemp)
}

/* BIT 0,A */
func instrCB__BIT_0_A(z80 *Z80) {
	z80.bit(0, z80.A)
}

/* BIT 1,B */
func instrCB__BIT_1_B(z80 *Z80) {
	z80.bit(1, z80.B)
}

/* BIT 1,C */
func instrCB__BIT_1_C(z80 *Z80) {
	z80.bit(1, z80.C)
}

/* BIT 1,D */
func instrCB__BIT_1_D(z80 *Z80) {
	z80.bit(1, z80.D)
}

/* BIT 1,E */
func instrCB__BIT_1_E(z80 *Z80) {
	z80.bit(1, z80.E)
}

/* BIT 1,H */
func instrCB__BIT_1_H(z80 *Z80) {
	z80.bit(1, z80.H)
}

/* BIT 1,L */
func instrCB__BIT_1_L(z80 *Z80) {
	z80.bit(1, z80.L)
}

/* BIT 1,(HL) */
func instrCB__BIT_1_iHL(z80 *Z80) {
	bytetemp := z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.bit(1, bytetemp)
}

/* BIT 1,A */
func instrCB__BIT_1_A(z80 *Z80) {
	z80.bit(1, z80.A)
}

/* BIT 2,B */
func instrCB__BIT_2_B(z80 *Z80) {
	z80.bit(2, z80.B)
}

/* BIT 2,C */
func instrCB__BIT_2_C(z80 *Z80) {
	z80.bit(2, z80.C)
}

/* BIT 2,D */
func instrCB__BIT_2_D(z80 *Z80) {
	z80.bit(2, z80.D)
}

/* BIT 2,E */
func instrCB__BIT_2_E(z80 *Z80) {
	z80.bit(2, z80.E)
}

/* BIT 2,H */
func instrCB__BIT_2_H(z80 *Z80) {
	z80.bit(2, z80.H)
}

/* BIT 2,L */
func instrCB__BIT_2_L(z80 *Z80) {
	z80.bit(2, z80.L)
}

/* BIT 2,(HL) */
func instrCB__BIT_2_iHL(z80 *Z80) {
	bytetemp := z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.bit(2, bytetemp)
}

/* BIT 2,A */
func instrCB__BIT_2_A(z80 *Z80) {
	z80.bit(2, z80.A)
}

/* BIT 3,B */
func instrCB__BIT_3_B(z80 *Z80) {
	z80.bit(3, z80.B)
}

/* BIT 3,C */
func instrCB__BIT_3_C(z80 *Z80) {
	z80.bit(3, z80.C)
}

/* BIT 3,D */
func instrCB__BIT_3_D(z80 *Z80) {
	z80.bit(3, z80.D)
}

/* BIT 3,E */
func instrCB__BIT_3_E(z80 *Z80) {
	z80.bit(3, z80.E)
}

/* BIT 3,H */
func instrCB__BIT_3_H(z80 *Z80) {
	z80.bit(3, z80.H)
}

/* BIT 3,L */
func instrCB__BIT_3_L(z80 *Z80) {
	z80.bit(3, z80.L)
}

/* BIT 3,(HL) */
func instrCB__BIT_3_iHL(z80 *Z80) {
	bytetemp := z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.bit(3, bytetemp)
}

/* BIT 3,A */
func instrCB__BIT_3_A(z80 *Z80) {
	z80.bit(3, z80.A)
}

/* BIT 4,B */
func instrCB__BIT_4_B(z80 *Z80) {
	z80.bit(4, z80.B)
}

/* BIT 4,C */
func instrCB__BIT_4_C(z80 *Z80) {
	z80.bit(4, z80.C)
}

/* BIT 4,D */
func instrCB__BIT_4_D(z80 *Z80) {
	z80.bit(4, z80.D)
}

/* BIT 4,E */
func instrCB__BIT_4_E(z80 *Z80) {
	z80.bit(4, z80.E)
}

/* BIT 4,H */
func instrCB__BIT_4_H(z80 *Z80) {
	z80.bit(4, z80.H)
}

/* BIT 4,L */
func instrCB__BIT_4_L(z80 *Z80) {
	z80.bit(4, z80.L)
}

/* BIT 4,(HL) */
func instrCB__BIT_4_iHL(z80 *Z80) {
	bytetemp := z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.bit(4, bytetemp)
}

/* BIT 4,A */
func instrCB__BIT_4_A(z80 *Z80) {
	z80.bit(4, z80.A)
}

/* BIT 5,B */
func instrCB__BIT_5_B(z80 *Z80) {
	z80.bit(5, z80.B)
}

/* BIT 5,C */
func instrCB__BIT_5_C(z80 *Z80) {
	z80.bit(5, z80.C)
}

/* BIT 5,D */
func instrCB__BIT_5_D(z80 *Z80) {
	z80.bit(5, z80.D)
}

/* BIT 5,E */
func instrCB__BIT_5_E(z80 *Z80) {
	z80.bit(5, z80.E)
}

/* BIT 5,H */
func instrCB__BIT_5_H(z80 *Z80) {
	z80.bit(5, z80.H)
}

/* BIT 5,L */
func instrCB__BIT_5_L(z80 *Z80) {
	z80.bit(5, z80.L)
}

/* BIT 5,(HL) */
func instrCB__BIT_5_iHL(z80 *Z80) {
	bytetemp := z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.bit(5, bytetemp)
}

/* BIT 5,A */
func instrCB__BIT_5_A(z80 *Z80) {
	z80.bit(5, z80.A)
}

/* BIT 6,B */
func instrCB__BIT_6_B(z80 *Z80) {
	z80.bit(6, z80.B)
}

/* BIT 6,C */
func instrCB__BIT_6_C(z80 *Z80) {
	z80.bit(6, z80.C)
}

/* BIT 6,D */
func instrCB__BIT_6_D(z80 *Z80) {
	z80.bit(6, z80.D)
}

/* BIT 6,E */
func instrCB__BIT_6_E(z80 *Z80) {
	z80.bit(6, z80.E)
}

/* BIT 6,H */
func instrCB__BIT_6_H(z80 *Z80) {
	z80.bit(6, z80.H)
}

/* BIT 6,L */
func instrCB__BIT_6_L(z80 *Z80) {
	z80.bit(6, z80.L)
}

/* BIT 6,(HL) */
func instrCB__BIT_6_iHL(z80 *Z80) {
	bytetemp := z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.bit(6, bytetemp)
}

/* BIT 6,A */
func instrCB__BIT_6_A(z80 *Z80) {
	z80.bit(6, z80.A)
}

/* BIT 7,B */
func instrCB__BIT_7_B(z80 *Z80) {
	z80.bit(7, z80.B)
}

/* BIT 7,C */
func instrCB__BIT_7_C(z80 *Z80) {
	z80.bit(7, z80.C)
}

/* BIT 7,D */
func instrCB__BIT_7_D(z80 *Z80) {
	z80.bit(7, z80.D)
}

/* BIT 7,E */
func instrCB__BIT_7_E(z80 *Z80) {
	z80.bit(7, z80.E)
}

/* BIT 7,H */
func instrCB__BIT_7_H(z80 *Z80) {
	z80.bit(7, z80.H)
}

/* BIT 7,L */
func instrCB__BIT_7_L(z80 *Z80) {
	z80.bit(7, z80.L)
}

/* BIT 7,(HL) */
func instrCB__BIT_7_iHL(z80 *Z80) {
	bytetemp := z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.bit(7, bytetemp)
}

/* BIT 7,A */
func instrCB__BIT_7_A(z80 *Z80) {
	z80.bit(7, z80.A)
}

/* RES 0,B */
func instrCB__RES_0_B(z80 *Z80) {
	z80.B &= 0xfe
}

/* RES 0,C */
func instrCB__RES_0_C(z80 *Z80) {
	z80.C &= 0xfe
}

/* RES 0,D */
func instrCB__RES_0_D(z80 *Z80) {
	z80.D &= 0xfe
}

/* RES 0,E */
func instrCB__RES_0_E(z80 *Z80) {
	z80.E &= 0xfe
}

/* RES 0,H */
func instrCB__RES_0_H(z80 *Z80) {
	z80.H &= 0xfe
}

/* RES 0,L */
func instrCB__RES_0_L(z80 *Z80) {
	z80.L &= 0xfe
}

/* RES 0,(HL) */
func instrCB__RES_0_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp&0xfe)
}

/* RES 0,A */
func instrCB__RES_0_A(z80 *Z80) {
	z80.A &= 0xfe
}

/* RES 1,B */
func instrCB__RES_1_B(z80 *Z80) {
	z80.B &= 0xfd
}

/* RES 1,C */
func instrCB__RES_1_C(z80 *Z80) {
	z80.C &= 0xfd
}

/* RES 1,D */
func instrCB__RES_1_D(z80 *Z80) {
	z80.D &= 0xfd
}

/* RES 1,E */
func instrCB__RES_1_E(z80 *Z80) {
	z80.E &= 0xfd
}

/* RES 1,H */
func instrCB__RES_1_H(z80 *Z80) {
	z80.H &= 0xfd
}

/* RES 1,L */
func instrCB__RES_1_L(z80 *Z80) {
	z80.L &= 0xfd
}

/* RES 1,(HL) */
func instrCB__RES_1_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp&0xfd)
}

/* RES 1,A */
func instrCB__RES_1_A(z80 *Z80) {
	z80.A &= 0xfd
}

/* RES 2,B */
func instrCB__RES_2_B(z80 *Z80) {
	z80.B &= 0xfb
}

/* RES 2,C */
func instrCB__RES_2_C(z80 *Z80) {
	z80.C &= 0xfb
}

/* RES 2,D */
func instrCB__RES_2_D(z80 *Z80) {
	z80.D &= 0xfb
}

/* RES 2,E */
func instrCB__RES_2_E(z80 *Z80) {
	z80.E &= 0xfb
}

/* RES 2,H */
func instrCB__RES_2_H(z80 *Z80) {
	z80.H &= 0xfb
}

/* RES 2,L */
func instrCB__RES_2_L(z80 *Z80) {
	z80.L &= 0xfb
}

/* RES 2,(HL) */
func instrCB__RES_2_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp&0xfb)
}

/* RES 2,A */
func instrCB__RES_2_A(z80 *Z80) {
	z80.A &= 0xfb
}

/* RES 3,B */
func instrCB__RES_3_B(z80 *Z80) {
	z80.B &= 0xf7
}

/* RES 3,C */
func instrCB__RES_3_C(z80 *Z80) {
	z80.C &= 0xf7
}

/* RES 3,D */
func instrCB__RES_3_D(z80 *Z80) {
	z80.D &= 0xf7
}

/* RES 3,E */
func instrCB__RES_3_E(z80 *Z80) {
	z80.E &= 0xf7
}

/* RES 3,H */
func instrCB__RES_3_H(z80 *Z80) {
	z80.H &= 0xf7
}

/* RES 3,L */
func instrCB__RES_3_L(z80 *Z80) {
	z80.L &= 0xf7
}

/* RES 3,(HL) */
func instrCB__RES_3_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp&0xf7)
}

/* RES 3,A */
func instrCB__RES_3_A(z80 *Z80) {
	z80.A &= 0xf7
}

/* RES 4,B */
func instrCB__RES_4_B(z80 *Z80) {
	z80.B &= 0xef
}

/* RES 4,C */
func instrCB__RES_4_C(z80 *Z80) {
	z80.C &= 0xef
}

/* RES 4,D */
func instrCB__RES_4_D(z80 *Z80) {
	z80.D &= 0xef
}

/* RES 4,E */
func instrCB__RES_4_E(z80 *Z80) {
	z80.E &= 0xef
}

/* RES 4,H */
func instrCB__RES_4_H(z80 *Z80) {
	z80.H &= 0xef
}

/* RES 4,L */
func instrCB__RES_4_L(z80 *Z80) {
	z80.L &= 0xef
}

/* RES 4,(HL) */
func instrCB__RES_4_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp&0xef)
}

/* RES 4,A */
func instrCB__RES_4_A(z80 *Z80) {
	z80.A &= 0xef
}

/* RES 5,B */
func instrCB__RES_5_B(z80 *Z80) {
	z80.B &= 0xdf
}

/* RES 5,C */
func instrCB__RES_5_C(z80 *Z80) {
	z80.C &= 0xdf
}

/* RES 5,D */
func instrCB__RES_5_D(z80 *Z80) {
	z80.D &= 0xdf
}

/* RES 5,E */
func instrCB__RES_5_E(z80 *Z80) {
	z80.E &= 0xdf
}

/* RES 5,H */
func instrCB__RES_5_H(z80 *Z80) {
	z80.H &= 0xdf
}

/* RES 5,L */
func instrCB__RES_5_L(z80 *Z80) {
	z80.L &= 0xdf
}

/* RES 5,(HL) */
func instrCB__RES_5_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp&0xdf)
}

/* RES 5,A */
func instrCB__RES_5_A(z80 *Z80) {
	z80.A &= 0xdf
}

/* RES 6,B */
func instrCB__RES_6_B(z80 *Z80) {
	z80.B &= 0xbf
}

/* RES 6,C */
func instrCB__RES_6_C(z80 *Z80) {
	z80.C &= 0xbf
}

/* RES 6,D */
func instrCB__RES_6_D(z80 *Z80) {
	z80.D &= 0xbf
}

/* RES 6,E */
func instrCB__RES_6_E(z80 *Z80) {
	z80.E &= 0xbf
}

/* RES 6,H */
func instrCB__RES_6_H(z80 *Z80) {
	z80.H &= 0xbf
}

/* RES 6,L */
func instrCB__RES_6_L(z80 *Z80) {
	z80.L &= 0xbf
}

/* RES 6,(HL) */
func instrCB__RES_6_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp&0xbf)
}

/* RES 6,A */
func instrCB__RES_6_A(z80 *Z80) {
	z80.A &= 0xbf
}

/* RES 7,B */
func instrCB__RES_7_B(z80 *Z80) {
	z80.B &= 0x7f
}

/* RES 7,C */
func instrCB__RES_7_C(z80 *Z80) {
	z80.C &= 0x7f
}

/* RES 7,D */
func instrCB__RES_7_D(z80 *Z80) {
	z80.D &= 0x7f
}

/* RES 7,E */
func instrCB__RES_7_E(z80 *Z80) {
	z80.E &= 0x7f
}

/* RES 7,H */
func instrCB__RES_7_H(z80 *Z80) {
	z80.H &= 0x7f
}

/* RES 7,L */
func instrCB__RES_7_L(z80 *Z80) {
	z80.L &= 0x7f
}

/* RES 7,(HL) */
func instrCB__RES_7_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp&0x7f)
}

/* RES 7,A */
func instrCB__RES_7_A(z80 *Z80) {
	z80.A &= 0x7f
}

/* SET 0,B */
func instrCB__SET_0_B(z80 *Z80) {
	z80.B |= 0x01
}

/* SET 0,C */
func instrCB__SET_0_C(z80 *Z80) {
	z80.C |= 0x01
}

/* SET 0,D */
func instrCB__SET_0_D(z80 *Z80) {
	z80.D |= 0x01
}

/* SET 0,E */
func instrCB__SET_0_E(z80 *Z80) {
	z80.E |= 0x01
}

/* SET 0,H */
func instrCB__SET_0_H(z80 *Z80) {
	z80.H |= 0x01
}

/* SET 0,L */
func instrCB__SET_0_L(z80 *Z80) {
	z80.L |= 0x01
}

/* SET 0,(HL) */
func instrCB__SET_0_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp|0x01)
}

/* SET 0,A */
func instrCB__SET_0_A(z80 *Z80) {
	z80.A |= 0x01
}

/* SET 1,B */
func instrCB__SET_1_B(z80 *Z80) {
	z80.B |= 0x02
}

/* SET 1,C */
func instrCB__SET_1_C(z80 *Z80) {
	z80.C |= 0x02
}

/* SET 1,D */
func instrCB__SET_1_D(z80 *Z80) {
	z80.D |= 0x02
}

/* SET 1,E */
func instrCB__SET_1_E(z80 *Z80) {
	z80.E |= 0x02
}

/* SET 1,H */
func instrCB__SET_1_H(z80 *Z80) {
	z80.H |= 0x02
}

/* SET 1,L */
func instrCB__SET_1_L(z80 *Z80) {
	z80.L |= 0x02
}

/* SET 1,(HL) */
func instrCB__SET_1_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp|0x02)
}

/* SET 1,A */
func instrCB__SET_1_A(z80 *Z80) {
	z80.A |= 0x02
}

/* SET 2,B */
func instrCB__SET_2_B(z80 *Z80) {
	z80.B |= 0x04
}

/* SET 2,C */
func instrCB__SET_2_C(z80 *Z80) {
	z80.C |= 0x04
}

/* SET 2,D */
func instrCB__SET_2_D(z80 *Z80) {
	z80.D |= 0x04
}

/* SET 2,E */
func instrCB__SET_2_E(z80 *Z80) {
	z80.E |= 0x04
}

/* SET 2,H */
func instrCB__SET_2_H(z80 *Z80) {
	z80.H |= 0x04
}

/* SET 2,L */
func instrCB__SET_2_L(z80 *Z80) {
	z80.L |= 0x04
}

/* SET 2,(HL) */
func instrCB__SET_2_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp|0x04)
}

/* SET 2,A */
func instrCB__SET_2_A(z80 *Z80) {
	z80.A |= 0x04
}

/* SET 3,B */
func instrCB__SET_3_B(z80 *Z80) {
	z80.B |= 0x08
}

/* SET 3,C */
func instrCB__SET_3_C(z80 *Z80) {
	z80.C |= 0x08
}

/* SET 3,D */
func instrCB__SET_3_D(z80 *Z80) {
	z80.D |= 0x08
}

/* SET 3,E */
func instrCB__SET_3_E(z80 *Z80) {
	z80.E |= 0x08
}

/* SET 3,H */
func instrCB__SET_3_H(z80 *Z80) {
	z80.H |= 0x08
}

/* SET 3,L */
func instrCB__SET_3_L(z80 *Z80) {
	z80.L |= 0x08
}

/* SET 3,(HL) */
func instrCB__SET_3_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp|0x08)
}

/* SET 3,A */
func instrCB__SET_3_A(z80 *Z80) {
	z80.A |= 0x08
}

/* SET 4,B */
func instrCB__SET_4_B(z80 *Z80) {
	z80.B |= 0x10
}

/* SET 4,C */
func instrCB__SET_4_C(z80 *Z80) {
	z80.C |= 0x10
}

/* SET 4,D */
func instrCB__SET_4_D(z80 *Z80) {
	z80.D |= 0x10
}

/* SET 4,E */
func instrCB__SET_4_E(z80 *Z80) {
	z80.E |= 0x10
}

/* SET 4,H */
func instrCB__SET_4_H(z80 *Z80) {
	z80.H |= 0x10
}

/* SET 4,L */
func instrCB__SET_4_L(z80 *Z80) {
	z80.L |= 0x10
}

/* SET 4,(HL) */
func instrCB__SET_4_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp|0x10)
}

/* SET 4,A */
func instrCB__SET_4_A(z80 *Z80) {
	z80.A |= 0x10
}

/* SET 5,B */
func instrCB__SET_5_B(z80 *Z80) {
	z80.B |= 0x20
}

/* SET 5,C */
func instrCB__SET_5_C(z80 *Z80) {
	z80.C |= 0x20
}

/* SET 5,D */
func instrCB__SET_5_D(z80 *Z80) {
	z80.D |= 0x20
}

/* SET 5,E */
func instrCB__SET_5_E(z80 *Z80) {
	z80.E |= 0x20
}

/* SET 5,H */
func instrCB__SET_5_H(z80 *Z80) {
	z80.H |= 0x20
}

/* SET 5,L */
func instrCB__SET_5_L(z80 *Z80) {
	z80.L |= 0x20
}

/* SET 5,(HL) */
func instrCB__SET_5_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp|0x20)
}

/* SET 5,A */
func instrCB__SET_5_A(z80 *Z80) {
	z80.A |= 0x20
}

/* SET 6,B */
func instrCB__SET_6_B(z80 *Z80) {
	z80.B |= 0x40
}

/* SET 6,C */
func instrCB__SET_6_C(z80 *Z80) {
	z80.C |= 0x40
}

/* SET 6,D */
func instrCB__SET_6_D(z80 *Z80) {
	z80.D |= 0x40
}

/* SET 6,E */
func instrCB__SET_6_E(z80 *Z80) {
	z80.E |= 0x40
}

/* SET 6,H */
func instrCB__SET_6_H(z80 *Z80) {
	z80.H |= 0x40
}

/* SET 6,L */
func instrCB__SET_6_L(z80 *Z80) {
	z80.L |= 0x40
}

/* SET 6,(HL) */
func instrCB__SET_6_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp|0x40)
}

/* SET 6,A */
func instrCB__SET_6_A(z80 *Z80) {
	z80.A |= 0x40
}

/* SET 7,B */
func instrCB__SET_7_B(z80 *Z80) {
	z80.B |= 0x80
}

/* SET 7,C */
func instrCB__SET_7_C(z80 *Z80) {
	z80.C |= 0x80
}

/* SET 7,D */
func instrCB__SET_7_D(z80 *Z80) {
	z80.D |= 0x80
}

/* SET 7,E */
func instrCB__SET_7_E(z80 *Z80) {
	z80.E |= 0x80
}

/* SET 7,H */
func instrCB__SET_7_H(z80 *Z80) {
	z80.H |= 0x80
}

/* SET 7,L */
func instrCB__SET_7_L(z80 *Z80) {
	z80.L |= 0x80
}

/* SET 7,(HL) */
func instrCB__SET_7_iHL(z80 *Z80) {
	var bytetemp byte = z80.memory.ReadByte(z80.HL())
	z80.memory.ContendReadNoMreq(z80.HL(), 1)
	z80.memory.WriteByte(z80.HL(), bytetemp|0x80)
}

/* SET 7,A */
func instrCB__SET_7_A(z80 *Z80) {
	z80.A |= 0x80
}
