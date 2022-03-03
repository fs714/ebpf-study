# ebpf-study
Study ebpf and xdp with golang based on `github.com/cilium/ebpf`

## 1. ENV Setup on Ubuntu
- Install Build Tools
```
sudo apt install build-essential clang llvm 
```
- Clone Project
```
git clone https://github.com/fs714/ebpf-study.git
cd ebpf-study
git submodule update --init --recursive
```

## 2. Run First Example
```
cd cd trace/kprobe/
rm examplekprobe_bpfel_x86.go examplekprobe_bpfel_x86.o
go generate
go build
sudo ./kprobe
```
