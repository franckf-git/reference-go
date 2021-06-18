# run
go run main.go fibo.go binet.go recursive.go

# cpu profile
go test -cpuprofile cpu.prof -bench recursive_test.go
# take 5-6 sec

# show profile
go tool pprof cpu.prof
# in cmd prompt : show process
top5 -cum