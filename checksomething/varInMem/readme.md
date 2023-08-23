

go build -o vars.o vars.go
go tool objdump -S vars.o > vars-o.s 