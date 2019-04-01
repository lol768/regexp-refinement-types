# -*- coding: utf-8 -*-
"""
    pygments.lexers.dafny
    ~~~~~~~~~~~~~~~~~~~~

    Lexers for Dafny.

    :copyright: Copyright 2006-2017 by the Pygments team, see AUTHORS.
    :license: BSD, see LICENSE for details.
"""

import re

from pygments.lexer import RegexLexer, include, bygroups, default, using, \
    this, words, combined
from pygments.token import Text, Comment, Operator, Keyword, Name, String, \
    Number, Punctuation, Other
from pygments.util import get_bool_opt, iteritems
import pygments.unistring as uni

__all__ = ['DafnyLexer']


class DafnyLexer(RegexLexer):
    """
    Pygments Lexer for Dafny files (.dfy).
    Broken af, but good enough for the dissertation report.
    .. versionadded:: 1.0
    """

    name = 'Dafny'
    aliases = ['dafny']
    filenames = ['*.dfy']
    mimetypes = ['text/dafny']

    flags = re.DOTALL | re.UNICODE | re.IGNORECASE | re.MULTILINE

    tokens = {
        'commentsandwhitespace': [
            (r'\s+', Text),
            (r'\/\/.*?\n', Comment.Single),
            (r'\/\*.*?\*\/', Comment.Multiline)
        ],
        'root': [
            include('commentsandwhitespace'),
            (r'(\.\d+|[0-9]+\.[0-9]*)([eE][-+]?[0-9]+)?', Number.Float),
            (r'[0-9]+', Number.Integer),
            (r'\.\.\.|=>', Punctuation),
            (r'\+\+|--|~|&&|\?|:|\|\||\\(?=\n)|'
             r'(<<|>>>?|==?|!=?|[-<>+*%&|^\/:])=?', Operator),
            (r'[{(\[;,]', Punctuation),
            (r'[})\].]', Punctuation),
            (r'(in|while|do|return|if|else|then|requires|print|decreases|ensures|var)\b', Keyword),
            (r'(function method\b|function\b|method\b)(\s*)(.*?)(\()',
             bygroups(Keyword.Declaration, Text, Name.Identifier, Punctuation), 'arguments'),
            (r'(function method\b|function\b|method\b)', Keyword.Declaration),

            (r'(true|false)\b', Keyword.Constant),
            (r'"(\\\\|\\"|[^"])*"', String.Double),
            (r'[A-Za-z][A-Za-z0-9]*', Name.Identifier),

        ],
        'arguments': [
            (r'\)', Punctuation, ('#pop', 'returntype')),
            (r'(.*?)(:)(\s*)([^\ )]*)(,?)(\s*)', bygroups(Name.Identifier, Punctuation, Text, Keyword.Type, Punctuation, Text)),
        ],
        'returntype': [
            (r'(\s*)(:)(\s*)(.*?)(\s)', bygroups(Text, Punctuation, Text, Keyword.Type, Text), '#pop'),

        ]

    }
