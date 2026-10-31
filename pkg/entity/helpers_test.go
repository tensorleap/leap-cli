package entity

import "testing"

type testEntity struct {
	name string
	id   string
}

var testEntityDesc = NewEntityDescriptor[testEntity](
	"secret",
	"secrets",
	func(e *testEntity) string { return e.name },
	func(e *testEntity) string { return e.id },
)

// The message reaches users verbatim (`leap secrets delete <name>` returns it),
// so its wording is the behaviour under test — it read "not found foo name:
// secret" while the format arguments were reversed.
func TestGetEntityByDisplayNameNotFoundMessage(t *testing.T) {
	_, err := GetEntityByDisplayName("foo", []testEntity{{name: "bar", id: "1"}}, testEntityDesc)
	if err == nil {
		t.Fatal("expected an error for a name that is absent")
	}
	if got, want := err.Error(), `no secret found with name "foo"`; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestGetEntityByDisplayNameFound(t *testing.T) {
	found, err := GetEntityByDisplayName("bar", []testEntity{{name: "bar", id: "1"}}, testEntityDesc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found.id != "1" {
		t.Fatalf("got id %q, want %q", found.id, "1")
	}
}
