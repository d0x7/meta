// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

// Package meta provides common application metadata for use with go build and
// ldflags. This package is intended to be imported, where variable values can
// be set by using -X arguments, tp the -ldflags argument, when running go
// build. See https://pkg.go.dev/cmd/go and https://pkg.go.dev/cmd/link.
//
// List of variable names:
//
//	xiam.li/meta.author
//	xiam.li/meta.author_url
//	xiam.li/meta.copyright
//	xiam.li/meta.date
//	xiam.li/meta.desc
//	xiam.li/meta.dev
//	xiam.li/meta.docs
//	xiam.li/meta.license
//	xiam.li/meta.license_url
//	xiam.li/meta.name
//	xiam.li/meta.note
//	xiam.li/meta.sha
//	xiam.li/meta.src
//	xiam.li/meta.title
//	xiam.li/meta.url
//	xiam.li/meta.version
package meta

import (
	u "net/url"
	"runtime"
	"time"
)

// Arch is the architecture target that the application is running on.
func Arch() string {
	return runtime.GOARCH
}

// author is the name of the application author. May contain their name, email
// address, or optionally both.
//
// Variable name:
//
//	xiam.li/meta.author
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.author=John Doe'"
//	-ldflags "-X 'xiam.li/meta.author=jdoe@example.com'"
//	-ldflags "-X 'xiam.li/meta.author=Jane Doe <jdoe@example.com>'"
var author string

var authorParsed, authorEmailParsed = mustAuthor("xiam.li/meta.author", author)

// Author is the name of the application author.
func Author() string {
	return authorParsed
}

// AuthorOr is the name of the application author, or the given default value if not set.
func AuthorOr(defaultValue string) string {
	if authorParsed == "" {
		return defaultValue
	}

	return authorParsed
}

// AuthorEmail is the email address for the application author.
func AuthorEmail() string {
	return authorEmailParsed
}

// AuthorEmailOr is the email address for the application author, or the given default value if not set.
func AuthorEmailOr(defaultValue string) string {
	if authorEmailParsed == "" {
		return defaultValue
	}

	return authorEmailParsed
}

// author_url is a URL for the application author. Typically links to the
// author's personal homepage or Github profile.
//
// Variable name:
//
//	xiam.li/meta.author_url
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.author_url=https://example.com/profile'"
var author_url string

var authorURLParsed = mustURL("xiam.li/meta.author_url", author_url)

// AuthorURL is the homepage URL for the application author.
func AuthorURL() *u.URL {
	return authorURLParsed
}

// AuthorURLOr is the homepage URL for the application author, or the given default value if not set.
func AuthorURLOr(defaultValue string) *u.URL {
	if authorURLParsed == nil {
		return mustURL("xiam.li/meta.author_url", defaultValue)
	}

	return authorURLParsed
}

// copyright is the copyright for the application. Typically the name if the
// author or organization, sometimes prefixed with a year or year range.
//
// Variable name:
//
//	xiam.li/meta.copyright
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.copyright=John Doe'"
//	-ldflags "-X 'xiam.li/meta.copyright=2021 Jane Doe'"
//	-ldflags "-X 'xiam.li/meta.copyright=2019-2021 Jim Doe'"
var copyright string

// Copyright is the copyright for the application.
func Copyright() string {
	return copyright
}

// CopyrightOr is the copyright for the application, or the given default value if not set.
func CopyrightOr(defaultValue string) string {
	if copyright == "" {
		return defaultValue
	}

	return copyright
}

// date is the time that the application was built. Supports several common
// formats.
//
// Variable name:
//
//	xiam.li/meta.date
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.date=$(date -R)'"
//	-ldflags "-X 'xiam.li/meta.date=Fri, 23 Aug 2019 11:00:00 -0700'"
//	-ldflags "-X 'xiam.li/meta.date=$(date -Iseconds)'"
//	-ldflags "-X 'xiam.li/meta.date=2019-08-23T11:00:00-07:00'"
//	-ldflags "-X 'xiam.li/meta.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)'"
//	-ldflags "-X 'xiam.li/meta.date=2019-08-23T18:00:00Z'"
var date string

var dateParsed = mustTime("xiam.li/meta.date", date)

// Date is the time at which the application was built.
func Date() *time.Time {
	return dateParsed
}

// DateOr is the time at which the application was built, or the given default value if not set.
func DateOr(defaultValue time.Time) *time.Time {
	if dateParsed == nil {
		return &defaultValue
	}

	return dateParsed
}

// DateFormat is the time at which the application was built, formatted using
// the given layout.
func DateFormat(layout string) string {
	if dateParsed == nil {
		return ""
	}

	return dateParsed.Format(layout)
}

// DateFormatOr is the time at which the application was built, formatted using
// the given layout, or the given default value if not set.
func DateFormatOr(layout string, defaultValue string) string {
	if dateParsed == nil {
		return defaultValue
	}

	return dateParsed.Format(layout)
}

// desc is a description for the application. Typically a longer statement
// describing what the application does.
//
// Variable name:
//
//	xiam.li/meta.desc
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.desc=A super simple demonstration application'"
var desc string

// Description is the description of the application.
func Description() string {
	return desc
}

// DescriptionOr is the description of the application, or the given default value if not set.
func DescriptionOr(defaultValue string) string {
	if desc == "" {
		return defaultValue
	}

	return desc
}

// dev is the development status for the application. An application in
// development mode may indicate that it's using experimental or untested
// features, and should be used with caution.
//
// Variable name:
//
//	xiam.li/meta.dev
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.dev=true'"
var dev string

var devParsed = mustBool("xiam.li/meta.dev", dev)

// Development is the development status for the application.
func Development() bool {
	return devParsed
}

// docs is a URL for application documentation. Typically links to a page where
// a user can find technical documentation.
//
// Variable name:
//
//	xiam.li/meta.docs
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.docs=https://example.com/demo/README.md'"
var docs string

var docsParsed = mustURL("xiam.li/meta.docs", docs)

// Docs is the documentation URL for the application.
func Docs() *u.URL {
	return docsParsed
}

// DocsOr is the documentation URL for the application, or the given default value if not set.
func DocsOr(defaultValue string) *u.URL {
	if docsParsed == nil {
		return mustURL("xiam.li/meta.docs", defaultValue)
	}

	return docsParsed
}

// Go is the version of the Go runtime that the application is running on.
func Go() string {
	return runtime.Version()
}

// license is the license identifier for the application. Should not the full
// license body, but one of the identifiers from https://spdx.org/licenses, so
// that the type of license can be easily determined.
//
// Variable name:
//
//	xiam.li/meta.license
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.license=Apache-2.0'"
//	-ldflags "-X 'xiam.li/meta.license=MIT'"
//	-ldflags "-X 'xiam.li/meta.license=WTFPL'"
var license string

// License is the license identifier for the application.
func License() string {
	return license
}

// LicenseOr is the license identifier for the application, or the given default value if not set.
func LicenseOr(defaultValue string) string {
	if license == "" {
		return defaultValue
	}

	return license
}

// license_url is a URL for the application license. Typically links to a page
// where the verbatim license body is available.
//
// Variable name:
//
//	xiam.li/meta.license_url
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.license_url=https://example.com/demo/LICENSE.txt'"
var license_url string

var licenseURLParsed = mustURL("xiam.li/meta.license_url", license_url)

// LicenseURL is the license URL for the application.
func LicenseURL() *u.URL {
	return licenseURLParsed
}

// LicenseURLOr is the license URL for the application, or the given default value if not set.
func LicenseURLOr(defaultValue string) *u.URL {
	if licenseURLParsed == nil {
		return mustURL("xiam.li/meta.license_url", defaultValue)
	}

	return licenseURLParsed
}

// name is the name of the application. Typically named the same as the binary,
// or for display in an error or help message.
//
// Variable name:
//
//	xiam.li/meta.name
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.name=demo-app'"
var name string

// Name is the name of the application.
func Name() string {
	return name
}

// NameOr is the name of the application, or the given default value if not set.
func NameOr(defaultValue string) string {
	if name == "" {
		return defaultValue
	}

	return name
}

// note is an arbitrary message for the application. Can be used to store a
// message about the build environment, release, etc.
//
// Variable name:
//
//	xiam.li/meta.note
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.note=Built on CI server ...'"
//	-ldflags "-X 'xiam.li/meta.note=This release is dedicated to ...'"
var note string

// Note is an arbitrary message for the application.
func Note() string {
	return note
}

// NoteOr is an arbitrary message for the application, or the given default value if not set.
func NoteOr(defaultValue string) string {
	if note == "" {
		return defaultValue
	}

	return note
}

// OS is the operating system target that the application is running on.
func OS() string {
	return runtime.GOOS
}

// sha is the git SHA that was used to build the application. A 40 character
// "long" SHA should be provided.
//
// Variable name:
//
//	xiam.li/meta.sha
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.sha=bb2fecbb4a287ea4c1f9887ca86dd0eb7ff28ec6'"
//	-ldflags "-X 'xiam.li/meta.sha=$(git rev-parse HEAD)'"
var sha string

var shaParsed = mustSHA("xiam.li/meta.sha", sha)

// SHA is the git SHA used to build the application.
func SHA() string {
	return shaParsed
}

// SHAOr is the git SHA used to build the application, or the given default value if not set.
func SHAOr(defaultValue string) string {
	if shaParsed == "" {
		return defaultValue
	}

	return shaParsed
}

// ShortSHA is the git "short" SHA used to build the application.
func ShortSHA() string {
	if shaParsed == "" {
		return ""
	}

	return shaParsed[:7]
}

// ShortSHAOr is the git "short" SHA used to build the application, or the given default value if not set.
func ShortSHAOr(defaultValue string) string {
	return SHAOr(defaultValue)[:7]
}

// src is a URL for the application source code. Typically links to a
// repository where a user can browse or clone the source code.
//
// Variable name:
//
//	xiam.li/meta.src
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.src=https://example.com/demo.git'"
var src string

var srcParsed = mustURL("xiam.li/meta.src", src)

// Source is the URL for the application source code.
func Source() *u.URL {
	return srcParsed
}

// SourceOr is the URL for the application source code, or the given default value if not set.
func SourceOr(defaultValue string) *u.URL {
	if srcParsed == nil {
		return mustURL("xiam.li/meta.src", defaultValue)
	}

	return srcParsed
}

// title is the title of the application. Typically a full or non-abbreviated
// form of the application name.
//
// Variable name:
//
//	xiam.li/meta.title
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.title=Demo Application'"
var title string

// Title is the title of the application.
func Title() string {
	return title
}

// TitleOr is the title of the application, or the given default value if not set.
func TitleOr(defaultValue string) string {
	if title == "" {
		return defaultValue
	}

	return title
}

// url is a URL for the application homepage. Typically links to a page where a
// user can learn more about the application.
//
// Variable name:
//
//	xiam.li/meta.url
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.url=https://example.com/demo'"
var url string

var urlParsed = mustURL("xiam.li/meta.url", url)

// URL is the homepage URL for the application.
func URL() *u.URL {
	return urlParsed
}

// URLOr is the homepage URL for the application, or the given default value if not set.
func URLOr(defaultValue string) *u.URL {
	if urlParsed == nil {
		return mustURL("xiam.li/meta.src", defaultValue)
	}

	return urlParsed
}

// version is the version slug for the application. The value can be used to
// point back to a specific tag or release. Supports semver, see
// https://semver.org.
//
// Variable name:
//
//	xiam.li/meta.version
//
// Examples:
//
//	-ldflags "-X 'xiam.li/meta.version=development'"
//	-ldflags "-X 'xiam.li/meta.version=v1.0.0'"
//	-ldflags "-X 'xiam.li/meta.version=$(git describe)'"
var version string

// Version is the version slug for the application.
func Version() string {
	return version
}

// VersionOr is the version slug for the application, or the given default value if not set.
func VersionOr(defaultValue string) string {
	if version == "" {
		versionMajor, versionMinor, versionPatch, versionPreRelease,
			versionBuild = mustSemver("xiam.li/version", defaultValue)

		return defaultValue
	}

	return version
}

var versionMajor, versionMinor, versionPatch, versionPreRelease, versionBuild = mustSemver("xiam.li/version", version)

// VersionMajor is the semver major version.
// See https://semver.org.
func VersionMajor() string {
	return versionMajor
}

// VersionMinor is the semver minor version.
// See https://semver.org.
func VersionMinor() string {
	return versionMinor
}

// VersionPatch is the semver patch version.
// See https://semver.org.
func VersionPatch() string {
	return versionPatch
}

// VersionPreRelease is the semver pre-release version.
// See https://semver.org.
func VersionPreRelease() string {
	return versionPreRelease
}

// VersionBuild is the semver build metadata version.
// See https://semver.org.
func VersionBuild() string {
	return versionBuild
}
