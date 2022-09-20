package sophosutm

import "testing"

func TestVersion(t *testing.T) {
    target, err := NewClient("", "")
    if err != nil {
        t.Fatal(err)
    }
    actual, err := target.Status().Version()
    if err != nil {
        t.Fatal(err)
    }
    if actual == nil {
        t.Fatal("actual is nil")
    }
    if actual.Utm == "" {
        t.Fatal("actual.Utm is empty")
    }
    if actual.Restd == "" {
        t.Fatal("actual.Restd is empty")
    }
    t.Logf("actual: %+v", actual)
}
