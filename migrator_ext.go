package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (m Migrator) DropConstraint(value interface{}, name string) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		constraint, _ := m.GuessConstraintInterfaceAndTable(stmt, name)
		if constraint != nil {
			name = constraint.GetName()
		}
		return m.DB.Exec("ALTER TABLE ? DROP CONSTRAINT ?", m.CurrentTable(stmt), clause.Column{Name: name}).Error
	})
}

func GetTableName(schema, tableName string) string {
	return fmt.Sprintf("%s.%s", schema, tableName)
}
