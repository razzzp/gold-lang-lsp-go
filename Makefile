parser:
	rm -r -f gold-lang-lsp/parser/antlr4
	java -jar antlr4/antlr-4.13.2-complete.jar -Dlanguage=Go -Xexact-output-dir -o gold-lang-lsp/parser -package parser antlr4/Gold.g4