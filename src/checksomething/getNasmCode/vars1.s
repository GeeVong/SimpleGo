TEXT main.main(SB) gofile../Users/hq/GeeVong/golangReview/src/checksomething/getNasmCode/vars.go
func main() {
  0xb8a			4883ec60		SUBQ $0x60, SP		
  0xb8e			48896c2458		MOVQ BP, 0x58(SP)	
  0xb93			488d6c2458		LEAQ 0x58(SP), BP	
	var data_stack_int int
  0xb98			48c744241000000000	MOVQ $0x0, 0x10(SP)	
	data_heap_slice := make([]int32, 10, 10)
  0xba1			440f117c2418		MOVUPS X15, 0x18(SP)	
  0xba7			440f117c2420		MOVUPS X15, 0x20(SP)	
  0xbad			440f117c2430		MOVUPS X15, 0x30(SP)	
  0xbb3			488d542418		LEAQ 0x18(SP), DX	
  0xbb8			8402			TESTB AL, 0(DX)		
  0xbba			eb00			JMP 0xbbc		
  0xbbc			4889542440		MOVQ DX, 0x40(SP)	
  0xbc1			48c74424480a000000	MOVQ $0xa, 0x48(SP)	
  0xbca			48c74424500a000000	MOVQ $0xa, 0x50(SP)	
	data_global_i = 10
  0xbd3			48c705000000000a000000	MOVQ $0xa, 0(IP)	[3:7]R_PCREL:main.data_global_i+-4	
	data_stack_int = 111
  0xbde			48c74424106f000000	MOVQ $0x6f, 0x10(SP)	
	data_global_i = +data_stack_int
  0xbe7			48c705000000006f000000	MOVQ $0x6f, 0(IP)	[3:7]R_PCREL:main.data_global_i+-4	
	data_heap_slice[1] = 1
  0xbf2			488b4c2448		MOVQ 0x48(SP), CX	
  0xbf7			488b542440		MOVQ 0x40(SP), DX	
  0xbfc			4883f901		CMPQ $0x1, CX		
  0xc00			7702			JA 0xc04		
  0xc02			eb11			JMP 0xc15		
  0xc04			c7420401000000		MOVL $0x1, 0x4(DX)	
	// 		0x0049 00073  MOVQ    $10, main.data_global_i(SB)
  0xc0b			488b6c2458		MOVQ 0x58(SP), BP	
  0xc10			4883c460		ADDQ $0x60, SP		
  0xc14			c3			RET			
	data_heap_slice[1] = 1
  0xc15			b801000000		MOVL $0x1, AX		
  0xc1a			e800000000		CALL 0xc1f		[1:5]R_CALL:runtime.panicIndex	
  0xc1f			90			NOPL			

TEXT main.init(SB) gofile../Users/hq/GeeVong/golangReview/src/checksomething/getNasmCode/vars.go
var data_global_i int
  0xc20			c3			RET			
