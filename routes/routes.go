package routes

type StepKind int

const (
	AttrQuoteClosed StepKind = iota + 1
	AttrQuote
	AttrMapInjection
	AttrSetter
	AttrValue
	AttrValueUnquoted
	Attr
	TailElementClosed
	TailElementSolidus
	TailElementSpace
	TailTag
	DescendantInjection
	FragmentClosed
	Fragment
	EmptyElementClosed
	EmptyElement
	Initial
	InjectionConfirmed
	InjectionSpace
	ElementClosed
	ElementSpace
	Element
	Tag
	Text
	// needed for comments and scripts
	AltText
	AltTextCloseSequence
	CommentText
)
