package ncfile

import (
	"bufio"
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
			want: new(NcScriptFile),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNcScriptFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNcScriptFile() = %v, want %v", got, tt.want)
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
		n       *NcScriptFile
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "正常系_ファイルの内容を読み込めること",
			n:    new(NcScriptFile),
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
			n:    new(NcScriptFile),
			args: args{
				"",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系_存在しないパス",
			n:    new(NcScriptFile),
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
			defer got.Close()
			s := bufio.NewScanner(got)
			var lines []string
			for s.Scan() {
				lines = append(lines, s.Text())
			}
			if !reflect.DeepEqual(lines, tt.want) {
				t.Errorf("NcScriptFile.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
