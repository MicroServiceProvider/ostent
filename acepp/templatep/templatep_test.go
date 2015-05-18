package templatep

import (
	"strings"
	"testing"
	"text/template"
)

func TestDotted(t *testing.T) {
	abc := "a.b.c"
	words := strings.Split(abc, ".")
	d := Dotted{}
	d.Append(words)
	d.Append(strings.Split(abc+".1", "."))
	d.Append(strings.Split(abc+".2", "."))
	if x := d.Find(append(words, "z")); x != nil {
		t.Errorf("Find returned non-nil")
	}
	if x := d.Find(words); x == nil {
		t.Errorf("Find returned nil")
	} else if s := x.Notation(); s != abc {
		t.Errorf("Notation: %q (expected %q)", s, abc)
	}
	if expected, s := `[]
  [a]
    [b]
      [c]
        [1]
        [2]
`, d.GoString(); s != expected {
		t.Errorf("GoString: %s (expected %s)", s, expected)
	}
}

func TestMkmap(t *testing.T) {
	words := strings.Split("a.b.c", ".")
	d := Dotted{}
	d.Append(words)
	l1 := d.Find(words)
	if l1 == nil {
		t.Errorf("Dotted.Find returned nil")
	}
	l1.Ranged = true
	l1.Keys = []string{"z"}
	l1.Decl = "DECL"

	w2 := strings.Split("a.b.z", ".")
	d.Append(w2)
	if l2 := d.Find(w2); l2 != nil {
		l2.Ranged = true
	}

	v := mkmap(d, false, -1)
	c := v.(Hash)["a"].(Hash)["b"].(Hash)["c"].([]map[string]string)[0]
	if expected := "{DECL.z}"; c["z"] != expected {
		t.Errorf("mkmap result mismatch: %q (expected %q)", c["z"], expected)
	}
	if z := v.(Hash)["a"].(Hash)["b"].(Hash)["z"].([]string); len(z) != 0 {
		t.Errorf("z is expected to be empty: %+v\n", z)
	}
}

func ExecutewDottedTest(t *testing.T, tm *template.Template, expected string) {
	d := Dotted{}
	d.Append([]string{"a"}) // non-string result from mkmap
	h := mkmap(d, true, 0)
	if _, ok := h.(string); ok {
		t.Errorf("mkmap expected to return non-string on %+v input", d)
	}
	s, err := StringExecute(tm, h)
	if err != nil {
		t.Fatal(err)
	}
	if s != expected {
		t.Errorf("Execute with mkmap: %q (expected %q)", s, expected)
	}
}

func TestDot(t *testing.T) {
	define := `[[define "define_withdot"]][[with .OVERRIDE]][[.]][[else]][[end]][[end]]`
	text1 := define + `[[template "define_withdot" .]]`
	text2 := define + `[[template "define_withdot" (dot . "rows")]]`
	tm1, err1 := template.New("withdot1").Funcs(DotFuncs).Delims("[[", "]]").Parse(text1)
	tm2, err2 := template.New("withdot2").Funcs(DotFuncs).Delims("[[", "]]").Parse(text2)
	if err1 != nil {
		t.Fatal(err1)
	}
	if err2 != nil {
		t.Fatal(err2)
	}
	ExecutewDottedTest(t, tm1, "")
	ExecutewDottedTest(t, tm2, "{rows}")
}

func DotValueText(t *testing.T, in, expected string) {
	h := dot(Hash{}, in)
	d, ok := h["OVERRIDE"]
	if !ok {
		t.Errorf("Getting \"dot\" from `dot' result is not okd.")
	}
	dv, ok := d.(dotValue)
	if !ok {
		t.Error("Cannot cast to dotValue")
	}
	if s := dv.String(); s != expected {
		t.Errorf("dotValue mismatch: %q (expected %q)", s, expected)
	}
	if v, ok := (*dv.hashp)["OVERRIDE"]; ok || v != nil {
		t.Errorf("Getting \"dot\" from `dot' result is okd: %+v.", v)
	}
}

func TestDotValue(t *testing.T) {
	DotValueText(t, "KEY", "{KEY}")
	DotValueText(t, "aHTML", "<span dangerouslySetInnerHTML={{__html: aHTML}} />")
}
