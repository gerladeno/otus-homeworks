package sqlstorage

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gerladeno/otus_homeworks/hw12_13_14_15_calendar/internal/common"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	// run ONLY on empty DB
	// goose -dir internal/storage/sql/migrations postgres "user=calendar password=calendar dbname=postgres sslmode=disable" down
	// goose -dir internal/storage/sql/migrations postgres "user=calendar password=calendar dbname=postgres sslmode=disable" up
	t.Skip()
	log := logrus.New()
	events, err := New(log, "host=localhost port=5432 user=calendar password=calendar dbname=postgres sslmode=disable")
	require.NoError(t, err)
	tt, err := time.Parse(common.PgTimestampFmt, "2020-01-01 00:00:00")
	require.NoError(t, err)
	id, err := events.CreateEvent(context.Background(), &common.Event{
		Title:      "First",
		StartTime:  tt,
		Duration:   0,
		InviteList: "blabla",
		Comment:    "first",
	})
	require.NoError(t, err)
	require.Equal(t, id, uint64(0))

	id, err = events.CreateEvent(context.Background(), &common.Event{
		Title:      "Second",
		StartTime:  tt,
		Duration:   0,
		InviteList: "blablablabla",
		Comment:    "Second",
	})
	require.NoError(t, err)
	require.Equal(t, id, uint64(1))

	test, err := events.ListEvents(context.Background())
	require.NoError(t, err)
	require.Len(t, test, 2)

	err = events.UpdateEvent(context.Background(), 0, &common.Event{
		Title:      "First edited",
		StartTime:  tt,
		Duration:   0,
		InviteList: "blabla edited",
		Comment:    "First edited",
	})
	require.NoError(t, err)

	err = events.DeleteEvent(context.Background(), 1)
	require.NoError(t, err)

	test, err = events.ListEvents(context.Background())
	require.NoError(t, err)
	require.Len(t, test, 1)
	elem, err := events.ReadEvent(context.Background(), 1)
	require.True(t, errors.Is(err, common.ErrNoSuchEvent))
	elem, err = events.ReadEvent(context.Background(), 0)
	require.NoError(t, err)
	require.Equal(t, elem.Title, "First edited")
	require.Equal(t, elem.StartTime, tt)
	require.Equal(t, elem.InviteList, "blabla edited")
	require.Equal(t, elem.Comment, "First edited")
	require.True(t, elem.Created.Before(elem.Updated))

	id, err = events.CreateEvent(context.Background(), &common.Event{})
	require.NoError(t, err)
	require.Equal(t, id, uint64(2))
}
