main.main STEXT nosplit size=150 args=0x0 locals=0x60 funcid=0x0 align=0x0
        0x0000 00000  TEXT    main.main(SB), NOSPLIT|ABIInternal, $96-0
        0x0000 00000  TEXT    main.main(SB), NOSPLIT|ABIInternal, $96-0
        0x0000 00000  SUBQ    $96, SP
        0x0004 00004  MOVQ    BP, 88(SP)
        0x0009 00009  LEAQ    88(SP), BP
        0x000e 00014  FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        0x000e 00014  FUNCDATA        $1, gclocals·h9/3ZXR9C8aF0T5QhErdzQ==(SB)
        0x000e 00014  MOVQ    $0, main.data_stack_int+16(SP)
        0x0017 00023  MOVUPS  X15, main..autotmp_2+24(SP)
        0x001d 00029  MOVUPS  X15, main..autotmp_2+32(SP)
        0x0023 00035  MOVUPS  X15, main..autotmp_2+48(SP)
        0x0029 00041  LEAQ    main..autotmp_2+24(SP), DX
        0x002e 00046  TESTB   AL, (DX)
        0x0030 00048  JMP     50
        0x0032 00050  MOVQ    DX, main.data_heap_slice+64(SP)
        0x0037 00055  MOVQ    $10, main.data_heap_slice+72(SP)
        0x0040 00064  MOVQ    $10, main.data_heap_slice+80(SP)
        0x0049 00073  MOVQ    $10, main.data_global_i(SB)
        0x0054 00084  MOVQ    $111, main.data_stack_int+16(SP)
        0x005d 00093  MOVQ    $111, main.data_global_i(SB)
        0x0068 00104  MOVQ    main.data_heap_slice+72(SP), CX
        0x006d 00109  MOVQ    main.data_heap_slice+64(SP), DX
        0x0072 00114  CMPQ    CX, $1
        0x0076 00118  JHI     122
        0x0078 00120  JMP     139
        0x007a 00122  MOVL    $1, 4(DX)
        0x0081 00129  MOVQ    88(SP), BP
        0x0086 00134  ADDQ    $96, SP
        0x008a 00138  RET
        0x008b 00139  MOVL    $1, AX
        0x0090 00144  PCDATA  $1, $0
        0x0090 00144  CALL    runtime.panicIndex(SB)
        0x0095 00149  XCHGL   AX, AX
        0x0000 48 83 ec 60 48 89 6c 24 58 48 8d 6c 24 58 48 c7  H..`H.l$XH.l$XH.
        0x0010 44 24 10 00 00 00 00 44 0f 11 7c 24 18 44 0f 11  D$.....D..|$.D..
        0x0020 7c 24 20 44 0f 11 7c 24 30 48 8d 54 24 18 84 02  |$ D..|$0H.T$...
        0x0030 eb 00 48 89 54 24 40 48 c7 44 24 48 0a 00 00 00  ..H.T$@H.D$H....
        0x0040 48 c7 44 24 50 0a 00 00 00 48 c7 05 00 00 00 00  H.D$P....H......
        0x0050 0a 00 00 00 48 c7 44 24 10 6f 00 00 00 48 c7 05  ....H.D$.o...H..
        0x0060 00 00 00 00 6f 00 00 00 48 8b 4c 24 48 48 8b 54  ....o...H.L$HH.T
        0x0070 24 40 48 83 f9 01 77 02 eb 11 c7 42 04 01 00 00  $@H...w....B....
        0x0080 00 48 8b 6c 24 58 48 83 c4 60 c3 b8 01 00 00 00  .H.l$XH..`......
        0x0090 e8 00 00 00 00 90                                ......
        rel 76+4 t=14 main.data_global_i+-4
        rel 96+4 t=14 main.data_global_i+-4
        rel 145+4 t=7 runtime.panicIndex+0
main.init STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0 align=0x0
        0x0000 00000   TEXT    main.init(SB), NOSPLIT|ABIInternal, $0-0
        0x0000 00000   FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        0x0000 00000   FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        0x0000 00000   RET
        0x0000 c3                                               .
go:cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
        0x0000 2d 4e 20 2d 6c 20 72 65 67 61 62 69              -N -l regabi
go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
main..inittask SNOPTRDATA size=24
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
main.data_global_i SNOPTRBSS size=8
go:info.main.data_global_i SDWARFVAR dupok size=35
        0x0000 0a 6d 61 69 6e 2e 64 61 74 61 5f 67 6c 6f 62 61  .main.data_globa
        0x0010 6c 5f 69 00 09 03 00 00 00 00 00 00 00 00 00 00  l_i.............
        0x0020 00 00 01                                         ...
        rel 22+8 t=1 main.data_global_i+0
        rel 30+4 t=31 go:info.int+0
type:.eqfunc40 SRODATA dupok size=16
        0x0000 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
        rel 0+8 t=1 runtime.memequal_varlen+0
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.0100000000000000 SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
type:.namedata.*[10]int32- SRODATA dupok size=12
        0x0000 00 0a 2a 5b 31 30 5d 69 6e 74 33 32              ..*[10]int32
type:*[10]int32 SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 ad 79 2e 67 08 08 08 36 00 00 00 00 00 00 00 00  .y.g...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.0100000000000000+0
        rel 40+4 t=5 type:.namedata.*[10]int32-+0
        rel 48+8 t=1 type:[10]int32+0
runtime.gcbits. SRODATA dupok size=0
type:[10]int32 SRODATA dupok size=72
        0x0000 28 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  (...............
        0x0010 21 cf f0 07 0a 04 04 11 00 00 00 00 00 00 00 00  !...............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 0a 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 type:.eqfunc40+0
        rel 32+8 t=1 runtime.gcbits.+0
        rel 40+4 t=5 type:.namedata.*[10]int32-+0
        rel 44+4 t=-32763 type:*[10]int32+0
        rel 48+8 t=1 type:int32+0
        rel 56+8 t=1 type:[]int32+0
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·h9/3ZXR9C8aF0T5QhErdzQ== SRODATA dupok size=9
        0x0000 01 00 00 00 03 00 00 00 00                       .........