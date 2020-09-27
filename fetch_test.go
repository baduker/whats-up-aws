package whats_up_aws

import "testing"

func Test_buildURL(t *testing.T) {
	type args struct {
		page int
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test full URL",
			args: args{
				page: 0,
				size: 25,
			},
			want: "https://aws.amazon.com/api/dirs/items/search?item.directoryId=whats-new&item." +
				"locale=en_US&sort_by=item.additionalFields.postDateTime&sort_order=desc&page=0&size=25",
		},
		{
			name: "Test full URL with page 1 & size 50",
			args: args{
				page: 1,
				size: 50,
			},
			want: "https://aws.amazon.com/api/dirs/items/search?item.directoryId=whats-new&item." +
				"locale=en_US&sort_by=item.additionalFields.postDateTime&sort_order=desc&page=1&size=50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildURL(tt.args.page, tt.args.size); got != tt.want {
				t.Errorf("buildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
