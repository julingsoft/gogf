package guidx

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    uint64
		wantErr bool
	}{
		{
			name:    "case1",
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
