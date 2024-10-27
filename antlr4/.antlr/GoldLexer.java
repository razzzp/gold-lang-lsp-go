// Generated from /home/razzz/dev/projects/gold-lang-lsp-go/antlr4/Gold.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class GoldLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, PARAMETER_MODIFIER=14, WHITESPACE=15, 
		INT_LITERAL=16, DEICMAL_LITERAL=17, STRING_LITERAL=18, ID=19;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "PARAMETER_MODIFIER", "WHITESPACE", 
			"INT_LITERAL", "DEICMAL_LITERAL", "STRING_LITERAL", "ID"
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


	public GoldLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Gold.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0013\u00a5\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\rw\b"+
		"\r\u0001\u000e\u0004\u000ez\b\u000e\u000b\u000e\f\u000e{\u0001\u000e\u0001"+
		"\u000e\u0001\u000f\u0003\u000f\u0081\b\u000f\u0001\u000f\u0004\u000f\u0084"+
		"\b\u000f\u000b\u000f\f\u000f\u0085\u0001\u0010\u0003\u0010\u0089\b\u0010"+
		"\u0001\u0010\u0004\u0010\u008c\b\u0010\u000b\u0010\f\u0010\u008d\u0001"+
		"\u0010\u0001\u0010\u0004\u0010\u0092\b\u0010\u000b\u0010\f\u0010\u0093"+
		"\u0001\u0011\u0001\u0011\u0005\u0011\u0098\b\u0011\n\u0011\f\u0011\u009b"+
		"\t\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0005\u0012\u00a1"+
		"\b\u0012\n\u0012\f\u0012\u00a4\t\u0012\u0001\u0099\u0000\u0013\u0001\u0001"+
		"\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f"+
		"\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f"+
		"\u001f\u0010!\u0011#\u0012%\u0013\u0001\u0000\u0015\u0002\u0000PPpp\u0002"+
		"\u0000RRrr\u0002\u0000OOoo\u0002\u0000CCcc\u0002\u0000EEee\u0002\u0000"+
		"DDdd\u0002\u0000UUuu\u0002\u0000NNnn\u0002\u0000FFff\u0002\u0000TTtt\u0002"+
		"\u0000IIii\u0002\u0000VVvv\u0002\u0000AAaa\u0002\u0000LLll\u0002\u0000"+
		"SSss\u0002\u0000MMmm\u0003\u0000\t\n\r\r  \u0002\u0000++--\u0001\u0000"+
		"09\u0003\u0000AZ__az\u0004\u000009AZ__az\u00ae\u0000\u0001\u0001\u0000"+
		"\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000"+
		"\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000"+
		"\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000"+
		"\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000"+
		"\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000"+
		"\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000"+
		"\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000"+
		"\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000\u0000"+
		"#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000\u0000\u0001\'\u0001"+
		"\u0000\u0000\u0000\u00031\u0001\u0000\u0000\u0000\u00056\u0001\u0000\u0000"+
		"\u0000\u00078\u0001\u0000\u0000\u0000\t:\u0001\u0000\u0000\u0000\u000b"+
		"<\u0001\u0000\u0000\u0000\r@\u0001\u0000\u0000\u0000\u000fH\u0001\u0000"+
		"\u0000\u0000\u0011Q\u0001\u0000\u0000\u0000\u0013V\u0001\u0000\u0000\u0000"+
		"\u0015X\u0001\u0000\u0000\u0000\u0017\\\u0001\u0000\u0000\u0000\u0019"+
		"b\u0001\u0000\u0000\u0000\u001bv\u0001\u0000\u0000\u0000\u001dy\u0001"+
		"\u0000\u0000\u0000\u001f\u0080\u0001\u0000\u0000\u0000!\u0088\u0001\u0000"+
		"\u0000\u0000#\u0095\u0001\u0000\u0000\u0000%\u009e\u0001\u0000\u0000\u0000"+
		"\'(\u0007\u0000\u0000\u0000()\u0007\u0001\u0000\u0000)*\u0007\u0002\u0000"+
		"\u0000*+\u0007\u0003\u0000\u0000+,\u0007\u0004\u0000\u0000,-\u0007\u0005"+
		"\u0000\u0000-.\u0007\u0006\u0000\u0000./\u0007\u0001\u0000\u0000/0\u0007"+
		"\u0004\u0000\u00000\u0002\u0001\u0000\u0000\u000012\u0007\u0000\u0000"+
		"\u000023\u0007\u0001\u0000\u000034\u0007\u0002\u0000\u000045\u0007\u0003"+
		"\u0000\u00005\u0004\u0001\u0000\u0000\u000067\u0005(\u0000\u00007\u0006"+
		"\u0001\u0000\u0000\u000089\u0005,\u0000\u00009\b\u0001\u0000\u0000\u0000"+
		":;\u0005)\u0000\u0000;\n\u0001\u0000\u0000\u0000<=\u0007\u0004\u0000\u0000"+
		"=>\u0007\u0007\u0000\u0000>?\u0007\u0005\u0000\u0000?\f\u0001\u0000\u0000"+
		"\u0000@A\u0007\u0004\u0000\u0000AB\u0007\u0007\u0000\u0000BC\u0007\u0005"+
		"\u0000\u0000CD\u0007\u0000\u0000\u0000DE\u0007\u0001\u0000\u0000EF\u0007"+
		"\u0002\u0000\u0000FG\u0007\u0003\u0000\u0000G\u000e\u0001\u0000\u0000"+
		"\u0000HI\u0007\b\u0000\u0000IJ\u0007\u0006\u0000\u0000JK\u0007\u0007\u0000"+
		"\u0000KL\u0007\u0003\u0000\u0000LM\u0007\t\u0000\u0000MN\u0007\n\u0000"+
		"\u0000NO\u0007\u0002\u0000\u0000OP\u0007\u0007\u0000\u0000P\u0010\u0001"+
		"\u0000\u0000\u0000QR\u0007\b\u0000\u0000RS\u0007\u0006\u0000\u0000ST\u0007"+
		"\u0007\u0000\u0000TU\u0007\u0003\u0000\u0000U\u0012\u0001\u0000\u0000"+
		"\u0000VW\u0005:\u0000\u0000W\u0014\u0001\u0000\u0000\u0000XY\u0007\u000b"+
		"\u0000\u0000YZ\u0007\f\u0000\u0000Z[\u0007\u0001\u0000\u0000[\u0016\u0001"+
		"\u0000\u0000\u0000\\]\u0007\u0003\u0000\u0000]^\u0007\r\u0000\u0000^_"+
		"\u0007\f\u0000\u0000_`\u0007\u000e\u0000\u0000`a\u0007\u000e\u0000\u0000"+
		"a\u0018\u0001\u0000\u0000\u0000bc\u0007\u000f\u0000\u0000cd\u0007\u0002"+
		"\u0000\u0000de\u0007\u0005\u0000\u0000ef\u0007\u0006\u0000\u0000fg\u0007"+
		"\r\u0000\u0000gh\u0007\u0004\u0000\u0000h\u001a\u0001\u0000\u0000\u0000"+
		"ij\u0007\n\u0000\u0000jk\u0007\u0007\u0000\u0000kl\u0007\u0002\u0000\u0000"+
		"lm\u0007\u0006\u0000\u0000mw\u0007\t\u0000\u0000no\u0007\u000b\u0000\u0000"+
		"op\u0007\f\u0000\u0000pw\u0007\u0001\u0000\u0000qr\u0007\u0003\u0000\u0000"+
		"rs\u0007\u0002\u0000\u0000st\u0007\u0007\u0000\u0000tu\u0007\u000e\u0000"+
		"\u0000uw\u0007\t\u0000\u0000vi\u0001\u0000\u0000\u0000vn\u0001\u0000\u0000"+
		"\u0000vq\u0001\u0000\u0000\u0000w\u001c\u0001\u0000\u0000\u0000xz\u0007"+
		"\u0010\u0000\u0000yx\u0001\u0000\u0000\u0000z{\u0001\u0000\u0000\u0000"+
		"{y\u0001\u0000\u0000\u0000{|\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000"+
		"\u0000}~\u0006\u000e\u0000\u0000~\u001e\u0001\u0000\u0000\u0000\u007f"+
		"\u0081\u0007\u0011\u0000\u0000\u0080\u007f\u0001\u0000\u0000\u0000\u0080"+
		"\u0081\u0001\u0000\u0000\u0000\u0081\u0083\u0001\u0000\u0000\u0000\u0082"+
		"\u0084\u0007\u0012\u0000\u0000\u0083\u0082\u0001\u0000\u0000\u0000\u0084"+
		"\u0085\u0001\u0000\u0000\u0000\u0085\u0083\u0001\u0000\u0000\u0000\u0085"+
		"\u0086\u0001\u0000\u0000\u0000\u0086 \u0001\u0000\u0000\u0000\u0087\u0089"+
		"\u0007\u0011\u0000\u0000\u0088\u0087\u0001\u0000\u0000\u0000\u0088\u0089"+
		"\u0001\u0000\u0000\u0000\u0089\u008b\u0001\u0000\u0000\u0000\u008a\u008c"+
		"\u0007\u0012\u0000\u0000\u008b\u008a\u0001\u0000\u0000\u0000\u008c\u008d"+
		"\u0001\u0000\u0000\u0000\u008d\u008b\u0001\u0000\u0000\u0000\u008d\u008e"+
		"\u0001\u0000\u0000\u0000\u008e\u008f\u0001\u0000\u0000\u0000\u008f\u0091"+
		"\u0005.\u0000\u0000\u0090\u0092\u0007\u0012\u0000\u0000\u0091\u0090\u0001"+
		"\u0000\u0000\u0000\u0092\u0093\u0001\u0000\u0000\u0000\u0093\u0091\u0001"+
		"\u0000\u0000\u0000\u0093\u0094\u0001\u0000\u0000\u0000\u0094\"\u0001\u0000"+
		"\u0000\u0000\u0095\u0099\u0005\'\u0000\u0000\u0096\u0098\t\u0000\u0000"+
		"\u0000\u0097\u0096\u0001\u0000\u0000\u0000\u0098\u009b\u0001\u0000\u0000"+
		"\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u0099\u0097\u0001\u0000\u0000"+
		"\u0000\u009a\u009c\u0001\u0000\u0000\u0000\u009b\u0099\u0001\u0000\u0000"+
		"\u0000\u009c\u009d\u0005\'\u0000\u0000\u009d$\u0001\u0000\u0000\u0000"+
		"\u009e\u00a2\u0007\u0013\u0000\u0000\u009f\u00a1\u0007\u0014\u0000\u0000"+
		"\u00a0\u009f\u0001\u0000\u0000\u0000\u00a1\u00a4\u0001\u0000\u0000\u0000"+
		"\u00a2\u00a0\u0001\u0000\u0000\u0000\u00a2\u00a3\u0001\u0000\u0000\u0000"+
		"\u00a3&\u0001\u0000\u0000\u0000\u00a4\u00a2\u0001\u0000\u0000\u0000\n"+
		"\u0000v{\u0080\u0085\u0088\u008d\u0093\u0099\u00a2\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}