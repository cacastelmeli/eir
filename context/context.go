package context

import "github.com/cacastelmeli/eir/util"

// Context used to be exposed to template-based code or files
type Context struct {
	// Type's name (e.g. `user`)
	TypeName Caseable

	// Use-case name (aka application service)
	UseCaseName Caseable

	// Repository implementation name
	RepositoryName Caseable

	// Domain's entity name
	EntityName Caseable
}

// ModulePath gets the module's path from a parsed modfile
func (ctx *Context) ModulePath() string {
	modFile, err := util.ParseModfile()

	if err != nil {
		panic(err)
	}

	return modFile.Module.Mod.Path
}

// FileNames converts every context field to
// its snake-cased version
func (ctx *Context) FileNames() *Context {
	return &Context{
		TypeName:       ctx.TypeName.Snake(),
		UseCaseName:    ctx.UseCaseName.Snake(),
		RepositoryName: ctx.RepositoryName.Snake(),
		EntityName:     ctx.EntityName.Snake(),
	}
}
