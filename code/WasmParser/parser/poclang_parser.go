// Code generated from PocLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PocLang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 64, 371,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 5,
	2, 78, 10, 2, 3, 2, 5, 2, 81, 10, 2, 3, 2, 5, 2, 84, 10, 2, 3, 2, 3, 2,
	5, 2, 88, 10, 2, 3, 2, 3, 2, 7, 2, 92, 10, 2, 12, 2, 14, 2, 95, 11, 2,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 104, 10, 4, 3, 5, 3, 5,
	3, 5, 5, 5, 109, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 136, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 149, 10, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 162,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 5, 11, 176, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 190, 10, 13,
	12, 13, 14, 13, 193, 11, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 7, 16, 214, 10, 16, 12, 16, 14, 16, 217, 11, 16, 6, 16,
	219, 10, 16, 13, 16, 14, 16, 220, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 5, 17, 232, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 245, 10, 17, 12,
	17, 14, 17, 248, 11, 17, 3, 18, 3, 18, 5, 18, 252, 10, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 262, 10, 19, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 271, 10, 21, 12, 21, 14,
	21, 274, 11, 21, 5, 21, 276, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 7, 21, 285, 10, 21, 12, 21, 14, 21, 288, 11, 21, 5, 21, 290,
	10, 21, 3, 21, 3, 21, 5, 21, 294, 10, 21, 3, 22, 3, 22, 5, 22, 298, 10,
	22, 7, 22, 300, 10, 22, 12, 22, 14, 22, 303, 11, 22, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 5, 24, 311, 10, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 5, 26, 318, 10, 26, 3, 27, 3, 27, 3, 27, 5, 27, 323, 10, 27, 3, 28,
	3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 335,
	10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 5, 32, 343, 10, 32, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 362, 10, 37, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 38, 5, 38, 369, 10, 38, 3, 38, 2, 3, 32, 39, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 2, 6, 4,
	2, 4, 4, 7, 7, 3, 2, 5, 6, 3, 2, 21, 25, 4, 2, 42, 42, 45, 45, 2, 385,
	2, 93, 3, 2, 2, 2, 4, 96, 3, 2, 2, 2, 6, 103, 3, 2, 2, 2, 8, 105, 3, 2,
	2, 2, 10, 110, 3, 2, 2, 2, 12, 135, 3, 2, 2, 2, 14, 137, 3, 2, 2, 2, 16,
	148, 3, 2, 2, 2, 18, 161, 3, 2, 2, 2, 20, 175, 3, 2, 2, 2, 22, 177, 3,
	2, 2, 2, 24, 181, 3, 2, 2, 2, 26, 202, 3, 2, 2, 2, 28, 207, 3, 2, 2, 2,
	30, 218, 3, 2, 2, 2, 32, 231, 3, 2, 2, 2, 34, 249, 3, 2, 2, 2, 36, 261,
	3, 2, 2, 2, 38, 263, 3, 2, 2, 2, 40, 293, 3, 2, 2, 2, 42, 301, 3, 2, 2,
	2, 44, 304, 3, 2, 2, 2, 46, 310, 3, 2, 2, 2, 48, 312, 3, 2, 2, 2, 50, 317,
	3, 2, 2, 2, 52, 322, 3, 2, 2, 2, 54, 324, 3, 2, 2, 2, 56, 327, 3, 2, 2,
	2, 58, 334, 3, 2, 2, 2, 60, 336, 3, 2, 2, 2, 62, 342, 3, 2, 2, 2, 64, 344,
	3, 2, 2, 2, 66, 348, 3, 2, 2, 2, 68, 353, 3, 2, 2, 2, 70, 355, 3, 2, 2,
	2, 72, 361, 3, 2, 2, 2, 74, 368, 3, 2, 2, 2, 76, 78, 7, 35, 2, 2, 77, 76,
	3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 81, 5, 6, 4, 2,
	80, 79, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 84, 7,
	35, 2, 2, 83, 82, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85,
	92, 7, 37, 2, 2, 86, 88, 7, 35, 2, 2, 87, 86, 3, 2, 2, 2, 87, 88, 3, 2,
	2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 7, 36, 2, 2, 90, 92, 7, 37, 2, 2, 91,
	77, 3, 2, 2, 2, 91, 87, 3, 2, 2, 2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2, 2,
	2, 93, 94, 3, 2, 2, 2, 94, 3, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 97, 5,
	2, 2, 2, 97, 5, 3, 2, 2, 2, 98, 104, 5, 34, 18, 2, 99, 104, 5, 8, 5, 2,
	100, 104, 5, 14, 8, 2, 101, 104, 5, 40, 21, 2, 102, 104, 5, 10, 6, 2, 103,
	98, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2, 103, 100, 3, 2, 2, 2, 103, 101, 3,
	2, 2, 2, 103, 102, 3, 2, 2, 2, 104, 7, 3, 2, 2, 2, 105, 108, 7, 8, 2, 2,
	106, 107, 7, 14, 2, 2, 107, 109, 5, 32, 17, 2, 108, 106, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 9, 3, 2, 2, 2, 110, 111, 7, 11, 2, 2, 111, 112, 7,
	14, 2, 2, 112, 113, 7, 28, 2, 2, 113, 114, 5, 32, 17, 2, 114, 115, 7, 29,
	2, 2, 115, 116, 7, 14, 2, 2, 116, 117, 7, 30, 2, 2, 117, 118, 7, 37, 2,
	2, 118, 119, 5, 2, 2, 2, 119, 120, 7, 31, 2, 2, 120, 121, 5, 12, 7, 2,
	121, 11, 3, 2, 2, 2, 122, 123, 7, 14, 2, 2, 123, 124, 7, 12, 2, 2, 124,
	125, 7, 14, 2, 2, 125, 126, 7, 30, 2, 2, 126, 127, 7, 37, 2, 2, 127, 128,
	5, 2, 2, 2, 128, 129, 7, 31, 2, 2, 129, 136, 3, 2, 2, 2, 130, 131, 7, 14,
	2, 2, 131, 132, 7, 12, 2, 2, 132, 133, 7, 14, 2, 2, 133, 136, 5, 10, 6,
	2, 134, 136, 3, 2, 2, 2, 135, 122, 3, 2, 2, 2, 135, 130, 3, 2, 2, 2, 135,
	134, 3, 2, 2, 2, 136, 13, 3, 2, 2, 2, 137, 138, 7, 9, 2, 2, 138, 139, 7,
	14, 2, 2, 139, 140, 7, 32, 2, 2, 140, 141, 7, 13, 2, 2, 141, 142, 7, 14,
	2, 2, 142, 143, 5, 18, 10, 2, 143, 15, 3, 2, 2, 2, 144, 149, 7, 15, 2,
	2, 145, 149, 7, 16, 2, 2, 146, 149, 7, 18, 2, 2, 147, 149, 7, 17, 2, 2,
	148, 144, 3, 2, 2, 2, 148, 145, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148,
	147, 3, 2, 2, 2, 149, 17, 3, 2, 2, 2, 150, 162, 5, 16, 9, 2, 151, 152,
	5, 16, 9, 2, 152, 153, 7, 27, 2, 2, 153, 154, 5, 20, 11, 2, 154, 155, 7,
	64, 2, 2, 155, 162, 3, 2, 2, 2, 156, 157, 5, 16, 9, 2, 157, 158, 7, 27,
	2, 2, 158, 159, 5, 22, 12, 2, 159, 160, 7, 64, 2, 2, 160, 162, 3, 2, 2,
	2, 161, 150, 3, 2, 2, 2, 161, 151, 3, 2, 2, 2, 161, 156, 3, 2, 2, 2, 162,
	19, 3, 2, 2, 2, 163, 164, 7, 58, 2, 2, 164, 165, 7, 62, 2, 2, 165, 176,
	7, 61, 2, 2, 166, 167, 7, 56, 2, 2, 167, 168, 7, 62, 2, 2, 168, 176, 7,
	61, 2, 2, 169, 170, 7, 59, 2, 2, 170, 171, 7, 62, 2, 2, 171, 176, 7, 61,
	2, 2, 172, 173, 7, 57, 2, 2, 173, 174, 7, 62, 2, 2, 174, 176, 7, 61, 2,
	2, 175, 163, 3, 2, 2, 2, 175, 166, 3, 2, 2, 2, 175, 169, 3, 2, 2, 2, 175,
	172, 3, 2, 2, 2, 176, 21, 3, 2, 2, 2, 177, 178, 7, 63, 2, 2, 178, 179,
	5, 44, 23, 2, 179, 180, 7, 51, 2, 2, 180, 23, 3, 2, 2, 2, 181, 182, 7,
	10, 2, 2, 182, 183, 7, 14, 2, 2, 183, 184, 7, 32, 2, 2, 184, 191, 7, 28,
	2, 2, 185, 190, 5, 26, 14, 2, 186, 187, 5, 26, 14, 2, 187, 188, 7, 38,
	2, 2, 188, 190, 3, 2, 2, 2, 189, 185, 3, 2, 2, 2, 189, 186, 3, 2, 2, 2,
	190, 193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192,
	194, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 195, 7, 29, 2, 2, 195, 196,
	7, 13, 2, 2, 196, 197, 7, 14, 2, 2, 197, 198, 5, 18, 10, 2, 198, 199, 7,
	14, 2, 2, 199, 200, 7, 30, 2, 2, 200, 201, 7, 37, 2, 2, 201, 25, 3, 2,
	2, 2, 202, 203, 7, 32, 2, 2, 203, 204, 7, 13, 2, 2, 204, 205, 7, 14, 2,
	2, 205, 206, 5, 18, 10, 2, 206, 27, 3, 2, 2, 2, 207, 208, 5, 24, 13, 2,
	208, 209, 5, 4, 3, 2, 209, 210, 7, 31, 2, 2, 210, 29, 3, 2, 2, 2, 211,
	215, 5, 28, 15, 2, 212, 214, 7, 37, 2, 2, 213, 212, 3, 2, 2, 2, 214, 217,
	3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 219, 3, 2,
	2, 2, 217, 215, 3, 2, 2, 2, 218, 211, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2,
	220, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222,
	223, 7, 2, 2, 3, 223, 31, 3, 2, 2, 2, 224, 225, 8, 17, 1, 2, 225, 232,
	5, 36, 19, 2, 226, 232, 5, 40, 21, 2, 227, 228, 7, 28, 2, 2, 228, 229,
	5, 32, 17, 2, 229, 230, 7, 29, 2, 2, 230, 232, 3, 2, 2, 2, 231, 224, 3,
	2, 2, 2, 231, 226, 3, 2, 2, 2, 231, 227, 3, 2, 2, 2, 232, 246, 3, 2, 2,
	2, 233, 234, 12, 8, 2, 2, 234, 235, 9, 2, 2, 2, 235, 245, 5, 32, 17, 9,
	236, 237, 12, 7, 2, 2, 237, 238, 9, 3, 2, 2, 238, 245, 5, 32, 17, 8, 239,
	240, 12, 5, 2, 2, 240, 241, 7, 14, 2, 2, 241, 242, 9, 4, 2, 2, 242, 243,
	7, 14, 2, 2, 243, 245, 5, 32, 17, 6, 244, 233, 3, 2, 2, 2, 244, 236, 3,
	2, 2, 2, 244, 239, 3, 2, 2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2,
	2, 246, 247, 3, 2, 2, 2, 247, 33, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249,
	251, 7, 32, 2, 2, 250, 252, 7, 35, 2, 2, 251, 250, 3, 2, 2, 2, 251, 252,
	3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 254, 7, 26, 2, 2, 254, 255, 5, 32,
	17, 2, 255, 35, 3, 2, 2, 2, 256, 262, 7, 3, 2, 2, 257, 262, 7, 39, 2, 2,
	258, 262, 7, 19, 2, 2, 259, 262, 7, 20, 2, 2, 260, 262, 5, 38, 20, 2, 261,
	256, 3, 2, 2, 2, 261, 257, 3, 2, 2, 2, 261, 258, 3, 2, 2, 2, 261, 259,
	3, 2, 2, 2, 261, 260, 3, 2, 2, 2, 262, 37, 3, 2, 2, 2, 263, 264, 7, 32,
	2, 2, 264, 39, 3, 2, 2, 2, 265, 266, 7, 32, 2, 2, 266, 275, 7, 28, 2, 2,
	267, 272, 5, 32, 17, 2, 268, 269, 7, 38, 2, 2, 269, 271, 5, 32, 17, 2,
	270, 268, 3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272,
	273, 3, 2, 2, 2, 273, 276, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 267,
	3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 294, 7, 29,
	2, 2, 278, 279, 7, 33, 2, 2, 279, 280, 5, 42, 22, 2, 280, 289, 7, 28, 2,
	2, 281, 286, 5, 32, 17, 2, 282, 283, 7, 38, 2, 2, 283, 285, 5, 32, 17,
	2, 284, 282, 3, 2, 2, 2, 285, 288, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 286,
	287, 3, 2, 2, 2, 287, 290, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 289, 281,
	3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 292, 7, 29,
	2, 2, 292, 294, 3, 2, 2, 2, 293, 265, 3, 2, 2, 2, 293, 278, 3, 2, 2, 2,
	294, 41, 3, 2, 2, 2, 295, 297, 7, 32, 2, 2, 296, 298, 7, 34, 2, 2, 297,
	296, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 300, 3, 2, 2, 2, 299, 295,
	3, 2, 2, 2, 300, 303, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301, 302, 3, 2,
	2, 2, 302, 43, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 304, 305, 5, 48, 25, 2,
	305, 306, 5, 46, 24, 2, 306, 45, 3, 2, 2, 2, 307, 308, 7, 49, 2, 2, 308,
	311, 5, 44, 23, 2, 309, 311, 3, 2, 2, 2, 310, 307, 3, 2, 2, 2, 310, 309,
	3, 2, 2, 2, 311, 47, 3, 2, 2, 2, 312, 313, 5, 52, 27, 2, 313, 314, 5, 50,
	26, 2, 314, 49, 3, 2, 2, 2, 315, 318, 5, 48, 25, 2, 316, 318, 3, 2, 2,
	2, 317, 315, 3, 2, 2, 2, 317, 316, 3, 2, 2, 2, 318, 51, 3, 2, 2, 2, 319,
	323, 5, 54, 28, 2, 320, 323, 5, 56, 29, 2, 321, 323, 5, 58, 30, 2, 322,
	319, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 321, 3, 2, 2, 2, 323, 53, 3,
	2, 2, 2, 324, 325, 5, 58, 30, 2, 325, 326, 7, 47, 2, 2, 326, 55, 3, 2,
	2, 2, 327, 328, 5, 58, 30, 2, 328, 329, 7, 48, 2, 2, 329, 57, 3, 2, 2,
	2, 330, 335, 5, 60, 31, 2, 331, 335, 7, 46, 2, 2, 332, 335, 5, 70, 36,
	2, 333, 335, 5, 62, 32, 2, 334, 330, 3, 2, 2, 2, 334, 331, 3, 2, 2, 2,
	334, 332, 3, 2, 2, 2, 334, 333, 3, 2, 2, 2, 335, 59, 3, 2, 2, 2, 336, 337,
	7, 43, 2, 2, 337, 338, 5, 44, 23, 2, 338, 339, 7, 44, 2, 2, 339, 61, 3,
	2, 2, 2, 340, 343, 5, 64, 33, 2, 341, 343, 5, 66, 34, 2, 342, 340, 3, 2,
	2, 2, 342, 341, 3, 2, 2, 2, 343, 63, 3, 2, 2, 2, 344, 345, 7, 40, 2, 2,
	345, 346, 5, 72, 37, 2, 346, 347, 7, 55, 2, 2, 347, 65, 3, 2, 2, 2, 348,
	349, 7, 40, 2, 2, 349, 350, 7, 50, 2, 2, 350, 351, 5, 72, 37, 2, 351, 352,
	7, 55, 2, 2, 352, 67, 3, 2, 2, 2, 353, 354, 7, 54, 2, 2, 354, 69, 3, 2,
	2, 2, 355, 356, 9, 5, 2, 2, 356, 71, 3, 2, 2, 2, 357, 362, 5, 74, 38, 2,
	358, 359, 5, 74, 38, 2, 359, 360, 5, 72, 37, 2, 360, 362, 3, 2, 2, 2, 361,
	357, 3, 2, 2, 2, 361, 358, 3, 2, 2, 2, 362, 73, 3, 2, 2, 2, 363, 364, 5,
	68, 35, 2, 364, 365, 7, 53, 2, 2, 365, 366, 5, 68, 35, 2, 366, 369, 3,
	2, 2, 2, 367, 369, 5, 68, 35, 2, 368, 363, 3, 2, 2, 2, 368, 367, 3, 2,
	2, 2, 369, 75, 3, 2, 2, 2, 37, 77, 80, 83, 87, 91, 93, 103, 108, 135, 148,
	161, 175, 189, 191, 215, 220, 231, 244, 246, 251, 261, 272, 275, 286, 289,
	293, 297, 301, 310, 317, 322, 334, 342, 361, 368,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "'return'", "'var'", "'function'", "'if'", "'else'",
	"':'", "", "'uint'", "'string'", "'bool'", "'void'", "'true'", "'false'",
	"", "", "", "", "'=='", "", "", "", "", "'{'", "'}'", "", "'!'", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "'|'", "'^'",
}
var symbolicNames = []string{
	"", "INT", "MULTIPLY", "ADD", "SUBTRACT", "DIVIDE", "RETURN", "VAR", "FUNCTION",
	"IF", "ELSE", "COLON", "SPACE", "UINT_T", "STRING_T", "BOOL_T", "VOID_T",
	"TRUE_LIT", "FALSE_LIT", "GT_CONDITION", "GE_CONDITION", "LT_CONDITION",
	"LE_CONDITION", "EQ_CONDITION", "EQ_ASSIGNMENT", "BEGIN_CONSTRAINT", "BEGIN_GROUP",
	"END_GROUP", "BEGIN_BODY", "END_BODY", "IDENTIFIER", "JAVA_BEGIN", "JAVA_SEP",
	"WS", "COMMENT_LINE", "NEWLINE", "ARG_SEP", "STRING_LITERAL", "BEGIN_RE_RANGE",
	"END_RE_RANGE", "CHARACTER", "BEGIN_RE_GROUP", "END_RE_GROUP", "MINUS",
	"DOT", "STAR", "PLUS", "ALTERNATION", "RANGE_NEGATE", "RE_DELIMITER_CLOSE",
	"META_CHAR", "RANGE_SEPARATOR", "RANGE_CHARACTER", "RANGE_TERMINATE", "GT",
	"GE", "LT", "LE", "EQ", "CONSTRAINT_UINT", "CONSTRAINT_SPACE", "RE_DELIMITER_OPEN",
	"END_CONSTRAINT",
}

var ruleNames = []string{
	"body", "function_body", "body_line", "return_stmt", "if_stmt", "optional_else",
	"var_decl", "type_keyword", "type_specifier", "int_constraint", "string_constraint",
	"function_sig", "argument_decl", "function", "program", "expr", "var_assignment",
	"value_ref", "identifier_ref", "function_call", "java_call", "re", "union_prime",
	"simple_re", "concat_prime", "basic_re", "kleene_star", "plus", "elementary_re",
	"group", "range_re", "positive_range", "negative_range", "lax_character",
	"character", "range_items", "range_item",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PocLang struct {
	*antlr.BaseParser
}

func NewPocLang(input antlr.TokenStream) *PocLang {
	this := new(PocLang)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PocLang.g4"

	return this
}

// PocLang tokens.
const (
	PocLangEOF                = antlr.TokenEOF
	PocLangINT                = 1
	PocLangMULTIPLY           = 2
	PocLangADD                = 3
	PocLangSUBTRACT           = 4
	PocLangDIVIDE             = 5
	PocLangRETURN             = 6
	PocLangVAR                = 7
	PocLangFUNCTION           = 8
	PocLangIF                 = 9
	PocLangELSE               = 10
	PocLangCOLON              = 11
	PocLangSPACE              = 12
	PocLangUINT_T             = 13
	PocLangSTRING_T           = 14
	PocLangBOOL_T             = 15
	PocLangVOID_T             = 16
	PocLangTRUE_LIT           = 17
	PocLangFALSE_LIT          = 18
	PocLangGT_CONDITION       = 19
	PocLangGE_CONDITION       = 20
	PocLangLT_CONDITION       = 21
	PocLangLE_CONDITION       = 22
	PocLangEQ_CONDITION       = 23
	PocLangEQ_ASSIGNMENT      = 24
	PocLangBEGIN_CONSTRAINT   = 25
	PocLangBEGIN_GROUP        = 26
	PocLangEND_GROUP          = 27
	PocLangBEGIN_BODY         = 28
	PocLangEND_BODY           = 29
	PocLangIDENTIFIER         = 30
	PocLangJAVA_BEGIN         = 31
	PocLangJAVA_SEP           = 32
	PocLangWS                 = 33
	PocLangCOMMENT_LINE       = 34
	PocLangNEWLINE            = 35
	PocLangARG_SEP            = 36
	PocLangSTRING_LITERAL     = 37
	PocLangBEGIN_RE_RANGE     = 38
	PocLangEND_RE_RANGE       = 39
	PocLangCHARACTER          = 40
	PocLangBEGIN_RE_GROUP     = 41
	PocLangEND_RE_GROUP       = 42
	PocLangMINUS              = 43
	PocLangDOT                = 44
	PocLangSTAR               = 45
	PocLangPLUS               = 46
	PocLangALTERNATION        = 47
	PocLangRANGE_NEGATE       = 48
	PocLangRE_DELIMITER_CLOSE = 49
	PocLangMETA_CHAR          = 50
	PocLangRANGE_SEPARATOR    = 51
	PocLangRANGE_CHARACTER    = 52
	PocLangRANGE_TERMINATE    = 53
	PocLangGT                 = 54
	PocLangGE                 = 55
	PocLangLT                 = 56
	PocLangLE                 = 57
	PocLangEQ                 = 58
	PocLangCONSTRAINT_UINT    = 59
	PocLangCONSTRAINT_SPACE   = 60
	PocLangRE_DELIMITER_OPEN  = 61
	PocLangEND_CONSTRAINT     = 62
)

// PocLang rules.
const (
	PocLangRULE_body              = 0
	PocLangRULE_function_body     = 1
	PocLangRULE_body_line         = 2
	PocLangRULE_return_stmt       = 3
	PocLangRULE_if_stmt           = 4
	PocLangRULE_optional_else     = 5
	PocLangRULE_var_decl          = 6
	PocLangRULE_type_keyword      = 7
	PocLangRULE_type_specifier    = 8
	PocLangRULE_int_constraint    = 9
	PocLangRULE_string_constraint = 10
	PocLangRULE_function_sig      = 11
	PocLangRULE_argument_decl     = 12
	PocLangRULE_function          = 13
	PocLangRULE_program           = 14
	PocLangRULE_expr              = 15
	PocLangRULE_var_assignment    = 16
	PocLangRULE_value_ref         = 17
	PocLangRULE_identifier_ref    = 18
	PocLangRULE_function_call     = 19
	PocLangRULE_java_call         = 20
	PocLangRULE_re                = 21
	PocLangRULE_union_prime       = 22
	PocLangRULE_simple_re         = 23
	PocLangRULE_concat_prime      = 24
	PocLangRULE_basic_re          = 25
	PocLangRULE_kleene_star       = 26
	PocLangRULE_plus              = 27
	PocLangRULE_elementary_re     = 28
	PocLangRULE_group             = 29
	PocLangRULE_range_re          = 30
	PocLangRULE_positive_range    = 31
	PocLangRULE_negative_range    = 32
	PocLangRULE_lax_character     = 33
	PocLangRULE_character         = 34
	PocLangRULE_range_items       = 35
	PocLangRULE_range_item        = 36
)

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(PocLangNEWLINE)
}

func (s *BodyContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(PocLangNEWLINE, i)
}

func (s *BodyContext) AllCOMMENT_LINE() []antlr.TerminalNode {
	return s.GetTokens(PocLangCOMMENT_LINE)
}

func (s *BodyContext) COMMENT_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PocLangCOMMENT_LINE, i)
}

func (s *BodyContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(PocLangWS)
}

func (s *BodyContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(PocLangWS, i)
}

func (s *BodyContext) AllBody_line() []IBody_lineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBody_lineContext)(nil)).Elem())
	var tst = make([]IBody_lineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBody_lineContext)
		}
	}

	return tst
}

func (s *BodyContext) Body_line(i int) IBody_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBody_lineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBody_lineContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PocLangRULE_body)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(PocLangRETURN-6))|(1<<(PocLangVAR-6))|(1<<(PocLangIF-6))|(1<<(PocLangIDENTIFIER-6))|(1<<(PocLangJAVA_BEGIN-6))|(1<<(PocLangWS-6))|(1<<(PocLangCOMMENT_LINE-6))|(1<<(PocLangNEWLINE-6)))) != 0 {
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			p.SetState(75)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(74)
					p.Match(PocLangWS)
				}

			}
			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PocLangRETURN)|(1<<PocLangVAR)|(1<<PocLangIF)|(1<<PocLangIDENTIFIER)|(1<<PocLangJAVA_BEGIN))) != 0 {
				{
					p.SetState(77)
					p.Body_line()
				}

			}
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == PocLangWS {
				{
					p.SetState(80)
					p.Match(PocLangWS)
				}

			}
			{
				p.SetState(83)
				p.Match(PocLangNEWLINE)
			}

		case 2:
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == PocLangWS {
				{
					p.SetState(84)
					p.Match(PocLangWS)
				}

			}
			{
				p.SetState(87)
				p.Match(PocLangCOMMENT_LINE)
			}
			{
				p.SetState(88)
				p.Match(PocLangNEWLINE)
			}

		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunction_bodyContext is an interface to support dynamic dispatch.
type IFunction_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_bodyContext differentiates from other interfaces.
	IsFunction_bodyContext()
}

type Function_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_bodyContext() *Function_bodyContext {
	var p = new(Function_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_function_body
	return p
}

func (*Function_bodyContext) IsFunction_bodyContext() {}

func NewFunction_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_bodyContext {
	var p = new(Function_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_function_body

	return p
}

func (s *Function_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_bodyContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *Function_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterFunction_body(s)
	}
}

func (s *Function_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitFunction_body(s)
	}
}

func (s *Function_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitFunction_body(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Function_body() (localctx IFunction_bodyContext) {
	localctx = NewFunction_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PocLangRULE_function_body)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Body()
	}

	return localctx
}

// IBody_lineContext is an interface to support dynamic dispatch.
type IBody_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBody_lineContext differentiates from other interfaces.
	IsBody_lineContext()
}

type Body_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBody_lineContext() *Body_lineContext {
	var p = new(Body_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_body_line
	return p
}

func (*Body_lineContext) IsBody_lineContext() {}

func NewBody_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Body_lineContext {
	var p = new(Body_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_body_line

	return p
}

func (s *Body_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Body_lineContext) Var_assignment() IVar_assignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_assignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_assignmentContext)
}

func (s *Body_lineContext) Return_stmt() IReturn_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *Body_lineContext) Var_decl() IVar_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Body_lineContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Body_lineContext) If_stmt() IIf_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *Body_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Body_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Body_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterBody_line(s)
	}
}

func (s *Body_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitBody_line(s)
	}
}

func (s *Body_lineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitBody_line(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Body_line() (localctx IBody_lineContext) {
	localctx = NewBody_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PocLangRULE_body_line)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(96)
			p.Var_assignment()
		}

	case 2:
		{
			p.SetState(97)
			p.Return_stmt()
		}

	case 3:
		{
			p.SetState(98)
			p.Var_decl()
		}

	case 4:
		{
			p.SetState(99)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(100)
			p.If_stmt()
		}

	}

	return localctx
}

// IReturn_stmtContext is an interface to support dynamic dispatch.
type IReturn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturn_stmtContext differentiates from other interfaces.
	IsReturn_stmtContext()
}

type Return_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_stmtContext() *Return_stmtContext {
	var p = new(Return_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_return_stmt
	return p
}

func (*Return_stmtContext) IsReturn_stmtContext() {}

func NewReturn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_stmtContext {
	var p = new(Return_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_return_stmt

	return p
}

func (s *Return_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_stmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(PocLangRETURN, 0)
}

func (s *Return_stmtContext) SPACE() antlr.TerminalNode {
	return s.GetToken(PocLangSPACE, 0)
}

func (s *Return_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterReturn_stmt(s)
	}
}

func (s *Return_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitReturn_stmt(s)
	}
}

func (s *Return_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitReturn_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Return_stmt() (localctx IReturn_stmtContext) {
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PocLangRULE_return_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(PocLangRETURN)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PocLangSPACE {
		{
			p.SetState(104)
			p.Match(PocLangSPACE)
		}
		{
			p.SetState(105)
			p.expr(0)
		}

	}

	return localctx
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_if_stmt
	return p
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) IF() antlr.TerminalNode {
	return s.GetToken(PocLangIF, 0)
}

func (s *If_stmtContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(PocLangSPACE)
}

func (s *If_stmtContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(PocLangSPACE, i)
}

func (s *If_stmtContext) BEGIN_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_GROUP, 0)
}

func (s *If_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_stmtContext) END_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangEND_GROUP, 0)
}

func (s *If_stmtContext) BEGIN_BODY() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_BODY, 0)
}

func (s *If_stmtContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(PocLangNEWLINE, 0)
}

func (s *If_stmtContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *If_stmtContext) END_BODY() antlr.TerminalNode {
	return s.GetToken(PocLangEND_BODY, 0)
}

func (s *If_stmtContext) Optional_else() IOptional_elseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_elseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_elseContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (s *If_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitIf_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PocLangRULE_if_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(PocLangIF)
	}
	{
		p.SetState(109)
		p.Match(PocLangSPACE)
	}
	{
		p.SetState(110)
		p.Match(PocLangBEGIN_GROUP)
	}
	{
		p.SetState(111)
		p.expr(0)
	}
	{
		p.SetState(112)
		p.Match(PocLangEND_GROUP)
	}
	{
		p.SetState(113)
		p.Match(PocLangSPACE)
	}
	{
		p.SetState(114)
		p.Match(PocLangBEGIN_BODY)
	}
	{
		p.SetState(115)
		p.Match(PocLangNEWLINE)
	}
	{
		p.SetState(116)
		p.Body()
	}
	{
		p.SetState(117)
		p.Match(PocLangEND_BODY)
	}
	{
		p.SetState(118)
		p.Optional_else()
	}

	return localctx
}

// IOptional_elseContext is an interface to support dynamic dispatch.
type IOptional_elseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptional_elseContext differentiates from other interfaces.
	IsOptional_elseContext()
}

type Optional_elseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptional_elseContext() *Optional_elseContext {
	var p = new(Optional_elseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_optional_else
	return p
}

func (*Optional_elseContext) IsOptional_elseContext() {}

func NewOptional_elseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Optional_elseContext {
	var p = new(Optional_elseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_optional_else

	return p
}

func (s *Optional_elseContext) GetParser() antlr.Parser { return s.parser }

func (s *Optional_elseContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(PocLangSPACE)
}

func (s *Optional_elseContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(PocLangSPACE, i)
}

func (s *Optional_elseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PocLangELSE, 0)
}

func (s *Optional_elseContext) BEGIN_BODY() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_BODY, 0)
}

func (s *Optional_elseContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(PocLangNEWLINE, 0)
}

func (s *Optional_elseContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *Optional_elseContext) END_BODY() antlr.TerminalNode {
	return s.GetToken(PocLangEND_BODY, 0)
}

func (s *Optional_elseContext) If_stmt() IIf_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *Optional_elseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Optional_elseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Optional_elseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterOptional_else(s)
	}
}

func (s *Optional_elseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitOptional_else(s)
	}
}

func (s *Optional_elseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitOptional_else(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Optional_else() (localctx IOptional_elseContext) {
	localctx = NewOptional_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PocLangRULE_optional_else)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(PocLangSPACE)
		}
		{
			p.SetState(121)
			p.Match(PocLangELSE)
		}
		{
			p.SetState(122)
			p.Match(PocLangSPACE)
		}
		{
			p.SetState(123)
			p.Match(PocLangBEGIN_BODY)
		}
		{
			p.SetState(124)
			p.Match(PocLangNEWLINE)
		}
		{
			p.SetState(125)
			p.Body()
		}
		{
			p.SetState(126)
			p.Match(PocLangEND_BODY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Match(PocLangSPACE)
		}
		{
			p.SetState(129)
			p.Match(PocLangELSE)
		}
		{
			p.SetState(130)
			p.Match(PocLangSPACE)
		}
		{
			p.SetState(131)
			p.If_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	}

	return localctx
}

// IVar_declContext is an interface to support dynamic dispatch.
type IVar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_declContext differentiates from other interfaces.
	IsVar_declContext()
}

type Var_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declContext() *Var_declContext {
	var p = new(Var_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_var_decl
	return p
}

func (*Var_declContext) IsVar_declContext() {}

func NewVar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declContext {
	var p = new(Var_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_var_decl

	return p
}

func (s *Var_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declContext) VAR() antlr.TerminalNode {
	return s.GetToken(PocLangVAR, 0)
}

func (s *Var_declContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(PocLangSPACE)
}

func (s *Var_declContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(PocLangSPACE, i)
}

func (s *Var_declContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PocLangIDENTIFIER, 0)
}

func (s *Var_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(PocLangCOLON, 0)
}

func (s *Var_declContext) Type_specifier() IType_specifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_specifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_specifierContext)
}

func (s *Var_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterVar_decl(s)
	}
}

func (s *Var_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitVar_decl(s)
	}
}

func (s *Var_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitVar_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Var_decl() (localctx IVar_declContext) {
	localctx = NewVar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PocLangRULE_var_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(PocLangVAR)
	}
	{
		p.SetState(136)
		p.Match(PocLangSPACE)
	}
	{
		p.SetState(137)
		p.Match(PocLangIDENTIFIER)
	}
	{
		p.SetState(138)
		p.Match(PocLangCOLON)
	}
	{
		p.SetState(139)
		p.Match(PocLangSPACE)
	}
	{
		p.SetState(140)
		p.Type_specifier()
	}

	return localctx
}

// IType_keywordContext is an interface to support dynamic dispatch.
type IType_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_keywordContext differentiates from other interfaces.
	IsType_keywordContext()
}

type Type_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_keywordContext() *Type_keywordContext {
	var p = new(Type_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_type_keyword
	return p
}

func (*Type_keywordContext) IsType_keywordContext() {}

func NewType_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_keywordContext {
	var p = new(Type_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_type_keyword

	return p
}

func (s *Type_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_keywordContext) CopyFrom(ctx *Type_keywordContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Type_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VoidTypeContext struct {
	*Type_keywordContext
}

func NewVoidTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidTypeContext {
	var p = new(VoidTypeContext)

	p.Type_keywordContext = NewEmptyType_keywordContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_keywordContext))

	return p
}

func (s *VoidTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidTypeContext) VOID_T() antlr.TerminalNode {
	return s.GetToken(PocLangVOID_T, 0)
}

func (s *VoidTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterVoidType(s)
	}
}

func (s *VoidTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitVoidType(s)
	}
}

func (s *VoidTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitVoidType(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolTypeContext struct {
	*Type_keywordContext
}

func NewBoolTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolTypeContext {
	var p = new(BoolTypeContext)

	p.Type_keywordContext = NewEmptyType_keywordContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_keywordContext))

	return p
}

func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTypeContext) BOOL_T() antlr.TerminalNode {
	return s.GetToken(PocLangBOOL_T, 0)
}

func (s *BoolTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterBoolType(s)
	}
}

func (s *BoolTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitBoolType(s)
	}
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringTypeContext struct {
	*Type_keywordContext
}

func NewStringTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringTypeContext {
	var p = new(StringTypeContext)

	p.Type_keywordContext = NewEmptyType_keywordContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_keywordContext))

	return p
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) STRING_T() antlr.TerminalNode {
	return s.GetToken(PocLangSTRING_T, 0)
}

func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnsignedIntTypeContext struct {
	*Type_keywordContext
}

func NewUnsignedIntTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnsignedIntTypeContext {
	var p = new(UnsignedIntTypeContext)

	p.Type_keywordContext = NewEmptyType_keywordContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_keywordContext))

	return p
}

func (s *UnsignedIntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnsignedIntTypeContext) UINT_T() antlr.TerminalNode {
	return s.GetToken(PocLangUINT_T, 0)
}

func (s *UnsignedIntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterUnsignedIntType(s)
	}
}

func (s *UnsignedIntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitUnsignedIntType(s)
	}
}

func (s *UnsignedIntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitUnsignedIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Type_keyword() (localctx IType_keywordContext) {
	localctx = NewType_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PocLangRULE_type_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PocLangUINT_T:
		localctx = NewUnsignedIntTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(PocLangUINT_T)
		}

	case PocLangSTRING_T:
		localctx = NewStringTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Match(PocLangSTRING_T)
		}

	case PocLangVOID_T:
		localctx = NewVoidTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.Match(PocLangVOID_T)
		}

	case PocLangBOOL_T:
		localctx = NewBoolTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(145)
			p.Match(PocLangBOOL_T)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IType_specifierContext is an interface to support dynamic dispatch.
type IType_specifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_specifierContext differentiates from other interfaces.
	IsType_specifierContext()
}

type Type_specifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_specifierContext() *Type_specifierContext {
	var p = new(Type_specifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_type_specifier
	return p
}

func (*Type_specifierContext) IsType_specifierContext() {}

func NewType_specifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_specifierContext {
	var p = new(Type_specifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_type_specifier

	return p
}

func (s *Type_specifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_specifierContext) Type_keyword() IType_keywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_keywordContext)
}

func (s *Type_specifierContext) BEGIN_CONSTRAINT() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_CONSTRAINT, 0)
}

func (s *Type_specifierContext) Int_constraint() IInt_constraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt_constraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt_constraintContext)
}

func (s *Type_specifierContext) END_CONSTRAINT() antlr.TerminalNode {
	return s.GetToken(PocLangEND_CONSTRAINT, 0)
}

func (s *Type_specifierContext) String_constraint() IString_constraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_constraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_constraintContext)
}

func (s *Type_specifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_specifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_specifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterType_specifier(s)
	}
}

func (s *Type_specifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitType_specifier(s)
	}
}

func (s *Type_specifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitType_specifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Type_specifier() (localctx IType_specifierContext) {
	localctx = NewType_specifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PocLangRULE_type_specifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Type_keyword()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Type_keyword()
		}
		{
			p.SetState(150)
			p.Match(PocLangBEGIN_CONSTRAINT)
		}
		{
			p.SetState(151)
			p.Int_constraint()
		}
		{
			p.SetState(152)
			p.Match(PocLangEND_CONSTRAINT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Type_keyword()
		}
		{
			p.SetState(155)
			p.Match(PocLangBEGIN_CONSTRAINT)
		}
		{
			p.SetState(156)
			p.String_constraint()
		}
		{
			p.SetState(157)
			p.Match(PocLangEND_CONSTRAINT)
		}

	}

	return localctx
}

// IInt_constraintContext is an interface to support dynamic dispatch.
type IInt_constraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInt_constraintContext differentiates from other interfaces.
	IsInt_constraintContext()
}

type Int_constraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt_constraintContext() *Int_constraintContext {
	var p = new(Int_constraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_int_constraint
	return p
}

func (*Int_constraintContext) IsInt_constraintContext() {}

func NewInt_constraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int_constraintContext {
	var p = new(Int_constraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_int_constraint

	return p
}

func (s *Int_constraintContext) GetParser() antlr.Parser { return s.parser }

func (s *Int_constraintContext) CopyFrom(ctx *Int_constraintContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Int_constraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int_constraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GreaterThanEqualsConstraintContext struct {
	*Int_constraintContext
}

func NewGreaterThanEqualsConstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanEqualsConstraintContext {
	var p = new(GreaterThanEqualsConstraintContext)

	p.Int_constraintContext = NewEmptyInt_constraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Int_constraintContext))

	return p
}

func (s *GreaterThanEqualsConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanEqualsConstraintContext) GE() antlr.TerminalNode {
	return s.GetToken(PocLangGE, 0)
}

func (s *GreaterThanEqualsConstraintContext) CONSTRAINT_SPACE() antlr.TerminalNode {
	return s.GetToken(PocLangCONSTRAINT_SPACE, 0)
}

func (s *GreaterThanEqualsConstraintContext) CONSTRAINT_UINT() antlr.TerminalNode {
	return s.GetToken(PocLangCONSTRAINT_UINT, 0)
}

func (s *GreaterThanEqualsConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterGreaterThanEqualsConstraint(s)
	}
}

func (s *GreaterThanEqualsConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitGreaterThanEqualsConstraint(s)
	}
}

func (s *GreaterThanEqualsConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitGreaterThanEqualsConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterThanConstraintContext struct {
	*Int_constraintContext
}

func NewGreaterThanConstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanConstraintContext {
	var p = new(GreaterThanConstraintContext)

	p.Int_constraintContext = NewEmptyInt_constraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Int_constraintContext))

	return p
}

func (s *GreaterThanConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanConstraintContext) GT() antlr.TerminalNode {
	return s.GetToken(PocLangGT, 0)
}

func (s *GreaterThanConstraintContext) CONSTRAINT_SPACE() antlr.TerminalNode {
	return s.GetToken(PocLangCONSTRAINT_SPACE, 0)
}

func (s *GreaterThanConstraintContext) CONSTRAINT_UINT() antlr.TerminalNode {
	return s.GetToken(PocLangCONSTRAINT_UINT, 0)
}

func (s *GreaterThanConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterGreaterThanConstraint(s)
	}
}

func (s *GreaterThanConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitGreaterThanConstraint(s)
	}
}

func (s *GreaterThanConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitGreaterThanConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessThanEqualsConstraintContext struct {
	*Int_constraintContext
}

func NewLessThanEqualsConstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanEqualsConstraintContext {
	var p = new(LessThanEqualsConstraintContext)

	p.Int_constraintContext = NewEmptyInt_constraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Int_constraintContext))

	return p
}

func (s *LessThanEqualsConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanEqualsConstraintContext) LE() antlr.TerminalNode {
	return s.GetToken(PocLangLE, 0)
}

func (s *LessThanEqualsConstraintContext) CONSTRAINT_SPACE() antlr.TerminalNode {
	return s.GetToken(PocLangCONSTRAINT_SPACE, 0)
}

func (s *LessThanEqualsConstraintContext) CONSTRAINT_UINT() antlr.TerminalNode {
	return s.GetToken(PocLangCONSTRAINT_UINT, 0)
}

func (s *LessThanEqualsConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterLessThanEqualsConstraint(s)
	}
}

func (s *LessThanEqualsConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitLessThanEqualsConstraint(s)
	}
}

func (s *LessThanEqualsConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitLessThanEqualsConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessThanConstraintContext struct {
	*Int_constraintContext
}

func NewLessThanConstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanConstraintContext {
	var p = new(LessThanConstraintContext)

	p.Int_constraintContext = NewEmptyInt_constraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Int_constraintContext))

	return p
}

func (s *LessThanConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanConstraintContext) LT() antlr.TerminalNode {
	return s.GetToken(PocLangLT, 0)
}

func (s *LessThanConstraintContext) CONSTRAINT_SPACE() antlr.TerminalNode {
	return s.GetToken(PocLangCONSTRAINT_SPACE, 0)
}

func (s *LessThanConstraintContext) CONSTRAINT_UINT() antlr.TerminalNode {
	return s.GetToken(PocLangCONSTRAINT_UINT, 0)
}

func (s *LessThanConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterLessThanConstraint(s)
	}
}

func (s *LessThanConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitLessThanConstraint(s)
	}
}

func (s *LessThanConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitLessThanConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Int_constraint() (localctx IInt_constraintContext) {
	localctx = NewInt_constraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PocLangRULE_int_constraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PocLangLT:
		localctx = NewLessThanConstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(PocLangLT)
		}
		{
			p.SetState(162)
			p.Match(PocLangCONSTRAINT_SPACE)
		}
		{
			p.SetState(163)
			p.Match(PocLangCONSTRAINT_UINT)
		}

	case PocLangGT:
		localctx = NewGreaterThanConstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(PocLangGT)
		}
		{
			p.SetState(165)
			p.Match(PocLangCONSTRAINT_SPACE)
		}
		{
			p.SetState(166)
			p.Match(PocLangCONSTRAINT_UINT)
		}

	case PocLangLE:
		localctx = NewLessThanEqualsConstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.Match(PocLangLE)
		}
		{
			p.SetState(168)
			p.Match(PocLangCONSTRAINT_SPACE)
		}
		{
			p.SetState(169)
			p.Match(PocLangCONSTRAINT_UINT)
		}

	case PocLangGE:
		localctx = NewGreaterThanEqualsConstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(170)
			p.Match(PocLangGE)
		}
		{
			p.SetState(171)
			p.Match(PocLangCONSTRAINT_SPACE)
		}
		{
			p.SetState(172)
			p.Match(PocLangCONSTRAINT_UINT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IString_constraintContext is an interface to support dynamic dispatch.
type IString_constraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_constraintContext differentiates from other interfaces.
	IsString_constraintContext()
}

type String_constraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_constraintContext() *String_constraintContext {
	var p = new(String_constraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_string_constraint
	return p
}

func (*String_constraintContext) IsString_constraintContext() {}

func NewString_constraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_constraintContext {
	var p = new(String_constraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_string_constraint

	return p
}

func (s *String_constraintContext) GetParser() antlr.Parser { return s.parser }

func (s *String_constraintContext) RE_DELIMITER_OPEN() antlr.TerminalNode {
	return s.GetToken(PocLangRE_DELIMITER_OPEN, 0)
}

func (s *String_constraintContext) Re() IReContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReContext)
}

func (s *String_constraintContext) RE_DELIMITER_CLOSE() antlr.TerminalNode {
	return s.GetToken(PocLangRE_DELIMITER_CLOSE, 0)
}

func (s *String_constraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_constraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_constraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterString_constraint(s)
	}
}

func (s *String_constraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitString_constraint(s)
	}
}

func (s *String_constraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitString_constraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) String_constraint() (localctx IString_constraintContext) {
	localctx = NewString_constraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PocLangRULE_string_constraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(PocLangRE_DELIMITER_OPEN)
	}
	{
		p.SetState(176)
		p.Re()
	}
	{
		p.SetState(177)
		p.Match(PocLangRE_DELIMITER_CLOSE)
	}

	return localctx
}

// IFunction_sigContext is an interface to support dynamic dispatch.
type IFunction_sigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_sigContext differentiates from other interfaces.
	IsFunction_sigContext()
}

type Function_sigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_sigContext() *Function_sigContext {
	var p = new(Function_sigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_function_sig
	return p
}

func (*Function_sigContext) IsFunction_sigContext() {}

func NewFunction_sigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_sigContext {
	var p = new(Function_sigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_function_sig

	return p
}

func (s *Function_sigContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_sigContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PocLangFUNCTION, 0)
}

func (s *Function_sigContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(PocLangSPACE)
}

func (s *Function_sigContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(PocLangSPACE, i)
}

func (s *Function_sigContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PocLangIDENTIFIER, 0)
}

func (s *Function_sigContext) BEGIN_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_GROUP, 0)
}

func (s *Function_sigContext) END_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangEND_GROUP, 0)
}

func (s *Function_sigContext) COLON() antlr.TerminalNode {
	return s.GetToken(PocLangCOLON, 0)
}

func (s *Function_sigContext) Type_specifier() IType_specifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_specifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_specifierContext)
}

func (s *Function_sigContext) BEGIN_BODY() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_BODY, 0)
}

func (s *Function_sigContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(PocLangNEWLINE, 0)
}

func (s *Function_sigContext) AllArgument_decl() []IArgument_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgument_declContext)(nil)).Elem())
	var tst = make([]IArgument_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgument_declContext)
		}
	}

	return tst
}

func (s *Function_sigContext) Argument_decl(i int) IArgument_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgument_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgument_declContext)
}

func (s *Function_sigContext) AllARG_SEP() []antlr.TerminalNode {
	return s.GetTokens(PocLangARG_SEP)
}

func (s *Function_sigContext) ARG_SEP(i int) antlr.TerminalNode {
	return s.GetToken(PocLangARG_SEP, i)
}

func (s *Function_sigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_sigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_sigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterFunction_sig(s)
	}
}

func (s *Function_sigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitFunction_sig(s)
	}
}

func (s *Function_sigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitFunction_sig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Function_sig() (localctx IFunction_sigContext) {
	localctx = NewFunction_sigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PocLangRULE_function_sig)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(PocLangFUNCTION)
	}
	{
		p.SetState(180)
		p.Match(PocLangSPACE)
	}
	{
		p.SetState(181)
		p.Match(PocLangIDENTIFIER)
	}
	{
		p.SetState(182)
		p.Match(PocLangBEGIN_GROUP)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PocLangIDENTIFIER {
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(183)
				p.Argument_decl()
			}

		case 2:
			{
				p.SetState(184)
				p.Argument_decl()
			}
			{
				p.SetState(185)
				p.Match(PocLangARG_SEP)
			}

		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(192)
		p.Match(PocLangEND_GROUP)
	}
	{
		p.SetState(193)
		p.Match(PocLangCOLON)
	}
	{
		p.SetState(194)
		p.Match(PocLangSPACE)
	}
	{
		p.SetState(195)
		p.Type_specifier()
	}
	{
		p.SetState(196)
		p.Match(PocLangSPACE)
	}
	{
		p.SetState(197)
		p.Match(PocLangBEGIN_BODY)
	}
	{
		p.SetState(198)
		p.Match(PocLangNEWLINE)
	}

	return localctx
}

// IArgument_declContext is an interface to support dynamic dispatch.
type IArgument_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgument_declContext differentiates from other interfaces.
	IsArgument_declContext()
}

type Argument_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgument_declContext() *Argument_declContext {
	var p = new(Argument_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_argument_decl
	return p
}

func (*Argument_declContext) IsArgument_declContext() {}

func NewArgument_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Argument_declContext {
	var p = new(Argument_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_argument_decl

	return p
}

func (s *Argument_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Argument_declContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PocLangIDENTIFIER, 0)
}

func (s *Argument_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(PocLangCOLON, 0)
}

func (s *Argument_declContext) SPACE() antlr.TerminalNode {
	return s.GetToken(PocLangSPACE, 0)
}

func (s *Argument_declContext) Type_specifier() IType_specifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_specifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_specifierContext)
}

func (s *Argument_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Argument_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Argument_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterArgument_decl(s)
	}
}

func (s *Argument_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitArgument_decl(s)
	}
}

func (s *Argument_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitArgument_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Argument_decl() (localctx IArgument_declContext) {
	localctx = NewArgument_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PocLangRULE_argument_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(PocLangIDENTIFIER)
	}
	{
		p.SetState(201)
		p.Match(PocLangCOLON)
	}
	{
		p.SetState(202)
		p.Match(PocLangSPACE)
	}
	{
		p.SetState(203)
		p.Type_specifier()
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Function_sig() IFunction_sigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_sigContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_sigContext)
}

func (s *FunctionContext) Function_body() IFunction_bodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_bodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_bodyContext)
}

func (s *FunctionContext) END_BODY() antlr.TerminalNode {
	return s.GetToken(PocLangEND_BODY, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PocLangRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Function_sig()
	}
	{
		p.SetState(206)
		p.Function_body()
	}
	{
		p.SetState(207)
		p.Match(PocLangEND_BODY)
	}

	return localctx
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(PocLangEOF, 0)
}

func (s *ProgramContext) AllFunction() []IFunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionContext)(nil)).Elem())
	var tst = make([]IFunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Function(i int) IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ProgramContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(PocLangNEWLINE)
}

func (s *ProgramContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(PocLangNEWLINE, i)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PocLangRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PocLangFUNCTION {
		{
			p.SetState(209)
			p.Function()
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PocLangNEWLINE {
			{
				p.SetState(210)
				p.Match(PocLangNEWLINE)
			}

			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(220)
		p.Match(PocLangEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Value_ref() IValue_refContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_refContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValue_refContext)
}

func (s *ExprContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ExprContext) BEGIN_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_GROUP, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) END_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangEND_GROUP, 0)
}

func (s *ExprContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(PocLangMULTIPLY, 0)
}

func (s *ExprContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(PocLangDIVIDE, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(PocLangADD, 0)
}

func (s *ExprContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(PocLangSUBTRACT, 0)
}

func (s *ExprContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(PocLangSPACE)
}

func (s *ExprContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(PocLangSPACE, i)
}

func (s *ExprContext) LT_CONDITION() antlr.TerminalNode {
	return s.GetToken(PocLangLT_CONDITION, 0)
}

func (s *ExprContext) GT_CONDITION() antlr.TerminalNode {
	return s.GetToken(PocLangGT_CONDITION, 0)
}

func (s *ExprContext) LE_CONDITION() antlr.TerminalNode {
	return s.GetToken(PocLangLE_CONDITION, 0)
}

func (s *ExprContext) GE_CONDITION() antlr.TerminalNode {
	return s.GetToken(PocLangGE_CONDITION, 0)
}

func (s *ExprContext) EQ_CONDITION() antlr.TerminalNode {
	return s.GetToken(PocLangEQ_CONDITION, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PocLang) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, PocLangRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(223)
			p.Value_ref()
		}

	case 2:
		{
			p.SetState(224)
			p.Function_call()
		}

	case 3:
		{
			p.SetState(225)
			p.Match(PocLangBEGIN_GROUP)
		}
		{
			p.SetState(226)
			p.expr(0)
		}
		{
			p.SetState(227)
			p.Match(PocLangEND_GROUP)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PocLangRULE_expr)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(232)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PocLangMULTIPLY || _la == PocLangDIVIDE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(233)
					p.expr(7)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PocLangRULE_expr)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(235)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PocLangADD || _la == PocLangSUBTRACT) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(236)
					p.expr(6)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PocLangRULE_expr)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(238)
					p.Match(PocLangSPACE)
				}
				{
					p.SetState(239)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PocLangGT_CONDITION)|(1<<PocLangGE_CONDITION)|(1<<PocLangLT_CONDITION)|(1<<PocLangLE_CONDITION)|(1<<PocLangEQ_CONDITION))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(240)
					p.Match(PocLangSPACE)
				}
				{
					p.SetState(241)
					p.expr(4)
				}

			}

		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IVar_assignmentContext is an interface to support dynamic dispatch.
type IVar_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_assignmentContext differentiates from other interfaces.
	IsVar_assignmentContext()
}

type Var_assignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_assignmentContext() *Var_assignmentContext {
	var p = new(Var_assignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_var_assignment
	return p
}

func (*Var_assignmentContext) IsVar_assignmentContext() {}

func NewVar_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignmentContext {
	var p = new(Var_assignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_var_assignment

	return p
}

func (s *Var_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PocLangIDENTIFIER, 0)
}

func (s *Var_assignmentContext) EQ_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(PocLangEQ_ASSIGNMENT, 0)
}

func (s *Var_assignmentContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Var_assignmentContext) WS() antlr.TerminalNode {
	return s.GetToken(PocLangWS, 0)
}

func (s *Var_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterVar_assignment(s)
	}
}

func (s *Var_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitVar_assignment(s)
	}
}

func (s *Var_assignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitVar_assignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Var_assignment() (localctx IVar_assignmentContext) {
	localctx = NewVar_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PocLangRULE_var_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(PocLangIDENTIFIER)
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PocLangWS {
		{
			p.SetState(248)
			p.Match(PocLangWS)
		}

	}
	{
		p.SetState(251)
		p.Match(PocLangEQ_ASSIGNMENT)
	}
	{
		p.SetState(252)
		p.expr(0)
	}

	return localctx
}

// IValue_refContext is an interface to support dynamic dispatch.
type IValue_refContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValue_refContext differentiates from other interfaces.
	IsValue_refContext()
}

type Value_refContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_refContext() *Value_refContext {
	var p = new(Value_refContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_value_ref
	return p
}

func (*Value_refContext) IsValue_refContext() {}

func NewValue_refContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_refContext {
	var p = new(Value_refContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_value_ref

	return p
}

func (s *Value_refContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_refContext) INT() antlr.TerminalNode {
	return s.GetToken(PocLangINT, 0)
}

func (s *Value_refContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PocLangSTRING_LITERAL, 0)
}

func (s *Value_refContext) TRUE_LIT() antlr.TerminalNode {
	return s.GetToken(PocLangTRUE_LIT, 0)
}

func (s *Value_refContext) FALSE_LIT() antlr.TerminalNode {
	return s.GetToken(PocLangFALSE_LIT, 0)
}

func (s *Value_refContext) Identifier_ref() IIdentifier_refContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_refContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_refContext)
}

func (s *Value_refContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_refContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_refContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterValue_ref(s)
	}
}

func (s *Value_refContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitValue_ref(s)
	}
}

func (s *Value_refContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitValue_ref(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Value_ref() (localctx IValue_refContext) {
	localctx = NewValue_refContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PocLangRULE_value_ref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PocLangINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.Match(PocLangINT)
		}

	case PocLangSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
			p.Match(PocLangSTRING_LITERAL)
		}

	case PocLangTRUE_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(256)
			p.Match(PocLangTRUE_LIT)
		}

	case PocLangFALSE_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(257)
			p.Match(PocLangFALSE_LIT)
		}

	case PocLangIDENTIFIER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(258)
			p.Identifier_ref()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifier_refContext is an interface to support dynamic dispatch.
type IIdentifier_refContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifier_refContext differentiates from other interfaces.
	IsIdentifier_refContext()
}

type Identifier_refContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifier_refContext() *Identifier_refContext {
	var p = new(Identifier_refContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_identifier_ref
	return p
}

func (*Identifier_refContext) IsIdentifier_refContext() {}

func NewIdentifier_refContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Identifier_refContext {
	var p = new(Identifier_refContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_identifier_ref

	return p
}

func (s *Identifier_refContext) GetParser() antlr.Parser { return s.parser }

func (s *Identifier_refContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PocLangIDENTIFIER, 0)
}

func (s *Identifier_refContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Identifier_refContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Identifier_refContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterIdentifier_ref(s)
	}
}

func (s *Identifier_refContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitIdentifier_ref(s)
	}
}

func (s *Identifier_refContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitIdentifier_ref(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Identifier_ref() (localctx IIdentifier_refContext) {
	localctx = NewIdentifier_refContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PocLangRULE_identifier_ref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(PocLangIDENTIFIER)
	}

	return localctx
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_function_call
	return p
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PocLangIDENTIFIER, 0)
}

func (s *Function_callContext) BEGIN_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_GROUP, 0)
}

func (s *Function_callContext) END_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangEND_GROUP, 0)
}

func (s *Function_callContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Function_callContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Function_callContext) AllARG_SEP() []antlr.TerminalNode {
	return s.GetTokens(PocLangARG_SEP)
}

func (s *Function_callContext) ARG_SEP(i int) antlr.TerminalNode {
	return s.GetToken(PocLangARG_SEP, i)
}

func (s *Function_callContext) JAVA_BEGIN() antlr.TerminalNode {
	return s.GetToken(PocLangJAVA_BEGIN, 0)
}

func (s *Function_callContext) Java_call() IJava_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJava_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJava_callContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PocLangRULE_function_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(291)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PocLangIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(PocLangIDENTIFIER)
		}
		{
			p.SetState(264)
			p.Match(PocLangBEGIN_GROUP)
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PocLangINT)|(1<<PocLangTRUE_LIT)|(1<<PocLangFALSE_LIT)|(1<<PocLangBEGIN_GROUP)|(1<<PocLangIDENTIFIER)|(1<<PocLangJAVA_BEGIN))) != 0) || _la == PocLangSTRING_LITERAL {
			{
				p.SetState(265)
				p.expr(0)
			}
			p.SetState(270)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == PocLangARG_SEP {
				{
					p.SetState(266)
					p.Match(PocLangARG_SEP)
				}
				{
					p.SetState(267)
					p.expr(0)
				}

				p.SetState(272)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(275)
			p.Match(PocLangEND_GROUP)
		}

	case PocLangJAVA_BEGIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.Match(PocLangJAVA_BEGIN)
		}
		{
			p.SetState(277)
			p.Java_call()
		}
		{
			p.SetState(278)
			p.Match(PocLangBEGIN_GROUP)
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PocLangINT)|(1<<PocLangTRUE_LIT)|(1<<PocLangFALSE_LIT)|(1<<PocLangBEGIN_GROUP)|(1<<PocLangIDENTIFIER)|(1<<PocLangJAVA_BEGIN))) != 0) || _la == PocLangSTRING_LITERAL {
			{
				p.SetState(279)
				p.expr(0)
			}
			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == PocLangARG_SEP {
				{
					p.SetState(280)
					p.Match(PocLangARG_SEP)
				}
				{
					p.SetState(281)
					p.expr(0)
				}

				p.SetState(286)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(289)
			p.Match(PocLangEND_GROUP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJava_callContext is an interface to support dynamic dispatch.
type IJava_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJava_callContext differentiates from other interfaces.
	IsJava_callContext()
}

type Java_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJava_callContext() *Java_callContext {
	var p = new(Java_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_java_call
	return p
}

func (*Java_callContext) IsJava_callContext() {}

func NewJava_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Java_callContext {
	var p = new(Java_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_java_call

	return p
}

func (s *Java_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Java_callContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PocLangIDENTIFIER)
}

func (s *Java_callContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PocLangIDENTIFIER, i)
}

func (s *Java_callContext) AllJAVA_SEP() []antlr.TerminalNode {
	return s.GetTokens(PocLangJAVA_SEP)
}

func (s *Java_callContext) JAVA_SEP(i int) antlr.TerminalNode {
	return s.GetToken(PocLangJAVA_SEP, i)
}

func (s *Java_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Java_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Java_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterJava_call(s)
	}
}

func (s *Java_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitJava_call(s)
	}
}

func (s *Java_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitJava_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Java_call() (localctx IJava_callContext) {
	localctx = NewJava_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PocLangRULE_java_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PocLangIDENTIFIER {
		{
			p.SetState(293)
			p.Match(PocLangIDENTIFIER)
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PocLangJAVA_SEP {
			{
				p.SetState(294)
				p.Match(PocLangJAVA_SEP)
			}

		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReContext is an interface to support dynamic dispatch.
type IReContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReContext differentiates from other interfaces.
	IsReContext()
}

type ReContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReContext() *ReContext {
	var p = new(ReContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_re
	return p
}

func (*ReContext) IsReContext() {}

func NewReContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReContext {
	var p = new(ReContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_re

	return p
}

func (s *ReContext) GetParser() antlr.Parser { return s.parser }

func (s *ReContext) Simple_re() ISimple_reContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_reContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_reContext)
}

func (s *ReContext) Union_prime() IUnion_primeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnion_primeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnion_primeContext)
}

func (s *ReContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterRe(s)
	}
}

func (s *ReContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitRe(s)
	}
}

func (s *ReContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitRe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Re() (localctx IReContext) {
	localctx = NewReContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PocLangRULE_re)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Simple_re()
	}
	{
		p.SetState(303)
		p.Union_prime()
	}

	return localctx
}

// IUnion_primeContext is an interface to support dynamic dispatch.
type IUnion_primeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnion_primeContext differentiates from other interfaces.
	IsUnion_primeContext()
}

type Union_primeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnion_primeContext() *Union_primeContext {
	var p = new(Union_primeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_union_prime
	return p
}

func (*Union_primeContext) IsUnion_primeContext() {}

func NewUnion_primeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_primeContext {
	var p = new(Union_primeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_union_prime

	return p
}

func (s *Union_primeContext) GetParser() antlr.Parser { return s.parser }

func (s *Union_primeContext) ALTERNATION() antlr.TerminalNode {
	return s.GetToken(PocLangALTERNATION, 0)
}

func (s *Union_primeContext) Re() IReContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReContext)
}

func (s *Union_primeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_primeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Union_primeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterUnion_prime(s)
	}
}

func (s *Union_primeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitUnion_prime(s)
	}
}

func (s *Union_primeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitUnion_prime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Union_prime() (localctx IUnion_primeContext) {
	localctx = NewUnion_primeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PocLangRULE_union_prime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(308)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PocLangALTERNATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.Match(PocLangALTERNATION)
		}
		{
			p.SetState(306)
			p.Re()
		}

	case PocLangEND_RE_GROUP, PocLangRE_DELIMITER_CLOSE:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimple_reContext is an interface to support dynamic dispatch.
type ISimple_reContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_reContext differentiates from other interfaces.
	IsSimple_reContext()
}

type Simple_reContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_reContext() *Simple_reContext {
	var p = new(Simple_reContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_simple_re
	return p
}

func (*Simple_reContext) IsSimple_reContext() {}

func NewSimple_reContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_reContext {
	var p = new(Simple_reContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_simple_re

	return p
}

func (s *Simple_reContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_reContext) Basic_re() IBasic_reContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasic_reContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasic_reContext)
}

func (s *Simple_reContext) Concat_prime() IConcat_primeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcat_primeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcat_primeContext)
}

func (s *Simple_reContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_reContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_reContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterSimple_re(s)
	}
}

func (s *Simple_reContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitSimple_re(s)
	}
}

func (s *Simple_reContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitSimple_re(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Simple_re() (localctx ISimple_reContext) {
	localctx = NewSimple_reContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PocLangRULE_simple_re)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Basic_re()
	}
	{
		p.SetState(311)
		p.Concat_prime()
	}

	return localctx
}

// IConcat_primeContext is an interface to support dynamic dispatch.
type IConcat_primeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcat_primeContext differentiates from other interfaces.
	IsConcat_primeContext()
}

type Concat_primeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcat_primeContext() *Concat_primeContext {
	var p = new(Concat_primeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_concat_prime
	return p
}

func (*Concat_primeContext) IsConcat_primeContext() {}

func NewConcat_primeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Concat_primeContext {
	var p = new(Concat_primeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_concat_prime

	return p
}

func (s *Concat_primeContext) GetParser() antlr.Parser { return s.parser }

func (s *Concat_primeContext) Simple_re() ISimple_reContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_reContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_reContext)
}

func (s *Concat_primeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Concat_primeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Concat_primeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterConcat_prime(s)
	}
}

func (s *Concat_primeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitConcat_prime(s)
	}
}

func (s *Concat_primeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitConcat_prime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Concat_prime() (localctx IConcat_primeContext) {
	localctx = NewConcat_primeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PocLangRULE_concat_prime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(315)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PocLangBEGIN_RE_RANGE, PocLangCHARACTER, PocLangBEGIN_RE_GROUP, PocLangMINUS, PocLangDOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.Simple_re()
		}

	case PocLangEND_RE_GROUP, PocLangALTERNATION, PocLangRE_DELIMITER_CLOSE:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBasic_reContext is an interface to support dynamic dispatch.
type IBasic_reContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasic_reContext differentiates from other interfaces.
	IsBasic_reContext()
}

type Basic_reContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasic_reContext() *Basic_reContext {
	var p = new(Basic_reContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_basic_re
	return p
}

func (*Basic_reContext) IsBasic_reContext() {}

func NewBasic_reContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Basic_reContext {
	var p = new(Basic_reContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_basic_re

	return p
}

func (s *Basic_reContext) GetParser() antlr.Parser { return s.parser }

func (s *Basic_reContext) Kleene_star() IKleene_starContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKleene_starContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKleene_starContext)
}

func (s *Basic_reContext) Plus() IPlusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusContext)
}

func (s *Basic_reContext) Elementary_re() IElementary_reContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementary_reContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementary_reContext)
}

func (s *Basic_reContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Basic_reContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Basic_reContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterBasic_re(s)
	}
}

func (s *Basic_reContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitBasic_re(s)
	}
}

func (s *Basic_reContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitBasic_re(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Basic_re() (localctx IBasic_reContext) {
	localctx = NewBasic_reContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PocLangRULE_basic_re)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Kleene_star()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(318)
			p.Plus()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.Elementary_re()
		}

	}

	return localctx
}

// IKleene_starContext is an interface to support dynamic dispatch.
type IKleene_starContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKleene_starContext differentiates from other interfaces.
	IsKleene_starContext()
}

type Kleene_starContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKleene_starContext() *Kleene_starContext {
	var p = new(Kleene_starContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_kleene_star
	return p
}

func (*Kleene_starContext) IsKleene_starContext() {}

func NewKleene_starContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kleene_starContext {
	var p = new(Kleene_starContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_kleene_star

	return p
}

func (s *Kleene_starContext) GetParser() antlr.Parser { return s.parser }

func (s *Kleene_starContext) Elementary_re() IElementary_reContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementary_reContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementary_reContext)
}

func (s *Kleene_starContext) STAR() antlr.TerminalNode {
	return s.GetToken(PocLangSTAR, 0)
}

func (s *Kleene_starContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kleene_starContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kleene_starContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterKleene_star(s)
	}
}

func (s *Kleene_starContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitKleene_star(s)
	}
}

func (s *Kleene_starContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitKleene_star(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Kleene_star() (localctx IKleene_starContext) {
	localctx = NewKleene_starContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PocLangRULE_kleene_star)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Elementary_re()
	}
	{
		p.SetState(323)
		p.Match(PocLangSTAR)
	}

	return localctx
}

// IPlusContext is an interface to support dynamic dispatch.
type IPlusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlusContext differentiates from other interfaces.
	IsPlusContext()
}

type PlusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlusContext() *PlusContext {
	var p = new(PlusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_plus
	return p
}

func (*PlusContext) IsPlusContext() {}

func NewPlusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlusContext {
	var p = new(PlusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_plus

	return p
}

func (s *PlusContext) GetParser() antlr.Parser { return s.parser }

func (s *PlusContext) Elementary_re() IElementary_reContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementary_reContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementary_reContext)
}

func (s *PlusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PocLangPLUS, 0)
}

func (s *PlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterPlus(s)
	}
}

func (s *PlusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitPlus(s)
	}
}

func (s *PlusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitPlus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Plus() (localctx IPlusContext) {
	localctx = NewPlusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PocLangRULE_plus)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Elementary_re()
	}
	{
		p.SetState(326)
		p.Match(PocLangPLUS)
	}

	return localctx
}

// IElementary_reContext is an interface to support dynamic dispatch.
type IElementary_reContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementary_reContext differentiates from other interfaces.
	IsElementary_reContext()
}

type Elementary_reContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementary_reContext() *Elementary_reContext {
	var p = new(Elementary_reContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_elementary_re
	return p
}

func (*Elementary_reContext) IsElementary_reContext() {}

func NewElementary_reContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elementary_reContext {
	var p = new(Elementary_reContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_elementary_re

	return p
}

func (s *Elementary_reContext) GetParser() antlr.Parser { return s.parser }

func (s *Elementary_reContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *Elementary_reContext) DOT() antlr.TerminalNode {
	return s.GetToken(PocLangDOT, 0)
}

func (s *Elementary_reContext) Character() ICharacterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterContext)
}

func (s *Elementary_reContext) Range_re() IRange_reContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_reContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_reContext)
}

func (s *Elementary_reContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elementary_reContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elementary_reContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterElementary_re(s)
	}
}

func (s *Elementary_reContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitElementary_re(s)
	}
}

func (s *Elementary_reContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitElementary_re(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Elementary_re() (localctx IElementary_reContext) {
	localctx = NewElementary_reContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PocLangRULE_elementary_re)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(332)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PocLangBEGIN_RE_GROUP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Group()
		}

	case PocLangDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Match(PocLangDOT)
		}

	case PocLangCHARACTER, PocLangMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(330)
			p.Character()
		}

	case PocLangBEGIN_RE_RANGE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(331)
			p.Range_re()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) BEGIN_RE_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_RE_GROUP, 0)
}

func (s *GroupContext) Re() IReContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReContext)
}

func (s *GroupContext) END_RE_GROUP() antlr.TerminalNode {
	return s.GetToken(PocLangEND_RE_GROUP, 0)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PocLangRULE_group)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(PocLangBEGIN_RE_GROUP)
	}
	{
		p.SetState(335)
		p.Re()
	}
	{
		p.SetState(336)
		p.Match(PocLangEND_RE_GROUP)
	}

	return localctx
}

// IRange_reContext is an interface to support dynamic dispatch.
type IRange_reContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_reContext differentiates from other interfaces.
	IsRange_reContext()
}

type Range_reContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_reContext() *Range_reContext {
	var p = new(Range_reContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_range_re
	return p
}

func (*Range_reContext) IsRange_reContext() {}

func NewRange_reContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_reContext {
	var p = new(Range_reContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_range_re

	return p
}

func (s *Range_reContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_reContext) Positive_range() IPositive_rangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositive_rangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPositive_rangeContext)
}

func (s *Range_reContext) Negative_range() INegative_rangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INegative_rangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INegative_rangeContext)
}

func (s *Range_reContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_reContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_reContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterRange_re(s)
	}
}

func (s *Range_reContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitRange_re(s)
	}
}

func (s *Range_reContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitRange_re(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Range_re() (localctx IRange_reContext) {
	localctx = NewRange_reContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PocLangRULE_range_re)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Positive_range()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)
			p.Negative_range()
		}

	}

	return localctx
}

// IPositive_rangeContext is an interface to support dynamic dispatch.
type IPositive_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPositive_rangeContext differentiates from other interfaces.
	IsPositive_rangeContext()
}

type Positive_rangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPositive_rangeContext() *Positive_rangeContext {
	var p = new(Positive_rangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_positive_range
	return p
}

func (*Positive_rangeContext) IsPositive_rangeContext() {}

func NewPositive_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Positive_rangeContext {
	var p = new(Positive_rangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_positive_range

	return p
}

func (s *Positive_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Positive_rangeContext) BEGIN_RE_RANGE() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_RE_RANGE, 0)
}

func (s *Positive_rangeContext) Range_items() IRange_itemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_itemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_itemsContext)
}

func (s *Positive_rangeContext) RANGE_TERMINATE() antlr.TerminalNode {
	return s.GetToken(PocLangRANGE_TERMINATE, 0)
}

func (s *Positive_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Positive_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Positive_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterPositive_range(s)
	}
}

func (s *Positive_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitPositive_range(s)
	}
}

func (s *Positive_rangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitPositive_range(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Positive_range() (localctx IPositive_rangeContext) {
	localctx = NewPositive_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PocLangRULE_positive_range)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(PocLangBEGIN_RE_RANGE)
	}
	{
		p.SetState(343)
		p.Range_items()
	}
	{
		p.SetState(344)
		p.Match(PocLangRANGE_TERMINATE)
	}

	return localctx
}

// INegative_rangeContext is an interface to support dynamic dispatch.
type INegative_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNegative_rangeContext differentiates from other interfaces.
	IsNegative_rangeContext()
}

type Negative_rangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNegative_rangeContext() *Negative_rangeContext {
	var p = new(Negative_rangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_negative_range
	return p
}

func (*Negative_rangeContext) IsNegative_rangeContext() {}

func NewNegative_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Negative_rangeContext {
	var p = new(Negative_rangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_negative_range

	return p
}

func (s *Negative_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Negative_rangeContext) BEGIN_RE_RANGE() antlr.TerminalNode {
	return s.GetToken(PocLangBEGIN_RE_RANGE, 0)
}

func (s *Negative_rangeContext) RANGE_NEGATE() antlr.TerminalNode {
	return s.GetToken(PocLangRANGE_NEGATE, 0)
}

func (s *Negative_rangeContext) Range_items() IRange_itemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_itemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_itemsContext)
}

func (s *Negative_rangeContext) RANGE_TERMINATE() antlr.TerminalNode {
	return s.GetToken(PocLangRANGE_TERMINATE, 0)
}

func (s *Negative_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Negative_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Negative_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterNegative_range(s)
	}
}

func (s *Negative_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitNegative_range(s)
	}
}

func (s *Negative_rangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitNegative_range(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Negative_range() (localctx INegative_rangeContext) {
	localctx = NewNegative_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PocLangRULE_negative_range)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Match(PocLangBEGIN_RE_RANGE)
	}
	{
		p.SetState(347)
		p.Match(PocLangRANGE_NEGATE)
	}
	{
		p.SetState(348)
		p.Range_items()
	}
	{
		p.SetState(349)
		p.Match(PocLangRANGE_TERMINATE)
	}

	return localctx
}

// ILax_characterContext is an interface to support dynamic dispatch.
type ILax_characterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLax_characterContext differentiates from other interfaces.
	IsLax_characterContext()
}

type Lax_characterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLax_characterContext() *Lax_characterContext {
	var p = new(Lax_characterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_lax_character
	return p
}

func (*Lax_characterContext) IsLax_characterContext() {}

func NewLax_characterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lax_characterContext {
	var p = new(Lax_characterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_lax_character

	return p
}

func (s *Lax_characterContext) GetParser() antlr.Parser { return s.parser }

func (s *Lax_characterContext) RANGE_CHARACTER() antlr.TerminalNode {
	return s.GetToken(PocLangRANGE_CHARACTER, 0)
}

func (s *Lax_characterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lax_characterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lax_characterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterLax_character(s)
	}
}

func (s *Lax_characterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitLax_character(s)
	}
}

func (s *Lax_characterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitLax_character(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Lax_character() (localctx ILax_characterContext) {
	localctx = NewLax_characterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PocLangRULE_lax_character)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(PocLangRANGE_CHARACTER)
	}

	return localctx
}

// ICharacterContext is an interface to support dynamic dispatch.
type ICharacterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterContext differentiates from other interfaces.
	IsCharacterContext()
}

type CharacterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterContext() *CharacterContext {
	var p = new(CharacterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_character
	return p
}

func (*CharacterContext) IsCharacterContext() {}

func NewCharacterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterContext {
	var p = new(CharacterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_character

	return p
}

func (s *CharacterContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(PocLangCHARACTER, 0)
}

func (s *CharacterContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PocLangMINUS, 0)
}

func (s *CharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterCharacter(s)
	}
}

func (s *CharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitCharacter(s)
	}
}

func (s *CharacterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitCharacter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Character() (localctx ICharacterContext) {
	localctx = NewCharacterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PocLangRULE_character)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PocLangCHARACTER || _la == PocLangMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRange_itemsContext is an interface to support dynamic dispatch.
type IRange_itemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_itemsContext differentiates from other interfaces.
	IsRange_itemsContext()
}

type Range_itemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_itemsContext() *Range_itemsContext {
	var p = new(Range_itemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_range_items
	return p
}

func (*Range_itemsContext) IsRange_itemsContext() {}

func NewRange_itemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_itemsContext {
	var p = new(Range_itemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_range_items

	return p
}

func (s *Range_itemsContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_itemsContext) Range_item() IRange_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_itemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_itemContext)
}

func (s *Range_itemsContext) Range_items() IRange_itemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_itemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_itemsContext)
}

func (s *Range_itemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_itemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_itemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterRange_items(s)
	}
}

func (s *Range_itemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitRange_items(s)
	}
}

func (s *Range_itemsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitRange_items(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Range_items() (localctx IRange_itemsContext) {
	localctx = NewRange_itemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PocLangRULE_range_items)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.Range_item()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.Range_item()
		}
		{
			p.SetState(357)
			p.Range_items()
		}

	}

	return localctx
}

// IRange_itemContext is an interface to support dynamic dispatch.
type IRange_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_itemContext differentiates from other interfaces.
	IsRange_itemContext()
}

type Range_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_itemContext() *Range_itemContext {
	var p = new(Range_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PocLangRULE_range_item
	return p
}

func (*Range_itemContext) IsRange_itemContext() {}

func NewRange_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_itemContext {
	var p = new(Range_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PocLangRULE_range_item

	return p
}

func (s *Range_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_itemContext) AllLax_character() []ILax_characterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILax_characterContext)(nil)).Elem())
	var tst = make([]ILax_characterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILax_characterContext)
		}
	}

	return tst
}

func (s *Range_itemContext) Lax_character(i int) ILax_characterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILax_characterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILax_characterContext)
}

func (s *Range_itemContext) RANGE_SEPARATOR() antlr.TerminalNode {
	return s.GetToken(PocLangRANGE_SEPARATOR, 0)
}

func (s *Range_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.EnterRange_item(s)
	}
}

func (s *Range_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PocLangListener); ok {
		listenerT.ExitRange_item(s)
	}
}

func (s *Range_itemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PocLangVisitor:
		return t.VisitRange_item(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PocLang) Range_item() (localctx IRange_itemContext) {
	localctx = NewRange_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PocLangRULE_range_item)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)
			p.Lax_character()
		}
		{
			p.SetState(362)
			p.Match(PocLangRANGE_SEPARATOR)
		}
		{
			p.SetState(363)
			p.Lax_character()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.Lax_character()
		}

	}

	return localctx
}

func (p *PocLang) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PocLang) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
