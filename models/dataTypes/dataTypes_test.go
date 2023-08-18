package datatypes

import (
	"reflect"
	"testing"
)

func Test_buildDataSlice(t *testing.T) {
	type args struct {
		schemajsn []byte
		datajsn   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []Data
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildDataSlice(tt.args.schemajsn, tt.args.datajsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildDataSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildDataSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
