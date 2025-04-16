# go-processor-template

Template for wasmVision processor written in Go

## How to build

```shell
tinygo build -o go-processor-template.wasm -target=wasip1 -buildmode=c-shared -scheduler=none --no-debug .
```
