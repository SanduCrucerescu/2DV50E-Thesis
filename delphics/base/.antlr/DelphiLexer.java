// Generated from /Users/ingmarfalk/uni/thesis/delphics/base/delphi.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class DelphiLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TokenIdentifier=1, TokenInteger=2, TokenReal=3, StringLiteral=4, ControlLiteral=5, 
		DigitSequence=6, HexDigitSequence=7, WS=8, UNICODE_BOM=9, COMMENT=10;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"TokenIdentifier", "TokenInteger", "TokenReal", "StringLiteral", "ControlLiteral", 
			"DigitSequence", "HexDigitSequence", "Digit", "HexDigit", "Letter", "Control", 
			"WS", "UNICODE_BOM", "COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, "'\\uFEFF'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TokenIdentifier", "TokenInteger", "TokenReal", "StringLiteral", 
			"ControlLiteral", "DigitSequence", "HexDigitSequence", "WS", "UNICODE_BOM", 
			"COMMENT"
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


	public DelphiLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "delphi.g4"; }

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
		"\u0004\u0000\n\u008b\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0001\u0000\u0001\u0000\u0003"+
		"\u0000 \b\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0005\u0000%\b\u0000"+
		"\n\u0000\f\u0000(\t\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0005\u00035\b\u0003\n\u0003\f\u00038\t\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0004\u0004\u0004=\b\u0004\u000b\u0004\f\u0004>\u0001"+
		"\u0005\u0004\u0005B\b\u0005\u000b\u0005\f\u0005C\u0001\u0006\u0004\u0006"+
		"G\b\u0006\u000b\u0006\f\u0006H\u0001\u0007\u0001\u0007\u0001\b\u0001\b"+
		"\u0003\bO\b\b\u0001\t\u0001\t\u0001\t\u0003\tT\b\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0003\n[\b\n\u0001\u000b\u0004\u000b^\b\u000b\u000b"+
		"\u000b\f\u000b_\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\r\u0001\r\u0001\r\u0001\r\u0005\rl\b\r\n\r\f\ro\t\r\u0001\r\u0003"+
		"\rr\b\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0005\ry\b\r\n\r\f\r|"+
		"\t\r\u0001\r\u0001\r\u0001\r\u0001\r\u0005\r\u0082\b\r\n\r\f\r\u0085\t"+
		"\r\u0001\r\u0003\r\u0088\b\r\u0001\r\u0001\r\u0002z\u0083\u0000\u000e"+
		"\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r"+
		"\u0007\u000f\u0000\u0011\u0000\u0013\u0000\u0015\u0000\u0017\b\u0019\t"+
		"\u001b\n\u0001\u0000\u0006\u0001\u0000\'\'\u0002\u0000AFaf\u0002\u0000"+
		"AZaz\u0001\u0000\u8000\ufeff\u8000\ufeff\u0003\u0000\t\n\f\r  \u0002\u0000"+
		"\n\n\r\r\u0099\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000"+
		"\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000"+
		"\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000"+
		"\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000"+
		"\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000"+
		"\u0001\u001f\u0001\u0000\u0000\u0000\u0003)\u0001\u0000\u0000\u0000\u0005"+
		"+\u0001\u0000\u0000\u0000\u0007/\u0001\u0000\u0000\u0000\t<\u0001\u0000"+
		"\u0000\u0000\u000bA\u0001\u0000\u0000\u0000\rF\u0001\u0000\u0000\u0000"+
		"\u000fJ\u0001\u0000\u0000\u0000\u0011N\u0001\u0000\u0000\u0000\u0013S"+
		"\u0001\u0000\u0000\u0000\u0015Z\u0001\u0000\u0000\u0000\u0017]\u0001\u0000"+
		"\u0000\u0000\u0019c\u0001\u0000\u0000\u0000\u001b\u0087\u0001\u0000\u0000"+
		"\u0000\u001d \u0003\u0013\t\u0000\u001e \u0005_\u0000\u0000\u001f\u001d"+
		"\u0001\u0000\u0000\u0000\u001f\u001e\u0001\u0000\u0000\u0000 &\u0001\u0000"+
		"\u0000\u0000!%\u0003\u0013\t\u0000\"%\u0003\u000f\u0007\u0000#%\u0005"+
		"_\u0000\u0000$!\u0001\u0000\u0000\u0000$\"\u0001\u0000\u0000\u0000$#\u0001"+
		"\u0000\u0000\u0000%(\u0001\u0000\u0000\u0000&$\u0001\u0000\u0000\u0000"+
		"&\'\u0001\u0000\u0000\u0000\'\u0002\u0001\u0000\u0000\u0000(&\u0001\u0000"+
		"\u0000\u0000)*\u0003\u000b\u0005\u0000*\u0004\u0001\u0000\u0000\u0000"+
		"+,\u0003\u000b\u0005\u0000,-\u0005.\u0000\u0000-.\u0003\u000b\u0005\u0000"+
		".\u0006\u0001\u0000\u0000\u0000/6\u0005\'\u0000\u000001\u0005\'\u0000"+
		"\u000012\u0005\'\u0000\u000025\u0005\'\u0000\u000035\b\u0000\u0000\u0000"+
		"40\u0001\u0000\u0000\u000043\u0001\u0000\u0000\u000058\u0001\u0000\u0000"+
		"\u000064\u0001\u0000\u0000\u000067\u0001\u0000\u0000\u000079\u0001\u0000"+
		"\u0000\u000086\u0001\u0000\u0000\u00009:\u0005\'\u0000\u0000:\b\u0001"+
		"\u0000\u0000\u0000;=\u0003\u0015\n\u0000<;\u0001\u0000\u0000\u0000=>\u0001"+
		"\u0000\u0000\u0000><\u0001\u0000\u0000\u0000>?\u0001\u0000\u0000\u0000"+
		"?\n\u0001\u0000\u0000\u0000@B\u0003\u000f\u0007\u0000A@\u0001\u0000\u0000"+
		"\u0000BC\u0001\u0000\u0000\u0000CA\u0001\u0000\u0000\u0000CD\u0001\u0000"+
		"\u0000\u0000D\f\u0001\u0000\u0000\u0000EG\u0003\u0011\b\u0000FE\u0001"+
		"\u0000\u0000\u0000GH\u0001\u0000\u0000\u0000HF\u0001\u0000\u0000\u0000"+
		"HI\u0001\u0000\u0000\u0000I\u000e\u0001\u0000\u0000\u0000JK\u000209\u0000"+
		"K\u0010\u0001\u0000\u0000\u0000LO\u0003\u000f\u0007\u0000MO\u0007\u0001"+
		"\u0000\u0000NL\u0001\u0000\u0000\u0000NM\u0001\u0000\u0000\u0000O\u0012"+
		"\u0001\u0000\u0000\u0000PT\u0007\u0002\u0000\u0000QR\u0002\u0080\u8000"+
		"\ufffe\u0000RT\b\u0003\u0000\u0000SP\u0001\u0000\u0000\u0000SQ\u0001\u0000"+
		"\u0000\u0000T\u0014\u0001\u0000\u0000\u0000UV\u0005#\u0000\u0000V[\u0003"+
		"\r\u0006\u0000WX\u0005#\u0000\u0000XY\u0005$\u0000\u0000Y[\u0003\r\u0006"+
		"\u0000ZU\u0001\u0000\u0000\u0000ZW\u0001\u0000\u0000\u0000[\u0016\u0001"+
		"\u0000\u0000\u0000\\^\u0007\u0004\u0000\u0000]\\\u0001\u0000\u0000\u0000"+
		"^_\u0001\u0000\u0000\u0000_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000"+
		"\u0000`a\u0001\u0000\u0000\u0000ab\u0006\u000b\u0000\u0000b\u0018\u0001"+
		"\u0000\u0000\u0000cd\u0005\u8000\ufeff\u0000\u0000de\u0001\u0000\u0000"+
		"\u0000ef\u0006\f\u0000\u0000f\u001a\u0001\u0000\u0000\u0000gh\u0005/\u0000"+
		"\u0000hi\u0005/\u0000\u0000im\u0001\u0000\u0000\u0000jl\b\u0005\u0000"+
		"\u0000kj\u0001\u0000\u0000\u0000lo\u0001\u0000\u0000\u0000mk\u0001\u0000"+
		"\u0000\u0000mn\u0001\u0000\u0000\u0000nq\u0001\u0000\u0000\u0000om\u0001"+
		"\u0000\u0000\u0000pr\u0005\r\u0000\u0000qp\u0001\u0000\u0000\u0000qr\u0001"+
		"\u0000\u0000\u0000rs\u0001\u0000\u0000\u0000s\u0088\u0005\n\u0000\u0000"+
		"tu\u0005(\u0000\u0000uv\u0005*\u0000\u0000vz\u0001\u0000\u0000\u0000w"+
		"y\t\u0000\u0000\u0000xw\u0001\u0000\u0000\u0000y|\u0001\u0000\u0000\u0000"+
		"z{\u0001\u0000\u0000\u0000zx\u0001\u0000\u0000\u0000{}\u0001\u0000\u0000"+
		"\u0000|z\u0001\u0000\u0000\u0000}~\u0005*\u0000\u0000~\u0088\u0005)\u0000"+
		"\u0000\u007f\u0083\u0005{\u0000\u0000\u0080\u0082\t\u0000\u0000\u0000"+
		"\u0081\u0080\u0001\u0000\u0000\u0000\u0082\u0085\u0001\u0000\u0000\u0000"+
		"\u0083\u0084\u0001\u0000\u0000\u0000\u0083\u0081\u0001\u0000\u0000\u0000"+
		"\u0084\u0086\u0001\u0000\u0000\u0000\u0085\u0083\u0001\u0000\u0000\u0000"+
		"\u0086\u0088\u0005}\u0000\u0000\u0087g\u0001\u0000\u0000\u0000\u0087t"+
		"\u0001\u0000\u0000\u0000\u0087\u007f\u0001\u0000\u0000\u0000\u0088\u0089"+
		"\u0001\u0000\u0000\u0000\u0089\u008a\u0006\r\u0000\u0000\u008a\u001c\u0001"+
		"\u0000\u0000\u0000\u0012\u0000\u001f$&46>CHNSZ_mqz\u0083\u0087\u0001\u0006"+
		"\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}