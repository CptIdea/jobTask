package app

import "testing"

func Test_emailValidator_validEmail(t *testing.T) {
	type fields struct {
		token  string
		secret string
	}

	testFields := fields{
		token:  "d6ba3449152602cbd63d215ed1ee612e6b3134eb",
		secret: "fbce4c61182373033f9184812ef91cb7180a05d6",
	}

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "Temp mail", fields: testFields, args: args{email: "qa05y@vmani.com"}, want: false, wantErr: false},
		{name: "No Token error", fields: fields{}, args: args{email: "pliyznik@yandex.ru"}, want: false, wantErr: true},
		{name: "Valid email", fields: testFields, args: args{email: "pliyznik@yandex.ru"}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := &emailValidator{
				Token:  tt.fields.token,
				Secret: tt.fields.secret,
			}
			got, err := ev.validEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("validEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
