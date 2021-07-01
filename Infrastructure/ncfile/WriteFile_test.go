package ncfile

import (
	"bufio"
	"nc-script-converter/Domain/alterationncscript"
	"os"
	"reflect"
	"testing"
)

func TestNewWritableNcScriptFile(t *testing.T) {
	tests := []struct {
		name string
		want alterationncscript.FileWriter
	}{
		{
			name: "正常系_オブジェクト生成できること",
			want: new(WritableNcScriptFile),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWritableNcScriptFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWritableNcScriptFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWritableNcScriptFile_WriteAll(t *testing.T) {
	type args struct {
		path     string
		contents []string
	}
	tests := []struct {
		name    string
		n       *WritableNcScriptFile
		args    args
		wantErr bool
	}{
		{
			name: "正常系_ファイルが保存されること",
			n:    new(WritableNcScriptFile),
			args: args{
				path: "./testdata/output_newScript",
				contents: []string{
					"",
					"(O4701)",
					"T16",
					"M6 Q0",
					"G91G0G28Z0",
					"G54",
					"G90G0X0Y0",
					"G0B0C0",
					"G0W0",
					"G43Z100.H16",
					"M01",
					"S4500M3",
					"M8",
					"G05.1Q1",
					"G05.1Q0",
					"G80",
					"M5",
					"M9",
					"G91G0G28Z0",
					"(M99)",
					"",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.WriteAll(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("WritableNcScriptFile.WriteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
			fp, _ := os.Open(tt.args.path)
			defer fp.Close()
			s := bufio.NewScanner(fp)
			var i int
			for s.Scan() {
				actual := s.Text()
				if actual != tt.args.contents[i] {
					t.Errorf("WritableNcScriptFile.WriteAll() contets error line: %d, actual: %s, expected: %s", i, actual, tt.args.contents[i])
				}
				i++
			}
		})
	}
}
