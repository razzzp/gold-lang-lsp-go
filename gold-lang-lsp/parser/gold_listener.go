// Code generated from antlr4/Gold.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gold
import "github.com/antlr4-go/antlr/v4"

// GoldListener is a complete listener for a parse tree produced by GoldParser.
type GoldListener interface {
	antlr.ParseTreeListener

	// EnterGoldProg is called when entering the goldProg production.
	EnterGoldProg(c *GoldProgContext)

	// EnterGoldDefinition is called when entering the goldDefinition production.
	EnterGoldDefinition(c *GoldDefinitionContext)

	// EnterGoldClassDefinition is called when entering the goldClassDefinition production.
	EnterGoldClassDefinition(c *GoldClassDefinitionContext)

	// EnterGoldModuleDefinition is called when entering the goldModuleDefinition production.
	EnterGoldModuleDefinition(c *GoldModuleDefinitionContext)

	// EnterGoldProcedureDefinition is called when entering the goldProcedureDefinition production.
	EnterGoldProcedureDefinition(c *GoldProcedureDefinitionContext)

	// EnterGoldFunctionDefinition is called when entering the goldFunctionDefinition production.
	EnterGoldFunctionDefinition(c *GoldFunctionDefinitionContext)

	// EnterParameterDefinition is called when entering the parameterDefinition production.
	EnterParameterDefinition(c *ParameterDefinitionContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVariableDefinition is called when entering the variableDefinition production.
	EnterVariableDefinition(c *VariableDefinitionContext)

	// EnterVarType is called when entering the varType production.
	EnterVarType(c *VarTypeContext)

	// ExitGoldProg is called when exiting the goldProg production.
	ExitGoldProg(c *GoldProgContext)

	// ExitGoldDefinition is called when exiting the goldDefinition production.
	ExitGoldDefinition(c *GoldDefinitionContext)

	// ExitGoldClassDefinition is called when exiting the goldClassDefinition production.
	ExitGoldClassDefinition(c *GoldClassDefinitionContext)

	// ExitGoldModuleDefinition is called when exiting the goldModuleDefinition production.
	ExitGoldModuleDefinition(c *GoldModuleDefinitionContext)

	// ExitGoldProcedureDefinition is called when exiting the goldProcedureDefinition production.
	ExitGoldProcedureDefinition(c *GoldProcedureDefinitionContext)

	// ExitGoldFunctionDefinition is called when exiting the goldFunctionDefinition production.
	ExitGoldFunctionDefinition(c *GoldFunctionDefinitionContext)

	// ExitParameterDefinition is called when exiting the parameterDefinition production.
	ExitParameterDefinition(c *ParameterDefinitionContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVariableDefinition is called when exiting the variableDefinition production.
	ExitVariableDefinition(c *VariableDefinitionContext)

	// ExitVarType is called when exiting the varType production.
	ExitVarType(c *VarTypeContext)
}
