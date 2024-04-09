package export

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/wagoodman/dive/dive/image/docker"
)

func Test_Export(t *testing.T) {
	result := docker.TestAnalysisFromArchive(t, "../../.data/test-docker-image.tar")

	export := NewExport(result)
	payload, err := export.Marshal()
	if err != nil {
		t.Errorf("Test_Export: unable to export analysis: %v", err)
	}

	expectedResult := `{
  "layer": [
    {
      "index": 0,
      "id": "28cfe03618aa2e914e81fdd90345245c15f4478e35252c06ca52d238fd3cc694",
      "digestId": "sha256:23bc2b70b2014dec0ac22f27bb93e9babd08cdd6f1115d0c955b9ff22b382f5a",
      "sizeBytes": 1154361,
      "command": "#(nop) ADD file:ce026b62356eec3ad1214f92be2c9dc063fe205bd5e600be3492c4dfb17148bd in / ",
      "fileList": [
        {
          "Path": "bin/[",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 1075464,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/[[",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/acpid",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/add-shell",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/addgroup",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/adduser",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/adjtimex",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ar",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/arch",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/arp",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/arping",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ash",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/awk",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/base64",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/basename",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/beep",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/blkdiscard",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/blkid",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/blockdev",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/bootchartd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/brctl",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/bunzip2",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/busybox",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/bzcat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/bzip2",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cal",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chattr",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chgrp",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chmod",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chown",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chpasswd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chpst",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chroot",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chrt",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/chvt",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cksum",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/clear",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cmp",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/comm",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/conspy",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cp",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cpio",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/crond",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/crontab",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cryptpw",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cttyhack",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/cut",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/date",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/deallocvt",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/delgroup",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/deluser",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/depmod",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/devmem",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/df",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dhcprelay",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/diff",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dirname",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dmesg",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dnsd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dnsdomainname",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dos2unix",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dpkg",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dpkg-deb",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/du",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dumpkmap",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/dumpleases",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/echo",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ed",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/egrep",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/eject",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/env",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/envdir",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/envuidgid",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ether-wake",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/expand",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/expr",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/factor",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fakeidentd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fallocate",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/false",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fatattr",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fbset",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fbsplash",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fdflush",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fdformat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fdisk",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fgconsole",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fgrep",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/find",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/findfs",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/flock",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fold",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/free",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/freeramdisk",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fsck",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fsck.minix",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fsfreeze",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fstrim",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fsync",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ftpd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ftpget",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ftpput",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/fuser",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/getconf",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 77880,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/getopt",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/getty",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/grep",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/groups",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/gunzip",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/gzip",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/halt",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/hd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/hdparm",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/head",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/hexdump",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/hexedit",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/hostid",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/hostname",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/httpd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/hush",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/hwclock",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/i2cdetect",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/i2cdump",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/i2cget",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/i2cset",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/id",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ifconfig",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ifdown",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ifenslave",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ifplugd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ifup",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/inetd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/init",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/insmod",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/install",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ionice",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/iostat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ip",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ipaddr",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ipcalc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ipcrm",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ipcs",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/iplink",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ipneigh",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/iproute",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/iprule",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/iptunnel",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/kbd_mode",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/kill",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/killall",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/killall5",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/klogd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/last",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/less",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/link",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/linux32",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/linux64",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/linuxrc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ln",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/loadfont",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/loadkmap",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/logger",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/login",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/logname",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/logread",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/losetup",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lpd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lpq",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lpr",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ls",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lsattr",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lsmod",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lsof",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lspci",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lsscsi",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lsusb",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lzcat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lzma",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/lzop",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/makedevs",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/makemime",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/man",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/md5sum",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mdev",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mesg",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/microcom",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mkdir",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mkdosfs",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mke2fs",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mkfifo",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mkfs.ext2",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mkfs.minix",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mkfs.vfat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mknod",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mkpasswd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mkswap",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mktemp",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/modinfo",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/modprobe",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/more",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mount",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mountpoint",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mpstat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mt",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/mv",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nameif",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nanddump",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nandwrite",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nbd-client",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/netstat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nice",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nl",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nmeter",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nohup",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nproc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nsenter",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nslookup",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ntpd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/nuke",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/od",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/openvt",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/partprobe",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/passwd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/paste",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/patch",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pgrep",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pidof",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ping",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ping6",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pipe_progress",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pivot_root",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pkill",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pmap",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/popmaildir",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/poweroff",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/powertop",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/printenv",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/printf",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ps",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pscan",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pstree",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pwd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/pwdx",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/raidautorun",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rdate",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rdev",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/readahead",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/readlink",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/readprofile",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/realpath",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/reboot",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/reformime",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/remove-shell",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/renice",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/reset",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/resize",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/resume",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rev",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rm",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rmdir",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rmmod",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/route",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rpm",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rpm2cpio",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rtcwake",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/run-init",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/run-parts",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/runlevel",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/runsv",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/runsvdir",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/rx",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/script",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/scriptreplay",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sed",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sendmail",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/seq",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setarch",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setconsole",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setfattr",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setfont",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setkeycodes",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setlogcons",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setpriv",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setserial",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setsid",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/setuidgid",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sh",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sha1sum",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sha256sum",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sha3sum",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sha512sum",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/showkey",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/shred",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/shuf",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/slattach",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sleep",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/smemcap",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/softlimit",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sort",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/split",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ssl_client",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/start-stop-daemon",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/stat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/strings",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/stty",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/su",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sulogin",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sum",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sv",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/svc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/svlogd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/svok",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/swapoff",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/swapon",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/switch_root",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sync",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/sysctl",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/syslogd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tac",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tail",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tar",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/taskset",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tcpsvd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tee",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/telnet",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/telnetd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/test",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tftp",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tftpd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/time",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/timeout",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/top",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/touch",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tr",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/traceroute",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/traceroute6",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/true",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/truncate",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tty",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ttysize",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/tunctl",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ubiattach",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ubidetach",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ubimkvol",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ubirename",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ubirmvol",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ubirsvol",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/ubiupdatevol",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/udhcpc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/udhcpd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/udpsvd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/uevent",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/umount",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/uname",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/unexpand",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/uniq",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/unix2dos",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/unlink",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/unlzma",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/unshare",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/unxz",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/unzip",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/uptime",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/users",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/usleep",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/uudecode",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/uuencode",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/vconfig",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/vi",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/vlock",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/volname",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/w",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/wall",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/watch",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/watchdog",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/wc",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/wget",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/which",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/who",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/whoami",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/whois",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/xargs",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/xxd",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/xz",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/xzcat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/yes",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/zcat",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin/zcip",
          "TypeFlag": 49,
          "Linkname": "bin/[",
          "Size": 0,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "bin",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "dev",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "etc/group",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 307,
          "Mode": 436,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "etc/localtime",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 127,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "etc/network/if-down.d",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "etc/network/if-post-down.d",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "etc/network/if-pre-up.d",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "etc/network/if-up.d",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "etc/network",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "etc/passwd",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 340,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "etc/shadow",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 243,
          "Mode": 384,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "etc",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "home",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 65534,
          "Gid": 65534,
          "IsDir": true
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "tmp",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2148532735,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "usr/sbin",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 1,
          "Gid": 1,
          "IsDir": true
        },
        {
          "Path": "usr",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "var/spool/mail",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 8,
          "Gid": 8,
          "IsDir": true
        },
        {
          "Path": "var/spool",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "var/www",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "var",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 1,
      "id": "1871059774abe6914075e4a919b778fa1561f577d620ae52438a9635e6241936",
      "digestId": "sha256:a65b7d7ac139a0e4337bc3c73ce511f937d6140ef61a0108f7d4b8aab8d67274",
      "sizeBytes": 6405,
      "command": "#(nop) ADD file:139c3708fb6261126453e34483abd8bf7b26ed16d952fd976994d68e72d93be2 in /somefile.txt ",
      "fileList": [
        {
          "Path": "somefile.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 436,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        }
      ]
    },
    {
      "index": 2,
      "id": "49fe2a475548bfa4d493fc796fce41f30704e3d4cbff3e45dd3e06f463236d1d",
      "digestId": "sha256:93e208d471756ffbac88cf9c25feb442007f221d3bd73231e27b747a0a68927c",
      "sizeBytes": 0,
      "command": "mkdir -p /root/example/really/nested",
      "fileList": [
        {
          "Path": "root/example/really/nested",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root/example/really",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root/example",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 3,
      "id": "80cd2ca1ffc89962b9349c80280c2bc551acbd11e09b16badb0669f8e2369020",
      "digestId": "sha256:4abad3abe3cb99ad7a492a9d9f6b3d66287c1646843c74128bbbec4f7be5aa9e",
      "sizeBytes": 6405,
      "command": "cp /somefile.txt /root/example/somefile1.txt",
      "fileList": [
        {
          "Path": "root/example/somefile1.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root/example",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 4,
      "id": "c99e2f8d3f6282668f0d30dc1db5e67a51d7a1dcd7ff6ddfa0f90760836778ec",
      "digestId": "sha256:14c9a6ffcb6a0f32d1035f97373b19608e2d307961d8be156321c3f1c1504cbf",
      "sizeBytes": 6405,
      "command": "chmod 444 /root/example/somefile1.txt",
      "fileList": [
        {
          "Path": "root/example/somefile1.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 292,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root/example",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 5,
      "id": "5eca617bdc3bc06134fe957a30da4c57adb7c340a6d749c8edc4c15861c928d7",
      "digestId": "sha256:778fb5770ef466f314e79cc9dc418eba76bfc0a64491ce7b167b76aa52c736c4",
      "sizeBytes": 6405,
      "command": "cp /somefile.txt /root/example/somefile2.txt",
      "fileList": [
        {
          "Path": "root/example/somefile2.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root/example",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 6,
      "id": "f07c3eb887572395408f8e11a07af945e4da5f02b3188bb06b93fad713ca0b99",
      "digestId": "sha256:f275b8a31a71deb521cc048e6021e2ff6fa52bedb25c9b7bbe129a0195ddca5f",
      "sizeBytes": 6405,
      "command": "cp /somefile.txt /root/example/somefile3.txt",
      "fileList": [
        {
          "Path": "root/example/somefile3.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root/example",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 7,
      "id": "461885fc22589158dee3c5b9f01cc41c87805439f58b4399d733b51aa305cbf9",
      "digestId": "sha256:dd1effc5eb19894c3e9b57411c98dd1cf30fa1de4253c7fae53c9cea67267d83",
      "sizeBytes": 6405,
      "command": "mv /root/example/somefile3.txt /root/saved.txt",
      "fileList": [
        {
          "Path": "root/example/.wh.somefile3.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 0,
          "Mode": 0,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root/example",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root/saved.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 8,
      "id": "a10327f68ffed4afcba78919052809a8f774978a6b87fc117d39c53c4842f72c",
      "digestId": "sha256:8d1869a0a066cdd12e48d648222866e77b5e2814f773bb3bd8774ab4052f0f1d",
      "sizeBytes": 6405,
      "command": "cp /root/saved.txt /root/.saved.txt",
      "fileList": [
        {
          "Path": "root/.saved.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 9,
      "id": "f2fc54e25cb7966dc9732ec671a77a1c5c104e732bd15ad44a2dc1ac42368f84",
      "digestId": "sha256:bc2e36423fa31a97223fd421f22c35466220fa160769abf697b8eb58c896b468",
      "sizeBytes": 0,
      "command": "rm -rf /root/example/",
      "fileList": [
        {
          "Path": "root/.wh.example",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 0,
          "Mode": 0,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 10,
      "id": "aad36d0b05e71c7e6d4dfe0ca9ed6be89e2e0d8995dafe83438299a314e91071",
      "digestId": "sha256:7f648d45ee7b6de2292162fba498b66cbaaf181da9004fcceef824c72dbae445",
      "sizeBytes": 2187,
      "command": "#(nop) ADD dir:7ec14b81316baa1a31c38c97686a8f030c98cba2035c968412749e33e0c4427e in /root/.data/ ",
      "fileList": [
        {
          "Path": "root/.data/tag.sh",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 917,
          "Mode": 509,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root/.data/test.sh",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 1270,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root/.data",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 11,
      "id": "3d4ad907517a021d86a4102d2764ad2161e4818bbd144e41d019bfc955434181",
      "digestId": "sha256:a4b8f95f266d5c063c9a9473c45f2f85ddc183e37941b5e6b6b9d3c00e8e0457",
      "sizeBytes": 6405,
      "command": "cp /root/saved.txt /tmp/saved.again1.txt",
      "fileList": [
        {
          "Path": "tmp/saved.again1.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "tmp",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2148532735,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 12,
      "id": "81b1b002d4b4c1325a9cad9990b5277e7f29f79e0f24582344c0891178f95905",
      "digestId": "sha256:22a44d45780a541e593a8862d80f3e14cb80b6bf76aa42ce68dc207a35bf3a4a",
      "sizeBytes": 6405,
      "command": "cp /root/saved.txt /root/.data/saved.again2.txt",
      "fileList": [
        {
          "Path": "root/.data/saved.again2.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root/.data",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    },
    {
      "index": 13,
      "id": "cfb35bb5c127d848739be5ca726057e6e2c77b2849f588e7aebb642c0d3d4b7b",
      "digestId": "sha256:ba689cac6a98c92d121fa5c9716a1bab526b8bb1fd6d43625c575b79e97300c5",
      "sizeBytes": 6405,
      "command": "chmod +x /root/saved.txt",
      "fileList": [
        {
          "Path": "root/saved.txt",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 6405,
          "Mode": 493,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "root",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484096,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        }
      ]
    }
  ],
  "image": {
    "sizeBytes": 1220598,
    "inefficientBytes": 32025,
    "efficiencyScore": 0.9844212134184309,
    "fileReference": [
      {
        "count": 2,
        "sizeBytes": 12810,
        "file": "/root/saved.txt"
      },
      {
        "count": 2,
        "sizeBytes": 12810,
        "file": "/root/example/somefile1.txt"
      },
      {
        "count": 2,
        "sizeBytes": 6405,
        "file": "/root/example/somefile3.txt"
      }
    ]
  }
}`
	actualResult := string(payload)
	if expectedResult != actualResult {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expectedResult, actualResult, false)

		t.Errorf("Test_Export: unexpected export result:\n%v", dmp.DiffPrettyText(diffs))
	}
}
