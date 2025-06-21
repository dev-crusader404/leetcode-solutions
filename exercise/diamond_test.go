package exercise

import "testing"

func TestGen(t *testing.T) {
	type args struct {
		char byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "C",
			args: args{
				char: 'C',
			},
			want: `  A  
 B B 
C   C
 B B 
  A  `,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Gen(tt.args.char)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Gen() = %v, want %v", got, tt.want)
			}
		})
	}
}
