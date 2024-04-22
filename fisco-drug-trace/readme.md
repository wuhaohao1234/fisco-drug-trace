```shell
 ./solc-0.8.11 --bin --abi -o ./fisco-drug-trace/contract/ ./fisco-drug-trace/contract/DrugTrace.sol 
./abigen --bin ./fisco-drug-trace/contract/DrugTrace.bin --abi ./fisco-drug-trace/contract/DrugTrace.abi --pkg sdk --type DrugTrace --out ./fisco-drug-trace/DrugTrace.go
```