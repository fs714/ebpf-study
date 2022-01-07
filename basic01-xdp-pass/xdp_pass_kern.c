// +build ignore

#include <linux/bpf.h>
#include "bpf_helpers.h"

char _license[] SEC("license") = "GPL";

SEC("xdp")
int xdp_prog_simple(struct xdp_md *ctx)
{
	return XDP_PASS;
}



