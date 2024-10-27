// Code generated from antlr4/Gold.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gold
import "github.com/antlr4-go/antlr/v4"

// BaseGoldListener is a complete listener for a parse tree produced by GoldParser.
type BaseGoldListener struct{}

var _ GoldListener = &BaseGoldListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoldListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoldListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoldListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoldListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGoldProg is called when production goldProg is entered.
func (s *BaseGoldListener) EnterGoldProg(ctx *GoldProgContext) {}

// ExitGoldProg is called when production goldProg is exited.
func (s *BaseGoldListener) ExitGoldProg(ctx *GoldProgContext) {}

// EnterGoldDefinition is called when production goldDefinition is entered.
func (s *BaseGoldListener) EnterGoldDefinition(ctx *GoldDefinitionContext) {}

// ExitGoldDefinition is called when production goldDefinition is exited.
func (s *BaseGoldListener) ExitGoldDefinition(ctx *GoldDefinitionContext) {}

// EnterGoldClassDefinition is called when production goldClassDefinition is entered.
func (s *BaseGoldListener) EnterGoldClassDefinition(ctx *GoldClassDefinitionContext) {}

// ExitGoldClassDefinition is called when production goldClassDefinition is exited.
func (s *BaseGoldListener) ExitGoldClassDefinition(ctx *GoldClassDefinitionContext) {}

// EnterGoldModuleDefinition is called when production goldModuleDefinition is entered.
func (s *BaseGoldListener) EnterGoldModuleDefinition(ctx *GoldModuleDefinitionContext) {}

// ExitGoldModuleDefinition is called when production goldModuleDefinition is exited.
func (s *BaseGoldListener) ExitGoldModuleDefinition(ctx *GoldModuleDefinitionContext) {}

// EnterGoldProcedureDefinition is called when production goldProcedureDefinition is entered.
func (s *BaseGoldListener) EnterGoldProcedureDefinition(ctx *GoldProcedureDefinitionContext) {}

// ExitGoldProcedureDefinition is called when production goldProcedureDefinition is exited.
func (s *BaseGoldListener) ExitGoldProcedureDefinition(ctx *GoldProcedureDefinitionContext) {}

// EnterGoldFunctionDefinition is called when production goldFunctionDefinition is entered.
func (s *BaseGoldListener) EnterGoldFunctionDefinition(ctx *GoldFunctionDefinitionContext) {}

// ExitGoldFunctionDefinition is called when production goldFunctionDefinition is exited.
func (s *BaseGoldListener) ExitGoldFunctionDefinition(ctx *GoldFunctionDefinitionContext) {}

// EnterParameterDefinition is called when production parameterDefinition is entered.
func (s *BaseGoldListener) EnterParameterDefinition(ctx *ParameterDefinitionContext) {}

// ExitParameterDefinition is called when production parameterDefinition is exited.
func (s *BaseGoldListener) ExitParameterDefinition(ctx *ParameterDefinitionContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGoldListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGoldListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableDefinition is called when production variableDefinition is entered.
func (s *BaseGoldListener) EnterVariableDefinition(ctx *VariableDefinitionContext) {}

// ExitVariableDefinition is called when production variableDefinition is exited.
func (s *BaseGoldListener) ExitVariableDefinition(ctx *VariableDefinitionContext) {}

// EnterVarType is called when production varType is entered.
func (s *BaseGoldListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BaseGoldListener) ExitVarType(ctx *VarTypeContext) {}
