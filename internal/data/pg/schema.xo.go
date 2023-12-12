// Package pg contains generated code for schema 'public'.
package pg

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/rarimo/worldcoin-relayer-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// Storage is the helper struct for database operations
type Storage struct {
	db *pgdb.DB
}

// New - returns new instance of storage
func New(db *pgdb.DB) *Storage {
	return &Storage{
		db,
	}
}

// DB - returns db used by Storage
func (s *Storage) DB() *pgdb.DB {
	return s.db
}

// Clone - returns new storage with clone of db
func (s *Storage) Clone() *Storage {
	return New(s.db.Clone())
}

// Transaction begins a transaction on repo.
func (s *Storage) Transaction(tx func() error) error {
	return s.db.Transaction(tx)
} // GistQ represents helper struct to access row of 'gists'.
type GistQ struct {
	db *pgdb.DB
}

// NewGistQ  - creates new instance
func NewGistQ(db *pgdb.DB) *GistQ {
	return &GistQ{
		db,
	}
}

// GistQ  - creates new instance of GistQ
func (s Storage) GistQ() *GistQ {
	return NewGistQ(s.DB())
}

var colsGist = `id, operation, confirmation`

// InsertCtx inserts a Gist to the database.
func (q GistQ) InsertCtx(ctx context.Context, g *data.Gist) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.gists (` +
		`id, operation, confirmation` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, g.ID, g.Operation, g.Confirmation)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a Gist to the database.
func (q GistQ) Insert(g *data.Gist) error {
	return q.InsertCtx(context.Background(), g)
}

// UpdateCtx updates a Gist in the database.
func (q GistQ) UpdateCtx(ctx context.Context, g *data.Gist) error {
	// update with composite primary key
	sqlstr := `UPDATE public.gists SET ` +
		`operation = $1, confirmation = $2 ` +
		`WHERE id = $3`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, g.Operation, g.Confirmation, g.ID)
	return errors.Wrap(err, "failed to execute update")
}

// Update updates a Gist in the database.
func (q GistQ) Update(g *data.Gist) error {
	return q.UpdateCtx(context.Background(), g)
}

// UpsertCtx performs an upsert for Gist.
func (q GistQ) UpsertCtx(ctx context.Context, g *data.Gist) error {
	// upsert
	sqlstr := `INSERT INTO public.gists (` +
		`id, operation, confirmation` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`operation = EXCLUDED.operation, confirmation = EXCLUDED.confirmation `
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, g.ID, g.Operation, g.Confirmation); err != nil {
		return errors.Wrap(err, "failed to execute upsert stmt")
	}
	return nil
}

// Upsert performs an upsert for Gist.
func (q GistQ) Upsert(g *data.Gist) error {
	return q.UpsertCtx(context.Background(), g)
}

// DeleteCtx deletes the Gist from the database.
func (q GistQ) DeleteCtx(ctx context.Context, g *data.Gist) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.gists ` +
		`WHERE id = $1`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, g.ID); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the Gist from the database.
func (q GistQ) Delete(g *data.Gist) error {
	return q.DeleteCtx(context.Background(), g)
} // GistTransitionQ represents helper struct to access row of 'gist_transitions'.
type GistTransitionQ struct {
	db *pgdb.DB
}

// NewGistTransitionQ  - creates new instance
func NewGistTransitionQ(db *pgdb.DB) *GistTransitionQ {
	return &GistTransitionQ{
		db,
	}
}

// GistTransitionQ  - creates new instance of GistTransitionQ
func (s Storage) GistTransitionQ() *GistTransitionQ {
	return NewGistTransitionQ(s.DB())
}

var colsGistTransition = `tx, gist, chain`

// InsertCtx inserts a GistTransition to the database.
func (q GistTransitionQ) InsertCtx(ctx context.Context, gt *data.GistTransition) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.gist_transitions (` +
		`tx, gist, chain` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, gt.Tx, gt.Gist, gt.Chain)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a GistTransition to the database.
func (q GistTransitionQ) Insert(gt *data.GistTransition) error {
	return q.InsertCtx(context.Background(), gt)
}

// UpdateCtx updates a GistTransition in the database.
func (q GistTransitionQ) UpdateCtx(ctx context.Context, gt *data.GistTransition) error {
	// update with composite primary key
	sqlstr := `UPDATE public.gist_transitions SET ` +
		`gist = $1, chain = $2 ` +
		`WHERE tx = $3`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, gt.Gist, gt.Chain, gt.Tx)
	return errors.Wrap(err, "failed to execute update")
}

// Update updates a GistTransition in the database.
func (q GistTransitionQ) Update(gt *data.GistTransition) error {
	return q.UpdateCtx(context.Background(), gt)
}

// UpsertCtx performs an upsert for GistTransition.
func (q GistTransitionQ) UpsertCtx(ctx context.Context, gt *data.GistTransition) error {
	// upsert
	sqlstr := `INSERT INTO public.gist_transitions (` +
		`tx, gist, chain` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (tx) DO ` +
		`UPDATE SET ` +
		`gist = EXCLUDED.gist, chain = EXCLUDED.chain `
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, gt.Tx, gt.Gist, gt.Chain); err != nil {
		return errors.Wrap(err, "failed to execute upsert stmt")
	}
	return nil
}

// Upsert performs an upsert for GistTransition.
func (q GistTransitionQ) Upsert(gt *data.GistTransition) error {
	return q.UpsertCtx(context.Background(), gt)
}

// DeleteCtx deletes the GistTransition from the database.
func (q GistTransitionQ) DeleteCtx(ctx context.Context, gt *data.GistTransition) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.gist_transitions ` +
		`WHERE tx = $1`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, gt.Tx); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the GistTransition from the database.
func (q GistTransitionQ) Delete(gt *data.GistTransition) error {
	return q.DeleteCtx(context.Background(), gt)
} // GorpMigrationQ represents helper struct to access row of 'gorp_migrations'.
type GorpMigrationQ struct {
	db *pgdb.DB
}

// NewGorpMigrationQ  - creates new instance
func NewGorpMigrationQ(db *pgdb.DB) *GorpMigrationQ {
	return &GorpMigrationQ{
		db,
	}
}

// GorpMigrationQ  - creates new instance of GorpMigrationQ
func (s Storage) GorpMigrationQ() *GorpMigrationQ {
	return NewGorpMigrationQ(s.DB())
}

var colsGorpMigration = `id, applied_at`

// InsertCtx inserts a GorpMigration to the database.
func (q GorpMigrationQ) InsertCtx(ctx context.Context, gm *data.GorpMigration) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.gorp_migrations (` +
		`id, applied_at` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, gm.ID, gm.AppliedAt)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a GorpMigration to the database.
func (q GorpMigrationQ) Insert(gm *data.GorpMigration) error {
	return q.InsertCtx(context.Background(), gm)
}

// UpdateCtx updates a GorpMigration in the database.
func (q GorpMigrationQ) UpdateCtx(ctx context.Context, gm *data.GorpMigration) error {
	// update with composite primary key
	sqlstr := `UPDATE public.gorp_migrations SET ` +
		`applied_at = $1 ` +
		`WHERE id = $2`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, gm.AppliedAt, gm.ID)
	return errors.Wrap(err, "failed to execute update")
}

// Update updates a GorpMigration in the database.
func (q GorpMigrationQ) Update(gm *data.GorpMigration) error {
	return q.UpdateCtx(context.Background(), gm)
}

// UpsertCtx performs an upsert for GorpMigration.
func (q GorpMigrationQ) UpsertCtx(ctx context.Context, gm *data.GorpMigration) error {
	// upsert
	sqlstr := `INSERT INTO public.gorp_migrations (` +
		`id, applied_at` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`applied_at = EXCLUDED.applied_at `
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, gm.ID, gm.AppliedAt); err != nil {
		return errors.Wrap(err, "failed to execute upsert stmt")
	}
	return nil
}

// Upsert performs an upsert for GorpMigration.
func (q GorpMigrationQ) Upsert(gm *data.GorpMigration) error {
	return q.UpsertCtx(context.Background(), gm)
}

// DeleteCtx deletes the GorpMigration from the database.
func (q GorpMigrationQ) DeleteCtx(ctx context.Context, gm *data.GorpMigration) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.gorp_migrations ` +
		`WHERE id = $1`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, gm.ID); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the GorpMigration from the database.
func (q GorpMigrationQ) Delete(gm *data.GorpMigration) error {
	return q.DeleteCtx(context.Background(), gm)
} // StateQ represents helper struct to access row of 'states'.
type StateQ struct {
	db *pgdb.DB
}

// NewStateQ  - creates new instance
func NewStateQ(db *pgdb.DB) *StateQ {
	return &StateQ{
		db,
	}
}

// StateQ  - creates new instance of StateQ
func (s Storage) StateQ() *StateQ {
	return NewStateQ(s.DB())
}

var colsState = `id, operation, confirmation`

// InsertCtx inserts a State to the database.
func (q StateQ) InsertCtx(ctx context.Context, s *data.State) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.states (` +
		`id, operation, confirmation` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, s.ID, s.Operation, s.Confirmation)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a State to the database.
func (q StateQ) Insert(s *data.State) error {
	return q.InsertCtx(context.Background(), s)
}

// UpdateCtx updates a State in the database.
func (q StateQ) UpdateCtx(ctx context.Context, s *data.State) error {
	// update with composite primary key
	sqlstr := `UPDATE public.states SET ` +
		`operation = $1, confirmation = $2 ` +
		`WHERE id = $3`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, s.Operation, s.Confirmation, s.ID)
	return errors.Wrap(err, "failed to execute update")
}

// Update updates a State in the database.
func (q StateQ) Update(s *data.State) error {
	return q.UpdateCtx(context.Background(), s)
}

// UpsertCtx performs an upsert for State.
func (q StateQ) UpsertCtx(ctx context.Context, s *data.State) error {
	// upsert
	sqlstr := `INSERT INTO public.states (` +
		`id, operation, confirmation` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`operation = EXCLUDED.operation, confirmation = EXCLUDED.confirmation `
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, s.ID, s.Operation, s.Confirmation); err != nil {
		return errors.Wrap(err, "failed to execute upsert stmt")
	}
	return nil
}

// Upsert performs an upsert for State.
func (q StateQ) Upsert(s *data.State) error {
	return q.UpsertCtx(context.Background(), s)
}

// DeleteCtx deletes the State from the database.
func (q StateQ) DeleteCtx(ctx context.Context, s *data.State) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.states ` +
		`WHERE id = $1`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, s.ID); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the State from the database.
func (q StateQ) Delete(s *data.State) error {
	return q.DeleteCtx(context.Background(), s)
} // TransitionQ represents helper struct to access row of 'transitions'.
type TransitionQ struct {
	db *pgdb.DB
}

// NewTransitionQ  - creates new instance
func NewTransitionQ(db *pgdb.DB) *TransitionQ {
	return &TransitionQ{
		db,
	}
}

// TransitionQ  - creates new instance of TransitionQ
func (s Storage) TransitionQ() *TransitionQ {
	return NewTransitionQ(s.DB())
}

var colsTransition = `tx, state, chain`

// InsertCtx inserts a Transition to the database.
func (q TransitionQ) InsertCtx(ctx context.Context, t *data.Transition) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.transitions (` +
		`tx, state, chain` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, t.Tx, t.State, t.Chain)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a Transition to the database.
func (q TransitionQ) Insert(t *data.Transition) error {
	return q.InsertCtx(context.Background(), t)
}

// UpdateCtx updates a Transition in the database.
func (q TransitionQ) UpdateCtx(ctx context.Context, t *data.Transition) error {
	// update with composite primary key
	sqlstr := `UPDATE public.transitions SET ` +
		`state = $1, chain = $2 ` +
		`WHERE tx = $3`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, t.State, t.Chain, t.Tx)
	return errors.Wrap(err, "failed to execute update")
}

// Update updates a Transition in the database.
func (q TransitionQ) Update(t *data.Transition) error {
	return q.UpdateCtx(context.Background(), t)
}

// UpsertCtx performs an upsert for Transition.
func (q TransitionQ) UpsertCtx(ctx context.Context, t *data.Transition) error {
	// upsert
	sqlstr := `INSERT INTO public.transitions (` +
		`tx, state, chain` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (tx) DO ` +
		`UPDATE SET ` +
		`state = EXCLUDED.state, chain = EXCLUDED.chain `
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, t.Tx, t.State, t.Chain); err != nil {
		return errors.Wrap(err, "failed to execute upsert stmt")
	}
	return nil
}

// Upsert performs an upsert for Transition.
func (q TransitionQ) Upsert(t *data.Transition) error {
	return q.UpsertCtx(context.Background(), t)
}

// DeleteCtx deletes the Transition from the database.
func (q TransitionQ) DeleteCtx(ctx context.Context, t *data.Transition) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.transitions ` +
		`WHERE tx = $1`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, t.Tx); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the Transition from the database.
func (q TransitionQ) Delete(t *data.Transition) error {
	return q.DeleteCtx(context.Background(), t)
}

// GistByOperationCtx retrieves a row from 'public.gists' as a Gist.
//
// Generated from index 'gists_operation_key'.
func (q GistQ) GistByOperationCtx(ctx context.Context, operation string, isForUpdate bool) (*data.Gist, error) {
	// query
	sqlstr := `SELECT ` +
		`id, operation, confirmation ` +
		`FROM public.gists ` +
		`WHERE operation = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.Gist
	err := q.db.GetRawContext(ctx, &res, sqlstr, operation)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// GistByOperation retrieves a row from 'public.gists' as a Gist.
//
// Generated from index 'gists_operation_key'.
func (q GistQ) GistByOperation(operation string, isForUpdate bool) (*data.Gist, error) {
	return q.GistByOperationCtx(context.Background(), operation, isForUpdate)
}

// GistByIDCtx retrieves a row from 'public.gists' as a Gist.
//
// Generated from index 'gists_pkey'.
func (q GistQ) GistByIDCtx(ctx context.Context, id string, isForUpdate bool) (*data.Gist, error) {
	// query
	sqlstr := `SELECT ` +
		`id, operation, confirmation ` +
		`FROM public.gists ` +
		`WHERE id = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.Gist
	err := q.db.GetRawContext(ctx, &res, sqlstr, id)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// GistByID retrieves a row from 'public.gists' as a Gist.
//
// Generated from index 'gists_pkey'.
func (q GistQ) GistByID(id string, isForUpdate bool) (*data.Gist, error) {
	return q.GistByIDCtx(context.Background(), id, isForUpdate)
}

// GistTransitionsByGistCtx retrieves a row from 'public.gist_transitions' as a GistTransition.
//
// Generated from index 'gist_transitions_index'.
func (q GistTransitionQ) GistTransitionsByGistCtx(ctx context.Context, gist string, isForUpdate bool) ([]data.GistTransition, error) {
	// query
	sqlstr := `SELECT ` +
		`tx, gist, chain ` +
		`FROM public.gist_transitions ` +
		`WHERE gist = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res []data.GistTransition
	err := q.db.SelectRawContext(ctx, &res, sqlstr, gist)
	if err != nil {
		return nil, errors.Wrap(err, "failed to exec select")
	}

	return res, nil
}

// GistTransitionsByGist retrieves a row from 'public.gist_transitions' as a GistTransition.
//
// Generated from index 'gist_transitions_index'.
func (q GistTransitionQ) GistTransitionsByGist(gist string, isForUpdate bool) ([]data.GistTransition, error) {
	return q.GistTransitionsByGistCtx(context.Background(), gist, isForUpdate)
}

// GistTransitionByTxCtx retrieves a row from 'public.gist_transitions' as a GistTransition.
//
// Generated from index 'gist_transitions_pkey'.
func (q GistTransitionQ) GistTransitionByTxCtx(ctx context.Context, tx string, isForUpdate bool) (*data.GistTransition, error) {
	// query
	sqlstr := `SELECT ` +
		`tx, gist, chain ` +
		`FROM public.gist_transitions ` +
		`WHERE tx = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.GistTransition
	err := q.db.GetRawContext(ctx, &res, sqlstr, tx)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// GistTransitionByTx retrieves a row from 'public.gist_transitions' as a GistTransition.
//
// Generated from index 'gist_transitions_pkey'.
func (q GistTransitionQ) GistTransitionByTx(tx string, isForUpdate bool) (*data.GistTransition, error) {
	return q.GistTransitionByTxCtx(context.Background(), tx, isForUpdate)
}

// GorpMigrationByIDCtx retrieves a row from 'public.gorp_migrations' as a GorpMigration.
//
// Generated from index 'gorp_migrations_pkey'.
func (q GorpMigrationQ) GorpMigrationByIDCtx(ctx context.Context, id string, isForUpdate bool) (*data.GorpMigration, error) {
	// query
	sqlstr := `SELECT ` +
		`id, applied_at ` +
		`FROM public.gorp_migrations ` +
		`WHERE id = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.GorpMigration
	err := q.db.GetRawContext(ctx, &res, sqlstr, id)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// GorpMigrationByID retrieves a row from 'public.gorp_migrations' as a GorpMigration.
//
// Generated from index 'gorp_migrations_pkey'.
func (q GorpMigrationQ) GorpMigrationByID(id string, isForUpdate bool) (*data.GorpMigration, error) {
	return q.GorpMigrationByIDCtx(context.Background(), id, isForUpdate)
}

// StateByOperationCtx retrieves a row from 'public.states' as a State.
//
// Generated from index 'states_operation_key'.
func (q StateQ) StateByOperationCtx(ctx context.Context, operation string, isForUpdate bool) (*data.State, error) {
	// query
	sqlstr := `SELECT ` +
		`id, operation, confirmation ` +
		`FROM public.states ` +
		`WHERE operation = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.State
	err := q.db.GetRawContext(ctx, &res, sqlstr, operation)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// StateByOperation retrieves a row from 'public.states' as a State.
//
// Generated from index 'states_operation_key'.
func (q StateQ) StateByOperation(operation string, isForUpdate bool) (*data.State, error) {
	return q.StateByOperationCtx(context.Background(), operation, isForUpdate)
}

// StateByIDCtx retrieves a row from 'public.states' as a State.
//
// Generated from index 'states_pkey'.
func (q StateQ) StateByIDCtx(ctx context.Context, id string, isForUpdate bool) (*data.State, error) {
	// query
	sqlstr := `SELECT ` +
		`id, operation, confirmation ` +
		`FROM public.states ` +
		`WHERE id = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.State
	err := q.db.GetRawContext(ctx, &res, sqlstr, id)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// StateByID retrieves a row from 'public.states' as a State.
//
// Generated from index 'states_pkey'.
func (q StateQ) StateByID(id string, isForUpdate bool) (*data.State, error) {
	return q.StateByIDCtx(context.Background(), id, isForUpdate)
}

// TransitionsByStateCtx retrieves a row from 'public.transitions' as a Transition.
//
// Generated from index 'transitions_index'.
func (q TransitionQ) TransitionsByStateCtx(ctx context.Context, state string, isForUpdate bool) ([]data.Transition, error) {
	// query
	sqlstr := `SELECT ` +
		`tx, state, chain ` +
		`FROM public.transitions ` +
		`WHERE state = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res []data.Transition
	err := q.db.SelectRawContext(ctx, &res, sqlstr, state)
	if err != nil {
		return nil, errors.Wrap(err, "failed to exec select")
	}

	return res, nil
}

// TransitionsByState retrieves a row from 'public.transitions' as a Transition.
//
// Generated from index 'transitions_index'.
func (q TransitionQ) TransitionsByState(state string, isForUpdate bool) ([]data.Transition, error) {
	return q.TransitionsByStateCtx(context.Background(), state, isForUpdate)
}

// TransitionByTxCtx retrieves a row from 'public.transitions' as a Transition.
//
// Generated from index 'transitions_pkey'.
func (q TransitionQ) TransitionByTxCtx(ctx context.Context, tx string, isForUpdate bool) (*data.Transition, error) {
	// query
	sqlstr := `SELECT ` +
		`tx, state, chain ` +
		`FROM public.transitions ` +
		`WHERE tx = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.Transition
	err := q.db.GetRawContext(ctx, &res, sqlstr, tx)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// TransitionByTx retrieves a row from 'public.transitions' as a Transition.
//
// Generated from index 'transitions_pkey'.
func (q TransitionQ) TransitionByTx(tx string, isForUpdate bool) (*data.Transition, error) {
	return q.TransitionByTxCtx(context.Background(), tx, isForUpdate)
}
