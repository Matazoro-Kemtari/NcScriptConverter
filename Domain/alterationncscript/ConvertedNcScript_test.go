package alterationncscript

import (
	"reflect"
	"testing"
)

func TestNewConvertedNcScript(t *testing.T) {
	tests := []struct {
		name string
		want *ConvertedNcScript
	}{
		{
			name: "正常系_オブジェクト生成できること",
			want: NewConvertedNcScript(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConvertedNcScript(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConvertedNcScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertedNcScript_Convert(t *testing.T) {
	type args struct {
		source        []string
		canOpenReview bool
	}
	tests := []struct {
		name    string
		c       *ConvertedNcScript
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "正常系_カッタースクリプトが変換されること",
			c:    NewConvertedNcScript(),
			args: args{
				[]string{
					"%",
					"O4701",
					"(T16)",
					"(S4500)",
					"X0.Y0.",
					"G90X0.Y0.",
					"G54",
					"X0.Y0.",
					"M99",
					"%",
				},
				false,
			},
			want: []string{
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
				"X0.Y0.",
				"G90X0.Y0.",
				"G49",
				"G54",
				"X0.Y0.",
				"G05.1Q0",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_センタードリルスクリプトが変換されること",
			c:    NewConvertedNcScript(),
			args: args{
				[]string{
					"%",
					"O4702",
					"(T12)",
					"(S2000)",
					"(G82)",
					"X0.Y0.",
					"G90",
					"X0.Y0.",
					"M99",
					"%",
				},
				false,
			},
			want: []string{
				"",
				"(O4702)",
				"T12",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H12",
				"M01",
				"S2000M3",
				"M8",
				"G98G82R2.0Z-1.0Q2.0P500F180L0",
				"X0.Y0.",
				"G90",
				"X0.Y0.",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_下穴ドリルスクリプトが変換されること",
			c:    NewConvertedNcScript(),
			args: args{
				[]string{
					"%",
					"O4702",
					"(T13)",
					"(S1800)",
					"(G83)",
					"X0.Y0.",
					"G90",
					"X0.Y0.",
					"M99",
					"%",
				},
				false,
			},
			want: []string{
				"",
				"(O4702)",
				"T13",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H13",
				"M01",
				"S1800M3",
				"M8",
				"G98G83R2.0 Z-45.Q2.0F180L0",
				"X0.Y0.",
				"G90",
				"X0.Y0.",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_リーマスクリプトが変換されること",
			c:    NewConvertedNcScript(),
			args: args{
				[]string{
					"%",
					"O4702",
					"(T15)",
					"(S1500)",
					"(G85)",
					"X0.Y0.",
					"G90",
					"X0.Y0.",
					"M99",
					"%",
				},
				false,
			},
			want: []string{
				"M00",
				"",
				"(O4702)",
				"T15",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H15",
				"M01",
				"S1500M3",
				"M8",
				"G98G85R2.0 Z-35.F150L0",
				"X0.Y0.",
				"G90",
				"X0.Y0.",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_カッタースクリプトが変換されること",
			c:    NewConvertedNcScript(),
			args: args{
				[]string{
					"%",
					"O4701",
					"(T16)",
					"(S4500)",
					"X0.Y0.",
					"G90X0.Y0.",
					"X0.Y0.",
					"M30",
					"%",
				},
				false,
			},
			want: []string{
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
				"X0.Y0.",
				"G90X0.Y0.",
				"X0.Y0.",
				"G91G0G28Z0",
				"G91G0G28B0",
				"G91G0G28C0",
				"(M30)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_オープンレビュー変換されること",
			c:    NewConvertedNcScript(),
			args: args{
				[]string{
					"%",
					"O4701",
					"(20T-1147 216A HON)",
					"(D50XR3 BXD ZENSYU ARA)",
					"(STARTPOINT X0. Y0. Z100.)",
					"(LASTPOINT Z100.)",
					"(TIME 3 MIN )",
					"(T16)",
					"(S4500)",
					"G91Z0.",
					"G0X-77.34Y121.073",
					"Z0.",
					"Z-81.",
					"G1Z-19.F2600.",
					"G2X45.5Y45.5I45.5J0.",
					"M30",
					"G3X5.Y5.I0.J5.",
					"G54",
					"G1Y30.",
					"G0G40Z139.",
					"Z0.",
					"G90X0.Y0.",
					"M99",
					"%",
				},
				true,
			},
			want: []string{
				"",
				"(O4701)",
				"(20T-1147 216A HON)",
				"(D50XR3 BXD ZENSYU ARA)",
				"(STARTPOINT X0. Y0. Z100.)",
				"(LASTPOINT Z100.)",
				"(TIME 3 MIN )",
				"T16",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H16",
				"M01",
				"(M99)",
				"",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Convert(tt.args.source, tt.args.canOpenReview)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertedNcScript.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertedNcScript.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
