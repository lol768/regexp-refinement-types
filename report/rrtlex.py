# -*- coding: utf-8 -*-
"""
    pygments.lexers.rrt
    ~~~~~~~~~~~~~~~~~~~~

    Lexer for RRT.
"""

import re

from pygments.lexer import RegexLexer, include, bygroups, default, using, \
    this, words, combined
from pygments.token import Text, Comment, Operator, Keyword, Name, String, \
    Number, Punctuation, Other
from pygments.util import get_bool_opt, iteritems
import pygments.unistring as uni

__all__ = ['RrtLexer']


class RrtLexer(RegexLexer):
    """
    Pygments Lexer for RegexRefinementType files (.rrt).
    Broken af, but good enough for the dissertation report.
    .. versionadded:: 1.0
    """

    name = 'RegexRefinementTypes'
    aliases = ['rrt']
    filenames = ['*.rrt']
    mimetypes = ['text/rrt']

    flags = re.DOTALL | re.UNICODE | re.IGNORECASE | re.MULTILINE

    tokens = {
        'commentsandwhitespace': [
            (r'\s+', Text),
            (r'\/\/.*?\n', Comment.Single),
        ],
        'root': [
            include('commentsandwhitespace'),
            (r'(\.\d+|[0-9]+\.[0-9]*)([eE][-+]?[0-9]+)?', Number.Float),
            (r'[0-9]+', Number.Integer),
            (r'(var)(\s+)([A-Za-z][A-Za-z0-9]*)(:)(\s+)', bygroups(Keyword, Text, Name.Identifier, Punctuation, Text), 'type'),

            (r'!(?:[A-Za-z][A-Za-z0-9]*\.?)+', Name.Function.Magic),

            (r'\.\.\.|=>', Punctuation),

            (r'\+\+|--|~|&&|\?|::|:|\|\||\\(?=\n)|'
             r'(<<|>>>?|==?|!=?|[-<>+*%&|^\/:])=?', Operator),
            (r'[{(\[;,]', Punctuation),
            (r'[})\].]', Punctuation),
            (r'(forall)(\s*)([A-Za-z][A-Za-z0-9]*)(\s*)(:)(\s*)(.*?)(\s+)(::)', bygroups(Keyword, Text, Name.Identifier, Text, Punctuation, Text, Keyword.Type, Text, Operator)),
            (r'(return|if|else|print|decreases|ensures|var)\b', Keyword),
            (r'(function\b)(\s*)(.*?)(\()',
             bygroups(Keyword.Declaration, Text, Name.Identifier, Punctuation), 'arguments'),

            (r'(true|false)\b', Keyword.Constant),
            (r'"(\\\\|\\"|[^"])*"', String.Double),
            (r'\'[^\']\'', String.Char),
            (r'[A-Za-z][A-Za-z0-9]*', Name.Identifier),

        ],
        'arguments': [
            (r'\)', Punctuation, ('#pop', 'returntype')),
            (r'(,)(\s*)', bygroups(Punctuation, Text)),
            (r'([^,\)]*?)(:)(\s*)', bygroups(Name.Identifier, Punctuation, Text), 'type'),
            (r'(\))', Punctuation, '#pop'),

        ],
        'returntype': [
            (r'(\s*)(:)(\s*)(.*?)(\s)', bygroups(Text, Punctuation, Text, Keyword.Type, Text), '#pop'),
        ],
        'type': [
            (r'uint', Keyword.Type, ('#pop', 'refinementint')),
            (r'string', Keyword.Type,  ('#pop', 'refinementstring')),
            (r'void', Keyword.Type, '#pop'),
        ],
        'refinementint': [
            (r'(\[)([<>]=?)(\s*)([0-9]+)(\])', bygroups(Punctuation, Punctuation, Text, Number.Integer, Punctuation), '#pop'),
            (r'\b', Text, '#pop')
        ],
        'refinementstring': [
            (r'(\[)(/.*?/)(\])', bygroups(Punctuation, String.Regex, Punctuation), '#pop'),
            (r'\b', Text, '#pop')
        ]

    }
