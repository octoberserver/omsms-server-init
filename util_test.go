package main

import "testing"

func expectUrl(t *testing.T, str string) {
	if !isURL(str) {
		t.Fatalf("Expected %s to be an url", str)
	}
}

func expectNotUrl(t *testing.T, str string) {
	if isURL(str) {
		t.Fatalf("Expected %s to not be an url", str)
	}
}

func TestIsURL1(t *testing.T) {
	expectUrl(t, "https://ducky")
}
func TestIsURL2(t *testing.T) {
	expectUrl(t, "https://ducky.com")
}
func TestIsURL3(t *testing.T) {
	expectUrl(t, "https://sub.ducky.com")
}
func TestIsURL4(t *testing.T) {
	expectUrl(t, "https://ducky.com/ducky")
}
func TestIsURL5(t *testing.T) {
	expectUrl(t, "https://ducky.net/no-ducky?quack=true&foo=bar")
}
func TestIsURL6(t *testing.T) {
	expectUrl(t, "http://ducky")
}
func TestIsNotURL1(t *testing.T) {
	expectNotUrl(t, "ssh://ducky")
}
func TestIsNotURL2(t *testing.T) {
	expectNotUrl(t, "ssh://ducky.com")
}
func TestIsNotURL3(t *testing.T) {
	expectNotUrl(t, "ssh://sub.ducky.com")
}
func TestIsNotURL4(t *testing.T) {
	expectNotUrl(t, "ssh://ducky.com/ducky")
}
func TestIsNotURL5(t *testing.T) {
	expectNotUrl(t, "https://ducky.net//no-ducky?quack=true&foo=bar")
}
func TestIsNotURL6(t *testing.T) {
	expectNotUrl(t, "https:///ducky.net/no-ducky")
}
func TestIsNotURL7(t *testing.T) {
	expectNotUrl(t, "https:///ducky.net?/no-ducky")
}

func shouldMatch(t *testing.T, str string, strs []string) {
	if !checkStringMatches(str, strs) {
		t.Fatalf("Expected %s to match in options: %s", str, strs)
	}
}

func TestCheckStringMatches1(t *testing.T) {
	shouldMatch(t, "str1", []string{"str1", "str2", "str3"})
}
func TestCheckStringMatches2(t *testing.T) {
	shouldMatch(t, "goose", []string{"duck", "goose", "goose"})
}
func TestCheckStringMatches3(t *testing.T) {
	shouldMatch(t, "duck", []string{"duck", "ducky", "duckery"})
}
