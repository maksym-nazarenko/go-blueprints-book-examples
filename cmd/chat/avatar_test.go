package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	_, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should retorn ErrNoAvatarURL when no value present")
	}
	testURL := "http://url-to=gravater/"
	client.userData = map[string]interface{}{
		"avatar_url": testURL,
	}
	url, err := authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GetAvatar should return no error")
	}
	if url != testURL {
		t.Error("GetURL return invalid value")
	}
}
