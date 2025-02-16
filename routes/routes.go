// Names based roughly on:
// https://html.spec.whatwg.org/multipage/parsing.html

package routes

import (
	"unicode"
)

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

func Route(glyph rune, prevKind StepKind) StepKind {
	switch prevKind {
	case Attr:
		return getKindFromAttribute(glyph)
	case AttrMapInjection:
		return getKindFromInjection(glyph)
	case AttrQuote:
		return getKindFromAttributeQuote(glyph)
	case AttrQuoteClosed:
		return getKindFromAttributeQuoteClosed(glyph)
	case AttrSetter:
		return getKindFromAttributeSetter(glyph)
	case AttrValue:
		return getKindFromAttributeQuote(glyph)
	case AttrValueUnquoted:
		return getKindFromAttributeValueUnquoted(glyph)
	case DescendantInjection:
		return getKindFromInjection(glyph)
	case Element:
		return getKindFromElement(glyph)
	case ElementSpace:
		return getKindFromElementSpace(glyph)
	case EmptyElement:
		return getKindFromEmptyElement(glyph)
	case InjectionSpace:
		return getKindFromInjection(glyph)
	case Tag:
		return getKindFromTag(glyph)
	case TailElementSolidus:
		return getKindFromTailElementSolidus(glyph)
	case TailElementSpace:
		return getKindFromTailElementSpace(glyph)
	case TailTag:
		return getKindFromTailTag(glyph)
	default:
		return getKindFromInitial(glyph)
	}
}

func getKindFromAttribute(glyph rune) StepKind {
	if unicode.IsSpace(glyph) {
		return ElementSpace
	}
	switch glyph {
	case '=':
		return AttrSetter
	case '>':
		return ElementClosed
	case '/':
		return EmptyElement
	case '{':
		return AttrMapInjection
	default:
		return Attr
	}
}

func getKindFromInjection(glyph rune) StepKind {
	switch glyph {
	case '}':
		return InjectionConfirmed
	default:
		return InjectionSpace
	}
}

func getKindFromAttributeQuote(glyph rune) StepKind {
	switch glyph {
	case '"':
		return AttrQuoteClosed
	default:
		return AttrValue
	}
}

func getKindFromAttributeQuoteClosed(glyph rune) StepKind {
	switch glyph {
	case '>':
		return ElementClosed
	case '/':
		return EmptyElement
	default:
		return ElementSpace
	}
}

func getKindFromAttributeSetter(glyph rune) StepKind {
	if unicode.IsSpace(glyph) {
		return AttrSetter
	}

	switch glyph {
	case '"':
		return AttrQuote
	default:
		return AttrValueUnquoted
	}
}

func getKindFromAttributeValueUnquoted(glyph rune) StepKind {
	if unicode.IsSpace(glyph) {
		return ElementSpace
	}

	switch glyph {
	case '>':
		return ElementClosed
	default:
		return AttrValueUnquoted
	}
}

func getKindFromElement(glyph rune) StepKind {
	if unicode.IsSpace(glyph) {
		return Element
	}

	switch glyph {
	case '>':
		return Fragment
	case '/':
		return TailElementSolidus
	default:
		return Tag
	}
}

func getKindFromElementSpace(glyph rune) StepKind {
	if unicode.IsSpace(glyph) {
		return ElementSpace
	}

	switch glyph {
	case '>':
		return ElementClosed
	case '/':
		return EmptyElement
	case '{':
		return AttrMapInjection
	default:
		return Attr
	}
}

func getKindFromEmptyElement(glyph rune) StepKind {
	switch glyph {
	case '>':
		return EmptyElementClosed
	default:
		return EmptyElement
	}
}

func getKindFromTag(glyph rune) StepKind {
	if unicode.IsSpace(glyph) {
		return ElementSpace
	}

	switch glyph {
	case '>':
		return ElementClosed
	case '/':
		return EmptyElement
	default:
		return Tag
	}
}

func getKindFromTailElementSolidus(glyph rune) StepKind {
	if unicode.IsSpace(glyph) {
		return TailElementSolidus
	}

	switch glyph {
	case '>':
		return FragmentClosed
	default:
		return TailTag
	}
}

func getKindFromTailTag(glyph rune) StepKind {
	if unicode.IsSpace(glyph) {
		return TailElementSpace
	}

	switch glyph {
	case '>':
		return TailElementClosed
	default:
		return TailTag
	}
}

func getKindFromTailElementSpace(glyph rune) StepKind {
	switch glyph {
	case '>':
		return TailElementClosed
	default:
		return TailElementSpace
	}
}

func getKindFromInitial(glyph rune) StepKind {
	switch glyph {
	case '<':
		return Element
	case '{':
		return DescendantInjection
	default:
		return Text
	}
}
