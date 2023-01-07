package logger

import "testing"

func init() {
	InitLog("test.log")
}

func TestDebug(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_debug_1",
			args: args{msg: "debug test abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.msg)
		})
	}
}

func TestDebugF(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DebugF(tt.args.format, tt.args.args...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.msg)
		})
	}
}

func TestErrorF(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorF(tt.args.format, tt.args.args...)
		})
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatal(tt.args.err)
		})
	}
}

func TestFatalF(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FatalF(tt.args.format, tt.args.args...)
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_info_1",
			args: args{msg: "info test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg)
		})
	}
}

func TestInfoF(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_infof_1",
			args: args{
				format: "Hello World! %s",
				args:   []interface{}{"zhang san"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoF(tt.args.format, tt.args.args...)
		})
	}
}

func TestInitLog(t *testing.T) {
	type args struct {
		logPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLog(tt.args.logPath)
		})
	}
}

func TestTrace(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Trace(tt.args.msg)
		})
	}
}

func TestTraceF(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TraceF(tt.args.format, tt.args.args...)
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.msg)
		})
	}
}

func TestWarnF(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarnF(tt.args.format, tt.args.args...)
		})
	}
}
