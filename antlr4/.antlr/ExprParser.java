// Generated from /home/razzz/dev/projects/gold-lang-lsp-go/antlr4/Expr.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class ExprParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, NEWLINE=4, INT_LITERAL=5, DEICMAL_LITERAL=6, ID=7;
	public static final int
		RULE_prog = 0, RULE_goldClass = 1, RULE_goldModule = 2, RULE_goldMethod = 3, 
		RULE_goldFunction = 4, RULE_goldClassDefinition = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "goldClass", "goldModule", "goldMethod", "goldFunction", "goldClassDefinition"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'class'", "'('", "')'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, "NEWLINE", "INT_LITERAL", "DEICMAL_LITERAL", 
			"ID"
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
	public String getGrammarFileName() { return "Expr.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ExprParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgContext extends ParserRuleContext {
		public GoldClassContext goldClass() {
			return getRuleContext(GoldClassContext.class,0);
		}
		public TerminalNode EOF() { return getToken(ExprParser.EOF, 0); }
		public GoldModuleContext goldModule() {
			return getRuleContext(GoldModuleContext.class,0);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		try {
			setState(18);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(12);
				goldClass();
				setState(13);
				match(EOF);
				}
				break;
			case EOF:
				enterOuterAlt(_localctx, 2);
				{
				setState(15);
				goldModule();
				setState(16);
				match(EOF);
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
	public static class GoldClassContext extends ParserRuleContext {
		public GoldClassDefinitionContext goldClassDefinition() {
			return getRuleContext(GoldClassDefinitionContext.class,0);
		}
		public GoldClassContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldClass; }
	}

	public final GoldClassContext goldClass() throws RecognitionException {
		GoldClassContext _localctx = new GoldClassContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_goldClass);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(20);
			goldClassDefinition();
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
	public static class GoldModuleContext extends ParserRuleContext {
		public GoldModuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldModule; }
	}

	public final GoldModuleContext goldModule() throws RecognitionException {
		GoldModuleContext _localctx = new GoldModuleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_goldModule);
		try {
			enterOuterAlt(_localctx, 1);
			{
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
	public static class GoldMethodContext extends ParserRuleContext {
		public GoldMethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldMethod; }
	}

	public final GoldMethodContext goldMethod() throws RecognitionException {
		GoldMethodContext _localctx = new GoldMethodContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_goldMethod);
		try {
			enterOuterAlt(_localctx, 1);
			{
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
	public static class GoldFunctionContext extends ParserRuleContext {
		public GoldFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldFunction; }
	}

	public final GoldFunctionContext goldFunction() throws RecognitionException {
		GoldFunctionContext _localctx = new GoldFunctionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_goldFunction);
		try {
			enterOuterAlt(_localctx, 1);
			{
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
		public List<TerminalNode> ID() { return getTokens(ExprParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(ExprParser.ID, i);
		}
		public GoldClassDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goldClassDefinition; }
	}

	public final GoldClassDefinitionContext goldClassDefinition() throws RecognitionException {
		GoldClassDefinitionContext _localctx = new GoldClassDefinitionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_goldClassDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(28);
			match(T__0);
			setState(29);
			((GoldClassDefinitionContext)_localctx).name = match(ID);
			{
			setState(30);
			match(T__1);
			setState(31);
			((GoldClassDefinitionContext)_localctx).parentClass = match(ID);
			setState(32);
			match(T__2);
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

	public static final String _serializedATN =
		"\u0004\u0001\u0007#\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0000\u0003\u0000\u0013\b\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0000\u0000\u0006\u0000\u0002\u0004\u0006\b\n\u0000\u0000\u001d"+
		"\u0000\u0012\u0001\u0000\u0000\u0000\u0002\u0014\u0001\u0000\u0000\u0000"+
		"\u0004\u0016\u0001\u0000\u0000\u0000\u0006\u0018\u0001\u0000\u0000\u0000"+
		"\b\u001a\u0001\u0000\u0000\u0000\n\u001c\u0001\u0000\u0000\u0000\f\r\u0003"+
		"\u0002\u0001\u0000\r\u000e\u0005\u0000\u0000\u0001\u000e\u0013\u0001\u0000"+
		"\u0000\u0000\u000f\u0010\u0003\u0004\u0002\u0000\u0010\u0011\u0005\u0000"+
		"\u0000\u0001\u0011\u0013\u0001\u0000\u0000\u0000\u0012\f\u0001\u0000\u0000"+
		"\u0000\u0012\u000f\u0001\u0000\u0000\u0000\u0013\u0001\u0001\u0000\u0000"+
		"\u0000\u0014\u0015\u0003\n\u0005\u0000\u0015\u0003\u0001\u0000\u0000\u0000"+
		"\u0016\u0017\u0001\u0000\u0000\u0000\u0017\u0005\u0001\u0000\u0000\u0000"+
		"\u0018\u0019\u0001\u0000\u0000\u0000\u0019\u0007\u0001\u0000\u0000\u0000"+
		"\u001a\u001b\u0001\u0000\u0000\u0000\u001b\t\u0001\u0000\u0000\u0000\u001c"+
		"\u001d\u0005\u0001\u0000\u0000\u001d\u001e\u0005\u0007\u0000\u0000\u001e"+
		"\u001f\u0005\u0002\u0000\u0000\u001f \u0005\u0007\u0000\u0000 !\u0005"+
		"\u0003\u0000\u0000!\u000b\u0001\u0000\u0000\u0000\u0001\u0012";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}