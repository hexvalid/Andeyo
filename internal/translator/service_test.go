package translator

import "testing"

func TestBuildReplyUnauthorizedChat(t *testing.T) {
	t.Parallel()

	svc := NewService([]int64{10})
	reply := svc.BuildReply(99, "selam yağmur")

	if reply != UnauthorizedMessage {
		t.Fatalf("unexpected unauthorized reply: %q", reply)
	}
}

func TestBuildReplyEncodesLatinInput(t *testing.T) {
	t.Parallel()

	svc := NewService([]int64{10})
	reply := svc.BuildReply(10, "Ab c")

	want := "🎄🔩     🌜"
	if reply != want {
		t.Fatalf("unexpected encoded text: got %q want %q", reply, want)
	}
}

func TestBuildReplyDecodesAndeyoInput(t *testing.T) {
	t.Parallel()

	svc := NewService([]int64{10})
	reply := svc.BuildReply(10, "🎄🔩     🌜")

	want := "ab c"
	if reply != want {
		t.Fatalf("unexpected decoded text: got %q want %q", reply, want)
	}
}
