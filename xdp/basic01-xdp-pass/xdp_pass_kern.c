//go:build ignore
// +build ignore

#include <vmlinux.h>
#include <bpf_helpers.h>

char LICENSE[] SEC("license") = "Dual BSD/GPL";


SEC("xdp/example_pass")
int xdp_pass(struct xdp_md *ctx)
{
	return XDP_PASS;
}



