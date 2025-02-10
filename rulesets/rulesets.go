package rulesets

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

type ServerRules struct{}

func (s ServerRules) GetInitialNamespace() string {
	return "html"
}

func (s ServerRules) TagIsComment(tag string) bool {
	return "!--" == tag
}

type ClientRules struct{}

func (c ClientRules) GetInitialNamespace() string {
	return "html"
}

func (s ClientRules) TagIsComment(tag string) bool {
	return "!--" == tag
}

type XmlRules struct{}

func (x XmlRules) GetInitialNamespace() string {
	return "xml"
}

func (s XmlRules) TagIsComment(tag string) bool {
	return "!--" == tag
}

func (s XmlRules) GetCloseSequenceFromTag(tag string) string {
	switch tag {
	case "!--":
		return "-->"
	case "!CDATA[[":
		return "]]>"
	default:
		return ""
	}
}

func (s XmlRules) GetTagFromCloseSequence(tag string) string {
	switch tag {
	case "-->":
		return "!--"
	case "]]>":
		return "!CDATA[["
	default:
		return ""
	}
}

func (s XmlRules) RespectIndentation(tag string) bool {
	return true
}

func (s XmlRules) TagIsBannedEl(tag string) bool {
	return false
}

func (s XmlRules) TagIsVoidEl(tag string) bool {
	return false
}

func (s XmlRules) TagIsNamespaceEl(tag string) bool {
	return false
}

func (s XmlRules) TagIsPreservedTextEl(tag string) bool {
	return false
}

func (s XmlRules) TagIsInlineEl(tag string) bool {
	return false
}

func isBannedEl(tag string) bool {
	switch tag {
	case "acronym":
		return true
	case "big":
		return true
	case "center":
		return true
	case "content":
		return true
	case "dir":
		return true
	case "font":
		return true
	case "frame":
		return true
	case "framset":
		return true
	case "image":
		return true
	case "marquee":
		return true
	case "menuitem":
		return true
	case "nobr":
		return true
	case "noembed":
		return true
	case "noframes":
		return true
	case "param":
		return true
	case "plaintext":
		return true
	case "rb":
		return true
	case "rtc":
		return true
	case "shadow":
		return true
	case "strike":
		return true
	case "tt":
		return true
	case "xmp":
		return true
	default:
		return false
	}
}

func isVoidEl(tag string) bool {
	switch tag {
	case "!--":
		return true
	case "!DOCTYPE":
		return true
	case "area":
		return true
	case "base":
		return true
	case "br":
		return true
	case "col":
		return true
	case "embed":
		return true
	case "hr":
		return true
	case "img":
		return true
	case "input":
		return true
	case "link":
		return true
	case "meta":
		return true
	case "param":
		return true
	case "source":
		return true
	case "track":
		return true
	case "wbr":
		return true
	default:
		return false
	}
}

func isNamespaceEl(tag string) bool {
	switch tag {
	case "html":
		return true
	case "math":
		return true
	case "svg":
		return true
	default:
		return false
	}
}

func isPreservedTextEl(tag string) bool {
	return "pre" == tag
}

func isInlineEl(tag string) bool {
	switch tag {
	case "abbr":
		return true
	case "area":
		return true
	case "audio":
		return true
	case "b":
		return true
	case "bdi":
		return true
	case "bdo":
		return true
	case "cite":
		return true
	case "code":
		return true
	case "data":
		return true
	case "dfn":
		return true
	case "em":
		return true
	case "embed":
		return true
	case "i":
		return true
	case "iframe":
		return true
	case "img":
		return true
	case "kbd":
		return true
	case "map":
		return true
	case "mark":
		return true
	case "object":
		return true
	case "picture":
		return true
	case "portal":
		return true
	case "q":
		return true
	case "rp":
		return true
	case "rt":
		return true
	case "ruby":
		return true
	case "s":
		return true
	case "samp":
		return true
	case "small":
		return true
	case "source":
		return true
	case "span":
		return true
	case "strong":
		return true
	case "sub":
		return true
	case "sup":
		return true
	case "time":
		return true
	case "track":
		return true
	case "u":
		return true
	case "var":
		return true
	case "video":
		return true
	case "wbr":
		return true
	default:
		return false
	}
}
