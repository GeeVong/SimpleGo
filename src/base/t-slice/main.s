runtime.printnl@plt-0x10:
 push   QWORD PTR [rip+0x7fca]        # 409ff0 <_GLOBAL_OFFSET_TABLE_+0x8>
 jmp    QWORD PTR [rip+0x7fcc]        # 409ff8 <_GLOBAL_OFFSET_TABLE_+0x10>
 nop    DWORD PTR [rax+0x0]
main:
 cmp    rsp,QWORD PTR fs:0x70
 jb     402324 <main+0x24>
 mov    rax,QWORD PTR [rip+0x7cae]        # 409fc0 <runtime_isarchive@Base>
 mov    BYTE PTR [rax],0x0
 mov    rax,QWORD PTR [rip+0x7c6c]        # 409f88 <runtime_isstarted@Base>
 cmp    BYTE PTR [rax],0x0
 je     402335 <main+0x35>
 xor    eax,eax
 ret
 mov    r10d,0x18
 xor    r11d,r11d
 call   4033b7 <__morestack>
 ret
 jmp    40230b <main+0xb>
 sub    rsp,0x18
 mov    BYTE PTR [rax],0x1
 mov    rax,QWORD PTR [rip+0x7c3d]        # 409f80 <runtime_iscgo@Base>
 cmp    BYTE PTR [rax],0x0
 jne    4023a9 <main+0xa9>
 mov    rdx,0x40a240
 mov    rax,0x40a188
 mov    QWORD PTR [rsp+0x8],rsi
 mov    DWORD PTR [rsp+0x4],edi
 mov    QWORD PTR [rax],rdx
 call   402050 <runtime.ginit@plt>
 call   4020c0 <runtime_cpuinit@plt>
 call   402250 <runtime.check@plt>
 mov    rsi,QWORD PTR [rsp+0x8]
 mov    edi,DWORD PTR [rsp+0x4]
 call   402290 <runtime.args@plt>
 call   4020e0 <runtime.osinit@plt>
 call   402080 <runtime.schedinit@plt>
 mov    rdi,QWORD PTR [rip+0x7c10]        # 409fa0 <runtime.main@Base>
 xor    esi,esi
 call   4021e0 <__go_go@plt>
 call   402280 <runtime_m@plt>
 mov    rdi,rax
 call   402130 <runtime_mstart@plt>
 call   402180 <abort@plt>
 mov    QWORD PTR [rsp+0x8],rsi
 mov    DWORD PTR [rsp+0x4],edi
 call   4021f0 <runtime.setIsCgo@plt>
 mov    rsi,QWORD PTR [rsp+0x8]
 mov    edi,DWORD PTR [rsp+0x4]
 jmp    402348 <main+0x48>
 cs nop WORD PTR [rax+rax*1+0x0]
 nop    DWORD PTR [rax+0x0]
main.sliceParam:
 cmp    rsp,QWORD PTR fs:0x70
 jae    4024d3 <main.sliceParam+0x1d>
 mov    r10d,0x98
 mov    r11d,0x18
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 push   r13
 push   r12
 push   rbx
 sub    rsp,0x78
 mov    QWORD PTR [rbp-0x88],rdi
 pxor   xmm0,xmm0
 movaps XMMWORD PTR [rbp-0x80],xmm0
 movq   QWORD PTR [rbp-0x70],xmm0
 mov    rax,QWORD PTR [rbp+0x18]
 test   rax,rax
 jg     40250a <main.sliceParam+0x54>
 mov    rsi,rax
 mov    edi,0x0
 call   4022a0 <runtime.goPanicIndex@plt>
 mov    rax,QWORD PTR [rbp+0x10]
 mov    DWORD PTR [rax],0x457
 mov    rax,QWORD PTR [rbp+0x10]
 mov    rdx,QWORD PTR [rbp+0x18]
 mov    QWORD PTR [rbp-0x60],rax
 mov    QWORD PTR [rbp-0x58],rdx
 mov    rax,QWORD PTR [rbp+0x20]
 mov    QWORD PTR [rbp-0x50],rax
 mov    rbx,QWORD PTR [rbp-0x58]
 lea    rax,[rbx+0x1]
 mov    rcx,QWORD PTR [rbp-0x50]
 mov    rdx,rax
 mov    rsi,rcx
 cmp    rsi,rdx
 jae    402578 <main.sliceParam+0xc2>
 mov    rdx,QWORD PTR [rbp-0x60]
 lea    rdi,[rbp-0x40]
 mov    r9,rax
 mov    r8,rcx
 mov    rcx,rbx
 mov    esi,0x4050e0
 call   402200 <runtime.growslice@plt>
 mov    rax,QWORD PTR [rbp-0x40]
 mov    rdx,QWORD PTR [rbp-0x38]
 mov    QWORD PTR [rbp+0x10],rax
 mov    QWORD PTR [rbp+0x18],rdx
 mov    rax,QWORD PTR [rbp-0x30]
 mov    QWORD PTR [rbp+0x20],rax
 jmp    40258c <main.sliceParam+0xd6>
 mov    rcx,QWORD PTR [rbp-0x60]
 mov    rdx,QWORD PTR [rbp-0x50]
 mov    QWORD PTR [rbp+0x10],rcx
 mov    QWORD PTR [rbp+0x18],rax
 mov    QWORD PTR [rbp+0x20],rdx
 mov    rax,QWORD PTR [rbp+0x10]
 mov    rdx,QWORD PTR [rbp+0x18]
 mov    QWORD PTR [rbp-0x60],rax
 mov    QWORD PTR [rbp-0x58],rdx
 mov    rax,QWORD PTR [rbp+0x20]
 mov    QWORD PTR [rbp-0x50],rax
 mov    rax,QWORD PTR [rbp-0x60]
 mov    rdx,rbx
 shl    rdx,0x2
 add    rax,rdx
 mov    DWORD PTR [rax],0x64
 mov    r12d,0x405011
 mov    r13d,0xb
 call   4022b0 <runtime.printlock@plt>
 mov    rax,r12
 mov    rdx,r13
 mov    rcx,r12
 mov    rbx,r13
 mov    rdx,rax
 mov    rax,rbx
 mov    rdi,rdx
 mov    rsi,rax
 call   4020f0 <runtime.printstring@plt>
 call   402260 <runtime.printsp@plt>
 sub    rsp,0x8
 sub    rsp,0x18
 mov    rcx,rsp
 mov    rax,QWORD PTR [rbp+0x10]
 mov    rdx,QWORD PTR [rbp+0x18]
 mov    QWORD PTR [rcx],rax
 mov    QWORD PTR [rcx+0x8],rdx
 mov    rax,QWORD PTR [rbp+0x20]
 mov    QWORD PTR [rcx+0x10],rax
 call   402110 <runtime.printslice@plt>
 add    rsp,0x20
 call   402030 <runtime.printnl@plt>
 call   4022d0 <runtime.printunlock@plt>
 mov    rax,QWORD PTR [rbp+0x10]
 mov    rdx,QWORD PTR [rbp+0x18]
 mov    QWORD PTR [rbp-0x80],rax
 mov    QWORD PTR [rbp-0x78],rdx
 mov    rax,QWORD PTR [rbp+0x20]
 mov    QWORD PTR [rbp-0x70],rax
 mov    rcx,QWORD PTR [rbp-0x88]
 mov    rax,QWORD PTR [rbp-0x80]
 mov    rdx,QWORD PTR [rbp-0x78]
 mov    QWORD PTR [rcx],rax
 mov    QWORD PTR [rcx+0x8],rdx
 mov    rax,QWORD PTR [rbp-0x70]
 mov    QWORD PTR [rcx+0x10],rax
 mov    rax,QWORD PTR [rbp-0x88]
 lea    rsp,[rbp-0x18]
 pop    rbx
 pop    r12
 pop    r13
 pop    rbp
 ret
main.main:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402685 <main.main+0x1d>
 mov    r10d,0xc8
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 push   rbx
 sub    rsp,0xb8
 pxor   xmm0,xmm0
 movaps XMMWORD PTR [rbp-0xa0],xmm0
 movaps XMMWORD PTR [rbp-0x90],xmm0
 movq   QWORD PTR [rbp-0x80],xmm0
 lea    rax,[rbp-0xa0]
 mov    QWORD PTR [rbp-0xc0],rax
 mov    QWORD PTR [rbp-0xb8],0x0
 mov    QWORD PTR [rbp-0xb0],0xa
 mov    rax,QWORD PTR [rbp-0xc0]
 mov    rdx,QWORD PTR [rbp-0xb8]
 mov    QWORD PTR [rbp-0x70],rax
 mov    QWORD PTR [rbp-0x68],rdx
 mov    rax,QWORD PTR [rbp-0xb0]
 mov    QWORD PTR [rbp-0x60],rax
 mov    rbx,QWORD PTR [rbp-0x68]
 lea    rax,[rbx+0x1]
 mov    rcx,QWORD PTR [rbp-0x60]
 mov    rdx,rax
 mov    rsi,rcx
 cmp    rsi,rdx
 jae    402742 <main.main+0xda>
 mov    rdx,QWORD PTR [rbp-0x70]
 lea    rdi,[rbp-0x30]
 mov    r9,rax
 mov    r8,rcx
 mov    rcx,rbx
 mov    esi,0x4050e0
 call   402200 <runtime.growslice@plt>
 mov    rax,QWORD PTR [rbp-0x30]
 mov    rdx,QWORD PTR [rbp-0x28]
 mov    QWORD PTR [rbp-0xc0],rax
 mov    QWORD PTR [rbp-0xb8],rdx
 mov    rax,QWORD PTR [rbp-0x20]
 mov    QWORD PTR [rbp-0xb0],rax
 jmp    40275f <main.main+0xf7>
 mov    rcx,QWORD PTR [rbp-0x70]
 mov    rdx,QWORD PTR [rbp-0x60]
 mov    QWORD PTR [rbp-0xc0],rcx
 mov    QWORD PTR [rbp-0xb8],rax
 mov    QWORD PTR [rbp-0xb0],rdx
 mov    rax,QWORD PTR [rbp-0xc0]
 mov    rdx,QWORD PTR [rbp-0xb8]
 mov    QWORD PTR [rbp-0x70],rax
 mov    QWORD PTR [rbp-0x68],rdx
 mov    rax,QWORD PTR [rbp-0xb0]
 mov    QWORD PTR [rbp-0x60],rax
 mov    rax,QWORD PTR [rbp-0x70]
 mov    rdx,rbx
 shl    rdx,0x2
 add    rax,rdx
 mov    DWORD PTR [rax],0x1
 lea    rsi,[rbp-0x50]
 sub    rsp,0x8
 sub    rsp,0x18
 mov    rcx,rsp
 mov    rax,QWORD PTR [rbp-0xc0]
 mov    rdx,QWORD PTR [rbp-0xb8]
 mov    QWORD PTR [rcx],rax
 mov    QWORD PTR [rcx+0x8],rdx
 mov    rax,QWORD PTR [rbp-0xb0]
 mov    QWORD PTR [rcx+0x10],rax
 mov    rdi,rsi
 call   4024b6 <main.sliceParam>
 add    rsp,0x20
 mov    rax,QWORD PTR [rbp-0x50]
 mov    rdx,QWORD PTR [rbp-0x48]
 mov    QWORD PTR [rbp-0xc0],rax
 mov    QWORD PTR [rbp-0xb8],rdx
 mov    rax,QWORD PTR [rbp-0x40]
 mov    QWORD PTR [rbp-0xb0],rax
 mov    rbx,QWORD PTR [rbp-0x8]
 leave
 ret
struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0empty..stub:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402813 <struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0empty..stub+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    BYTE PTR [rbp-0x1],0x0
 cmp    QWORD PTR [rbp-0x18],0x0
 jne    40282f <struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0empty..stub+0x39>
 call   4020d0 <runtime.panicmem@plt>
 mov    rax,QWORD PTR [rbp-0x18]
 mov    rdi,rax
 call   4022c0 <runtime.gList.empty@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0pop..stub:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402861 <struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0pop..stub+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x8],0x0
 cmp    QWORD PTR [rbp-0x18],0x0
 jne    402881 <struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0pop..stub+0x3d>
 call   4020d0 <runtime.panicmem@plt>
 mov    rax,QWORD PTR [rbp-0x18]
 mov    rdi,rax
 call   402230 <runtime.gList.pop@plt>
 mov    QWORD PTR [rbp-0x8],rax
 mov    rax,QWORD PTR [rbp-0x8]
 leave
 ret
struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0push..stub:
 cmp    rsp,QWORD PTR fs:0x70
 jae    4028b4 <struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0push..stub+0x1d>
 mov    r10d,0x18
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x10
 mov    QWORD PTR [rbp-0x8],rdi
 mov    QWORD PTR [rbp-0x10],rsi
 cmp    QWORD PTR [rbp-0x8],0x0
 jne    4028d0 <struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0push..stub+0x39>
 call   4020d0 <runtime.panicmem@plt>
 mov    rax,QWORD PTR [rbp-0x8]
 mov    rdx,rax
 mov    rax,QWORD PTR [rbp-0x10]
 mov    rsi,rax
 mov    rdi,rdx
 call   402060 <runtime.gList.push@plt>
 leave
 ret
struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0pushAll..stub:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402905 <struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0pushAll..stub+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x8],rdi
 mov    rcx,rdx
 mov    rax,rsi
 mov    rdx,rdi
 mov    rdx,rcx
 mov    QWORD PTR [rbp-0x20],rax
 mov    QWORD PTR [rbp-0x18],rdx
 cmp    QWORD PTR [rbp-0x8],0x0
 jne    402931 <struct_4runtime_0gList_cruntime_0n_bint32_5.runtime_0pushAll..stub+0x49>
 call   4020d0 <runtime.panicmem@plt>
 mov    rax,QWORD PTR [rbp-0x8]
 mov    rcx,rax
 mov    rdx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    rsi,rdx
 mov    rdx,rax
 mov    rdi,rcx
 call   402040 <runtime.gList.pushAll@plt>
 leave
 ret
main.struct_4runtime_0gList_cruntime_0n_bint32_5..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    40296d <main.struct_4runtime_0gList_cruntime_0n_bint32_5..eq+0x1d>
 mov    r10d,0x8
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x18]
 mov    rdx,QWORD PTR [rbp-0x20]
 mov    eax,0x1
 test   al,al
 je     40299d <main.struct_4runtime_0gList_cruntime_0n_bint32_5..eq+0x4d>
 mov    rax,QWORD PTR [rcx]
 mov    rsi,rax
 mov    rax,QWORD PTR [rdx]
 cmp    rsi,rax
 sete   al
 xor    eax,0x1
 test   al,al
 je     4029ae <main.struct_4runtime_0gList_cruntime_0n_bint32_5..eq+0x5e>
 mov    BYTE PTR [rbp-0x1],0x0
 movzx  eax,BYTE PTR [rbp-0x1]
 jmp    4029ca <main.struct_4runtime_0gList_cruntime_0n_bint32_5..eq+0x7a>
 mov    ecx,DWORD PTR [rcx+0x8]
 mov    eax,DWORD PTR [rdx+0x8]
 cmp    ecx,eax
 je     4029c2 <main.struct_4runtime_0gList_cruntime_0n_bint32_5..eq+0x72>
 mov    BYTE PTR [rbp-0x1],0x0
 movzx  eax,BYTE PTR [rbp-0x1]
 jmp    4029ca <main.struct_4runtime_0gList_cruntime_0n_bint32_5..eq+0x7a>
 mov    BYTE PTR [rbp-0x1],0x1
 movzx  eax,BYTE PTR [rbp-0x1]
 pop    rbp
 ret
main._632_7uintptr..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    4029e9 <main._632_7uintptr..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x100
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._6256_7uint64..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402a3b <main._6256_7uint64..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x800
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402a8d <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x1d>
 mov    r10d,0xa8
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 push   rbx
 sub    rsp,0x98
 mov    QWORD PTR [rbp-0x98],rdi
 mov    QWORD PTR [rbp-0xa0],rsi
 mov    BYTE PTR [rbp-0x11],0x0
 mov    rsi,QWORD PTR [rbp-0x98]
 mov    rdi,QWORD PTR [rbp-0xa0]
 mov    r8d,0x3d
 mov    eax,0x0
 jmp    402bda <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x16a>
 nop
 mov    edx,0x1
 test   dl,dl
 je     402b9e <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x12e>
 test   rax,rax
 js     402b34 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0xc4>
 cmp    rax,0x3c
 jg     402b34 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0xc4>
 mov    rdx,rax
 add    rdx,rdx
 add    rdx,rax
 shl    rdx,0x3
 add    rdx,rsi
 mov    rcx,QWORD PTR [rdx]
 mov    rbx,QWORD PTR [rdx+0x8]
 mov    QWORD PTR [rbp-0x90],rcx
 mov    QWORD PTR [rbp-0x88],rbx
 mov    rdx,QWORD PTR [rdx+0x10]
 mov    QWORD PTR [rbp-0x80],rdx
 mov    rcx,QWORD PTR [rbp-0x90]
 mov    rbx,QWORD PTR [rbp-0x88]
 mov    QWORD PTR [rbp-0x50],rcx
 mov    QWORD PTR [rbp-0x48],rbx
 mov    rdx,QWORD PTR [rbp-0x80]
 mov    QWORD PTR [rbp-0x40],rdx
 test   rax,rax
 jns    402b41 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0xd1>
 jmp    402b47 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0xd7>
 mov    esi,0x3d
 mov    rdi,rax
 call   4022a0 <runtime.goPanicIndex@plt>
 cmp    rax,0x3c
 jle    402b54 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0xe4>
 mov    esi,0x3d
 mov    rdi,rax
 call   4022a0 <runtime.goPanicIndex@plt>
 mov    rdx,rax
 add    rdx,rdx
 add    rdx,rax
 shl    rdx,0x3
 add    rdx,rdi
 mov    rcx,QWORD PTR [rdx]
 mov    rbx,QWORD PTR [rdx+0x8]
 mov    QWORD PTR [rbp-0x70],rcx
 mov    QWORD PTR [rbp-0x68],rbx
 mov    rdx,QWORD PTR [rdx+0x10]
 mov    QWORD PTR [rbp-0x60],rdx
 mov    rcx,QWORD PTR [rbp-0x70]
 mov    rbx,QWORD PTR [rbp-0x68]
 mov    QWORD PTR [rbp-0x30],rcx
 mov    QWORD PTR [rbp-0x28],rbx
 mov    rdx,QWORD PTR [rbp-0x60]
 mov    QWORD PTR [rbp-0x20],rdx
 mov    ecx,DWORD PTR [rbp-0x50]
 mov    edx,DWORD PTR [rbp-0x30]
 cmp    ecx,edx
 sete   dl
 test   dl,dl
 je     402bb3 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x143>
 mov    rcx,QWORD PTR [rbp-0x88]
 mov    rdx,QWORD PTR [rbp-0x68]
 cmp    rcx,rdx
 sete   dl
 test   dl,dl
 je     402bc5 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x155>
 mov    rcx,QWORD PTR [rbp-0x80]
 mov    rdx,QWORD PTR [rbp-0x60]
 cmp    rcx,rdx
 sete   dl
 xor    edx,0x1
 test   dl,dl
 je     402bd6 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x166>
 mov    BYTE PTR [rbp-0x11],0x0
 movzx  eax,BYTE PTR [rbp-0x11]
 jmp    402beb <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x17b>
 add    rax,0x1
 cmp    rax,r8
 jl     402ac9 <main._661_7struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x59>
 mov    BYTE PTR [rbp-0x11],0x1
 movzx  eax,BYTE PTR [rbp-0x11]
 mov    rbx,QWORD PTR [rbp-0x8]
 leave
 ret
main.struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402c0e <main.struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x1d>
 mov    r10d,0x8
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rdx,QWORD PTR [rbp-0x18]
 mov    rax,QWORD PTR [rbp-0x20]
 mov    esi,DWORD PTR [rdx]
 mov    ecx,DWORD PTR [rax]
 cmp    esi,ecx
 je     402c38 <main.struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x47>
 mov    BYTE PTR [rbp-0x1],0x0
 movzx  eax,BYTE PTR [rbp-0x1]
 jmp    402c6e <main.struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x7d>
 mov    rsi,QWORD PTR [rdx+0x8]
 mov    rcx,QWORD PTR [rax+0x8]
 cmp    rsi,rcx
 je     402c4f <main.struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x5e>
 mov    BYTE PTR [rbp-0x1],0x0
 movzx  eax,BYTE PTR [rbp-0x1]
 jmp    402c6e <main.struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x7d>
 mov    rdx,QWORD PTR [rdx+0x10]
 mov    rax,QWORD PTR [rax+0x10]
 cmp    rdx,rax
 je     402c66 <main.struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x75>
 mov    BYTE PTR [rbp-0x1],0x0
 movzx  eax,BYTE PTR [rbp-0x1]
 jmp    402c6e <main.struct_4Size_buint32_cMallocs_buint64_cFrees_buint64_5..eq+0x7d>
 mov    BYTE PTR [rbp-0x1],0x1
 movzx  eax,BYTE PTR [rbp-0x1]
 pop    rbp
 ret
main._6122_7uintptr..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402c8d <main._6122_7uintptr..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x3d0
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._68_7uint64..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402cdf <main._68_7uint64..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x40
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._6128_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402d31 <main._6128_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x80
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._64096_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402d83 <main._64096_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x1000
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._668_7uint16..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402dd5 <main._668_7uint16..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x88
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._633_7float64..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402e27 <main._633_7float64..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rdx,QWORD PTR [rbp-0x18]
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    esi,0x21
 mov    eax,0x0
 jmp    402ea8 <main._633_7float64..eq+0x9e>
 nop
 test   rax,rax
 js     402e62 <main._633_7float64..eq+0x58>
 cmp    rax,0x20
 jg     402e62 <main._633_7float64..eq+0x58>
 test   rax,rax
 jns    402e6f <main._633_7float64..eq+0x65>
 jmp    402e8d <main._633_7float64..eq+0x83>
 mov    esi,0x21
 mov    rdi,rax
 call   4022a0 <runtime.goPanicIndex@plt>
 cmp    rax,0x20
 jg     402e8d <main._633_7float64..eq+0x83>
 movsd  xmm0,QWORD PTR [rdx+rax*8]
 movsd  xmm1,QWORD PTR [rcx+rax*8]
 ucomisd xmm0,xmm1
 jp     402e9a <main._633_7float64..eq+0x90>
 ucomisd xmm0,xmm1
 je     402ea4 <main._633_7float64..eq+0x9a>
 jmp    402e9a <main._633_7float64..eq+0x90>
 mov    esi,0x21
 mov    rdi,rax
 call   4022a0 <runtime.goPanicIndex@plt>
 mov    BYTE PTR [rbp-0x1],0x0
 movzx  eax,BYTE PTR [rbp-0x1]
 jmp    402eb5 <main._633_7float64..eq+0xab>
 add    rax,0x1
 cmp    rax,rsi
 jl     402e4f <main._633_7float64..eq+0x45>
 mov    BYTE PTR [rbp-0x1],0x1
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._665_7uint32..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402ed4 <main._665_7uint32..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x104
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._64_7uintptr..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402f26 <main._64_7uintptr..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x20
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._65_7uint..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402f78 <main._65_7uint..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x28
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._6512_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    402fca <main._6512_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x200
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._6249_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    40301c <main._6249_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0xf9
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._6129_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    40306e <main._6129_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x81
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._632_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    4030c0 <main._632_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x20
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._627_7string..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    403112 <main._627_7string..eq+0x1d>
 mov    r10d,0x48
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 push   r14
 push   r13
 push   r12
 push   rbx
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x38],rdi
 mov    QWORD PTR [rbp-0x40],rsi
 mov    BYTE PTR [rbp-0x21],0x0
 mov    r12,QWORD PTR [rbp-0x38]
 mov    r13,QWORD PTR [rbp-0x40]
 mov    r14d,0x1b
 mov    ebx,0x0
 jmp    4031de <main._627_7string..eq+0xe9>
 nop
 test   rbx,rbx
 js     403158 <main._627_7string..eq+0x63>
 cmp    rbx,0x1a
 jg     403158 <main._627_7string..eq+0x63>
 test   rbx,rbx
 jns    403165 <main._627_7string..eq+0x70>
 jmp    40319a <main._627_7string..eq+0xa5>
 mov    esi,0x1b
 mov    rdi,rbx
 call   4022a0 <runtime.goPanicIndex@plt>
 cmp    rbx,0x1a
 jg     40319a <main._627_7string..eq+0xa5>
 mov    rax,rbx
 shl    rax,0x4
 add    rax,r12
 mov    rsi,QWORD PTR [rax]
 mov    rdi,QWORD PTR [rax+0x8]
 mov    rax,rbx
 shl    rax,0x4
 add    rax,r13
 mov    rdx,QWORD PTR [rax+0x8]
 mov    rax,QWORD PTR [rax]
 mov    r8,rdi
 mov    rcx,rdx
 cmp    r8,rcx
 jne    4031d0 <main._627_7string..eq+0xdb>
 jmp    4031a7 <main._627_7string..eq+0xb2>
 mov    esi,0x1b
 mov    rdi,rbx
 call   4022a0 <runtime.goPanicIndex@plt>
 mov    r8,rsi
 mov    rcx,rax
 cmp    r8,rcx
 je     4031da <main._627_7string..eq+0xe5>
 mov    rcx,rdx
 mov    r8,rcx
 mov    rcx,rax
 mov    rax,rsi
 mov    rdx,r8
 mov    rsi,rcx
 mov    rdi,rax
 call   402240 <memcmp@plt>
 test   eax,eax
 je     4031da <main._627_7string..eq+0xe5>
 mov    BYTE PTR [rbp-0x21],0x0
 movzx  eax,BYTE PTR [rbp-0x21]
 jmp    4031ef <main._627_7string..eq+0xfa>
 add    rbx,0x1
 cmp    rbx,r14
 jl     403145 <main._627_7string..eq+0x50>
 mov    BYTE PTR [rbp-0x21],0x1
 movzx  eax,BYTE PTR [rbp-0x21]
 add    rsp,0x20
 pop    rbx
 pop    r12
 pop    r13
 pop    r14
 pop    rbp
 ret
main._61024_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    403219 <main._61024_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x400
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._62_7int32..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    40326b <main._62_7int32..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x8
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._664_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    4032bd <main._664_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x40
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret
main._6256_7uint8..eq:
 cmp    rsp,QWORD PTR fs:0x70
 jae    40330f <main._6256_7uint8..eq+0x1d>
 mov    r10d,0x28
 mov    r11d,0x0
 call   4033b7 <__morestack>
 ret
 push   rbp
 mov    rbp,rsp
 sub    rsp,0x20
 mov    QWORD PTR [rbp-0x18],rdi
 mov    QWORD PTR [rbp-0x20],rsi
 mov    BYTE PTR [rbp-0x1],0x0
 mov    rcx,QWORD PTR [rbp-0x20]
 mov    rax,QWORD PTR [rbp-0x18]
 mov    edx,0x100
 mov    rsi,rcx
 mov    rdi,rax
 call   402140 <runtime.memequal@plt>
 mov    BYTE PTR [rbp-0x1],al
 movzx  eax,BYTE PTR [rbp-0x1]
 leave
 ret