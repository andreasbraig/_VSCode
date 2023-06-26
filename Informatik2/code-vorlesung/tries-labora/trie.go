package trieslabora

type Element struct {
	Key   string
	left  *Element
	right *Element
	Value string
}

func New() *Element {
	return &Element{}
}

func (e Element) HasValue() bool {
	return e.Value != ""
}

func (e *Element) SetValue(value string) {
	//TODO
	e.Value = value
}

func (e *Element) AddValueForPath(path, value string) {

	if path == "" {

		e.SetValue(value)
		return
	}

	head, tail := path[:1], path[1:]
	if head == "@" {
		e.AddValueForPath(tail, value)
		return
	}

	if head == "a" {

		if e.left == nil {
			e.left = New()
		}
		e.left.AddValueForPath(tail, value)

	} else {

		if e.right == nil {
			e.right = New()
		}
		e.right.AddValueForPath(tail, value)

	}

}
