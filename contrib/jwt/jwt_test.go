package jwt

import "testing"

func TestCreateToken(t *testing.T) {
	type args struct {
		uid uint
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				uid: 20,
				key: "abc",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.args.uid, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		tokenString string
		key         string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDg1OTI0MDIsInVpZCI6MjB9.l-3eG-1vRRgjWbWKtag2NjOj3LaH_f5-Y-OPTl99Is0",
				key:         "abc",
			},
			want:    20,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.args.tokenString, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
