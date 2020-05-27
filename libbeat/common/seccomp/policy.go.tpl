// Code generated by seccomp-profiler - DO NOT EDIT.

// {{ printf "+build linux,"}}{{.GOARCH}}

package {{.Package}}

import (
	"github.com/elastic/go-seccomp-bpf"

	beat "github.com/elastic/beats/v7/libbeat/common/seccomp"
)

func init() {
	beat.MustRegisterPolicy(&seccomp.Policy{
		DefaultAction: seccomp.ActionErrno,
		Syscalls: []seccomp.SyscallGroup{
			{
				Action: seccomp.ActionAllow,
				Names: []string{
{{- range $syscall := .SyscallNames}}
					"{{ $syscall }}",
{{- end}}
				},
			},
		},
	})
}
