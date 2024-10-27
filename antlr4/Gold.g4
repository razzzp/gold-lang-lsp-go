
grammar Gold;

options{caseInsensitive=true;}

goldProg: goldDefinition+ EOF;

// class/module level definitions
goldDefinition: 
    goldClassDefinition
    | goldModuleDefinition
    | goldFunctionDefinition
    | goldProcedureDefinition
    ;

goldClassDefinition: 'class' name=ID ('(' parentClass=ID ')')?;
goldModuleDefinition: 'module' name=ID;


goldProcedureDefinition: 
    ('procedure'|'proc') name=ID ('(' params+=parameterDefinition (',' params+=parameterDefinition)? ')')?
    statements+=statement*
    ('end'|'endproc');
goldFunctionDefinition: 'function'|'func';

parameterDefinition: PARAMETER_MODIFIER? name=ID (':' varType)?;

statement:
    variableDefinition
    ;

// statements
variableDefinition: 'var' name=ID ':' varType;

varType:
    ID;

PARAMETER_MODIFIER: 'inout'|'var'|'const';
WHITESPACE: [ \t\r\n]+ -> skip;
INT_LITERAL: [-+]?[0-9]+;
DEICMAL_LITERAL: [-+]?[0-9]+ '.' [0-9]+;
STRING_LITERAL: '\'' .*? '\'';
ID: [_a-z][_a-z0-9]*;
