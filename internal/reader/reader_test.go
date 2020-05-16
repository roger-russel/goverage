package reader

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/docker/docker/pkg/testutil/assert"
)

func Test_splitContent(t *testing.T) {

	type args struct {
		c          CoverStruct
		data       string
		file       string
		duplicated map[string]bool
	}
	tests := []struct {
		name string
		args args
		want CoverStruct
	}{
		{
			name: "simple test",
			args: args{
				c:          make(CoverStruct, 0),
				data:       "14.34,26.2 7 1",
				file:       "main.go",
				duplicated: make(map[string]bool),
			},
			want: func() (c CoverStruct) {
				count := 1
				c = make(CoverStruct, 0)

				c["main.go"] = &LineCover{
					NumberOfStatements: 0,
					Report:             make(map[int]map[int]int, 0),
				}

				c["main.go"].NumberOfStatements = 7

				c["main.go"].Report[14] = make(map[int]int, 0)
				c["main.go"].Report[14][34] = count

				c["main.go"].Report[26] = make(map[int]int, 0)
				c["main.go"].Report[26][2] += count * -1

				return c
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitContent(tt.args.c, tt.args.file, tt.args.data, tt.args.duplicated)
			assert.DeepEqual(t, tt.args.c, tt.want)
		})
	}
}

func TestReadFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want CoverStruct
	}{
		{
			name: "small-report test",
			args: args{
				filename: "../../tests/fixtures/small-report.txt",
			},
			want: func() (c CoverStruct) {
				c = make(CoverStruct, 0)

				c["goverage-test-crud/internal/pkg/api/routers/routers.go"] = &LineCover{
					NumberOfStatements: 9,
					Report:             make(map[int]map[int]int, 0),
				}

				c["goverage-test-crud/cmd/crud/main.go"] = &LineCover{
					NumberOfStatements: 10,
					Report:             make(map[int]map[int]int, 0),
				}

				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[14] = make(map[int]int, 0)
				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[14][34]++

				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[26] = make(map[int]int, 0)
				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[26][2]--

				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[28] = make(map[int]int, 0)
				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[28][132] += 4

				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[29] = make(map[int]int, 0)
				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[29][54] -= 4
				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[29][54] += 0

				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[31] = make(map[int]int, 0)
				c["goverage-test-crud/internal/pkg/api/routers/routers.go"].Report[31][3] += 0

				c["goverage-test-crud/cmd/crud/main.go"].Report[23] = make(map[int]int, 0)
				c["goverage-test-crud/cmd/crud/main.go"].Report[23][13]++

				c["goverage-test-crud/cmd/crud/main.go"].Report[29] = make(map[int]int, 0)
				c["goverage-test-crud/cmd/crud/main.go"].Report[29][16]--
				c["goverage-test-crud/cmd/crud/main.go"].Report[29][16] += 0

				c["goverage-test-crud/cmd/crud/main.go"].Report[31] = make(map[int]int, 0)
				c["goverage-test-crud/cmd/crud/main.go"].Report[31][3] -= 0

				c["goverage-test-crud/cmd/crud/main.go"].Report[33] = make(map[int]int, 0)
				c["goverage-test-crud/cmd/crud/main.go"].Report[33][2]++

				c["goverage-test-crud/cmd/crud/main.go"].Report[41] = make(map[int]int, 0)
				c["goverage-test-crud/cmd/crud/main.go"].Report[41][16]--
				return c

			}(),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ReadFile(tt.args.filename)

			if !reflect.DeepEqual(got, tt.want) {
				jGot, _ := json.Marshal(got)
				jWant, _ := json.Marshal(tt.want)
				t.Errorf("ReadFile() = %v, want %v", string(jGot), string(jWant))
			}
		})
	}
}
