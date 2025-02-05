package coyote

type RulesetInterface interface {
	GetInitialNamespace() string
	TagIsComment(string) bool
	GetCloseSequenceFromAltTextTag(string) string
	GetTagFromCloseSequence(string) string
	RespectIndentation(string) bool
	TagIsBannedEl(string) bool
	TagIsVoidEl(string) bool
	TagIsNamespaceEl(string) bool
	TagIsPreservedTextEl(string) bool
	TagIsInlineEl(string) bool
}
