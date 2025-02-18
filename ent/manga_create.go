// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/browningluke/mangathr/v2/ent/chapter"
	"github.com/browningluke/mangathr/v2/ent/manga"
)

// MangaCreate is the builder for creating a Manga entity.
type MangaCreate struct {
	config
	mutation *MangaMutation
	hooks    []Hook
}

// SetMangaID sets the "MangaID" field.
func (mc *MangaCreate) SetMangaID(s string) *MangaCreate {
	mc.mutation.SetMangaID(s)
	return mc
}

// SetSource sets the "Source" field.
func (mc *MangaCreate) SetSource(s string) *MangaCreate {
	mc.mutation.SetSource(s)
	return mc
}

// SetTitle sets the "Title" field.
func (mc *MangaCreate) SetTitle(s string) *MangaCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetMapping sets the "Mapping" field.
func (mc *MangaCreate) SetMapping(s string) *MangaCreate {
	mc.mutation.SetMapping(s)
	return mc
}

// SetRegisteredOn sets the "RegisteredOn" field.
func (mc *MangaCreate) SetRegisteredOn(t time.Time) *MangaCreate {
	mc.mutation.SetRegisteredOn(t)
	return mc
}

// SetFilteredGroups sets the "FilteredGroups" field.
func (mc *MangaCreate) SetFilteredGroups(s []string) *MangaCreate {
	mc.mutation.SetFilteredGroups(s)
	return mc
}

// SetExcludedGroups sets the "ExcludedGroups" field.
func (mc *MangaCreate) SetExcludedGroups(s []string) *MangaCreate {
	mc.mutation.SetExcludedGroups(s)
	return mc
}

// AddChapterIDs adds the "Chapters" edge to the Chapter entity by IDs.
func (mc *MangaCreate) AddChapterIDs(ids ...int) *MangaCreate {
	mc.mutation.AddChapterIDs(ids...)
	return mc
}

// AddChapters adds the "Chapters" edges to the Chapter entity.
func (mc *MangaCreate) AddChapters(c ...*Chapter) *MangaCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mc.AddChapterIDs(ids...)
}

// Mutation returns the MangaMutation object of the builder.
func (mc *MangaCreate) Mutation() *MangaMutation {
	return mc.mutation
}

// Save creates the Manga in the database.
func (mc *MangaCreate) Save(ctx context.Context) (*Manga, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MangaCreate) SaveX(ctx context.Context) *Manga {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MangaCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MangaCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MangaCreate) defaults() {
	if _, ok := mc.mutation.FilteredGroups(); !ok {
		v := manga.DefaultFilteredGroups
		mc.mutation.SetFilteredGroups(v)
	}
	if _, ok := mc.mutation.ExcludedGroups(); !ok {
		v := manga.DefaultExcludedGroups
		mc.mutation.SetExcludedGroups(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MangaCreate) check() error {
	if _, ok := mc.mutation.MangaID(); !ok {
		return &ValidationError{Name: "MangaID", err: errors.New(`ent: missing required field "Manga.MangaID"`)}
	}
	if _, ok := mc.mutation.Source(); !ok {
		return &ValidationError{Name: "Source", err: errors.New(`ent: missing required field "Manga.Source"`)}
	}
	if _, ok := mc.mutation.Title(); !ok {
		return &ValidationError{Name: "Title", err: errors.New(`ent: missing required field "Manga.Title"`)}
	}
	if _, ok := mc.mutation.Mapping(); !ok {
		return &ValidationError{Name: "Mapping", err: errors.New(`ent: missing required field "Manga.Mapping"`)}
	}
	if _, ok := mc.mutation.RegisteredOn(); !ok {
		return &ValidationError{Name: "RegisteredOn", err: errors.New(`ent: missing required field "Manga.RegisteredOn"`)}
	}
	return nil
}

func (mc *MangaCreate) sqlSave(ctx context.Context) (*Manga, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MangaCreate) createSpec() (*Manga, *sqlgraph.CreateSpec) {
	var (
		_node = &Manga{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(manga.Table, sqlgraph.NewFieldSpec(manga.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.MangaID(); ok {
		_spec.SetField(manga.FieldMangaID, field.TypeString, value)
		_node.MangaID = value
	}
	if value, ok := mc.mutation.Source(); ok {
		_spec.SetField(manga.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if value, ok := mc.mutation.Title(); ok {
		_spec.SetField(manga.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := mc.mutation.Mapping(); ok {
		_spec.SetField(manga.FieldMapping, field.TypeString, value)
		_node.Mapping = value
	}
	if value, ok := mc.mutation.RegisteredOn(); ok {
		_spec.SetField(manga.FieldRegisteredOn, field.TypeTime, value)
		_node.RegisteredOn = value
	}
	if value, ok := mc.mutation.FilteredGroups(); ok {
		_spec.SetField(manga.FieldFilteredGroups, field.TypeJSON, value)
		_node.FilteredGroups = value
	}
	if value, ok := mc.mutation.ExcludedGroups(); ok {
		_spec.SetField(manga.FieldExcludedGroups, field.TypeJSON, value)
		_node.ExcludedGroups = value
	}
	if nodes := mc.mutation.ChaptersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manga.ChaptersTable,
			Columns: []string{manga.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MangaCreateBulk is the builder for creating many Manga entities in bulk.
type MangaCreateBulk struct {
	config
	err      error
	builders []*MangaCreate
}

// Save creates the Manga entities in the database.
func (mcb *MangaCreateBulk) Save(ctx context.Context) ([]*Manga, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Manga, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MangaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MangaCreateBulk) SaveX(ctx context.Context) []*Manga {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MangaCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MangaCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
