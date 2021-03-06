package calendar

// Code generated by askew. DO NOT EDIT.

import (
	"syscall/js"

	"github.com/QuestScreen/api/web"
	askew "github.com/flyx/askew/runtime"
)

var αcalendarUITemplate = js.Global().Get("document").Call("createElement", "template")

func init() {
	αcalendarUITemplate.Set("innerHTML", `
  <table class="qstut-calendar">
    <tbody>
      <!--embed(days)-->
      <!--embed(months)-->
      <!--embed(years)-->
    </tbody>
  </table>
`)
}

// calendarUI is a DOM component autogenerated by Askew
type calendarUI struct {
	αcd    askew.ComponentData
	srv    web.Server
	years  intRow
	months stringRow
	days   intRow
}

// FirstNode returns the first DOM node of this component.
// It implements the askew.Component interface.
func (o *calendarUI) FirstNode() js.Value {
	return o.αcd.First()
}

// askewInit initializes the component, discarding all previous information.
// The component is initially a DocumentFragment until it gets inserted into
// the main document. It can be manipulated both before and after insertion.
func (o *calendarUI) askewInit(srv web.Server) {
	o.αcd.Init(αcalendarUITemplate.Get("content").Call("cloneNode", true))
	o.srv = srv

	{
		container := o.αcd.Walk(1, 1)
		o.years.Init(YearRow)
		o.years.InsertInto(container, container.Get("childNodes").Index(5))
		o.years.Controller = o
	}
	{
		container := o.αcd.Walk(1, 1)
		o.months.Init(MonthRow)
		o.months.InsertInto(container, container.Get("childNodes").Index(3))
		o.months.Controller = o
	}
	{
		container := o.αcd.Walk(1, 1)
		o.days.Init(DayRow)
		o.days.InsertInto(container, container.Get("childNodes").Index(1))
		o.days.Controller = o
	}
}

// InsertInto inserts this component into the given object.
// The component will be in inserted state afterwards.
//
// The component will be inserted in front of 'before', or at the end if 'before' is 'js.Undefined()'.
func (o *calendarUI) InsertInto(parent js.Value, before js.Value) {
	o.αcd.DoInsert(parent, before)
}

// Extract removes this component from its current parent.
// The component will be in initial state afterwards.
func (o *calendarUI) Extract() {
	o.αcd.DoExtract()
}

// Destroy destroys this element (and all contained components). If it is
// currently inserted anywhere, it gets removed before.
func (o *calendarUI) Destroy() {
	o.years.Destroy()
	o.months.Destroy()
	o.days.Destroy()
	o.αcd.DoDestroy()
}

// intRowController can be implemented to handle external events
// generated by intRow
type intRowController interface {
	step(kind RowKind, amount int)
}

var αintRowTemplate = js.Global().Get("document").Call("createElement", "template")

func init() {
	αintRowTemplate.Set("innerHTML", `
  <!--controller-->
  
  <tr><td colspan="7"></td></tr>
  <tr>
    <td>
      <button><i class="fas fa-angle-double-left"></i> 10</button>
    </td>
    <td>
      <button><i class="fas fa-angle-double-left"></i> 3</button>
    </td>
    <td>
      <button><i class="fas fa-angle-left"></i> 1</button>
    </td>
    <td></td>
    <td>
      <button>1 <i class="fas fa-angle-right"></i></button>
    </td>
    <td>
      <button>3 <i class="fas fa-angle-double-right"></i></button>
    </td>
    <td>
      <button>10 <i class="fas fa-angle-double-right"></i></button>
    </td>
  </tr>

`)
}

// intRow is a DOM component autogenerated by Askew
type intRow struct {
	αcd askew.ComponentData
	// Controller is the adapter for events generated from this component.
	// if nil, events that would be passed to the controller will not be handled.
	Controller intRowController
	value      askew.IntValue
	kind       RowKind
}

// newIntRow creates a new component and initializes it with the given parameters.
func newIntRow(kind RowKind) *intRow {
	ret := new(intRow)
	ret.askewInit(kind)
	return ret
}

// Init initializes the component with the given arguments.
func (o *intRow) Init(kind RowKind) {
	o.askewInit(kind)
}

// FirstNode returns the first DOM node of this component.
// It implements the askew.Component interface.
func (o *intRow) FirstNode() js.Value {
	return o.αcd.First()
}

// askewInit initializes the component, discarding all previous information.
// The component is initially a DocumentFragment until it gets inserted into
// the main document. It can be manipulated both before and after insertion.
func (o *intRow) askewInit(kind RowKind) {
	o.αcd.Init(αintRowTemplate.Get("content").Call("cloneNode", true))
	o.kind = kind

	o.value.BoundValue = askew.NewBoundProperty(&o.αcd, "textContent", 5, 7)
	{
		block := o.αcd.Walk()
		{
			tmp := askew.BoundPropertyAt(
				askew.WalkPath(block, 3, 0), "textContent")
			askew.Assign(tmp, kind.String())
		}
	}
	{
		src := o.αcd.Walk(5, 1, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, -10)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 3, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, -3)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 5, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, -1)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 9, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, 1)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 11, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, 3)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 13, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, 10)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
}

// InsertInto inserts this component into the given object.
// The component will be in inserted state afterwards.
//
// The component will be inserted in front of 'before', or at the end if 'before' is 'js.Undefined()'.
func (o *intRow) InsertInto(parent js.Value, before js.Value) {
	o.αcd.DoInsert(parent, before)
}

// Extract removes this component from its current parent.
// The component will be in initial state afterwards.
func (o *intRow) Extract() {
	o.αcd.DoExtract()
}

// Destroy destroys this element (and all contained components). If it is
// currently inserted anywhere, it gets removed before.
func (o *intRow) Destroy() {
	o.αcd.DoDestroy()
}

// stringRowController can be implemented to handle external events
// generated by stringRow
type stringRowController interface {
	step(kind RowKind, amount int)
}

var αstringRowTemplate = js.Global().Get("document").Call("createElement", "template")

func init() {
	αstringRowTemplate.Set("innerHTML", `
  <!--controller-->
  
  <tr><td colspan="7"></td></tr>
  <tr>
    <td>
      <button><i class="fas fa-angle-double-left"></i> 10</button>
    </td>
    <td>
      <button><i class="fas fa-angle-double-left"></i> 3</button>
    </td>
    <td>
      <button><i class="fas fa-angle-left"></i> 1</button>
    </td>
    <td></td>
    <td>
      <button>1 <i class="fas fa-angle-right"></i></button>
    </td>
    <td>
      <button>3 <i class="fas fa-angle-double-right"></i></button>
    </td>
    <td>
      <button>10 <i class="fas fa-angle-double-right"></i></button>
    </td>
  </tr>

`)
}

// stringRow is a DOM component autogenerated by Askew
type stringRow struct {
	αcd askew.ComponentData
	// Controller is the adapter for events generated from this component.
	// if nil, events that would be passed to the controller will not be handled.
	Controller stringRowController
	value      askew.StringValue
	kind       RowKind
}

// newStringRow creates a new component and initializes it with the given parameters.
func newStringRow(kind RowKind) *stringRow {
	ret := new(stringRow)
	ret.askewInit(kind)
	return ret
}

// Init initializes the component with the given arguments.
func (o *stringRow) Init(kind RowKind) {
	o.askewInit(kind)
}

// FirstNode returns the first DOM node of this component.
// It implements the askew.Component interface.
func (o *stringRow) FirstNode() js.Value {
	return o.αcd.First()
}

// askewInit initializes the component, discarding all previous information.
// The component is initially a DocumentFragment until it gets inserted into
// the main document. It can be manipulated both before and after insertion.
func (o *stringRow) askewInit(kind RowKind) {
	o.αcd.Init(αstringRowTemplate.Get("content").Call("cloneNode", true))
	o.kind = kind

	o.value.BoundValue = askew.NewBoundProperty(&o.αcd, "textContent", 5, 7)
	{
		block := o.αcd.Walk()
		{
			tmp := askew.BoundPropertyAt(
				askew.WalkPath(block, 3, 0), "textContent")
			askew.Assign(tmp, kind.String())
		}
	}
	{
		src := o.αcd.Walk(5, 1, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, -10)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 3, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, -3)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 5, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, -1)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 9, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, 1)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 11, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, 3)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 13, 1)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {

				go o.Controller.step(o.kind, 10)
				arguments[0].Call("preventDefault")
				return nil
			})
			src.Call("addEventListener", "click", wrapper)
		}
	}
}

// InsertInto inserts this component into the given object.
// The component will be in inserted state afterwards.
//
// The component will be inserted in front of 'before', or at the end if 'before' is 'js.Undefined()'.
func (o *stringRow) InsertInto(parent js.Value, before js.Value) {
	o.αcd.DoInsert(parent, before)
}

// Extract removes this component from its current parent.
// The component will be in initial state afterwards.
func (o *stringRow) Extract() {
	o.αcd.DoExtract()
}

// Destroy destroys this element (and all contained components). If it is
// currently inserted anywhere, it gets removed before.
func (o *stringRow) Destroy() {
	o.αcd.DoDestroy()
}

// calendarUIList is a list of calendarUI whose manipulation methods auto-update
// the corresponding nodes in the document.
type calendarUIList struct {
	αmgr   askew.ListManager
	αitems []*calendarUI
}

// Init initializes the list, discarding previous data.
// The list's items will be placed in the given container, starting at the
// given index.
func (l *calendarUIList) Init(container js.Value, index int) {
	l.αmgr = askew.CreateListManager(container, index)
	l.αitems = nil
}

// Len returns the number of items in the list.
func (l *calendarUIList) Len() int {
	return len(l.αitems)
}

// Item returns the item at the current index.
func (l *calendarUIList) Item(index int) *calendarUI {
	return l.αitems[index]
}

// Append appends the given item to the list.
func (l *calendarUIList) Append(item *calendarUI) {
	if item == nil {
		panic("cannot append nil to list")
	}
	l.αmgr.Append(item)
	l.αitems = append(l.αitems, item)
	return
}

// Insert inserts the given item at the given index into the list.
func (l *calendarUIList) Insert(index int, item *calendarUI) {
	var prev js.Value
	if index < len(l.αitems) {
		prev = l.αitems[index].αcd.First()
	}
	if item == nil {
		panic("cannot insert nil into list")
	}
	l.αmgr.Insert(item, prev)
	l.αitems = append(l.αitems, nil)
	copy(l.αitems[index+1:], l.αitems[index:])
	l.αitems[index] = item
	return
}

// Remove removes the item at the given index from the list and returns it.
func (l *calendarUIList) Remove(index int) *calendarUI {
	item := l.αitems[index]
	item.Extract()
	copy(l.αitems[index:], l.αitems[index+1:])
	l.αitems = l.αitems[:len(l.αitems)-1]
	return item
}

// Destroy destroys the item at the given index and removes it from the list.
func (l *calendarUIList) Destroy(index int) {
	item := l.αitems[index]
	item.Destroy()
	copy(l.αitems[index:], l.αitems[index+1:])
	l.αitems = l.αitems[:len(l.αitems)-1]
}

// DestroyAll destroys all items in the list and empties it.
func (l *calendarUIList) DestroyAll() {
	for _, item := range l.αitems {
		item.Destroy()
	}
	l.αitems = l.αitems[:0]
}

// intRowList is a list of intRow whose manipulation methods auto-update
// the corresponding nodes in the document.
type intRowList struct {
	αmgr              askew.ListManager
	αitems            []*intRow
	DefaultController intRowController
}

// Init initializes the list, discarding previous data.
// The list's items will be placed in the given container, starting at the
// given index.
func (l *intRowList) Init(container js.Value, index int) {
	l.αmgr = askew.CreateListManager(container, index)
	l.αitems = nil
}

// Len returns the number of items in the list.
func (l *intRowList) Len() int {
	return len(l.αitems)
}

// Item returns the item at the current index.
func (l *intRowList) Item(index int) *intRow {
	return l.αitems[index]
}

// Append appends the given item to the list.
func (l *intRowList) Append(item *intRow) {
	if item == nil {
		panic("cannot append nil to list")
	}
	l.αmgr.Append(item)
	l.αitems = append(l.αitems, item)
	item.Controller = l.DefaultController
	return
}

// Insert inserts the given item at the given index into the list.
func (l *intRowList) Insert(index int, item *intRow) {
	var prev js.Value
	if index < len(l.αitems) {
		prev = l.αitems[index].αcd.First()
	}
	if item == nil {
		panic("cannot insert nil into list")
	}
	l.αmgr.Insert(item, prev)
	l.αitems = append(l.αitems, nil)
	copy(l.αitems[index+1:], l.αitems[index:])
	l.αitems[index] = item
	item.Controller = l.DefaultController
	return
}

// Remove removes the item at the given index from the list and returns it.
func (l *intRowList) Remove(index int) *intRow {
	item := l.αitems[index]
	item.Extract()
	copy(l.αitems[index:], l.αitems[index+1:])
	l.αitems = l.αitems[:len(l.αitems)-1]
	return item
}

// Destroy destroys the item at the given index and removes it from the list.
func (l *intRowList) Destroy(index int) {
	item := l.αitems[index]
	item.Destroy()
	copy(l.αitems[index:], l.αitems[index+1:])
	l.αitems = l.αitems[:len(l.αitems)-1]
}

// DestroyAll destroys all items in the list and empties it.
func (l *intRowList) DestroyAll() {
	for _, item := range l.αitems {
		item.Destroy()
	}
	l.αitems = l.αitems[:0]
}

// stringRowList is a list of stringRow whose manipulation methods auto-update
// the corresponding nodes in the document.
type stringRowList struct {
	αmgr              askew.ListManager
	αitems            []*stringRow
	DefaultController stringRowController
}

// Init initializes the list, discarding previous data.
// The list's items will be placed in the given container, starting at the
// given index.
func (l *stringRowList) Init(container js.Value, index int) {
	l.αmgr = askew.CreateListManager(container, index)
	l.αitems = nil
}

// Len returns the number of items in the list.
func (l *stringRowList) Len() int {
	return len(l.αitems)
}

// Item returns the item at the current index.
func (l *stringRowList) Item(index int) *stringRow {
	return l.αitems[index]
}

// Append appends the given item to the list.
func (l *stringRowList) Append(item *stringRow) {
	if item == nil {
		panic("cannot append nil to list")
	}
	l.αmgr.Append(item)
	l.αitems = append(l.αitems, item)
	item.Controller = l.DefaultController
	return
}

// Insert inserts the given item at the given index into the list.
func (l *stringRowList) Insert(index int, item *stringRow) {
	var prev js.Value
	if index < len(l.αitems) {
		prev = l.αitems[index].αcd.First()
	}
	if item == nil {
		panic("cannot insert nil into list")
	}
	l.αmgr.Insert(item, prev)
	l.αitems = append(l.αitems, nil)
	copy(l.αitems[index+1:], l.αitems[index:])
	l.αitems[index] = item
	item.Controller = l.DefaultController
	return
}

// Remove removes the item at the given index from the list and returns it.
func (l *stringRowList) Remove(index int) *stringRow {
	item := l.αitems[index]
	item.Extract()
	copy(l.αitems[index:], l.αitems[index+1:])
	l.αitems = l.αitems[:len(l.αitems)-1]
	return item
}

// Destroy destroys the item at the given index and removes it from the list.
func (l *stringRowList) Destroy(index int) {
	item := l.αitems[index]
	item.Destroy()
	copy(l.αitems[index:], l.αitems[index+1:])
	l.αitems = l.αitems[:len(l.αitems)-1]
}

// DestroyAll destroys all items in the list and empties it.
func (l *stringRowList) DestroyAll() {
	for _, item := range l.αitems {
		item.Destroy()
	}
	l.αitems = l.αitems[:0]
}

// OptionalcalendarUI is a nillable embeddable container for calendarUI.
type OptionalcalendarUI struct {
	αcur *calendarUI
	αmgr askew.ListManager
}

// Init initializes the container to be empty.
// The contained item, if any, will be placed in the given container at the
// given index.
func (o *OptionalcalendarUI) Init(container js.Value, index int) {
	o.αmgr = askew.CreateListManager(container, index)
	o.αcur = nil
}

// Item returns the current item, or nil if no item is assigned
func (o *OptionalcalendarUI) Item() *calendarUI {
	return o.αcur
}

// Set sets the contained item destroying the current one.
// Give nil as value to simply destroy the current item.
func (o *OptionalcalendarUI) Set(value *calendarUI) {
	if o.αcur != nil {
		o.αcur.Destroy()
	}
	o.αcur = value
	if value != nil {
		o.αmgr.Append(value)
	}
}

// Remove removes the current item and returns it.
// Returns nil if there is no current item.
func (o *OptionalcalendarUI) Remove() askew.Component {
	if o.αcur != nil {
		ret := o.αcur
		ret.Extract()
		o.αcur = nil
		return ret
	}
	return nil
}

// OptionalintRow is a nillable embeddable container for intRow.
type OptionalintRow struct {
	αcur              *intRow
	αmgr              askew.ListManager
	DefaultController intRowController
}

// Init initializes the container to be empty.
// The contained item, if any, will be placed in the given container at the
// given index.
func (o *OptionalintRow) Init(container js.Value, index int) {
	o.αmgr = askew.CreateListManager(container, index)
	o.αcur = nil
}

// Item returns the current item, or nil if no item is assigned
func (o *OptionalintRow) Item() *intRow {
	return o.αcur
}

// Set sets the contained item destroying the current one.
// Give nil as value to simply destroy the current item.
func (o *OptionalintRow) Set(value *intRow) {
	if o.αcur != nil {
		o.αcur.Destroy()
	}
	o.αcur = value
	if value != nil {
		o.αmgr.Append(value)
		value.Controller = o.DefaultController
	}
}

// Remove removes the current item and returns it.
// Returns nil if there is no current item.
func (o *OptionalintRow) Remove() askew.Component {
	if o.αcur != nil {
		ret := o.αcur
		ret.Extract()
		o.αcur = nil
		return ret
	}
	return nil
}

// OptionalstringRow is a nillable embeddable container for stringRow.
type OptionalstringRow struct {
	αcur              *stringRow
	αmgr              askew.ListManager
	DefaultController stringRowController
}

// Init initializes the container to be empty.
// The contained item, if any, will be placed in the given container at the
// given index.
func (o *OptionalstringRow) Init(container js.Value, index int) {
	o.αmgr = askew.CreateListManager(container, index)
	o.αcur = nil
}

// Item returns the current item, or nil if no item is assigned
func (o *OptionalstringRow) Item() *stringRow {
	return o.αcur
}

// Set sets the contained item destroying the current one.
// Give nil as value to simply destroy the current item.
func (o *OptionalstringRow) Set(value *stringRow) {
	if o.αcur != nil {
		o.αcur.Destroy()
	}
	o.αcur = value
	if value != nil {
		o.αmgr.Append(value)
		value.Controller = o.DefaultController
	}
}

// Remove removes the current item and returns it.
// Returns nil if there is no current item.
func (o *OptionalstringRow) Remove() askew.Component {
	if o.αcur != nil {
		ret := o.αcur
		ret.Extract()
		o.αcur = nil
		return ret
	}
	return nil
}
