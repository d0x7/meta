// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

package meta

import (
	"fmt"
	u "net/url"
	"runtime"
	"testing"
	"time"
)

func TestMeta(t *testing.T) { //nolint:funlen
	t.Parallel()

	expectedDate := time.Date(2019, 8, 23, 18, 0, 0, 0, time.UTC)
	expectedURL := u.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/page",
	}

	tests := []struct {
		flags    map[string]string
		assertfn func(*testing.T, *info)
		panics   bool
	}{
		{
			// Validate that the test program does not panic when no
			// definitions are given.
		},
		{
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, runtime.GOARCH, actual.Arch)
			},
		},
		{
			// Value for xiam.li/meta.author.
			flags: map[string]string{
				"xiam.li/meta.author": "Jane Doe <jdoe@example.com>",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "Jane Doe", actual.Author)
				equalString(t, "jdoe@example.com", actual.AuthorEmail)
			},
		},
		{
			// Value for xiam.li/meta.author_url that is valid.
			flags: map[string]string{
				"xiam.li/meta.author_url": "https://example.com/page",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalURL(t, &expectedURL, actual.AuthorURL)
			},
		},
		{
			// Value for xiam.li/meta.author_url that causes a panic.
			flags: map[string]string{
				"xiam.li/meta.author_url": "example.com/page",
			},
			panics: true,
		},
		{
			// Value for xiam.li/meta.copyright.
			flags: map[string]string{
				"xiam.li/meta.copyright": "2021 Jane Doe",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "2021 Jane Doe", actual.Copyright)
			},
		},
		{
			// Value for xiam.li/meta.date that is valid.
			flags: map[string]string{
				"xiam.li/meta.date": "Fri, 23 Aug 2019 11:00:00 -0700",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalTime(t, &expectedDate, actual.Date)
				equalString(t, "2019-08-23T18:00:00Z", actual.DateFormat)
			},
		},
		{
			// Value for xiam.li/meta.date that causes a panic.
			flags: map[string]string{
				"xiam.li/meta.date": "tomorrow",
			},
			panics: true,
		},
		{
			// Value for xiam.li/meta.desc.
			flags: map[string]string{
				"xiam.li/meta.desc": "Example description",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "Example description", actual.Description)
			},
		},
		{
			// Value for xiam.li/meta.dev.
			flags: map[string]string{
				"xiam.li/meta.dev": "true",
			},
			assertfn: func(t *testing.T, actual *info) {
				if true != actual.Development {
					t.Fatalf("expected %v but got %v", true, actual)
				}
			},
		},
		{
			// Value for xiam.li/meta.docs that is valid.
			flags: map[string]string{
				"xiam.li/meta.docs": "https://example.com/page",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalURL(t, &expectedURL, actual.Docs)
			},
		},
		{
			// Value for xiam.li/meta.docs that causes a panic.
			flags: map[string]string{
				"xiam.li/meta.docs": "example.com/page",
			},
			panics: true,
		},
		{
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, runtime.Version(), actual.Go)
			},
		},
		{
			// Value for xiam.li/meta.license.
			flags: map[string]string{
				"xiam.li/meta.license": "MIT",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "MIT", actual.License)
			},
		},
		{
			// Value for xiam.li/meta.license_url that is valid.
			flags: map[string]string{
				"xiam.li/meta.license_url": "https://example.com/page",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalURL(t, &expectedURL, actual.LicenseURL)
			},
		},
		{
			// Value for xiam.li/meta.license_url that causes a panic.
			flags: map[string]string{
				"xiam.li/meta.license_url": "example.com/page",
			},
			panics: true,
		},
		{
			// Value for xiam.li/meta.name.
			flags: map[string]string{
				"xiam.li/meta.name": "demo-app",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "demo-app", actual.Name)
			},
		},
		{
			// Value for xiam.li/meta.note.
			flags: map[string]string{
				"xiam.li/meta.note": "Example note",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "Example note", actual.Note)
			},
		},
		{
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, runtime.GOOS, actual.OS)
			},
		},
		{
			// Value for xiam.li/meta.sha that is valid.
			flags: map[string]string{
				"xiam.li/meta.sha": "bb2fecbb4a287ea4c1f9887ca86dd0eb7ff28ec6",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "bb2fecbb4a287ea4c1f9887ca86dd0eb7ff28ec6", actual.SHA)
				equalString(t, "bb2fecb", actual.ShortSHA)
			},
		},
		{
			// Value for xiam.li/meta.sha that causes a panic.
			flags: map[string]string{
				"xiam.li/meta.sha": "HEAD",
			},
			panics: true,
		},
		{
			// Value for xiam.li/meta.src that is valid.
			flags: map[string]string{
				"xiam.li/meta.src": "https://example.com/page",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalURL(t, &expectedURL, actual.Source)
			},
		},
		{
			// Value for xiam.li/meta.src that causes a panic.
			flags: map[string]string{
				"xiam.li/meta.src": "example.com/page",
			},
			panics: true,
		},
		{
			// Value for xiam.li/meta.title.
			flags: map[string]string{
				"xiam.li/meta.title": "Example title",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "Example title", actual.Title)
			},
		},
		{
			// Value for xiam.li/meta.url that is valid.
			flags: map[string]string{
				"xiam.li/meta.url": "https://example.com/page",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalURL(t, &expectedURL, actual.URL)
			},
		},
		{
			// Value for xiam.li/meta.url that causes a panic.
			flags: map[string]string{
				"xiam.li/meta.url": "example.com/page",
			},
			panics: true,
		},
		{
			// Value for xiam.li/meta.version.
			flags: map[string]string{
				"xiam.li/meta.version": "v1.2.3-rc.456+build.789",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "v1.2.3-rc.456+build.789", actual.Version)
				equalString(t, "1", actual.VersionMajor)
				equalString(t, "2", actual.VersionMinor)
				equalString(t, "3", actual.VersionPatch)
				equalString(t, "rc.456", actual.VersionPreRelease)
				equalString(t, "build.789", actual.VersionBuild)
			},
		},
		{
			// Value for xiam.li/meta.version.
			flags: map[string]string{
				"xiam.li/meta.version": "v1.2.3",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "v1.2.3", actual.Version)
				equalString(t, "1", actual.VersionMajor)
				equalString(t, "2", actual.VersionMinor)
				equalString(t, "3", actual.VersionPatch)
				equalString(t, "", actual.VersionPreRelease)
				equalString(t, "", actual.VersionBuild)
			},
		},
		{
			// Value for xiam.li/meta.version.
			flags: map[string]string{
				"xiam.li/meta.version": "latest",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalString(t, "latest", actual.Version)
				equalString(t, "", actual.VersionMajor)
				equalString(t, "", actual.VersionMinor)
				equalString(t, "", actual.VersionPatch)
				equalString(t, "", actual.VersionPreRelease)
				equalString(t, "", actual.VersionBuild)
			},
		},
	}

	for index, test := range tests {
		test := test

		t.Run(fmt.Sprint(index), func(t *testing.T) {
			t.Parallel()

			actual, failed := execTestJSON(t, test.flags)
			switch {
			case failed && !test.panics:
				t.Error("expected test success but got failure")
			case !failed && test.panics:
				t.Error("expected test failure but got success")
			case !failed && !test.panics && test.assertfn != nil:
				test.assertfn(t, actual)
			}
		})
	}
}
