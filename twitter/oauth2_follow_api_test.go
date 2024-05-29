package twitter_test

import (
	"reflect"
	"testing"

	"github.com/gophero/goal/twitter"
)

var followoApi = &twitter.Oauth2FollowApi{}

var target_userid = "1779099104049876992"
var self_userid = "1776491059318792192"

func TestOauth2FollowApi_Follow(t *testing.T) {
	type args struct {
		accessToken  string
		userId       string
		targetUserId string
	}
	tests := []struct {
		name    string
		o       *twitter.Oauth2FollowApi
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
			o := &twitter.Oauth2FollowApi{}
			got, err := o.Follow(tt.args.accessToken, tt.args.userId, tt.args.targetUserId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Oauth2FollowApi.Follow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Oauth2FollowApi.Follow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOauth2FollowApi_UnFollow(t *testing.T) {
}
