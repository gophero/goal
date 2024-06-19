package twitter_test

import (
	"reflect"
	"testing"

	"github.com/gophero/goal/twitter"
)

var followoApi = &twitter.OAuth2FollowApi{}

var target_userid = "1655224265766240257"
var self_userid = "1519255422300868609"

func TestOauth2FollowApi_Follow(t *testing.T) {
	type args struct {
		accessToken  string
		userId       string
		targetUserId string
	}
	tests := []struct {
		name    string
		o       *twitter.OAuth2FollowApi
		args    args
		want    twitter.FollowRet
		wantErr bool
	}{
		// Add test cases.
		{name: "success", o: followoApi, args: args{
			userId:       self_userid,
			accessToken:  testEnv.accessToken,
			targetUserId: target_userid,
		}, want: twitter.FollowRet{Following: true, PendingFollow: false}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &twitter.OAuth2FollowApi{}
			got, err := o.Follow(tt.args.accessToken, tt.args.userId, tt.args.targetUserId)
			if (err != nil) != tt.wantErr {
				t.Errorf("OAuth2FollowApi.Follow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuth2FollowApi.Follow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOauth2FollowApi_UnFollow(t *testing.T) {
}
