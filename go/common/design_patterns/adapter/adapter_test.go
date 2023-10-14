package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var player MusicPlayer
	player = &PlayerAdaptor{}
	player.play("mp3", "死了都要爱")
	player.play("wma", "滴滴")
	player.play("mp4", "复仇者联盟")
}
