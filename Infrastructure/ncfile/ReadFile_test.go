package ncfile

import (
	"nc-script-converter/Domain/alterationncscript"
	"reflect"
	"testing"
)

func TestNewNcScriptFile(t *testing.T) {
	tests := []struct {
		name string
		want alterationncscript.FileReader
	}{
		{
			name: "正常系_オブジェクト生成できること",
			want: new(ReadableNcScriptFile),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReadableNcScriptFile(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("NewNcScriptFile() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestNcScriptFile_ReadAll(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		n       *ReadableNcScriptFile
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "正常系_ファイルの内容を読み込めること",
			n:    new(ReadableNcScriptFile),
			args: args{
				"./testdata/test.txt",
			},
			want: []string{
				"test",
				"2nd Line",
			},
			wantErr: false,
		},
		{
			name: "異常系_パスがフランク",
			n:    new(ReadableNcScriptFile),
			args: args{
				"",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系_存在しないパス",
			n:    new(ReadableNcScriptFile),
			args: args{
				"./testdata/no",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.ReadAll(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NcScriptFile.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NcScriptFile.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadableNcScriptFile_FileExist(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		n    *ReadableNcScriptFile
		args args
		want bool
	}{
		{
			name: "正常系_ファイルが存在すること",
			n:    new(ReadableNcScriptFile),
			args: args{file: "testdata/dummy.csv"},
			want: true,
		},
		{
			name: "正常系_存在しないファイル",
			n:    new(ReadableNcScriptFile),
			args: args{file: "testdata/nothing.csv"},
			want: false,
		},
		{
			name: "正常系_ディレクトリ",
			n:    new(ReadableNcScriptFile),
			args: args{file: "testdata/dir"},
			want: false,
		},
		{
			name: "正常系_ディレクトリ(／付き)",
			n:    new(ReadableNcScriptFile),
			args: args{file: "testdata/dir/"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.FileExist(tt.args.file); got != tt.want {
				t.Errorf("ReadableNcScriptFile.FileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
