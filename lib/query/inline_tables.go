package query

import (
	"context"
	"strings"

	"github.com/mithrandie/csvq/lib/parser"
)

type InlineTableMap map[string]*View

func (it InlineTableMap) Set(ctx context.Context, scope *ReferenceScope, inlineTable parser.InlineTable) error {
	if it.Exists(inlineTable.Name) {
		return NewInLineTableRedefinedError(inlineTable.Name)
	}

	scope = scope.CreateNode()
	if inlineTable.IsRecursive() {
		if scope.RecursiveTable != nil {
			return NewNestedRecursionError(inlineTable.Name)
		}

		scope.RecursiveTable = &inlineTable
	}

	view, err := Select(ctx, scope, inlineTable.Query)
	scope.CloseCurrentNode()
	if err != nil {
		return err
	}

	err = view.Header.Update(inlineTable.Name.Literal, inlineTable.Fields)
	if err != nil {
		if _, ok := err.(*FieldLengthNotMatchError); ok {
			return NewInlineTableFieldLengthError(inlineTable.Query, inlineTable.Name, len(inlineTable.Fields))
		}
		return err
	}

	view.FileInfo = nil
	it[strings.ToUpper(inlineTable.Name.Literal)] = view

	return nil
}

func (it InlineTableMap) Exists(name parser.Identifier) bool {
	uname := strings.ToUpper(name.Literal)
	_, ok := it[uname]
	return ok
}

func (it InlineTableMap) Get(name parser.Identifier) (*View, error) {
	uname := strings.ToUpper(name.Literal)
	if view, ok := it[uname]; ok {
		return view.Copy(), nil
	}
	return nil, NewUndefinedInLineTableError(name)
}

func (it InlineTableMap) Clear() {
	for k := range it {
		delete(it, k)
	}
}
