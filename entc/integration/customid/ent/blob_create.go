// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/blob"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// BlobCreate is the builder for creating a Blob entity.
type BlobCreate struct {
	config
	mutation *BlobMutation
	hooks    []Hook
}

// SetUUID sets the uuid field.
func (bc *BlobCreate) SetUUID(u uuid.UUID) *BlobCreate {
	bc.mutation.SetUUID(u)
	return bc
}

// SetID sets the id field.
func (bc *BlobCreate) SetID(u uuid.UUID) *BlobCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetParentID sets the parent edge to Blob by id.
func (bc *BlobCreate) SetParentID(id uuid.UUID) *BlobCreate {
	bc.mutation.SetParentID(id)
	return bc
}

// SetNillableParentID sets the parent edge to Blob by id if the given value is not nil.
func (bc *BlobCreate) SetNillableParentID(id *uuid.UUID) *BlobCreate {
	if id != nil {
		bc = bc.SetParentID(*id)
	}
	return bc
}

// SetParent sets the parent edge to Blob.
func (bc *BlobCreate) SetParent(b *Blob) *BlobCreate {
	return bc.SetParentID(b.ID)
}

// AddLinkIDs adds the links edge to Blob by ids.
func (bc *BlobCreate) AddLinkIDs(ids ...uuid.UUID) *BlobCreate {
	bc.mutation.AddLinkIDs(ids...)
	return bc
}

// AddLinks adds the links edges to Blob.
func (bc *BlobCreate) AddLinks(b ...*Blob) *BlobCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddLinkIDs(ids...)
}

// Save creates the Blob in the database.
func (bc *BlobCreate) Save(ctx context.Context) (*Blob, error) {
	if _, ok := bc.mutation.UUID(); !ok {
		v := blob.DefaultUUID()
		bc.mutation.SetUUID(v)
	}
	var (
		err  error
		node *Blob
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlobCreate) SaveX(ctx context.Context) *Blob {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BlobCreate) sqlSave(ctx context.Context) (*Blob, error) {
	var (
		b     = &Blob{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		b.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: blob.FieldUUID,
		})
		b.UUID = value
	}
	if nodes := bc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}
