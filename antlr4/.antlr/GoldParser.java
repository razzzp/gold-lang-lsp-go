// Generated from /home/razzz/dev/projects/gold-lang-lsp-go/antlr4/Gold.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class GoldParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, PARAMETER_MODIFIER=14, WHITESPACE=15, 
		INT_LITERAL=16, DEICMAL_LITERAL=17, STRING_LITERAL=18, ID=19;
	public static final int
		RULE_goldProg = 0, RULE_goldDefinition = 1, RULE_goldClassDefinition = 2, 
		RULE_goldModuleDefinition = 3, RULE_goldProcedureDefinition = 4, RULE_goldFunctionDefinition = 5, 
		RULE_parameterDefinition = 6, RULE_statement = 7, RULE_variableDefinition = 8, 
		RULE_varType = 9;
	private static String[] makeRuleNames() {
		return new String[] {
			"goldProg", "goldDefinition", "goldClassDefinition", "goldModuleDefinition", 
			"goldProcedureDefinition", "goldFunctionDefinition", "parameterDefinition", 
			"statement", "variableDefinition", "varType"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'class'", "'('", "')'", "'module'", "'procedure'", "'proc'", "','", 
			"'end'", "'endproc'", "'function'", "'func'", "':'", "'var'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, "PARAMETER_MODIFIER", "WHITESPACE", "INT_LITERAL", "DEICMAL_LITERAL", 
			"STRING_LITERAL", "ID"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Gold.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GoldParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GoldProgContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(GoldParser.EOF, 0); }
		public List<GoldDefinitionContext> goldDefinition() {
			return getRuleContexts(GoldDefinitionContext.class);
		}
		public GoldDefinitionContext goldDefinition(int i) {
			return getRuleContext(GoldDefinitionContext.class,i);
		}
		public GoldProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldProg; }
	}

	public final GoldProgContext goldProg() throws RecognitionException {
		GoldProgContext _localctx = new GoldProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_goldProg);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(21); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(20);
				goldDefinition();
				}
				}
				setState(23); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 3186L) != 0) );
			setState(25);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GoldDefinitionContext extends ParserRuleContext {
		public GoldClassDefinitionContext goldClassDefinition() {
			return getRuleContext(GoldClassDefinitionContext.class,0);
		}
		public GoldModuleDefinitionContext goldModuleDefinition() {
			return getRuleContext(GoldModuleDefinitionContext.class,0);
		}
		public GoldFunctionDefinitionContext goldFunctionDefinition() {
			return getRuleContext(GoldFunctionDefinitionContext.class,0);
		}
		public GoldProcedureDefinitionContext goldProcedureDefinition() {
			return getRuleContext(GoldProcedureDefinitionContext.class,0);
		}
		public GoldDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldDefinition; }
	}

	public final GoldDefinitionContext goldDefinition() throws RecognitionException {
		GoldDefinitionContext _localctx = new GoldDefinitionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_goldDefinition);
		try {
			setState(31);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				goldClassDefinition();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(28);
				goldModuleDefinition();
				}
				break;
			case T__9:
			case T__10:
				enterOuterAlt(_localctx, 3);
				{
				setState(29);
				goldFunctionDefinition();
				}
				break;
			case T__4:
			case T__5:
				enterOuterAlt(_localctx, 4);
				{
				setState(30);
				goldProcedureDefinition();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GoldClassDefinitionContext extends ParserRuleContext {
		public Token name;
		public Token parentClass;
		public List<TerminalNode> ID() { return getTokens(GoldParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GoldParser.ID, i);
		}
		public GoldClassDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldClassDefinition; }
	}

	public final GoldClassDefinitionContext goldClassDefinition() throws RecognitionException {
		GoldClassDefinitionContext _localctx = new GoldClassDefinitionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_goldClassDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(33);
			match(T__0);
			setState(34);
			((GoldClassDefinitionContext)_localctx).name = match(ID);
			setState(38);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__1) {
				{
				setState(35);
				match(T__1);
				setState(36);
				((GoldClassDefinitionContext)_localctx).parentClass = match(ID);
				setState(37);
				match(T__2);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GoldModuleDefinitionContext extends ParserRuleContext {
		public Token name;
		public TerminalNode ID() { return getToken(GoldParser.ID, 0); }
		public GoldModuleDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldModuleDefinition; }
	}

	public final GoldModuleDefinitionContext goldModuleDefinition() throws RecognitionException {
		GoldModuleDefinitionContext _localctx = new GoldModuleDefinitionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_goldModuleDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			match(T__3);
			setState(41);
			((GoldModuleDefinitionContext)_localctx).name = match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GoldProcedureDefinitionContext extends ParserRuleContext {
		public Token name;
		public ParameterDefinitionContext parameterDefinition;
		public List<ParameterDefinitionContext> params = new ArrayList<ParameterDefinitionContext>();
		public StatementContext statement;
		public List<StatementContext> statements = new ArrayList<StatementContext>();
		public TerminalNode ID() { return getToken(GoldParser.ID, 0); }
		public List<ParameterDefinitionContext> parameterDefinition() {
			return getRuleContexts(ParameterDefinitionContext.class);
		}
		public ParameterDefinitionContext parameterDefinition(int i) {
			return getRuleContext(ParameterDefinitionContext.class,i);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public GoldProcedureDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldProcedureDefinition; }
	}

	public final GoldProcedureDefinitionContext goldProcedureDefinition() throws RecognitionException {
		GoldProcedureDefinitionContext _localctx = new GoldProcedureDefinitionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_goldProcedureDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			_la = _input.LA(1);
			if ( !(_la==T__4 || _la==T__5) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(44);
			((GoldProcedureDefinitionContext)_localctx).name = match(ID);
			setState(53);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__1) {
				{
				setState(45);
				match(T__1);
				setState(46);
				((GoldProcedureDefinitionContext)_localctx).parameterDefinition = parameterDefinition();
				((GoldProcedureDefinitionContext)_localctx).params.add(((GoldProcedureDefinitionContext)_localctx).parameterDefinition);
				setState(49);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__6) {
					{
					setState(47);
					match(T__6);
					setState(48);
					((GoldProcedureDefinitionContext)_localctx).parameterDefinition = parameterDefinition();
					((GoldProcedureDefinitionContext)_localctx).params.add(((GoldProcedureDefinitionContext)_localctx).parameterDefinition);
					}
				}

				setState(51);
				match(T__2);
				}
			}

			setState(58);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(55);
				((GoldProcedureDefinitionContext)_localctx).statement = statement();
				((GoldProcedureDefinitionContext)_localctx).statements.add(((GoldProcedureDefinitionContext)_localctx).statement);
				}
				}
				setState(60);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(61);
			_la = _input.LA(1);
			if ( !(_la==T__7 || _la==T__8) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GoldFunctionDefinitionContext extends ParserRuleContext {
		public GoldFunctionDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldFunctionDefinition; }
	}

	public final GoldFunctionDefinitionContext goldFunctionDefinition() throws RecognitionException {
		GoldFunctionDefinitionContext _localctx = new GoldFunctionDefinitionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_goldFunctionDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(63);
			_la = _input.LA(1);
			if ( !(_la==T__9 || _la==T__10) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParameterDefinitionContext extends ParserRuleContext {
		public Token name;
		public TerminalNode ID() { return getToken(GoldParser.ID, 0); }
		public TerminalNode PARAMETER_MODIFIER() { return getToken(GoldParser.PARAMETER_MODIFIER, 0); }
		public VarTypeContext varType() {
			return getRuleContext(VarTypeContext.class,0);
		}
		public ParameterDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterDefinition; }
	}

	public final ParameterDefinitionContext parameterDefinition() throws RecognitionException {
		ParameterDefinitionContext _localctx = new ParameterDefinitionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_parameterDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PARAMETER_MODIFIER) {
				{
				setState(65);
				match(PARAMETER_MODIFIER);
				}
			}

			setState(68);
			((ParameterDefinitionContext)_localctx).name = match(ID);
			setState(71);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(69);
				match(T__11);
				setState(70);
				varType();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public VariableDefinitionContext variableDefinition() {
			return getRuleContext(VariableDefinitionContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			variableDefinition();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariableDefinitionContext extends ParserRuleContext {
		public Token name;
		public VarTypeContext varType() {
			return getRuleContext(VarTypeContext.class,0);
		}
		public TerminalNode ID() { return getToken(GoldParser.ID, 0); }
		public VariableDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDefinition; }
	}

	public final VariableDefinitionContext variableDefinition() throws RecognitionException {
		VariableDefinitionContext _localctx = new VariableDefinitionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_variableDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			match(T__12);
			setState(76);
			((VariableDefinitionContext)_localctx).name = match(ID);
			setState(77);
			match(T__11);
			setState(78);
			varType();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarTypeContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GoldParser.ID, 0); }
		public VarTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varType; }
	}

	public final VarTypeContext varType() throws RecognitionException {
		VarTypeContext _localctx = new VarTypeContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_varType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u0013S\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0001\u0000\u0004\u0000\u0016\b\u0000\u000b"+
		"\u0000\f\u0000\u0017\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0003\u0001 \b\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002\'\b\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u00042\b\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u00046\b\u0004\u0001\u0004\u0005\u00049\b\u0004\n\u0004\f\u0004<\t\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0006\u0003\u0006"+
		"C\b\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006H\b\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001"+
		"\t\u0001\t\u0000\u0000\n\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0000\u0003\u0001\u0000\u0005\u0006\u0001\u0000\b\t\u0001\u0000\n\u000b"+
		"R\u0000\u0015\u0001\u0000\u0000\u0000\u0002\u001f\u0001\u0000\u0000\u0000"+
		"\u0004!\u0001\u0000\u0000\u0000\u0006(\u0001\u0000\u0000\u0000\b+\u0001"+
		"\u0000\u0000\u0000\n?\u0001\u0000\u0000\u0000\fB\u0001\u0000\u0000\u0000"+
		"\u000eI\u0001\u0000\u0000\u0000\u0010K\u0001\u0000\u0000\u0000\u0012P"+
		"\u0001\u0000\u0000\u0000\u0014\u0016\u0003\u0002\u0001\u0000\u0015\u0014"+
		"\u0001\u0000\u0000\u0000\u0016\u0017\u0001\u0000\u0000\u0000\u0017\u0015"+
		"\u0001\u0000\u0000\u0000\u0017\u0018\u0001\u0000\u0000\u0000\u0018\u0019"+
		"\u0001\u0000\u0000\u0000\u0019\u001a\u0005\u0000\u0000\u0001\u001a\u0001"+
		"\u0001\u0000\u0000\u0000\u001b \u0003\u0004\u0002\u0000\u001c \u0003\u0006"+
		"\u0003\u0000\u001d \u0003\n\u0005\u0000\u001e \u0003\b\u0004\u0000\u001f"+
		"\u001b\u0001\u0000\u0000\u0000\u001f\u001c\u0001\u0000\u0000\u0000\u001f"+
		"\u001d\u0001\u0000\u0000\u0000\u001f\u001e\u0001\u0000\u0000\u0000 \u0003"+
		"\u0001\u0000\u0000\u0000!\"\u0005\u0001\u0000\u0000\"&\u0005\u0013\u0000"+
		"\u0000#$\u0005\u0002\u0000\u0000$%\u0005\u0013\u0000\u0000%\'\u0005\u0003"+
		"\u0000\u0000&#\u0001\u0000\u0000\u0000&\'\u0001\u0000\u0000\u0000\'\u0005"+
		"\u0001\u0000\u0000\u0000()\u0005\u0004\u0000\u0000)*\u0005\u0013\u0000"+
		"\u0000*\u0007\u0001\u0000\u0000\u0000+,\u0007\u0000\u0000\u0000,5\u0005"+
		"\u0013\u0000\u0000-.\u0005\u0002\u0000\u0000.1\u0003\f\u0006\u0000/0\u0005"+
		"\u0007\u0000\u000002\u0003\f\u0006\u00001/\u0001\u0000\u0000\u000012\u0001"+
		"\u0000\u0000\u000023\u0001\u0000\u0000\u000034\u0005\u0003\u0000\u0000"+
		"46\u0001\u0000\u0000\u00005-\u0001\u0000\u0000\u000056\u0001\u0000\u0000"+
		"\u00006:\u0001\u0000\u0000\u000079\u0003\u000e\u0007\u000087\u0001\u0000"+
		"\u0000\u00009<\u0001\u0000\u0000\u0000:8\u0001\u0000\u0000\u0000:;\u0001"+
		"\u0000\u0000\u0000;=\u0001\u0000\u0000\u0000<:\u0001\u0000\u0000\u0000"+
		"=>\u0007\u0001\u0000\u0000>\t\u0001\u0000\u0000\u0000?@\u0007\u0002\u0000"+
		"\u0000@\u000b\u0001\u0000\u0000\u0000AC\u0005\u000e\u0000\u0000BA\u0001"+
		"\u0000\u0000\u0000BC\u0001\u0000\u0000\u0000CD\u0001\u0000\u0000\u0000"+
		"DG\u0005\u0013\u0000\u0000EF\u0005\f\u0000\u0000FH\u0003\u0012\t\u0000"+
		"GE\u0001\u0000\u0000\u0000GH\u0001\u0000\u0000\u0000H\r\u0001\u0000\u0000"+
		"\u0000IJ\u0003\u0010\b\u0000J\u000f\u0001\u0000\u0000\u0000KL\u0005\r"+
		"\u0000\u0000LM\u0005\u0013\u0000\u0000MN\u0005\f\u0000\u0000NO\u0003\u0012"+
		"\t\u0000O\u0011\u0001\u0000\u0000\u0000PQ\u0005\u0013\u0000\u0000Q\u0013"+
		"\u0001\u0000\u0000\u0000\b\u0017\u001f&15:BG";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}