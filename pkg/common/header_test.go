package common

import (
	"strings"
	"testing"
)

type CanonicalKeyTestCase = struct {
	input    string
	expected string
}

type HeaderAction = struct {
	action   string
	key      string
	values   []string
	expected interface{}
}

type HeaderTestCase = struct {
	initial Header
	actions []HeaderAction
	desc    string
}

var canonicalTestSet = []CanonicalKeyTestCase{
	{"content-type", "Content-Type"},
	{"Content-type", "Content-Type"},
	{"content-length", "Content-Length"},
	{"content-Length", "Content-Length"},
	{"content-LENGTH", "Content-Length"},
	{"CONtent-length", "Content-Length"},
	{"pragma", "Pragma"},
	{"PRAgma", "Pragma"},
}

var headerTestSet = []HeaderTestCase{
	{
		initial: Header{},
		actions: []HeaderAction{
			{
				action:   "Has",
				key:      "Content-Type",
				values:   nil,
				expected: false,
			},
		},
		desc: "Simple check if existed == false",
	},
	{
		initial: Header{"Content-Type": []string{"application/json"}},
		actions: []HeaderAction{
			{
				action:   "Has",
				key:      "Content-Type",
				values:   nil,
				expected: true,
			},
		},
		desc: "Simple check if existed == true",
	},
	{
		initial: Header{"Content-Type": []string{"application/json"}},
		actions: []HeaderAction{
			{
				action:   "Get",
				key:      "Content-Type",
				values:   nil,
				expected: "application/json",
			},
		},
		desc: "Simple Get",
	},
	{
		initial: Header{"Content-Type": []string{"application/json"}},
		actions: []HeaderAction{
			{
				action:   "Set",
				key:      "Content-Type",
				values:   []string{"application/text"},
				expected: nil,
			},
			{
				action:   "Get",
				key:      "Content-Type",
				values:   nil,
				expected: "application/text",
			},
		},
		desc: "Simple Set/Get",
	},
	{
		initial: Header{"Content-Type": []string{"application/json"}},
		actions: []HeaderAction{
			{
				action:   "SetOne",
				key:      "content-type",
				values:   []string{"application/text"},
				expected: nil,
			},
			{
				action:   "Get",
				key:      "Content-Type",
				values:   nil,
				expected: "application/text",
			},
		},
		desc: "Simple SetOne/Get (with canonical key transformation)",
	},
	{
		initial: Header{"Content-Type": []string{"application/json"}},
		actions: []HeaderAction{
			{
				action:   "Add",
				key:      "Content-Type",
				values:   []string{"application/text"},
				expected: nil,
			},
			{
				action:   "Get",
				key:      "Content-Type",
				values:   nil,
				expected: "application/json",
			},
		},
		desc: "Simple Add/Get",
	},
	{
		initial: Header{"Content-Type": []string{"application/json"}},
		actions: []HeaderAction{
			{
				action:   "Add",
				key:      "Content-Type",
				values:   []string{"application/text"},
				expected: nil,
			},
			{
				action:   "GetAll",
				key:      "Content-Type",
				values:   nil,
				expected: "application/json;application/text",
			},
		},
		desc: "Simple Add/GetAll",
	},
	{
		initial: Header{"Content-Type": []string{"application/json"}},
		actions: []HeaderAction{
			{
				action:   "Delete",
				key:      "content-type",
				values:   nil,
				expected: nil,
			},
			{
				action:   "Get",
				key:      "Content-Type",
				values:   nil,
				expected: "",
			},
		},
		desc: "Simple Delete/Get",
	},
}

func TestCanonicalKey(t *testing.T) {
	for _, tc := range canonicalTestSet {
		ck := CanonicalKey(tc.input)
		if ck != tc.expected {
			t.Errorf("Expected %q. Got %q", tc.expected, ck)
		}
	}
}

func TestHeaderModification(t *testing.T) {
	for _, tc := range headerTestSet {
		h := tc.initial
		for _, a := range tc.actions {
			switch a.action {
			case "Has":
				if out := h.Has(a.key); out != a.expected {
					t.Errorf("Expected %t. Got %t", a.expected, out)
				}
			case "Get":
				if out := h.Get(a.key); out != a.expected {
					t.Errorf("Expected %v. Got %v", a.expected, out)
				}
			case "GetAll":
				if out := h.GetAll(a.key); strings.Join(out, ";") != a.expected {
					t.Errorf("Expected %v. Got %v", a.expected, strings.Join(out, ";"))
				}
			case "Set":
				h.Set(a.key, a.values)
			case "SetOne":
				h.SetOne(a.key, a.values[0])
			case "Add":
				h.Add(a.key, a.values[0])
			case "Delete":
				h.Delete(a.key)
			}
		}
	}
}
