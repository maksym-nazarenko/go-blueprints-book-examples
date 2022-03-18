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

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"email": "MyEmailAddress@example.com",
	}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar should not return an error")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("Gravatar wrongly returned %s", url)
	}

}
