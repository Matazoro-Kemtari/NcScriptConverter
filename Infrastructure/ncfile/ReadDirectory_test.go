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
			if got := NewNcScriptDir(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNcScriptDir() = %v, want %v", got, tt.want)
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
				"testdata\\dummy.csv",
				"testdata\\test.txt",
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
				"testdata\\dummy.csv",
				"testdata\\test.txt",
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
