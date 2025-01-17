package snowflake

import (
	"fmt"
)

type (
	AllGrantType   string
	AllGrantTarget string
)

const (
	AllGrantTypeSchema           AllGrantType = "SCHEMA"
	AllGrantTypeTable            AllGrantType = "TABLE"
	AllGrantTypeView             AllGrantType = "VIEW"
	AllGrantTypeMaterializedView AllGrantType = "MATERIALIZED VIEW"
	AllGrantTypeStage            AllGrantType = "STAGE"
	AllGrantTypeExternalTable    AllGrantType = "EXTERNAL TABLE"
	AllGrantTypeFileFormat       AllGrantType = "FILE FORMAT"
	AllGrantTypeFunction         AllGrantType = "FUNCTION"
	AllGrantTypeProcedure        AllGrantType = "PROCEDURE"
	AllGrantTypeSequence         AllGrantType = "SEQUENCE"
	AllGrantTypeStream           AllGrantType = "STREAM"
	AllGrantTypePipe             AllGrantType = "PIPE"
	AllGrantTypeTask             AllGrantType = "TASK"
)

const (
	AllGrantTargetSchema   AllGrantTarget = "SCHEMA"
	AllGrantTargetDatabase AllGrantTarget = "DATABASE"
)

// AllGrantBuilder abstracts the creation of ExistingGrantExecutables.
type AllGrantBuilder struct {
	name           string
	qualifiedName  string
	allGrantType   AllGrantType
	allGrantTarget AllGrantTarget
}

func getNameAndQualifiedNameForAllGrants(db, schema string) (string, string, AllGrantTarget) {
	name := schema
	AllGrantTarget := AllGrantTargetSchema
	qualifiedName := fmt.Sprintf(`"%v"."%v"`, db, schema)

	if schema == "" {
		name = db
		AllGrantTarget = AllGrantTargetDatabase
		qualifiedName = fmt.Sprintf(`"%v"`, db)
	}

	return name, qualifiedName, AllGrantTarget
}

// Name returns the object name for this FutureGrantBuilder.
func (fgb *AllGrantBuilder) Name() string {
	return fgb.name
}

func (fgb *AllGrantBuilder) GrantType() string {
	return string(fgb.allGrantType)
}

// ExistingSchemaGrant returns a pointer to a AllGrantBuilder for a schema.
func ExistingSchemaGrant(db string) GrantBuilder {
	return &AllGrantBuilder{
		name:           db,
		qualifiedName:  fmt.Sprintf(`"%v"`, db),
		allGrantType:   AllGrantTypeSchema,
		allGrantTarget: AllGrantTargetDatabase,
	}
}

// AllTableGrant returns a pointer to a AllGrantBuilder for a table.
func AllTableGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeTable,
		allGrantTarget: target,
	}
}

// ExistingViewGrant returns a pointer to a AllGrantBuilder for a view.
func ExistingViewGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeView,
		allGrantTarget: target,
	}
}

// ExistingMaterializedViewGrant returns a pointer to a AllGrantBuilder for a view.
func ExistingMaterializedViewGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeMaterializedView,
		allGrantTarget: target,
	}
}

// ExistingStageGrant returns a pointer to a AllGrantBuilder for a stage.
func ExistingStageGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeStage,
		allGrantTarget: target,
	}
}

// ExistingExternalTableGrant returns a pointer to a AllGrantBuilder for a external table.
func ExistingExternalTableGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeExternalTable,
		allGrantTarget: target,
	}
}

// ExistingFileFormatGrant returns a pointer to a AllGrantBuilder for a file format.
func ExistingFileFormatGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeFileFormat,
		allGrantTarget: target,
	}
}

// ExistingFunctionGrant returns a pointer to a AllGrantBuilder for a function.
func ExistingFunctionGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeFunction,
		allGrantTarget: target,
	}
}

// ExistingProcedureGrant returns a pointer to a AllGrantBuilder for a procedure.
func ExistingProcedureGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeProcedure,
		allGrantTarget: target,
	}
}

// ExistingSequenceGrant returns a pointer to a AllGrantBuilder for a sequence.
func ExistingSequenceGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeSequence,
		allGrantTarget: target,
	}
}

// ExistingStreamGrant returns a pointer to a AllGrantBuilder for a stream.
func ExistingStreamGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeStream,
		allGrantTarget: target,
	}
}

// ExistingPipeGrant returns a pointer to a AllGrantBuilder for a pipe.
func ExistingPipeGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypePipe,
		allGrantTarget: target,
	}
}

// ExistingTaskGrant returns a pointer to a AllGrantBuilder for a task.
func ExistingTaskGrant(db, schema string) GrantBuilder {
	name, qualifiedName, target := getNameAndQualifiedNameForAllGrants(db, schema)
	return &AllGrantBuilder{
		name:           name,
		qualifiedName:  qualifiedName,
		allGrantType:   AllGrantTypeTask,
		allGrantTarget: target,
	}
}

// Show returns the SQL that will show all privileges on the grant.
func (fgb *AllGrantBuilder) Show() string {
	return fmt.Sprintf(`SHOW ALL GRANTS IN %v %v`, fgb.allGrantTarget, fgb.qualifiedName)
}

// Role returns a pointer to a ExistingGrantExecutable for a role.
func (fgb *AllGrantBuilder) Role(n string) GrantExecutable {
	return &ExistingGrantExecutable{
		granteeName:    n,
		grantName:      fgb.qualifiedName,
		allGrantType:   fgb.allGrantType,
		allGrantTarget: fgb.allGrantTarget,
	}
}

// Share is not implemented because all objects cannot be granted to shares.
func (fgb *AllGrantBuilder) Share(n string) GrantExecutable {
	return nil
}

// ExistingGrantExecutable abstracts the creation of SQL queries to build all grants for
// different all grant types.
type ExistingGrantExecutable struct {
	grantName      string
	granteeName    string
	allGrantType   AllGrantType
	allGrantTarget AllGrantTarget
}

// Grant returns the SQL that will grant all privileges on the grant to the grantee.
func (fge *ExistingGrantExecutable) Grant(p string, w bool) string {
	var template string
	if w {
		template = `GRANT %v ON ALL %vS IN %v %v TO ROLE "%v" WITH GRANT OPTION`
	} else {
		template = `GRANT %v ON ALL %vS IN %v %v TO ROLE "%v"`
	}
	return fmt.Sprintf(template,
		p, fge.allGrantType, fge.allGrantTarget, fge.grantName, fge.granteeName)
}

// Revoke returns the SQL that will revoke all privileges on the grant from the grantee.
func (fge *ExistingGrantExecutable) Revoke(p string) []string {
	// TODO: has no effect for ALL GRANTS
	return []string{
		fmt.Sprintf(`REVOKE %v ON ALL %vS IN %v %v FROM ROLE "%v"`,
			p, fge.allGrantType, fge.allGrantTarget, fge.grantName, fge.granteeName),
	}
}

// Show returns the SQL that will show all all grants on the schema.
func (fge *ExistingGrantExecutable) Show() string {
	return fmt.Sprintf(`SHOW ALL GRANTS IN %v %v`, fge.allGrantTarget, fge.grantName)
}
