package tag_info

import "github.com/w-lfpup/coyote-go/rulesets"

type TextFormat int

const (
	Block TextFormat = iota + 1
	Initial
	Inline
	Root
)

type TagInfo struct {
	Namespace string
	Tag string
	TextFormat TextFormat // "block" | "Initial"
	IndentCount int16
	VoidEl bool
	InlineEl bool
	PreservedTextPath bool
	BannedPath bool
}

func New(rules rulesets.RulesetInterface, tag string) TagInfo {
	namespace := rules.GetInitialNamespace();
	if rules.TagIsNamespaceEl(tag) {
		namespace = tag
	}

	return TagInfo {
		Namespace: namespace,
		Tag: tag,
		TextFormat: Root,
		IndentCount: 0,
		VoidEl: rules.TagIsVoidEl(tag),
		InlineEl: rules.TagIsInlineEl(tag),
		PreservedTextPath: rules.TagIsPreservedTextEl(tag),
		BannedPath: rules.TagIsBannedEl(tag),
	}
}

func From(rules rulesets.RulesetInterface, prevTagInfo TagInfo, tag string) TagInfo {
	tagInfo := New(rules, tag)

	tagInfo.Namespace = prevTagInfo.Namespace
	tagInfo.IndentCount = prevTagInfo.IndentCount
	tagInfo.TextFormat = Initial

	if rules.TagIsNamespaceEl(tag) {
		tagInfo.Namespace = tag
	}

	if rules.TagIsPreservedTextEl(prevTagInfo.Tag) {
		tagInfo.PreservedTextPath = true
	}

	if rules.TagIsBannedEl(tag) {
		tagInfo.BannedPath = true
	}

	if (!rules.TagIsVoidEl(prevTagInfo.Tag) && !rules.TagIsInlineEl(tag)) {
		tagInfo.IndentCount += 1
	}

	return tagInfo
}

// rules, tag

// create()

// from()
