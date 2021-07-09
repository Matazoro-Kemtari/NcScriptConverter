package ncfile

import (
	"nc-script-converter/Domain/alterationncscript"
	"reflect"
	"testing"
)

func TestNewNcScriptDir(t *testing.T) {
	tests := []struct {
		name string
		want alterationncscript.DirViewer
	}{
		{
			name: "正常系_オブジェクト生成できること",
			want: new(NcScriptDir),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNcScriptDir(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("NewNcScriptDir() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestNcScriptDir_FetchDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		n       *NcScriptDir
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "正常系_ファイルの一覧が取得できること",
			n:    new(NcScriptDir),
			args: args{
				path: "./testdata",
			},
			want: []string{
				"dummy.csv",
				"output_newScript",
				"test.txt",
			},
			wantErr: false,
		},
		{
			name: "正常系_ファイルの一覧が取得できること(/付き)",
			n:    new(NcScriptDir),
			args: args{
				path: "./testdata/",
			},
			want: []string{
				"dummy.csv",
				"output_newScript",
				"test.txt",
			},
			wantErr: false,
		},
		{
			name: "異常系_パスにブランク",
			n:    new(NcScriptDir),
			args: args{
				path: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.FetchDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NcScriptDir.FetchDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NcScriptDir.FetchDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNcScriptDir_DirExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		n    *NcScriptDir
		args args
		want bool
	}{
		{
			name: "正常系_存在するファイル",
			n:    new(NcScriptDir),
			args: args{path: "testdata/dummy.csv"},
			want: false,
		},
		{
			name: "正常系_存在しないファイル",
			n:    new(NcScriptDir),
			args: args{path: "testdata/nothing.csv"},
			want: false,
		},
		{
			name: "正常系_存在するディレクトリ",
			n:    new(NcScriptDir),
			args: args{path: "testdata/dir"},
			want: true,
		},
		{
			name: "正常系_存在するディレクトリ(／付き)",
			n:    new(NcScriptDir),
			args: args{path: "testdata/dir/"},
			want: true,
		},
		{
			name: "正常系_存在しないディレクトリ",
			n:    new(NcScriptDir),
			args: args{path: "testdata/nothingdir"},
			want: false,
		},
		{
			name: "正常系_存在しないディレクトリ(／付き)",
			n:    new(NcScriptDir),
			args: args{path: "testdata/nothingdir/"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.DirExist(tt.args.path); got != tt.want {
				t.Errorf("NcScriptDir.DirExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
