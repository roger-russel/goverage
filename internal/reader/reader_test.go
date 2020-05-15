package reader

import (
	"testing"

	"github.com/docker/docker/pkg/testutil/assert"
)

func Test_splitContent(t *testing.T) {

	type args struct {
		c    CoverStruct
		data string
		file string
	}
	tests := []struct {
		name string
		args args
		want CoverStruct
	}{
		{
			name: "simple test",
			args: args{
				c:    make(CoverStruct, 0),
				data: "14.34,26.2 7 1",
				file: "main.go",
			},
			want: func() (c CoverStruct) {
				count := 1
				c = make(CoverStruct, 0)
				c["main.go"] = make(map[int]map[int]int, 0)
				c["main.go"][14] = make(map[int]int, 0)
				c["main.go"][14][34] = count
				c["main.go"][26] = make(map[int]int, 0)
				c["main.go"][26][2] += count * -1
				return c
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitContent(tt.args.c, tt.args.file, tt.args.data)
			assert.DeepEqual(t, tt.args.c, tt.want)
		})
	}
}
