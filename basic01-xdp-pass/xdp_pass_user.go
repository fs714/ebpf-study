//go:build linux
// +build linux

package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target amd64 -cc clang-10 bpf_fs ./xdp_pass_kern.c -- -nostdinc -Wall -Werror -I../headers -I/usr/src/linux-headers-5.4.0-92-generic/include -I/usr/src/linux-headers-5.4.0-92-generic/arch/x86/include/generated/uapi -I/usr/src/linux-headers-5.4.0-92/arch/x86/include/uapi

func main() {

}
