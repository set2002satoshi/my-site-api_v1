package auth

import (
	"testing"
)

func TestComparisonPassAndHash(t *testing.T) {
	type args struct {
		CurrentPassword string
		EntryPassword   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				CurrentPassword: "$2a$14$pxE/ywxntCpndHQuSXFXi.ttiCmfQuK70uH4gNl78PSuoAYeNykTu",
				EntryPassword:   "Password",
			},
			want: false,
		},
		{
			name: "ng",
			args: args{
				CurrentPassword: "$2a$14$pxE/ywxntCpndHQuSXFXi.ttiCmfQuK70uH4gNl78PSuoAYeNykTu",
				EntryPassword:   "aaa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparisonPassAndHash(tt.args.CurrentPassword, tt.args.EntryPassword); got != tt.want {
				t.Errorf("ComparisonPassAndHash = %v, want %v", got, tt.want)
			}
		})
	}
}
