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
		RULE_goldProg = 0, RULE_goldDefinittion = 1, RULE_goldProcedureDefinition = 2, 
		RULE_goldFunctionDefinition = 3, RULE_parameterDefinition = 4, RULE_statement = 5, 
		RULE_variableDefinition = 6, RULE_goldClassDefinition = 7, RULE_goldModuleDefinition = 8, 
		RULE_varType = 9;
	private static String[] makeRuleNames() {
		return new String[] {
			"goldProg", "goldDefinittion", "goldProcedureDefinition", "goldFunctionDefinition", 
			"parameterDefinition", "statement", "variableDefinition", "goldClassDefinition", 
			"goldModuleDefinition", "varType"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'procedure'", "'proc'", "'('", "','", "')'", "'end'", "'endproc'", 
			"'function'", "'func'", "':'", "'var'", "'class'", "'module'"
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
		public List<GoldDefinittionContext> goldDefinittion() {
			return getRuleContexts(GoldDefinittionContext.class);
		}
		public GoldDefinittionContext goldDefinittion(int i) {
			return getRuleContext(GoldDefinittionContext.class,i);
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
				goldDefinittion();
				}
				}
				setState(23); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 13190L) != 0) );
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
	public static class GoldDefinittionContext extends ParserRuleContext {
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
		public GoldDefinittionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldDefinittion; }
	}

	public final GoldDefinittionContext goldDefinittion() throws RecognitionException {
		GoldDefinittionContext _localctx = new GoldDefinittionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_goldDefinittion);
		try {
			setState(31);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__11:
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				goldClassDefinition();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 2);
				{
				setState(28);
				goldModuleDefinition();
				}
				break;
			case T__7:
			case T__8:
				enterOuterAlt(_localctx, 3);
				{
				setState(29);
				goldFunctionDefinition();
				}
				break;
			case T__0:
			case T__1:
			case T__6:
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
	public static class GoldProcedureDefinitionContext extends ParserRuleContext {
		public Token name;
		public ParameterDefinitionContext parameterDefinition;
		public List<ParameterDefinitionContext> params = new ArrayList<ParameterDefinitionContext>();
		public TerminalNode ID() { return getToken(GoldParser.ID, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<ParameterDefinitionContext> parameterDefinition() {
			return getRuleContexts(ParameterDefinitionContext.class);
		}
		public ParameterDefinitionContext parameterDefinition(int i) {
			return getRuleContext(ParameterDefinitionContext.class,i);
		}
		public GoldProcedureDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldProcedureDefinition; }
	}

	public final GoldProcedureDefinitionContext goldProcedureDefinition() throws RecognitionException {
		GoldProcedureDefinitionContext _localctx = new GoldProcedureDefinitionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_goldProcedureDefinition);
		int _la;
		try {
			setState(54);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(33);
				match(T__0);
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				match(T__1);
				setState(35);
				((GoldProcedureDefinitionContext)_localctx).name = match(ID);
				setState(44);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__2) {
					{
					setState(36);
					match(T__2);
					setState(37);
					((GoldProcedureDefinitionContext)_localctx).parameterDefinition = parameterDefinition();
					((GoldProcedureDefinitionContext)_localctx).params.add(((GoldProcedureDefinitionContext)_localctx).parameterDefinition);
					setState(40);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==T__3) {
						{
						setState(38);
						match(T__3);
						setState(39);
						((GoldProcedureDefinitionContext)_localctx).parameterDefinition = parameterDefinition();
						((GoldProcedureDefinitionContext)_localctx).params.add(((GoldProcedureDefinitionContext)_localctx).parameterDefinition);
						}
					}

					setState(42);
					match(T__4);
					}
				}

				setState(49);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__10) {
					{
					{
					setState(46);
					statement();
					}
					}
					setState(51);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(52);
				match(T__5);
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 3);
				{
				setState(53);
				match(T__6);
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
	public static class GoldFunctionDefinitionContext extends ParserRuleContext {
		public GoldFunctionDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldFunctionDefinition; }
	}

	public final GoldFunctionDefinitionContext goldFunctionDefinition() throws RecognitionException {
		GoldFunctionDefinitionContext _localctx = new GoldFunctionDefinitionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_goldFunctionDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
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
		enterRule(_localctx, 8, RULE_parameterDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PARAMETER_MODIFIER) {
				{
				setState(58);
				match(PARAMETER_MODIFIER);
				}
			}

			setState(61);
			((ParameterDefinitionContext)_localctx).name = match(ID);
			setState(64);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__9) {
				{
				setState(62);
				match(T__9);
				setState(63);
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
		enterRule(_localctx, 10, RULE_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
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
		enterRule(_localctx, 12, RULE_variableDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			match(T__10);
			setState(69);
			((VariableDefinitionContext)_localctx).name = match(ID);
			setState(70);
			match(T__9);
			setState(71);
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
		enterRule(_localctx, 14, RULE_goldClassDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			match(T__11);
			setState(74);
			((GoldClassDefinitionContext)_localctx).name = match(ID);
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__2) {
				{
				setState(75);
				match(T__2);
				setState(76);
				((GoldClassDefinitionContext)_localctx).parentClass = match(ID);
				setState(77);
				match(T__4);
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
		enterRule(_localctx, 16, RULE_goldModuleDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			match(T__12);
			setState(81);
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
			setState(83);
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
		"\u0004\u0001\u0013V\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0001\u0000\u0004\u0000\u0016\b\u0000\u000b"+
		"\u0000\f\u0000\u0017\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0003\u0001 \b\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002)\b"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002-\b\u0002\u0001\u0002\u0005"+
		"\u00020\b\u0002\n\u0002\f\u00023\t\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u00027\b\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0003\u0004<\b\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004A\b\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007"+
		"O\b\u0007\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0000\u0000"+
		"\n\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0000\u0001\u0001\u0000"+
		"\b\tW\u0000\u0015\u0001\u0000\u0000\u0000\u0002\u001f\u0001\u0000\u0000"+
		"\u0000\u00046\u0001\u0000\u0000\u0000\u00068\u0001\u0000\u0000\u0000\b"+
		";\u0001\u0000\u0000\u0000\nB\u0001\u0000\u0000\u0000\fD\u0001\u0000\u0000"+
		"\u0000\u000eI\u0001\u0000\u0000\u0000\u0010P\u0001\u0000\u0000\u0000\u0012"+
		"S\u0001\u0000\u0000\u0000\u0014\u0016\u0003\u0002\u0001\u0000\u0015\u0014"+
		"\u0001\u0000\u0000\u0000\u0016\u0017\u0001\u0000\u0000\u0000\u0017\u0015"+
		"\u0001\u0000\u0000\u0000\u0017\u0018\u0001\u0000\u0000\u0000\u0018\u0019"+
		"\u0001\u0000\u0000\u0000\u0019\u001a\u0005\u0000\u0000\u0001\u001a\u0001"+
		"\u0001\u0000\u0000\u0000\u001b \u0003\u000e\u0007\u0000\u001c \u0003\u0010"+
		"\b\u0000\u001d \u0003\u0006\u0003\u0000\u001e \u0003\u0004\u0002\u0000"+
		"\u001f\u001b\u0001\u0000\u0000\u0000\u001f\u001c\u0001\u0000\u0000\u0000"+
		"\u001f\u001d\u0001\u0000\u0000\u0000\u001f\u001e\u0001\u0000\u0000\u0000"+
		" \u0003\u0001\u0000\u0000\u0000!7\u0005\u0001\u0000\u0000\"#\u0005\u0002"+
		"\u0000\u0000#,\u0005\u0013\u0000\u0000$%\u0005\u0003\u0000\u0000%(\u0003"+
		"\b\u0004\u0000&\'\u0005\u0004\u0000\u0000\')\u0003\b\u0004\u0000(&\u0001"+
		"\u0000\u0000\u0000()\u0001\u0000\u0000\u0000)*\u0001\u0000\u0000\u0000"+
		"*+\u0005\u0005\u0000\u0000+-\u0001\u0000\u0000\u0000,$\u0001\u0000\u0000"+
		"\u0000,-\u0001\u0000\u0000\u0000-1\u0001\u0000\u0000\u0000.0\u0003\n\u0005"+
		"\u0000/.\u0001\u0000\u0000\u000003\u0001\u0000\u0000\u00001/\u0001\u0000"+
		"\u0000\u000012\u0001\u0000\u0000\u000024\u0001\u0000\u0000\u000031\u0001"+
		"\u0000\u0000\u000047\u0005\u0006\u0000\u000057\u0005\u0007\u0000\u0000"+
		"6!\u0001\u0000\u0000\u00006\"\u0001\u0000\u0000\u000065\u0001\u0000\u0000"+
		"\u00007\u0005\u0001\u0000\u0000\u000089\u0007\u0000\u0000\u00009\u0007"+
		"\u0001\u0000\u0000\u0000:<\u0005\u000e\u0000\u0000;:\u0001\u0000\u0000"+
		"\u0000;<\u0001\u0000\u0000\u0000<=\u0001\u0000\u0000\u0000=@\u0005\u0013"+
		"\u0000\u0000>?\u0005\n\u0000\u0000?A\u0003\u0012\t\u0000@>\u0001\u0000"+
		"\u0000\u0000@A\u0001\u0000\u0000\u0000A\t\u0001\u0000\u0000\u0000BC\u0003"+
		"\f\u0006\u0000C\u000b\u0001\u0000\u0000\u0000DE\u0005\u000b\u0000\u0000"+
		"EF\u0005\u0013\u0000\u0000FG\u0005\n\u0000\u0000GH\u0003\u0012\t\u0000"+
		"H\r\u0001\u0000\u0000\u0000IJ\u0005\f\u0000\u0000JN\u0005\u0013\u0000"+
		"\u0000KL\u0005\u0003\u0000\u0000LM\u0005\u0013\u0000\u0000MO\u0005\u0005"+
		"\u0000\u0000NK\u0001\u0000\u0000\u0000NO\u0001\u0000\u0000\u0000O\u000f"+
		"\u0001\u0000\u0000\u0000PQ\u0005\r\u0000\u0000QR\u0005\u0013\u0000\u0000"+
		"R\u0011\u0001\u0000\u0000\u0000ST\u0005\u0013\u0000\u0000T\u0013\u0001"+
		"\u0000\u0000\u0000\t\u0017\u001f(,16;@N";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}