package server

import (
	"testing"
)

var testTables = &Tables{
	Tables: []*Table{
		{TableNo: 1},
		{TableNo: 2},
		{TableNo: 3},
	},
}

func TestTableExists(t *testing.T) {
	if !testTables.TableExists(2) {
		t.Errorf("Error: returned false for an existing table")
	}

	if testTables.TableExists(4) {
		t.Errorf("Error: returned true for a non-existing table")
	}
}

func TestGetTable(t *testing.T) {
	table := testTables.Get(2)
	if table == nil {
		t.Errorf("Error: returned nil for an existing table")
	} else if table.TableNo != 2 {
		t.Errorf("Error: returned incorrect table")
	}

	table = testTables.Get(4)
	if table != nil {
		t.Errorf("Error: returned non-nil for a non-existing table")
	}
}

func TestAddTable(t *testing.T) {
	newTable := &Table{TableNo: 10}
	testTables.Add(newTable)
	if testTables.Get(newTable.TableNo) == nil {
		t.Errorf("Error: failed to add table")
	}
}

func TestRemoveTable(t *testing.T) {
	tableToRemove := &Table{TableNo: 2}
	testTables.RemoveTable(tableToRemove)
	if testTables.Get(tableToRemove.TableNo) != nil {
		t.Errorf("Error: failed to remove table")
	}
}
