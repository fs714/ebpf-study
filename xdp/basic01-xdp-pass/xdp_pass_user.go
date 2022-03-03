//go:build linux
// +build linux

package main

import (
	"github.com/cilium/ebpf/rlimit"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target amd64 -cc clang-10 exampleXdp ./xdp_pass_kern.c -- -nostdinc -Wall -Werror -I../../include -I../../libbpf/src

func main() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}

	objs := exampleXdpObjects{}
	if err := loadExampleXdpObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	// Now you can attach XDP programs to network interfaces with Cilium eBPF library.
	// This feature has added to the library since version 0.8.0 and it requires at
	// least Linux 5.9. (It uses bpf_link to attach, not netlink)
	//iface, err := net.InterfaceByName("enp0s3")
	//if err != nil {
	//	log.Fatalf("interface by name: %v", err)
	//}
	//
	//l, err := link.AttachXDP(link.XDPOptions{
	//	Program:   objs.XdpPass,
	//	Interface: iface.Index,
	//})
	//if err != nil {
	//	log.Fatalf("attach XDP: %v", err)
	//}
	//defer l.Close()

	link, err := netlink.LinkByName("enp0s3")
	if err != nil {
		log.Fatalf("link by name: %v", err)
	}

	err = netlink.LinkSetXdpFdWithFlags(link, objs.XdpPass.FD(), nl.XDP_FLAGS_SKB_MODE)
	if err != nil {
		log.Fatalf("link set XDP fd: %v", err)
	}

	<-signalCh
	_ = netlink.LinkSetXdpFdWithFlags(link, -1, nl.XDP_FLAGS_SKB_MODE)
}
