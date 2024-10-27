
grammar Gold;

options{caseInsensitive=true;}

goldProg: goldDefinittion+ EOF;

// class/module level definitions
goldDefinittion: 
    goldClassDefinition
    | goldModuleDefinition
    | goldFunctionDefinition
    | goldProcedureDefinition
    ;  

goldProcedureDefinition: 
    'procedure'|'proc' name=ID ('(' params+=parameterDefinition (',' params+=parameterDefinition)? ')')?
    statement*
    'end'|'endproc';
goldFunctionDefinition: 'function'|'func';

parameterDefinition: PARAMETER_MODIFIER? name=ID (':' varType)?;

statement:
    variableDefinition
    ;

// statements
variableDefinition: 'var' name=ID ':' varType;

goldClassDefinition: 'class' name=ID ('(' parentClass=ID ')')?;
goldModuleDefinition: 'module' name=ID;

varType:
    ID;

PARAMETER_MODIFIER: 'inout'|'var'|'const';
WHITESPACE: [ \t\r\n]+ -> skip;
INT_LITERAL: [-+]?[0-9]+;
DEICMAL_LITERAL: [-+]?[0-9]+ '.' [0-9]+;
STRING_LITERAL: '\'' .*? '\'';
ID: [_a-z][_a-z0-9]*;
