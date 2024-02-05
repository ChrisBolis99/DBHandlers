package DBMLParser

import (
	"reflect"
	"testing"
)

func TestParseDBML(t *testing.T) {
	tests := []struct {
		name    string
		dbml    string
		want    Schema
		wantErr bool
	}{
		{
			name: "empty DBML",
			dbml: "",
			want: Schema{},
		},
		{
			name: "single table with columns",
			dbml: `Table users {
  					id: int [pk, notNull],
  					name: varchar [unique]
		}`,
			want: Schema{
				Tables: []Table{
					{
						Name: "users",
						Columns: []Column{
							{Name: "id", Type: "int", PrimaryKey: true, NotNull: true},
							{Name: "name", Type: "varchar", Unique: true},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDBML(tt.dbml)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDBML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDBML() got = %v, want %v", got, tt.want)
			}
		})
	}
}
